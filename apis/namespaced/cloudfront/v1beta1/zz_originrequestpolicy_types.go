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

type CookiesConfigCookiesInitParameters struct {

	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type CookiesConfigCookiesObservation struct {

	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type CookiesConfigCookiesParameters struct {

	// +kubebuilder:validation:Optional
	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersConfigHeadersInitParameters struct {

	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersConfigHeadersObservation struct {

	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersConfigHeadersParameters struct {

	// +kubebuilder:validation:Optional
	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type OriginRequestPolicyCookiesConfigInitParameters struct {
	CookieBehavior *string `json:"cookieBehavior,omitempty" tf:"cookie_behavior,omitempty"`

	Cookies []CookiesConfigCookiesInitParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type OriginRequestPolicyCookiesConfigObservation struct {
	CookieBehavior *string `json:"cookieBehavior,omitempty" tf:"cookie_behavior,omitempty"`

	Cookies []CookiesConfigCookiesObservation `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type OriginRequestPolicyCookiesConfigParameters struct {

	// +kubebuilder:validation:Optional
	CookieBehavior *string `json:"cookieBehavior" tf:"cookie_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Cookies []CookiesConfigCookiesParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type OriginRequestPolicyHeadersConfigInitParameters struct {
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	Headers []HeadersConfigHeadersInitParameters `json:"headers,omitempty" tf:"headers,omitempty"`
}

type OriginRequestPolicyHeadersConfigObservation struct {
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	Headers []HeadersConfigHeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`
}

type OriginRequestPolicyHeadersConfigParameters struct {

	// +kubebuilder:validation:Optional
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Headers []HeadersConfigHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`
}

type OriginRequestPolicyInitParameters struct {

	// Comment to describe the origin request policy.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig []OriginRequestPolicyCookiesConfigInitParameters `json:"cookiesConfig,omitempty" tf:"cookies_config,omitempty"`

	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig []OriginRequestPolicyHeadersConfigInitParameters `json:"headersConfig,omitempty" tf:"headers_config,omitempty"`

	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig []OriginRequestPolicyQueryStringsConfigInitParameters `json:"queryStringsConfig,omitempty" tf:"query_strings_config,omitempty"`
}

type OriginRequestPolicyObservation struct {

	// Comment to describe the origin request policy.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig []OriginRequestPolicyCookiesConfigObservation `json:"cookiesConfig,omitempty" tf:"cookies_config,omitempty"`

	// The current version of the origin request policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig []OriginRequestPolicyHeadersConfigObservation `json:"headersConfig,omitempty" tf:"headers_config,omitempty"`

	// The identifier for the origin request policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig []OriginRequestPolicyQueryStringsConfigObservation `json:"queryStringsConfig,omitempty" tf:"query_strings_config,omitempty"`
}

type OriginRequestPolicyParameters struct {

	// Comment to describe the origin request policy.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	// +kubebuilder:validation:Optional
	CookiesConfig []OriginRequestPolicyCookiesConfigParameters `json:"cookiesConfig,omitempty" tf:"cookies_config,omitempty"`

	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	// +kubebuilder:validation:Optional
	HeadersConfig []OriginRequestPolicyHeadersConfigParameters `json:"headersConfig,omitempty" tf:"headers_config,omitempty"`

	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	// +kubebuilder:validation:Optional
	QueryStringsConfig []OriginRequestPolicyQueryStringsConfigParameters `json:"queryStringsConfig,omitempty" tf:"query_strings_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type OriginRequestPolicyQueryStringsConfigInitParameters struct {
	QueryStringBehavior *string `json:"queryStringBehavior,omitempty" tf:"query_string_behavior,omitempty"`

	QueryStrings []QueryStringsConfigQueryStringsInitParameters `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type OriginRequestPolicyQueryStringsConfigObservation struct {
	QueryStringBehavior *string `json:"queryStringBehavior,omitempty" tf:"query_string_behavior,omitempty"`

	QueryStrings []QueryStringsConfigQueryStringsObservation `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type OriginRequestPolicyQueryStringsConfigParameters struct {

	// +kubebuilder:validation:Optional
	QueryStringBehavior *string `json:"queryStringBehavior" tf:"query_string_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	QueryStrings []QueryStringsConfigQueryStringsParameters `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type QueryStringsConfigQueryStringsInitParameters struct {

	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type QueryStringsConfigQueryStringsObservation struct {

	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type QueryStringsConfigQueryStringsParameters struct {

	// +kubebuilder:validation:Optional
	// +listType=set
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

// OriginRequestPolicySpec defines the desired state of OriginRequestPolicy
type OriginRequestPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginRequestPolicyParameters `json:"forProvider"`
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
	InitProvider OriginRequestPolicyInitParameters `json:"initProvider,omitempty"`
}

// OriginRequestPolicyStatus defines the observed state of OriginRequestPolicy.
type OriginRequestPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginRequestPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OriginRequestPolicy is the Schema for the OriginRequestPolicys API. Determines the values that CloudFront includes in requests that it sends to the origin.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type OriginRequestPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cookiesConfig) || (has(self.initProvider) && has(self.initProvider.cookiesConfig))",message="spec.forProvider.cookiesConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.headersConfig) || (has(self.initProvider) && has(self.initProvider.headersConfig))",message="spec.forProvider.headersConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.queryStringsConfig) || (has(self.initProvider) && has(self.initProvider.queryStringsConfig))",message="spec.forProvider.queryStringsConfig is a required parameter"
	Spec   OriginRequestPolicySpec   `json:"spec"`
	Status OriginRequestPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginRequestPolicyList contains a list of OriginRequestPolicys
type OriginRequestPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginRequestPolicy `json:"items"`
}

// Repository type metadata.
var (
	OriginRequestPolicy_Kind             = "OriginRequestPolicy"
	OriginRequestPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginRequestPolicy_Kind}.String()
	OriginRequestPolicy_KindAPIVersion   = OriginRequestPolicy_Kind + "." + CRDGroupVersion.String()
	OriginRequestPolicy_GroupVersionKind = CRDGroupVersion.WithKind(OriginRequestPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&OriginRequestPolicy{}, &OriginRequestPolicyList{})
}
