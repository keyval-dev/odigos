package config

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/odigos-io/odigos/common"
)

const (
	qrynHost      = "QRYN_URL"
	qrynAPIKey    = "QRYN_API_KEY"
	qrynAPISecret = "${QRYN_API_SECRET}"
)

type Qryn struct{}

func (g *Qryn) DestType() common.DestinationType {
	return common.QrynDestinationType
}

func (g *Qryn) ModifyConfig(dest ExporterConfigurer, currentConfig *Config) error {
	if !g.requiredVarsExists(dest) {
		return errors.New("Qryn config is missing required variables")
	}
	apiKey, apiSecret := g.authData(dest)
	if apiKey == "" || apiSecret == "" {
		return errors.New("Qryn API key or secret not set")
	}

	baseURL, err := parseURL(dest.GetConfig()[qrynHost], apiKey, apiSecret)
	if err != nil {
		return errors.New("Qryn API host is not a valid")
	}

	if isMetricsEnabled(dest) {
		rwExporterName := "prometheusremotewrite/qryn-" + dest.GetID()
		currentConfig.Exporters[rwExporterName] = GenericMap{
			"endpoint": fmt.Sprintf("%s/api/v1/prom/remote/write", baseURL),
		}
		metricsPipelineName := "metrics/qryn-" + dest.GetID()
		currentConfig.Service.Pipelines[metricsPipelineName] = Pipeline{
			Exporters: []string{rwExporterName},
		}
	}

	if isTracingEnabled(dest) {
		exporterName := "otlp/qryn-" + dest.GetID()
		currentConfig.Exporters[exporterName] = GenericMap{
			"endpoint": fmt.Sprintf("%s/tempo/spans", baseURL),
		}
		tracesPipelineName := "traces/qryn-" + dest.GetID()
		currentConfig.Service.Pipelines[tracesPipelineName] = Pipeline{
			Exporters: []string{exporterName},
		}
	}

	if isLoggingEnabled(dest) {
		lokiExporterName := "loki/qryn-" + dest.GetID()
		currentConfig.Exporters[lokiExporterName] = GenericMap{
			"endpoint": fmt.Sprintf("%s/loki/api/v1/push", baseURL),
			"labels": GenericMap{
				"attributes": GenericMap{
					"k8s.container.name": "k8s_container_name",
					"k8s.pod.name":       "k8s_pod_name",
					"k8s.namespace.name": "k8s_namespace_name",
				},
			},
		}
		logsPipelineName := "logs/qryn-" + dest.GetID()
		currentConfig.Service.Pipelines[logsPipelineName] = Pipeline{
			Exporters: []string{lokiExporterName},
		}
	}

	return nil
}

func (g *Qryn) requiredVarsExists(dest ExporterConfigurer) bool {
	if _, ok := dest.GetConfig()[qrynHost]; !ok {
		return false
	}
	return true
}

func (g *Qryn) authData(dest ExporterConfigurer) (string, string) {
	var key string
	if k, ok := dest.GetConfig()[qrynAPIKey]; ok {
		key = k
	}
	return key, qrynAPISecret
}

func parseURL(rawURL, apiKey, apiSecret string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		return parseURL(fmt.Sprintf("https://%s", rawURL), apiKey, apiSecret)
	}

	return fmt.Sprintf("https://%s:%s@%s", apiKey, apiSecret, u.Host), nil
}
