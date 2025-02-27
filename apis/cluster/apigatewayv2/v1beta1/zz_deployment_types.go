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

type DeploymentInitParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/apigatewayv2/v1beta2.API
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// Description for the deployment resource. Must be less than or equal to 1024 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type DeploymentObservation struct {

	// API identifier.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Whether the deployment was automatically released.
	AutoDeployed *bool `json:"autoDeployed,omitempty" tf:"auto_deployed,omitempty"`

	// Description for the deployment resource. Must be less than or equal to 1024 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DeploymentParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/apigatewayv2/v1beta2.API
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// Description for the deployment resource. Must be less than or equal to 1024 characters in length.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentParameters `json:"forProvider"`
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
	InitProvider DeploymentInitParameters `json:"initProvider,omitempty"`
}

// DeploymentStatus defines the observed state of Deployment.
type DeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Deployment is the Schema for the Deployments API. Manages an Amazon API Gateway Version 2 deployment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSpec   `json:"spec"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Repository type metadata.
var (
	Deployment_Kind             = "Deployment"
	Deployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Deployment_Kind}.String()
	Deployment_KindAPIVersion   = Deployment_Kind + "." + CRDGroupVersion.String()
	Deployment_GroupVersionKind = CRDGroupVersion.WithKind(Deployment_Kind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
