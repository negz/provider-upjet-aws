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

type BuildInitParameters struct {

	// Name of the build
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system that the game server binaries are built to run on. Valid values: WINDOWS_2012, AMAZON_LINUX, AMAZON_LINUX_2, WINDOWS_2016, AMAZON_LINUX_2023.
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// Information indicating where your game build files are stored. See below.
	StorageLocation *StorageLocationInitParameters `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Version that is associated with this build.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type BuildObservation struct {

	// GameLift Build ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// GameLift Build ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the build
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system that the game server binaries are built to run on. Valid values: WINDOWS_2012, AMAZON_LINUX, AMAZON_LINUX_2, WINDOWS_2016, AMAZON_LINUX_2023.
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// Information indicating where your game build files are stored. See below.
	StorageLocation *StorageLocationObservation `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Version that is associated with this build.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type BuildParameters struct {

	// Name of the build
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system that the game server binaries are built to run on. Valid values: WINDOWS_2012, AMAZON_LINUX, AMAZON_LINUX_2, WINDOWS_2016, AMAZON_LINUX_2023.
	// +kubebuilder:validation:Optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Information indicating where your game build files are stored. See below.
	// +kubebuilder:validation:Optional
	StorageLocation *StorageLocationParameters `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Version that is associated with this build.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type StorageLocationInitParameters struct {

	// Name of your S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Name of the zip file containing your build files.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Object
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("key",false)
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Reference to a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeyRef *v1.Reference `json:"keyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeySelector *v1.Selector `json:"keySelector,omitempty" tf:"-"`

	// A specific version of the file. If not set, the latest version of the file is retrieved.
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`

	// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

type StorageLocationObservation struct {

	// Name of your S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Name of the zip file containing your build files.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A specific version of the file. If not set, the latest version of the file is retrieved.
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`

	// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type StorageLocationParameters struct {

	// Name of your S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Name of the zip file containing your build files.
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

	// A specific version of the file. If not set, the latest version of the file is retrieved.
	// +kubebuilder:validation:Optional
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`

	// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
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
}

// BuildSpec defines the desired state of Build
type BuildSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BuildParameters `json:"forProvider"`
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
	InitProvider BuildInitParameters `json:"initProvider,omitempty"`
}

// BuildStatus defines the observed state of Build.
type BuildStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BuildObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Build is the Schema for the Builds API. Provides a GameLift Build resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.operatingSystem) || (has(self.initProvider) && has(self.initProvider.operatingSystem))",message="spec.forProvider.operatingSystem is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageLocation) || (has(self.initProvider) && has(self.initProvider.storageLocation))",message="spec.forProvider.storageLocation is a required parameter"
	Spec   BuildSpec   `json:"spec"`
	Status BuildStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BuildList contains a list of Builds
type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Build `json:"items"`
}

// Repository type metadata.
var (
	Build_Kind             = "Build"
	Build_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Build_Kind}.String()
	Build_KindAPIVersion   = Build_Kind + "." + CRDGroupVersion.String()
	Build_GroupVersionKind = CRDGroupVersion.WithKind(Build_Kind)
)

func init() {
	SchemeBuilder.Register(&Build{}, &BuildList{})
}
