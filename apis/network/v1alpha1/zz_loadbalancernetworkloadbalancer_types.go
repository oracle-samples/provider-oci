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

type IPAddressesObservation struct {
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	ReservedIP []ReservedIPObservation `json:"reservedIp,omitempty" tf:"reserved_ip,omitempty"`
}

type IPAddressesParameters struct {
}

type LoadBalancerNetworkLoadBalancerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddresses []IPAddressesObservation `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type LoadBalancerNetworkLoadBalancerParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	IsPreserveSourceDestination *bool `json:"isPreserveSourceDestination,omitempty" tf:"is_preserve_source_destination,omitempty"`

	// +kubebuilder:validation:Optional
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIds []*string `json:"networkSecurityGroupIds,omitempty" tf:"network_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	NlbIPVersion *string `json:"nlbIpVersion,omitempty" tf:"nlb_ip_version,omitempty"`

	// +kubebuilder:validation:Optional
	ReservedIps []ReservedIpsParameters `json:"reservedIps,omitempty" tf:"reserved_ips,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type ReservedIPObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReservedIPParameters struct {
}

type ReservedIpsObservation struct {
}

type ReservedIpsParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

// LoadBalancerNetworkLoadBalancerSpec defines the desired state of LoadBalancerNetworkLoadBalancer
type LoadBalancerNetworkLoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerNetworkLoadBalancerParameters `json:"forProvider"`
}

// LoadBalancerNetworkLoadBalancerStatus defines the observed state of LoadBalancerNetworkLoadBalancer.
type LoadBalancerNetworkLoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerNetworkLoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNetworkLoadBalancer is the Schema for the LoadBalancerNetworkLoadBalancers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type LoadBalancerNetworkLoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerNetworkLoadBalancerSpec   `json:"spec"`
	Status            LoadBalancerNetworkLoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNetworkLoadBalancerList contains a list of LoadBalancerNetworkLoadBalancers
type LoadBalancerNetworkLoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerNetworkLoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerNetworkLoadBalancer_Kind             = "LoadBalancerNetworkLoadBalancer"
	LoadBalancerNetworkLoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerNetworkLoadBalancer_Kind}.String()
	LoadBalancerNetworkLoadBalancer_KindAPIVersion   = LoadBalancerNetworkLoadBalancer_Kind + "." + CRDGroupVersion.String()
	LoadBalancerNetworkLoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerNetworkLoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerNetworkLoadBalancer{}, &LoadBalancerNetworkLoadBalancerList{})
}
