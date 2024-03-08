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

type ACLInitParameters struct {

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Set of MemoryDB user names to be included in this ACL.
	// +listType=set
	UserNames []*string `json:"userNames,omitempty" tf:"user_names,omitempty"`
}

type ACLObservation struct {

	// The ARN of the ACL.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Same as name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The minimum engine version supported by the ACL.
	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty" tf:"minimum_engine_version,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Set of MemoryDB user names to be included in this ACL.
	// +listType=set
	UserNames []*string `json:"userNames,omitempty" tf:"user_names,omitempty"`
}

type ACLParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Set of MemoryDB user names to be included in this ACL.
	// +kubebuilder:validation:Optional
	// +listType=set
	UserNames []*string `json:"userNames,omitempty" tf:"user_names,omitempty"`
}

// ACLSpec defines the desired state of ACL
type ACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ACLParameters `json:"forProvider"`
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
	InitProvider ACLInitParameters `json:"initProvider,omitempty"`
}

// ACLStatus defines the observed state of ACL.
type ACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ACL is the Schema for the ACLs API. Provides a MemoryDB ACL.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ACLSpec   `json:"spec"`
	Status            ACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ACLList contains a list of ACLs
type ACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ACL `json:"items"`
}

// Repository type metadata.
var (
	ACL_Kind             = "ACL"
	ACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ACL_Kind}.String()
	ACL_KindAPIVersion   = ACL_Kind + "." + CRDGroupVersion.String()
	ACL_GroupVersionKind = CRDGroupVersion.WithKind(ACL_Kind)
)

func init() {
	SchemeBuilder.Register(&ACL{}, &ACLList{})
}
