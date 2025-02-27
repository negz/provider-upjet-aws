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

type AccessPointInitParameters struct {

	// AWS account ID for the owner of the bucket for which you want to create an access point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Name of an AWS Partition S3 General Purpose Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountID *string `json:"bucketAccountId,omitempty" tf:"bucket_account_id,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Name you want to assign to this access point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing policy from your configuration or setting policy to null or an empty string (i.e., policy = "") will not delete the policy since it could have been set by aws_s3control_access_point_policy. To remove the policy, set it to "{}" (an empty JSON document).
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Configuration block to manage the PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration []PublicAccessBlockConfigurationInitParameters `json:"publicAccessBlockConfiguration,omitempty" tf:"public_access_block_configuration,omitempty"`

	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VPCConfiguration []VPCConfigurationInitParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type AccessPointObservation struct {

	// AWS account ID for the owner of the bucket for which you want to create an access point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Alias of the S3 Access Point.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// ARN of the S3 Access Point.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Name of an AWS Partition S3 General Purpose Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountID *string `json:"bucketAccountId,omitempty" tf:"bucket_account_id,omitempty"`

	// DNS domain name of the S3 Access Point in the format name-account_id.s3-accesspoint.region.amazonaws.com.
	// Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// VPC endpoints for the S3 Access Point.
	// +mapType=granular
	Endpoints map[string]*string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy *bool `json:"hasPublicAccessPolicy,omitempty" tf:"has_public_access_policy,omitempty"`

	// For Access Point of an AWS Partition S3 Bucket, the AWS account ID and access point name separated by a colon (:). For S3 on Outposts Bucket, the ARN of the Access Point.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name you want to assign to this access point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether this access point allows access from the public Internet. Values are VPC (the access point doesn't allow access from the public Internet) and Internet (the access point allows access from the public Internet, subject to the access point and bucket access policies).
	NetworkOrigin *string `json:"networkOrigin,omitempty" tf:"network_origin,omitempty"`

	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing policy from your configuration or setting policy to null or an empty string (i.e., policy = "") will not delete the policy since it could have been set by aws_s3control_access_point_policy. To remove the policy, set it to "{}" (an empty JSON document).
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Configuration block to manage the PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration []PublicAccessBlockConfigurationObservation `json:"publicAccessBlockConfiguration,omitempty" tf:"public_access_block_configuration,omitempty"`

	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VPCConfiguration []VPCConfigurationObservation `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type AccessPointParameters struct {

	// AWS account ID for the owner of the bucket for which you want to create an access point.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Name of an AWS Partition S3 General Purpose Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// AWS account ID associated with the S3 bucket associated with this access point.
	// +kubebuilder:validation:Optional
	BucketAccountID *string `json:"bucketAccountId,omitempty" tf:"bucket_account_id,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Name you want to assign to this access point.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing policy from your configuration or setting policy to null or an empty string (i.e., policy = "") will not delete the policy since it could have been set by aws_s3control_access_point_policy. To remove the policy, set it to "{}" (an empty JSON document).
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Configuration block to manage the PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	// +kubebuilder:validation:Optional
	PublicAccessBlockConfiguration []PublicAccessBlockConfigurationParameters `json:"publicAccessBlockConfiguration,omitempty" tf:"public_access_block_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	// +kubebuilder:validation:Optional
	VPCConfiguration []VPCConfigurationParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type PublicAccessBlockConfigurationInitParameters struct {

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type PublicAccessBlockConfigurationObservation struct {

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type PublicAccessBlockConfigurationParameters struct {

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	// +kubebuilder:validation:Optional
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	// +kubebuilder:validation:Optional
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	// +kubebuilder:validation:Optional
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	// +kubebuilder:validation:Optional
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type VPCConfigurationInitParameters struct {

	// This access point will only allow connections from the specified VPC ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type VPCConfigurationObservation struct {

	// This access point will only allow connections from the specified VPC ID.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigurationParameters struct {

	// This access point will only allow connections from the specified VPC ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// AccessPointSpec defines the desired state of AccessPoint
type AccessPointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPointParameters `json:"forProvider"`
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
	InitProvider AccessPointInitParameters `json:"initProvider,omitempty"`
}

// AccessPointStatus defines the observed state of AccessPoint.
type AccessPointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessPoint is the Schema for the AccessPoints API. Manages an S3 Access Point.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AccessPointSpec   `json:"spec"`
	Status AccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPointList contains a list of AccessPoints
type AccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPoint `json:"items"`
}

// Repository type metadata.
var (
	AccessPoint_Kind             = "AccessPoint"
	AccessPoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPoint_Kind}.String()
	AccessPoint_KindAPIVersion   = AccessPoint_Kind + "." + CRDGroupVersion.String()
	AccessPoint_GroupVersionKind = CRDGroupVersion.WithKind(AccessPoint_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPoint{}, &AccessPointList{})
}
