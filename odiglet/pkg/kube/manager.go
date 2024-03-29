package kube

import (
	"context"

	"github.com/keyval-dev/odigos/common"
	"github.com/keyval-dev/odigos/odiglet/pkg/ebpf"
	"github.com/keyval-dev/odigos/odiglet/pkg/kube/instrumentation_ebpf"
	"github.com/keyval-dev/odigos/odiglet/pkg/kube/runtime_details"
	"github.com/keyval-dev/odigos/odiglet/pkg/log"
	ctrl "sigs.k8s.io/controller-runtime"

	odigosv1 "github.com/keyval-dev/odigos/api/odigos/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(odigosv1.AddToScheme(scheme))
}

func CreateManager() (ctrl.Manager, error) {
	log.Logger.V(0).Info("Starting reconcileres for runtime details")
	ctrl.SetLogger(log.Logger)
	return manager.New(config.GetConfigOrDie(), manager.Options{
		Scheme: scheme,
		Metrics: metricsserver.Options{
			BindAddress: "0",
		},
	})
}

func StartManager(ctx context.Context, mgr ctrl.Manager) error {
	go func() {
		err := mgr.Start(ctx)
		if err != nil {
			log.Logger.Error(err, "error starting kube manager")
		}
	}()

	return nil
}

func SetupWithManager(mgr ctrl.Manager, ebpfDirectors map[common.ProgrammingLanguage]ebpf.Director) error {
	err := runtime_details.SetupWithManager(mgr)
	if err != nil {
		return err
	}

	err = instrumentation_ebpf.SetupWithManager(mgr, ebpfDirectors)
	if err != nil {
		return err
	}

	return nil
}
