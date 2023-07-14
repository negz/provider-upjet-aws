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

type BucketOwnershipControlsObservation struct {

	// Name of the bucket that you want to associate this access point with.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// S3 Bucket name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block(s) with Ownership Controls rules. Detailed below.
	Rule []BucketOwnershipControlsRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketOwnershipControlsParameters struct {

	// Name of the bucket that you want to associate this access point with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block(s) with Ownership Controls rules. Detailed below.
	// +kubebuilder:validation:Optional
	Rule []BucketOwnershipControlsRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketOwnershipControlsRuleObservation struct {

	// Object ownership. Valid values: BucketOwnerPreferred, ObjectWriter or BucketOwnerEnforced
	ObjectOwnership *string `json:"objectOwnership,omitempty" tf:"object_ownership,omitempty"`
}

type BucketOwnershipControlsRuleParameters struct {

	// Object ownership. Valid values: BucketOwnerPreferred, ObjectWriter or BucketOwnerEnforced
	// +kubebuilder:validation:Required
	ObjectOwnership *string `json:"objectOwnership" tf:"object_ownership,omitempty"`
}

// BucketOwnershipControlsSpec defines the desired state of BucketOwnershipControls
type BucketOwnershipControlsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketOwnershipControlsParameters `json:"forProvider"`
}

// BucketOwnershipControlsStatus defines the observed state of BucketOwnershipControls.
type BucketOwnershipControlsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketOwnershipControlsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketOwnershipControls is the Schema for the BucketOwnershipControlss API. Manages S3 Bucket Ownership Controls.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketOwnershipControls struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule)",message="rule is a required parameter"
	Spec   BucketOwnershipControlsSpec   `json:"spec"`
	Status BucketOwnershipControlsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketOwnershipControlsList contains a list of BucketOwnershipControlss
type BucketOwnershipControlsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketOwnershipControls `json:"items"`
}

// Repository type metadata.
var (
	BucketOwnershipControls_Kind             = "BucketOwnershipControls"
	BucketOwnershipControls_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketOwnershipControls_Kind}.String()
	BucketOwnershipControls_KindAPIVersion   = BucketOwnershipControls_Kind + "." + CRDGroupVersion.String()
	BucketOwnershipControls_GroupVersionKind = CRDGroupVersion.WithKind(BucketOwnershipControls_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketOwnershipControls{}, &BucketOwnershipControlsList{})
}
