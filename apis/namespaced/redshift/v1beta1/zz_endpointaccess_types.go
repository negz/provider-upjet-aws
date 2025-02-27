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

type EndpointAccessInitParameters struct {

	// The cluster identifier of the cluster to access.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta2.Cluster
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// The Amazon Web Services account ID of the owner of the cluster. This is only required if the cluster is in another Amazon Web Services account.
	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	// The subnet group from which Amazon Redshift chooses the subnet to deploy the endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.SubnetGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// Reference to a SubnetGroup in redshift to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameRef *v1.Reference `json:"subnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup in redshift to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameSelector *v1.Selector `json:"subnetGroupNameSelector,omitempty" tf:"-"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +listType=set
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type EndpointAccessObservation struct {

	// The DNS address of the endpoint.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The cluster identifier of the cluster to access.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// The Redshift-managed VPC endpoint name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The port number on which the cluster accepts incoming connections.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The Amazon Web Services account ID of the owner of the cluster. This is only required if the cluster is in another Amazon Web Services account.
	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	// The subnet group from which Amazon Redshift chooses the subnet to deploy the endpoint.
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// The connection endpoint for connecting to an Amazon Redshift cluster through the proxy. See details below.
	VPCEndpoint []VPCEndpointObservation `json:"vpcEndpoint,omitempty" tf:"vpc_endpoint,omitempty"`

	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	// +listType=set
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type EndpointAccessParameters struct {

	// The cluster identifier of the cluster to access.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta2.Cluster
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Web Services account ID of the owner of the cluster. This is only required if the cluster is in another Amazon Web Services account.
	// +kubebuilder:validation:Optional
	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	// The subnet group from which Amazon Redshift chooses the subnet to deploy the endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.SubnetGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// Reference to a SubnetGroup in redshift to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameRef *v1.Reference `json:"subnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup in redshift to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameSelector *v1.Selector `json:"subnetGroupNameSelector,omitempty" tf:"-"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type NetworkInterfaceInitParameters struct {
}

type NetworkInterfaceObservation struct {

	// The Availability Zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The network interface identifier.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The IPv4 address of the network interface within the subnet.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The subnet identifier.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NetworkInterfaceParameters struct {
}

type VPCEndpointInitParameters struct {
}

type VPCEndpointObservation struct {

	// One or more network interfaces of the endpoint. Also known as an interface endpoint. See details below.
	NetworkInterface []NetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The VPC identifier that the endpoint is associated.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCEndpointParameters struct {
}

// EndpointAccessSpec defines the desired state of EndpointAccess
type EndpointAccessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointAccessParameters `json:"forProvider"`
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
	InitProvider EndpointAccessInitParameters `json:"initProvider,omitempty"`
}

// EndpointAccessStatus defines the observed state of EndpointAccess.
type EndpointAccessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointAccessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EndpointAccess is the Schema for the EndpointAccesss API. Provides a Redshift Endpoint Access resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type EndpointAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointAccessSpec   `json:"spec"`
	Status            EndpointAccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointAccessList contains a list of EndpointAccesss
type EndpointAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointAccess `json:"items"`
}

// Repository type metadata.
var (
	EndpointAccess_Kind             = "EndpointAccess"
	EndpointAccess_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointAccess_Kind}.String()
	EndpointAccess_KindAPIVersion   = EndpointAccess_Kind + "." + CRDGroupVersion.String()
	EndpointAccess_GroupVersionKind = CRDGroupVersion.WithKind(EndpointAccess_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointAccess{}, &EndpointAccessList{})
}
