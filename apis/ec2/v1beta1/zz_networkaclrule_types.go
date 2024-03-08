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

type NetworkACLRuleInitParameters struct {

	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default false.
	Egress *bool `json:"egress,omitempty" tf:"egress,omitempty"`

	// The from port to match.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IPv6 CIDR block to allow or deny.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocolE.g., -1
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocolE.g., -1
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The ID of the network ACL.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.NetworkACL
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	NetworkACLID *string `json:"networkAclId,omitempty" tf:"network_acl_id,omitempty"`

	// Reference to a NetworkACL in ec2 to populate networkAclId.
	// +kubebuilder:validation:Optional
	NetworkACLIDRef *v1.Reference `json:"networkAclIdRef,omitempty" tf:"-"`

	// Selector for a NetworkACL in ec2 to populate networkAclId.
	// +kubebuilder:validation:Optional
	NetworkACLIDSelector *v1.Selector `json:"networkAclIdSelector,omitempty" tf:"-"`

	// The protocol. A value of -1 means all protocols.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: allow | deny
	RuleAction *string `json:"ruleAction,omitempty" tf:"rule_action,omitempty"`

	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber *float64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`

	// The to port to match.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type NetworkACLRuleObservation struct {

	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default false.
	Egress *bool `json:"egress,omitempty" tf:"egress,omitempty"`

	// The from port to match.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The ID of the network ACL Rule
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IPv6 CIDR block to allow or deny.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocolE.g., -1
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocolE.g., -1
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The ID of the network ACL.
	NetworkACLID *string `json:"networkAclId,omitempty" tf:"network_acl_id,omitempty"`

	// The protocol. A value of -1 means all protocols.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: allow | deny
	RuleAction *string `json:"ruleAction,omitempty" tf:"rule_action,omitempty"`

	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber *float64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`

	// The to port to match.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type NetworkACLRuleParameters struct {

	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default false.
	// +kubebuilder:validation:Optional
	Egress *bool `json:"egress,omitempty" tf:"egress,omitempty"`

	// The from port to match.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The IPv6 CIDR block to allow or deny.
	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocolE.g., -1
	// +kubebuilder:validation:Optional
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocolE.g., -1
	// +kubebuilder:validation:Optional
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// The ID of the network ACL.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.NetworkACL
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkACLID *string `json:"networkAclId,omitempty" tf:"network_acl_id,omitempty"`

	// Reference to a NetworkACL in ec2 to populate networkAclId.
	// +kubebuilder:validation:Optional
	NetworkACLIDRef *v1.Reference `json:"networkAclIdRef,omitempty" tf:"-"`

	// Selector for a NetworkACL in ec2 to populate networkAclId.
	// +kubebuilder:validation:Optional
	NetworkACLIDSelector *v1.Selector `json:"networkAclIdSelector,omitempty" tf:"-"`

	// The protocol. A value of -1 means all protocols.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: allow | deny
	// +kubebuilder:validation:Optional
	RuleAction *string `json:"ruleAction,omitempty" tf:"rule_action,omitempty"`

	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	// +kubebuilder:validation:Optional
	RuleNumber *float64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`

	// The to port to match.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

// NetworkACLRuleSpec defines the desired state of NetworkACLRule
type NetworkACLRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkACLRuleParameters `json:"forProvider"`
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
	InitProvider NetworkACLRuleInitParameters `json:"initProvider,omitempty"`
}

// NetworkACLRuleStatus defines the observed state of NetworkACLRule.
type NetworkACLRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkACLRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NetworkACLRule is the Schema for the NetworkACLRules API. Provides an network ACL Rule resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkACLRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleAction) || (has(self.initProvider) && has(self.initProvider.ruleAction))",message="spec.forProvider.ruleAction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleNumber) || (has(self.initProvider) && has(self.initProvider.ruleNumber))",message="spec.forProvider.ruleNumber is a required parameter"
	Spec   NetworkACLRuleSpec   `json:"spec"`
	Status NetworkACLRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkACLRuleList contains a list of NetworkACLRules
type NetworkACLRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkACLRule `json:"items"`
}

// Repository type metadata.
var (
	NetworkACLRule_Kind             = "NetworkACLRule"
	NetworkACLRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkACLRule_Kind}.String()
	NetworkACLRule_KindAPIVersion   = NetworkACLRule_Kind + "." + CRDGroupVersion.String()
	NetworkACLRule_GroupVersionKind = CRDGroupVersion.WithKind(NetworkACLRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkACLRule{}, &NetworkACLRuleList{})
}
