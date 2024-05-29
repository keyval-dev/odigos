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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstrumentationInstanceSpec struct {
}

// +kubebuilder:validation:Enum=instrumentation;sampler;exporter
type InstrumentationLibraryType string

const (
	InstrumentationLibraryTypeInstrumentation InstrumentationLibraryType = "instrumentation"
	InstrumentationLibraryTypeSampler 	    InstrumentationLibraryType = "sampler"
	InstrumentationLibraryTypeExporter 	    InstrumentationLibraryType = "exporter"
)

// Attribute is a key-value pair that describes a component or instrumentation
type Attribute struct {
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Key string `json:"key"`
	// +required
	// +kubebuilder:validation:Required
	Value string `json:"value"`
}

// InstrumentationLibraryStatus defines the observed state of an InstrumentationLibrary.
// if a library is not active/disable, it should not be included in the status
type InstrumentationLibraryStatus struct {
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// for example ("net/http", "@opentelemetry/instrumentation-redis")
	Name string `json:"name"`
	// +required
	// +kubebuilder:validation:Required
	Type InstrumentationLibraryType `json:"type"`
	// Attributes that identify the component.
	// The combination of (Name, Type, IdentifyingAttributes) must be unique.
	IdentifyingAttributes []Attribute `json:"identifyingAttributes,omitempty"`
	// Attributes that do not necessarily identify the component but help describe
    // its characteristics.
	NonIdentifyingAttributes []Attribute `json:"nonIdentifyingAttributes,omitempty"`
	Healthy *bool `json:"healthy,omitempty"`
	// message is a human readable message indicating details about the component health.
	// can be omitted if healthy is true
	// +kubebuilder:validation:MaxLength=32768
	Message string `json:"message,omitempty"`
	// reason contains a programmatic identifier indicating the reason for the SDK status.
	// Producers of specific condition types may define expected values and meanings for this field,
	// and whether the values are considered a guaranteed API.
	Reason string `json:"reason,omitempty"`
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	LastStatusTime metav1.Time `json:"lastStatusTime"`
}


// InstrumentationInstanceStatus defines the observed state of InstrumentationInstance
// If the instrumentation is not active, this CR should be deleted
type InstrumentationInstanceStatus struct {
	// Attributes that do not necessarily identify the SDK but help describe
    // its characteristics.
	NonIdentifyingAttributes []Attribute `json:"nonIdentifyingAttributes,omitempty"`
	Healthy *bool `json:"healthy,omitempty"`
	// message is a human readable message indicating details about the SDK general health.
	// can be omitted if healthy is true
	// +kubebuilder:validation:MaxLength=32768
	Message string `json:"message,omitempty"`
	// reason contains a programmatic identifier indicating the reason for the component status.
	// Producers of specific condition types may define expected values and meanings for this field,
	// and whether the values are considered a guaranteed API.
	Reason string `json:"reason,omitempty"`
	// +required
	// +kubebuilder:validation:Required
	StartTime metav1.Time `json:"startTime"`
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	LastStatusTime metav1.Time `json:"lastStatusTime"`
	Components []InstrumentationLibraryStatus `json:"components,omitempty"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// InstrumentationInstance is the Schema for the InstrumentationInstances API
type InstrumentationInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstrumentationInstanceSpec   `json:"spec,omitempty"`
	Status InstrumentationInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InstrumentationInstanceList contains a list of InstrumentationInstance
type InstrumentationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstrumentationInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstrumentationInstance{}, &InstrumentationInstanceList{})
}
