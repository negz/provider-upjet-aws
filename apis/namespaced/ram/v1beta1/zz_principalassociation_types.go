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

type PrincipalAssociationInitParameters struct {

	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/organizations/v1beta1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// Reference to a Organization in organizations to populate principal.
	// +kubebuilder:validation:Optional
	PrincipalRef *v1.Reference `json:"principalRef,omitempty" tf:"-"`

	// Selector for a Organization in organizations to populate principal.
	// +kubebuilder:validation:Optional
	PrincipalSelector *v1.Selector `json:"principalSelector,omitempty" tf:"-"`

	// The Amazon Resource Name (ARN) of the resource share.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ram/v1beta1.ResourceShare
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ResourceShareArn *string `json:"resourceShareArn,omitempty" tf:"resource_share_arn,omitempty"`

	// Reference to a ResourceShare in ram to populate resourceShareArn.
	// +kubebuilder:validation:Optional
	ResourceShareArnRef *v1.Reference `json:"resourceShareArnRef,omitempty" tf:"-"`

	// Selector for a ResourceShare in ram to populate resourceShareArn.
	// +kubebuilder:validation:Optional
	ResourceShareArnSelector *v1.Selector `json:"resourceShareArnSelector,omitempty" tf:"-"`
}

type PrincipalAssociationObservation struct {

	// The Amazon Resource Name (ARN) of the Resource Share and the principal, separated by a comma.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string `json:"resourceShareArn,omitempty" tf:"resource_share_arn,omitempty"`
}

type PrincipalAssociationParameters struct {

	// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/organizations/v1beta1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// Reference to a Organization in organizations to populate principal.
	// +kubebuilder:validation:Optional
	PrincipalRef *v1.Reference `json:"principalRef,omitempty" tf:"-"`

	// Selector for a Organization in organizations to populate principal.
	// +kubebuilder:validation:Optional
	PrincipalSelector *v1.Selector `json:"principalSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the resource share.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ram/v1beta1.ResourceShare
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceShareArn *string `json:"resourceShareArn,omitempty" tf:"resource_share_arn,omitempty"`

	// Reference to a ResourceShare in ram to populate resourceShareArn.
	// +kubebuilder:validation:Optional
	ResourceShareArnRef *v1.Reference `json:"resourceShareArnRef,omitempty" tf:"-"`

	// Selector for a ResourceShare in ram to populate resourceShareArn.
	// +kubebuilder:validation:Optional
	ResourceShareArnSelector *v1.Selector `json:"resourceShareArnSelector,omitempty" tf:"-"`
}

// PrincipalAssociationSpec defines the desired state of PrincipalAssociation
type PrincipalAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrincipalAssociationParameters `json:"forProvider"`
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
	InitProvider PrincipalAssociationInitParameters `json:"initProvider,omitempty"`
}

// PrincipalAssociationStatus defines the observed state of PrincipalAssociation.
type PrincipalAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrincipalAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrincipalAssociation is the Schema for the PrincipalAssociations API. Provides a Resource Access Manager (RAM) principal association.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type PrincipalAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrincipalAssociationSpec   `json:"spec"`
	Status            PrincipalAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrincipalAssociationList contains a list of PrincipalAssociations
type PrincipalAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrincipalAssociation `json:"items"`
}

// Repository type metadata.
var (
	PrincipalAssociation_Kind             = "PrincipalAssociation"
	PrincipalAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrincipalAssociation_Kind}.String()
	PrincipalAssociation_KindAPIVersion   = PrincipalAssociation_Kind + "." + CRDGroupVersion.String()
	PrincipalAssociation_GroupVersionKind = CRDGroupVersion.WithKind(PrincipalAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&PrincipalAssociation{}, &PrincipalAssociationList{})
}
