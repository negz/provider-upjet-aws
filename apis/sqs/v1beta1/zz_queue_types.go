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

type QueueInitParameters struct {

	// Enables content-based deduplication for FIFO queues. For more information, see the related documentation
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue (default).
	DeduplicationScope *string `json:"deduplicationScope,omitempty" tf:"deduplication_scope,omitempty"`

	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// Boolean designating a FIFO queue. If not set, it defaults to false making it standard.
	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue,omitempty"`

	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue (default) and perMessageGroupId.
	FifoThroughputLimit *string `json:"fifoThroughputLimit,omitempty" tf:"fifo_throughput_limit,omitempty"`

	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KMSDataKeyReusePeriodSeconds *float64 `json:"kmsDataKeyReusePeriodSeconds,omitempty" tf:"kms_data_key_reuse_period_seconds,omitempty"`

	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see Key Terms.
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize *float64 `json:"maxMessageSize,omitempty" tf:"max_message_size,omitempty"`

	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds *float64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds,omitempty"`

	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the .fifo suffix. Conflicts with name_prefix
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The JSON policy for the SQS queue.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds *float64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds,omitempty"`

	// The JSON policy to set up the Dead Letter Queue redrive permission, see AWS docs.
	RedriveAllowPolicy *string `json:"redriveAllowPolicy,omitempty" tf:"redrive_allow_policy,omitempty"`

	// The JSON policy to set up the Dead Letter Queue, see AWS docs. Note: when specifying maxReceiveCount, you must specify it as an integer (5), and not a string ("5").
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys. See Encryption at rest.
	SqsManagedSseEnabled *bool `json:"sqsManagedSseEnabled,omitempty" tf:"sqs_managed_sse_enabled,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see AWS docs.
	VisibilityTimeoutSeconds *float64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds,omitempty"`
}

type QueueObservation struct {

	// The ARN of the SQS queue
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Enables content-based deduplication for FIFO queues. For more information, see the related documentation
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue (default).
	DeduplicationScope *string `json:"deduplicationScope,omitempty" tf:"deduplication_scope,omitempty"`

	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// Boolean designating a FIFO queue. If not set, it defaults to false making it standard.
	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue,omitempty"`

	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue (default) and perMessageGroupId.
	FifoThroughputLimit *string `json:"fifoThroughputLimit,omitempty" tf:"fifo_throughput_limit,omitempty"`

	// The URL for the created Amazon SQS queue.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KMSDataKeyReusePeriodSeconds *float64 `json:"kmsDataKeyReusePeriodSeconds,omitempty" tf:"kms_data_key_reuse_period_seconds,omitempty"`

	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see Key Terms.
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize *float64 `json:"maxMessageSize,omitempty" tf:"max_message_size,omitempty"`

	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds *float64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds,omitempty"`

	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the .fifo suffix. Conflicts with name_prefix
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The JSON policy for the SQS queue.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds *float64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds,omitempty"`

	// The JSON policy to set up the Dead Letter Queue redrive permission, see AWS docs.
	RedriveAllowPolicy *string `json:"redriveAllowPolicy,omitempty" tf:"redrive_allow_policy,omitempty"`

	// The JSON policy to set up the Dead Letter Queue, see AWS docs. Note: when specifying maxReceiveCount, you must specify it as an integer (5), and not a string ("5").
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys. See Encryption at rest.
	SqsManagedSseEnabled *bool `json:"sqsManagedSseEnabled,omitempty" tf:"sqs_managed_sse_enabled,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Same as id: The URL for the created Amazon SQS queue.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see AWS docs.
	VisibilityTimeoutSeconds *float64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds,omitempty"`
}

type QueueParameters struct {

	// Enables content-based deduplication for FIFO queues. For more information, see the related documentation
	// +kubebuilder:validation:Optional
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue (default).
	// +kubebuilder:validation:Optional
	DeduplicationScope *string `json:"deduplicationScope,omitempty" tf:"deduplication_scope,omitempty"`

	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	// +kubebuilder:validation:Optional
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// Boolean designating a FIFO queue. If not set, it defaults to false making it standard.
	// +kubebuilder:validation:Optional
	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue,omitempty"`

	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue (default) and perMessageGroupId.
	// +kubebuilder:validation:Optional
	FifoThroughputLimit *string `json:"fifoThroughputLimit,omitempty" tf:"fifo_throughput_limit,omitempty"`

	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	// +kubebuilder:validation:Optional
	KMSDataKeyReusePeriodSeconds *float64 `json:"kmsDataKeyReusePeriodSeconds,omitempty" tf:"kms_data_key_reuse_period_seconds,omitempty"`

	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see Key Terms.
	// +kubebuilder:validation:Optional
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	// +kubebuilder:validation:Optional
	MaxMessageSize *float64 `json:"maxMessageSize,omitempty" tf:"max_message_size,omitempty"`

	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	// +kubebuilder:validation:Optional
	MessageRetentionSeconds *float64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds,omitempty"`

	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the .fifo suffix. Conflicts with name_prefix
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The JSON policy for the SQS queue.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	// +kubebuilder:validation:Optional
	ReceiveWaitTimeSeconds *float64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds,omitempty"`

	// The JSON policy to set up the Dead Letter Queue redrive permission, see AWS docs.
	// +kubebuilder:validation:Optional
	RedriveAllowPolicy *string `json:"redriveAllowPolicy,omitempty" tf:"redrive_allow_policy,omitempty"`

	// The JSON policy to set up the Dead Letter Queue, see AWS docs. Note: when specifying maxReceiveCount, you must specify it as an integer (5), and not a string ("5").
	// +kubebuilder:validation:Optional
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys. See Encryption at rest.
	// +kubebuilder:validation:Optional
	SqsManagedSseEnabled *bool `json:"sqsManagedSseEnabled,omitempty" tf:"sqs_managed_sse_enabled,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see AWS docs.
	// +kubebuilder:validation:Optional
	VisibilityTimeoutSeconds *float64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
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
	InitProvider QueueInitParameters `json:"initProvider,omitempty"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Queue is the Schema for the Queues API. Provides a SQS resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec"`
	Status            QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
