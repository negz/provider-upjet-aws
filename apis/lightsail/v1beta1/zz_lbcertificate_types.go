// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DomainValidationRecordsInitParameters struct {
}

type DomainValidationRecordsObservation struct {

	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The SSL/TLS certificate name.
	ResourceRecordName *string `json:"resourceRecordName,omitempty" tf:"resource_record_name,omitempty"`

	ResourceRecordType *string `json:"resourceRecordType,omitempty" tf:"resource_record_type,omitempty"`

	ResourceRecordValue *string `json:"resourceRecordValue,omitempty" tf:"resource_record_value,omitempty"`
}

type DomainValidationRecordsParameters struct {
}

type LBCertificateInitParameters struct {

	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Set of domains that should be SANs in the issued certificate. domain_name attribute is automatically added as a Subject Alternative Name.
	// +listType=set
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`
}

type LBCertificateObservation struct {

	// The ARN of the lightsail certificate.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The timestamp when the instance was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	DomainValidationRecords []DomainValidationRecordsObservation `json:"domainValidationRecords,omitempty" tf:"domain_validation_records,omitempty"`

	// A combination of attributes to create a unique id: lb_name,name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The load balancer name where you want to create the SSL/TLS certificate.
	LBName *string `json:"lbName,omitempty" tf:"lb_name,omitempty"`

	// Set of domains that should be SANs in the issued certificate. domain_name attribute is automatically added as a Subject Alternative Name.
	// +listType=set
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	SupportCode *string `json:"supportCode,omitempty" tf:"support_code,omitempty"`
}

type LBCertificateParameters struct {

	// The domain name (e.g., example.com) for your SSL/TLS certificate.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The load balancer name where you want to create the SSL/TLS certificate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lightsail/v1beta1.LB
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LBName *string `json:"lbName,omitempty" tf:"lb_name,omitempty"`

	// Reference to a LB in lightsail to populate lbName.
	// +kubebuilder:validation:Optional
	LBNameRef *v1.Reference `json:"lbNameRef,omitempty" tf:"-"`

	// Selector for a LB in lightsail to populate lbName.
	// +kubebuilder:validation:Optional
	LBNameSelector *v1.Selector `json:"lbNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Set of domains that should be SANs in the issued certificate. domain_name attribute is automatically added as a Subject Alternative Name.
	// +kubebuilder:validation:Optional
	// +listType=set
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`
}

// LBCertificateSpec defines the desired state of LBCertificate
type LBCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBCertificateParameters `json:"forProvider"`
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
	InitProvider LBCertificateInitParameters `json:"initProvider,omitempty"`
}

// LBCertificateStatus defines the observed state of LBCertificate.
type LBCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LBCertificate is the Schema for the LBCertificates API. Provides a Lightsail Load Balancer
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBCertificateSpec   `json:"spec"`
	Status            LBCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBCertificateList contains a list of LBCertificates
type LBCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBCertificate `json:"items"`
}

// Repository type metadata.
var (
	LBCertificate_Kind             = "LBCertificate"
	LBCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBCertificate_Kind}.String()
	LBCertificate_KindAPIVersion   = LBCertificate_Kind + "." + CRDGroupVersion.String()
	LBCertificate_GroupVersionKind = CRDGroupVersion.WithKind(LBCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&LBCertificate{}, &LBCertificateList{})
}
