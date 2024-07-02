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
	"github.com/odigos-io/odigos/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RedactionType string

const (
	CreditCardRedaction RedactionType = "CREDIT_CARD"
)

type RedactionAttribute struct {
	Redact              bool          `json:"redact"`
	RedactAttributeName RedactionType `json:"redactAttributeName"`
}

// RedactionSpec defines the desired state of Redaction action
type RedactionSpec struct {
	ActionName string                       `json:"actionName,omitempty"`
	Notes      string                       `json:"notes,omitempty"`
	Disabled   bool                         `json:"disabled,omitempty"`
	Signals    []common.ObservabilitySignal `json:"signals"`

	RedactionAttributes []RedactionAttribute `json:"redactionAttributes"`
}

// RedactionStatus defines the observed state of Redaction action
type RedactionStatus struct {
	// Represents the observations of a redaction's current state.
	// Known .status.conditions.type are: "Available", "Progressing"
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=redactions,scope=Namespaced,shortName=aci

// Redaction is the Schema for the redaction odigos action API
type Redaction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedactionSpec   `json:"spec,omitempty"`
	Status RedactionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RedactionList contains a list of Redaction
type RedactionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Redaction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Redaction{}, &RedactionList{})
}
