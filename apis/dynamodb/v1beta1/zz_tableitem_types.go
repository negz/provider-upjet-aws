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

type TableItemInitParameters struct {

	// Hash key to use for lookups and identification of the item
	HashKey *string `json:"hashKey,omitempty" tf:"hash_key,omitempty"`

	// JSON representation of a map of attribute name/value pairs, one for each attribute. Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item *string `json:"item,omitempty" tf:"item,omitempty"`

	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

	// Name of the table to contain the item.
	// +crossplane:generate:reference:type=Table
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`
}

type TableItemObservation struct {

	// Hash key to use for lookups and identification of the item
	HashKey *string `json:"hashKey,omitempty" tf:"hash_key,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// JSON representation of a map of attribute name/value pairs, one for each attribute. Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item *string `json:"item,omitempty" tf:"item,omitempty"`

	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

	// Name of the table to contain the item.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type TableItemParameters struct {

	// Hash key to use for lookups and identification of the item
	// +kubebuilder:validation:Optional
	HashKey *string `json:"hashKey,omitempty" tf:"hash_key,omitempty"`

	// JSON representation of a map of attribute name/value pairs, one for each attribute. Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	// +kubebuilder:validation:Optional
	Item *string `json:"item,omitempty" tf:"item,omitempty"`

	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	// +kubebuilder:validation:Optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the table to contain the item.
	// +crossplane:generate:reference:type=Table
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`
}

// TableItemSpec defines the desired state of TableItem
type TableItemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableItemParameters `json:"forProvider"`
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
	InitProvider TableItemInitParameters `json:"initProvider,omitempty"`
}

// TableItemStatus defines the observed state of TableItem.
type TableItemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableItemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TableItem is the Schema for the TableItems API. Provides a DynamoDB table item resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TableItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hashKey) || (has(self.initProvider) && has(self.initProvider.hashKey))",message="spec.forProvider.hashKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.item) || (has(self.initProvider) && has(self.initProvider.item))",message="spec.forProvider.item is a required parameter"
	Spec   TableItemSpec   `json:"spec"`
	Status TableItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableItemList contains a list of TableItems
type TableItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableItem `json:"items"`
}

// Repository type metadata.
var (
	TableItem_Kind             = "TableItem"
	TableItem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableItem_Kind}.String()
	TableItem_KindAPIVersion   = TableItem_Kind + "." + CRDGroupVersion.String()
	TableItem_GroupVersionKind = CRDGroupVersion.WithKind(TableItem_Kind)
)

func init() {
	SchemeBuilder.Register(&TableItem{}, &TableItemList{})
}
