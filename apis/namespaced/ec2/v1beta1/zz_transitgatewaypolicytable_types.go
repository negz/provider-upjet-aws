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

type TransitGatewayPolicyTableInitParameters struct {

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// EC2 Transit Gateway identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.TransitGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Reference to a TransitGateway in ec2 to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway in ec2 to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`
}

type TransitGatewayPolicyTableObservation struct {

	// EC2 Transit Gateway Policy Table Amazon Resource Name (ARN).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// EC2 Transit Gateway Policy Table identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The state of the EC2 Transit Gateway Policy Table.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// EC2 Transit Gateway identifier.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`
}

type TransitGatewayPolicyTableParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// EC2 Transit Gateway identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.TransitGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Reference to a TransitGateway in ec2 to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway in ec2 to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`
}

// TransitGatewayPolicyTableSpec defines the desired state of TransitGatewayPolicyTable
type TransitGatewayPolicyTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayPolicyTableParameters `json:"forProvider"`
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
	InitProvider TransitGatewayPolicyTableInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayPolicyTableStatus defines the observed state of TransitGatewayPolicyTable.
type TransitGatewayPolicyTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayPolicyTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TransitGatewayPolicyTable is the Schema for the TransitGatewayPolicyTables API. Manages an EC2 Transit Gateway Policy Table
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type TransitGatewayPolicyTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayPolicyTableSpec   `json:"spec"`
	Status            TransitGatewayPolicyTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPolicyTableList contains a list of TransitGatewayPolicyTables
type TransitGatewayPolicyTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayPolicyTable `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayPolicyTable_Kind             = "TransitGatewayPolicyTable"
	TransitGatewayPolicyTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayPolicyTable_Kind}.String()
	TransitGatewayPolicyTable_KindAPIVersion   = TransitGatewayPolicyTable_Kind + "." + CRDGroupVersion.String()
	TransitGatewayPolicyTable_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayPolicyTable_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayPolicyTable{}, &TransitGatewayPolicyTableList{})
}
