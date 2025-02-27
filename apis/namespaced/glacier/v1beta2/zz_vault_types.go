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

type NotificationInitParameters struct {

	// You can configure a vault to publish a notification for ArchiveRetrievalCompleted and InventoryRetrievalCompleted events.
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// The SNS Topic ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	SnsTopic *string `json:"snsTopic,omitempty" tf:"sns_topic,omitempty"`

	// Reference to a Topic in sns to populate snsTopic.
	// +kubebuilder:validation:Optional
	SnsTopicRef *v1.Reference `json:"snsTopicRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopic.
	// +kubebuilder:validation:Optional
	SnsTopicSelector *v1.Selector `json:"snsTopicSelector,omitempty" tf:"-"`
}

type NotificationObservation struct {

	// You can configure a vault to publish a notification for ArchiveRetrievalCompleted and InventoryRetrievalCompleted events.
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// The SNS Topic ARN.
	SnsTopic *string `json:"snsTopic,omitempty" tf:"sns_topic,omitempty"`
}

type NotificationParameters struct {

	// You can configure a vault to publish a notification for ArchiveRetrievalCompleted and InventoryRetrievalCompleted events.
	// +kubebuilder:validation:Optional
	// +listType=set
	Events []*string `json:"events" tf:"events,omitempty"`

	// The SNS Topic ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SnsTopic *string `json:"snsTopic,omitempty" tf:"sns_topic,omitempty"`

	// Reference to a Topic in sns to populate snsTopic.
	// +kubebuilder:validation:Optional
	SnsTopicRef *v1.Reference `json:"snsTopicRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopic.
	// +kubebuilder:validation:Optional
	SnsTopicSelector *v1.Selector `json:"snsTopicSelector,omitempty" tf:"-"`
}

type VaultInitParameters struct {

	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or file function is helpful here. Use the Glacier Developer Guide for more information on Glacier Vault Policy
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The notifications for the Vault. Fields documented below.
	Notification *NotificationInitParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VaultObservation struct {

	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or file function is helpful here. Use the Glacier Developer Guide for more information on Glacier Vault Policy
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The ARN of the vault.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URI of the vault that was created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The notifications for the Vault. Fields documented below.
	Notification *NotificationObservation `json:"notification,omitempty" tf:"notification,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VaultParameters struct {

	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or file function is helpful here. Use the Glacier Developer Guide for more information on Glacier Vault Policy
	// +kubebuilder:validation:Optional
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The notifications for the Vault. Fields documented below.
	// +kubebuilder:validation:Optional
	Notification *NotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VaultSpec defines the desired state of Vault
type VaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultParameters `json:"forProvider"`
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
	InitProvider VaultInitParameters `json:"initProvider,omitempty"`
}

// VaultStatus defines the observed state of Vault.
type VaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Vault is the Schema for the Vaults API. Provides a Glacier Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Vault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultSpec   `json:"spec"`
	Status            VaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultList contains a list of Vaults
type VaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vault `json:"items"`
}

// Repository type metadata.
var (
	Vault_Kind             = "Vault"
	Vault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vault_Kind}.String()
	Vault_KindAPIVersion   = Vault_Kind + "." + CRDGroupVersion.String()
	Vault_GroupVersionKind = CRDGroupVersion.WithKind(Vault_Kind)
)

func init() {
	SchemeBuilder.Register(&Vault{}, &VaultList{})
}
