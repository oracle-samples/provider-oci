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

type PrivateIPObservation struct {
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`

	IsReserved *bool `json:"isReserved,omitempty" tf:"is_reserved,omitempty"`

	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type PrivateIPParameters struct {

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	HostnameLabel *string `json:"hostnameLabel,omitempty" tf:"hostname_label,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +crossplane:generate:reference:type=Vlan
	// +kubebuilder:validation:Optional
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// +kubebuilder:validation:Optional
	VlanIDRef *v1.Reference `json:"vlanIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VlanIDSelector *v1.Selector `json:"vlanIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=VnicAttachment
	// +kubebuilder:validation:Optional
	VnicID *string `json:"vnicId,omitempty" tf:"vnic_id,omitempty"`

	// +kubebuilder:validation:Optional
	VnicIDRef *v1.Reference `json:"vnicIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VnicIDSelector *v1.Selector `json:"vnicIdSelector,omitempty" tf:"-"`
}

// PrivateIPSpec defines the desired state of PrivateIP
type PrivateIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateIPParameters `json:"forProvider"`
}

// PrivateIPStatus defines the observed state of PrivateIP.
type PrivateIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateIP is the Schema for the PrivateIPs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type PrivateIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateIPSpec   `json:"spec"`
	Status            PrivateIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateIPList contains a list of PrivateIPs
type PrivateIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateIP `json:"items"`
}

// Repository type metadata.
var (
	PrivateIP_Kind             = "PrivateIP"
	PrivateIP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateIP_Kind}.String()
	PrivateIP_KindAPIVersion   = PrivateIP_Kind + "." + CRDGroupVersion.String()
	PrivateIP_GroupVersionKind = CRDGroupVersion.WithKind(PrivateIP_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateIP{}, &PrivateIPList{})
}
