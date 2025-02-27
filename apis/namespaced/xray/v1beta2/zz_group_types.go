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

type GroupInitParameters struct {

	// The filter expression defining criteria by which to group traces. more info can be found in official docs.
	FilterExpression *string `json:"filterExpression,omitempty" tf:"filter_expression,omitempty"`

	// The name of the group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// Configuration options for enabling insights.
	InsightsConfiguration *InsightsConfigurationInitParameters `json:"insightsConfiguration,omitempty" tf:"insights_configuration,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GroupObservation struct {

	// The ARN of the Group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The filter expression defining criteria by which to group traces. more info can be found in official docs.
	FilterExpression *string `json:"filterExpression,omitempty" tf:"filter_expression,omitempty"`

	// The name of the group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// The ARN of the Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration options for enabling insights.
	InsightsConfiguration *InsightsConfigurationObservation `json:"insightsConfiguration,omitempty" tf:"insights_configuration,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GroupParameters struct {

	// The filter expression defining criteria by which to group traces. more info can be found in official docs.
	// +kubebuilder:validation:Optional
	FilterExpression *string `json:"filterExpression,omitempty" tf:"filter_expression,omitempty"`

	// The name of the group.
	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// Configuration options for enabling insights.
	// +kubebuilder:validation:Optional
	InsightsConfiguration *InsightsConfigurationParameters `json:"insightsConfiguration,omitempty" tf:"insights_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type InsightsConfigurationInitParameters struct {

	// Specifies whether insights are enabled.
	InsightsEnabled *bool `json:"insightsEnabled,omitempty" tf:"insights_enabled,omitempty"`

	// Specifies whether insight notifications are enabled.
	NotificationsEnabled *bool `json:"notificationsEnabled,omitempty" tf:"notifications_enabled,omitempty"`
}

type InsightsConfigurationObservation struct {

	// Specifies whether insights are enabled.
	InsightsEnabled *bool `json:"insightsEnabled,omitempty" tf:"insights_enabled,omitempty"`

	// Specifies whether insight notifications are enabled.
	NotificationsEnabled *bool `json:"notificationsEnabled,omitempty" tf:"notifications_enabled,omitempty"`
}

type InsightsConfigurationParameters struct {

	// Specifies whether insights are enabled.
	// +kubebuilder:validation:Optional
	InsightsEnabled *bool `json:"insightsEnabled" tf:"insights_enabled,omitempty"`

	// Specifies whether insight notifications are enabled.
	// +kubebuilder:validation:Optional
	NotificationsEnabled *bool `json:"notificationsEnabled,omitempty" tf:"notifications_enabled,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Group is the Schema for the Groups API. Creates and manages an AWS XRay Group.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filterExpression) || (has(self.initProvider) && has(self.initProvider.filterExpression))",message="spec.forProvider.filterExpression is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupName) || (has(self.initProvider) && has(self.initProvider.groupName))",message="spec.forProvider.groupName is a required parameter"
	Spec   GroupSpec   `json:"spec"`
	Status GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
