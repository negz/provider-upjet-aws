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

type SigningCertificateInitParameters struct {

	// encoded format.
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// –   The status you want to assign to the certificate. Active means that the certificate can be used for programmatic calls to Amazon Web Services Inactive means that the certificate cannot be used.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// –  The name of the user the signing certificate is for.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type SigningCertificateObservation struct {

	// encoded format.
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// The ID for the signing certificate.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The certificate_id:user_name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –   The status you want to assign to the certificate. Active means that the certificate can be used for programmatic calls to Amazon Web Services Inactive means that the certificate cannot be used.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// –  The name of the user the signing certificate is for.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type SigningCertificateParameters struct {

	// encoded format.
	// +kubebuilder:validation:Optional
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// –   The status you want to assign to the certificate. Active means that the certificate can be used for programmatic calls to Amazon Web Services Inactive means that the certificate cannot be used.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// –  The name of the user the signing certificate is for.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// SigningCertificateSpec defines the desired state of SigningCertificate
type SigningCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SigningCertificateParameters `json:"forProvider"`
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
	InitProvider SigningCertificateInitParameters `json:"initProvider,omitempty"`
}

// SigningCertificateStatus defines the observed state of SigningCertificate.
type SigningCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SigningCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SigningCertificate is the Schema for the SigningCertificates API. Provides an IAM Signing Certificate
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type SigningCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificateBody) || (has(self.initProvider) && has(self.initProvider.certificateBody))",message="spec.forProvider.certificateBody is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userName) || (has(self.initProvider) && has(self.initProvider.userName))",message="spec.forProvider.userName is a required parameter"
	Spec   SigningCertificateSpec   `json:"spec"`
	Status SigningCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SigningCertificateList contains a list of SigningCertificates
type SigningCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SigningCertificate `json:"items"`
}

// Repository type metadata.
var (
	SigningCertificate_Kind             = "SigningCertificate"
	SigningCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SigningCertificate_Kind}.String()
	SigningCertificate_KindAPIVersion   = SigningCertificate_Kind + "." + CRDGroupVersion.String()
	SigningCertificate_GroupVersionKind = CRDGroupVersion.WithKind(SigningCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&SigningCertificate{}, &SigningCertificateList{})
}
