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

type UserPolicyAttachmentInitParameters struct {

	// The ARN of the policy you want to apply
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Policy
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// Reference to a Policy in iam to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// Selector for a Policy in iam to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`

	// The user the policy should be applied to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.User
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User in iam to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

type UserPolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the policy you want to apply
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// The user the policy should be applied to
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type UserPolicyAttachmentParameters struct {

	// The ARN of the policy you want to apply
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Policy
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// Reference to a Policy in iam to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// Selector for a Policy in iam to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`

	// The user the policy should be applied to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.User
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User in iam to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// UserPolicyAttachmentSpec defines the desired state of UserPolicyAttachment
type UserPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPolicyAttachmentParameters `json:"forProvider"`
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
	InitProvider UserPolicyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// UserPolicyAttachmentStatus defines the observed state of UserPolicyAttachment.
type UserPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserPolicyAttachment is the Schema for the UserPolicyAttachments API. Attaches a Managed IAM Policy to an IAM user
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type UserPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPolicyAttachmentSpec   `json:"spec"`
	Status            UserPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPolicyAttachmentList contains a list of UserPolicyAttachments
type UserPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	UserPolicyAttachment_Kind             = "UserPolicyAttachment"
	UserPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPolicyAttachment_Kind}.String()
	UserPolicyAttachment_KindAPIVersion   = UserPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	UserPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(UserPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPolicyAttachment{}, &UserPolicyAttachmentList{})
}
