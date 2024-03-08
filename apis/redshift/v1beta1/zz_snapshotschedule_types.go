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

type SnapshotScheduleInitParameters struct {

	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example cron(30 12 *) or rate(12 hours).
	// +listType=set
	Definitions []*string `json:"definitions,omitempty" tf:"definitions,omitempty"`

	// The description of the snapshot schedule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SnapshotScheduleObservation struct {

	// Amazon Resource Name (ARN) of the Redshift Snapshot Schedule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example cron(30 12 *) or rate(12 hours).
	// +listType=set
	Definitions []*string `json:"definitions,omitempty" tf:"definitions,omitempty"`

	// The description of the snapshot schedule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SnapshotScheduleParameters struct {

	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example cron(30 12 *) or rate(12 hours).
	// +kubebuilder:validation:Optional
	// +listType=set
	Definitions []*string `json:"definitions,omitempty" tf:"definitions,omitempty"`

	// The description of the snapshot schedule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SnapshotScheduleSpec defines the desired state of SnapshotSchedule
type SnapshotScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotScheduleParameters `json:"forProvider"`
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
	InitProvider SnapshotScheduleInitParameters `json:"initProvider,omitempty"`
}

// SnapshotScheduleStatus defines the observed state of SnapshotSchedule.
type SnapshotScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SnapshotSchedule is the Schema for the SnapshotSchedules API. Provides an Redshift Snapshot Schedule resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SnapshotSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.definitions) || (has(self.initProvider) && has(self.initProvider.definitions))",message="spec.forProvider.definitions is a required parameter"
	Spec   SnapshotScheduleSpec   `json:"spec"`
	Status SnapshotScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotScheduleList contains a list of SnapshotSchedules
type SnapshotScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotSchedule `json:"items"`
}

// Repository type metadata.
var (
	SnapshotSchedule_Kind             = "SnapshotSchedule"
	SnapshotSchedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotSchedule_Kind}.String()
	SnapshotSchedule_KindAPIVersion   = SnapshotSchedule_Kind + "." + CRDGroupVersion.String()
	SnapshotSchedule_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotSchedule_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotSchedule{}, &SnapshotScheduleList{})
}
