package instrumentation_ebpf

import (
	"context"
	"errors"

	odigosv1 "github.com/keyval-dev/odigos/api/odigos/v1alpha1"
	"github.com/keyval-dev/odigos/common"
	"github.com/keyval-dev/odigos/common/utils"
	"github.com/keyval-dev/odigos/odiglet/pkg/ebpf"
	"github.com/keyval-dev/odigos/odiglet/pkg/kube"
	"github.com/keyval-dev/odigos/odiglet/pkg/process"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// PodWorkload represents the higher-level controller managing a specific Pod within a Kubernetes cluster.
// It contains essential details about the controller such as its Name, Namespace, and Kind.
// 'Kind' refers to the type of controller, which can be a Deployment, StatefulSet, or DaemonSet.
// This struct is useful for identifying and interacting with the overarching entity
// that governs the lifecycle and behavior of a Pod, especially in contexts where
// understanding the relationship between a Pod and its controlling workload is crucial.
type PodWorkload struct {
	Name      string
	Namespace string
	Kind      string
}

func ApplyEbpfToPodWorkload(ctx context.Context, kubeClient client.Client, directors map[common.ProgrammingLanguage]ebpf.Director, podWorkload *PodWorkload) error {
	logger := log.FromContext(ctx)
	ebpfInstrumented, matchLabels, err := isEbpfInstrumented(ctx, kubeClient, podWorkload)
	if err != nil {
		logger.Error(err, "error checking if pod is ebpf instrumented")
		return err
	}
	if !ebpfInstrumented {
		cleanupEbpf(directors, types.NamespacedName{
			Namespace: podWorkload.Namespace,
			Name:      podWorkload.Name,
		})
		return nil
	}

	pods, err := kube.GetRunningPods(ctx, matchLabels, podWorkload.Namespace, kubeClient)
	if err != nil {
		logger.Error(err, "error fetching running pods")
		return err
	}

	if len(pods) == 0 {
		return nil
	}

	runtimeDetails, err := getRuntimeDetails(ctx, kubeClient, podWorkload)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// Probably shutdown in progress, cleanup will be done as soon as the pod object is deleted
			return nil
		}
		return err
	}

	for _, pod := range pods {
		err = instrumentPodWithEbpf(ctx, &pod, directors, runtimeDetails)
		if err != nil {
			logger.Error(err, "error instrumenting pod")
			return err
		}
	}

	return nil
}

func cleanupEbpf(directors map[common.ProgrammingLanguage]ebpf.Director, name types.NamespacedName) {
	// cleanup using all available directors
	// the Cleanup method is idempotent, so no harm in calling it multiple times
	for _, director := range directors {
		director.Cleanup(name)
	}
}

func instrumentPodWithEbpf(ctx context.Context, pod *corev1.Pod, directors map[common.ProgrammingLanguage]ebpf.Director, runtimeDetails *odigosv1.InstrumentedApplication) error {
	logger := log.FromContext(ctx)
	podUid := string(pod.UID)
	for _, container := range runtimeDetails.Spec.Languages {

		director := directors[container.Language]
		if director == nil {
			return errors.New("no director found for language " + string(container.Language))
		}

		appName := container.ContainerName
		if len(runtimeDetails.Spec.Languages) == 1 && len(runtimeDetails.OwnerReferences) > 0 {
			appName = runtimeDetails.OwnerReferences[0].Name
		}

		details, err := process.FindAllInContainer(podUid, container.ContainerName)
		if err != nil {
			logger.Error(err, "error finding processes")
			return err
		}

		for _, d := range details {
			err = director.Instrument(d.ProcessID, types.NamespacedName{
				Namespace: pod.Namespace,
				Name:      pod.Name,
			}, appName)

			if err != nil {
				logger.Error(err, "error instrumenting process", "pid", d.ProcessID)
				return err
			}
		}
	}
	return nil
}

func isEbpfInstrumented(ctx context.Context, kubeClient client.Client, podWorkload *PodWorkload) (bool, map[string]string, error) {
	// TODO: this is better done with a dynamic client
	switch podWorkload.Kind {
	case "Deployment":
		var dep appsv1.Deployment
		err := kubeClient.Get(ctx, client.ObjectKey{
			Namespace: podWorkload.Namespace,
			Name:      podWorkload.Name,
		}, &dep)
		return hasEbpfInstrumentationAnnotation(&dep), dep.Spec.Selector.MatchLabels, err
	case "DaemonSet":
		var ds appsv1.DaemonSet
		err := kubeClient.Get(ctx, client.ObjectKey{
			Namespace: podWorkload.Namespace,
			Name:      podWorkload.Name,
		}, &ds)
		return hasEbpfInstrumentationAnnotation(&ds), ds.Spec.Selector.MatchLabels, err
	case "StatefulSet":
		var sts appsv1.StatefulSet
		err := kubeClient.Get(ctx, client.ObjectKey{
			Namespace: podWorkload.Namespace,
			Name:      podWorkload.Name,
		}, &sts)
		return hasEbpfInstrumentationAnnotation(&sts), sts.Spec.Selector.MatchLabels, err
	default:
		return false, nil, errors.New("unknown pod workload kind")
	}
}

func getRuntimeDetails(ctx context.Context, kubeClient client.Client, podWorkload *PodWorkload) (*odigosv1.InstrumentedApplication, error) {
	instrumentedApplicationName := utils.GetRuntimeObjectName(podWorkload.Name, podWorkload.Kind)

	var runtimeDetails odigosv1.InstrumentedApplication
	err := kubeClient.Get(ctx, client.ObjectKey{
		Namespace: podWorkload.Namespace,
		Name:      instrumentedApplicationName,
	}, &runtimeDetails)
	if err != nil {
		return nil, err
	}

	return &runtimeDetails, nil
}
