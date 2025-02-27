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

	// A map that contains user attributes and attribute values to be set for the user.
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the client_metadata value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see Customizing User Pool Workflows with Lambda Triggers.
	// +mapType=granular
	ClientMetadata map[string]*string `json:"clientMetadata,omitempty" tf:"client_metadata,omitempty"`

	// A list of mediums to the welcome message will be sent through. Allowed values are EMAIL and SMS. If it's provided, make sure you have also specified email attribute for the EMAIL medium and phone_number for the SMS. More than one value can be specified. Amazon Cognito does not store the desired_delivery_mediums value. Defaults to ["SMS"].
	// +listType=set
	DesiredDeliveryMediums []*string `json:"desiredDeliveryMediums,omitempty" tf:"desired_delivery_mediums,omitempty"`

	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the enabled value. The behavior can be changed with message_action argument. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If this parameter is set to True and the phone_number or email address specified in the attributes parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the force_alias_creation value. Defaults to false.
	ForceAliasCreation *bool `json:"forceAliasCreation,omitempty" tf:"force_alias_creation,omitempty"`

	// Set to RESEND to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to SUPPRESS to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the message_action value.
	MessageAction *string `json:"messageAction,omitempty" tf:"message_action,omitempty"`

	// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only temporary_password value. You can suppress sending the welcome message with the message_action argument. Amazon Cognito does not store the password value. Conflicts with temporary_password.
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The user's temporary password. Conflicts with password.
	TemporaryPasswordSecretRef *v1.SecretKeySelector `json:"temporaryPasswordSecretRef,omitempty" tf:"-"`

	// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the validation_data value. For more information, see Customizing User Pool Workflows with Lambda Triggers.
	// +mapType=granular
	ValidationData map[string]*string `json:"validationData,omitempty" tf:"validation_data,omitempty"`
}

type UserObservation struct {

	// A map that contains user attributes and attribute values to be set for the user.
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the client_metadata value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see Customizing User Pool Workflows with Lambda Triggers.
	// +mapType=granular
	ClientMetadata map[string]*string `json:"clientMetadata,omitempty" tf:"client_metadata,omitempty"`

	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// A list of mediums to the welcome message will be sent through. Allowed values are EMAIL and SMS. If it's provided, make sure you have also specified email attribute for the EMAIL medium and phone_number for the SMS. More than one value can be specified. Amazon Cognito does not store the desired_delivery_mediums value. Defaults to ["SMS"].
	// +listType=set
	DesiredDeliveryMediums []*string `json:"desiredDeliveryMediums,omitempty" tf:"desired_delivery_mediums,omitempty"`

	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the enabled value. The behavior can be changed with message_action argument. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If this parameter is set to True and the phone_number or email address specified in the attributes parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the force_alias_creation value. Defaults to false.
	ForceAliasCreation *bool `json:"forceAliasCreation,omitempty" tf:"force_alias_creation,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastModifiedDate *string `json:"lastModifiedDate,omitempty" tf:"last_modified_date,omitempty"`

	// Set to RESEND to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to SUPPRESS to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the message_action value.
	MessageAction *string `json:"messageAction,omitempty" tf:"message_action,omitempty"`

	// +listType=set
	MfaSettingList []*string `json:"mfaSettingList,omitempty" tf:"mfa_setting_list,omitempty"`

	PreferredMfaSetting *string `json:"preferredMfaSetting,omitempty" tf:"preferred_mfa_setting,omitempty"`

	// current user status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// unique user id that is never reassignable to another user.
	Sub *string `json:"sub,omitempty" tf:"sub,omitempty"`

	// The user pool ID for the user pool where the user will be created.
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the validation_data value. For more information, see Customizing User Pool Workflows with Lambda Triggers.
	// +mapType=granular
	ValidationData map[string]*string `json:"validationData,omitempty" tf:"validation_data,omitempty"`
}

type UserParameters struct {

	// A map that contains user attributes and attribute values to be set for the user.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// A map of custom key-value pairs that you can provide as input for any custom workflows that user creation triggers. Amazon Cognito does not store the client_metadata value. This data is available only to Lambda triggers that are assigned to a user pool to support custom workflows. If your user pool configuration does not include triggers, the ClientMetadata parameter serves no purpose. For more information, see Customizing User Pool Workflows with Lambda Triggers.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ClientMetadata map[string]*string `json:"clientMetadata,omitempty" tf:"client_metadata,omitempty"`

	// A list of mediums to the welcome message will be sent through. Allowed values are EMAIL and SMS. If it's provided, make sure you have also specified email attribute for the EMAIL medium and phone_number for the SMS. More than one value can be specified. Amazon Cognito does not store the desired_delivery_mediums value. Defaults to ["SMS"].
	// +kubebuilder:validation:Optional
	// +listType=set
	DesiredDeliveryMediums []*string `json:"desiredDeliveryMediums,omitempty" tf:"desired_delivery_mediums,omitempty"`

	// Specifies whether the user should be enabled after creation. The welcome message will be sent regardless of the enabled value. The behavior can be changed with message_action argument. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If this parameter is set to True and the phone_number or email address specified in the attributes parameter already exists as an alias with a different user, Amazon Cognito will migrate the alias from the previous user to the newly created user. The previous user will no longer be able to log in using that alias. Amazon Cognito does not store the force_alias_creation value. Defaults to false.
	// +kubebuilder:validation:Optional
	ForceAliasCreation *bool `json:"forceAliasCreation,omitempty" tf:"force_alias_creation,omitempty"`

	// Set to RESEND to resend the invitation message to a user that already exists and reset the expiration limit on the user's account. Set to SUPPRESS to suppress sending the message. Only one value can be specified. Amazon Cognito does not store the message_action value.
	// +kubebuilder:validation:Optional
	MessageAction *string `json:"messageAction,omitempty" tf:"message_action,omitempty"`

	// The user's permanent password. This password must conform to the password policy specified by user pool the user belongs to. The welcome message always contains only temporary_password value. You can suppress sending the welcome message with the message_action argument. Amazon Cognito does not store the password value. Conflicts with temporary_password.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The user's temporary password. Conflicts with password.
	// +kubebuilder:validation:Optional
	TemporaryPasswordSecretRef *v1.SecretKeySelector `json:"temporaryPasswordSecretRef,omitempty" tf:"-"`

	// The user pool ID for the user pool where the user will be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/cognitoidp/v1beta2.UserPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`

	// The user's validation data. This is an array of name-value pairs that contain user attributes and attribute values that you can use for custom validation, such as restricting the types of user accounts that can be registered. Amazon Cognito does not store the validation_data value. For more information, see Customizing User Pool Workflows with Lambda Triggers.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValidationData map[string]*string `json:"validationData,omitempty" tf:"validation_data,omitempty"`
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

// User is the Schema for the Users API. Provides a Cognito User resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
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
