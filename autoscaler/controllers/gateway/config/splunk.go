package config

import (
	"errors"
	"fmt"

	commonconf "github.com/odigos-io/odigos/autoscaler/controllers/common"
	"github.com/odigos-io/odigos/common"
)

const (
	splunkRealm = "SPLUNK_REALM"
)

type Splunk struct{}

func (s *Splunk) DestType() common.DestinationType {
	return common.SplunkDestinationType
}

func (s *Splunk) ModifyConfig(dest common.ExporterConfigurer, currentConfig *commonconf.Config) error {
	realm, exists := dest.GetConfig()[splunkRealm]
	if !exists {
		return errors.New("Splunk realm not specified, gateway will not be configured for Splunk")
	}

	if isTracingEnabled(dest) {
		exporterName := "sapm/" + dest.GetName()
		currentConfig.Exporters[exporterName] = commonconf.GenericMap{
			"access_token": "${SPLUNK_ACCESS_TOKEN}",
			"endpoint":     fmt.Sprintf("https://ingest.%s.signalfx.com/v2/trace", realm),
		}

		tracesPipelineName := "traces/splunk-" + dest.GetName()
		currentConfig.Service.Pipelines[tracesPipelineName] = commonconf.Pipeline{
			Exporters: []string{exporterName},
		}
	}

	return nil
}
