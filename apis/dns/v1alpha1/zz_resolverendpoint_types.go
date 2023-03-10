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

type ResolverEndpointObservation struct {
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Self *string `json:"self,omitempty" tf:"self,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type ResolverEndpointParameters struct {

	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// +kubebuilder:validation:Optional
	ForwardingAddress *string `json:"forwardingAddress,omitempty" tf:"forwarding_address,omitempty"`

	// +kubebuilder:validation:Required
	IsForwarding *bool `json:"isForwarding" tf:"is_forwarding,omitempty"`

	// +kubebuilder:validation:Required
	IsListening *bool `json:"isListening" tf:"is_listening,omitempty"`

	// +kubebuilder:validation:Optional
	ListeningAddress *string `json:"listeningAddress,omitempty" tf:"listening_address,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.NetworkSecurityGroup
	// +kubebuilder:validation:Optional
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// +kubebuilder:validation:Optional
	NsgIdsRefs []v1.Reference `json:"nsgIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NsgIdsSelector *v1.Selector `json:"nsgIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Resolver
	// +kubebuilder:validation:Optional
	ResolverID *string `json:"resolverId,omitempty" tf:"resolver_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResolverIDRef *v1.Reference `json:"resolverIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResolverIDSelector *v1.Selector `json:"resolverIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// ResolverEndpointSpec defines the desired state of ResolverEndpoint
type ResolverEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResolverEndpointParameters `json:"forProvider"`
}

// ResolverEndpointStatus defines the observed state of ResolverEndpoint.
type ResolverEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResolverEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverEndpoint is the Schema for the ResolverEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type ResolverEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverEndpointSpec   `json:"spec"`
	Status            ResolverEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverEndpointList contains a list of ResolverEndpoints
type ResolverEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResolverEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ResolverEndpoint_Kind             = "ResolverEndpoint"
	ResolverEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResolverEndpoint_Kind}.String()
	ResolverEndpoint_KindAPIVersion   = ResolverEndpoint_Kind + "." + CRDGroupVersion.String()
	ResolverEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ResolverEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ResolverEndpoint{}, &ResolverEndpointList{})
}
