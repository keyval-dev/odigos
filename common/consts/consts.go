package consts

import (
	"errors"
)

const (
	CurrentNamespaceEnvVar         = "CURRENT_NS"
	DefaultNamespace               = "odigos-system"
	DefaultOdigosConfigurationName = "odigos-config"
	OTLPPort                       = 4317
	OTLPHttpPort                   = 4318
	OdigosInstrumentationLabel     = "odigos-instrumentation"
	InstrumentationEnabled         = "enabled"
	InstrumentationDisabled        = "disabled"
	OdigosReportedNameAnnotation   = "odigos.io/reported-name"
	EbpfInstrumentationAnnotation  = "instrumentation.odigos.io/ebpf" // deprecated.
	// Used to store the original value of the environment variable in the pod manifest.
	// This is used to restore the original value when an instrumentation is removed
	// or odigos is uninstalled.
	// Should only be used for environment variables that are modified by odigos.
	ManifestEnvOriginalValAnnotation = "odigos.io/manifest-env-original-val"
)

var (
	PodsNotFoundErr = errors.New("could not find a ready pod")
)

var (
	SystemNamespaces = []string{DefaultNamespace, "kube-system", "local-path-storage", "istio-system", "linkerd", "kube-node-lease"}
)
