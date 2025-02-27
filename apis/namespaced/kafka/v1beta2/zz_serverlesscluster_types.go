// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IAMInitParameters struct {

	// Whether SASL/IAM authentication is enabled or not.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type IAMObservation struct {

	// Whether SASL/IAM authentication is enabled or not.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type IAMParameters struct {

	// Whether SASL/IAM authentication is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type ServerlessClusterClientAuthenticationInitParameters struct {

	// Details for client authentication using SASL. See below.
	Sasl *ServerlessClusterClientAuthenticationSaslInitParameters `json:"sasl,omitempty" tf:"sasl,omitempty"`
}

type ServerlessClusterClientAuthenticationObservation struct {

	// Details for client authentication using SASL. See below.
	Sasl *ServerlessClusterClientAuthenticationSaslObservation `json:"sasl,omitempty" tf:"sasl,omitempty"`
}

type ServerlessClusterClientAuthenticationParameters struct {

	// Details for client authentication using SASL. See below.
	// +kubebuilder:validation:Optional
	Sasl *ServerlessClusterClientAuthenticationSaslParameters `json:"sasl" tf:"sasl,omitempty"`
}

type ServerlessClusterClientAuthenticationSaslInitParameters struct {

	// Details for client authentication using IAM. See below.
	IAM *IAMInitParameters `json:"iam,omitempty" tf:"iam,omitempty"`
}

type ServerlessClusterClientAuthenticationSaslObservation struct {

	// Details for client authentication using IAM. See below.
	IAM *IAMObservation `json:"iam,omitempty" tf:"iam,omitempty"`
}

type ServerlessClusterClientAuthenticationSaslParameters struct {

	// Details for client authentication using IAM. See below.
	// +kubebuilder:validation:Optional
	IAM *IAMParameters `json:"iam" tf:"iam,omitempty"`
}

type ServerlessClusterInitParameters struct {

	// Specifies client authentication information for the serverless cluster. See below.
	ClientAuthentication *ServerlessClusterClientAuthenticationInitParameters `json:"clientAuthentication,omitempty" tf:"client_authentication,omitempty"`

	// The name of the serverless cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// VPC configuration information. See below.
	VPCConfig []VPCConfigInitParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ServerlessClusterObservation struct {

	// The ARN of the serverless cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies client authentication information for the serverless cluster. See below.
	ClientAuthentication *ServerlessClusterClientAuthenticationObservation `json:"clientAuthentication,omitempty" tf:"client_authentication,omitempty"`

	// The name of the serverless cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// UUID of the serverless cluster, for use in IAM policies.
	ClusterUUID *string `json:"clusterUuid,omitempty" tf:"cluster_uuid,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// VPC configuration information. See below.
	VPCConfig []VPCConfigObservation `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ServerlessClusterParameters struct {

	// Specifies client authentication information for the serverless cluster. See below.
	// +kubebuilder:validation:Optional
	ClientAuthentication *ServerlessClusterClientAuthenticationParameters `json:"clientAuthentication,omitempty" tf:"client_authentication,omitempty"`

	// The name of the serverless cluster.
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// VPC configuration information. See below.
	// +kubebuilder:validation:Optional
	VPCConfig []VPCConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type VPCConfigInitParameters struct {

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies up to five security groups that control inbound and outbound traffic for the serverless cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of subnets in at least two different Availability Zones that host your client applications.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type VPCConfigObservation struct {

	// Specifies up to five security groups that control inbound and outbound traffic for the serverless cluster.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A list of subnets in at least two different Availability Zones that host your client applications.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type VPCConfigParameters struct {

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies up to five security groups that control inbound and outbound traffic for the serverless cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of subnets in at least two different Availability Zones that host your client applications.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

// ServerlessClusterSpec defines the desired state of ServerlessCluster
type ServerlessClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerlessClusterParameters `json:"forProvider"`
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
	InitProvider ServerlessClusterInitParameters `json:"initProvider,omitempty"`
}

// ServerlessClusterStatus defines the observed state of ServerlessCluster.
type ServerlessClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerlessClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ServerlessCluster is the Schema for the ServerlessClusters API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServerlessCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientAuthentication) || (has(self.initProvider) && has(self.initProvider.clientAuthentication))",message="spec.forProvider.clientAuthentication is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterName) || (has(self.initProvider) && has(self.initProvider.clusterName))",message="spec.forProvider.clusterName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcConfig) || (has(self.initProvider) && has(self.initProvider.vpcConfig))",message="spec.forProvider.vpcConfig is a required parameter"
	Spec   ServerlessClusterSpec   `json:"spec"`
	Status ServerlessClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerlessClusterList contains a list of ServerlessClusters
type ServerlessClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerlessCluster `json:"items"`
}

// Repository type metadata.
var (
	ServerlessCluster_Kind             = "ServerlessCluster"
	ServerlessCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerlessCluster_Kind}.String()
	ServerlessCluster_KindAPIVersion   = ServerlessCluster_Kind + "." + CRDGroupVersion.String()
	ServerlessCluster_GroupVersionKind = CRDGroupVersion.WithKind(ServerlessCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerlessCluster{}, &ServerlessClusterList{})
}
