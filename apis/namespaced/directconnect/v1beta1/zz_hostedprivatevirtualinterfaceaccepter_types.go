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

type HostedPrivateVirtualInterfaceAccepterInitParameters struct {

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the virtual private gateway to which to connect the virtual interface.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// Reference to a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDRef *v1.Reference `json:"vpnGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDSelector *v1.Selector `json:"vpnGatewayIdSelector,omitempty" tf:"-"`

	// The ID of the Direct Connect virtual interface to accept.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.HostedPrivateVirtualInterface
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`

	// Reference to a HostedPrivateVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDRef *v1.Reference `json:"virtualInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a HostedPrivateVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDSelector *v1.Selector `json:"virtualInterfaceIdSelector,omitempty" tf:"-"`
}

type HostedPrivateVirtualInterfaceAccepterObservation struct {

	// The ARN of the virtual interface.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// The ID of the virtual interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the virtual private gateway to which to connect the virtual interface.
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`
}

type HostedPrivateVirtualInterfaceAccepterParameters struct {

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	// +kubebuilder:validation:Optional
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the virtual private gateway to which to connect the virtual interface.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// Reference to a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDRef *v1.Reference `json:"vpnGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDSelector *v1.Selector `json:"vpnGatewayIdSelector,omitempty" tf:"-"`

	// The ID of the Direct Connect virtual interface to accept.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.HostedPrivateVirtualInterface
	// +kubebuilder:validation:Optional
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`

	// Reference to a HostedPrivateVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDRef *v1.Reference `json:"virtualInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a HostedPrivateVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDSelector *v1.Selector `json:"virtualInterfaceIdSelector,omitempty" tf:"-"`
}

// HostedPrivateVirtualInterfaceAccepterSpec defines the desired state of HostedPrivateVirtualInterfaceAccepter
type HostedPrivateVirtualInterfaceAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedPrivateVirtualInterfaceAccepterParameters `json:"forProvider"`
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
	InitProvider HostedPrivateVirtualInterfaceAccepterInitParameters `json:"initProvider,omitempty"`
}

// HostedPrivateVirtualInterfaceAccepterStatus defines the observed state of HostedPrivateVirtualInterfaceAccepter.
type HostedPrivateVirtualInterfaceAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedPrivateVirtualInterfaceAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HostedPrivateVirtualInterfaceAccepter is the Schema for the HostedPrivateVirtualInterfaceAccepters API. Provides a resource to manage the accepter's side of a Direct Connect hosted private virtual interface.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type HostedPrivateVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedPrivateVirtualInterfaceAccepterSpec   `json:"spec"`
	Status            HostedPrivateVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPrivateVirtualInterfaceAccepterList contains a list of HostedPrivateVirtualInterfaceAccepters
type HostedPrivateVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedPrivateVirtualInterfaceAccepter `json:"items"`
}

// Repository type metadata.
var (
	HostedPrivateVirtualInterfaceAccepter_Kind             = "HostedPrivateVirtualInterfaceAccepter"
	HostedPrivateVirtualInterfaceAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedPrivateVirtualInterfaceAccepter_Kind}.String()
	HostedPrivateVirtualInterfaceAccepter_KindAPIVersion   = HostedPrivateVirtualInterfaceAccepter_Kind + "." + CRDGroupVersion.String()
	HostedPrivateVirtualInterfaceAccepter_GroupVersionKind = CRDGroupVersion.WithKind(HostedPrivateVirtualInterfaceAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedPrivateVirtualInterfaceAccepter{}, &HostedPrivateVirtualInterfaceAccepterList{})
}
