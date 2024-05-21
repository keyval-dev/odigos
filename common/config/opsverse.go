package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/odigos-io/odigos/common"
)

const (
	opsverseLogsUrl    = "OPSVERSE_LOGS_URL"
	opsverseMetricsUrl = "OPSVERSE_METRICS_URL"
	opsverseTracesUrl  = "OPSVERSE_TRACES_URL"
	opsverseUserName   = "OPSVERSE_USERNAME"
)

type OpsVerse struct{}

func (g *OpsVerse) DestType() common.DestinationType {
	return common.OpsVerseDestinationType
}

func (g *OpsVerse) ModifyConfig(dest ExporterConfigurer, currentConfig *Config) error {
	var err error
	if isMetricsEnabled(dest) {
		e := g.isMetricsVarsExists(dest)
		if e != nil {
			err = errors.Join(err, e)
		} else {
			url := fmt.Sprintf("%s/api/v1/write", dest.GetConfig()[opsverseMetricsUrl])
			rwExporterName := "prometheusremotewrite/opsverse-" + dest.GetName()
			currentConfig.Exporters[rwExporterName] = GenericMap{
				"endpoint": url,
				"headers": GenericMap{
					"Authorization": fmt.Sprintf("Basic %s", "${OPSVERSE_AUTH_TOKEN}"),
				},
			}

			metricsPipelineName := "metrics/opsverse-" + dest.GetName()
			currentConfig.Service.Pipelines[metricsPipelineName] = Pipeline{
				Exporters: []string{rwExporterName},
			}
		}
	}

	if isTracingEnabled(dest) {
		e := g.isTracingVarsExists(dest)
		if e != nil {
			err = errors.Join(err, e)
		} else {
			url := dest.GetConfig()[opsverseTracesUrl]
			url = strings.TrimPrefix(url, "http://")
			url = strings.TrimPrefix(url, "https://")
			url = fmt.Sprintf("%s:443", url)
			exporterName := "otlp/opsverse-" + dest.GetName()
			currentConfig.Exporters[exporterName] = GenericMap{
				"endpoint": url,
				"headers": GenericMap{
					"authorization": "Basic ${OPSVERSE_AUTH_TOKEN}",
				},
			}

			currentConfig.Service.Pipelines["traces/opsverse"] = Pipeline{
				Exporters: []string{exporterName},
			}
		}
	}

	if isLoggingEnabled(dest) {
		e := g.isLogsVarsExists(dest)
		if e != nil {
			err = errors.Join(err, e)
		} else {
			url := fmt.Sprintf("%s/loki/api/v1/push", dest.GetConfig()[opsverseLogsUrl])

			lokiExporterName := "loki/opsverse-" + dest.GetName()
			currentConfig.Exporters[lokiExporterName] = GenericMap{
				"endpoint": url,
				"headers": GenericMap{
					"Authorization": fmt.Sprintf("Basic %s", "${OPSVERSE_AUTH_TOKEN}"),
				},
				"labels": GenericMap{
					"attributes": GenericMap{
						"k8s.container.name": "k8s_container_name",
						"k8s.pod.name":       "k8s_pod_name",
						"k8s.namespace.name": "k8s_namespace_name",
					},
				},
			}

			logsPipelineName := "logs/opsverse-" + dest.GetName()
			currentConfig.Service.Pipelines[logsPipelineName] = Pipeline{
				Exporters: []string{lokiExporterName},
			}
		}
	}

	return err
}

func (g *OpsVerse) isTracingVarsExists(dest ExporterConfigurer) error {
	_, exists := dest.GetConfig()[opsverseTracesUrl]
	if !exists {
		return errors.New("OpsVerse OTLP tracing endpoint not specified, gateway will not be configured for tracing")
	}

	_, exists = dest.GetConfig()[opsverseUserName]
	if !exists {
		return errors.New("OpsVerse user not specified, gateway will not be configured for traces")
	}

	return nil
}

func (g *OpsVerse) isLogsVarsExists(dest ExporterConfigurer) error {
	_, exists := dest.GetConfig()[opsverseLogsUrl]
	if !exists {
		return errors.New("OpsVerse logs endpoint not specified, gateway will not be configured for logs")
	}

	_, exists = dest.GetConfig()[opsverseUserName]
	if !exists {
		return errors.New("OpsVerse user not specified, gateway will not be configured for logs")
	}

	return nil
}

func (g *OpsVerse) isMetricsVarsExists(dest ExporterConfigurer) error {
	_, exists := dest.GetConfig()[opsverseMetricsUrl]
	if !exists {
		return errors.New("OpsVerse metrics endpoint not specified, gateway will not be configured for metrics")
	}

	_, exists = dest.GetConfig()[opsverseUserName]
	if !exists {
		return errors.New("OpsVerse user not specified, gateway will not be configured for metrics")
	}

	return nil
}
