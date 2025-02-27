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

type RegistryPolicyInitParameters struct {

	// The policy document. This is a JSON formatted string
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type RegistryPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy document. This is a JSON formatted string
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The registry ID where the registry was created.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`
}

type RegistryPolicyParameters struct {

	// The policy document. This is a JSON formatted string
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// RegistryPolicySpec defines the desired state of RegistryPolicy
type RegistryPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryPolicyParameters `json:"forProvider"`
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
	InitProvider RegistryPolicyInitParameters `json:"initProvider,omitempty"`
}

// RegistryPolicyStatus defines the observed state of RegistryPolicy.
type RegistryPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegistryPolicy is the Schema for the RegistryPolicys API. Provides an Elastic Container Registry Policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type RegistryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   RegistryPolicySpec   `json:"spec"`
	Status RegistryPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryPolicyList contains a list of RegistryPolicys
type RegistryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryPolicy `json:"items"`
}

// Repository type metadata.
var (
	RegistryPolicy_Kind             = "RegistryPolicy"
	RegistryPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegistryPolicy_Kind}.String()
	RegistryPolicy_KindAPIVersion   = RegistryPolicy_Kind + "." + CRDGroupVersion.String()
	RegistryPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RegistryPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RegistryPolicy{}, &RegistryPolicyList{})
}
