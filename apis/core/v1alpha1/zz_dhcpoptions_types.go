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

type DHCPOptionsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type DHCPOptionsParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	DomainNameType *string `json:"domainNameType,omitempty" tf:"domain_name_type,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Required
	Options []OptionsParameters `json:"options" tf:"options,omitempty"`

	// +crossplane:generate:reference:type=Vcn
	// +kubebuilder:validation:Optional
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// +kubebuilder:validation:Optional
	VcnIDRef *v1.Reference `json:"vcnIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VcnIDSelector *v1.Selector `json:"vcnIdSelector,omitempty" tf:"-"`
}

type OptionsObservation struct {
}

type OptionsParameters struct {

	// +kubebuilder:validation:Optional
	CustomDNSServers []*string `json:"customDnsServers,omitempty" tf:"custom_dns_servers,omitempty"`

	// +kubebuilder:validation:Optional
	SearchDomainNames []*string `json:"searchDomainNames,omitempty" tf:"search_domain_names,omitempty"`

	// +kubebuilder:validation:Optional
	ServerType *string `json:"serverType,omitempty" tf:"server_type,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// DHCPOptionsSpec defines the desired state of DHCPOptions
type DHCPOptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DHCPOptionsParameters `json:"forProvider"`
}

// DHCPOptionsStatus defines the observed state of DHCPOptions.
type DHCPOptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DHCPOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DHCPOptions is the Schema for the DHCPOptionss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type DHCPOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DHCPOptionsSpec   `json:"spec"`
	Status            DHCPOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DHCPOptionsList contains a list of DHCPOptionss
type DHCPOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DHCPOptions `json:"items"`
}

// Repository type metadata.
var (
	DHCPOptions_Kind             = "DHCPOptions"
	DHCPOptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DHCPOptions_Kind}.String()
	DHCPOptions_KindAPIVersion   = DHCPOptions_Kind + "." + CRDGroupVersion.String()
	DHCPOptions_GroupVersionKind = CRDGroupVersion.WithKind(DHCPOptions_Kind)
)

func init() {
	SchemeBuilder.Register(&DHCPOptions{}, &DHCPOptionsList{})
}
