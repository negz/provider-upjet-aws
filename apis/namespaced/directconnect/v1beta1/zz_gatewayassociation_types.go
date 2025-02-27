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

type GatewayAssociationInitParameters struct {

	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	// +listType=set
	AllowedPrefixes []*string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes,omitempty"`

	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	AssociatedGatewayID *string `json:"associatedGatewayId,omitempty" tf:"associated_gateway_id,omitempty"`

	// Reference to a VPNGateway in ec2 to populate associatedGatewayId.
	// +kubebuilder:validation:Optional
	AssociatedGatewayIDRef *v1.Reference `json:"associatedGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in ec2 to populate associatedGatewayId.
	// +kubebuilder:validation:Optional
	AssociatedGatewayIDSelector *v1.Selector `json:"associatedGatewayIdSelector,omitempty" tf:"-"`

	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountID *string `json:"associatedGatewayOwnerAccountId,omitempty" tf:"associated_gateway_owner_account_id,omitempty"`

	// The ID of the Direct Connect gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/directconnect/v1beta1.Gateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// Reference to a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDRef *v1.Reference `json:"dxGatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDSelector *v1.Selector `json:"dxGatewayIdSelector,omitempty" tf:"-"`

	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalID *string `json:"proposalId,omitempty" tf:"proposal_id,omitempty"`

	// The ID of the Direct Connect gateway association resource.
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type GatewayAssociationObservation struct {

	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	// +listType=set
	AllowedPrefixes []*string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes,omitempty"`

	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayID *string `json:"associatedGatewayId,omitempty" tf:"associated_gateway_id,omitempty"`

	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountID *string `json:"associatedGatewayOwnerAccountId,omitempty" tf:"associated_gateway_owner_account_id,omitempty"`

	// The type of the associated gateway, transitGateway or virtualPrivateGateway.
	AssociatedGatewayType *string `json:"associatedGatewayType,omitempty" tf:"associated_gateway_type,omitempty"`

	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationID *string `json:"dxGatewayAssociationId,omitempty" tf:"dx_gateway_association_id,omitempty"`

	// The ID of the Direct Connect gateway.
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountID *string `json:"dxGatewayOwnerAccountId,omitempty" tf:"dx_gateway_owner_account_id,omitempty"`

	// The ID of the Direct Connect gateway association resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalID *string `json:"proposalId,omitempty" tf:"proposal_id,omitempty"`

	// The ID of the Direct Connect gateway association resource.
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type GatewayAssociationParameters struct {

	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedPrefixes []*string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes,omitempty"`

	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AssociatedGatewayID *string `json:"associatedGatewayId,omitempty" tf:"associated_gateway_id,omitempty"`

	// Reference to a VPNGateway in ec2 to populate associatedGatewayId.
	// +kubebuilder:validation:Optional
	AssociatedGatewayIDRef *v1.Reference `json:"associatedGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in ec2 to populate associatedGatewayId.
	// +kubebuilder:validation:Optional
	AssociatedGatewayIDSelector *v1.Selector `json:"associatedGatewayIdSelector,omitempty" tf:"-"`

	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	// +kubebuilder:validation:Optional
	AssociatedGatewayOwnerAccountID *string `json:"associatedGatewayOwnerAccountId,omitempty" tf:"associated_gateway_owner_account_id,omitempty"`

	// The ID of the Direct Connect gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/directconnect/v1beta1.Gateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// Reference to a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDRef *v1.Reference `json:"dxGatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDSelector *v1.Selector `json:"dxGatewayIdSelector,omitempty" tf:"-"`

	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	// +kubebuilder:validation:Optional
	ProposalID *string `json:"proposalId,omitempty" tf:"proposal_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the Direct Connect gateway association resource.
	// +kubebuilder:validation:Optional
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`
}

// GatewayAssociationSpec defines the desired state of GatewayAssociation
type GatewayAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayAssociationParameters `json:"forProvider"`
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
	InitProvider GatewayAssociationInitParameters `json:"initProvider,omitempty"`
}

// GatewayAssociationStatus defines the observed state of GatewayAssociation.
type GatewayAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GatewayAssociation is the Schema for the GatewayAssociations API. Associates a Direct Connect Gateway with a VGW or transit gateway.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type GatewayAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayAssociationSpec   `json:"spec"`
	Status            GatewayAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayAssociationList contains a list of GatewayAssociations
type GatewayAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayAssociation `json:"items"`
}

// Repository type metadata.
var (
	GatewayAssociation_Kind             = "GatewayAssociation"
	GatewayAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayAssociation_Kind}.String()
	GatewayAssociation_KindAPIVersion   = GatewayAssociation_Kind + "." + CRDGroupVersion.String()
	GatewayAssociation_GroupVersionKind = CRDGroupVersion.WithKind(GatewayAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayAssociation{}, &GatewayAssociationList{})
}
