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

type RepositoryPolicyInitParameters struct {

	// The policy document. This is a JSON formatted string
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Name of the repository to apply the policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ecr/v1beta2.Repository
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in ecr to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in ecr to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

type RepositoryPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy document. This is a JSON formatted string
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The registry ID where the repository was created.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Name of the repository to apply the policy.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`
}

type RepositoryPolicyParameters struct {

	// The policy document. This is a JSON formatted string
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the repository to apply the policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ecr/v1beta2.Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in ecr to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in ecr to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

// RepositoryPolicySpec defines the desired state of RepositoryPolicy
type RepositoryPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryPolicyParameters `json:"forProvider"`
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
	InitProvider RepositoryPolicyInitParameters `json:"initProvider,omitempty"`
}

// RepositoryPolicyStatus defines the observed state of RepositoryPolicy.
type RepositoryPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RepositoryPolicy is the Schema for the RepositoryPolicys API. Provides an Elastic Container Registry Repository Policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type RepositoryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   RepositoryPolicySpec   `json:"spec"`
	Status RepositoryPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryPolicyList contains a list of RepositoryPolicys
type RepositoryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryPolicy `json:"items"`
}

// Repository type metadata.
var (
	RepositoryPolicy_Kind             = "RepositoryPolicy"
	RepositoryPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryPolicy_Kind}.String()
	RepositoryPolicy_KindAPIVersion   = RepositoryPolicy_Kind + "." + CRDGroupVersion.String()
	RepositoryPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryPolicy{}, &RepositoryPolicyList{})
}
