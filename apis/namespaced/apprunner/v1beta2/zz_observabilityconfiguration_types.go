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

type ObservabilityConfigurationInitParameters struct {

	// Name of the observability configuration.
	ObservabilityConfigurationName *string `json:"observabilityConfigurationName,omitempty" tf:"observability_configuration_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
	TraceConfiguration *TraceConfigurationInitParameters `json:"traceConfiguration,omitempty" tf:"trace_configuration,omitempty"`
}

type ObservabilityConfigurationObservation struct {

	// ARN of this observability configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the observability configuration has the highest observability_configuration_revision among all configurations that share the same observability_configuration_name.
	Latest *bool `json:"latest,omitempty" tf:"latest,omitempty"`

	// Name of the observability configuration.
	ObservabilityConfigurationName *string `json:"observabilityConfigurationName,omitempty" tf:"observability_configuration_name,omitempty"`

	// The revision of this observability configuration.
	ObservabilityConfigurationRevision *float64 `json:"observabilityConfigurationRevision,omitempty" tf:"observability_configuration_revision,omitempty"`

	// Current state of the observability configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
	TraceConfiguration *TraceConfigurationObservation `json:"traceConfiguration,omitempty" tf:"trace_configuration,omitempty"`
}

type ObservabilityConfigurationParameters struct {

	// Name of the observability configuration.
	// +kubebuilder:validation:Optional
	ObservabilityConfigurationName *string `json:"observabilityConfigurationName,omitempty" tf:"observability_configuration_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
	// +kubebuilder:validation:Optional
	TraceConfiguration *TraceConfigurationParameters `json:"traceConfiguration,omitempty" tf:"trace_configuration,omitempty"`
}

type TraceConfigurationInitParameters struct {

	// Implementation provider chosen for tracing App Runner services. Valid values: AWSXRAY.
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type TraceConfigurationObservation struct {

	// Implementation provider chosen for tracing App Runner services. Valid values: AWSXRAY.
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type TraceConfigurationParameters struct {

	// Implementation provider chosen for tracing App Runner services. Valid values: AWSXRAY.
	// +kubebuilder:validation:Optional
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

// ObservabilityConfigurationSpec defines the desired state of ObservabilityConfiguration
type ObservabilityConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObservabilityConfigurationParameters `json:"forProvider"`
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
	InitProvider ObservabilityConfigurationInitParameters `json:"initProvider,omitempty"`
}

// ObservabilityConfigurationStatus defines the observed state of ObservabilityConfiguration.
type ObservabilityConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObservabilityConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ObservabilityConfiguration is the Schema for the ObservabilityConfigurations API. Manages an App Runner Observability Configuration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type ObservabilityConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.observabilityConfigurationName) || (has(self.initProvider) && has(self.initProvider.observabilityConfigurationName))",message="spec.forProvider.observabilityConfigurationName is a required parameter"
	Spec   ObservabilityConfigurationSpec   `json:"spec"`
	Status ObservabilityConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObservabilityConfigurationList contains a list of ObservabilityConfigurations
type ObservabilityConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObservabilityConfiguration `json:"items"`
}

// Repository type metadata.
var (
	ObservabilityConfiguration_Kind             = "ObservabilityConfiguration"
	ObservabilityConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObservabilityConfiguration_Kind}.String()
	ObservabilityConfiguration_KindAPIVersion   = ObservabilityConfiguration_Kind + "." + CRDGroupVersion.String()
	ObservabilityConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(ObservabilityConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&ObservabilityConfiguration{}, &ObservabilityConfigurationList{})
}
