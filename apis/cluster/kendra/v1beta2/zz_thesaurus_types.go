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

type ThesaurusInitParameters struct {

	// The description for a thesaurus.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the index for a thesaurus.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kendra/v1beta2.Index
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`

	// Reference to a Index in kendra to populate indexId.
	// +kubebuilder:validation:Optional
	IndexIDRef *v1.Reference `json:"indexIdRef,omitempty" tf:"-"`

	// Selector for a Index in kendra to populate indexId.
	// +kubebuilder:validation:Optional
	IndexIDSelector *v1.Selector `json:"indexIdSelector,omitempty" tf:"-"`

	// The name for the thesaurus.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The S3 path where your thesaurus file sits in S3. Detailed below.
	SourceS3Path *ThesaurusSourceS3PathInitParameters `json:"sourceS3Path,omitempty" tf:"source_s3_path,omitempty"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ThesaurusObservation struct {

	// ARN of the thesaurus.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description for a thesaurus.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique identifiers of the thesaurus and index separated by a slash (/).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identifier of the index for a thesaurus.
	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`

	// The name for the thesaurus.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The S3 path where your thesaurus file sits in S3. Detailed below.
	SourceS3Path *ThesaurusSourceS3PathObservation `json:"sourceS3Path,omitempty" tf:"source_s3_path,omitempty"`

	// The current status of the thesaurus.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The unique identifiers of the thesaurus and index separated by a slash (/).
	ThesaurusID *string `json:"thesaurusId,omitempty" tf:"thesaurus_id,omitempty"`
}

type ThesaurusParameters struct {

	// The description for a thesaurus.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the index for a thesaurus.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kendra/v1beta2.Index
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`

	// Reference to a Index in kendra to populate indexId.
	// +kubebuilder:validation:Optional
	IndexIDRef *v1.Reference `json:"indexIdRef,omitempty" tf:"-"`

	// Selector for a Index in kendra to populate indexId.
	// +kubebuilder:validation:Optional
	IndexIDSelector *v1.Selector `json:"indexIdSelector,omitempty" tf:"-"`

	// The name for the thesaurus.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The S3 path where your thesaurus file sits in S3. Detailed below.
	// +kubebuilder:validation:Optional
	SourceS3Path *ThesaurusSourceS3PathParameters `json:"sourceS3Path,omitempty" tf:"source_s3_path,omitempty"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ThesaurusSourceS3PathInitParameters struct {

	// The name of the S3 bucket that contains the file.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The name of the file.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Object
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("key",false)
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Reference to a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeyRef *v1.Reference `json:"keyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeySelector *v1.Selector `json:"keySelector,omitempty" tf:"-"`
}

type ThesaurusSourceS3PathObservation struct {

	// The name of the S3 bucket that contains the file.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The name of the file.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type ThesaurusSourceS3PathParameters struct {

	// The name of the S3 bucket that contains the file.
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

	// The name of the file.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Object
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("key",false)
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Reference to a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeyRef *v1.Reference `json:"keyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeySelector *v1.Selector `json:"keySelector,omitempty" tf:"-"`
}

// ThesaurusSpec defines the desired state of Thesaurus
type ThesaurusSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ThesaurusParameters `json:"forProvider"`
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
	InitProvider ThesaurusInitParameters `json:"initProvider,omitempty"`
}

// ThesaurusStatus defines the observed state of Thesaurus.
type ThesaurusStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ThesaurusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Thesaurus is the Schema for the Thesauruss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws},path=thesaurus
type Thesaurus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceS3Path) || (has(self.initProvider) && has(self.initProvider.sourceS3Path))",message="spec.forProvider.sourceS3Path is a required parameter"
	Spec   ThesaurusSpec   `json:"spec"`
	Status ThesaurusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThesaurusList contains a list of Thesauruss
type ThesaurusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Thesaurus `json:"items"`
}

// Repository type metadata.
var (
	Thesaurus_Kind             = "Thesaurus"
	Thesaurus_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Thesaurus_Kind}.String()
	Thesaurus_KindAPIVersion   = Thesaurus_Kind + "." + CRDGroupVersion.String()
	Thesaurus_GroupVersionKind = CRDGroupVersion.WithKind(Thesaurus_Kind)
)

func init() {
	SchemeBuilder.Register(&Thesaurus{}, &ThesaurusList{})
}
