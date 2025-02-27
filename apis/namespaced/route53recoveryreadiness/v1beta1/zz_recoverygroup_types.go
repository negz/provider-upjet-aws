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

type RecoveryGroupInitParameters struct {

	// List of cell arns to add as nested fault domains within this recovery group
	Cells []*string `json:"cells,omitempty" tf:"cells,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RecoveryGroupObservation struct {

	// ARN of the recovery group
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// List of cell arns to add as nested fault domains within this recovery group
	Cells []*string `json:"cells,omitempty" tf:"cells,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RecoveryGroupParameters struct {

	// List of cell arns to add as nested fault domains within this recovery group
	// +kubebuilder:validation:Optional
	Cells []*string `json:"cells,omitempty" tf:"cells,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RecoveryGroupSpec defines the desired state of RecoveryGroup
type RecoveryGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecoveryGroupParameters `json:"forProvider"`
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
	InitProvider RecoveryGroupInitParameters `json:"initProvider,omitempty"`
}

// RecoveryGroupStatus defines the observed state of RecoveryGroup.
type RecoveryGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecoveryGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RecoveryGroup is the Schema for the RecoveryGroups API. Provides an AWS Route 53 Recovery Readiness Recovery Group
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type RecoveryGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryGroupSpec   `json:"spec"`
	Status            RecoveryGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecoveryGroupList contains a list of RecoveryGroups
type RecoveryGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecoveryGroup `json:"items"`
}

// Repository type metadata.
var (
	RecoveryGroup_Kind             = "RecoveryGroup"
	RecoveryGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RecoveryGroup_Kind}.String()
	RecoveryGroup_KindAPIVersion   = RecoveryGroup_Kind + "." + CRDGroupVersion.String()
	RecoveryGroup_GroupVersionKind = CRDGroupVersion.WithKind(RecoveryGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&RecoveryGroup{}, &RecoveryGroupList{})
}
