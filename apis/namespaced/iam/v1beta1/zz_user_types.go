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

type UserInitParameters struct {

	// when destroying this user, destroy even if it
	// has non-provider-managed iam access keys, login profile or mfa devices. without force_destroy
	// a user with non-provider-managed access keys and login profile will fail to be destroyed.
	// delete user even if it has non-provider-managed iam access keys, login profile or mfa devices
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Path in which to create the user.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type UserObservation struct {

	// The ARN assigned by AWS for this user.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// when destroying this user, destroy even if it
	// has non-provider-managed iam access keys, login profile or mfa devices. without force_destroy
	// a user with non-provider-managed access keys and login profile will fail to be destroyed.
	// delete user even if it has non-provider-managed iam access keys, login profile or mfa devices
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The user's name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Path in which to create the user.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The unique ID assigned by AWS.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type UserParameters struct {

	// when destroying this user, destroy even if it
	// has non-provider-managed iam access keys, login profile or mfa devices. without force_destroy
	// a user with non-provider-managed access keys and login profile will fail to be destroyed.
	// delete user even if it has non-provider-managed iam access keys, login profile or mfa devices
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Path in which to create the user.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The ARN of the policy that is used to set the permissions boundary for the user.
	// +kubebuilder:validation:Optional
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// User is the Schema for the Users API. Provides an IAM user.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
