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

type AlertManagerDefinitionInitParameters struct {

	// the alert manager definition that you want to be applied. See more in AWS Docs.
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// ID of the prometheus workspace the alert manager definition should be linked to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/amp/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in amp to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in amp to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

type AlertManagerDefinitionObservation struct {

	// the alert manager definition that you want to be applied. See more in AWS Docs.
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the prometheus workspace the alert manager definition should be linked to
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type AlertManagerDefinitionParameters struct {

	// the alert manager definition that you want to be applied. See more in AWS Docs.
	// +kubebuilder:validation:Optional
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the prometheus workspace the alert manager definition should be linked to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/amp/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in amp to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in amp to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// AlertManagerDefinitionSpec defines the desired state of AlertManagerDefinition
type AlertManagerDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertManagerDefinitionParameters `json:"forProvider"`
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
	InitProvider AlertManagerDefinitionInitParameters `json:"initProvider,omitempty"`
}

// AlertManagerDefinitionStatus defines the observed state of AlertManagerDefinition.
type AlertManagerDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertManagerDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AlertManagerDefinition is the Schema for the AlertManagerDefinitions API. Manages an Amazon Managed Service for Prometheus (AMP) Alert Manager Definition
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AlertManagerDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.definition) || (has(self.initProvider) && has(self.initProvider.definition))",message="spec.forProvider.definition is a required parameter"
	Spec   AlertManagerDefinitionSpec   `json:"spec"`
	Status AlertManagerDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertManagerDefinitionList contains a list of AlertManagerDefinitions
type AlertManagerDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertManagerDefinition `json:"items"`
}

// Repository type metadata.
var (
	AlertManagerDefinition_Kind             = "AlertManagerDefinition"
	AlertManagerDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlertManagerDefinition_Kind}.String()
	AlertManagerDefinition_KindAPIVersion   = AlertManagerDefinition_Kind + "." + CRDGroupVersion.String()
	AlertManagerDefinition_GroupVersionKind = CRDGroupVersion.WithKind(AlertManagerDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&AlertManagerDefinition{}, &AlertManagerDefinitionList{})
}
