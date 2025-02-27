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

type BGPPeerInitParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/directconnect/v1beta1.PrivateVirtualInterface
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`

	// Reference to a PrivateVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDRef *v1.Reference `json:"virtualInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a PrivateVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDSelector *v1.Selector `json:"virtualInterfaceIdSelector,omitempty" tf:"-"`
}

type BGPPeerObservation struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The Direct Connect endpoint on which the BGP peer terminates.
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the BGP peer.
	BGPPeerID *string `json:"bgpPeerId,omitempty" tf:"bgp_peer_id,omitempty"`

	// The Up/Down state of the BGP peer.
	BGPStatus *string `json:"bgpStatus,omitempty" tf:"bgp_status,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The ID of the BGP peer resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`
}

type BGPPeerParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	// +kubebuilder:validation:Optional
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	// +kubebuilder:validation:Optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	// +kubebuilder:validation:Optional
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	// +kubebuilder:validation:Optional
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/directconnect/v1beta1.PrivateVirtualInterface
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`

	// Reference to a PrivateVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDRef *v1.Reference `json:"virtualInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a PrivateVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDSelector *v1.Selector `json:"virtualInterfaceIdSelector,omitempty" tf:"-"`
}

// BGPPeerSpec defines the desired state of BGPPeer
type BGPPeerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BGPPeerParameters `json:"forProvider"`
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
	InitProvider BGPPeerInitParameters `json:"initProvider,omitempty"`
}

// BGPPeerStatus defines the observed state of BGPPeer.
type BGPPeerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BGPPeerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BGPPeer is the Schema for the BGPPeers API. Provides a Direct Connect BGP peer resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type BGPPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addressFamily) || (has(self.initProvider) && has(self.initProvider.addressFamily))",message="spec.forProvider.addressFamily is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bgpAsn) || (has(self.initProvider) && has(self.initProvider.bgpAsn))",message="spec.forProvider.bgpAsn is a required parameter"
	Spec   BGPPeerSpec   `json:"spec"`
	Status BGPPeerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BGPPeerList contains a list of BGPPeers
type BGPPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BGPPeer `json:"items"`
}

// Repository type metadata.
var (
	BGPPeer_Kind             = "BGPPeer"
	BGPPeer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BGPPeer_Kind}.String()
	BGPPeer_KindAPIVersion   = BGPPeer_Kind + "." + CRDGroupVersion.String()
	BGPPeer_GroupVersionKind = CRDGroupVersion.WithKind(BGPPeer_Kind)
)

func init() {
	SchemeBuilder.Register(&BGPPeer{}, &BGPPeerList{})
}
