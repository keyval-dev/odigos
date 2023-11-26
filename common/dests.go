package common

type DestinationType string

const (
	MiddlewareDestinationType    DestinationType = "middleware"
	GrafanaDestinationType       DestinationType = "grafana"
	DatadogDestinationType       DestinationType = "datadog"
	HoneycombDestinationType     DestinationType = "honeycomb"
	NewRelicDestinationType      DestinationType = "newrelic"
	LogzioDestinationType        DestinationType = "logzio"
	PrometheusDestinationType    DestinationType = "prometheus"
	LokiDestinationType          DestinationType = "loki"
	TempoDestinationType         DestinationType = "tempo"
	JaegerDestinationType        DestinationType = "jaeger"
	ElasticsearchDestinationType DestinationType = "elasticsearch"
	GenericOTLPDestinationType   DestinationType = "otlp"
	SignozDestinationType        DestinationType = "signoz"
	QrynDestinationType          DestinationType = "qryn"
	OpsVerseDestinationType      DestinationType = "opsverse"
	SplunkDestinationType        DestinationType = "splunk"
	LightstepDestinationType     DestinationType = "lightstep"
	GoogleCloudDestinationType   DestinationType = "googlecloud"
	SentryDestinationType        DestinationType = "sentry"
	GCSDestinationType           DestinationType = "gcs"
	AWSS3DestinationType         DestinationType = "s3"
	AzureBlobDestinationType     DestinationType = "azureblob"
	DynatraceDestinationType     DestinationType = "dynatrace"
	ChronosphereDestinationType  DestinationType = "chronosphere"
	ElasticAPMDestinationType    DestinationType = "elasticapm"
	AxiomDestinationType         DestinationType = "axiom"
	SumoLogicDestinationType     DestinationType = "sumologic"
)
