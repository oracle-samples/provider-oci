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

type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedBackendsObservation struct {
}

type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedBackendsParameters struct {

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	IsBackup *bool `json:"isBackup,omitempty" tf:"is_backup,omitempty"`

	// +kubebuilder:validation:Optional
	IsDrain *bool `json:"isDrain,omitempty" tf:"is_drain,omitempty"`

	// +kubebuilder:validation:Optional
	IsOffline *bool `json:"isOffline,omitempty" tf:"is_offline,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedHealthCheckerObservation struct {
}

type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedHealthCheckerParameters struct {

	// +kubebuilder:validation:Optional
	IntervalInMillis *float64 `json:"intervalInMillis,omitempty" tf:"interval_in_millis,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	RequestData *string `json:"requestData,omitempty" tf:"request_data,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseBodyRegex *string `json:"responseBodyRegex,omitempty" tf:"response_body_regex,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseData *string `json:"responseData,omitempty" tf:"response_data,omitempty"`

	// +kubebuilder:validation:Optional
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// +kubebuilder:validation:Optional
	ReturnCode *float64 `json:"returnCode,omitempty" tf:"return_code,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutInMillis *float64 `json:"timeoutInMillis,omitempty" tf:"timeout_in_millis,omitempty"`

	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedParameters struct {

	// +kubebuilder:validation:Optional
	Backends []LoadBalancerNetworkLoadBalancersBackendSetsUnifiedBackendsParameters `json:"backends,omitempty" tf:"backends,omitempty"`

	// +kubebuilder:validation:Required
	HealthChecker []LoadBalancerNetworkLoadBalancersBackendSetsUnifiedHealthCheckerParameters `json:"healthChecker" tf:"health_checker,omitempty"`

	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// +kubebuilder:validation:Optional
	IsPreserveSource *bool `json:"isPreserveSource,omitempty" tf:"is_preserve_source,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=LoadBalancerNetworkLoadBalancer
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerID *string `json:"networkLoadBalancerId,omitempty" tf:"network_load_balancer_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkLoadBalancerIDRef *v1.Reference `json:"networkLoadBalancerIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkLoadBalancerIDSelector *v1.Selector `json:"networkLoadBalancerIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`
}

// LoadBalancerNetworkLoadBalancersBackendSetsUnifiedSpec defines the desired state of LoadBalancerNetworkLoadBalancersBackendSetsUnified
type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerNetworkLoadBalancersBackendSetsUnifiedParameters `json:"forProvider"`
}

// LoadBalancerNetworkLoadBalancersBackendSetsUnifiedStatus defines the observed state of LoadBalancerNetworkLoadBalancersBackendSetsUnified.
type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerNetworkLoadBalancersBackendSetsUnifiedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNetworkLoadBalancersBackendSetsUnified is the Schema for the LoadBalancerNetworkLoadBalancersBackendSetsUnifieds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type LoadBalancerNetworkLoadBalancersBackendSetsUnified struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerNetworkLoadBalancersBackendSetsUnifiedSpec   `json:"spec"`
	Status            LoadBalancerNetworkLoadBalancersBackendSetsUnifiedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNetworkLoadBalancersBackendSetsUnifiedList contains a list of LoadBalancerNetworkLoadBalancersBackendSetsUnifieds
type LoadBalancerNetworkLoadBalancersBackendSetsUnifiedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerNetworkLoadBalancersBackendSetsUnified `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerNetworkLoadBalancersBackendSetsUnified_Kind             = "LoadBalancerNetworkLoadBalancersBackendSetsUnified"
	LoadBalancerNetworkLoadBalancersBackendSetsUnified_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerNetworkLoadBalancersBackendSetsUnified_Kind}.String()
	LoadBalancerNetworkLoadBalancersBackendSetsUnified_KindAPIVersion   = LoadBalancerNetworkLoadBalancersBackendSetsUnified_Kind + "." + CRDGroupVersion.String()
	LoadBalancerNetworkLoadBalancersBackendSetsUnified_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerNetworkLoadBalancersBackendSetsUnified_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerNetworkLoadBalancersBackendSetsUnified{}, &LoadBalancerNetworkLoadBalancersBackendSetsUnifiedList{})
}
