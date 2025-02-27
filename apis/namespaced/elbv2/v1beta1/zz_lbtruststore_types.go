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

type LBTrustStoreInitParameters struct {

	// S3 Bucket name holding the client certificate CA bundle.
	CACertificatesBundleS3Bucket *string `json:"caCertificatesBundleS3Bucket,omitempty" tf:"ca_certificates_bundle_s3_bucket,omitempty"`

	// S3 object key holding the client certificate CA bundle.
	CACertificatesBundleS3Key *string `json:"caCertificatesBundleS3Key,omitempty" tf:"ca_certificates_bundle_s3_key,omitempty"`

	// Version Id of CA bundle S3 bucket object, if versioned, defaults to latest if omitted.
	CACertificatesBundleS3ObjectVersion *string `json:"caCertificatesBundleS3ObjectVersion,omitempty" tf:"ca_certificates_bundle_s3_object_version,omitempty"`

	// Name of the Trust Store. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LBTrustStoreObservation struct {

	// ARN of the Trust Store (matches id).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN suffix for use with CloudWatch Metrics.
	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix,omitempty"`

	// S3 Bucket name holding the client certificate CA bundle.
	CACertificatesBundleS3Bucket *string `json:"caCertificatesBundleS3Bucket,omitempty" tf:"ca_certificates_bundle_s3_bucket,omitempty"`

	// S3 object key holding the client certificate CA bundle.
	CACertificatesBundleS3Key *string `json:"caCertificatesBundleS3Key,omitempty" tf:"ca_certificates_bundle_s3_key,omitempty"`

	// Version Id of CA bundle S3 bucket object, if versioned, defaults to latest if omitted.
	CACertificatesBundleS3ObjectVersion *string `json:"caCertificatesBundleS3ObjectVersion,omitempty" tf:"ca_certificates_bundle_s3_object_version,omitempty"`

	// ARN of the Trust Store (matches arn).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Trust Store. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LBTrustStoreParameters struct {

	// S3 Bucket name holding the client certificate CA bundle.
	// +kubebuilder:validation:Optional
	CACertificatesBundleS3Bucket *string `json:"caCertificatesBundleS3Bucket,omitempty" tf:"ca_certificates_bundle_s3_bucket,omitempty"`

	// S3 object key holding the client certificate CA bundle.
	// +kubebuilder:validation:Optional
	CACertificatesBundleS3Key *string `json:"caCertificatesBundleS3Key,omitempty" tf:"ca_certificates_bundle_s3_key,omitempty"`

	// Version Id of CA bundle S3 bucket object, if versioned, defaults to latest if omitted.
	// +kubebuilder:validation:Optional
	CACertificatesBundleS3ObjectVersion *string `json:"caCertificatesBundleS3ObjectVersion,omitempty" tf:"ca_certificates_bundle_s3_object_version,omitempty"`

	// Name of the Trust Store. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LBTrustStoreSpec defines the desired state of LBTrustStore
type LBTrustStoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBTrustStoreParameters `json:"forProvider"`
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
	InitProvider LBTrustStoreInitParameters `json:"initProvider,omitempty"`
}

// LBTrustStoreStatus defines the observed state of LBTrustStore.
type LBTrustStoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBTrustStoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LBTrustStore is the Schema for the LBTrustStores API. Provides a Trust Store resource for use with Load Balancers.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type LBTrustStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.caCertificatesBundleS3Bucket) || (has(self.initProvider) && has(self.initProvider.caCertificatesBundleS3Bucket))",message="spec.forProvider.caCertificatesBundleS3Bucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.caCertificatesBundleS3Key) || (has(self.initProvider) && has(self.initProvider.caCertificatesBundleS3Key))",message="spec.forProvider.caCertificatesBundleS3Key is a required parameter"
	Spec   LBTrustStoreSpec   `json:"spec"`
	Status LBTrustStoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBTrustStoreList contains a list of LBTrustStores
type LBTrustStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBTrustStore `json:"items"`
}

// Repository type metadata.
var (
	LBTrustStore_Kind             = "LBTrustStore"
	LBTrustStore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBTrustStore_Kind}.String()
	LBTrustStore_KindAPIVersion   = LBTrustStore_Kind + "." + CRDGroupVersion.String()
	LBTrustStore_GroupVersionKind = CRDGroupVersion.WithKind(LBTrustStore_Kind)
)

func init() {
	SchemeBuilder.Register(&LBTrustStore{}, &LBTrustStoreList{})
}
