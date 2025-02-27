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

type RepositoryPermissionsPolicyInitParameters struct {

	// The account number of the AWS account that owns the domain.
	DomainOwner *string `json:"domainOwner,omitempty" tf:"domain_owner,omitempty"`

	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision *string `json:"policyRevision,omitempty" tf:"policy_revision,omitempty"`
}

type RepositoryPermissionsPolicyObservation struct {

	// The name of the domain on which to set the resource policy.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The account number of the AWS account that owns the domain.
	DomainOwner *string `json:"domainOwner,omitempty" tf:"domain_owner,omitempty"`

	// The ARN of the resource associated with the resource policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision *string `json:"policyRevision,omitempty" tf:"policy_revision,omitempty"`

	// The name of the repository to set the resource policy on.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// The ARN of the resource associated with the resource policy.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type RepositoryPermissionsPolicyParameters struct {

	// The name of the domain on which to set the resource policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codeartifact/v1beta1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("domain",true)
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The account number of the AWS account that owns the domain.
	// +kubebuilder:validation:Optional
	DomainOwner *string `json:"domainOwner,omitempty" tf:"domain_owner,omitempty"`

	// Reference to a Domain in codeartifact to populate domain.
	// +kubebuilder:validation:Optional
	DomainRef *v1.Reference `json:"domainRef,omitempty" tf:"-"`

	// Selector for a Domain in codeartifact to populate domain.
	// +kubebuilder:validation:Optional
	DomainSelector *v1.Selector `json:"domainSelector,omitempty" tf:"-"`

	// A JSON policy string to be set as the access control resource policy on the provided domain.
	// +kubebuilder:validation:Optional
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	// +kubebuilder:validation:Optional
	PolicyRevision *string `json:"policyRevision,omitempty" tf:"policy_revision,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the repository to set the resource policy on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codeartifact/v1beta1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("repository",true)
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in codeartifact to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in codeartifact to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

// RepositoryPermissionsPolicySpec defines the desired state of RepositoryPermissionsPolicy
type RepositoryPermissionsPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryPermissionsPolicyParameters `json:"forProvider"`
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
	InitProvider RepositoryPermissionsPolicyInitParameters `json:"initProvider,omitempty"`
}

// RepositoryPermissionsPolicyStatus defines the observed state of RepositoryPermissionsPolicy.
type RepositoryPermissionsPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryPermissionsPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RepositoryPermissionsPolicy is the Schema for the RepositoryPermissionsPolicys API. Provides a CodeArtifact Repository Permissions Policy resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type RepositoryPermissionsPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyDocument) || (has(self.initProvider) && has(self.initProvider.policyDocument))",message="spec.forProvider.policyDocument is a required parameter"
	Spec   RepositoryPermissionsPolicySpec   `json:"spec"`
	Status RepositoryPermissionsPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryPermissionsPolicyList contains a list of RepositoryPermissionsPolicys
type RepositoryPermissionsPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryPermissionsPolicy `json:"items"`
}

// Repository type metadata.
var (
	RepositoryPermissionsPolicy_Kind             = "RepositoryPermissionsPolicy"
	RepositoryPermissionsPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryPermissionsPolicy_Kind}.String()
	RepositoryPermissionsPolicy_KindAPIVersion   = RepositoryPermissionsPolicy_Kind + "." + CRDGroupVersion.String()
	RepositoryPermissionsPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryPermissionsPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryPermissionsPolicy{}, &RepositoryPermissionsPolicyList{})
}
