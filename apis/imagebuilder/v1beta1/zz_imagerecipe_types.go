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

type BlockDeviceMappingEBSInitParameters struct {

	// Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Number of Input/Output (I/O) operations per second to provision for an io1 or io2 volume.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Identifier of the EC2 Volume Snapshot.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// For GP3 volumes only. The throughput in MiB/s that the volume supports.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of the volume, in GiB.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of the volume. For example, gp2 or io2.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type BlockDeviceMappingEBSObservation struct {

	// Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Number of Input/Output (I/O) operations per second to provision for an io1 or io2 volume.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Identifier of the EC2 Volume Snapshot.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// For GP3 volumes only. The throughput in MiB/s that the volume supports.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of the volume, in GiB.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of the volume. For example, gp2 or io2.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type BlockDeviceMappingEBSParameters struct {

	// Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
	// +kubebuilder:validation:Optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	// +kubebuilder:validation:Optional
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Number of Input/Output (I/O) operations per second to provision for an io1 or io2 volume.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Identifier of the EC2 Volume Snapshot.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// For GP3 volumes only. The throughput in MiB/s that the volume supports.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of the volume, in GiB.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of the volume. For example, gp2 or io2.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ComponentParameterInitParameters struct {

	// The name of the component parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value for the named component parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ComponentParameterObservation struct {

	// The name of the component parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value for the named component parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ComponentParameterParameters struct {

	// The name of the component parameter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value for the named component parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type ImageRecipeBlockDeviceMappingInitParameters struct {

	// Name of the device. For example, /dev/sda or /dev/xvdb.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Configuration block with Elastic Block Storage (EBS) block device mapping settings. Detailed below.
	EBS []BlockDeviceMappingEBSInitParameters `json:"ebs,omitempty" tf:"ebs,omitempty"`

	// Set to true to remove a mapping from the parent image.
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// Virtual device name. For example, ephemeral0. Instance store volumes are numbered starting from 0.
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type ImageRecipeBlockDeviceMappingObservation struct {

	// Name of the device. For example, /dev/sda or /dev/xvdb.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Configuration block with Elastic Block Storage (EBS) block device mapping settings. Detailed below.
	EBS []BlockDeviceMappingEBSObservation `json:"ebs,omitempty" tf:"ebs,omitempty"`

	// Set to true to remove a mapping from the parent image.
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// Virtual device name. For example, ephemeral0. Instance store volumes are numbered starting from 0.
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type ImageRecipeBlockDeviceMappingParameters struct {

	// Name of the device. For example, /dev/sda or /dev/xvdb.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Configuration block with Elastic Block Storage (EBS) block device mapping settings. Detailed below.
	// +kubebuilder:validation:Optional
	EBS []BlockDeviceMappingEBSParameters `json:"ebs,omitempty" tf:"ebs,omitempty"`

	// Set to true to remove a mapping from the parent image.
	// +kubebuilder:validation:Optional
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// Virtual device name. For example, ephemeral0. Instance store volumes are numbered starting from 0.
	// +kubebuilder:validation:Optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type ImageRecipeComponentInitParameters struct {

	// Amazon Resource Name (ARN) of the Image Builder Component to associate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.Component
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ComponentArn *string `json:"componentArn,omitempty" tf:"component_arn,omitempty"`

	// Reference to a Component in imagebuilder to populate componentArn.
	// +kubebuilder:validation:Optional
	ComponentArnRef *v1.Reference `json:"componentArnRef,omitempty" tf:"-"`

	// Selector for a Component in imagebuilder to populate componentArn.
	// +kubebuilder:validation:Optional
	ComponentArnSelector *v1.Selector `json:"componentArnSelector,omitempty" tf:"-"`

	// Configuration block(s) for parameters to configure the component. Detailed below.
	Parameter []ComponentParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type ImageRecipeComponentObservation struct {

	// Amazon Resource Name (ARN) of the Image Builder Component to associate.
	ComponentArn *string `json:"componentArn,omitempty" tf:"component_arn,omitempty"`

	// Configuration block(s) for parameters to configure the component. Detailed below.
	Parameter []ComponentParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type ImageRecipeComponentParameters struct {

	// Amazon Resource Name (ARN) of the Image Builder Component to associate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.Component
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ComponentArn *string `json:"componentArn,omitempty" tf:"component_arn,omitempty"`

	// Reference to a Component in imagebuilder to populate componentArn.
	// +kubebuilder:validation:Optional
	ComponentArnRef *v1.Reference `json:"componentArnRef,omitempty" tf:"-"`

	// Selector for a Component in imagebuilder to populate componentArn.
	// +kubebuilder:validation:Optional
	ComponentArnSelector *v1.Selector `json:"componentArnSelector,omitempty" tf:"-"`

	// Configuration block(s) for parameters to configure the component. Detailed below.
	// +kubebuilder:validation:Optional
	Parameter []ComponentParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type ImageRecipeInitParameters struct {

	// Configuration block(s) with block device mappings for the image recipe. Detailed below.
	BlockDeviceMapping []ImageRecipeBlockDeviceMappingInitParameters `json:"blockDeviceMapping,omitempty" tf:"block_device_mapping,omitempty"`

	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	Component []ImageRecipeComponentInitParameters `json:"component,omitempty" tf:"component,omitempty"`

	// Description of the image recipe.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the image recipe.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The image recipe uses this image as a base from which to build your customized image. The value can be the base image ARN or an AMI ID.
	ParentImage *string `json:"parentImage,omitempty" tf:"parent_image,omitempty"`

	// Configuration block for the Systems Manager Agent installed by default by Image Builder. Detailed below.
	SystemsManagerAgent []SystemsManagerAgentInitParameters `json:"systemsManagerAgent,omitempty" tf:"systems_manager_agent,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`

	// The semantic version of the image recipe, which specifies the version in the following format, with numeric values in each position to indicate a specific version: major.minor.patch. For example: 1.0.0.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type ImageRecipeObservation struct {

	// Amazon Resource Name (ARN) of the image recipe.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block(s) with block device mappings for the image recipe. Detailed below.
	BlockDeviceMapping []ImageRecipeBlockDeviceMappingObservation `json:"blockDeviceMapping,omitempty" tf:"block_device_mapping,omitempty"`

	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	Component []ImageRecipeComponentObservation `json:"component,omitempty" tf:"component,omitempty"`

	// Date the image recipe was created.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// Description of the image recipe.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the image recipe.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner of the image recipe.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The image recipe uses this image as a base from which to build your customized image. The value can be the base image ARN or an AMI ID.
	ParentImage *string `json:"parentImage,omitempty" tf:"parent_image,omitempty"`

	// Platform of the image recipe.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Configuration block for the Systems Manager Agent installed by default by Image Builder. Detailed below.
	SystemsManagerAgent []SystemsManagerAgentObservation `json:"systemsManagerAgent,omitempty" tf:"systems_manager_agent,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`

	// The semantic version of the image recipe, which specifies the version in the following format, with numeric values in each position to indicate a specific version: major.minor.patch. For example: 1.0.0.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type ImageRecipeParameters struct {

	// Configuration block(s) with block device mappings for the image recipe. Detailed below.
	// +kubebuilder:validation:Optional
	BlockDeviceMapping []ImageRecipeBlockDeviceMappingParameters `json:"blockDeviceMapping,omitempty" tf:"block_device_mapping,omitempty"`

	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	// +kubebuilder:validation:Optional
	Component []ImageRecipeComponentParameters `json:"component,omitempty" tf:"component,omitempty"`

	// Description of the image recipe.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the image recipe.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The image recipe uses this image as a base from which to build your customized image. The value can be the base image ARN or an AMI ID.
	// +kubebuilder:validation:Optional
	ParentImage *string `json:"parentImage,omitempty" tf:"parent_image,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for the Systems Manager Agent installed by default by Image Builder. Detailed below.
	// +kubebuilder:validation:Optional
	SystemsManagerAgent []SystemsManagerAgentParameters `json:"systemsManagerAgent,omitempty" tf:"systems_manager_agent,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	// +kubebuilder:validation:Optional
	UserDataBase64 *string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`

	// The semantic version of the image recipe, which specifies the version in the following format, with numeric values in each position to indicate a specific version: major.minor.patch. For example: 1.0.0.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The working directory to be used during build and test workflows.
	// +kubebuilder:validation:Optional
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type SystemsManagerAgentInitParameters struct {

	// Whether to remove the Systems Manager Agent after the image has been built. Defaults to false.
	UninstallAfterBuild *bool `json:"uninstallAfterBuild,omitempty" tf:"uninstall_after_build,omitempty"`
}

type SystemsManagerAgentObservation struct {

	// Whether to remove the Systems Manager Agent after the image has been built. Defaults to false.
	UninstallAfterBuild *bool `json:"uninstallAfterBuild,omitempty" tf:"uninstall_after_build,omitempty"`
}

type SystemsManagerAgentParameters struct {

	// Whether to remove the Systems Manager Agent after the image has been built. Defaults to false.
	// +kubebuilder:validation:Optional
	UninstallAfterBuild *bool `json:"uninstallAfterBuild" tf:"uninstall_after_build,omitempty"`
}

// ImageRecipeSpec defines the desired state of ImageRecipe
type ImageRecipeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageRecipeParameters `json:"forProvider"`
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
	InitProvider ImageRecipeInitParameters `json:"initProvider,omitempty"`
}

// ImageRecipeStatus defines the observed state of ImageRecipe.
type ImageRecipeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageRecipeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImageRecipe is the Schema for the ImageRecipes API. Manage an Image Builder Image Recipe
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ImageRecipe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.component) || (has(self.initProvider) && has(self.initProvider.component))",message="spec.forProvider.component is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parentImage) || (has(self.initProvider) && has(self.initProvider.parentImage))",message="spec.forProvider.parentImage is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   ImageRecipeSpec   `json:"spec"`
	Status ImageRecipeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageRecipeList contains a list of ImageRecipes
type ImageRecipeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageRecipe `json:"items"`
}

// Repository type metadata.
var (
	ImageRecipe_Kind             = "ImageRecipe"
	ImageRecipe_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageRecipe_Kind}.String()
	ImageRecipe_KindAPIVersion   = ImageRecipe_Kind + "." + CRDGroupVersion.String()
	ImageRecipe_GroupVersionKind = CRDGroupVersion.WithKind(ImageRecipe_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageRecipe{}, &ImageRecipeList{})
}
