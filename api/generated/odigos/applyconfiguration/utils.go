/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	odigosv1alpha1 "github.com/odigos-io/odigos/api/generated/odigos/applyconfiguration/odigos/v1alpha1"
	v1alpha1 "github.com/odigos-io/odigos/api/odigos/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=odigos.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("Attribute"):
		return &odigosv1alpha1.AttributeApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CollectorGatewayConfiguration"):
		return &odigosv1alpha1.CollectorGatewayConfigurationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CollectorsGroup"):
		return &odigosv1alpha1.CollectorsGroupApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CollectorsGroupSpec"):
		return &odigosv1alpha1.CollectorsGroupSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CollectorsGroupStatus"):
		return &odigosv1alpha1.CollectorsGroupStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ConfigOption"):
		return &odigosv1alpha1.ConfigOptionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Destination"):
		return &odigosv1alpha1.DestinationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DestinationSpec"):
		return &odigosv1alpha1.DestinationSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DestinationStatus"):
		return &odigosv1alpha1.DestinationStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EnvVar"):
		return &odigosv1alpha1.EnvVarApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InstrumentationInstance"):
		return &odigosv1alpha1.InstrumentationInstanceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InstrumentationInstanceStatus"):
		return &odigosv1alpha1.InstrumentationInstanceStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InstrumentationLibraryOptions"):
		return &odigosv1alpha1.InstrumentationLibraryOptionsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InstrumentationLibraryStatus"):
		return &odigosv1alpha1.InstrumentationLibraryStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InstrumentedApplication"):
		return &odigosv1alpha1.InstrumentedApplicationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InstrumentedApplicationSpec"):
		return &odigosv1alpha1.InstrumentedApplicationSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InstrumentedApplicationStatus"):
		return &odigosv1alpha1.InstrumentedApplicationStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OdigosConfiguration"):
		return &odigosv1alpha1.OdigosConfigurationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OdigosConfigurationSpec"):
		return &odigosv1alpha1.OdigosConfigurationSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OptionByContainer"):
		return &odigosv1alpha1.OptionByContainerApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Processor"):
		return &odigosv1alpha1.ProcessorApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ProcessorSpec"):
		return &odigosv1alpha1.ProcessorSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RuntimeDetailsByContainer"):
		return &odigosv1alpha1.RuntimeDetailsByContainerApplyConfiguration{}

	}
	return nil
}
