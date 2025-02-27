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

type ApplicationSnapshotInitParameters struct {

	// The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kinesisanalyticsv2/v1beta2.Application
	ApplicationName *string `json:"applicationName,omitempty" tf:"application_name,omitempty"`

	// Reference to a Application in kinesisanalyticsv2 to populate applicationName.
	// +kubebuilder:validation:Optional
	ApplicationNameRef *v1.Reference `json:"applicationNameRef,omitempty" tf:"-"`

	// Selector for a Application in kinesisanalyticsv2 to populate applicationName.
	// +kubebuilder:validation:Optional
	ApplicationNameSelector *v1.Selector `json:"applicationNameSelector,omitempty" tf:"-"`
}

type ApplicationSnapshotObservation struct {

	// The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
	ApplicationName *string `json:"applicationName,omitempty" tf:"application_name,omitempty"`

	// The current application version ID when the snapshot was created.
	ApplicationVersionID *float64 `json:"applicationVersionId,omitempty" tf:"application_version_id,omitempty"`

	// The application snapshot identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp of the application snapshot.
	SnapshotCreationTimestamp *string `json:"snapshotCreationTimestamp,omitempty" tf:"snapshot_creation_timestamp,omitempty"`
}

type ApplicationSnapshotParameters struct {

	// The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kinesisanalyticsv2/v1beta2.Application
	// +kubebuilder:validation:Optional
	ApplicationName *string `json:"applicationName,omitempty" tf:"application_name,omitempty"`

	// Reference to a Application in kinesisanalyticsv2 to populate applicationName.
	// +kubebuilder:validation:Optional
	ApplicationNameRef *v1.Reference `json:"applicationNameRef,omitempty" tf:"-"`

	// Selector for a Application in kinesisanalyticsv2 to populate applicationName.
	// +kubebuilder:validation:Optional
	ApplicationNameSelector *v1.Selector `json:"applicationNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ApplicationSnapshotSpec defines the desired state of ApplicationSnapshot
type ApplicationSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationSnapshotParameters `json:"forProvider"`
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
	InitProvider ApplicationSnapshotInitParameters `json:"initProvider,omitempty"`
}

// ApplicationSnapshotStatus defines the observed state of ApplicationSnapshot.
type ApplicationSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ApplicationSnapshot is the Schema for the ApplicationSnapshots API. Manages a Kinesis Analytics v2 Application Snapshot.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type ApplicationSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSnapshotSpec   `json:"spec"`
	Status            ApplicationSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationSnapshotList contains a list of ApplicationSnapshots
type ApplicationSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationSnapshot `json:"items"`
}

// Repository type metadata.
var (
	ApplicationSnapshot_Kind             = "ApplicationSnapshot"
	ApplicationSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationSnapshot_Kind}.String()
	ApplicationSnapshot_KindAPIVersion   = ApplicationSnapshot_Kind + "." + CRDGroupVersion.String()
	ApplicationSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationSnapshot{}, &ApplicationSnapshotList{})
}
