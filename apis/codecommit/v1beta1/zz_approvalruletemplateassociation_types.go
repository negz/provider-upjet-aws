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

type ApprovalRuleTemplateAssociationInitParameters struct {
}

type ApprovalRuleTemplateAssociationObservation struct {

	// The name for the approval rule template.
	ApprovalRuleTemplateName *string `json:"approvalRuleTemplateName,omitempty" tf:"approval_rule_template_name,omitempty"`

	// The name of the approval rule template and name of the repository, separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the repository that you want to associate with the template.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`
}

type ApprovalRuleTemplateAssociationParameters struct {

	// The name for the approval rule template.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codecommit/v1beta1.ApprovalRuleTemplate
	// +kubebuilder:validation:Optional
	ApprovalRuleTemplateName *string `json:"approvalRuleTemplateName,omitempty" tf:"approval_rule_template_name,omitempty"`

	// Reference to a ApprovalRuleTemplate in codecommit to populate approvalRuleTemplateName.
	// +kubebuilder:validation:Optional
	ApprovalRuleTemplateNameRef *v1.Reference `json:"approvalRuleTemplateNameRef,omitempty" tf:"-"`

	// Selector for a ApprovalRuleTemplate in codecommit to populate approvalRuleTemplateName.
	// +kubebuilder:validation:Optional
	ApprovalRuleTemplateNameSelector *v1.Selector `json:"approvalRuleTemplateNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the repository that you want to associate with the template.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codecommit/v1beta1.Repository
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Reference to a Repository in codecommit to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameRef *v1.Reference `json:"repositoryNameRef,omitempty" tf:"-"`

	// Selector for a Repository in codecommit to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameSelector *v1.Selector `json:"repositoryNameSelector,omitempty" tf:"-"`
}

// ApprovalRuleTemplateAssociationSpec defines the desired state of ApprovalRuleTemplateAssociation
type ApprovalRuleTemplateAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApprovalRuleTemplateAssociationParameters `json:"forProvider"`
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
	InitProvider ApprovalRuleTemplateAssociationInitParameters `json:"initProvider,omitempty"`
}

// ApprovalRuleTemplateAssociationStatus defines the observed state of ApprovalRuleTemplateAssociation.
type ApprovalRuleTemplateAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApprovalRuleTemplateAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ApprovalRuleTemplateAssociation is the Schema for the ApprovalRuleTemplateAssociations API. Associates a CodeCommit Approval Rule Template with a Repository.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApprovalRuleTemplateAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApprovalRuleTemplateAssociationSpec   `json:"spec"`
	Status            ApprovalRuleTemplateAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApprovalRuleTemplateAssociationList contains a list of ApprovalRuleTemplateAssociations
type ApprovalRuleTemplateAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApprovalRuleTemplateAssociation `json:"items"`
}

// Repository type metadata.
var (
	ApprovalRuleTemplateAssociation_Kind             = "ApprovalRuleTemplateAssociation"
	ApprovalRuleTemplateAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApprovalRuleTemplateAssociation_Kind}.String()
	ApprovalRuleTemplateAssociation_KindAPIVersion   = ApprovalRuleTemplateAssociation_Kind + "." + CRDGroupVersion.String()
	ApprovalRuleTemplateAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ApprovalRuleTemplateAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ApprovalRuleTemplateAssociation{}, &ApprovalRuleTemplateAssociationList{})
}
