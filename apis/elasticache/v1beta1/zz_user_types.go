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

type AuthenticationModeInitParameters struct {

	// Specifies the authentication type. Possible options are: password, no-password-required or iam.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuthenticationModeObservation struct {
	PasswordCount *float64 `json:"passwordCount,omitempty" tf:"password_count,omitempty"`

	// Specifies the authentication type. Possible options are: password, no-password-required or iam.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuthenticationModeParameters struct {

	// Specifies the passwords to use for authentication if type is set to password.
	// +kubebuilder:validation:Optional
	PasswordsSecretRef *[]v1.SecretKeySelector `json:"passwordsSecretRef,omitempty" tf:"-"`

	// Specifies the authentication type. Possible options are: password, no-password-required or iam.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type UserInitParameters struct {

	// Access permissions string used for this user. See Specifying Permissions Using an Access String for more details.
	AccessString *string `json:"accessString,omitempty" tf:"access_string,omitempty"`

	// Denotes the user's authentication properties. Detailed below.
	AuthenticationMode []AuthenticationModeInitParameters `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The current supported value is REDIS.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Indicates a password is not required for this user.
	NoPasswordRequired *bool `json:"noPasswordRequired,omitempty" tf:"no_password_required,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The username of the user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UserObservation struct {

	// Access permissions string used for this user. See Specifying Permissions Using an Access String for more details.
	AccessString *string `json:"accessString,omitempty" tf:"access_string,omitempty"`

	// The ARN of the created ElastiCache User.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Denotes the user's authentication properties. Detailed below.
	AuthenticationMode []AuthenticationModeObservation `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The current supported value is REDIS.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates a password is not required for this user.
	NoPasswordRequired *bool `json:"noPasswordRequired,omitempty" tf:"no_password_required,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The username of the user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UserParameters struct {

	// Access permissions string used for this user. See Specifying Permissions Using an Access String for more details.
	// +kubebuilder:validation:Optional
	AccessString *string `json:"accessString,omitempty" tf:"access_string,omitempty"`

	// Denotes the user's authentication properties. Detailed below.
	// +kubebuilder:validation:Optional
	AuthenticationMode []AuthenticationModeParameters `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The current supported value is REDIS.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Indicates a password is not required for this user.
	// +kubebuilder:validation:Optional
	NoPasswordRequired *bool `json:"noPasswordRequired,omitempty" tf:"no_password_required,omitempty"`

	// Passwords used for this user. You can create up to two passwords for each user.
	// +kubebuilder:validation:Optional
	PasswordsSecretRef *[]v1.SecretKeySelector `json:"passwordsSecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The username of the user.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
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

// User is the Schema for the Users API. Provides an ElastiCache user.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessString) || (has(self.initProvider) && has(self.initProvider.accessString))",message="spec.forProvider.accessString is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || (has(self.initProvider) && has(self.initProvider.engine))",message="spec.forProvider.engine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userName) || (has(self.initProvider) && has(self.initProvider.userName))",message="spec.forProvider.userName is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
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
