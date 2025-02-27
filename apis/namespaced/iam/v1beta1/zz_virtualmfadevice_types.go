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

type VirtualMfaDeviceInitParameters struct {

	// –  The path for the virtual MFA device.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name of the virtual MFA device. Use with path to uniquely identify a virtual MFA device.
	VirtualMfaDeviceName *string `json:"virtualMfaDeviceName,omitempty" tf:"virtual_mfa_device_name,omitempty"`
}

type VirtualMfaDeviceObservation struct {

	// The Amazon Resource Name (ARN) specifying the virtual mfa device.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The base32 seed defined as specified in RFC3548. The base_32_string_seed is base64-encoded.
	Base32StringSeed *string `json:"base32StringSeed,omitempty" tf:"base_32_string_seed,omitempty"`

	// The date and time when the virtual MFA device was enabled.
	EnableDate *string `json:"enableDate,omitempty" tf:"enable_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  The path for the virtual MFA device.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// A QR code PNG image that encodes otpauth://totp/$virtualMFADeviceName@$AccountName?secret=$Base32String where $virtualMFADeviceName is one of the create call arguments. AccountName is the user name if set (otherwise, the account ID), and Base32String is the seed in base32 format.
	QrCodePng *string `json:"qrCodePng,omitempty" tf:"qr_code_png,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The associated IAM User name if the virtual MFA device is enabled.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// The name of the virtual MFA device. Use with path to uniquely identify a virtual MFA device.
	VirtualMfaDeviceName *string `json:"virtualMfaDeviceName,omitempty" tf:"virtual_mfa_device_name,omitempty"`
}

type VirtualMfaDeviceParameters struct {

	// –  The path for the virtual MFA device.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The name of the virtual MFA device. Use with path to uniquely identify a virtual MFA device.
	// +kubebuilder:validation:Optional
	VirtualMfaDeviceName *string `json:"virtualMfaDeviceName,omitempty" tf:"virtual_mfa_device_name,omitempty"`
}

// VirtualMfaDeviceSpec defines the desired state of VirtualMfaDevice
type VirtualMfaDeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualMfaDeviceParameters `json:"forProvider"`
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
	InitProvider VirtualMfaDeviceInitParameters `json:"initProvider,omitempty"`
}

// VirtualMfaDeviceStatus defines the observed state of VirtualMfaDevice.
type VirtualMfaDeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualMfaDeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VirtualMfaDevice is the Schema for the VirtualMfaDevices API. Provides an IAM Virtual MFA Device
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type VirtualMfaDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.virtualMfaDeviceName) || (has(self.initProvider) && has(self.initProvider.virtualMfaDeviceName))",message="spec.forProvider.virtualMfaDeviceName is a required parameter"
	Spec   VirtualMfaDeviceSpec   `json:"spec"`
	Status VirtualMfaDeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMfaDeviceList contains a list of VirtualMfaDevices
type VirtualMfaDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMfaDevice `json:"items"`
}

// Repository type metadata.
var (
	VirtualMfaDevice_Kind             = "VirtualMfaDevice"
	VirtualMfaDevice_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualMfaDevice_Kind}.String()
	VirtualMfaDevice_KindAPIVersion   = VirtualMfaDevice_Kind + "." + CRDGroupVersion.String()
	VirtualMfaDevice_GroupVersionKind = CRDGroupVersion.WithKind(VirtualMfaDevice_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualMfaDevice{}, &VirtualMfaDeviceList{})
}
