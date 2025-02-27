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

type CertificateValidationInitParameters struct {

	// ARN of the certificate that is being validated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/acm/v1beta2.Certificate
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Reference to a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnRef *v1.Reference `json:"certificateArnRef,omitempty" tf:"-"`

	// Selector for a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnSelector *v1.Selector `json:"certificateArnSelector,omitempty" tf:"-"`

	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	// +listType=set
	ValidationRecordFqdns []*string `json:"validationRecordFqdns,omitempty" tf:"validation_record_fqdns,omitempty"`
}

type CertificateValidationObservation struct {

	// ARN of the certificate that is being validated.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Time at which the certificate was issued
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	// +listType=set
	ValidationRecordFqdns []*string `json:"validationRecordFqdns,omitempty" tf:"validation_record_fqdns,omitempty"`
}

type CertificateValidationParameters struct {

	// ARN of the certificate that is being validated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/acm/v1beta2.Certificate
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Reference to a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnRef *v1.Reference `json:"certificateArnRef,omitempty" tf:"-"`

	// Selector for a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnSelector *v1.Selector `json:"certificateArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	// +kubebuilder:validation:Optional
	// +listType=set
	ValidationRecordFqdns []*string `json:"validationRecordFqdns,omitempty" tf:"validation_record_fqdns,omitempty"`
}

// CertificateValidationSpec defines the desired state of CertificateValidation
type CertificateValidationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateValidationParameters `json:"forProvider"`
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
	InitProvider CertificateValidationInitParameters `json:"initProvider,omitempty"`
}

// CertificateValidationStatus defines the observed state of CertificateValidation.
type CertificateValidationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateValidationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CertificateValidation is the Schema for the CertificateValidations API. Waits for and checks successful validation of an ACM certificate.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CertificateValidation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateValidationSpec   `json:"spec"`
	Status            CertificateValidationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateValidationList contains a list of CertificateValidations
type CertificateValidationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateValidation `json:"items"`
}

// Repository type metadata.
var (
	CertificateValidation_Kind             = "CertificateValidation"
	CertificateValidation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateValidation_Kind}.String()
	CertificateValidation_KindAPIVersion   = CertificateValidation_Kind + "." + CRDGroupVersion.String()
	CertificateValidation_GroupVersionKind = CRDGroupVersion.WithKind(CertificateValidation_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateValidation{}, &CertificateValidationList{})
}
