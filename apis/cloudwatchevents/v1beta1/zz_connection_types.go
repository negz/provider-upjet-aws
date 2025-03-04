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

type APIKeyInitParameters struct {

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type APIKeyObservation struct {

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type APIKeyParameters struct {

	// Header Name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Required
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type AuthParametersInitParameters struct {

	// Parameters used for API_KEY authorization. An API key to include in the header for each authentication request. A maximum of 1 are allowed. Conflicts with basic and oauth. Documented below.
	APIKey []APIKeyInitParameters `json:"apiKey,omitempty" tf:"api_key,omitempty"`

	// Parameters used for BASIC authorization. A maximum of 1 are allowed. Conflicts with api_key and oauth. Documented below.
	Basic []BasicInitParameters `json:"basic,omitempty" tf:"basic,omitempty"`

	// Invocation Http Parameters are additional credentials used to sign each Invocation of the ApiDestination created from this Connection. If the ApiDestination Rule Target has additional HttpParameters, the values will be merged together, with the Connection Invocation Http Parameters taking precedence. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
	InvocationHTTPParameters []InvocationHTTPParametersInitParameters `json:"invocationHttpParameters,omitempty" tf:"invocation_http_parameters,omitempty"`

	// Parameters used for OAUTH_CLIENT_CREDENTIALS authorization. A maximum of 1 are allowed. Conflicts with basic and api_key. Documented below.
	Oauth []OauthInitParameters `json:"oauth,omitempty" tf:"oauth,omitempty"`
}

type AuthParametersObservation struct {

	// Parameters used for API_KEY authorization. An API key to include in the header for each authentication request. A maximum of 1 are allowed. Conflicts with basic and oauth. Documented below.
	APIKey []APIKeyObservation `json:"apiKey,omitempty" tf:"api_key,omitempty"`

	// Parameters used for BASIC authorization. A maximum of 1 are allowed. Conflicts with api_key and oauth. Documented below.
	Basic []BasicObservation `json:"basic,omitempty" tf:"basic,omitempty"`

	// Invocation Http Parameters are additional credentials used to sign each Invocation of the ApiDestination created from this Connection. If the ApiDestination Rule Target has additional HttpParameters, the values will be merged together, with the Connection Invocation Http Parameters taking precedence. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
	InvocationHTTPParameters []InvocationHTTPParametersObservation `json:"invocationHttpParameters,omitempty" tf:"invocation_http_parameters,omitempty"`

	// Parameters used for OAUTH_CLIENT_CREDENTIALS authorization. A maximum of 1 are allowed. Conflicts with basic and api_key. Documented below.
	Oauth []OauthObservation `json:"oauth,omitempty" tf:"oauth,omitempty"`
}

type AuthParametersParameters struct {

	// Parameters used for API_KEY authorization. An API key to include in the header for each authentication request. A maximum of 1 are allowed. Conflicts with basic and oauth. Documented below.
	// +kubebuilder:validation:Optional
	APIKey []APIKeyParameters `json:"apiKey,omitempty" tf:"api_key,omitempty"`

	// Parameters used for BASIC authorization. A maximum of 1 are allowed. Conflicts with api_key and oauth. Documented below.
	// +kubebuilder:validation:Optional
	Basic []BasicParameters `json:"basic,omitempty" tf:"basic,omitempty"`

	// Invocation Http Parameters are additional credentials used to sign each Invocation of the ApiDestination created from this Connection. If the ApiDestination Rule Target has additional HttpParameters, the values will be merged together, with the Connection Invocation Http Parameters taking precedence. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
	// +kubebuilder:validation:Optional
	InvocationHTTPParameters []InvocationHTTPParametersParameters `json:"invocationHttpParameters,omitempty" tf:"invocation_http_parameters,omitempty"`

	// Parameters used for OAUTH_CLIENT_CREDENTIALS authorization. A maximum of 1 are allowed. Conflicts with basic and api_key. Documented below.
	// +kubebuilder:validation:Optional
	Oauth []OauthParameters `json:"oauth,omitempty" tf:"oauth,omitempty"`
}

type BasicInitParameters struct {

	// A username for the authorization.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BasicObservation struct {

	// A username for the authorization.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BasicParameters struct {

	// A password for the authorization. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// A username for the authorization.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type BodyInitParameters struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type BodyObservation struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type BodyParameters struct {

	// Specified whether the value is secret.
	// +kubebuilder:validation:Optional
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type ClientParametersInitParameters struct {

	// The client ID for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`
}

type ClientParametersObservation struct {

	// The client ID for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`
}

type ClientParametersParameters struct {

	// The client ID for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The client secret for the credentials to use for authorization. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`
}

type ConnectionInitParameters struct {

	// Parameters used for authorization. A maximum of 1 are allowed. Documented below.
	AuthParameters []AuthParametersInitParameters `json:"authParameters,omitempty" tf:"auth_parameters,omitempty"`

	// Choose the type of authorization to use for the connection. One of API_KEY,BASIC,OAUTH_CLIENT_CREDENTIALS.
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Enter a description for the connection. Maximum of 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type ConnectionObservation struct {

	// The Amazon Resource Name (ARN) of the connection.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Parameters used for authorization. A maximum of 1 are allowed. Documented below.
	AuthParameters []AuthParametersObservation `json:"authParameters,omitempty" tf:"auth_parameters,omitempty"`

	// Choose the type of authorization to use for the connection. One of API_KEY,BASIC,OAUTH_CLIENT_CREDENTIALS.
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Enter a description for the connection. Maximum of 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`
}

type ConnectionParameters struct {

	// Parameters used for authorization. A maximum of 1 are allowed. Documented below.
	// +kubebuilder:validation:Optional
	AuthParameters []AuthParametersParameters `json:"authParameters,omitempty" tf:"auth_parameters,omitempty"`

	// Choose the type of authorization to use for the connection. One of API_KEY,BASIC,OAUTH_CLIENT_CREDENTIALS.
	// +kubebuilder:validation:Optional
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Enter a description for the connection. Maximum of 512 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type HeaderInitParameters struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type HeaderObservation struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type HeaderParameters struct {

	// Specified whether the value is secret.
	// +kubebuilder:validation:Optional
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type InvocationHTTPParametersInitParameters struct {

	// Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Body []BodyInitParameters `json:"body,omitempty" tf:"body,omitempty"`

	// Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Header []HeaderInitParameters `json:"header,omitempty" tf:"header,omitempty"`

	// Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	QueryString []QueryStringInitParameters `json:"queryString,omitempty" tf:"query_string,omitempty"`
}

type InvocationHTTPParametersObservation struct {

	// Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Body []BodyObservation `json:"body,omitempty" tf:"body,omitempty"`

	// Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Header []HeaderObservation `json:"header,omitempty" tf:"header,omitempty"`

	// Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	QueryString []QueryStringObservation `json:"queryString,omitempty" tf:"query_string,omitempty"`
}

type InvocationHTTPParametersParameters struct {

	// Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	// +kubebuilder:validation:Optional
	Body []BodyParameters `json:"body,omitempty" tf:"body,omitempty"`

	// Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	// +kubebuilder:validation:Optional
	Header []HeaderParameters `json:"header,omitempty" tf:"header,omitempty"`

	// Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	// +kubebuilder:validation:Optional
	QueryString []QueryStringParameters `json:"queryString,omitempty" tf:"query_string,omitempty"`
}

type OauthHTTPParametersBodyInitParameters struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type OauthHTTPParametersBodyObservation struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type OauthHTTPParametersBodyParameters struct {

	// Specified whether the value is secret.
	// +kubebuilder:validation:Optional
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type OauthHTTPParametersHeaderInitParameters struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type OauthHTTPParametersHeaderObservation struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type OauthHTTPParametersHeaderParameters struct {

	// Specified whether the value is secret.
	// +kubebuilder:validation:Optional
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type OauthHTTPParametersInitParameters struct {

	// Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Body []OauthHTTPParametersBodyInitParameters `json:"body,omitempty" tf:"body,omitempty"`

	// Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Header []OauthHTTPParametersHeaderInitParameters `json:"header,omitempty" tf:"header,omitempty"`

	// Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	QueryString []OauthHTTPParametersQueryStringInitParameters `json:"queryString,omitempty" tf:"query_string,omitempty"`
}

type OauthHTTPParametersObservation struct {

	// Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Body []OauthHTTPParametersBodyObservation `json:"body,omitempty" tf:"body,omitempty"`

	// Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Header []OauthHTTPParametersHeaderObservation `json:"header,omitempty" tf:"header,omitempty"`

	// Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	QueryString []OauthHTTPParametersQueryStringObservation `json:"queryString,omitempty" tf:"query_string,omitempty"`
}

type OauthHTTPParametersParameters struct {

	// Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	// +kubebuilder:validation:Optional
	Body []OauthHTTPParametersBodyParameters `json:"body,omitempty" tf:"body,omitempty"`

	// Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	// +kubebuilder:validation:Optional
	Header []OauthHTTPParametersHeaderParameters `json:"header,omitempty" tf:"header,omitempty"`

	// Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	// +kubebuilder:validation:Optional
	QueryString []OauthHTTPParametersQueryStringParameters `json:"queryString,omitempty" tf:"query_string,omitempty"`
}

type OauthHTTPParametersQueryStringInitParameters struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type OauthHTTPParametersQueryStringObservation struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type OauthHTTPParametersQueryStringParameters struct {

	// Specified whether the value is secret.
	// +kubebuilder:validation:Optional
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type OauthInitParameters struct {

	// The URL to the authorization endpoint.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// Contains the client parameters for OAuth authorization. Contains the following two parameters.
	ClientParameters []ClientParametersInitParameters `json:"clientParameters,omitempty" tf:"client_parameters,omitempty"`

	// A password for the authorization. Created and stored in AWS Secrets Manager.
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// OAuth Http Parameters are additional credentials used to sign the request to the authorization endpoint to exchange the OAuth Client information for an access token. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
	OauthHTTPParameters []OauthHTTPParametersInitParameters `json:"oauthHttpParameters,omitempty" tf:"oauth_http_parameters,omitempty"`
}

type OauthObservation struct {

	// The URL to the authorization endpoint.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// Contains the client parameters for OAuth authorization. Contains the following two parameters.
	ClientParameters []ClientParametersObservation `json:"clientParameters,omitempty" tf:"client_parameters,omitempty"`

	// A password for the authorization. Created and stored in AWS Secrets Manager.
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// OAuth Http Parameters are additional credentials used to sign the request to the authorization endpoint to exchange the OAuth Client information for an access token. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
	OauthHTTPParameters []OauthHTTPParametersObservation `json:"oauthHttpParameters,omitempty" tf:"oauth_http_parameters,omitempty"`
}

type OauthParameters struct {

	// The URL to the authorization endpoint.
	// +kubebuilder:validation:Optional
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint,omitempty"`

	// Contains the client parameters for OAuth authorization. Contains the following two parameters.
	// +kubebuilder:validation:Optional
	ClientParameters []ClientParametersParameters `json:"clientParameters,omitempty" tf:"client_parameters,omitempty"`

	// A password for the authorization. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod" tf:"http_method,omitempty"`

	// OAuth Http Parameters are additional credentials used to sign the request to the authorization endpoint to exchange the OAuth Client information for an access token. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
	// +kubebuilder:validation:Optional
	OauthHTTPParameters []OauthHTTPParametersParameters `json:"oauthHttpParameters" tf:"oauth_http_parameters,omitempty"`
}

type QueryStringInitParameters struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type QueryStringObservation struct {

	// Specified whether the value is secret.
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type QueryStringParameters struct {

	// Specified whether the value is secret.
	// +kubebuilder:validation:Optional
	IsValueSecret *bool `json:"isValueSecret,omitempty" tf:"is_value_secret,omitempty"`

	// Header Name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
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
	InitProvider ConnectionInitParameters `json:"initProvider,omitempty"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Connection is the Schema for the Connections API. Provides an EventBridge connection resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authParameters) || (has(self.initProvider) && has(self.initProvider.authParameters))",message="spec.forProvider.authParameters is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authorizationType) || (has(self.initProvider) && has(self.initProvider.authorizationType))",message="spec.forProvider.authorizationType is a required parameter"
	Spec   ConnectionSpec   `json:"spec"`
	Status ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
