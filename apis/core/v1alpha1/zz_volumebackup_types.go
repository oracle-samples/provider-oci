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

type VolumeBackupObservation struct {
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	SizeInGbs *string `json:"sizeInGbs,omitempty" tf:"size_in_gbs,omitempty"`

	SizeInMbs *string `json:"sizeInMbs,omitempty" tf:"size_in_mbs,omitempty"`

	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	SourceVolumeBackupID *string `json:"sourceVolumeBackupId,omitempty" tf:"source_volume_backup_id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeRequestReceived *string `json:"timeRequestReceived,omitempty" tf:"time_request_received,omitempty"`

	UniqueSizeInGbs *string `json:"uniqueSizeInGbs,omitempty" tf:"unique_size_in_gbs,omitempty"`

	UniqueSizeInMbs *string `json:"uniqueSizeInMbs,omitempty" tf:"unique_size_in_mbs,omitempty"`
}

type VolumeBackupParameters struct {

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
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDetails []VolumeBackupSourceDetailsParameters `json:"sourceDetails,omitempty" tf:"source_details,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +crossplane:generate:reference:type=Volume
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

type VolumeBackupSourceDetailsObservation struct {
}

type VolumeBackupSourceDetailsParameters struct {

	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	VolumeBackupID *string `json:"volumeBackupId" tf:"volume_backup_id,omitempty"`
}

// VolumeBackupSpec defines the desired state of VolumeBackup
type VolumeBackupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeBackupParameters `json:"forProvider"`
}

// VolumeBackupStatus defines the observed state of VolumeBackup.
type VolumeBackupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeBackupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeBackup is the Schema for the VolumeBackups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type VolumeBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeBackupSpec   `json:"spec"`
	Status            VolumeBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeBackupList contains a list of VolumeBackups
type VolumeBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeBackup `json:"items"`
}

// Repository type metadata.
var (
	VolumeBackup_Kind             = "VolumeBackup"
	VolumeBackup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeBackup_Kind}.String()
	VolumeBackup_KindAPIVersion   = VolumeBackup_Kind + "." + CRDGroupVersion.String()
	VolumeBackup_GroupVersionKind = CRDGroupVersion.WithKind(VolumeBackup_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeBackup{}, &VolumeBackupList{})
}
