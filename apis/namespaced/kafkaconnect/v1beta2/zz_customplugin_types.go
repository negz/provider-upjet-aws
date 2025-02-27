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

type CustomPluginInitParameters_2 struct {

	// The type of the plugin file. Allowed values are ZIP and JAR.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// A summary description of the custom plugin.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Information about the location of a custom plugin. See location Block for details.
	Location *LocationInitParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CustomPluginObservation_2 struct {

	// the Amazon Resource Name (ARN) of the custom plugin.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The type of the plugin file. Allowed values are ZIP and JAR.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// A summary description of the custom plugin.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// an ID of the latest successfully created revision of the custom plugin.
	LatestRevision *float64 `json:"latestRevision,omitempty" tf:"latest_revision,omitempty"`

	// Information about the location of a custom plugin. See location Block for details.
	Location *LocationObservation `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the custom plugin..
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// the state of the custom plugin.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CustomPluginParameters_2 struct {

	// The type of the plugin file. Allowed values are ZIP and JAR.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// A summary description of the custom plugin.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Information about the location of a custom plugin. See location Block for details.
	// +kubebuilder:validation:Optional
	Location *LocationParameters `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the custom plugin..
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LocationInitParameters struct {

	// Information of the plugin file stored in Amazon S3. See s3 Block for details..
	S3 *LocationS3InitParameters `json:"s3,omitempty" tf:"s3,omitempty"`
}

type LocationObservation struct {

	// Information of the plugin file stored in Amazon S3. See s3 Block for details..
	S3 *LocationS3Observation `json:"s3,omitempty" tf:"s3,omitempty"`
}

type LocationParameters struct {

	// Information of the plugin file stored in Amazon S3. See s3 Block for details..
	// +kubebuilder:validation:Optional
	S3 *LocationS3Parameters `json:"s3" tf:"s3,omitempty"`
}

type LocationS3InitParameters struct {

	// The Amazon Resource Name (ARN) of an S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnRef *v1.Reference `json:"bucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnSelector *v1.Selector `json:"bucketArnSelector,omitempty" tf:"-"`

	// The file key for an object in an S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Object
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("key",false)
	FileKey *string `json:"fileKey,omitempty" tf:"file_key,omitempty"`

	// Reference to a Object in s3 to populate fileKey.
	// +kubebuilder:validation:Optional
	FileKeyRef *v1.Reference `json:"fileKeyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate fileKey.
	// +kubebuilder:validation:Optional
	FileKeySelector *v1.Selector `json:"fileKeySelector,omitempty" tf:"-"`

	// The version of an object in an S3 bucket.
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`
}

type LocationS3Observation struct {

	// The Amazon Resource Name (ARN) of an S3 bucket.
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// The file key for an object in an S3 bucket.
	FileKey *string `json:"fileKey,omitempty" tf:"file_key,omitempty"`

	// The version of an object in an S3 bucket.
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`
}

type LocationS3Parameters struct {

	// The Amazon Resource Name (ARN) of an S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnRef *v1.Reference `json:"bucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnSelector *v1.Selector `json:"bucketArnSelector,omitempty" tf:"-"`

	// The file key for an object in an S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Object
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("key",false)
	// +kubebuilder:validation:Optional
	FileKey *string `json:"fileKey,omitempty" tf:"file_key,omitempty"`

	// Reference to a Object in s3 to populate fileKey.
	// +kubebuilder:validation:Optional
	FileKeyRef *v1.Reference `json:"fileKeyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate fileKey.
	// +kubebuilder:validation:Optional
	FileKeySelector *v1.Selector `json:"fileKeySelector,omitempty" tf:"-"`

	// The version of an object in an S3 bucket.
	// +kubebuilder:validation:Optional
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`
}

// CustomPluginSpec defines the desired state of CustomPlugin
type CustomPluginSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomPluginParameters_2 `json:"forProvider"`
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
	InitProvider CustomPluginInitParameters_2 `json:"initProvider,omitempty"`
}

// CustomPluginStatus defines the observed state of CustomPlugin.
type CustomPluginStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomPluginObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CustomPlugin is the Schema for the CustomPlugins API. Provides an Amazon MSK Connect custom plugin resource. This resource can be Created, Observed and Deleted, but not Updated. AWS does not currently provide update APIs.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type CustomPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.contentType) || (has(self.initProvider) && has(self.initProvider.contentType))",message="spec.forProvider.contentType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   CustomPluginSpec   `json:"spec"`
	Status CustomPluginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomPluginList contains a list of CustomPlugins
type CustomPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomPlugin `json:"items"`
}

// Repository type metadata.
var (
	CustomPlugin_Kind             = "CustomPlugin"
	CustomPlugin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomPlugin_Kind}.String()
	CustomPlugin_KindAPIVersion   = CustomPlugin_Kind + "." + CRDGroupVersion.String()
	CustomPlugin_GroupVersionKind = CRDGroupVersion.WithKind(CustomPlugin_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomPlugin{}, &CustomPluginList{})
}
