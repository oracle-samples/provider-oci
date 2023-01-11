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

type ApplicationListsObservation struct {
}

type ApplicationListsParameters struct {

	// +kubebuilder:validation:Optional
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumPort *float64 `json:"maximumPort,omitempty" tf:"maximum_port,omitempty"`

	// +kubebuilder:validation:Optional
	MinimumPort *float64 `json:"minimumPort,omitempty" tf:"minimum_port,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Destinations []*string `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// +kubebuilder:validation:Optional
	Sources []*string `json:"sources,omitempty" tf:"sources,omitempty"`
}

type DecryptionProfilesObservation struct {
}

type DecryptionProfilesParameters struct {

	// +kubebuilder:validation:Optional
	AreCertificateExtensionsRestricted *bool `json:"areCertificateExtensionsRestricted,omitempty" tf:"are_certificate_extensions_restricted,omitempty"`

	// +kubebuilder:validation:Optional
	IsAutoIncludeAltName *bool `json:"isAutoIncludeAltName,omitempty" tf:"is_auto_include_alt_name,omitempty"`

	// +kubebuilder:validation:Optional
	IsExpiredCertificateBlocked *bool `json:"isExpiredCertificateBlocked,omitempty" tf:"is_expired_certificate_blocked,omitempty"`

	// +kubebuilder:validation:Optional
	IsOutOfCapacityBlocked *bool `json:"isOutOfCapacityBlocked,omitempty" tf:"is_out_of_capacity_blocked,omitempty"`

	// +kubebuilder:validation:Optional
	IsRevocationStatusTimeoutBlocked *bool `json:"isRevocationStatusTimeoutBlocked,omitempty" tf:"is_revocation_status_timeout_blocked,omitempty"`

	// +kubebuilder:validation:Optional
	IsUnknownRevocationStatusBlocked *bool `json:"isUnknownRevocationStatusBlocked,omitempty" tf:"is_unknown_revocation_status_blocked,omitempty"`

	// +kubebuilder:validation:Optional
	IsUnsupportedCipherBlocked *bool `json:"isUnsupportedCipherBlocked,omitempty" tf:"is_unsupported_cipher_blocked,omitempty"`

	// +kubebuilder:validation:Optional
	IsUnsupportedVersionBlocked *bool `json:"isUnsupportedVersionBlocked,omitempty" tf:"is_unsupported_version_blocked,omitempty"`

	// +kubebuilder:validation:Optional
	IsUntrustedIssuerBlocked *bool `json:"isUntrustedIssuerBlocked,omitempty" tf:"is_untrusted_issuer_blocked,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type DecryptionRulesObservation struct {
}

type DecryptionRulesParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Condition []ConditionParameters `json:"condition" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	DecryptionProfile *string `json:"decryptionProfile,omitempty" tf:"decryption_profile,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`
}

type FirewallNetworkFirewallPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsFirewallAttached *bool `json:"isFirewallAttached,omitempty" tf:"is_firewall_attached,omitempty"`

	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type FirewallNetworkFirewallPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ApplicationLists []ApplicationListsParameters `json:"applicationLists,omitempty" tf:"application_lists,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DecryptionProfiles []DecryptionProfilesParameters `json:"decryptionProfiles,omitempty" tf:"decryption_profiles,omitempty"`

	// +kubebuilder:validation:Optional
	DecryptionRules []DecryptionRulesParameters `json:"decryptionRules,omitempty" tf:"decryption_rules,omitempty"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressLists []IPAddressListsParameters `json:"ipAddressLists,omitempty" tf:"ip_address_lists,omitempty"`

	// +kubebuilder:validation:Optional
	MappedSecrets []MappedSecretsParameters `json:"mappedSecrets,omitempty" tf:"mapped_secrets,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityRules []SecurityRulesParameters `json:"securityRules,omitempty" tf:"security_rules,omitempty"`

	// +kubebuilder:validation:Optional
	URLLists []URLListsParameters `json:"urlLists,omitempty" tf:"url_lists,omitempty"`
}

type IPAddressListsObservation struct {
}

type IPAddressListsParameters struct {

	// +kubebuilder:validation:Required
	IPAddressListName *string `json:"ipAddressListName" tf:"ip_address_list_name,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressListValue []*string `json:"ipAddressListValue,omitempty" tf:"ip_address_list_value,omitempty"`
}

type MappedSecretsObservation struct {
}

type MappedSecretsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	VaultSecretID *string `json:"vaultSecretId,omitempty" tf:"vault_secret_id,omitempty"`

	// +kubebuilder:validation:Optional
	VersionNumber *float64 `json:"versionNumber,omitempty" tf:"version_number,omitempty"`
}

type SecurityRulesConditionObservation struct {
}

type SecurityRulesConditionParameters struct {

	// +kubebuilder:validation:Optional
	Applications []*string `json:"applications,omitempty" tf:"applications,omitempty"`

	// +kubebuilder:validation:Optional
	Destinations []*string `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// +kubebuilder:validation:Optional
	Sources []*string `json:"sources,omitempty" tf:"sources,omitempty"`

	// +kubebuilder:validation:Optional
	Urls []*string `json:"urls,omitempty" tf:"urls,omitempty"`
}

type SecurityRulesObservation struct {
}

type SecurityRulesParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Condition []SecurityRulesConditionParameters `json:"condition" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Inspection *string `json:"inspection,omitempty" tf:"inspection,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type URLListsObservation struct {
}

type URLListsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// FirewallNetworkFirewallPolicySpec defines the desired state of FirewallNetworkFirewallPolicy
type FirewallNetworkFirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallNetworkFirewallPolicyParameters `json:"forProvider"`
}

// FirewallNetworkFirewallPolicyStatus defines the observed state of FirewallNetworkFirewallPolicy.
type FirewallNetworkFirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallNetworkFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallNetworkFirewallPolicy is the Schema for the FirewallNetworkFirewallPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type FirewallNetworkFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallNetworkFirewallPolicySpec   `json:"spec"`
	Status            FirewallNetworkFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallNetworkFirewallPolicyList contains a list of FirewallNetworkFirewallPolicys
type FirewallNetworkFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallNetworkFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	FirewallNetworkFirewallPolicy_Kind             = "FirewallNetworkFirewallPolicy"
	FirewallNetworkFirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallNetworkFirewallPolicy_Kind}.String()
	FirewallNetworkFirewallPolicy_KindAPIVersion   = FirewallNetworkFirewallPolicy_Kind + "." + CRDGroupVersion.String()
	FirewallNetworkFirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FirewallNetworkFirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallNetworkFirewallPolicy{}, &FirewallNetworkFirewallPolicyList{})
}