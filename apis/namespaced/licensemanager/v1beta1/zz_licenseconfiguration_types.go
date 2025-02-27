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

type LicenseConfigurationInitParameters struct {

	// Description of the license configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Number of licenses managed by the license configuration.
	LicenseCount *float64 `json:"licenseCount,omitempty" tf:"license_count,omitempty"`

	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool `json:"licenseCountHardLimit,omitempty" tf:"license_count_hard_limit,omitempty"`

	// Dimension to use to track license inventory. Specify either vCPU, Instance, Core or Socket.
	LicenseCountingType *string `json:"licenseCountingType,omitempty" tf:"license_counting_type,omitempty"`

	// Array of configured License Manager rules.
	LicenseRules []*string `json:"licenseRules,omitempty" tf:"license_rules,omitempty"`

	// Name of the license configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LicenseConfigurationObservation struct {

	// The license configuration ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the license configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The license configuration ARN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Number of licenses managed by the license configuration.
	LicenseCount *float64 `json:"licenseCount,omitempty" tf:"license_count,omitempty"`

	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool `json:"licenseCountHardLimit,omitempty" tf:"license_count_hard_limit,omitempty"`

	// Dimension to use to track license inventory. Specify either vCPU, Instance, Core or Socket.
	LicenseCountingType *string `json:"licenseCountingType,omitempty" tf:"license_counting_type,omitempty"`

	// Array of configured License Manager rules.
	LicenseRules []*string `json:"licenseRules,omitempty" tf:"license_rules,omitempty"`

	// Name of the license configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Account ID of the owner of the license configuration.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LicenseConfigurationParameters struct {

	// Description of the license configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Number of licenses managed by the license configuration.
	// +kubebuilder:validation:Optional
	LicenseCount *float64 `json:"licenseCount,omitempty" tf:"license_count,omitempty"`

	// Sets the number of available licenses as a hard limit.
	// +kubebuilder:validation:Optional
	LicenseCountHardLimit *bool `json:"licenseCountHardLimit,omitempty" tf:"license_count_hard_limit,omitempty"`

	// Dimension to use to track license inventory. Specify either vCPU, Instance, Core or Socket.
	// +kubebuilder:validation:Optional
	LicenseCountingType *string `json:"licenseCountingType,omitempty" tf:"license_counting_type,omitempty"`

	// Array of configured License Manager rules.
	// +kubebuilder:validation:Optional
	LicenseRules []*string `json:"licenseRules,omitempty" tf:"license_rules,omitempty"`

	// Name of the license configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LicenseConfigurationSpec defines the desired state of LicenseConfiguration
type LicenseConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LicenseConfigurationParameters `json:"forProvider"`
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
	InitProvider LicenseConfigurationInitParameters `json:"initProvider,omitempty"`
}

// LicenseConfigurationStatus defines the observed state of LicenseConfiguration.
type LicenseConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LicenseConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LicenseConfiguration is the Schema for the LicenseConfigurations API. Provides a License Manager license configuration resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type LicenseConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.licenseCountingType) || (has(self.initProvider) && has(self.initProvider.licenseCountingType))",message="spec.forProvider.licenseCountingType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   LicenseConfigurationSpec   `json:"spec"`
	Status LicenseConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseConfigurationList contains a list of LicenseConfigurations
type LicenseConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseConfiguration `json:"items"`
}

// Repository type metadata.
var (
	LicenseConfiguration_Kind             = "LicenseConfiguration"
	LicenseConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LicenseConfiguration_Kind}.String()
	LicenseConfiguration_KindAPIVersion   = LicenseConfiguration_Kind + "." + CRDGroupVersion.String()
	LicenseConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(LicenseConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&LicenseConfiguration{}, &LicenseConfigurationList{})
}
