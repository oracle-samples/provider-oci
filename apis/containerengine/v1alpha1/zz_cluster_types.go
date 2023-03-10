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

type AddOnsObservation struct {
}

type AddOnsParameters struct {

	// +kubebuilder:validation:Optional
	IsKubernetesDashboardEnabled *bool `json:"isKubernetesDashboardEnabled,omitempty" tf:"is_kubernetes_dashboard_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IsTillerEnabled *bool `json:"isTillerEnabled,omitempty" tf:"is_tiller_enabled,omitempty"`
}

type AdmissionControllerOptionsObservation struct {
}

type AdmissionControllerOptionsParameters struct {

	// +kubebuilder:validation:Optional
	IsPodSecurityPolicyEnabled *bool `json:"isPodSecurityPolicyEnabled,omitempty" tf:"is_pod_security_policy_enabled,omitempty"`
}

type ClusterObservation struct {
	AvailableKubernetesUpgrades []*string `json:"availableKubernetesUpgrades,omitempty" tf:"available_kubernetes_upgrades,omitempty"`

	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	ClusterPodNetworkOptions []ClusterPodNetworkOptionsParameters `json:"clusterPodNetworkOptions,omitempty" tf:"cluster_pod_network_options,omitempty"`

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
	EndpointConfig []EndpointConfigParameters `json:"endpointConfig,omitempty" tf:"endpoint_config,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	ImagePolicyConfig []ImagePolicyConfigParameters `json:"imagePolicyConfig,omitempty" tf:"image_policy_config,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Required
	KubernetesVersion *string `json:"kubernetesVersion" tf:"kubernetes_version,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.Vcn
	// +kubebuilder:validation:Optional
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// +kubebuilder:validation:Optional
	VcnIDRef *v1.Reference `json:"vcnIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VcnIDSelector *v1.Selector `json:"vcnIdSelector,omitempty" tf:"-"`
}

type ClusterPodNetworkOptionsObservation struct {
}

type ClusterPodNetworkOptionsParameters struct {

	// +kubebuilder:validation:Required
	CniType *string `json:"cniType" tf:"cni_type,omitempty"`
}

type EndpointConfigObservation struct {
}

type EndpointConfigParameters struct {

	// +kubebuilder:validation:Optional
	IsPublicIPEnabled *bool `json:"isPublicIpEnabled,omitempty" tf:"is_public_ip_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.NetworkSecurityGroup
	// +kubebuilder:validation:Optional
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// +kubebuilder:validation:Optional
	NsgIdsRefs []v1.Reference `json:"nsgIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NsgIdsSelector *v1.Selector `json:"nsgIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type EndpointsObservation struct {
	Kubernetes *string `json:"kubernetes,omitempty" tf:"kubernetes,omitempty"`

	PrivateEndpoint *string `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	PublicEndpoint *string `json:"publicEndpoint,omitempty" tf:"public_endpoint,omitempty"`

	VcnHostnameEndpoint *string `json:"vcnHostnameEndpoint,omitempty" tf:"vcn_hostname_endpoint,omitempty"`
}

type EndpointsParameters struct {
}

type ImagePolicyConfigObservation struct {
}

type ImagePolicyConfigParameters struct {

	// +kubebuilder:validation:Optional
	IsPolicyEnabled *bool `json:"isPolicyEnabled,omitempty" tf:"is_policy_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KeyDetails []KeyDetailsParameters `json:"keyDetails,omitempty" tf:"key_details,omitempty"`
}

type KeyDetailsObservation struct {
}

type KeyDetailsParameters struct {

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type KubernetesNetworkConfigObservation struct {
}

type KubernetesNetworkConfigParameters struct {

	// +kubebuilder:validation:Optional
	PodsCidr *string `json:"podsCidr,omitempty" tf:"pods_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	ServicesCidr *string `json:"servicesCidr,omitempty" tf:"services_cidr,omitempty"`
}

type MetadataObservation struct {
	CreatedByUserID *string `json:"createdByUserId,omitempty" tf:"created_by_user_id,omitempty"`

	CreatedByWorkRequestID *string `json:"createdByWorkRequestId,omitempty" tf:"created_by_work_request_id,omitempty"`

	DeletedByUserID *string `json:"deletedByUserId,omitempty" tf:"deleted_by_user_id,omitempty"`

	DeletedByWorkRequestID *string `json:"deletedByWorkRequestId,omitempty" tf:"deleted_by_work_request_id,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeDeleted *string `json:"timeDeleted,omitempty" tf:"time_deleted,omitempty"`

	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`

	UpdatedByUserID *string `json:"updatedByUserId,omitempty" tf:"updated_by_user_id,omitempty"`

	UpdatedByWorkRequestID *string `json:"updatedByWorkRequestId,omitempty" tf:"updated_by_work_request_id,omitempty"`
}

type MetadataParameters struct {
}

type OptionsObservation struct {
}

type OptionsParameters struct {

	// +kubebuilder:validation:Optional
	AddOns []AddOnsParameters `json:"addOns,omitempty" tf:"add_ons,omitempty"`

	// +kubebuilder:validation:Optional
	AdmissionControllerOptions []AdmissionControllerOptionsParameters `json:"admissionControllerOptions,omitempty" tf:"admission_controller_options,omitempty"`

	// +kubebuilder:validation:Optional
	KubernetesNetworkConfig []KubernetesNetworkConfigParameters `json:"kubernetesNetworkConfig,omitempty" tf:"kubernetes_network_config,omitempty"`

	// +kubebuilder:validation:Optional
	PersistentVolumeConfig []PersistentVolumeConfigParameters `json:"persistentVolumeConfig,omitempty" tf:"persistent_volume_config,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceLBConfig []ServiceLBConfigParameters `json:"serviceLbConfig,omitempty" tf:"service_lb_config,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceLBSubnetIdSelector *v1.Selector `json:"serviceLbSubnetIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1.Subnet
	// +crossplane:generate:reference:refFieldName=ServiceLBSubnetIdsRef
	// +crossplane:generate:reference:selectorFieldName=ServiceLBSubnetIdSelector
	// +kubebuilder:validation:Optional
	ServiceLBSubnetIds []*string `json:"serviceLbSubnetIds,omitempty" tf:"service_lb_subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceLBSubnetIdsRef []v1.Reference `json:"serviceLbSubnetIdsRef,omitempty" tf:"-"`
}

type PersistentVolumeConfigObservation struct {
}

type PersistentVolumeConfigParameters struct {

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`
}

type ServiceLBConfigObservation struct {
}

type ServiceLBConfigParameters struct {

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
