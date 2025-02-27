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

type RoutingControlInitParameters struct {

	// ARN of the cluster in which this routing control will reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/route53recoverycontrolconfig/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	ClusterArn *string `json:"clusterArn,omitempty" tf:"cluster_arn,omitempty"`

	// Reference to a Cluster in route53recoverycontrolconfig to populate clusterArn.
	// +kubebuilder:validation:Optional
	ClusterArnRef *v1.Reference `json:"clusterArnRef,omitempty" tf:"-"`

	// Selector for a Cluster in route53recoverycontrolconfig to populate clusterArn.
	// +kubebuilder:validation:Optional
	ClusterArnSelector *v1.Selector `json:"clusterArnSelector,omitempty" tf:"-"`

	// ARN of the control panel in which this routing control will reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/route53recoverycontrolconfig/v1beta1.ControlPanel
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	ControlPanelArn *string `json:"controlPanelArn,omitempty" tf:"control_panel_arn,omitempty"`

	// Reference to a ControlPanel in route53recoverycontrolconfig to populate controlPanelArn.
	// +kubebuilder:validation:Optional
	ControlPanelArnRef *v1.Reference `json:"controlPanelArnRef,omitempty" tf:"-"`

	// Selector for a ControlPanel in route53recoverycontrolconfig to populate controlPanelArn.
	// +kubebuilder:validation:Optional
	ControlPanelArnSelector *v1.Selector `json:"controlPanelArnSelector,omitempty" tf:"-"`

	// The name describing the routing control.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RoutingControlObservation struct {

	// ARN of the routing control.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the cluster in which this routing control will reside.
	ClusterArn *string `json:"clusterArn,omitempty" tf:"cluster_arn,omitempty"`

	// ARN of the control panel in which this routing control will reside.
	ControlPanelArn *string `json:"controlPanelArn,omitempty" tf:"control_panel_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name describing the routing control.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Status of routing control. PENDING when it is being created/updated, PENDING_DELETION when it is being deleted, and DEPLOYED otherwise.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RoutingControlParameters struct {

	// ARN of the cluster in which this routing control will reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/route53recoverycontrolconfig/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	ClusterArn *string `json:"clusterArn,omitempty" tf:"cluster_arn,omitempty"`

	// Reference to a Cluster in route53recoverycontrolconfig to populate clusterArn.
	// +kubebuilder:validation:Optional
	ClusterArnRef *v1.Reference `json:"clusterArnRef,omitempty" tf:"-"`

	// Selector for a Cluster in route53recoverycontrolconfig to populate clusterArn.
	// +kubebuilder:validation:Optional
	ClusterArnSelector *v1.Selector `json:"clusterArnSelector,omitempty" tf:"-"`

	// ARN of the control panel in which this routing control will reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/route53recoverycontrolconfig/v1beta1.ControlPanel
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	ControlPanelArn *string `json:"controlPanelArn,omitempty" tf:"control_panel_arn,omitempty"`

	// Reference to a ControlPanel in route53recoverycontrolconfig to populate controlPanelArn.
	// +kubebuilder:validation:Optional
	ControlPanelArnRef *v1.Reference `json:"controlPanelArnRef,omitempty" tf:"-"`

	// Selector for a ControlPanel in route53recoverycontrolconfig to populate controlPanelArn.
	// +kubebuilder:validation:Optional
	ControlPanelArnSelector *v1.Selector `json:"controlPanelArnSelector,omitempty" tf:"-"`

	// The name describing the routing control.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// RoutingControlSpec defines the desired state of RoutingControl
type RoutingControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoutingControlParameters `json:"forProvider"`
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
	InitProvider RoutingControlInitParameters `json:"initProvider,omitempty"`
}

// RoutingControlStatus defines the observed state of RoutingControl.
type RoutingControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoutingControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RoutingControl is the Schema for the RoutingControls API. Provides an AWS Route 53 Recovery Control Config Routing Control
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type RoutingControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RoutingControlSpec   `json:"spec"`
	Status RoutingControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingControlList contains a list of RoutingControls
type RoutingControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoutingControl `json:"items"`
}

// Repository type metadata.
var (
	RoutingControl_Kind             = "RoutingControl"
	RoutingControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoutingControl_Kind}.String()
	RoutingControl_KindAPIVersion   = RoutingControl_Kind + "." + CRDGroupVersion.String()
	RoutingControl_GroupVersionKind = CRDGroupVersion.WithKind(RoutingControl_Kind)
)

func init() {
	SchemeBuilder.Register(&RoutingControl{}, &RoutingControlList{})
}
