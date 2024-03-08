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

type AccountTakeoverRiskConfigurationInitParameters struct {

	// Account takeover risk configuration actions. See details below.
	Actions []ActionsInitParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// The notify configuration used to construct email notifications. See details below.
	NotifyConfiguration []NotifyConfigurationInitParameters `json:"notifyConfiguration,omitempty" tf:"notify_configuration,omitempty"`
}

type AccountTakeoverRiskConfigurationObservation struct {

	// Account takeover risk configuration actions. See details below.
	Actions []ActionsObservation `json:"actions,omitempty" tf:"actions,omitempty"`

	// The notify configuration used to construct email notifications. See details below.
	NotifyConfiguration []NotifyConfigurationObservation `json:"notifyConfiguration,omitempty" tf:"notify_configuration,omitempty"`
}

type AccountTakeoverRiskConfigurationParameters struct {

	// Account takeover risk configuration actions. See details below.
	// +kubebuilder:validation:Optional
	Actions []ActionsParameters `json:"actions" tf:"actions,omitempty"`

	// The notify configuration used to construct email notifications. See details below.
	// +kubebuilder:validation:Optional
	NotifyConfiguration []NotifyConfigurationParameters `json:"notifyConfiguration" tf:"notify_configuration,omitempty"`
}

type ActionsInitParameters struct {

	// Action to take for a high risk. See action block below.
	HighAction []HighActionInitParameters `json:"highAction,omitempty" tf:"high_action,omitempty"`

	// Action to take for a low risk. See action block below.
	LowAction []LowActionInitParameters `json:"lowAction,omitempty" tf:"low_action,omitempty"`

	// Action to take for a medium risk. See action block below.
	MediumAction []MediumActionInitParameters `json:"mediumAction,omitempty" tf:"medium_action,omitempty"`
}

type ActionsObservation struct {

	// Action to take for a high risk. See action block below.
	HighAction []HighActionObservation `json:"highAction,omitempty" tf:"high_action,omitempty"`

	// Action to take for a low risk. See action block below.
	LowAction []LowActionObservation `json:"lowAction,omitempty" tf:"low_action,omitempty"`

	// Action to take for a medium risk. See action block below.
	MediumAction []MediumActionObservation `json:"mediumAction,omitempty" tf:"medium_action,omitempty"`
}

type ActionsParameters struct {

	// Action to take for a high risk. See action block below.
	// +kubebuilder:validation:Optional
	HighAction []HighActionParameters `json:"highAction,omitempty" tf:"high_action,omitempty"`

	// Action to take for a low risk. See action block below.
	// +kubebuilder:validation:Optional
	LowAction []LowActionParameters `json:"lowAction,omitempty" tf:"low_action,omitempty"`

	// Action to take for a medium risk. See action block below.
	// +kubebuilder:validation:Optional
	MediumAction []MediumActionParameters `json:"mediumAction,omitempty" tf:"medium_action,omitempty"`
}

type BlockEmailInitParameters struct {

	// The email HTML body.
	HTMLBody *string `json:"htmlBody,omitempty" tf:"html_body,omitempty"`

	// The email subject.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email text body.
	TextBody *string `json:"textBody,omitempty" tf:"text_body,omitempty"`
}

type BlockEmailObservation struct {

	// The email HTML body.
	HTMLBody *string `json:"htmlBody,omitempty" tf:"html_body,omitempty"`

	// The email subject.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email text body.
	TextBody *string `json:"textBody,omitempty" tf:"text_body,omitempty"`
}

type BlockEmailParameters struct {

	// The email HTML body.
	// +kubebuilder:validation:Optional
	HTMLBody *string `json:"htmlBody" tf:"html_body,omitempty"`

	// The email subject.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject" tf:"subject,omitempty"`

	// The email text body.
	// +kubebuilder:validation:Optional
	TextBody *string `json:"textBody" tf:"text_body,omitempty"`
}

type CompromisedCredentialsRiskConfigurationActionsInitParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`
}

type CompromisedCredentialsRiskConfigurationActionsObservation struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`
}

type CompromisedCredentialsRiskConfigurationActionsParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	// +kubebuilder:validation:Optional
	EventAction *string `json:"eventAction" tf:"event_action,omitempty"`
}

type CompromisedCredentialsRiskConfigurationInitParameters struct {

	// The compromised credentials risk configuration actions. See details below.
	Actions []CompromisedCredentialsRiskConfigurationActionsInitParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// Perform the action for these events. The default is to perform all events if no event filter is specified. Valid values are SIGN_IN, PASSWORD_CHANGE, and SIGN_UP.
	// +listType=set
	EventFilter []*string `json:"eventFilter,omitempty" tf:"event_filter,omitempty"`
}

type CompromisedCredentialsRiskConfigurationObservation struct {

	// The compromised credentials risk configuration actions. See details below.
	Actions []CompromisedCredentialsRiskConfigurationActionsObservation `json:"actions,omitempty" tf:"actions,omitempty"`

	// Perform the action for these events. The default is to perform all events if no event filter is specified. Valid values are SIGN_IN, PASSWORD_CHANGE, and SIGN_UP.
	// +listType=set
	EventFilter []*string `json:"eventFilter,omitempty" tf:"event_filter,omitempty"`
}

type CompromisedCredentialsRiskConfigurationParameters struct {

	// The compromised credentials risk configuration actions. See details below.
	// +kubebuilder:validation:Optional
	Actions []CompromisedCredentialsRiskConfigurationActionsParameters `json:"actions" tf:"actions,omitempty"`

	// Perform the action for these events. The default is to perform all events if no event filter is specified. Valid values are SIGN_IN, PASSWORD_CHANGE, and SIGN_UP.
	// +kubebuilder:validation:Optional
	// +listType=set
	EventFilter []*string `json:"eventFilter,omitempty" tf:"event_filter,omitempty"`
}

type HighActionInitParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Whether to send a notification.
	Notify *bool `json:"notify,omitempty" tf:"notify,omitempty"`
}

type HighActionObservation struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Whether to send a notification.
	Notify *bool `json:"notify,omitempty" tf:"notify,omitempty"`
}

type HighActionParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	// +kubebuilder:validation:Optional
	EventAction *string `json:"eventAction" tf:"event_action,omitempty"`

	// Whether to send a notification.
	// +kubebuilder:validation:Optional
	Notify *bool `json:"notify" tf:"notify,omitempty"`
}

type LowActionInitParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Whether to send a notification.
	Notify *bool `json:"notify,omitempty" tf:"notify,omitempty"`
}

type LowActionObservation struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Whether to send a notification.
	Notify *bool `json:"notify,omitempty" tf:"notify,omitempty"`
}

type LowActionParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	// +kubebuilder:validation:Optional
	EventAction *string `json:"eventAction" tf:"event_action,omitempty"`

	// Whether to send a notification.
	// +kubebuilder:validation:Optional
	Notify *bool `json:"notify" tf:"notify,omitempty"`
}

type MediumActionInitParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Whether to send a notification.
	Notify *bool `json:"notify,omitempty" tf:"notify,omitempty"`
}

type MediumActionObservation struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	EventAction *string `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Whether to send a notification.
	Notify *bool `json:"notify,omitempty" tf:"notify,omitempty"`
}

type MediumActionParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	// +kubebuilder:validation:Optional
	EventAction *string `json:"eventAction" tf:"event_action,omitempty"`

	// Whether to send a notification.
	// +kubebuilder:validation:Optional
	Notify *bool `json:"notify" tf:"notify,omitempty"`
}

type MfaEmailInitParameters struct {

	// The email HTML body.
	HTMLBody *string `json:"htmlBody,omitempty" tf:"html_body,omitempty"`

	// The email subject.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email text body.
	TextBody *string `json:"textBody,omitempty" tf:"text_body,omitempty"`
}

type MfaEmailObservation struct {

	// The email HTML body.
	HTMLBody *string `json:"htmlBody,omitempty" tf:"html_body,omitempty"`

	// The email subject.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email text body.
	TextBody *string `json:"textBody,omitempty" tf:"text_body,omitempty"`
}

type MfaEmailParameters struct {

	// The email HTML body.
	// +kubebuilder:validation:Optional
	HTMLBody *string `json:"htmlBody" tf:"html_body,omitempty"`

	// The email subject.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject" tf:"subject,omitempty"`

	// The email text body.
	// +kubebuilder:validation:Optional
	TextBody *string `json:"textBody" tf:"text_body,omitempty"`
}

type NoActionEmailInitParameters struct {

	// The email HTML body.
	HTMLBody *string `json:"htmlBody,omitempty" tf:"html_body,omitempty"`

	// The email subject.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email text body.
	TextBody *string `json:"textBody,omitempty" tf:"text_body,omitempty"`
}

type NoActionEmailObservation struct {

	// The email HTML body.
	HTMLBody *string `json:"htmlBody,omitempty" tf:"html_body,omitempty"`

	// The email subject.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// The email text body.
	TextBody *string `json:"textBody,omitempty" tf:"text_body,omitempty"`
}

type NoActionEmailParameters struct {

	// The email HTML body.
	// +kubebuilder:validation:Optional
	HTMLBody *string `json:"htmlBody" tf:"html_body,omitempty"`

	// The email subject.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject" tf:"subject,omitempty"`

	// The email text body.
	// +kubebuilder:validation:Optional
	TextBody *string `json:"textBody" tf:"text_body,omitempty"`
}

type NotifyConfigurationInitParameters struct {

	// Email template used when a detected risk event is blocked. See notify email type below.
	BlockEmail []BlockEmailInitParameters `json:"blockEmail,omitempty" tf:"block_email,omitempty"`

	// The email address that is sending the email. The address must be either individually verified with Amazon Simple Email Service, or from a domain that has been verified with Amazon SES.
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// The multi-factor authentication (MFA) email template used when MFA is challenged as part of a detected risk. See notify email type below.
	MfaEmail []MfaEmailInitParameters `json:"mfaEmail,omitempty" tf:"mfa_email,omitempty"`

	// The email template used when a detected risk event is allowed. See notify email type below.
	NoActionEmail []NoActionEmailInitParameters `json:"noActionEmail,omitempty" tf:"no_action_email,omitempty"`

	// The destination to which the receiver of an email should reply to.
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`

	// The Amazon Resource Name (ARN) of the identity that is associated with the sending authorization policy. This identity permits Amazon Cognito to send for the email address specified in the From parameter.
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`
}

type NotifyConfigurationObservation struct {

	// Email template used when a detected risk event is blocked. See notify email type below.
	BlockEmail []BlockEmailObservation `json:"blockEmail,omitempty" tf:"block_email,omitempty"`

	// The email address that is sending the email. The address must be either individually verified with Amazon Simple Email Service, or from a domain that has been verified with Amazon SES.
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// The multi-factor authentication (MFA) email template used when MFA is challenged as part of a detected risk. See notify email type below.
	MfaEmail []MfaEmailObservation `json:"mfaEmail,omitempty" tf:"mfa_email,omitempty"`

	// The email template used when a detected risk event is allowed. See notify email type below.
	NoActionEmail []NoActionEmailObservation `json:"noActionEmail,omitempty" tf:"no_action_email,omitempty"`

	// The destination to which the receiver of an email should reply to.
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`

	// The Amazon Resource Name (ARN) of the identity that is associated with the sending authorization policy. This identity permits Amazon Cognito to send for the email address specified in the From parameter.
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`
}

type NotifyConfigurationParameters struct {

	// Email template used when a detected risk event is blocked. See notify email type below.
	// +kubebuilder:validation:Optional
	BlockEmail []BlockEmailParameters `json:"blockEmail,omitempty" tf:"block_email,omitempty"`

	// The email address that is sending the email. The address must be either individually verified with Amazon Simple Email Service, or from a domain that has been verified with Amazon SES.
	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// The multi-factor authentication (MFA) email template used when MFA is challenged as part of a detected risk. See notify email type below.
	// +kubebuilder:validation:Optional
	MfaEmail []MfaEmailParameters `json:"mfaEmail,omitempty" tf:"mfa_email,omitempty"`

	// The email template used when a detected risk event is allowed. See notify email type below.
	// +kubebuilder:validation:Optional
	NoActionEmail []NoActionEmailParameters `json:"noActionEmail,omitempty" tf:"no_action_email,omitempty"`

	// The destination to which the receiver of an email should reply to.
	// +kubebuilder:validation:Optional
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`

	// The Amazon Resource Name (ARN) of the identity that is associated with the sending authorization policy. This identity permits Amazon Cognito to send for the email address specified in the From parameter.
	// +kubebuilder:validation:Optional
	SourceArn *string `json:"sourceArn" tf:"source_arn,omitempty"`
}

type RiskConfigurationInitParameters struct {

	// The account takeover risk configuration. See details below.
	AccountTakeoverRiskConfiguration []AccountTakeoverRiskConfigurationInitParameters `json:"accountTakeoverRiskConfiguration,omitempty" tf:"account_takeover_risk_configuration,omitempty"`

	// The app client ID. When the client ID is not provided, the same risk configuration is applied to all the clients in the User Pool.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The compromised credentials risk configuration. See details below.
	CompromisedCredentialsRiskConfiguration []CompromisedCredentialsRiskConfigurationInitParameters `json:"compromisedCredentialsRiskConfiguration,omitempty" tf:"compromised_credentials_risk_configuration,omitempty"`

	// The configuration to override the risk decision. See details below.
	RiskExceptionConfiguration []RiskExceptionConfigurationInitParameters `json:"riskExceptionConfiguration,omitempty" tf:"risk_exception_configuration,omitempty"`

	// The user pool ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

type RiskConfigurationObservation struct {

	// The account takeover risk configuration. See details below.
	AccountTakeoverRiskConfiguration []AccountTakeoverRiskConfigurationObservation `json:"accountTakeoverRiskConfiguration,omitempty" tf:"account_takeover_risk_configuration,omitempty"`

	// The app client ID. When the client ID is not provided, the same risk configuration is applied to all the clients in the User Pool.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The compromised credentials risk configuration. See details below.
	CompromisedCredentialsRiskConfiguration []CompromisedCredentialsRiskConfigurationObservation `json:"compromisedCredentialsRiskConfiguration,omitempty" tf:"compromised_credentials_risk_configuration,omitempty"`

	// The user pool ID or the user pool ID and Client Id separated by a : if the configuration is client specific.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The configuration to override the risk decision. See details below.
	RiskExceptionConfiguration []RiskExceptionConfigurationObservation `json:"riskExceptionConfiguration,omitempty" tf:"risk_exception_configuration,omitempty"`

	// The user pool ID.
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`
}

type RiskConfigurationParameters struct {

	// The account takeover risk configuration. See details below.
	// +kubebuilder:validation:Optional
	AccountTakeoverRiskConfiguration []AccountTakeoverRiskConfigurationParameters `json:"accountTakeoverRiskConfiguration,omitempty" tf:"account_takeover_risk_configuration,omitempty"`

	// The app client ID. When the client ID is not provided, the same risk configuration is applied to all the clients in the User Pool.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The compromised credentials risk configuration. See details below.
	// +kubebuilder:validation:Optional
	CompromisedCredentialsRiskConfiguration []CompromisedCredentialsRiskConfigurationParameters `json:"compromisedCredentialsRiskConfiguration,omitempty" tf:"compromised_credentials_risk_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The configuration to override the risk decision. See details below.
	// +kubebuilder:validation:Optional
	RiskExceptionConfiguration []RiskExceptionConfigurationParameters `json:"riskExceptionConfiguration,omitempty" tf:"risk_exception_configuration,omitempty"`

	// The user pool ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

type RiskExceptionConfigurationInitParameters struct {

	// Overrides the risk decision to always block the pre-authentication requests.
	// The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
	// Can contain a maximum of 200 items.
	// +listType=set
	BlockedIPRangeList []*string `json:"blockedIpRangeList,omitempty" tf:"blocked_ip_range_list,omitempty"`

	// Risk detection isn't performed on the IP addresses in this range list.
	// The IP range is in CIDR notation.
	// Can contain a maximum of 200 items.
	// +listType=set
	SkippedIPRangeList []*string `json:"skippedIpRangeList,omitempty" tf:"skipped_ip_range_list,omitempty"`
}

type RiskExceptionConfigurationObservation struct {

	// Overrides the risk decision to always block the pre-authentication requests.
	// The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
	// Can contain a maximum of 200 items.
	// +listType=set
	BlockedIPRangeList []*string `json:"blockedIpRangeList,omitempty" tf:"blocked_ip_range_list,omitempty"`

	// Risk detection isn't performed on the IP addresses in this range list.
	// The IP range is in CIDR notation.
	// Can contain a maximum of 200 items.
	// +listType=set
	SkippedIPRangeList []*string `json:"skippedIpRangeList,omitempty" tf:"skipped_ip_range_list,omitempty"`
}

type RiskExceptionConfigurationParameters struct {

	// Overrides the risk decision to always block the pre-authentication requests.
	// The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
	// Can contain a maximum of 200 items.
	// +kubebuilder:validation:Optional
	// +listType=set
	BlockedIPRangeList []*string `json:"blockedIpRangeList,omitempty" tf:"blocked_ip_range_list,omitempty"`

	// Risk detection isn't performed on the IP addresses in this range list.
	// The IP range is in CIDR notation.
	// Can contain a maximum of 200 items.
	// +kubebuilder:validation:Optional
	// +listType=set
	SkippedIPRangeList []*string `json:"skippedIpRangeList,omitempty" tf:"skipped_ip_range_list,omitempty"`
}

// RiskConfigurationSpec defines the desired state of RiskConfiguration
type RiskConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RiskConfigurationParameters `json:"forProvider"`
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
	InitProvider RiskConfigurationInitParameters `json:"initProvider,omitempty"`
}

// RiskConfigurationStatus defines the observed state of RiskConfiguration.
type RiskConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RiskConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RiskConfiguration is the Schema for the RiskConfigurations API. Provides a Cognito Risk Configuration resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RiskConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RiskConfigurationSpec   `json:"spec"`
	Status            RiskConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RiskConfigurationList contains a list of RiskConfigurations
type RiskConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RiskConfiguration `json:"items"`
}

// Repository type metadata.
var (
	RiskConfiguration_Kind             = "RiskConfiguration"
	RiskConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RiskConfiguration_Kind}.String()
	RiskConfiguration_KindAPIVersion   = RiskConfiguration_Kind + "." + CRDGroupVersion.String()
	RiskConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(RiskConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&RiskConfiguration{}, &RiskConfigurationList{})
}
