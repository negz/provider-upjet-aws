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

type NetworkACLEgressInitParameters struct {
}

type NetworkACLEgressObservation struct {

	// The action to take.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The CIDR block to match. This must be a
	// valid network mask.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The from port to match.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IPv6 CIDR block.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The ICMP type code to be used. Default 0.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to be used. Default 0.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The protocol to match. If using the -1 'all'
	// protocol, you must specify a from and to port of 0.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The rule number. Used for ordering.
	RuleNo *float64 `json:"ruleNo,omitempty" tf:"rule_no,omitempty"`

	// The to port to match.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type NetworkACLEgressParameters struct {
}

type NetworkACLIngressInitParameters struct {
}

type NetworkACLIngressObservation struct {

	// The action to take.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The CIDR block to match. This must be a
	// valid network mask.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The from port to match.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IPv6 CIDR block.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The ICMP type code to be used. Default 0.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to be used. Default 0.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The protocol to match. If using the -1 'all'
	// protocol, you must specify a from and to port of 0.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The rule number. Used for ordering.
	RuleNo *float64 `json:"ruleNo,omitempty" tf:"rule_no,omitempty"`

	// The to port to match.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type NetworkACLIngressParameters struct {
}

type NetworkACLInitParameters struct {

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of Subnet IDs to apply the ACL to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the associated VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type NetworkACLObservation struct {

	// The ARN of the network ACL
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies an egress rule. Parameters defined below.
	// This argument is processed in attribute-as-blocks mode.
	Egress []NetworkACLEgressObservation `json:"egress,omitempty" tf:"egress,omitempty"`

	// The ID of the network ACL
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies an ingress rule. Parameters defined below.
	// This argument is processed in attribute-as-blocks mode.
	Ingress []NetworkACLIngressObservation `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// The ID of the AWS account that owns the network ACL.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A list of Subnet IDs to apply the ACL to
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the associated VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type NetworkACLParameters struct {

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

	// A list of Subnet IDs to apply the ACL to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the associated VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// NetworkACLSpec defines the desired state of NetworkACL
type NetworkACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkACLParameters `json:"forProvider"`
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
	InitProvider NetworkACLInitParameters `json:"initProvider,omitempty"`
}

// NetworkACLStatus defines the observed state of NetworkACL.
type NetworkACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NetworkACL is the Schema for the NetworkACLs API. Provides an network ACL resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type NetworkACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkACLSpec   `json:"spec"`
	Status            NetworkACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkACLList contains a list of NetworkACLs
type NetworkACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkACL `json:"items"`
}

// Repository type metadata.
var (
	NetworkACL_Kind             = "NetworkACL"
	NetworkACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkACL_Kind}.String()
	NetworkACL_KindAPIVersion   = NetworkACL_Kind + "." + CRDGroupVersion.String()
	NetworkACL_GroupVersionKind = CRDGroupVersion.WithKind(NetworkACL_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkACL{}, &NetworkACLList{})
}
