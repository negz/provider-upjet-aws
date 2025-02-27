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

type RegistryScanningConfigurationInitParameters struct {

	// One or multiple blocks specifying scanning rules to determine which repository filters are used and at what frequency scanning will occur. See below for schema.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// the scanning type to set for the registry. Can be either ENHANCED or BASIC.
	ScanType *string `json:"scanType,omitempty" tf:"scan_type,omitempty"`
}

type RegistryScanningConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The registry ID the scanning configuration applies to.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// One or multiple blocks specifying scanning rules to determine which repository filters are used and at what frequency scanning will occur. See below for schema.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// the scanning type to set for the registry. Can be either ENHANCED or BASIC.
	ScanType *string `json:"scanType,omitempty" tf:"scan_type,omitempty"`
}

type RegistryScanningConfigurationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// One or multiple blocks specifying scanning rules to determine which repository filters are used and at what frequency scanning will occur. See below for schema.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// the scanning type to set for the registry. Can be either ENHANCED or BASIC.
	// +kubebuilder:validation:Optional
	ScanType *string `json:"scanType,omitempty" tf:"scan_type,omitempty"`
}

type RepositoryFilterInitParameters struct {
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`
}

type RepositoryFilterObservation struct {
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`
}

type RepositoryFilterParameters struct {

	// +kubebuilder:validation:Optional
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	FilterType *string `json:"filterType" tf:"filter_type,omitempty"`
}

type RuleInitParameters struct {

	// One or more repository filter blocks, containing a filter  and a filter_type .
	RepositoryFilter []RepositoryFilterInitParameters `json:"repositoryFilter,omitempty" tf:"repository_filter,omitempty"`

	// The frequency that scans are performed at for a private registry. Can be SCAN_ON_PUSH, CONTINUOUS_SCAN, or MANUAL.
	ScanFrequency *string `json:"scanFrequency,omitempty" tf:"scan_frequency,omitempty"`
}

type RuleObservation struct {

	// One or more repository filter blocks, containing a filter  and a filter_type .
	RepositoryFilter []RepositoryFilterObservation `json:"repositoryFilter,omitempty" tf:"repository_filter,omitempty"`

	// The frequency that scans are performed at for a private registry. Can be SCAN_ON_PUSH, CONTINUOUS_SCAN, or MANUAL.
	ScanFrequency *string `json:"scanFrequency,omitempty" tf:"scan_frequency,omitempty"`
}

type RuleParameters struct {

	// One or more repository filter blocks, containing a filter  and a filter_type .
	// +kubebuilder:validation:Optional
	RepositoryFilter []RepositoryFilterParameters `json:"repositoryFilter" tf:"repository_filter,omitempty"`

	// The frequency that scans are performed at for a private registry. Can be SCAN_ON_PUSH, CONTINUOUS_SCAN, or MANUAL.
	// +kubebuilder:validation:Optional
	ScanFrequency *string `json:"scanFrequency" tf:"scan_frequency,omitempty"`
}

// RegistryScanningConfigurationSpec defines the desired state of RegistryScanningConfiguration
type RegistryScanningConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryScanningConfigurationParameters `json:"forProvider"`
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
	InitProvider RegistryScanningConfigurationInitParameters `json:"initProvider,omitempty"`
}

// RegistryScanningConfigurationStatus defines the observed state of RegistryScanningConfiguration.
type RegistryScanningConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryScanningConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegistryScanningConfiguration is the Schema for the RegistryScanningConfigurations API. Provides an Elastic Container Registry Scanning Configuration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type RegistryScanningConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scanType) || (has(self.initProvider) && has(self.initProvider.scanType))",message="spec.forProvider.scanType is a required parameter"
	Spec   RegistryScanningConfigurationSpec   `json:"spec"`
	Status RegistryScanningConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryScanningConfigurationList contains a list of RegistryScanningConfigurations
type RegistryScanningConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryScanningConfiguration `json:"items"`
}

// Repository type metadata.
var (
	RegistryScanningConfiguration_Kind             = "RegistryScanningConfiguration"
	RegistryScanningConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegistryScanningConfiguration_Kind}.String()
	RegistryScanningConfiguration_KindAPIVersion   = RegistryScanningConfiguration_Kind + "." + CRDGroupVersion.String()
	RegistryScanningConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(RegistryScanningConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&RegistryScanningConfiguration{}, &RegistryScanningConfigurationList{})
}
