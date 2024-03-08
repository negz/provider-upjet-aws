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

type InlinePolicyInitParameters struct {

	// Friendly name of the role. See IAM Identifiers for more information.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Policy document as a JSON formatted string.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type InlinePolicyObservation struct {

	// Friendly name of the role. See IAM Identifiers for more information.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Policy document as a JSON formatted string.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type InlinePolicyParameters struct {

	// Friendly name of the role. See IAM Identifiers for more information.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Policy document as a JSON formatted string.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type RoleInitParameters struct {

	// Policy that grants an entity permission to assume the role.
	AssumeRolePolicy *string `json:"assumeRolePolicy,omitempty" tf:"assume_role_policy,omitempty"`

	// Description of the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to force detaching any policies the role has before destroying it. Defaults to false.
	ForceDetachPolicies *bool `json:"forceDetachPolicies,omitempty" tf:"force_detach_policies,omitempty"`

	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, Crossplane will not manage any inline policies in this resource. Configuring one empty block (i.e., inline_policy {}) will cause Crossplane to remove all inline policies added out of band on apply.
	InlinePolicy []InlinePolicyInitParameters `json:"inlinePolicy,omitempty" tf:"inline_policy,omitempty"`

	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, Crossplane will ignore policy attachments to this resource. When configured, Crossplane will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., managed_policy_arns = []) will cause Crossplane to remove all managed policy attachments.
	// +listType=set
	ManagedPolicyArns []*string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns,omitempty"`

	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration *float64 `json:"maxSessionDuration,omitempty" tf:"max_session_duration,omitempty"`

	// Path to the role. See IAM Identifiers for more information.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RoleObservation struct {

	// Amazon Resource Name (ARN) specifying the role.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Policy that grants an entity permission to assume the role.
	AssumeRolePolicy *string `json:"assumeRolePolicy,omitempty" tf:"assume_role_policy,omitempty"`

	// Creation date of the IAM role.
	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	// Description of the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to force detaching any policies the role has before destroying it. Defaults to false.
	ForceDetachPolicies *bool `json:"forceDetachPolicies,omitempty" tf:"force_detach_policies,omitempty"`

	// Name of the role.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, Crossplane will not manage any inline policies in this resource. Configuring one empty block (i.e., inline_policy {}) will cause Crossplane to remove all inline policies added out of band on apply.
	InlinePolicy []InlinePolicyObservation `json:"inlinePolicy,omitempty" tf:"inline_policy,omitempty"`

	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, Crossplane will ignore policy attachments to this resource. When configured, Crossplane will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., managed_policy_arns = []) will cause Crossplane to remove all managed policy attachments.
	// +listType=set
	ManagedPolicyArns []*string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns,omitempty"`

	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration *float64 `json:"maxSessionDuration,omitempty" tf:"max_session_duration,omitempty"`

	// Path to the role. See IAM Identifiers for more information.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Stable and unique string identifying the role.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type RoleParameters struct {

	// Policy that grants an entity permission to assume the role.
	// +kubebuilder:validation:Optional
	AssumeRolePolicy *string `json:"assumeRolePolicy,omitempty" tf:"assume_role_policy,omitempty"`

	// Description of the role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to force detaching any policies the role has before destroying it. Defaults to false.
	// +kubebuilder:validation:Optional
	ForceDetachPolicies *bool `json:"forceDetachPolicies,omitempty" tf:"force_detach_policies,omitempty"`

	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, Crossplane will not manage any inline policies in this resource. Configuring one empty block (i.e., inline_policy {}) will cause Crossplane to remove all inline policies added out of band on apply.
	// +kubebuilder:validation:Optional
	InlinePolicy []InlinePolicyParameters `json:"inlinePolicy,omitempty" tf:"inline_policy,omitempty"`

	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, Crossplane will ignore policy attachments to this resource. When configured, Crossplane will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., managed_policy_arns = []) will cause Crossplane to remove all managed policy attachments.
	// +kubebuilder:validation:Optional
	// +listType=set
	ManagedPolicyArns []*string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns,omitempty"`

	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	// +kubebuilder:validation:Optional
	MaxSessionDuration *float64 `json:"maxSessionDuration,omitempty" tf:"max_session_duration,omitempty"`

	// Path to the role. See IAM Identifiers for more information.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// ARN of the policy that is used to set the permissions boundary for the role.
	// +kubebuilder:validation:Optional
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
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
	InitProvider RoleInitParameters `json:"initProvider,omitempty"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Role is the Schema for the Roles API. Provides an IAM role.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.assumeRolePolicy) || (has(self.initProvider) && has(self.initProvider.assumeRolePolicy))",message="spec.forProvider.assumeRolePolicy is a required parameter"
	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
