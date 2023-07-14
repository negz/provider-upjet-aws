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

type ByteMatchSetObservation struct {

	// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
	ByteMatchTuples []ByteMatchTuplesObservation `json:"byteMatchTuples,omitempty" tf:"byte_match_tuples,omitempty"`

	// The ID of the WAF ByteMatchSet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description of the ByteMatchSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ByteMatchSetParameters struct {

	// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
	// +kubebuilder:validation:Optional
	ByteMatchTuples []ByteMatchTuplesParameters `json:"byteMatchTuples,omitempty" tf:"byte_match_tuples,omitempty"`

	// The name or description of the ByteMatchSet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type ByteMatchTuplesObservation struct {

	// Settings for the ByteMatchTuple. FieldToMatch documented below.
	FieldToMatch []FieldToMatchObservation `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Within the portion of a web request that you want to search.
	PositionalConstraint *string `json:"positionalConstraint,omitempty" tf:"positional_constraint,omitempty"`

	// The value that you want AWS WAF to search for. The maximum length of the value is 50 bytes.
	TargetString *string `json:"targetString,omitempty" tf:"target_string,omitempty"`

	// The formatting way for web request.
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

type ByteMatchTuplesParameters struct {

	// Settings for the ByteMatchTuple. FieldToMatch documented below.
	// +kubebuilder:validation:Required
	FieldToMatch []FieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// Within the portion of a web request that you want to search.
	// +kubebuilder:validation:Required
	PositionalConstraint *string `json:"positionalConstraint" tf:"positional_constraint,omitempty"`

	// The value that you want AWS WAF to search for. The maximum length of the value is 50 bytes.
	// +kubebuilder:validation:Optional
	TargetString *string `json:"targetString,omitempty" tf:"target_string,omitempty"`

	// The formatting way for web request.
	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

type FieldToMatchObservation struct {

	// When the value of Type is HEADER, enter the name of the header that you want AWS WAF to search, for example, User-Agent or Referer. If the value of Type is any other value, omit Data.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FieldToMatchParameters struct {

	// When the value of Type is HEADER, enter the name of the header that you want AWS WAF to search, for example, User-Agent or Referer. If the value of Type is any other value, omit Data.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// ByteMatchSetSpec defines the desired state of ByteMatchSet
type ByteMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ByteMatchSetParameters `json:"forProvider"`
}

// ByteMatchSetStatus defines the observed state of ByteMatchSet.
type ByteMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ByteMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ByteMatchSet is the Schema for the ByteMatchSets API. Provides a AWS WAF Regional ByteMatchSet resource for use with ALB.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ByteMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   ByteMatchSetSpec   `json:"spec"`
	Status ByteMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ByteMatchSetList contains a list of ByteMatchSets
type ByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ByteMatchSet `json:"items"`
}

// Repository type metadata.
var (
	ByteMatchSet_Kind             = "ByteMatchSet"
	ByteMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ByteMatchSet_Kind}.String()
	ByteMatchSet_KindAPIVersion   = ByteMatchSet_Kind + "." + CRDGroupVersion.String()
	ByteMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(ByteMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ByteMatchSet{}, &ByteMatchSetList{})
}
