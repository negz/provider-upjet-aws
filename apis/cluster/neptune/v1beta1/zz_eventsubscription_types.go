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

type EventSubscriptionInitParameters struct {

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a source_type that you want to subscribe to. Run aws neptune describe-event-categories to find all the event categories.
	// +listType=set
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// The ARN of the SNS topic to send events to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// Reference to a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnRef *v1.Reference `json:"snsTopicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnSelector *v1.Selector `json:"snsTopicArnSelector,omitempty" tf:"-"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/neptune/v1beta1.ClusterInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +listType=set
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// References to ClusterInstance in neptune to populate sourceIds.
	// +kubebuilder:validation:Optional
	SourceIdsRefs []v1.Reference `json:"sourceIdsRefs,omitempty" tf:"-"`

	// Selector for a list of ClusterInstance in neptune to populate sourceIds.
	// +kubebuilder:validation:Optional
	SourceIdsSelector *v1.Selector `json:"sourceIdsSelector,omitempty" tf:"-"`

	// The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot. If not set, all sources will be subscribed to.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventSubscriptionObservation struct {

	// The Amazon Resource Name of the Neptune event notification subscription.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The AWS customer account associated with the Neptune event notification subscription.
	CustomerAwsID *string `json:"customerAwsId,omitempty" tf:"customer_aws_id,omitempty"`

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a source_type that you want to subscribe to. Run aws neptune describe-event-categories to find all the event categories.
	// +listType=set
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// The name of the Neptune event notification subscription.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the SNS topic to send events to.
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	// +listType=set
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot. If not set, all sources will be subscribed to.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventSubscriptionParameters struct {

	// A boolean flag to enable/disable the subscription. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a source_type that you want to subscribe to. Run aws neptune describe-event-categories to find all the event categories.
	// +kubebuilder:validation:Optional
	// +listType=set
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the SNS topic to send events to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// Reference to a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnRef *v1.Reference `json:"snsTopicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnSelector *v1.Selector `json:"snsTopicArnSelector,omitempty" tf:"-"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/neptune/v1beta1.ClusterInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// References to ClusterInstance in neptune to populate sourceIds.
	// +kubebuilder:validation:Optional
	SourceIdsRefs []v1.Reference `json:"sourceIdsRefs,omitempty" tf:"-"`

	// Selector for a list of ClusterInstance in neptune to populate sourceIds.
	// +kubebuilder:validation:Optional
	SourceIdsSelector *v1.Selector `json:"sourceIdsSelector,omitempty" tf:"-"`

	// The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot. If not set, all sources will be subscribed to.
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EventSubscriptionSpec defines the desired state of EventSubscription
type EventSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventSubscriptionParameters `json:"forProvider"`
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
	InitProvider EventSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// EventSubscriptionStatus defines the observed state of EventSubscription.
type EventSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EventSubscription is the Schema for the EventSubscriptions API. Provides a Neptune event subscription resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSubscriptionSpec   `json:"spec"`
	Status            EventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSubscriptionList contains a list of EventSubscriptions
type EventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSubscription `json:"items"`
}

// Repository type metadata.
var (
	EventSubscription_Kind             = "EventSubscription"
	EventSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventSubscription_Kind}.String()
	EventSubscription_KindAPIVersion   = EventSubscription_Kind + "." + CRDGroupVersion.String()
	EventSubscription_GroupVersionKind = CRDGroupVersion.WithKind(EventSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&EventSubscription{}, &EventSubscriptionList{})
}
