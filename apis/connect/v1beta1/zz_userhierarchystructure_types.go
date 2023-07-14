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

type HierarchyStructureObservation struct {

	// A block that defines the details of level five. The level block is documented below.
	LevelFive []LevelFiveObservation `json:"levelFive,omitempty" tf:"level_five,omitempty"`

	// A block that defines the details of level four. The level block is documented below.
	LevelFour []LevelFourObservation `json:"levelFour,omitempty" tf:"level_four,omitempty"`

	// A block that defines the details of level one. The level block is documented below.
	LevelOne []LevelOneObservation `json:"levelOne,omitempty" tf:"level_one,omitempty"`

	// A block that defines the details of level three. The level block is documented below.
	LevelThree []LevelThreeObservation `json:"levelThree,omitempty" tf:"level_three,omitempty"`

	// A block that defines the details of level two. The level block is documented below.
	LevelTwo []LevelTwoObservation `json:"levelTwo,omitempty" tf:"level_two,omitempty"`
}

type HierarchyStructureParameters struct {

	// A block that defines the details of level five. The level block is documented below.
	// +kubebuilder:validation:Optional
	LevelFive []LevelFiveParameters `json:"levelFive,omitempty" tf:"level_five,omitempty"`

	// A block that defines the details of level four. The level block is documented below.
	// +kubebuilder:validation:Optional
	LevelFour []LevelFourParameters `json:"levelFour,omitempty" tf:"level_four,omitempty"`

	// A block that defines the details of level one. The level block is documented below.
	// +kubebuilder:validation:Optional
	LevelOne []LevelOneParameters `json:"levelOne,omitempty" tf:"level_one,omitempty"`

	// A block that defines the details of level three. The level block is documented below.
	// +kubebuilder:validation:Optional
	LevelThree []LevelThreeParameters `json:"levelThree,omitempty" tf:"level_three,omitempty"`

	// A block that defines the details of level two. The level block is documented below.
	// +kubebuilder:validation:Optional
	LevelTwo []LevelTwoParameters `json:"levelTwo,omitempty" tf:"level_two,omitempty"`
}

type LevelFiveObservation struct {

	// The Amazon Resource Name (ARN) of the hierarchy level.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the hosting Amazon Connect Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the user hierarchy level. Must not be more than 50 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LevelFiveParameters struct {

	// The name of the user hierarchy level. Must not be more than 50 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LevelFourObservation struct {

	// The Amazon Resource Name (ARN) of the hierarchy level.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the hosting Amazon Connect Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the user hierarchy level. Must not be more than 50 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LevelFourParameters struct {

	// The name of the user hierarchy level. Must not be more than 50 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LevelOneObservation struct {

	// The Amazon Resource Name (ARN) of the hierarchy level.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the hosting Amazon Connect Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the user hierarchy level. Must not be more than 50 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LevelOneParameters struct {

	// The name of the user hierarchy level. Must not be more than 50 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LevelThreeObservation struct {

	// The Amazon Resource Name (ARN) of the hierarchy level.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the hosting Amazon Connect Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the user hierarchy level. Must not be more than 50 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LevelThreeParameters struct {

	// The name of the user hierarchy level. Must not be more than 50 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LevelTwoObservation struct {

	// The Amazon Resource Name (ARN) of the hierarchy level.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the hosting Amazon Connect Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the user hierarchy level. Must not be more than 50 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LevelTwoParameters struct {

	// The name of the user hierarchy level. Must not be more than 50 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type UserHierarchyStructureObservation struct {

	// A block that defines the hierarchy structure's levels. The hierarchy_structure block is documented below.
	HierarchyStructure []HierarchyStructureObservation `json:"hierarchyStructure,omitempty" tf:"hierarchy_structure,omitempty"`

	// The identifier of the hosting Amazon Connect Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type UserHierarchyStructureParameters struct {

	// A block that defines the hierarchy structure's levels. The hierarchy_structure block is documented below.
	// +kubebuilder:validation:Optional
	HierarchyStructure []HierarchyStructureParameters `json:"hierarchyStructure,omitempty" tf:"hierarchy_structure,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// UserHierarchyStructureSpec defines the desired state of UserHierarchyStructure
type UserHierarchyStructureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserHierarchyStructureParameters `json:"forProvider"`
}

// UserHierarchyStructureStatus defines the observed state of UserHierarchyStructure.
type UserHierarchyStructureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserHierarchyStructureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserHierarchyStructure is the Schema for the UserHierarchyStructures API. Provides details about a specific Amazon Connect User Hierarchy Structure
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserHierarchyStructure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hierarchyStructure)",message="hierarchyStructure is a required parameter"
	Spec   UserHierarchyStructureSpec   `json:"spec"`
	Status UserHierarchyStructureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserHierarchyStructureList contains a list of UserHierarchyStructures
type UserHierarchyStructureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserHierarchyStructure `json:"items"`
}

// Repository type metadata.
var (
	UserHierarchyStructure_Kind             = "UserHierarchyStructure"
	UserHierarchyStructure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserHierarchyStructure_Kind}.String()
	UserHierarchyStructure_KindAPIVersion   = UserHierarchyStructure_Kind + "." + CRDGroupVersion.String()
	UserHierarchyStructure_GroupVersionKind = CRDGroupVersion.WithKind(UserHierarchyStructure_Kind)
)

func init() {
	SchemeBuilder.Register(&UserHierarchyStructure{}, &UserHierarchyStructureList{})
}
