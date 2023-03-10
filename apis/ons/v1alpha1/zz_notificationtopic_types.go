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

type NotificationTopicObservation struct {
	APIEndpoint *string `json:"apiEndpoint,omitempty" tf:"api_endpoint,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ShortTopicID *string `json:"shortTopicId,omitempty" tf:"short_topic_id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TopicID *string `json:"topicId,omitempty" tf:"topic_id,omitempty"`
}

type NotificationTopicParameters struct {

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
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// NotificationTopicSpec defines the desired state of NotificationTopic
type NotificationTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationTopicParameters `json:"forProvider"`
}

// NotificationTopicStatus defines the observed state of NotificationTopic.
type NotificationTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationTopic is the Schema for the NotificationTopics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type NotificationTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationTopicSpec   `json:"spec"`
	Status            NotificationTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationTopicList contains a list of NotificationTopics
type NotificationTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationTopic `json:"items"`
}

// Repository type metadata.
var (
	NotificationTopic_Kind             = "NotificationTopic"
	NotificationTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationTopic_Kind}.String()
	NotificationTopic_KindAPIVersion   = NotificationTopic_Kind + "." + CRDGroupVersion.String()
	NotificationTopic_GroupVersionKind = CRDGroupVersion.WithKind(NotificationTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationTopic{}, &NotificationTopicList{})
}
