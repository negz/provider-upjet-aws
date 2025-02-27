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

type BucketPublicAccessBlockInitParameters struct {

	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to false. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to false. Enabling this setting does not affect the existing bucket policy. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// S3 Bucket to which this Public Access Block configuration should be applied.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to false. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to false. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type BucketPublicAccessBlockObservation struct {

	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to false. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to false. Enabling this setting does not affect the existing bucket policy. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Name of the S3 bucket the configuration is attached to
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to false. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to false. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type BucketPublicAccessBlockParameters struct {

	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to false. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	// +kubebuilder:validation:Optional
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to false. Enabling this setting does not affect the existing bucket policy. When set to true causes Amazon S3 to:
	// +kubebuilder:validation:Optional
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// S3 Bucket to which this Public Access Block configuration should be applied.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to false. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	// +kubebuilder:validation:Optional
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to false. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	// +kubebuilder:validation:Optional
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

// BucketPublicAccessBlockSpec defines the desired state of BucketPublicAccessBlock
type BucketPublicAccessBlockSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketPublicAccessBlockParameters `json:"forProvider"`
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
	InitProvider BucketPublicAccessBlockInitParameters `json:"initProvider,omitempty"`
}

// BucketPublicAccessBlockStatus defines the observed state of BucketPublicAccessBlock.
type BucketPublicAccessBlockStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketPublicAccessBlockObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketPublicAccessBlock is the Schema for the BucketPublicAccessBlocks API. Manages S3 bucket-level Public Access Block Configuration
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketPublicAccessBlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketPublicAccessBlockSpec   `json:"spec"`
	Status            BucketPublicAccessBlockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketPublicAccessBlockList contains a list of BucketPublicAccessBlocks
type BucketPublicAccessBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketPublicAccessBlock `json:"items"`
}

// Repository type metadata.
var (
	BucketPublicAccessBlock_Kind             = "BucketPublicAccessBlock"
	BucketPublicAccessBlock_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketPublicAccessBlock_Kind}.String()
	BucketPublicAccessBlock_KindAPIVersion   = BucketPublicAccessBlock_Kind + "." + CRDGroupVersion.String()
	BucketPublicAccessBlock_GroupVersionKind = CRDGroupVersion.WithKind(BucketPublicAccessBlock_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketPublicAccessBlock{}, &BucketPublicAccessBlockList{})
}
