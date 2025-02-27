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

type AccountPublicAccessBlockInitParameters struct {

	// AWS account ID to configure.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to false. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to false. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to false. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to false. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type AccountPublicAccessBlockObservation struct {

	// AWS account ID to configure.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to false. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to false. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// AWS account ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to false. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to false. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type AccountPublicAccessBlockParameters struct {

	// AWS account ID to configure.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to false. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	// +kubebuilder:validation:Optional
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to false. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	// +kubebuilder:validation:Optional
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to false. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	// +kubebuilder:validation:Optional
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to false. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	// +kubebuilder:validation:Optional
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

// AccountPublicAccessBlockSpec defines the desired state of AccountPublicAccessBlock
type AccountPublicAccessBlockSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountPublicAccessBlockParameters `json:"forProvider"`
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
	InitProvider AccountPublicAccessBlockInitParameters `json:"initProvider,omitempty"`
}

// AccountPublicAccessBlockStatus defines the observed state of AccountPublicAccessBlock.
type AccountPublicAccessBlockStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountPublicAccessBlockObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccountPublicAccessBlock is the Schema for the AccountPublicAccessBlocks API. Manages S3 account-level Public Access Block Configuration
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type AccountPublicAccessBlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountPublicAccessBlockSpec   `json:"spec"`
	Status            AccountPublicAccessBlockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountPublicAccessBlockList contains a list of AccountPublicAccessBlocks
type AccountPublicAccessBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountPublicAccessBlock `json:"items"`
}

// Repository type metadata.
var (
	AccountPublicAccessBlock_Kind             = "AccountPublicAccessBlock"
	AccountPublicAccessBlock_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountPublicAccessBlock_Kind}.String()
	AccountPublicAccessBlock_KindAPIVersion   = AccountPublicAccessBlock_Kind + "." + CRDGroupVersion.String()
	AccountPublicAccessBlock_GroupVersionKind = CRDGroupVersion.WithKind(AccountPublicAccessBlock_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountPublicAccessBlock{}, &AccountPublicAccessBlockList{})
}
