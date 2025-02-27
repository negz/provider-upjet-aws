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

type CiphertextInitParameters struct {

	// An optional mapping that makes up the encryption context.
	// +mapType=granular
	Context map[string]*string `json:"context,omitempty" tf:"context,omitempty"`

	// Globally unique key ID for the customer master key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kms/v1beta1.Key
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`

	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	PlaintextSecretRef v1.SecretKeySelector `json:"plaintextSecretRef" tf:"-"`
}

type CiphertextObservation struct {

	// Base64 encoded ciphertext
	CiphertextBlob *string `json:"ciphertextBlob,omitempty" tf:"ciphertext_blob,omitempty"`

	// An optional mapping that makes up the encryption context.
	// +mapType=granular
	Context map[string]*string `json:"context,omitempty" tf:"context,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Globally unique key ID for the customer master key.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type CiphertextParameters struct {

	// An optional mapping that makes up the encryption context.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Context map[string]*string `json:"context,omitempty" tf:"context,omitempty"`

	// Globally unique key ID for the customer master key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`

	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	// +kubebuilder:validation:Optional
	PlaintextSecretRef v1.SecretKeySelector `json:"plaintextSecretRef" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// CiphertextSpec defines the desired state of Ciphertext
type CiphertextSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CiphertextParameters `json:"forProvider"`
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
	InitProvider CiphertextInitParameters `json:"initProvider,omitempty"`
}

// CiphertextStatus defines the observed state of Ciphertext.
type CiphertextStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CiphertextObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Ciphertext is the Schema for the Ciphertexts API. Provides ciphertext encrypted using a KMS key
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Ciphertext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plaintextSecretRef)",message="spec.forProvider.plaintextSecretRef is a required parameter"
	Spec   CiphertextSpec   `json:"spec"`
	Status CiphertextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CiphertextList contains a list of Ciphertexts
type CiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ciphertext `json:"items"`
}

// Repository type metadata.
var (
	Ciphertext_Kind             = "Ciphertext"
	Ciphertext_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ciphertext_Kind}.String()
	Ciphertext_KindAPIVersion   = Ciphertext_Kind + "." + CRDGroupVersion.String()
	Ciphertext_GroupVersionKind = CRDGroupVersion.WithKind(Ciphertext_Kind)
)

func init() {
	SchemeBuilder.Register(&Ciphertext{}, &CiphertextList{})
}
