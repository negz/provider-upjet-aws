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

type GlobalNodeGroupsInitParameters struct {
}

type GlobalNodeGroupsObservation struct {

	// The ID of the global node group.
	GlobalNodeGroupID *string `json:"globalNodeGroupId,omitempty" tf:"global_node_group_id,omitempty"`

	// The keyspace for this node group.
	Slots *string `json:"slots,omitempty" tf:"slots,omitempty"`
}

type GlobalNodeGroupsParameters struct {
}

type GlobalReplicationGroupInitParameters struct {

	// Specifies whether read-only replicas will be automatically promoted to read/write primary if the existing primary fails.
	// When creating, by default the Global Replication Group inherits the automatic failover setting of the primary replication group.
	AutomaticFailoverEnabled *bool `json:"automaticFailoverEnabled,omitempty" tf:"automatic_failover_enabled,omitempty"`

	// The instance class used.
	// See AWS documentation for information on supported node types
	// and guidance on selecting node types.
	// When creating, by default the Global Replication Group inherits the node type of the primary replication group.
	CacheNodeType *string `json:"cacheNodeType,omitempty" tf:"cache_node_type,omitempty"`

	// Redis version to use for the Global Replication Group.
	// When creating, by default the Global Replication Group inherits the version of the primary replication group.
	// If a version is specified, the Global Replication Group and all member replication groups will be upgraded to this version.
	// Cannot be downgraded without replacing the Global Replication Group and all member replication groups.
	// When the version is 7 or higher, the major and minor version should be set, e.g., 7.2.
	// When the version is 6, the major and minor version can be set, e.g., 6.2,
	// or the minor version can be unspecified which will use the latest version at creation time, e.g., 6.x.
	// The actual engine version used is returned in the attribute engine_version_actual, see Attribute Reference below.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// created description for the global replication group.
	GlobalReplicationGroupDescription *string `json:"globalReplicationGroupDescription,omitempty" tf:"global_replication_group_description,omitempty"`

	// –  The suffix name of a Global Datastore. If global_replication_group_id_suffix is changed, creates a new resource.
	GlobalReplicationGroupIDSuffix *string `json:"globalReplicationGroupIdSuffix,omitempty" tf:"global_replication_group_id_suffix,omitempty"`

	// The number of node groups (shards) on the global replication group.
	NumNodeGroups *float64 `json:"numNodeGroups,omitempty" tf:"num_node_groups,omitempty"`

	// An ElastiCache Parameter Group to use for the Global Replication Group.
	// Required when upgrading a major engine version, but will be ignored if left configured after the upgrade is complete.
	// Specifying without a major version upgrade will fail.
	// Note that ElastiCache creates a copy of this parameter group for each member replication group.
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// –  The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If primary_replication_group_id is changed, creates a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticache/v1beta2.ReplicationGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PrimaryReplicationGroupID *string `json:"primaryReplicationGroupId,omitempty" tf:"primary_replication_group_id,omitempty"`

	// Reference to a ReplicationGroup in elasticache to populate primaryReplicationGroupId.
	// +kubebuilder:validation:Optional
	PrimaryReplicationGroupIDRef *v1.Reference `json:"primaryReplicationGroupIdRef,omitempty" tf:"-"`

	// Selector for a ReplicationGroup in elasticache to populate primaryReplicationGroupId.
	// +kubebuilder:validation:Optional
	PrimaryReplicationGroupIDSelector *v1.Selector `json:"primaryReplicationGroupIdSelector,omitempty" tf:"-"`
}

type GlobalReplicationGroupObservation struct {

	// The ARN of the ElastiCache Global Replication Group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A flag that indicate whether the encryption at rest is enabled.
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`

	// A flag that indicate whether AuthToken (password) is enabled.
	AuthTokenEnabled *bool `json:"authTokenEnabled,omitempty" tf:"auth_token_enabled,omitempty"`

	// Specifies whether read-only replicas will be automatically promoted to read/write primary if the existing primary fails.
	// When creating, by default the Global Replication Group inherits the automatic failover setting of the primary replication group.
	AutomaticFailoverEnabled *bool `json:"automaticFailoverEnabled,omitempty" tf:"automatic_failover_enabled,omitempty"`

	// The instance class used.
	// See AWS documentation for information on supported node types
	// and guidance on selecting node types.
	// When creating, by default the Global Replication Group inherits the node type of the primary replication group.
	CacheNodeType *string `json:"cacheNodeType,omitempty" tf:"cache_node_type,omitempty"`

	// Indicates whether the Global Datastore is cluster enabled.
	ClusterEnabled *bool `json:"clusterEnabled,omitempty" tf:"cluster_enabled,omitempty"`

	// The name of the cache engine to be used for the clusters in this global replication group.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Redis version to use for the Global Replication Group.
	// When creating, by default the Global Replication Group inherits the version of the primary replication group.
	// If a version is specified, the Global Replication Group and all member replication groups will be upgraded to this version.
	// Cannot be downgraded without replacing the Global Replication Group and all member replication groups.
	// When the version is 7 or higher, the major and minor version should be set, e.g., 7.2.
	// When the version is 6, the major and minor version can be set, e.g., 6.2,
	// or the minor version can be unspecified which will use the latest version at creation time, e.g., 6.x.
	// The actual engine version used is returned in the attribute engine_version_actual, see Attribute Reference below.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The full version number of the cache engine running on the members of this global replication group.
	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	// Set of node groups (shards) on the global replication group.
	// Has the values:
	GlobalNodeGroups []GlobalNodeGroupsObservation `json:"globalNodeGroups,omitempty" tf:"global_node_groups,omitempty"`

	// created description for the global replication group.
	GlobalReplicationGroupDescription *string `json:"globalReplicationGroupDescription,omitempty" tf:"global_replication_group_description,omitempty"`

	// The full ID of the global replication group.
	GlobalReplicationGroupID *string `json:"globalReplicationGroupId,omitempty" tf:"global_replication_group_id,omitempty"`

	// –  The suffix name of a Global Datastore. If global_replication_group_id_suffix is changed, creates a new resource.
	GlobalReplicationGroupIDSuffix *string `json:"globalReplicationGroupIdSuffix,omitempty" tf:"global_replication_group_id_suffix,omitempty"`

	// The ID of the ElastiCache Global Replication Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The number of node groups (shards) on the global replication group.
	NumNodeGroups *float64 `json:"numNodeGroups,omitempty" tf:"num_node_groups,omitempty"`

	// An ElastiCache Parameter Group to use for the Global Replication Group.
	// Required when upgrading a major engine version, but will be ignored if left configured after the upgrade is complete.
	// Specifying without a major version upgrade will fail.
	// Note that ElastiCache creates a copy of this parameter group for each member replication group.
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// –  The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If primary_replication_group_id is changed, creates a new resource.
	PrimaryReplicationGroupID *string `json:"primaryReplicationGroupId,omitempty" tf:"primary_replication_group_id,omitempty"`

	// A flag that indicates whether the encryption in transit is enabled.
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
}

type GlobalReplicationGroupParameters struct {

	// Specifies whether read-only replicas will be automatically promoted to read/write primary if the existing primary fails.
	// When creating, by default the Global Replication Group inherits the automatic failover setting of the primary replication group.
	// +kubebuilder:validation:Optional
	AutomaticFailoverEnabled *bool `json:"automaticFailoverEnabled,omitempty" tf:"automatic_failover_enabled,omitempty"`

	// The instance class used.
	// See AWS documentation for information on supported node types
	// and guidance on selecting node types.
	// When creating, by default the Global Replication Group inherits the node type of the primary replication group.
	// +kubebuilder:validation:Optional
	CacheNodeType *string `json:"cacheNodeType,omitempty" tf:"cache_node_type,omitempty"`

	// Redis version to use for the Global Replication Group.
	// When creating, by default the Global Replication Group inherits the version of the primary replication group.
	// If a version is specified, the Global Replication Group and all member replication groups will be upgraded to this version.
	// Cannot be downgraded without replacing the Global Replication Group and all member replication groups.
	// When the version is 7 or higher, the major and minor version should be set, e.g., 7.2.
	// When the version is 6, the major and minor version can be set, e.g., 6.2,
	// or the minor version can be unspecified which will use the latest version at creation time, e.g., 6.x.
	// The actual engine version used is returned in the attribute engine_version_actual, see Attribute Reference below.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// created description for the global replication group.
	// +kubebuilder:validation:Optional
	GlobalReplicationGroupDescription *string `json:"globalReplicationGroupDescription,omitempty" tf:"global_replication_group_description,omitempty"`

	// –  The suffix name of a Global Datastore. If global_replication_group_id_suffix is changed, creates a new resource.
	// +kubebuilder:validation:Optional
	GlobalReplicationGroupIDSuffix *string `json:"globalReplicationGroupIdSuffix,omitempty" tf:"global_replication_group_id_suffix,omitempty"`

	// The number of node groups (shards) on the global replication group.
	// +kubebuilder:validation:Optional
	NumNodeGroups *float64 `json:"numNodeGroups,omitempty" tf:"num_node_groups,omitempty"`

	// An ElastiCache Parameter Group to use for the Global Replication Group.
	// Required when upgrading a major engine version, but will be ignored if left configured after the upgrade is complete.
	// Specifying without a major version upgrade will fail.
	// Note that ElastiCache creates a copy of this parameter group for each member replication group.
	// +kubebuilder:validation:Optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// –  The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If primary_replication_group_id is changed, creates a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elasticache/v1beta2.ReplicationGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrimaryReplicationGroupID *string `json:"primaryReplicationGroupId,omitempty" tf:"primary_replication_group_id,omitempty"`

	// Reference to a ReplicationGroup in elasticache to populate primaryReplicationGroupId.
	// +kubebuilder:validation:Optional
	PrimaryReplicationGroupIDRef *v1.Reference `json:"primaryReplicationGroupIdRef,omitempty" tf:"-"`

	// Selector for a ReplicationGroup in elasticache to populate primaryReplicationGroupId.
	// +kubebuilder:validation:Optional
	PrimaryReplicationGroupIDSelector *v1.Selector `json:"primaryReplicationGroupIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// GlobalReplicationGroupSpec defines the desired state of GlobalReplicationGroup
type GlobalReplicationGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalReplicationGroupParameters `json:"forProvider"`
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
	InitProvider GlobalReplicationGroupInitParameters `json:"initProvider,omitempty"`
}

// GlobalReplicationGroupStatus defines the observed state of GlobalReplicationGroup.
type GlobalReplicationGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalReplicationGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GlobalReplicationGroup is the Schema for the GlobalReplicationGroups API. Provides an ElastiCache Global Replication Group resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type GlobalReplicationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.globalReplicationGroupIdSuffix) || (has(self.initProvider) && has(self.initProvider.globalReplicationGroupIdSuffix))",message="spec.forProvider.globalReplicationGroupIdSuffix is a required parameter"
	Spec   GlobalReplicationGroupSpec   `json:"spec"`
	Status GlobalReplicationGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalReplicationGroupList contains a list of GlobalReplicationGroups
type GlobalReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalReplicationGroup `json:"items"`
}

// Repository type metadata.
var (
	GlobalReplicationGroup_Kind             = "GlobalReplicationGroup"
	GlobalReplicationGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalReplicationGroup_Kind}.String()
	GlobalReplicationGroup_KindAPIVersion   = GlobalReplicationGroup_Kind + "." + CRDGroupVersion.String()
	GlobalReplicationGroup_GroupVersionKind = CRDGroupVersion.WithKind(GlobalReplicationGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalReplicationGroup{}, &GlobalReplicationGroupList{})
}
