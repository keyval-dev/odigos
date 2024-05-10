package endpoints

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odigos-io/odigos/frontend/kube"
	"golang.org/x/sync/errgroup"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GetApplicationsInNamespaceRequest struct {
	Namespace string `uri:"namespace" binding:"required"`
}

type GetApplicationsInNamespaceResponse struct {
	Applications []GetApplicationItemInNamespace `json:"applications"`
}

type WorkloadKind string

const (
	WorkloadKindDeployment  WorkloadKind = "Deployment"
	WorkloadKindStatefulSet WorkloadKind = "StatefulSet"
	WorkloadKindDaemonSet   WorkloadKind = "DaemonSet"
)

type GetApplicationItemInNamespace struct {
	Name                      string       `json:"name"`
	Kind                      WorkloadKind `json:"kind"`
	Instances                 int          `json:"instances"`
	AppInstrumentationLabeled *bool        `json:"app_instrumentation_labeled"`
	NsInstrumentationLabeled  *bool        `json:"ns_instrumentation_labeled"`
	InstrumentationEffective  bool         `json:"instrumentation_effective"`
}

type GetApplicationItem struct {
	// nameSpace is used when querying all the namespaces, the response can be grouped/filtered by namespace
	nameSpace string
	nsItem GetApplicationItemInNamespace
}

func GetApplicationsInNamespace(c *gin.Context) {
	var request GetApplicationsInNamespaceRequest
	if err := c.ShouldBindUri(&request); err != nil {
		returnError(c, err)
		return
	}

	ctx := c.Request.Context()
	namespace, err := kube.DefaultClient.CoreV1().Namespaces().Get(ctx, request.Namespace, metav1.GetOptions{})
	if err != nil {
		returnError(c, err)
		return
	}

	items, err := getApplicationsInNamespace(ctx, namespace.Name, map[string]*bool{namespace.Name: isObjectLabeledForInstrumentation(namespace.ObjectMeta)})
	if err != nil {
		returnError(c, err)
		return
	}

	apps := make([]GetApplicationItemInNamespace, len(items))
	for i, item := range items {
		apps[i] = item.nsItem
	}

	c.JSON(http.StatusOK, GetApplicationsInNamespaceResponse{
		Applications: apps,
	})
}

// getApplicationsInNamespace returns all applications in the namespace and their instrumentation status.
// nsName can be an empty string to get applications in all namespaces.
// nsInstrumentedMap is a map of namespace name to a boolean pointer indicating if the namespace is instrumented.
func getApplicationsInNamespace(ctx context.Context, nsName string, nsInstrumentedMap map[string]*bool) ([]GetApplicationItem, error) {
	g, ctx := errgroup.WithContext(ctx)
	var (
		deps []GetApplicationItem
		ss   []GetApplicationItem
		dss  []GetApplicationItem
	)

	g.Go(func() error {
		var err error
		deps, err = getDeployments(nsName, ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ss, err = getStatefulSets(nsName, ctx)
		return err
	})

	g.Go(func() error {
		var err error
		dss, err = getDaemonSets(nsName, ctx)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	items := make([]GetApplicationItem, len(deps)+len(ss)+len(dss))
	copy(items, deps)
	copy(items[len(deps):], ss)
	copy(items[len(deps)+len(ss):], dss)

	for i := range items {
		item := &items[i]
		// check if the entire namespace is instrumented
		// as it affects the applications in the namespace
		// which use this label to determine if they should be instrumented
		nsInstrumentationLabeled := nsInstrumentedMap[item.nameSpace]
		item.nsItem.NsInstrumentationLabeled = nsInstrumentationLabeled
		appInstrumented := (item.nsItem.AppInstrumentationLabeled != nil && *item.nsItem.AppInstrumentationLabeled)
		appInstrumentationInherited := item.nsItem.AppInstrumentationLabeled == nil
		nsInstrumented := (nsInstrumentationLabeled != nil && *nsInstrumentationLabeled)
		item.nsItem.InstrumentationEffective = appInstrumented || (appInstrumentationInherited && nsInstrumented)
	}

	return items, nil
}

func getDeployments(namespace string, ctx context.Context) ([]GetApplicationItem, error) {
	deps, err := kube.DefaultClient.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	response := make([]GetApplicationItem, len(deps.Items))
	for i, dep := range deps.Items {
		appInstrumentationLabeled := isObjectLabeledForInstrumentation(dep.ObjectMeta)
		response[i] = GetApplicationItem{
			nameSpace:                 dep.Namespace,
			nsItem: GetApplicationItemInNamespace {
				Name:                      dep.Name,
				Kind:                      WorkloadKindDeployment,
				Instances:                 int(dep.Status.AvailableReplicas),
				AppInstrumentationLabeled: appInstrumentationLabeled,
			},
		}
	}

	return response, nil
}

func getStatefulSets(namespace string, ctx context.Context) ([]GetApplicationItem, error) {
	ss, err := kube.DefaultClient.AppsV1().StatefulSets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	response := make([]GetApplicationItem, len(ss.Items))
	for i, s := range ss.Items {
		response[i] = GetApplicationItem{
			nameSpace: s.Namespace,
			nsItem: GetApplicationItemInNamespace {
				Name:      s.Name,
				Kind:      WorkloadKindStatefulSet,
				Instances: int(s.Status.ReadyReplicas),
			},
		}
	}

	return response, nil
}

func getDaemonSets(namespace string, ctx context.Context) ([]GetApplicationItem, error) {
	dss, err := kube.DefaultClient.AppsV1().DaemonSets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	response := make([]GetApplicationItem, len(dss.Items))
	for i, ds := range dss.Items {
		response[i] = GetApplicationItem{
			nameSpace: ds.Namespace,
			nsItem: GetApplicationItemInNamespace {
				Name:      ds.Name,
				Kind:      WorkloadKindDaemonSet,
				Instances: int(ds.Status.NumberReady),
			},
		}
	}

	return response, nil
}
