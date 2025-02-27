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

type ConfigurationRecorderInitParameters struct {

	// Recording group - see below.
	RecordingGroup *RecordingGroupInitParameters `json:"recordingGroup,omitempty" tf:"recording_group,omitempty"`

	// Recording mode - see below.
	RecordingMode *RecordingModeInitParameters `json:"recordingMode,omitempty" tf:"recording_mode,omitempty"`

	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See AWS Docs for more details.
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

type ConfigurationRecorderObservation struct {

	// Name of the recorder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Recording group - see below.
	RecordingGroup *RecordingGroupObservation `json:"recordingGroup,omitempty" tf:"recording_group,omitempty"`

	// Recording mode - see below.
	RecordingMode *RecordingModeObservation `json:"recordingMode,omitempty" tf:"recording_mode,omitempty"`

	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See AWS Docs for more details.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type ConfigurationRecorderParameters struct {

	// Recording group - see below.
	// +kubebuilder:validation:Optional
	RecordingGroup *RecordingGroupParameters `json:"recordingGroup,omitempty" tf:"recording_group,omitempty"`

	// Recording mode - see below.
	// +kubebuilder:validation:Optional
	RecordingMode *RecordingModeParameters `json:"recordingMode,omitempty" tf:"recording_mode,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See AWS Docs for more details.
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

type ExclusionByResourceTypesInitParameters struct {

	// A list that specifies the types of AWS resources for which AWS Config excludes records configuration changes. See relevant part of AWS Docs for available types.
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type ExclusionByResourceTypesObservation struct {

	// A list that specifies the types of AWS resources for which AWS Config excludes records configuration changes. See relevant part of AWS Docs for available types.
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type ExclusionByResourceTypesParameters struct {

	// A list that specifies the types of AWS resources for which AWS Config excludes records configuration changes. See relevant part of AWS Docs for available types.
	// +kubebuilder:validation:Optional
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type RecordingGroupInitParameters struct {

	// Specifies whether AWS Config records configuration changes for every supported type of regional resource (which includes any new type that will become supported in the future). Conflicts with resource_types. Defaults to true.
	AllSupported *bool `json:"allSupported,omitempty" tf:"all_supported,omitempty"`

	// An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.To use this option, you must set the useOnly field of RecordingStrategy to EXCLUSION_BY_RESOURCE_TYPES Requires all_supported = false. Conflicts with resource_types.
	ExclusionByResourceTypes []ExclusionByResourceTypesInitParameters `json:"exclusionByResourceTypes,omitempty" tf:"exclusion_by_resource_types,omitempty"`

	// Specifies whether AWS Config includes all supported types of global resources with the resources that it records. Requires all_supported = true. Conflicts with resource_types.
	IncludeGlobalResourceTypes *bool `json:"includeGlobalResourceTypes,omitempty" tf:"include_global_resource_types,omitempty"`

	// Recording Strategy. Detailed below.
	RecordingStrategy []RecordingStrategyInitParameters `json:"recordingStrategy,omitempty" tf:"recording_strategy,omitempty"`

	// A list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, AWS::EC2::Instance or AWS::CloudTrail::Trail). See relevant part of AWS Docs for available types. In order to use this attribute, all_supported must be set to false.
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type RecordingGroupObservation struct {

	// Specifies whether AWS Config records configuration changes for every supported type of regional resource (which includes any new type that will become supported in the future). Conflicts with resource_types. Defaults to true.
	AllSupported *bool `json:"allSupported,omitempty" tf:"all_supported,omitempty"`

	// An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.To use this option, you must set the useOnly field of RecordingStrategy to EXCLUSION_BY_RESOURCE_TYPES Requires all_supported = false. Conflicts with resource_types.
	ExclusionByResourceTypes []ExclusionByResourceTypesObservation `json:"exclusionByResourceTypes,omitempty" tf:"exclusion_by_resource_types,omitempty"`

	// Specifies whether AWS Config includes all supported types of global resources with the resources that it records. Requires all_supported = true. Conflicts with resource_types.
	IncludeGlobalResourceTypes *bool `json:"includeGlobalResourceTypes,omitempty" tf:"include_global_resource_types,omitempty"`

	// Recording Strategy. Detailed below.
	RecordingStrategy []RecordingStrategyObservation `json:"recordingStrategy,omitempty" tf:"recording_strategy,omitempty"`

	// A list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, AWS::EC2::Instance or AWS::CloudTrail::Trail). See relevant part of AWS Docs for available types. In order to use this attribute, all_supported must be set to false.
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type RecordingGroupParameters struct {

	// Specifies whether AWS Config records configuration changes for every supported type of regional resource (which includes any new type that will become supported in the future). Conflicts with resource_types. Defaults to true.
	// +kubebuilder:validation:Optional
	AllSupported *bool `json:"allSupported,omitempty" tf:"all_supported,omitempty"`

	// An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.To use this option, you must set the useOnly field of RecordingStrategy to EXCLUSION_BY_RESOURCE_TYPES Requires all_supported = false. Conflicts with resource_types.
	// +kubebuilder:validation:Optional
	ExclusionByResourceTypes []ExclusionByResourceTypesParameters `json:"exclusionByResourceTypes,omitempty" tf:"exclusion_by_resource_types,omitempty"`

	// Specifies whether AWS Config includes all supported types of global resources with the resources that it records. Requires all_supported = true. Conflicts with resource_types.
	// +kubebuilder:validation:Optional
	IncludeGlobalResourceTypes *bool `json:"includeGlobalResourceTypes,omitempty" tf:"include_global_resource_types,omitempty"`

	// Recording Strategy. Detailed below.
	// +kubebuilder:validation:Optional
	RecordingStrategy []RecordingStrategyParameters `json:"recordingStrategy,omitempty" tf:"recording_strategy,omitempty"`

	// A list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, AWS::EC2::Instance or AWS::CloudTrail::Trail). See relevant part of AWS Docs for available types. In order to use this attribute, all_supported must be set to false.
	// +kubebuilder:validation:Optional
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type RecordingModeInitParameters struct {

	// Default reecording frequency. CONTINUOUS or DAILY.
	RecordingFrequency *string `json:"recordingFrequency,omitempty" tf:"recording_frequency,omitempty"`

	// Recording mode overrides. Detailed below.
	RecordingModeOverride *RecordingModeOverrideInitParameters `json:"recordingModeOverride,omitempty" tf:"recording_mode_override,omitempty"`
}

type RecordingModeObservation struct {

	// Default reecording frequency. CONTINUOUS or DAILY.
	RecordingFrequency *string `json:"recordingFrequency,omitempty" tf:"recording_frequency,omitempty"`

	// Recording mode overrides. Detailed below.
	RecordingModeOverride *RecordingModeOverrideObservation `json:"recordingModeOverride,omitempty" tf:"recording_mode_override,omitempty"`
}

type RecordingModeOverrideInitParameters struct {

	// A description you provide of the override.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Default reecording frequency. CONTINUOUS or DAILY.
	RecordingFrequency *string `json:"recordingFrequency,omitempty" tf:"recording_frequency,omitempty"`

	// A list that specifies the types of AWS resources for which AWS Config excludes records configuration changes. See relevant part of AWS Docs for available types.
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type RecordingModeOverrideObservation struct {

	// A description you provide of the override.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Default reecording frequency. CONTINUOUS or DAILY.
	RecordingFrequency *string `json:"recordingFrequency,omitempty" tf:"recording_frequency,omitempty"`

	// A list that specifies the types of AWS resources for which AWS Config excludes records configuration changes. See relevant part of AWS Docs for available types.
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type RecordingModeOverrideParameters struct {

	// A description you provide of the override.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Default reecording frequency. CONTINUOUS or DAILY.
	// +kubebuilder:validation:Optional
	RecordingFrequency *string `json:"recordingFrequency" tf:"recording_frequency,omitempty"`

	// A list that specifies the types of AWS resources for which AWS Config excludes records configuration changes. See relevant part of AWS Docs for available types.
	// +kubebuilder:validation:Optional
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes" tf:"resource_types,omitempty"`
}

type RecordingModeParameters struct {

	// Default reecording frequency. CONTINUOUS or DAILY.
	// +kubebuilder:validation:Optional
	RecordingFrequency *string `json:"recordingFrequency,omitempty" tf:"recording_frequency,omitempty"`

	// Recording mode overrides. Detailed below.
	// +kubebuilder:validation:Optional
	RecordingModeOverride *RecordingModeOverrideParameters `json:"recordingModeOverride,omitempty" tf:"recording_mode_override,omitempty"`
}

type RecordingStrategyInitParameters struct {
	UseOnly *string `json:"useOnly,omitempty" tf:"use_only,omitempty"`
}

type RecordingStrategyObservation struct {
	UseOnly *string `json:"useOnly,omitempty" tf:"use_only,omitempty"`
}

type RecordingStrategyParameters struct {

	// +kubebuilder:validation:Optional
	UseOnly *string `json:"useOnly,omitempty" tf:"use_only,omitempty"`
}

// ConfigurationRecorderSpec defines the desired state of ConfigurationRecorder
type ConfigurationRecorderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationRecorderParameters `json:"forProvider"`
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
	InitProvider ConfigurationRecorderInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationRecorderStatus defines the observed state of ConfigurationRecorder.
type ConfigurationRecorderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationRecorderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ConfigurationRecorder is the Schema for the ConfigurationRecorders API. Provides an AWS Config Configuration Recorder.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type ConfigurationRecorder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationRecorderSpec   `json:"spec"`
	Status            ConfigurationRecorderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationRecorderList contains a list of ConfigurationRecorders
type ConfigurationRecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationRecorder `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationRecorder_Kind             = "ConfigurationRecorder"
	ConfigurationRecorder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigurationRecorder_Kind}.String()
	ConfigurationRecorder_KindAPIVersion   = ConfigurationRecorder_Kind + "." + CRDGroupVersion.String()
	ConfigurationRecorder_GroupVersionKind = CRDGroupVersion.WithKind(ConfigurationRecorder_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigurationRecorder{}, &ConfigurationRecorderList{})
}
