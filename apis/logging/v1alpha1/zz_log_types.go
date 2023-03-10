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

type ConfigurationObservation struct {
}

type ConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`
}

type LogObservation struct {
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TenancyID *string `json:"tenancyId,omitempty" tf:"tenancy_id,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeLastModified *string `json:"timeLastModified,omitempty" tf:"time_last_modified,omitempty"`
}

type LogParameters struct {

	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// +crossplane:generate:reference:type=LogGroup
	// +kubebuilder:validation:Optional
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroupIDRef *v1.Reference `json:"logGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LogGroupIDSelector *v1.Selector `json:"logGroupIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	LogType *string `json:"logType" tf:"log_type,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionDuration *float64 `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Required
	Category *string `json:"category" tf:"category,omitempty"`

	// +kubebuilder:validation:Required
	Resource *string `json:"resource" tf:"resource,omitempty"`

	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`

	// +kubebuilder:validation:Required
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

// LogSpec defines the desired state of Log
type LogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogParameters `json:"forProvider"`
}

// LogStatus defines the observed state of Log.
type LogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Log is the Schema for the Logs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type Log struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogSpec   `json:"spec"`
	Status            LogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogList contains a list of Logs
type LogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Log `json:"items"`
}

// Repository type metadata.
var (
	Log_Kind             = "Log"
	Log_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Log_Kind}.String()
	Log_KindAPIVersion   = Log_Kind + "." + CRDGroupVersion.String()
	Log_GroupVersionKind = CRDGroupVersion.WithKind(Log_Kind)
)

func init() {
	SchemeBuilder.Register(&Log{}, &LogList{})
}
