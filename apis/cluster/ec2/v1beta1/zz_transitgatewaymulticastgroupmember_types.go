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

type TransitGatewayMulticastGroupMemberInitParameters struct {

	// The IP address assigned to the transit gateway multicast group.
	GroupIPAddress *string `json:"groupIpAddress,omitempty" tf:"group_ip_address,omitempty"`

	// The group members' network interface ID to register with the transit gateway multicast group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface in ec2 to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface in ec2 to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// The ID of the transit gateway multicast domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.TransitGatewayMulticastDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TransitGatewayMulticastDomainID *string `json:"transitGatewayMulticastDomainId,omitempty" tf:"transit_gateway_multicast_domain_id,omitempty"`

	// Reference to a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDRef *v1.Reference `json:"transitGatewayMulticastDomainIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDSelector *v1.Selector `json:"transitGatewayMulticastDomainIdSelector,omitempty" tf:"-"`
}

type TransitGatewayMulticastGroupMemberObservation struct {

	// The IP address assigned to the transit gateway multicast group.
	GroupIPAddress *string `json:"groupIpAddress,omitempty" tf:"group_ip_address,omitempty"`

	// EC2 Transit Gateway Multicast Group Member identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The group members' network interface ID to register with the transit gateway multicast group.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainID *string `json:"transitGatewayMulticastDomainId,omitempty" tf:"transit_gateway_multicast_domain_id,omitempty"`
}

type TransitGatewayMulticastGroupMemberParameters struct {

	// The IP address assigned to the transit gateway multicast group.
	// +kubebuilder:validation:Optional
	GroupIPAddress *string `json:"groupIpAddress,omitempty" tf:"group_ip_address,omitempty"`

	// The group members' network interface ID to register with the transit gateway multicast group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface in ec2 to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface in ec2 to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the transit gateway multicast domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.TransitGatewayMulticastDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainID *string `json:"transitGatewayMulticastDomainId,omitempty" tf:"transit_gateway_multicast_domain_id,omitempty"`

	// Reference to a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDRef *v1.Reference `json:"transitGatewayMulticastDomainIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDSelector *v1.Selector `json:"transitGatewayMulticastDomainIdSelector,omitempty" tf:"-"`
}

// TransitGatewayMulticastGroupMemberSpec defines the desired state of TransitGatewayMulticastGroupMember
type TransitGatewayMulticastGroupMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayMulticastGroupMemberParameters `json:"forProvider"`
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
	InitProvider TransitGatewayMulticastGroupMemberInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayMulticastGroupMemberStatus defines the observed state of TransitGatewayMulticastGroupMember.
type TransitGatewayMulticastGroupMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayMulticastGroupMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TransitGatewayMulticastGroupMember is the Schema for the TransitGatewayMulticastGroupMembers API. Manages an EC2 Transit Gateway Multicast Group Member
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayMulticastGroupMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupIpAddress) || (has(self.initProvider) && has(self.initProvider.groupIpAddress))",message="spec.forProvider.groupIpAddress is a required parameter"
	Spec   TransitGatewayMulticastGroupMemberSpec   `json:"spec"`
	Status TransitGatewayMulticastGroupMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayMulticastGroupMemberList contains a list of TransitGatewayMulticastGroupMembers
type TransitGatewayMulticastGroupMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayMulticastGroupMember `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayMulticastGroupMember_Kind             = "TransitGatewayMulticastGroupMember"
	TransitGatewayMulticastGroupMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayMulticastGroupMember_Kind}.String()
	TransitGatewayMulticastGroupMember_KindAPIVersion   = TransitGatewayMulticastGroupMember_Kind + "." + CRDGroupVersion.String()
	TransitGatewayMulticastGroupMember_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayMulticastGroupMember_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayMulticastGroupMember{}, &TransitGatewayMulticastGroupMemberList{})
}
