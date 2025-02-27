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

type SharedDirectoryInitParameters struct {

	// Identifier of the Managed Microsoft AD directory that you want to share with other accounts.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta1.Directory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDRef *v1.Reference `json:"directoryIdRef,omitempty" tf:"-"`

	// Selector for a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDSelector *v1.Selector `json:"directoryIdSelector,omitempty" tf:"-"`

	// Method used when sharing a directory. Valid values are ORGANIZATIONS and HANDSHAKE. Default is HANDSHAKE.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Message sent by the directory owner to the directory consumer to help the directory consumer administrator determine whether to approve or reject the share invitation.
	NotesSecretRef *v1.SecretKeySelector `json:"notesSecretRef,omitempty" tf:"-"`

	// Identifier for the directory consumer account with whom the directory is to be shared. See below.
	Target []TargetInitParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type SharedDirectoryObservation struct {

	// Identifier of the Managed Microsoft AD directory that you want to share with other accounts.
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Identifier of the shared directory.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Method used when sharing a directory. Valid values are ORGANIZATIONS and HANDSHAKE. Default is HANDSHAKE.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Identifier of the directory that is stored in the directory consumer account that corresponds to the shared directory in the owner account.
	SharedDirectoryID *string `json:"sharedDirectoryId,omitempty" tf:"shared_directory_id,omitempty"`

	// Identifier for the directory consumer account with whom the directory is to be shared. See below.
	Target []TargetObservation `json:"target,omitempty" tf:"target,omitempty"`
}

type SharedDirectoryParameters struct {

	// Identifier of the Managed Microsoft AD directory that you want to share with other accounts.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta1.Directory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDRef *v1.Reference `json:"directoryIdRef,omitempty" tf:"-"`

	// Selector for a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDSelector *v1.Selector `json:"directoryIdSelector,omitempty" tf:"-"`

	// Method used when sharing a directory. Valid values are ORGANIZATIONS and HANDSHAKE. Default is HANDSHAKE.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Message sent by the directory owner to the directory consumer to help the directory consumer administrator determine whether to approve or reject the share invitation.
	// +kubebuilder:validation:Optional
	NotesSecretRef *v1.SecretKeySelector `json:"notesSecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Identifier for the directory consumer account with whom the directory is to be shared. See below.
	// +kubebuilder:validation:Optional
	Target []TargetParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type TargetInitParameters struct {

	// Identifier of the directory consumer account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of identifier to be used in the id field. Valid value is ACCOUNT. Default is ACCOUNT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TargetObservation struct {

	// Identifier of the directory consumer account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of identifier to be used in the id field. Valid value is ACCOUNT. Default is ACCOUNT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TargetParameters struct {

	// Identifier of the directory consumer account.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// Type of identifier to be used in the id field. Valid value is ACCOUNT. Default is ACCOUNT.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SharedDirectorySpec defines the desired state of SharedDirectory
type SharedDirectorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedDirectoryParameters `json:"forProvider"`
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
	InitProvider SharedDirectoryInitParameters `json:"initProvider,omitempty"`
}

// SharedDirectoryStatus defines the observed state of SharedDirectory.
type SharedDirectoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedDirectoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SharedDirectory is the Schema for the SharedDirectorys API. Manages a directory in your account (directory owner) shared with another account (directory consumer).
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type SharedDirectory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.target) || (has(self.initProvider) && has(self.initProvider.target))",message="spec.forProvider.target is a required parameter"
	Spec   SharedDirectorySpec   `json:"spec"`
	Status SharedDirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedDirectoryList contains a list of SharedDirectorys
type SharedDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedDirectory `json:"items"`
}

// Repository type metadata.
var (
	SharedDirectory_Kind             = "SharedDirectory"
	SharedDirectory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedDirectory_Kind}.String()
	SharedDirectory_KindAPIVersion   = SharedDirectory_Kind + "." + CRDGroupVersion.String()
	SharedDirectory_GroupVersionKind = CRDGroupVersion.WithKind(SharedDirectory_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedDirectory{}, &SharedDirectoryList{})
}
