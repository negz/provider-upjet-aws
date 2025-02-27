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

type ArchiveInitParameters struct {

	// The description of the new event archive.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the event_source_arn.
	EventPattern *string `json:"eventPattern,omitempty" tf:"event_pattern,omitempty"`

	// Event bus source ARN from where these events should be archived.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta1.Bus
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	EventSourceArn *string `json:"eventSourceArn,omitempty" tf:"event_source_arn,omitempty"`

	// Reference to a Bus in cloudwatchevents to populate eventSourceArn.
	// +kubebuilder:validation:Optional
	EventSourceArnRef *v1.Reference `json:"eventSourceArnRef,omitempty" tf:"-"`

	// Selector for a Bus in cloudwatchevents to populate eventSourceArn.
	// +kubebuilder:validation:Optional
	EventSourceArnSelector *v1.Selector `json:"eventSourceArnSelector,omitempty" tf:"-"`

	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

type ArchiveObservation struct {

	// The Amazon Resource Name (ARN) of the event archive.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the new event archive.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the event_source_arn.
	EventPattern *string `json:"eventPattern,omitempty" tf:"event_pattern,omitempty"`

	// Event bus source ARN from where these events should be archived.
	EventSourceArn *string `json:"eventSourceArn,omitempty" tf:"event_source_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

type ArchiveParameters struct {

	// The description of the new event archive.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the event_source_arn.
	// +kubebuilder:validation:Optional
	EventPattern *string `json:"eventPattern,omitempty" tf:"event_pattern,omitempty"`

	// Event bus source ARN from where these events should be archived.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta1.Bus
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	EventSourceArn *string `json:"eventSourceArn,omitempty" tf:"event_source_arn,omitempty"`

	// Reference to a Bus in cloudwatchevents to populate eventSourceArn.
	// +kubebuilder:validation:Optional
	EventSourceArnRef *v1.Reference `json:"eventSourceArnRef,omitempty" tf:"-"`

	// Selector for a Bus in cloudwatchevents to populate eventSourceArn.
	// +kubebuilder:validation:Optional
	EventSourceArnSelector *v1.Selector `json:"eventSourceArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	// +kubebuilder:validation:Optional
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

// ArchiveSpec defines the desired state of Archive
type ArchiveSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ArchiveParameters `json:"forProvider"`
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
	InitProvider ArchiveInitParameters `json:"initProvider,omitempty"`
}

// ArchiveStatus defines the observed state of Archive.
type ArchiveStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ArchiveObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Archive is the Schema for the Archives API. Provides an EventBridge event archive resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Archive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ArchiveSpec   `json:"spec"`
	Status            ArchiveStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ArchiveList contains a list of Archives
type ArchiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Archive `json:"items"`
}

// Repository type metadata.
var (
	Archive_Kind             = "Archive"
	Archive_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Archive_Kind}.String()
	Archive_KindAPIVersion   = Archive_Kind + "." + CRDGroupVersion.String()
	Archive_GroupVersionKind = CRDGroupVersion.WithKind(Archive_Kind)
)

func init() {
	SchemeBuilder.Register(&Archive{}, &ArchiveList{})
}
