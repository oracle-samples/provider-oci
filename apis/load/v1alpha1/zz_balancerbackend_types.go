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

type BalancerBackendObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type BalancerBackendParameters struct {

	// +crossplane:generate:reference:type=BalancerBackendSet
	// +kubebuilder:validation:Optional
	BackendsetName *string `json:"backendsetName,omitempty" tf:"backendset_name,omitempty"`

	// +kubebuilder:validation:Optional
	BackendsetNameRef *v1.Reference `json:"backendsetNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BackendsetNameSelector *v1.Selector `json:"backendsetNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// +kubebuilder:validation:Optional
	Drain *bool `json:"drain,omitempty" tf:"drain,omitempty"`

	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// +crossplane:generate:reference:type=BalancerLoadBalancer
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Offline *bool `json:"offline,omitempty" tf:"offline,omitempty"`

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// BalancerBackendSpec defines the desired state of BalancerBackend
type BalancerBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerBackendParameters `json:"forProvider"`
}

// BalancerBackendStatus defines the observed state of BalancerBackend.
type BalancerBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerBackend is the Schema for the BalancerBackends API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type BalancerBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BalancerBackendSpec   `json:"spec"`
	Status            BalancerBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerBackendList contains a list of BalancerBackends
type BalancerBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerBackend `json:"items"`
}

// Repository type metadata.
var (
	BalancerBackend_Kind             = "BalancerBackend"
	BalancerBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerBackend_Kind}.String()
	BalancerBackend_KindAPIVersion   = BalancerBackend_Kind + "." + CRDGroupVersion.String()
	BalancerBackend_GroupVersionKind = CRDGroupVersion.WithKind(BalancerBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerBackend{}, &BalancerBackendList{})
}
