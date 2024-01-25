package config

import (
	"fmt"

	odigosv1 "github.com/keyval-dev/odigos/api/odigos/v1alpha1"
	commonconf "github.com/keyval-dev/odigos/autoscaler/controllers/common"
	"github.com/keyval-dev/odigos/common"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	s3BucketKey = "S3_BUCKET"
	s3RegionKey = "S3_REGION"
)

var (
	ErrS3BucketNotSpecified = fmt.Errorf("s3 bucket not specified")
	ErrS3RegionNotSpecified = fmt.Errorf("s3 region not specified")
)

type AWSS3 struct{}

func (s *AWSS3) DestType() common.DestinationType {
	return common.AWSS3DestinationType
}

func (s *AWSS3) ModifyConfig(dest *odigosv1.Destination, currentConfig *commonconf.Config) {
	bucket, ok := dest.Spec.Data[s3BucketKey]
	if !ok {
		ctrl.Log.Error(ErrS3BucketNotSpecified, "bucket not specified")
		return
	}

	region, ok := dest.Spec.Data[s3RegionKey]
	if !ok {
		ctrl.Log.Error(ErrS3RegionNotSpecified, "region not specified")
		return
	}

	if isLoggingEnabled(dest) {
		exporterName := "s3/" + dest.Name
		currentConfig.Exporters[exporterName] = commonconf.GenericMap{
			"settings": commonconf.GenericMap{
				"bucket": bucket,
				"region": region,
			},
		}

		logsPipelineName := "logs/s3-" + dest.Name
		currentConfig.Service.Pipelines[logsPipelineName] = commonconf.Pipeline{
			Receivers:  []string{"otlp"},
			Processors: []string{"batch"},
			Exporters:  []string{exporterName},
		}
	}

	if isTracingEnabled(dest) {
		exporterName := "s3/" + dest.Name
		currentConfig.Exporters[exporterName] = commonconf.GenericMap{
			"settings": commonconf.GenericMap{
				"bucket": bucket,
				"region": region,
			},
		}

		tracesPipelineName := "traces/s3-" + dest.Name
		currentConfig.Service.Pipelines[tracesPipelineName] = commonconf.Pipeline{
			Receivers:  []string{"otlp"},
			Processors: []string{"batch"},
			Exporters:  []string{exporterName},
		}
	}
}
