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

type CriterionInitParameters struct {

	// List of string values to be evaluated.
	Equals []*string `json:"equals,omitempty" tf:"equals,omitempty"`

	// The name of the field to be evaluated. The full list of field names can be found in AWS documentation.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	GreaterThan *string `json:"greaterThan,omitempty" tf:"greater_than,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	GreaterThanOrEqual *string `json:"greaterThanOrEqual,omitempty" tf:"greater_than_or_equal,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	LessThan *string `json:"lessThan,omitempty" tf:"less_than,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	LessThanOrEqual *string `json:"lessThanOrEqual,omitempty" tf:"less_than_or_equal,omitempty"`

	// List of string values to be evaluated.
	NotEquals []*string `json:"notEquals,omitempty" tf:"not_equals,omitempty"`
}

type CriterionObservation struct {

	// List of string values to be evaluated.
	Equals []*string `json:"equals,omitempty" tf:"equals,omitempty"`

	// The name of the field to be evaluated. The full list of field names can be found in AWS documentation.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	GreaterThan *string `json:"greaterThan,omitempty" tf:"greater_than,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	GreaterThanOrEqual *string `json:"greaterThanOrEqual,omitempty" tf:"greater_than_or_equal,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	LessThan *string `json:"lessThan,omitempty" tf:"less_than,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	LessThanOrEqual *string `json:"lessThanOrEqual,omitempty" tf:"less_than_or_equal,omitempty"`

	// List of string values to be evaluated.
	NotEquals []*string `json:"notEquals,omitempty" tf:"not_equals,omitempty"`
}

type CriterionParameters struct {

	// List of string values to be evaluated.
	// +kubebuilder:validation:Optional
	Equals []*string `json:"equals,omitempty" tf:"equals,omitempty"`

	// The name of the field to be evaluated. The full list of field names can be found in AWS documentation.
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	// +kubebuilder:validation:Optional
	GreaterThan *string `json:"greaterThan,omitempty" tf:"greater_than,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	// +kubebuilder:validation:Optional
	GreaterThanOrEqual *string `json:"greaterThanOrEqual,omitempty" tf:"greater_than_or_equal,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	// +kubebuilder:validation:Optional
	LessThan *string `json:"lessThan,omitempty" tf:"less_than,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in RFC 3339 format.
	// +kubebuilder:validation:Optional
	LessThanOrEqual *string `json:"lessThanOrEqual,omitempty" tf:"less_than_or_equal,omitempty"`

	// List of string values to be evaluated.
	// +kubebuilder:validation:Optional
	NotEquals []*string `json:"notEquals,omitempty" tf:"not_equals,omitempty"`
}

type FilterInitParameters struct {

	// Specifies the action that is to be applied to the findings that match the filter. Can be one of ARCHIVE or NOOP.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Description of the filter.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Represents the criteria to be used in the filter for querying findings. Contains one or more criterion blocks, documented below.
	FindingCriteria *FindingCriteriaInitParameters `json:"findingCriteria,omitempty" tf:"finding_criteria,omitempty"`

	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	Rank *float64 `json:"rank,omitempty" tf:"rank,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FilterObservation struct {

	// Specifies the action that is to be applied to the findings that match the filter. Can be one of ARCHIVE or NOOP.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ARN of the GuardDuty filter.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the filter.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of a GuardDuty detector, attached to your account.
	DetectorID *string `json:"detectorId,omitempty" tf:"detector_id,omitempty"`

	// Represents the criteria to be used in the filter for querying findings. Contains one or more criterion blocks, documented below.
	FindingCriteria *FindingCriteriaObservation `json:"findingCriteria,omitempty" tf:"finding_criteria,omitempty"`

	// A compound field, consisting of the ID of the GuardDuty detector and the name of the filter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	Rank *float64 `json:"rank,omitempty" tf:"rank,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FilterParameters struct {

	// Specifies the action that is to be applied to the findings that match the filter. Can be one of ARCHIVE or NOOP.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Description of the filter.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of a GuardDuty detector, attached to your account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/guardduty/v1beta2.Detector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DetectorID *string `json:"detectorId,omitempty" tf:"detector_id,omitempty"`

	// Reference to a Detector in guardduty to populate detectorId.
	// +kubebuilder:validation:Optional
	DetectorIDRef *v1.Reference `json:"detectorIdRef,omitempty" tf:"-"`

	// Selector for a Detector in guardduty to populate detectorId.
	// +kubebuilder:validation:Optional
	DetectorIDSelector *v1.Selector `json:"detectorIdSelector,omitempty" tf:"-"`

	// Represents the criteria to be used in the filter for querying findings. Contains one or more criterion blocks, documented below.
	// +kubebuilder:validation:Optional
	FindingCriteria *FindingCriteriaParameters `json:"findingCriteria,omitempty" tf:"finding_criteria,omitempty"`

	// Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.
	// +kubebuilder:validation:Optional
	Rank *float64 `json:"rank,omitempty" tf:"rank,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FindingCriteriaInitParameters struct {
	Criterion []CriterionInitParameters `json:"criterion,omitempty" tf:"criterion,omitempty"`
}

type FindingCriteriaObservation struct {
	Criterion []CriterionObservation `json:"criterion,omitempty" tf:"criterion,omitempty"`
}

type FindingCriteriaParameters struct {

	// +kubebuilder:validation:Optional
	Criterion []CriterionParameters `json:"criterion" tf:"criterion,omitempty"`
}

// FilterSpec defines the desired state of Filter
type FilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FilterParameters `json:"forProvider"`
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
	InitProvider FilterInitParameters `json:"initProvider,omitempty"`
}

// FilterStatus defines the observed state of Filter.
type FilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Filter is the Schema for the Filters API. Provides a resource to manage a GuardDuty filter
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Filter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.findingCriteria) || (has(self.initProvider) && has(self.initProvider.findingCriteria))",message="spec.forProvider.findingCriteria is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rank) || (has(self.initProvider) && has(self.initProvider.rank))",message="spec.forProvider.rank is a required parameter"
	Spec   FilterSpec   `json:"spec"`
	Status FilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FilterList contains a list of Filters
type FilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Filter `json:"items"`
}

// Repository type metadata.
var (
	Filter_Kind             = "Filter"
	Filter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Filter_Kind}.String()
	Filter_KindAPIVersion   = Filter_Kind + "." + CRDGroupVersion.String()
	Filter_GroupVersionKind = CRDGroupVersion.WithKind(Filter_Kind)
)

func init() {
	SchemeBuilder.Register(&Filter{}, &FilterList{})
}
