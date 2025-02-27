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

type DomainMailFromInitParameters struct {

	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to UseDefaultValue. See the SES API documentation for more information.
	BehaviorOnMxFailure *string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`

	// Verified domain name or email identity to generate DKIM tokens for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ses/v1beta1.DomainIdentity
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Reference to a DomainIdentity in ses to populate domain.
	// +kubebuilder:validation:Optional
	DomainRef *v1.Reference `json:"domainRef,omitempty" tf:"-"`

	// Selector for a DomainIdentity in ses to populate domain.
	// +kubebuilder:validation:Optional
	DomainSelector *v1.Selector `json:"domainSelector,omitempty" tf:"-"`

	// Subdomain (of above domain) which is to be used as MAIL FROM address
	MailFromDomain *string `json:"mailFromDomain,omitempty" tf:"mail_from_domain,omitempty"`
}

type DomainMailFromObservation struct {

	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to UseDefaultValue. See the SES API documentation for more information.
	BehaviorOnMxFailure *string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`

	// Verified domain name or email identity to generate DKIM tokens for.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The domain name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Subdomain (of above domain) which is to be used as MAIL FROM address
	MailFromDomain *string `json:"mailFromDomain,omitempty" tf:"mail_from_domain,omitempty"`
}

type DomainMailFromParameters struct {

	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to UseDefaultValue. See the SES API documentation for more information.
	// +kubebuilder:validation:Optional
	BehaviorOnMxFailure *string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`

	// Verified domain name or email identity to generate DKIM tokens for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ses/v1beta1.DomainIdentity
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Reference to a DomainIdentity in ses to populate domain.
	// +kubebuilder:validation:Optional
	DomainRef *v1.Reference `json:"domainRef,omitempty" tf:"-"`

	// Selector for a DomainIdentity in ses to populate domain.
	// +kubebuilder:validation:Optional
	DomainSelector *v1.Selector `json:"domainSelector,omitempty" tf:"-"`

	// Subdomain (of above domain) which is to be used as MAIL FROM address
	// +kubebuilder:validation:Optional
	MailFromDomain *string `json:"mailFromDomain,omitempty" tf:"mail_from_domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainMailFromSpec defines the desired state of DomainMailFrom
type DomainMailFromSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainMailFromParameters `json:"forProvider"`
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
	InitProvider DomainMailFromInitParameters `json:"initProvider,omitempty"`
}

// DomainMailFromStatus defines the observed state of DomainMailFrom.
type DomainMailFromStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainMailFromObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DomainMailFrom is the Schema for the DomainMailFroms API. Provides an SES domain MAIL FROM resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainMailFrom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mailFromDomain) || (has(self.initProvider) && has(self.initProvider.mailFromDomain))",message="spec.forProvider.mailFromDomain is a required parameter"
	Spec   DomainMailFromSpec   `json:"spec"`
	Status DomainMailFromStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMailFromList contains a list of DomainMailFroms
type DomainMailFromList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainMailFrom `json:"items"`
}

// Repository type metadata.
var (
	DomainMailFrom_Kind             = "DomainMailFrom"
	DomainMailFrom_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainMailFrom_Kind}.String()
	DomainMailFrom_KindAPIVersion   = DomainMailFrom_Kind + "." + CRDGroupVersion.String()
	DomainMailFrom_GroupVersionKind = CRDGroupVersion.WithKind(DomainMailFrom_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainMailFrom{}, &DomainMailFromList{})
}
