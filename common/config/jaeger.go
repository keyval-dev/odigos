package config

import (
	"errors"

	"github.com/odigos-io/odigos/common"
)

var (
	ErrorJaegerTracingDisabled = errors.New("attempting to configure Jaeger tracing, but tracing is disabled")
	ErrorJaegerMissingURL      = errors.New("missing Jaeger JAEGER_URL config")
	ErrorJaegerNoTls           = errors.New("jaeger destination only supports non tls connections")
)

const (
	jaegerUrlKey = "JAEGER_URL"
)

type Jaeger struct{}

func (j *Jaeger) DestType() common.DestinationType {
	return common.JaegerDestinationType
}

func (j *Jaeger) ModifyConfig(dest ExporterConfigurer, currentConfig *Config) error {

	if !isTracingEnabled(dest) {
		return ErrorJaegerTracingDisabled
	}

	url, urlExist := dest.GetConfig()[jaegerUrlKey]
	if !urlExist {
		return ErrorJaegerMissingURL
	}

	grpcEndpoint, err := parseUnencryptedOtlpGrpcUrl(url)
	if err != nil {
		return err
	}

	exporterName := "otlp/jaeger-" + dest.GetName()
	currentConfig.Exporters[exporterName] = GenericMap{
		"endpoint": grpcEndpoint,
		"tls": GenericMap{
			"insecure": true,
		},
	}

	pipelineName := "traces/jaeger-" + dest.GetName()
	currentConfig.Service.Pipelines[pipelineName] = Pipeline{
		Exporters: []string{exporterName},
	}
	return nil
}
