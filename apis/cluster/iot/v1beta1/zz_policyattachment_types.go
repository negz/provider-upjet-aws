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

type PolicyAttachmentInitParameters struct {

	// The name of the policy to attach.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iot/v1beta1.Policy
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Reference to a Policy in iot to populate policy.
	// +kubebuilder:validation:Optional
	PolicyRef *v1.Reference `json:"policyRef,omitempty" tf:"-"`

	// Selector for a Policy in iot to populate policy.
	// +kubebuilder:validation:Optional
	PolicySelector *v1.Selector `json:"policySelector,omitempty" tf:"-"`

	// The identity to which the policy is attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iot/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Reference to a Certificate in iot to populate target.
	// +kubebuilder:validation:Optional
	TargetRef *v1.Reference `json:"targetRef,omitempty" tf:"-"`

	// Selector for a Certificate in iot to populate target.
	// +kubebuilder:validation:Optional
	TargetSelector *v1.Selector `json:"targetSelector,omitempty" tf:"-"`
}

type PolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the policy to attach.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The identity to which the policy is attached.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type PolicyAttachmentParameters struct {

	// The name of the policy to attach.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iot/v1beta1.Policy
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Reference to a Policy in iot to populate policy.
	// +kubebuilder:validation:Optional
	PolicyRef *v1.Reference `json:"policyRef,omitempty" tf:"-"`

	// Selector for a Policy in iot to populate policy.
	// +kubebuilder:validation:Optional
	PolicySelector *v1.Selector `json:"policySelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The identity to which the policy is attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iot/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Reference to a Certificate in iot to populate target.
	// +kubebuilder:validation:Optional
	TargetRef *v1.Reference `json:"targetRef,omitempty" tf:"-"`

	// Selector for a Certificate in iot to populate target.
	// +kubebuilder:validation:Optional
	TargetSelector *v1.Selector `json:"targetSelector,omitempty" tf:"-"`
}

// PolicyAttachmentSpec defines the desired state of PolicyAttachment
type PolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyAttachmentParameters `json:"forProvider"`
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
	InitProvider PolicyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// PolicyAttachmentStatus defines the observed state of PolicyAttachment.
type PolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PolicyAttachment is the Schema for the PolicyAttachments API. Provides an IoT policy attachment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyAttachmentSpec   `json:"spec"`
	Status            PolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyAttachmentList contains a list of PolicyAttachments
type PolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	PolicyAttachment_Kind             = "PolicyAttachment"
	PolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyAttachment_Kind}.String()
	PolicyAttachment_KindAPIVersion   = PolicyAttachment_Kind + "." + CRDGroupVersion.String()
	PolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(PolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyAttachment{}, &PolicyAttachmentList{})
}
