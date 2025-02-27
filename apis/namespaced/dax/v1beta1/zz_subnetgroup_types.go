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

type SubnetGroupInitParameters struct {

	// A description of the subnet group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// –  A list of VPC subnet IDs for the subnet group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type SubnetGroupObservation struct {

	// A description of the subnet group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the subnet group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  A list of VPC subnet IDs for the subnet group.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// – VPC ID of the subnet group.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type SubnetGroupParameters struct {

	// A description of the subnet group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// –  A list of VPC subnet IDs for the subnet group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

// SubnetGroupSpec defines the desired state of SubnetGroup
type SubnetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetGroupParameters `json:"forProvider"`
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
	InitProvider SubnetGroupInitParameters `json:"initProvider,omitempty"`
}

// SubnetGroupStatus defines the observed state of SubnetGroup.
type SubnetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubnetGroup is the Schema for the SubnetGroups API. Provides an DAX Subnet Group resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type SubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetGroupSpec   `json:"spec"`
	Status            SubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetGroupList contains a list of SubnetGroups
type SubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetGroup `json:"items"`
}

// Repository type metadata.
var (
	SubnetGroup_Kind             = "SubnetGroup"
	SubnetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetGroup_Kind}.String()
	SubnetGroup_KindAPIVersion   = SubnetGroup_Kind + "." + CRDGroupVersion.String()
	SubnetGroup_GroupVersionKind = CRDGroupVersion.WithKind(SubnetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetGroup{}, &SubnetGroupList{})
}
