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

type ResolverConfigInitParameters struct {

	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: ENABLE, DISABLE.
	AutodefinedReverseFlag *string `json:"autodefinedReverseFlag,omitempty" tf:"autodefined_reverse_flag,omitempty"`

	// The ID of the VPC that the configuration is for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a VPC in ec2 to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

type ResolverConfigObservation struct {

	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: ENABLE, DISABLE.
	AutodefinedReverseFlag *string `json:"autodefinedReverseFlag,omitempty" tf:"autodefined_reverse_flag,omitempty"`

	// The ID of the resolver configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS account ID of the owner of the VPC that this resolver configuration applies to.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The ID of the VPC that the configuration is for.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type ResolverConfigParameters struct {

	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: ENABLE, DISABLE.
	// +kubebuilder:validation:Optional
	AutodefinedReverseFlag *string `json:"autodefinedReverseFlag,omitempty" tf:"autodefined_reverse_flag,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the VPC that the configuration is for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a VPC in ec2 to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

// ResolverConfigSpec defines the desired state of ResolverConfig
type ResolverConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResolverConfigParameters `json:"forProvider"`
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
	InitProvider ResolverConfigInitParameters `json:"initProvider,omitempty"`
}

// ResolverConfigStatus defines the observed state of ResolverConfig.
type ResolverConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResolverConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ResolverConfig is the Schema for the ResolverConfigs API. Provides a Route 53 Resolver config resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type ResolverConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autodefinedReverseFlag) || (has(self.initProvider) && has(self.initProvider.autodefinedReverseFlag))",message="spec.forProvider.autodefinedReverseFlag is a required parameter"
	Spec   ResolverConfigSpec   `json:"spec"`
	Status ResolverConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverConfigList contains a list of ResolverConfigs
type ResolverConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResolverConfig `json:"items"`
}

// Repository type metadata.
var (
	ResolverConfig_Kind             = "ResolverConfig"
	ResolverConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResolverConfig_Kind}.String()
	ResolverConfig_KindAPIVersion   = ResolverConfig_Kind + "." + CRDGroupVersion.String()
	ResolverConfig_GroupVersionKind = CRDGroupVersion.WithKind(ResolverConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ResolverConfig{}, &ResolverConfigList{})
}
