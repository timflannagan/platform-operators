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

// PlatformOperatorSpec defines the desired state of PlatformOperator
type PlatformOperatorSpec struct {
	// Package specifies the name of the package to be installed from the provided CatalogSource.
	// Package is required and must equal the exact name of the package in the catalog.
	Package string `json:"package"`
	// Channel specifies the current channel that's being used.
	Channel string `json:"channel"`
}

// PlatformOperatorStatus defines the observed state of PlatformOperator
type PlatformOperatorStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// PlatformOperator is the Schema for the platformoperators API
type PlatformOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlatformOperatorSpec   `json:"spec,omitempty"`
	Status PlatformOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PlatformOperatorList contains a list of PlatformOperator
type PlatformOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlatformOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PlatformOperator{}, &PlatformOperatorList{})
}
