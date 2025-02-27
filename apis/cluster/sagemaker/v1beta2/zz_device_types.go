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

type DeviceDeviceInitParameters struct {

	// A description for the device.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the device.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Amazon Web Services Internet of Things (IoT) object name.
	IotThingName *string `json:"iotThingName,omitempty" tf:"iot_thing_name,omitempty"`
}

type DeviceDeviceObservation struct {

	// A description for the device.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the device.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Amazon Web Services Internet of Things (IoT) object name.
	IotThingName *string `json:"iotThingName,omitempty" tf:"iot_thing_name,omitempty"`
}

type DeviceDeviceParameters struct {

	// A description for the device.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the device.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// Amazon Web Services Internet of Things (IoT) object name.
	// +kubebuilder:validation:Optional
	IotThingName *string `json:"iotThingName,omitempty" tf:"iot_thing_name,omitempty"`
}

type DeviceInitParameters struct {

	// The device to register with SageMaker Edge Manager. See Device details below.
	Device *DeviceDeviceInitParameters `json:"device,omitempty" tf:"device,omitempty"`

	// The name of the Device Fleet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/sagemaker/v1beta2.DeviceFleet
	DeviceFleetName *string `json:"deviceFleetName,omitempty" tf:"device_fleet_name,omitempty"`

	// Reference to a DeviceFleet in sagemaker to populate deviceFleetName.
	// +kubebuilder:validation:Optional
	DeviceFleetNameRef *v1.Reference `json:"deviceFleetNameRef,omitempty" tf:"-"`

	// Selector for a DeviceFleet in sagemaker to populate deviceFleetName.
	// +kubebuilder:validation:Optional
	DeviceFleetNameSelector *v1.Selector `json:"deviceFleetNameSelector,omitempty" tf:"-"`
}

type DeviceObservation struct {
	AgentVersion *string `json:"agentVersion,omitempty" tf:"agent_version,omitempty"`

	// The Amazon Resource Name (ARN) assigned by AWS to this Device.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The device to register with SageMaker Edge Manager. See Device details below.
	Device *DeviceDeviceObservation `json:"device,omitempty" tf:"device,omitempty"`

	// The name of the Device Fleet.
	DeviceFleetName *string `json:"deviceFleetName,omitempty" tf:"device_fleet_name,omitempty"`

	// The id is constructed from device-fleet-name/device-name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DeviceParameters struct {

	// The device to register with SageMaker Edge Manager. See Device details below.
	// +kubebuilder:validation:Optional
	Device *DeviceDeviceParameters `json:"device,omitempty" tf:"device,omitempty"`

	// The name of the Device Fleet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/sagemaker/v1beta2.DeviceFleet
	// +kubebuilder:validation:Optional
	DeviceFleetName *string `json:"deviceFleetName,omitempty" tf:"device_fleet_name,omitempty"`

	// Reference to a DeviceFleet in sagemaker to populate deviceFleetName.
	// +kubebuilder:validation:Optional
	DeviceFleetNameRef *v1.Reference `json:"deviceFleetNameRef,omitempty" tf:"-"`

	// Selector for a DeviceFleet in sagemaker to populate deviceFleetName.
	// +kubebuilder:validation:Optional
	DeviceFleetNameSelector *v1.Selector `json:"deviceFleetNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DeviceSpec defines the desired state of Device
type DeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceParameters `json:"forProvider"`
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
	InitProvider DeviceInitParameters `json:"initProvider,omitempty"`
}

// DeviceStatus defines the observed state of Device.
type DeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Device is the Schema for the Devices API. Provides a SageMaker Device resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Device struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.device) || (has(self.initProvider) && has(self.initProvider.device))",message="spec.forProvider.device is a required parameter"
	Spec   DeviceSpec   `json:"spec"`
	Status DeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceList contains a list of Devices
type DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Device `json:"items"`
}

// Repository type metadata.
var (
	Device_Kind             = "Device"
	Device_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Device_Kind}.String()
	Device_KindAPIVersion   = Device_Kind + "." + CRDGroupVersion.String()
	Device_GroupVersionKind = CRDGroupVersion.WithKind(Device_Kind)
)

func init() {
	SchemeBuilder.Register(&Device{}, &DeviceList{})
}
