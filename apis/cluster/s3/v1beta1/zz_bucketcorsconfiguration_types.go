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

type BucketCorsConfigurationCorsRuleInitParameters struct {

	// Set of Headers that are specified in the Access-Control-Request-Headers header.
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Set of HTTP methods that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, and DELETE.
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Set of origins you want customers to be able to access the bucket from.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// Set of headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).
	// +listType=set
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type BucketCorsConfigurationCorsRuleObservation struct {

	// Set of Headers that are specified in the Access-Control-Request-Headers header.
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Set of HTTP methods that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, and DELETE.
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Set of origins you want customers to be able to access the bucket from.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// Set of headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).
	// +listType=set
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type BucketCorsConfigurationCorsRuleParameters struct {

	// Set of Headers that are specified in the Access-Control-Request-Headers header.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Set of HTTP methods that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, and DELETE.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// Set of origins you want customers to be able to access the bucket from.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Set of headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).
	// +kubebuilder:validation:Optional
	// +listType=set
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Time in seconds that your browser is to cache the preflight response for the specified resource.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type BucketCorsConfigurationInitParameters struct {

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

	// Set of origins and methods (cross-origin access that you want to allow). See below. You can configure up to 100 rules.
	CorsRule []BucketCorsConfigurationCorsRuleInitParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`
}

type BucketCorsConfigurationObservation struct {

	// Name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Set of origins and methods (cross-origin access that you want to allow). See below. You can configure up to 100 rules.
	CorsRule []BucketCorsConfigurationCorsRuleObservation `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketCorsConfigurationParameters struct {

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

	// Set of origins and methods (cross-origin access that you want to allow). See below. You can configure up to 100 rules.
	// +kubebuilder:validation:Optional
	CorsRule []BucketCorsConfigurationCorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// Account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// BucketCorsConfigurationSpec defines the desired state of BucketCorsConfiguration
type BucketCorsConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketCorsConfigurationParameters `json:"forProvider"`
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
	InitProvider BucketCorsConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketCorsConfigurationStatus defines the observed state of BucketCorsConfiguration.
type BucketCorsConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketCorsConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketCorsConfiguration is the Schema for the BucketCorsConfigurations API. Provides an S3 bucket CORS configuration resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketCorsConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.corsRule) || (has(self.initProvider) && has(self.initProvider.corsRule))",message="spec.forProvider.corsRule is a required parameter"
	Spec   BucketCorsConfigurationSpec   `json:"spec"`
	Status BucketCorsConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketCorsConfigurationList contains a list of BucketCorsConfigurations
type BucketCorsConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketCorsConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketCorsConfiguration_Kind             = "BucketCorsConfiguration"
	BucketCorsConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketCorsConfiguration_Kind}.String()
	BucketCorsConfiguration_KindAPIVersion   = BucketCorsConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketCorsConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketCorsConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketCorsConfiguration{}, &BucketCorsConfigurationList{})
}
