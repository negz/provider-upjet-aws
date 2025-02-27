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

type StudioLifecycleConfigInitParameters struct {

	// The App type that the Lifecycle Configuration is attached to. Valid values are JupyterServer, JupyterLab, CodeEditor and KernelGateway.
	StudioLifecycleConfigAppType *string `json:"studioLifecycleConfigAppType,omitempty" tf:"studio_lifecycle_config_app_type,omitempty"`

	// The content of your Studio Lifecycle Configuration script. This content must be base64 encoded.
	StudioLifecycleConfigContent *string `json:"studioLifecycleConfigContent,omitempty" tf:"studio_lifecycle_config_content,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StudioLifecycleConfigObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Studio Lifecycle Config.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the Studio Lifecycle Config.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The App type that the Lifecycle Configuration is attached to. Valid values are JupyterServer, JupyterLab, CodeEditor and KernelGateway.
	StudioLifecycleConfigAppType *string `json:"studioLifecycleConfigAppType,omitempty" tf:"studio_lifecycle_config_app_type,omitempty"`

	// The content of your Studio Lifecycle Configuration script. This content must be base64 encoded.
	StudioLifecycleConfigContent *string `json:"studioLifecycleConfigContent,omitempty" tf:"studio_lifecycle_config_content,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StudioLifecycleConfigParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The App type that the Lifecycle Configuration is attached to. Valid values are JupyterServer, JupyterLab, CodeEditor and KernelGateway.
	// +kubebuilder:validation:Optional
	StudioLifecycleConfigAppType *string `json:"studioLifecycleConfigAppType,omitempty" tf:"studio_lifecycle_config_app_type,omitempty"`

	// The content of your Studio Lifecycle Configuration script. This content must be base64 encoded.
	// +kubebuilder:validation:Optional
	StudioLifecycleConfigContent *string `json:"studioLifecycleConfigContent,omitempty" tf:"studio_lifecycle_config_content,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StudioLifecycleConfigSpec defines the desired state of StudioLifecycleConfig
type StudioLifecycleConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StudioLifecycleConfigParameters `json:"forProvider"`
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
	InitProvider StudioLifecycleConfigInitParameters `json:"initProvider,omitempty"`
}

// StudioLifecycleConfigStatus defines the observed state of StudioLifecycleConfig.
type StudioLifecycleConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StudioLifecycleConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StudioLifecycleConfig is the Schema for the StudioLifecycleConfigs API. Provides a SageMaker Studio Lifecycle Config resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type StudioLifecycleConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.studioLifecycleConfigAppType) || (has(self.initProvider) && has(self.initProvider.studioLifecycleConfigAppType))",message="spec.forProvider.studioLifecycleConfigAppType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.studioLifecycleConfigContent) || (has(self.initProvider) && has(self.initProvider.studioLifecycleConfigContent))",message="spec.forProvider.studioLifecycleConfigContent is a required parameter"
	Spec   StudioLifecycleConfigSpec   `json:"spec"`
	Status StudioLifecycleConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StudioLifecycleConfigList contains a list of StudioLifecycleConfigs
type StudioLifecycleConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StudioLifecycleConfig `json:"items"`
}

// Repository type metadata.
var (
	StudioLifecycleConfig_Kind             = "StudioLifecycleConfig"
	StudioLifecycleConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StudioLifecycleConfig_Kind}.String()
	StudioLifecycleConfig_KindAPIVersion   = StudioLifecycleConfig_Kind + "." + CRDGroupVersion.String()
	StudioLifecycleConfig_GroupVersionKind = CRDGroupVersion.WithKind(StudioLifecycleConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&StudioLifecycleConfig{}, &StudioLifecycleConfigList{})
}
