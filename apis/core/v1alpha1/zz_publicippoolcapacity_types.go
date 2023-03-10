/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PublicIPPoolCapacityObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PublicIPPoolCapacityParameters struct {

	// +kubebuilder:validation:Required
	ByoipID *string `json:"byoipId" tf:"byoip_id,omitempty"`

	// +kubebuilder:validation:Required
	CidrBlock *string `json:"cidrBlock" tf:"cidr_block,omitempty"`

	// +crossplane:generate:reference:type=PublicIPPool
	// +kubebuilder:validation:Optional
	PublicIPPoolID *string `json:"publicIpPoolId,omitempty" tf:"public_ip_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPPoolIDRef *v1.Reference `json:"publicIpPoolIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PublicIPPoolIDSelector *v1.Selector `json:"publicIpPoolIdSelector,omitempty" tf:"-"`
}

// PublicIPPoolCapacitySpec defines the desired state of PublicIPPoolCapacity
type PublicIPPoolCapacitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicIPPoolCapacityParameters `json:"forProvider"`
}

// PublicIPPoolCapacityStatus defines the observed state of PublicIPPoolCapacity.
type PublicIPPoolCapacityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicIPPoolCapacityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIPPoolCapacity is the Schema for the PublicIPPoolCapacitys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type PublicIPPoolCapacity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicIPPoolCapacitySpec   `json:"spec"`
	Status            PublicIPPoolCapacityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIPPoolCapacityList contains a list of PublicIPPoolCapacitys
type PublicIPPoolCapacityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicIPPoolCapacity `json:"items"`
}

// Repository type metadata.
var (
	PublicIPPoolCapacity_Kind             = "PublicIPPoolCapacity"
	PublicIPPoolCapacity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicIPPoolCapacity_Kind}.String()
	PublicIPPoolCapacity_KindAPIVersion   = PublicIPPoolCapacity_Kind + "." + CRDGroupVersion.String()
	PublicIPPoolCapacity_GroupVersionKind = CRDGroupVersion.WithKind(PublicIPPoolCapacity_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicIPPoolCapacity{}, &PublicIPPoolCapacityList{})
}
