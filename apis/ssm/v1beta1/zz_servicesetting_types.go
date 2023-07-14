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

type ServiceSettingObservation struct {

	// ARN of the service setting.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the service setting.
	SettingID *string `json:"settingId,omitempty" tf:"setting_id,omitempty"`

	// Value of the service setting.
	SettingValue *string `json:"settingValue,omitempty" tf:"setting_value,omitempty"`

	// Status of the service setting. Value can be Default, Customized or PendingUpdate.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ServiceSettingParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the service setting.
	// +kubebuilder:validation:Optional
	SettingID *string `json:"settingId,omitempty" tf:"setting_id,omitempty"`

	// Value of the service setting.
	// +kubebuilder:validation:Optional
	SettingValue *string `json:"settingValue,omitempty" tf:"setting_value,omitempty"`
}

// ServiceSettingSpec defines the desired state of ServiceSetting
type ServiceSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceSettingParameters `json:"forProvider"`
}

// ServiceSettingStatus defines the observed state of ServiceSetting.
type ServiceSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceSetting is the Schema for the ServiceSettings API. Defines how a user interacts with or uses a service or a feature of a service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServiceSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.settingId)",message="settingId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.settingValue)",message="settingValue is a required parameter"
	Spec   ServiceSettingSpec   `json:"spec"`
	Status ServiceSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceSettingList contains a list of ServiceSettings
type ServiceSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceSetting `json:"items"`
}

// Repository type metadata.
var (
	ServiceSetting_Kind             = "ServiceSetting"
	ServiceSetting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceSetting_Kind}.String()
	ServiceSetting_KindAPIVersion   = ServiceSetting_Kind + "." + CRDGroupVersion.String()
	ServiceSetting_GroupVersionKind = CRDGroupVersion.WithKind(ServiceSetting_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceSetting{}, &ServiceSettingList{})
}
