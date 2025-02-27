// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BucketWebsiteConfigurationInitParameters struct {

	// Name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Name of the error document for the website. See below.
	ErrorDocument *ErrorDocumentInitParameters `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Name of the index document for the website. See below.
	IndexDocument *IndexDocumentInitParameters `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// Redirect behavior for every request to this bucket's website endpoint. See below. Conflicts with error_document, index_document, and routing_rule.
	RedirectAllRequestsTo *RedirectAllRequestsToInitParameters `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// List of rules that define when a redirect is applied and the redirect behavior. See below.
	RoutingRule []RoutingRuleInitParameters `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// JSON array containing routing rules
	// describing redirect behavior and when redirects are applied. Use this parameter when your routing rules contain empty String values ("") as seen in the example above.
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type BucketWebsiteConfigurationObservation struct {

	// Name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Name of the error document for the website. See below.
	ErrorDocument *ErrorDocumentObservation `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the index document for the website. See below.
	IndexDocument *IndexDocumentObservation `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// Redirect behavior for every request to this bucket's website endpoint. See below. Conflicts with error_document, index_document, and routing_rule.
	RedirectAllRequestsTo *RedirectAllRequestsToObservation `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// List of rules that define when a redirect is applied and the redirect behavior. See below.
	RoutingRule []RoutingRuleObservation `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// JSON array containing routing rules
	// describing redirect behavior and when redirects are applied. Use this parameter when your routing rules contain empty String values ("") as seen in the example above.
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`

	// Domain of the website endpoint. This is used to create Route 53 alias records.
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// Website endpoint.
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type BucketWebsiteConfigurationParameters struct {

	// Name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Name of the error document for the website. See below.
	// +kubebuilder:validation:Optional
	ErrorDocument *ErrorDocumentParameters `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Name of the index document for the website. See below.
	// +kubebuilder:validation:Optional
	IndexDocument *IndexDocumentParameters `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// Redirect behavior for every request to this bucket's website endpoint. See below. Conflicts with error_document, index_document, and routing_rule.
	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo *RedirectAllRequestsToParameters `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of rules that define when a redirect is applied and the redirect behavior. See below.
	// +kubebuilder:validation:Optional
	RoutingRule []RoutingRuleParameters `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// JSON array containing routing rules
	// describing redirect behavior and when redirects are applied. Use this parameter when your routing rules contain empty String values ("") as seen in the example above.
	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type ConditionInitParameters struct {

	// HTTP error code when the redirect is applied. If specified with key_prefix_equals, then both must be true for the redirect to be applied.
	HTTPErrorCodeReturnedEquals *string `json:"httpErrorCodeReturnedEquals,omitempty" tf:"http_error_code_returned_equals,omitempty"`

	// Object key name prefix when the redirect is applied. If specified with http_error_code_returned_equals, then both must be true for the redirect to be applied.
	KeyPrefixEquals *string `json:"keyPrefixEquals,omitempty" tf:"key_prefix_equals,omitempty"`
}

type ConditionObservation struct {

	// HTTP error code when the redirect is applied. If specified with key_prefix_equals, then both must be true for the redirect to be applied.
	HTTPErrorCodeReturnedEquals *string `json:"httpErrorCodeReturnedEquals,omitempty" tf:"http_error_code_returned_equals,omitempty"`

	// Object key name prefix when the redirect is applied. If specified with http_error_code_returned_equals, then both must be true for the redirect to be applied.
	KeyPrefixEquals *string `json:"keyPrefixEquals,omitempty" tf:"key_prefix_equals,omitempty"`
}

type ConditionParameters struct {

	// HTTP error code when the redirect is applied. If specified with key_prefix_equals, then both must be true for the redirect to be applied.
	// +kubebuilder:validation:Optional
	HTTPErrorCodeReturnedEquals *string `json:"httpErrorCodeReturnedEquals,omitempty" tf:"http_error_code_returned_equals,omitempty"`

	// Object key name prefix when the redirect is applied. If specified with http_error_code_returned_equals, then both must be true for the redirect to be applied.
	// +kubebuilder:validation:Optional
	KeyPrefixEquals *string `json:"keyPrefixEquals,omitempty" tf:"key_prefix_equals,omitempty"`
}

type ErrorDocumentInitParameters struct {

	// Object key name to use when a 4XX class error occurs.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type ErrorDocumentObservation struct {

	// Object key name to use when a 4XX class error occurs.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type ErrorDocumentParameters struct {

	// Object key name to use when a 4XX class error occurs.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`
}

type IndexDocumentInitParameters struct {

	// Suffix that is appended to a request that is for a directory on the website endpoint.
	// For example, if the suffix is index.html and you make a request to samplebucket/images/, the data that is returned will be for the object with the key name images/index.html.
	// The suffix must not be empty and must not include a slash character.
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type IndexDocumentObservation struct {

	// Suffix that is appended to a request that is for a directory on the website endpoint.
	// For example, if the suffix is index.html and you make a request to samplebucket/images/, the data that is returned will be for the object with the key name images/index.html.
	// The suffix must not be empty and must not include a slash character.
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type IndexDocumentParameters struct {

	// Suffix that is appended to a request that is for a directory on the website endpoint.
	// For example, if the suffix is index.html and you make a request to samplebucket/images/, the data that is returned will be for the object with the key name images/index.html.
	// The suffix must not be empty and must not include a slash character.
	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix" tf:"suffix,omitempty"`
}

type RedirectAllRequestsToInitParameters struct {

	// Name of the host where requests are redirected.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type RedirectAllRequestsToObservation struct {

	// Name of the host where requests are redirected.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type RedirectAllRequestsToParameters struct {

	// Name of the host where requests are redirected.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type RedirectInitParameters struct {

	// HTTP redirect code to use on the response.
	HTTPRedirectCode *string `json:"httpRedirectCode,omitempty" tf:"http_redirect_code,omitempty"`

	// Name of the host where requests are redirected.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix docs/ (objects in the docs/ folder) to documents/, you can set a condition block with key_prefix_equals set to docs/ and in the redirect set replace_key_prefix_with to /documents.
	ReplaceKeyPrefixWith *string `json:"replaceKeyPrefixWith,omitempty" tf:"replace_key_prefix_with,omitempty"`

	// Specific object key to use in the redirect request. For example, redirect request to error.html.
	ReplaceKeyWith *string `json:"replaceKeyWith,omitempty" tf:"replace_key_with,omitempty"`
}

type RedirectObservation struct {

	// HTTP redirect code to use on the response.
	HTTPRedirectCode *string `json:"httpRedirectCode,omitempty" tf:"http_redirect_code,omitempty"`

	// Name of the host where requests are redirected.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix docs/ (objects in the docs/ folder) to documents/, you can set a condition block with key_prefix_equals set to docs/ and in the redirect set replace_key_prefix_with to /documents.
	ReplaceKeyPrefixWith *string `json:"replaceKeyPrefixWith,omitempty" tf:"replace_key_prefix_with,omitempty"`

	// Specific object key to use in the redirect request. For example, redirect request to error.html.
	ReplaceKeyWith *string `json:"replaceKeyWith,omitempty" tf:"replace_key_with,omitempty"`
}

type RedirectParameters struct {

	// HTTP redirect code to use on the response.
	// +kubebuilder:validation:Optional
	HTTPRedirectCode *string `json:"httpRedirectCode,omitempty" tf:"http_redirect_code,omitempty"`

	// Name of the host where requests are redirected.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix docs/ (objects in the docs/ folder) to documents/, you can set a condition block with key_prefix_equals set to docs/ and in the redirect set replace_key_prefix_with to /documents.
	// +kubebuilder:validation:Optional
	ReplaceKeyPrefixWith *string `json:"replaceKeyPrefixWith,omitempty" tf:"replace_key_prefix_with,omitempty"`

	// Specific object key to use in the redirect request. For example, redirect request to error.html.
	// +kubebuilder:validation:Optional
	ReplaceKeyWith *string `json:"replaceKeyWith,omitempty" tf:"replace_key_with,omitempty"`
}

type RoutingRuleInitParameters struct {

	// Configuration block for describing a condition that must be met for the specified redirect to apply. See below.
	Condition *ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Configuration block for redirect information. See below.
	Redirect *RedirectInitParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`
}

type RoutingRuleObservation struct {

	// Configuration block for describing a condition that must be met for the specified redirect to apply. See below.
	Condition *ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Configuration block for redirect information. See below.
	Redirect *RedirectObservation `json:"redirect,omitempty" tf:"redirect,omitempty"`
}

type RoutingRuleParameters struct {

	// Configuration block for describing a condition that must be met for the specified redirect to apply. See below.
	// +kubebuilder:validation:Optional
	Condition *ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Configuration block for redirect information. See below.
	// +kubebuilder:validation:Optional
	Redirect *RedirectParameters `json:"redirect" tf:"redirect,omitempty"`
}

// BucketWebsiteConfigurationSpec defines the desired state of BucketWebsiteConfiguration
type BucketWebsiteConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketWebsiteConfigurationParameters `json:"forProvider"`
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
	InitProvider BucketWebsiteConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketWebsiteConfigurationStatus defines the observed state of BucketWebsiteConfiguration.
type BucketWebsiteConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketWebsiteConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BucketWebsiteConfiguration is the Schema for the BucketWebsiteConfigurations API. Provides an S3 bucket website configuration resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type BucketWebsiteConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketWebsiteConfigurationSpec   `json:"spec"`
	Status            BucketWebsiteConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketWebsiteConfigurationList contains a list of BucketWebsiteConfigurations
type BucketWebsiteConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketWebsiteConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketWebsiteConfiguration_Kind             = "BucketWebsiteConfiguration"
	BucketWebsiteConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketWebsiteConfiguration_Kind}.String()
	BucketWebsiteConfiguration_KindAPIVersion   = BucketWebsiteConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketWebsiteConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketWebsiteConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketWebsiteConfiguration{}, &BucketWebsiteConfigurationList{})
}
