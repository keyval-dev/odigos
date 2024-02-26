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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/keyval-dev/odigos/frontend/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Destinations returns a DestinationInformer.
	Destinations() DestinationInformer
	// InstrumentedApplications returns a InstrumentedApplicationInformer.
	InstrumentedApplications() InstrumentedApplicationInformer
	// OdigosConfigurations returns a OdigosConfigurationInformer.
	OdigosConfigurations() OdigosConfigurationInformer
	// Processors returns a ProcessorInformer.
	Processors() ProcessorInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Destinations returns a DestinationInformer.
func (v *version) Destinations() DestinationInformer {
	return &destinationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstrumentedApplications returns a InstrumentedApplicationInformer.
func (v *version) InstrumentedApplications() InstrumentedApplicationInformer {
	return &instrumentedApplicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OdigosConfigurations returns a OdigosConfigurationInformer.
func (v *version) OdigosConfigurations() OdigosConfigurationInformer {
	return &odigosConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Processors returns a ProcessorInformer.
func (v *version) Processors() ProcessorInformer {
	return &processorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
