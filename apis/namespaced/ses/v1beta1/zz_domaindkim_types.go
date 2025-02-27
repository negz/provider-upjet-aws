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

type DomainDKIMInitParameters struct {
}

type DomainDKIMObservation struct {

	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// Find out more about verifying domains in Amazon SES
	// in the AWS SES docs.
	DKIMTokens []*string `json:"dkimTokens,omitempty" tf:"dkim_tokens,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainDKIMParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainDKIMSpec defines the desired state of DomainDKIM
type DomainDKIMSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainDKIMParameters `json:"forProvider"`
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
	InitProvider DomainDKIMInitParameters `json:"initProvider,omitempty"`
}

// DomainDKIMStatus defines the observed state of DomainDKIM.
type DomainDKIMStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainDKIMObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DomainDKIM is the Schema for the DomainDKIMs API. Provides an SES domain DKIM generation resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type DomainDKIM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainDKIMSpec   `json:"spec"`
	Status            DomainDKIMStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainDKIMList contains a list of DomainDKIMs
type DomainDKIMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainDKIM `json:"items"`
}

// Repository type metadata.
var (
	DomainDKIM_Kind             = "DomainDKIM"
	DomainDKIM_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainDKIM_Kind}.String()
	DomainDKIM_KindAPIVersion   = DomainDKIM_Kind + "." + CRDGroupVersion.String()
	DomainDKIM_GroupVersionKind = CRDGroupVersion.WithKind(DomainDKIM_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainDKIM{}, &DomainDKIMList{})
}
