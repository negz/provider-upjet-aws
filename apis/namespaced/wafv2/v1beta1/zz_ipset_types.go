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

type IPSetInitParameters struct {

	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for /0.
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// A friendly description of the IP set.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify IPV4 or IPV6. Valid values are IPV4 or IPV6.
	IPAddressVersion *string `json:"ipAddressVersion,omitempty" tf:"ip_address_version,omitempty"`

	// A friendly name of the IP set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are CLOUDFRONT or REGIONAL. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IPSetObservation struct {

	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for /0.
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// The Amazon Resource Name (ARN) of the IP set.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A friendly description of the IP set.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique identifier for the IP set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specify IPV4 or IPV6. Valid values are IPV4 or IPV6.
	IPAddressVersion *string `json:"ipAddressVersion,omitempty" tf:"ip_address_version,omitempty"`

	LockToken *string `json:"lockToken,omitempty" tf:"lock_token,omitempty"`

	// A friendly name of the IP set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are CLOUDFRONT or REGIONAL. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type IPSetParameters struct {

	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for /0.
	// +kubebuilder:validation:Optional
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// A friendly description of the IP set.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify IPV4 or IPV6. Valid values are IPV4 or IPV6.
	// +kubebuilder:validation:Optional
	IPAddressVersion *string `json:"ipAddressVersion,omitempty" tf:"ip_address_version,omitempty"`

	// A friendly name of the IP set.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are CLOUDFRONT or REGIONAL. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// IPSetSpec defines the desired state of IPSet
type IPSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPSetParameters `json:"forProvider"`
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
	InitProvider IPSetInitParameters `json:"initProvider,omitempty"`
}

// IPSetStatus defines the observed state of IPSet.
type IPSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IPSet is the Schema for the IPSets API. Provides an AWS WAFv2 IP Set resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type IPSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipAddressVersion) || (has(self.initProvider) && has(self.initProvider.ipAddressVersion))",message="spec.forProvider.ipAddressVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope) || (has(self.initProvider) && has(self.initProvider.scope))",message="spec.forProvider.scope is a required parameter"
	Spec   IPSetSpec   `json:"spec"`
	Status IPSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPSetList contains a list of IPSets
type IPSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPSet `json:"items"`
}

// Repository type metadata.
var (
	IPSet_Kind             = "IPSet"
	IPSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPSet_Kind}.String()
	IPSet_KindAPIVersion   = IPSet_Kind + "." + CRDGroupVersion.String()
	IPSet_GroupVersionKind = CRDGroupVersion.WithKind(IPSet_Kind)
)

func init() {
	SchemeBuilder.Register(&IPSet{}, &IPSetList{})
}
