// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CredentialsInitParameters struct {

	// RFC2617 compliant password associated with the SIP credentials.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// RFC2617 compliant username associated with the SIP credentials.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsObservation struct {

	// RFC2617 compliant username associated with the SIP credentials.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsParameters struct {

	// RFC2617 compliant password associated with the SIP credentials.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// RFC2617 compliant username associated with the SIP credentials.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type VoiceConnectorTerminationCredentialsInitParameters struct {

	// List of termination SIP credentials.
	Credentials []CredentialsInitParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// Amazon Chime Voice Connector ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/chime/v1beta1.VoiceConnector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`

	// Reference to a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDRef *v1.Reference `json:"voiceConnectorIdRef,omitempty" tf:"-"`

	// Selector for a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDSelector *v1.Selector `json:"voiceConnectorIdSelector,omitempty" tf:"-"`
}

type VoiceConnectorTerminationCredentialsObservation struct {

	// List of termination SIP credentials.
	Credentials []CredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// Amazon Chime Voice Connector ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Chime Voice Connector ID.
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`
}

type VoiceConnectorTerminationCredentialsParameters struct {

	// List of termination SIP credentials.
	// +kubebuilder:validation:Optional
	Credentials []CredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Chime Voice Connector ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/chime/v1beta1.VoiceConnector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`

	// Reference to a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDRef *v1.Reference `json:"voiceConnectorIdRef,omitempty" tf:"-"`

	// Selector for a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDSelector *v1.Selector `json:"voiceConnectorIdSelector,omitempty" tf:"-"`
}

// VoiceConnectorTerminationCredentialsSpec defines the desired state of VoiceConnectorTerminationCredentials
type VoiceConnectorTerminationCredentialsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VoiceConnectorTerminationCredentialsParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VoiceConnectorTerminationCredentialsInitParameters `json:"initProvider,omitempty"`
}

// VoiceConnectorTerminationCredentialsStatus defines the observed state of VoiceConnectorTerminationCredentials.
type VoiceConnectorTerminationCredentialsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VoiceConnectorTerminationCredentialsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VoiceConnectorTerminationCredentials is the Schema for the VoiceConnectorTerminationCredentialss API. Adds termination SIP credentials for the specified Amazon Chime Voice Connector.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type VoiceConnectorTerminationCredentials struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentials) || (has(self.initProvider) && has(self.initProvider.credentials))",message="spec.forProvider.credentials is a required parameter"
	Spec   VoiceConnectorTerminationCredentialsSpec   `json:"spec"`
	Status VoiceConnectorTerminationCredentialsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorTerminationCredentialsList contains a list of VoiceConnectorTerminationCredentialss
type VoiceConnectorTerminationCredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VoiceConnectorTerminationCredentials `json:"items"`
}

// Repository type metadata.
var (
	VoiceConnectorTerminationCredentials_Kind             = "VoiceConnectorTerminationCredentials"
	VoiceConnectorTerminationCredentials_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VoiceConnectorTerminationCredentials_Kind}.String()
	VoiceConnectorTerminationCredentials_KindAPIVersion   = VoiceConnectorTerminationCredentials_Kind + "." + CRDGroupVersion.String()
	VoiceConnectorTerminationCredentials_GroupVersionKind = CRDGroupVersion.WithKind(VoiceConnectorTerminationCredentials_Kind)
)

func init() {
	SchemeBuilder.Register(&VoiceConnectorTerminationCredentials{}, &VoiceConnectorTerminationCredentialsList{})
}
