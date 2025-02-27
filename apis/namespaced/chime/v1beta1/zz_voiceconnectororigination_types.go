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

type RouteInitParameters struct {

	// The FQDN or IP address to contact for origination traffic.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The designated origination route port. Defaults to 5060.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The priority associated with the host, with 1 being the highest priority. Higher priority hosts are attempted first.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The protocol to use for the origination route. Encryption-enabled Amazon Chime Voice Connectors use TCP protocol by default.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The weight associated with the host. If hosts are equal in priority, calls are redistributed among them based on their relative weight.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RouteObservation struct {

	// The FQDN or IP address to contact for origination traffic.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The designated origination route port. Defaults to 5060.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The priority associated with the host, with 1 being the highest priority. Higher priority hosts are attempted first.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The protocol to use for the origination route. Encryption-enabled Amazon Chime Voice Connectors use TCP protocol by default.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The weight associated with the host. If hosts are equal in priority, calls are redistributed among them based on their relative weight.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RouteParameters struct {

	// The FQDN or IP address to contact for origination traffic.
	// +kubebuilder:validation:Optional
	Host *string `json:"host" tf:"host,omitempty"`

	// The designated origination route port. Defaults to 5060.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The priority associated with the host, with 1 being the highest priority. Higher priority hosts are attempted first.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// The protocol to use for the origination route. Encryption-enabled Amazon Chime Voice Connectors use TCP protocol by default.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The weight associated with the host. If hosts are equal in priority, calls are redistributed among them based on their relative weight.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

type VoiceConnectorOriginationInitParameters struct {

	// When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
	Route []RouteInitParameters `json:"route,omitempty" tf:"route,omitempty"`

	// The Amazon Chime Voice Connector ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/chime/v1beta1.VoiceConnector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`

	// Reference to a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDRef *v1.Reference `json:"voiceConnectorIdRef,omitempty" tf:"-"`

	// Selector for a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDSelector *v1.Selector `json:"voiceConnectorIdSelector,omitempty" tf:"-"`
}

type VoiceConnectorOriginationObservation struct {

	// When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The Amazon Chime Voice Connector ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
	Route []RouteObservation `json:"route,omitempty" tf:"route,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`
}

type VoiceConnectorOriginationParameters struct {

	// When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
	// +kubebuilder:validation:Optional
	Route []RouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// The Amazon Chime Voice Connector ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/chime/v1beta1.VoiceConnector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`

	// Reference to a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDRef *v1.Reference `json:"voiceConnectorIdRef,omitempty" tf:"-"`

	// Selector for a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDSelector *v1.Selector `json:"voiceConnectorIdSelector,omitempty" tf:"-"`
}

// VoiceConnectorOriginationSpec defines the desired state of VoiceConnectorOrigination
type VoiceConnectorOriginationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VoiceConnectorOriginationParameters `json:"forProvider"`
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
	InitProvider VoiceConnectorOriginationInitParameters `json:"initProvider,omitempty"`
}

// VoiceConnectorOriginationStatus defines the observed state of VoiceConnectorOrigination.
type VoiceConnectorOriginationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VoiceConnectorOriginationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VoiceConnectorOrigination is the Schema for the VoiceConnectorOriginations API. Enable origination settings to control inbound calling to your SIP infrastructure.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type VoiceConnectorOrigination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.route) || (has(self.initProvider) && has(self.initProvider.route))",message="spec.forProvider.route is a required parameter"
	Spec   VoiceConnectorOriginationSpec   `json:"spec"`
	Status VoiceConnectorOriginationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorOriginationList contains a list of VoiceConnectorOriginations
type VoiceConnectorOriginationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VoiceConnectorOrigination `json:"items"`
}

// Repository type metadata.
var (
	VoiceConnectorOrigination_Kind             = "VoiceConnectorOrigination"
	VoiceConnectorOrigination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VoiceConnectorOrigination_Kind}.String()
	VoiceConnectorOrigination_KindAPIVersion   = VoiceConnectorOrigination_Kind + "." + CRDGroupVersion.String()
	VoiceConnectorOrigination_GroupVersionKind = CRDGroupVersion.WithKind(VoiceConnectorOrigination_Kind)
)

func init() {
	SchemeBuilder.Register(&VoiceConnectorOrigination{}, &VoiceConnectorOriginationList{})
}
