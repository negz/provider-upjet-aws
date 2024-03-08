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

type SecretPolicyInitParameters struct {

	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Valid JSON document representing a resource policy. Unlike aws_secretsmanager_secret, where policy can be set to "{}" to delete the policy, "{}" is not a valid policy since policy is required.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Secret ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`

	// Reference to a Secret in secretsmanager to populate secretArn.
	// +kubebuilder:validation:Optional
	SecretArnRef *v1.Reference `json:"secretArnRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate secretArn.
	// +kubebuilder:validation:Optional
	SecretArnSelector *v1.Selector `json:"secretArnSelector,omitempty" tf:"-"`
}

type SecretPolicyObservation struct {

	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Amazon Resource Name (ARN) of the secret.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Valid JSON document representing a resource policy. Unlike aws_secretsmanager_secret, where policy can be set to "{}" to delete the policy, "{}" is not a valid policy since policy is required.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Secret ARN.
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`
}

type SecretPolicyParameters struct {

	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	// +kubebuilder:validation:Optional
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Valid JSON document representing a resource policy. Unlike aws_secretsmanager_secret, where policy can be set to "{}" to delete the policy, "{}" is not a valid policy since policy is required.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Secret ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`

	// Reference to a Secret in secretsmanager to populate secretArn.
	// +kubebuilder:validation:Optional
	SecretArnRef *v1.Reference `json:"secretArnRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate secretArn.
	// +kubebuilder:validation:Optional
	SecretArnSelector *v1.Selector `json:"secretArnSelector,omitempty" tf:"-"`
}

// SecretPolicySpec defines the desired state of SecretPolicy
type SecretPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretPolicyParameters `json:"forProvider"`
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
	InitProvider SecretPolicyInitParameters `json:"initProvider,omitempty"`
}

// SecretPolicyStatus defines the observed state of SecretPolicy.
type SecretPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretPolicy is the Schema for the SecretPolicys API. Provides a resource to manage AWS Secrets Manager secret policy
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecretPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   SecretPolicySpec   `json:"spec"`
	Status SecretPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretPolicyList contains a list of SecretPolicys
type SecretPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretPolicy `json:"items"`
}

// Repository type metadata.
var (
	SecretPolicy_Kind             = "SecretPolicy"
	SecretPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretPolicy_Kind}.String()
	SecretPolicy_KindAPIVersion   = SecretPolicy_Kind + "." + CRDGroupVersion.String()
	SecretPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SecretPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretPolicy{}, &SecretPolicyList{})
}
