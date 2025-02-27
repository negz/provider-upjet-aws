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

type InvocationInitParameters struct {

	// Name of the lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta2.Function
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Reference to a Function in lambda to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameRef *v1.Reference `json:"functionNameRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameSelector *v1.Selector `json:"functionNameSelector,omitempty" tf:"-"`

	// JSON payload to the lambda function.
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// Lifecycle scope of the resource to manage. Valid values are CREATE_ONLY and CRUD. Defaults to CREATE_ONLY. CREATE_ONLY will invoke the function only on creation or replacement. CRUD will invoke the function on each lifecycle event, and augment the input JSON payload with additional lifecycle information.
	LifecycleScope *string `json:"lifecycleScope,omitempty" tf:"lifecycle_scope,omitempty"`

	// Qualifier (i.e., version) of the lambda function. Defaults to $LATEST.
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// The JSON key used to store lifecycle information in the input JSON payload. Defaults to tf. This additional key is only included when lifecycle_scope is set to CRUD.
	TerraformKey *string `json:"terraformKey,omitempty" tf:"terraform_key,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger a re-invocation.
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type InvocationObservation struct {

	// Name of the lambda function.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// JSON payload to the lambda function.
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// Lifecycle scope of the resource to manage. Valid values are CREATE_ONLY and CRUD. Defaults to CREATE_ONLY. CREATE_ONLY will invoke the function only on creation or replacement. CRUD will invoke the function on each lifecycle event, and augment the input JSON payload with additional lifecycle information.
	LifecycleScope *string `json:"lifecycleScope,omitempty" tf:"lifecycle_scope,omitempty"`

	// Qualifier (i.e., version) of the lambda function. Defaults to $LATEST.
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// String result of the lambda function invocation.
	Result *string `json:"result,omitempty" tf:"result,omitempty"`

	// The JSON key used to store lifecycle information in the input JSON payload. Defaults to tf. This additional key is only included when lifecycle_scope is set to CRUD.
	TerraformKey *string `json:"terraformKey,omitempty" tf:"terraform_key,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger a re-invocation.
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type InvocationParameters struct {

	// Name of the lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta2.Function
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Reference to a Function in lambda to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameRef *v1.Reference `json:"functionNameRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameSelector *v1.Selector `json:"functionNameSelector,omitempty" tf:"-"`

	// JSON payload to the lambda function.
	// +kubebuilder:validation:Optional
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// Lifecycle scope of the resource to manage. Valid values are CREATE_ONLY and CRUD. Defaults to CREATE_ONLY. CREATE_ONLY will invoke the function only on creation or replacement. CRUD will invoke the function on each lifecycle event, and augment the input JSON payload with additional lifecycle information.
	// +kubebuilder:validation:Optional
	LifecycleScope *string `json:"lifecycleScope,omitempty" tf:"lifecycle_scope,omitempty"`

	// Qualifier (i.e., version) of the lambda function. Defaults to $LATEST.
	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The JSON key used to store lifecycle information in the input JSON payload. Defaults to tf. This additional key is only included when lifecycle_scope is set to CRUD.
	// +kubebuilder:validation:Optional
	TerraformKey *string `json:"terraformKey,omitempty" tf:"terraform_key,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger a re-invocation.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

// InvocationSpec defines the desired state of Invocation
type InvocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InvocationParameters `json:"forProvider"`
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
	InitProvider InvocationInitParameters `json:"initProvider,omitempty"`
}

// InvocationStatus defines the observed state of Invocation.
type InvocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InvocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Invocation is the Schema for the Invocations API. Invoke AWS Lambda Function
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Invocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.input) || (has(self.initProvider) && has(self.initProvider.input))",message="spec.forProvider.input is a required parameter"
	Spec   InvocationSpec   `json:"spec"`
	Status InvocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InvocationList contains a list of Invocations
type InvocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Invocation `json:"items"`
}

// Repository type metadata.
var (
	Invocation_Kind             = "Invocation"
	Invocation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Invocation_Kind}.String()
	Invocation_KindAPIVersion   = Invocation_Kind + "." + CRDGroupVersion.String()
	Invocation_GroupVersionKind = CRDGroupVersion.WithKind(Invocation_Kind)
)

func init() {
	SchemeBuilder.Register(&Invocation{}, &InvocationList{})
}
