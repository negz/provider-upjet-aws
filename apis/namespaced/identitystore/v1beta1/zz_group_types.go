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

type ExternalIdsInitParameters struct {
}

type ExternalIdsObservation struct {

	// The identifier issued to this resource by an external identity provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The issuer for an external identifier.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type ExternalIdsParameters struct {
}

type GroupInitParameters struct {

	// A string containing the description of the group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type GroupObservation struct {

	// A string containing the description of the group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A string containing the name of the group. This value is commonly displayed when the group is referenced.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A list of external IDs that contains the identifiers issued to this resource by an external identity provider. See External IDs below.
	ExternalIds []ExternalIdsObservation `json:"externalIds,omitempty" tf:"external_ids,omitempty"`

	// The identifier of the newly created group in the identity store.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The identifier issued to this resource by an external identity provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The globally unique identifier for the identity store.
	IdentityStoreID *string `json:"identityStoreId,omitempty" tf:"identity_store_id,omitempty"`
}

type GroupParameters struct {

	// A string containing the description of the group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A string containing the name of the group. This value is commonly displayed when the group is referenced.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The globally unique identifier for the identity store.
	// +kubebuilder:validation:Required
	IdentityStoreID *string `json:"identityStoreId" tf:"identity_store_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Group is the Schema for the Groups API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
