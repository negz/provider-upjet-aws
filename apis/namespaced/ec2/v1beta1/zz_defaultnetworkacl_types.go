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

type DefaultNetworkACLInitParameters struct {

	// Network ACL ID to manage. This attribute is exported from aws_vpc, or manually found via the AWS Console.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("default_network_acl_id",true)
	DefaultNetworkACLID *string `json:"defaultNetworkAclId,omitempty" tf:"default_network_acl_id,omitempty"`

	// Reference to a VPC in ec2 to populate defaultNetworkAclId.
	// +kubebuilder:validation:Optional
	DefaultNetworkACLIDRef *v1.Reference `json:"defaultNetworkAclIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate defaultNetworkAclId.
	// +kubebuilder:validation:Optional
	DefaultNetworkACLIDSelector *v1.Selector `json:"defaultNetworkAclIdSelector,omitempty" tf:"-"`

	// Configuration block for an egress rule. Detailed below.
	Egress []EgressInitParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// Configuration block for an ingress rule. Detailed below.
	Ingress []IngressInitParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultNetworkACLObservation struct {

	// ARN of the Default Network ACL
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Network ACL ID to manage. This attribute is exported from aws_vpc, or manually found via the AWS Console.
	DefaultNetworkACLID *string `json:"defaultNetworkAclId,omitempty" tf:"default_network_acl_id,omitempty"`

	// Configuration block for an egress rule. Detailed below.
	Egress []EgressObservation `json:"egress,omitempty" tf:"egress,omitempty"`

	// ID of the Default Network ACL
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block for an ingress rule. Detailed below.
	Ingress []IngressObservation `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// ID of the AWS account that owns the Default Network ACL
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// ID of the associated VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultNetworkACLParameters struct {

	// Network ACL ID to manage. This attribute is exported from aws_vpc, or manually found via the AWS Console.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("default_network_acl_id",true)
	// +kubebuilder:validation:Optional
	DefaultNetworkACLID *string `json:"defaultNetworkAclId,omitempty" tf:"default_network_acl_id,omitempty"`

	// Reference to a VPC in ec2 to populate defaultNetworkAclId.
	// +kubebuilder:validation:Optional
	DefaultNetworkACLIDRef *v1.Reference `json:"defaultNetworkAclIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate defaultNetworkAclId.
	// +kubebuilder:validation:Optional
	DefaultNetworkACLIDSelector *v1.Selector `json:"defaultNetworkAclIdSelector,omitempty" tf:"-"`

	// Configuration block for an egress rule. Detailed below.
	// +kubebuilder:validation:Optional
	Egress []EgressParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// Configuration block for an ingress rule. Detailed below.
	// +kubebuilder:validation:Optional
	Ingress []IngressParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

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

	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
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
}

type EgressInitParameters struct {

	// The action to take.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The CIDR block to match. This must be a valid network mask.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The from port to match.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IPv6 CIDR block.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The ICMP type code to be used. Default 0.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to be used. Default 0.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The rule number. Used for ordering.
	RuleNo *float64 `json:"ruleNo,omitempty" tf:"rule_no,omitempty"`

	// The to port to match.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type EgressObservation struct {

	// The action to take.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The CIDR block to match. This must be a valid network mask.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The from port to match.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IPv6 CIDR block.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The ICMP type code to be used. Default 0.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to be used. Default 0.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The rule number. Used for ordering.
	RuleNo *float64 `json:"ruleNo,omitempty" tf:"rule_no,omitempty"`

	// The to port to match.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type EgressParameters struct {

	// The action to take.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The CIDR block to match. This must be a valid network mask.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The from port to match.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// The IPv6 CIDR block.
	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The ICMP type code to be used. Default 0.
	// +kubebuilder:validation:Optional
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to be used. Default 0.
	// +kubebuilder:validation:Optional
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The rule number. Used for ordering.
	// +kubebuilder:validation:Optional
	RuleNo *float64 `json:"ruleNo" tf:"rule_no,omitempty"`

	// The to port to match.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort" tf:"to_port,omitempty"`
}

type IngressInitParameters struct {

	// The action to take.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The CIDR block to match. This must be a valid network mask.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.DefaultVPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("cidr_block",true)
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// Reference to a DefaultVPC in ec2 to populate cidrBlock.
	// +kubebuilder:validation:Optional
	CidrBlockRef *v1.Reference `json:"cidrBlockRef,omitempty" tf:"-"`

	// Selector for a DefaultVPC in ec2 to populate cidrBlock.
	// +kubebuilder:validation:Optional
	CidrBlockSelector *v1.Selector `json:"cidrBlockSelector,omitempty" tf:"-"`

	// The from port to match.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IPv6 CIDR block.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The ICMP type code to be used. Default 0.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to be used. Default 0.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The rule number. Used for ordering.
	RuleNo *float64 `json:"ruleNo,omitempty" tf:"rule_no,omitempty"`

	// The to port to match.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type IngressObservation struct {

	// The action to take.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The CIDR block to match. This must be a valid network mask.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The from port to match.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IPv6 CIDR block.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The ICMP type code to be used. Default 0.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to be used. Default 0.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The rule number. Used for ordering.
	RuleNo *float64 `json:"ruleNo,omitempty" tf:"rule_no,omitempty"`

	// The to port to match.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type IngressParameters struct {

	// The action to take.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The CIDR block to match. This must be a valid network mask.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.DefaultVPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("cidr_block",true)
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// Reference to a DefaultVPC in ec2 to populate cidrBlock.
	// +kubebuilder:validation:Optional
	CidrBlockRef *v1.Reference `json:"cidrBlockRef,omitempty" tf:"-"`

	// Selector for a DefaultVPC in ec2 to populate cidrBlock.
	// +kubebuilder:validation:Optional
	CidrBlockSelector *v1.Selector `json:"cidrBlockSelector,omitempty" tf:"-"`

	// The from port to match.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// The IPv6 CIDR block.
	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The ICMP type code to be used. Default 0.
	// +kubebuilder:validation:Optional
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// The ICMP type to be used. Default 0.
	// +kubebuilder:validation:Optional
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The rule number. Used for ordering.
	// +kubebuilder:validation:Optional
	RuleNo *float64 `json:"ruleNo" tf:"rule_no,omitempty"`

	// The to port to match.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort" tf:"to_port,omitempty"`
}

// DefaultNetworkACLSpec defines the desired state of DefaultNetworkACL
type DefaultNetworkACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultNetworkACLParameters `json:"forProvider"`
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
	InitProvider DefaultNetworkACLInitParameters `json:"initProvider,omitempty"`
}

// DefaultNetworkACLStatus defines the observed state of DefaultNetworkACL.
type DefaultNetworkACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultNetworkACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DefaultNetworkACL is the Schema for the DefaultNetworkACLs API. Manage a default network ACL.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type DefaultNetworkACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultNetworkACLSpec   `json:"spec"`
	Status            DefaultNetworkACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultNetworkACLList contains a list of DefaultNetworkACLs
type DefaultNetworkACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultNetworkACL `json:"items"`
}

// Repository type metadata.
var (
	DefaultNetworkACL_Kind             = "DefaultNetworkACL"
	DefaultNetworkACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultNetworkACL_Kind}.String()
	DefaultNetworkACL_KindAPIVersion   = DefaultNetworkACL_Kind + "." + CRDGroupVersion.String()
	DefaultNetworkACL_GroupVersionKind = CRDGroupVersion.WithKind(DefaultNetworkACL_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultNetworkACL{}, &DefaultNetworkACLList{})
}
