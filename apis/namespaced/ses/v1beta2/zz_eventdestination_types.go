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

type CloudwatchDestinationInitParameters struct {

	// The default value for the event
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// The name for the dimension
	DimensionName *string `json:"dimensionName,omitempty" tf:"dimension_name,omitempty"`

	// The source for the value. May be any of "messageTag", "emailHeader" or "linkTag".
	ValueSource *string `json:"valueSource,omitempty" tf:"value_source,omitempty"`
}

type CloudwatchDestinationObservation struct {

	// The default value for the event
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// The name for the dimension
	DimensionName *string `json:"dimensionName,omitempty" tf:"dimension_name,omitempty"`

	// The source for the value. May be any of "messageTag", "emailHeader" or "linkTag".
	ValueSource *string `json:"valueSource,omitempty" tf:"value_source,omitempty"`
}

type CloudwatchDestinationParameters struct {

	// The default value for the event
	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue" tf:"default_value,omitempty"`

	// The name for the dimension
	// +kubebuilder:validation:Optional
	DimensionName *string `json:"dimensionName" tf:"dimension_name,omitempty"`

	// The source for the value. May be any of "messageTag", "emailHeader" or "linkTag".
	// +kubebuilder:validation:Optional
	ValueSource *string `json:"valueSource" tf:"value_source,omitempty"`
}

type EventDestinationInitParameters struct {

	// CloudWatch destination for the events
	CloudwatchDestination []CloudwatchDestinationInitParameters `json:"cloudwatchDestination,omitempty" tf:"cloudwatch_destination,omitempty"`

	// The name of the configuration set
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ses/v1beta2.ConfigurationSet
	ConfigurationSetName *string `json:"configurationSetName,omitempty" tf:"configuration_set_name,omitempty"`

	// Reference to a ConfigurationSet in ses to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameRef *v1.Reference `json:"configurationSetNameRef,omitempty" tf:"-"`

	// Selector for a ConfigurationSet in ses to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameSelector *v1.Selector `json:"configurationSetNameSelector,omitempty" tf:"-"`

	// If true, the event destination will be enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Send the events to a kinesis firehose destination
	KinesisDestination *KinesisDestinationInitParameters `json:"kinesisDestination,omitempty" tf:"kinesis_destination,omitempty"`

	// A list of matching types. May be any of "send", "reject", "bounce", "complaint", "delivery", "open", "click", or "renderingFailure".
	// +listType=set
	MatchingTypes []*string `json:"matchingTypes,omitempty" tf:"matching_types,omitempty"`

	// Send the events to an SNS Topic destination
	SnsDestination *SnsDestinationInitParameters `json:"snsDestination,omitempty" tf:"sns_destination,omitempty"`
}

type EventDestinationObservation struct {

	// The SES event destination ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// CloudWatch destination for the events
	CloudwatchDestination []CloudwatchDestinationObservation `json:"cloudwatchDestination,omitempty" tf:"cloudwatch_destination,omitempty"`

	// The name of the configuration set
	ConfigurationSetName *string `json:"configurationSetName,omitempty" tf:"configuration_set_name,omitempty"`

	// If true, the event destination will be enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The SES event destination name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Send the events to a kinesis firehose destination
	KinesisDestination *KinesisDestinationObservation `json:"kinesisDestination,omitempty" tf:"kinesis_destination,omitempty"`

	// A list of matching types. May be any of "send", "reject", "bounce", "complaint", "delivery", "open", "click", or "renderingFailure".
	// +listType=set
	MatchingTypes []*string `json:"matchingTypes,omitempty" tf:"matching_types,omitempty"`

	// Send the events to an SNS Topic destination
	SnsDestination *SnsDestinationObservation `json:"snsDestination,omitempty" tf:"sns_destination,omitempty"`
}

type EventDestinationParameters struct {

	// CloudWatch destination for the events
	// +kubebuilder:validation:Optional
	CloudwatchDestination []CloudwatchDestinationParameters `json:"cloudwatchDestination,omitempty" tf:"cloudwatch_destination,omitempty"`

	// The name of the configuration set
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ses/v1beta2.ConfigurationSet
	// +kubebuilder:validation:Optional
	ConfigurationSetName *string `json:"configurationSetName,omitempty" tf:"configuration_set_name,omitempty"`

	// Reference to a ConfigurationSet in ses to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameRef *v1.Reference `json:"configurationSetNameRef,omitempty" tf:"-"`

	// Selector for a ConfigurationSet in ses to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameSelector *v1.Selector `json:"configurationSetNameSelector,omitempty" tf:"-"`

	// If true, the event destination will be enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Send the events to a kinesis firehose destination
	// +kubebuilder:validation:Optional
	KinesisDestination *KinesisDestinationParameters `json:"kinesisDestination,omitempty" tf:"kinesis_destination,omitempty"`

	// A list of matching types. May be any of "send", "reject", "bounce", "complaint", "delivery", "open", "click", or "renderingFailure".
	// +kubebuilder:validation:Optional
	// +listType=set
	MatchingTypes []*string `json:"matchingTypes,omitempty" tf:"matching_types,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Send the events to an SNS Topic destination
	// +kubebuilder:validation:Optional
	SnsDestination *SnsDestinationParameters `json:"snsDestination,omitempty" tf:"sns_destination,omitempty"`
}

type KinesisDestinationInitParameters struct {

	// The ARN of the role that has permissions to access the Kinesis Stream
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The ARN of the Kinesis Stream
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta2.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",false)
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnRef *v1.Reference `json:"streamArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnSelector *v1.Selector `json:"streamArnSelector,omitempty" tf:"-"`
}

type KinesisDestinationObservation struct {

	// The ARN of the role that has permissions to access the Kinesis Stream
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The ARN of the Kinesis Stream
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`
}

type KinesisDestinationParameters struct {

	// The ARN of the role that has permissions to access the Kinesis Stream
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The ARN of the Kinesis Stream
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta2.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnRef *v1.Reference `json:"streamArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnSelector *v1.Selector `json:"streamArnSelector,omitempty" tf:"-"`
}

type SnsDestinationInitParameters struct {

	// The ARN of the SNS topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

type SnsDestinationObservation struct {

	// The ARN of the SNS topic
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SnsDestinationParameters struct {

	// The ARN of the SNS topic
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

// EventDestinationSpec defines the desired state of EventDestination
type EventDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventDestinationParameters `json:"forProvider"`
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
	InitProvider EventDestinationInitParameters `json:"initProvider,omitempty"`
}

// EventDestinationStatus defines the observed state of EventDestination.
type EventDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// EventDestination is the Schema for the EventDestinations API. Provides an SES event destination
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type EventDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.matchingTypes) || (has(self.initProvider) && has(self.initProvider.matchingTypes))",message="spec.forProvider.matchingTypes is a required parameter"
	Spec   EventDestinationSpec   `json:"spec"`
	Status EventDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventDestinationList contains a list of EventDestinations
type EventDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventDestination `json:"items"`
}

// Repository type metadata.
var (
	EventDestination_Kind             = "EventDestination"
	EventDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventDestination_Kind}.String()
	EventDestination_KindAPIVersion   = EventDestination_Kind + "." + CRDGroupVersion.String()
	EventDestination_GroupVersionKind = CRDGroupVersion.WithKind(EventDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&EventDestination{}, &EventDestinationList{})
}
