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

type ImagePipelineImageScanningConfigurationInitParameters struct {

	// Configuration block with ECR configuration for image scanning. Detailed below.
	EcrConfiguration []ImageScanningConfigurationEcrConfigurationInitParameters `json:"ecrConfiguration,omitempty" tf:"ecr_configuration,omitempty"`

	// Whether image scans are enabled. Defaults to false.
	ImageScanningEnabled *bool `json:"imageScanningEnabled,omitempty" tf:"image_scanning_enabled,omitempty"`
}

type ImagePipelineImageScanningConfigurationObservation struct {

	// Configuration block with ECR configuration for image scanning. Detailed below.
	EcrConfiguration []ImageScanningConfigurationEcrConfigurationObservation `json:"ecrConfiguration,omitempty" tf:"ecr_configuration,omitempty"`

	// Whether image scans are enabled. Defaults to false.
	ImageScanningEnabled *bool `json:"imageScanningEnabled,omitempty" tf:"image_scanning_enabled,omitempty"`
}

type ImagePipelineImageScanningConfigurationParameters struct {

	// Configuration block with ECR configuration for image scanning. Detailed below.
	// +kubebuilder:validation:Optional
	EcrConfiguration []ImageScanningConfigurationEcrConfigurationParameters `json:"ecrConfiguration,omitempty" tf:"ecr_configuration,omitempty"`

	// Whether image scans are enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	ImageScanningEnabled *bool `json:"imageScanningEnabled,omitempty" tf:"image_scanning_enabled,omitempty"`
}

type ImagePipelineImageTestsConfigurationInitParameters struct {

	// Whether image tests are enabled. Defaults to true.
	ImageTestsEnabled *bool `json:"imageTestsEnabled,omitempty" tf:"image_tests_enabled,omitempty"`

	// Number of minutes before image tests time out. Valid values are between 60 and 1440. Defaults to 720.
	TimeoutMinutes *float64 `json:"timeoutMinutes,omitempty" tf:"timeout_minutes,omitempty"`
}

type ImagePipelineImageTestsConfigurationObservation struct {

	// Whether image tests are enabled. Defaults to true.
	ImageTestsEnabled *bool `json:"imageTestsEnabled,omitempty" tf:"image_tests_enabled,omitempty"`

	// Number of minutes before image tests time out. Valid values are between 60 and 1440. Defaults to 720.
	TimeoutMinutes *float64 `json:"timeoutMinutes,omitempty" tf:"timeout_minutes,omitempty"`
}

type ImagePipelineImageTestsConfigurationParameters struct {

	// Whether image tests are enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	ImageTestsEnabled *bool `json:"imageTestsEnabled,omitempty" tf:"image_tests_enabled,omitempty"`

	// Number of minutes before image tests time out. Valid values are between 60 and 1440. Defaults to 720.
	// +kubebuilder:validation:Optional
	TimeoutMinutes *float64 `json:"timeoutMinutes,omitempty" tf:"timeout_minutes,omitempty"`
}

type ImagePipelineInitParameters struct {

	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn *string `json:"containerRecipeArn,omitempty" tf:"container_recipe_arn,omitempty"`

	// Description of the image pipeline.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn *string `json:"distributionConfigurationArn,omitempty" tf:"distribution_configuration_arn,omitempty"`

	// Whether additional information about the image being created is collected. Defaults to true.
	EnhancedImageMetadataEnabled *bool `json:"enhancedImageMetadataEnabled,omitempty" tf:"enhanced_image_metadata_enabled,omitempty"`

	// Amazon Resource Name (ARN) of the image recipe.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.ImageRecipe
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ImageRecipeArn *string `json:"imageRecipeArn,omitempty" tf:"image_recipe_arn,omitempty"`

	// Reference to a ImageRecipe in imagebuilder to populate imageRecipeArn.
	// +kubebuilder:validation:Optional
	ImageRecipeArnRef *v1.Reference `json:"imageRecipeArnRef,omitempty" tf:"-"`

	// Selector for a ImageRecipe in imagebuilder to populate imageRecipeArn.
	// +kubebuilder:validation:Optional
	ImageRecipeArnSelector *v1.Selector `json:"imageRecipeArnSelector,omitempty" tf:"-"`

	// Configuration block with image scanning configuration. Detailed below.
	ImageScanningConfiguration []ImagePipelineImageScanningConfigurationInitParameters `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration,omitempty"`

	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration []ImagePipelineImageTestsConfigurationInitParameters `json:"imageTestsConfiguration,omitempty" tf:"image_tests_configuration,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.InfrastructureConfiguration
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	InfrastructureConfigurationArn *string `json:"infrastructureConfigurationArn,omitempty" tf:"infrastructure_configuration_arn,omitempty"`

	// Reference to a InfrastructureConfiguration in imagebuilder to populate infrastructureConfigurationArn.
	// +kubebuilder:validation:Optional
	InfrastructureConfigurationArnRef *v1.Reference `json:"infrastructureConfigurationArnRef,omitempty" tf:"-"`

	// Selector for a InfrastructureConfiguration in imagebuilder to populate infrastructureConfigurationArn.
	// +kubebuilder:validation:Optional
	InfrastructureConfigurationArnSelector *v1.Selector `json:"infrastructureConfigurationArnSelector,omitempty" tf:"-"`

	// Name of the image pipeline.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration block with schedule settings. Detailed below.
	Schedule []ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Status of the image pipeline. Valid values are DISABLED and ENABLED. Defaults to ENABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ImagePipelineObservation struct {

	// Amazon Resource Name (ARN) of the image pipeline.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn *string `json:"containerRecipeArn,omitempty" tf:"container_recipe_arn,omitempty"`

	// Date the image pipeline was created.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// Date the image pipeline was last run.
	DateLastRun *string `json:"dateLastRun,omitempty" tf:"date_last_run,omitempty"`

	// Date the image pipeline will run next.
	DateNextRun *string `json:"dateNextRun,omitempty" tf:"date_next_run,omitempty"`

	// Date the image pipeline was updated.
	DateUpdated *string `json:"dateUpdated,omitempty" tf:"date_updated,omitempty"`

	// Description of the image pipeline.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn *string `json:"distributionConfigurationArn,omitempty" tf:"distribution_configuration_arn,omitempty"`

	// Whether additional information about the image being created is collected. Defaults to true.
	EnhancedImageMetadataEnabled *bool `json:"enhancedImageMetadataEnabled,omitempty" tf:"enhanced_image_metadata_enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn *string `json:"imageRecipeArn,omitempty" tf:"image_recipe_arn,omitempty"`

	// Configuration block with image scanning configuration. Detailed below.
	ImageScanningConfiguration []ImagePipelineImageScanningConfigurationObservation `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration,omitempty"`

	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration []ImagePipelineImageTestsConfigurationObservation `json:"imageTestsConfiguration,omitempty" tf:"image_tests_configuration,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn *string `json:"infrastructureConfigurationArn,omitempty" tf:"infrastructure_configuration_arn,omitempty"`

	// Name of the image pipeline.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Platform of the image pipeline.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Configuration block with schedule settings. Detailed below.
	Schedule []ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Status of the image pipeline. Valid values are DISABLED and ENABLED. Defaults to ENABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ImagePipelineParameters struct {

	// Amazon Resource Name (ARN) of the container recipe.
	// +kubebuilder:validation:Optional
	ContainerRecipeArn *string `json:"containerRecipeArn,omitempty" tf:"container_recipe_arn,omitempty"`

	// Description of the image pipeline.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	// +kubebuilder:validation:Optional
	DistributionConfigurationArn *string `json:"distributionConfigurationArn,omitempty" tf:"distribution_configuration_arn,omitempty"`

	// Whether additional information about the image being created is collected. Defaults to true.
	// +kubebuilder:validation:Optional
	EnhancedImageMetadataEnabled *bool `json:"enhancedImageMetadataEnabled,omitempty" tf:"enhanced_image_metadata_enabled,omitempty"`

	// Amazon Resource Name (ARN) of the image recipe.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.ImageRecipe
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ImageRecipeArn *string `json:"imageRecipeArn,omitempty" tf:"image_recipe_arn,omitempty"`

	// Reference to a ImageRecipe in imagebuilder to populate imageRecipeArn.
	// +kubebuilder:validation:Optional
	ImageRecipeArnRef *v1.Reference `json:"imageRecipeArnRef,omitempty" tf:"-"`

	// Selector for a ImageRecipe in imagebuilder to populate imageRecipeArn.
	// +kubebuilder:validation:Optional
	ImageRecipeArnSelector *v1.Selector `json:"imageRecipeArnSelector,omitempty" tf:"-"`

	// Configuration block with image scanning configuration. Detailed below.
	// +kubebuilder:validation:Optional
	ImageScanningConfiguration []ImagePipelineImageScanningConfigurationParameters `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration,omitempty"`

	// Configuration block with image tests configuration. Detailed below.
	// +kubebuilder:validation:Optional
	ImageTestsConfiguration []ImagePipelineImageTestsConfigurationParameters `json:"imageTestsConfiguration,omitempty" tf:"image_tests_configuration,omitempty"`

	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.InfrastructureConfiguration
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	InfrastructureConfigurationArn *string `json:"infrastructureConfigurationArn,omitempty" tf:"infrastructure_configuration_arn,omitempty"`

	// Reference to a InfrastructureConfiguration in imagebuilder to populate infrastructureConfigurationArn.
	// +kubebuilder:validation:Optional
	InfrastructureConfigurationArnRef *v1.Reference `json:"infrastructureConfigurationArnRef,omitempty" tf:"-"`

	// Selector for a InfrastructureConfiguration in imagebuilder to populate infrastructureConfigurationArn.
	// +kubebuilder:validation:Optional
	InfrastructureConfigurationArnSelector *v1.Selector `json:"infrastructureConfigurationArnSelector,omitempty" tf:"-"`

	// Name of the image pipeline.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block with schedule settings. Detailed below.
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Status of the image pipeline. Valid values are DISABLED and ENABLED. Defaults to ENABLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ImageScanningConfigurationEcrConfigurationInitParameters struct {

	// Key-value map of resource tags.
	// +listType=set
	ContainerTags []*string `json:"containerTags,omitempty" tf:"container_tags,omitempty"`

	// The name of the repository to scan
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`
}

type ImageScanningConfigurationEcrConfigurationObservation struct {

	// Key-value map of resource tags.
	// +listType=set
	ContainerTags []*string `json:"containerTags,omitempty" tf:"container_tags,omitempty"`

	// The name of the repository to scan
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`
}

type ImageScanningConfigurationEcrConfigurationParameters struct {

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +listType=set
	ContainerTags []*string `json:"containerTags,omitempty" tf:"container_tags,omitempty"`

	// The name of the repository to scan
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`
}

type ScheduleInitParameters struct {

	// Condition when the pipeline should trigger a new image build. Valid values are EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE and EXPRESSION_MATCH_ONLY. Defaults to EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE.
	PipelineExecutionStartCondition *string `json:"pipelineExecutionStartCondition,omitempty" tf:"pipeline_execution_start_condition,omitempty"`

	// Cron expression of how often the pipeline start condition is evaluated. For example, cron(0 0 * * ? *) is evaluated every day at midnight UTC. Configurations using the five field syntax that was previously accepted by the API, such as cron(0 0 * * *), must be updated to the six field syntax. For more information, see the Image Builder User Guide.
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`

	// The timezone that applies to the scheduling expression. For example, "Etc/UTC", "America/Los_Angeles" in the IANA timezone format. If not specified this defaults to UTC.
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type ScheduleObservation struct {

	// Condition when the pipeline should trigger a new image build. Valid values are EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE and EXPRESSION_MATCH_ONLY. Defaults to EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE.
	PipelineExecutionStartCondition *string `json:"pipelineExecutionStartCondition,omitempty" tf:"pipeline_execution_start_condition,omitempty"`

	// Cron expression of how often the pipeline start condition is evaluated. For example, cron(0 0 * * ? *) is evaluated every day at midnight UTC. Configurations using the five field syntax that was previously accepted by the API, such as cron(0 0 * * *), must be updated to the six field syntax. For more information, see the Image Builder User Guide.
	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression,omitempty"`

	// The timezone that applies to the scheduling expression. For example, "Etc/UTC", "America/Los_Angeles" in the IANA timezone format. If not specified this defaults to UTC.
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type ScheduleParameters struct {

	// Condition when the pipeline should trigger a new image build. Valid values are EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE and EXPRESSION_MATCH_ONLY. Defaults to EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE.
	// +kubebuilder:validation:Optional
	PipelineExecutionStartCondition *string `json:"pipelineExecutionStartCondition,omitempty" tf:"pipeline_execution_start_condition,omitempty"`

	// Cron expression of how often the pipeline start condition is evaluated. For example, cron(0 0 * * ? *) is evaluated every day at midnight UTC. Configurations using the five field syntax that was previously accepted by the API, such as cron(0 0 * * *), must be updated to the six field syntax. For more information, see the Image Builder User Guide.
	// +kubebuilder:validation:Optional
	ScheduleExpression *string `json:"scheduleExpression" tf:"schedule_expression,omitempty"`

	// The timezone that applies to the scheduling expression. For example, "Etc/UTC", "America/Los_Angeles" in the IANA timezone format. If not specified this defaults to UTC.
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

// ImagePipelineSpec defines the desired state of ImagePipeline
type ImagePipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImagePipelineParameters `json:"forProvider"`
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
	InitProvider ImagePipelineInitParameters `json:"initProvider,omitempty"`
}

// ImagePipelineStatus defines the observed state of ImagePipeline.
type ImagePipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImagePipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImagePipeline is the Schema for the ImagePipelines API. Manages an Image Builder Image Pipeline
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ImagePipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ImagePipelineSpec   `json:"spec"`
	Status ImagePipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImagePipelineList contains a list of ImagePipelines
type ImagePipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagePipeline `json:"items"`
}

// Repository type metadata.
var (
	ImagePipeline_Kind             = "ImagePipeline"
	ImagePipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImagePipeline_Kind}.String()
	ImagePipeline_KindAPIVersion   = ImagePipeline_Kind + "." + CRDGroupVersion.String()
	ImagePipeline_GroupVersionKind = CRDGroupVersion.WithKind(ImagePipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&ImagePipeline{}, &ImagePipelineList{})
}
