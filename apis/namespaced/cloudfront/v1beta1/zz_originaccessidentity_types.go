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

type OriginAccessIdentityInitParameters struct {

	// An optional comment for the origin access identity.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`
}

type OriginAccessIdentityObservation struct {

	// Internal value used by CloudFront to allow future
	// updates to the origin access identity.
	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`

	// A shortcut to the full path for the
	// origin access identity to use in CloudFront, see below.
	CloudfrontAccessIdentityPath *string `json:"cloudfrontAccessIdentityPath,omitempty" tf:"cloudfront_access_identity_path,omitempty"`

	// An optional comment for the origin access identity.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The current version of the origin access identity's information.
	// For example: E2QWRUHAPOMQZL.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A pre-generated ARN for use in S3 bucket policies (see below).
	// Example: arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity E2QWRUHAPOMQZL.
	IAMArn *string `json:"iamArn,omitempty" tf:"iam_arn,omitempty"`

	// The identifier for the distribution. For example: EDFDVBD632BHDS5.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon S3 canonical user ID for the origin
	// access identity, which you use when giving the origin access identity read
	// permission to an object in Amazon S3.
	S3CanonicalUserID *string `json:"s3CanonicalUserId,omitempty" tf:"s3_canonical_user_id,omitempty"`
}

type OriginAccessIdentityParameters struct {

	// An optional comment for the origin access identity.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// OriginAccessIdentitySpec defines the desired state of OriginAccessIdentity
type OriginAccessIdentitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginAccessIdentityParameters `json:"forProvider"`
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
	InitProvider OriginAccessIdentityInitParameters `json:"initProvider,omitempty"`
}

// OriginAccessIdentityStatus defines the observed state of OriginAccessIdentity.
type OriginAccessIdentityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginAccessIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OriginAccessIdentity is the Schema for the OriginAccessIdentitys API. Provides a CloudFront origin access identity.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OriginAccessIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OriginAccessIdentitySpec   `json:"spec"`
	Status            OriginAccessIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginAccessIdentityList contains a list of OriginAccessIdentitys
type OriginAccessIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginAccessIdentity `json:"items"`
}

// Repository type metadata.
var (
	OriginAccessIdentity_Kind             = "OriginAccessIdentity"
	OriginAccessIdentity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginAccessIdentity_Kind}.String()
	OriginAccessIdentity_KindAPIVersion   = OriginAccessIdentity_Kind + "." + CRDGroupVersion.String()
	OriginAccessIdentity_GroupVersionKind = CRDGroupVersion.WithKind(OriginAccessIdentity_Kind)
)

func init() {
	SchemeBuilder.Register(&OriginAccessIdentity{}, &OriginAccessIdentityList{})
}
