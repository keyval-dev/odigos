package resources

import (
	odigosv1 "github.com/keyval-dev/odigos/api/odigos/v1alpha1"
	"github.com/keyval-dev/odigos/cli/pkg/kube"
)

func CreateResourceManagers(client *kube.Client, odigosNs string, version string, isOdigosCloud bool, config *odigosv1.OdigosConfigurationSpec) []ResourceManager {

	// Note - the order is important.
	// If resource A depends on resource B, then A must be installed after B.
	resourceManager := []ResourceManager{
		NewOdigosDeploymentResourceManager(client, odigosNs, version),
		NewOdigosConfigResourceManager(client, odigosNs, version, config),
		NewOwnTelemetryResourceManager(client, odigosNs, version, isOdigosCloud),
		NewDataCollectionResourceManager(client, odigosNs, version, config),
		NewInstrumentorResourceManager(client, odigosNs, version, config),
		NewSchedulerResourceManager(client, odigosNs, version, config),
		NewOdigletResourceManager(client, odigosNs, version, config),
		NewAutoScalerResourceManager(client, odigosNs, version, config),
	}

	if isOdigosCloud {
		resourceManager = append(resourceManager, NewKeyvalProxyResourceManager(client, odigosNs, version, config))
	}

	return resourceManager
}
