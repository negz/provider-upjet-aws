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

type ListenerInitParameters struct {

	// The Amazon Resource Name (ARN) of your accelerator.
	// +crossplane:generate:reference:type=Accelerator
	AcceleratorArn *string `json:"acceleratorArn,omitempty" tf:"accelerator_arn,omitempty"`

	// Reference to a Accelerator to populate acceleratorArn.
	// +kubebuilder:validation:Optional
	AcceleratorArnRef *v1.Reference `json:"acceleratorArnRef,omitempty" tf:"-"`

	// Selector for a Accelerator to populate acceleratorArn.
	// +kubebuilder:validation:Optional
	AcceleratorArnSelector *v1.Selector `json:"acceleratorArnSelector,omitempty" tf:"-"`

	// Direct all requests from a user to the same endpoint. Valid values are NONE, SOURCE_IP. Default: NONE. If NONE, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If SOURCE_IP, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity *string `json:"clientAffinity,omitempty" tf:"client_affinity,omitempty"`

	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRange []PortRangeInitParameters `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The protocol for the connections from clients to the accelerator. Valid values are TCP, UDP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type ListenerObservation struct {

	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn *string `json:"acceleratorArn,omitempty" tf:"accelerator_arn,omitempty"`

	// Direct all requests from a user to the same endpoint. Valid values are NONE, SOURCE_IP. Default: NONE. If NONE, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If SOURCE_IP, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity *string `json:"clientAffinity,omitempty" tf:"client_affinity,omitempty"`

	// The Amazon Resource Name (ARN) of the listener.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRange []PortRangeObservation `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The protocol for the connections from clients to the accelerator. Valid values are TCP, UDP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type ListenerParameters struct {

	// The Amazon Resource Name (ARN) of your accelerator.
	// +crossplane:generate:reference:type=Accelerator
	// +kubebuilder:validation:Optional
	AcceleratorArn *string `json:"acceleratorArn,omitempty" tf:"accelerator_arn,omitempty"`

	// Reference to a Accelerator to populate acceleratorArn.
	// +kubebuilder:validation:Optional
	AcceleratorArnRef *v1.Reference `json:"acceleratorArnRef,omitempty" tf:"-"`

	// Selector for a Accelerator to populate acceleratorArn.
	// +kubebuilder:validation:Optional
	AcceleratorArnSelector *v1.Selector `json:"acceleratorArnSelector,omitempty" tf:"-"`

	// Direct all requests from a user to the same endpoint. Valid values are NONE, SOURCE_IP. Default: NONE. If NONE, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If SOURCE_IP, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	// +kubebuilder:validation:Optional
	ClientAffinity *string `json:"clientAffinity,omitempty" tf:"client_affinity,omitempty"`

	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	// +kubebuilder:validation:Optional
	PortRange []PortRangeParameters `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The protocol for the connections from clients to the accelerator. Valid values are TCP, UDP.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type PortRangeInitParameters struct {

	// The first port in the range of ports, inclusive.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The last port in the range of ports, inclusive.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type PortRangeObservation struct {

	// The first port in the range of ports, inclusive.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The last port in the range of ports, inclusive.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type PortRangeParameters struct {

	// The first port in the range of ports, inclusive.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The last port in the range of ports, inclusive.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

// ListenerSpec defines the desired state of Listener
type ListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListenerParameters `json:"forProvider"`
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
	InitProvider ListenerInitParameters `json:"initProvider,omitempty"`
}

// ListenerStatus defines the observed state of Listener.
type ListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Listener is the Schema for the Listeners API. Provides a Global Accelerator listener.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Listener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.portRange) || (has(self.initProvider) && has(self.initProvider.portRange))",message="spec.forProvider.portRange is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	Spec   ListenerSpec   `json:"spec"`
	Status ListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListenerList contains a list of Listeners
type ListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Listener `json:"items"`
}

// Repository type metadata.
var (
	Listener_Kind             = "Listener"
	Listener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Listener_Kind}.String()
	Listener_KindAPIVersion   = Listener_Kind + "." + CRDGroupVersion.String()
	Listener_GroupVersionKind = CRDGroupVersion.WithKind(Listener_Kind)
)

func init() {
	SchemeBuilder.Register(&Listener{}, &ListenerList{})
}
