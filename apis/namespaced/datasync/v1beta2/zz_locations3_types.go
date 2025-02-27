// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LocationS3InitParameters struct {

	// A list of DataSync Agent ARNs with which this location will be associated.
	// +listType=set
	AgentArns []*string `json:"agentArns,omitempty" tf:"agent_arns,omitempty"`

	// Amazon Resource Name (ARN) of the S3 Bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	S3BucketArn *string `json:"s3BucketArn,omitempty" tf:"s3_bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate s3BucketArn.
	// +kubebuilder:validation:Optional
	S3BucketArnRef *v1.Reference `json:"s3BucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3BucketArn.
	// +kubebuilder:validation:Optional
	S3BucketArnSelector *v1.Selector `json:"s3BucketArnSelector,omitempty" tf:"-"`

	// Configuration block containing information for connecting to S3.
	S3Config *S3ConfigInitParameters `json:"s3Config,omitempty" tf:"s3_config,omitempty"`

	// The Amazon S3 storage class that you want to store your files in when this location is used as a task destination. Valid values
	S3StorageClass *string `json:"s3StorageClass,omitempty" tf:"s3_storage_class,omitempty"`

	// Prefix to perform actions as source or destination.
	Subdirectory *string `json:"subdirectory,omitempty" tf:"subdirectory,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LocationS3Observation struct {

	// A list of DataSync Agent ARNs with which this location will be associated.
	// +listType=set
	AgentArns []*string `json:"agentArns,omitempty" tf:"agent_arns,omitempty"`

	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Amazon Resource Name (ARN) of the DataSync Location.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Resource Name (ARN) of the S3 Bucket.
	S3BucketArn *string `json:"s3BucketArn,omitempty" tf:"s3_bucket_arn,omitempty"`

	// Configuration block containing information for connecting to S3.
	S3Config *S3ConfigObservation `json:"s3Config,omitempty" tf:"s3_config,omitempty"`

	// The Amazon S3 storage class that you want to store your files in when this location is used as a task destination. Valid values
	S3StorageClass *string `json:"s3StorageClass,omitempty" tf:"s3_storage_class,omitempty"`

	// Prefix to perform actions as source or destination.
	Subdirectory *string `json:"subdirectory,omitempty" tf:"subdirectory,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type LocationS3Parameters struct {

	// A list of DataSync Agent ARNs with which this location will be associated.
	// +kubebuilder:validation:Optional
	// +listType=set
	AgentArns []*string `json:"agentArns,omitempty" tf:"agent_arns,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) of the S3 Bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	S3BucketArn *string `json:"s3BucketArn,omitempty" tf:"s3_bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate s3BucketArn.
	// +kubebuilder:validation:Optional
	S3BucketArnRef *v1.Reference `json:"s3BucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3BucketArn.
	// +kubebuilder:validation:Optional
	S3BucketArnSelector *v1.Selector `json:"s3BucketArnSelector,omitempty" tf:"-"`

	// Configuration block containing information for connecting to S3.
	// +kubebuilder:validation:Optional
	S3Config *S3ConfigParameters `json:"s3Config,omitempty" tf:"s3_config,omitempty"`

	// The Amazon S3 storage class that you want to store your files in when this location is used as a task destination. Valid values
	// +kubebuilder:validation:Optional
	S3StorageClass *string `json:"s3StorageClass,omitempty" tf:"s3_storage_class,omitempty"`

	// Prefix to perform actions as source or destination.
	// +kubebuilder:validation:Optional
	Subdirectory *string `json:"subdirectory,omitempty" tf:"subdirectory,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type S3ConfigInitParameters struct {

	// ARN of the IAM Role used to connect to the S3 Bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	BucketAccessRoleArn *string `json:"bucketAccessRoleArn,omitempty" tf:"bucket_access_role_arn,omitempty"`

	// Reference to a Role in iam to populate bucketAccessRoleArn.
	// +kubebuilder:validation:Optional
	BucketAccessRoleArnRef *v1.Reference `json:"bucketAccessRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate bucketAccessRoleArn.
	// +kubebuilder:validation:Optional
	BucketAccessRoleArnSelector *v1.Selector `json:"bucketAccessRoleArnSelector,omitempty" tf:"-"`
}

type S3ConfigObservation struct {

	// ARN of the IAM Role used to connect to the S3 Bucket.
	BucketAccessRoleArn *string `json:"bucketAccessRoleArn,omitempty" tf:"bucket_access_role_arn,omitempty"`
}

type S3ConfigParameters struct {

	// ARN of the IAM Role used to connect to the S3 Bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	BucketAccessRoleArn *string `json:"bucketAccessRoleArn,omitempty" tf:"bucket_access_role_arn,omitempty"`

	// Reference to a Role in iam to populate bucketAccessRoleArn.
	// +kubebuilder:validation:Optional
	BucketAccessRoleArnRef *v1.Reference `json:"bucketAccessRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate bucketAccessRoleArn.
	// +kubebuilder:validation:Optional
	BucketAccessRoleArnSelector *v1.Selector `json:"bucketAccessRoleArnSelector,omitempty" tf:"-"`
}

// LocationS3Spec defines the desired state of LocationS3
type LocationS3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocationS3Parameters `json:"forProvider"`
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
	InitProvider LocationS3InitParameters `json:"initProvider,omitempty"`
}

// LocationS3Status defines the observed state of LocationS3.
type LocationS3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocationS3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// LocationS3 is the Schema for the LocationS3s API. Manages an AWS DataSync S3 Location
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type LocationS3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.s3Config) || (has(self.initProvider) && has(self.initProvider.s3Config))",message="spec.forProvider.s3Config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subdirectory) || (has(self.initProvider) && has(self.initProvider.subdirectory))",message="spec.forProvider.subdirectory is a required parameter"
	Spec   LocationS3Spec   `json:"spec"`
	Status LocationS3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocationS3List contains a list of LocationS3s
type LocationS3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocationS3 `json:"items"`
}

// Repository type metadata.
var (
	LocationS3_Kind             = "LocationS3"
	LocationS3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocationS3_Kind}.String()
	LocationS3_KindAPIVersion   = LocationS3_Kind + "." + CRDGroupVersion.String()
	LocationS3_GroupVersionKind = CRDGroupVersion.WithKind(LocationS3_Kind)
)

func init() {
	SchemeBuilder.Register(&LocationS3{}, &LocationS3List{})
}
