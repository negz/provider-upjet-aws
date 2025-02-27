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

type ImageVersionInitParameters struct {

	// The registry path of the container image on which this image version is based.
	BaseImage *string `json:"baseImage,omitempty" tf:"base_image,omitempty"`

	// The name of the image. Must be unique to your account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.Image
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Reference to a Image in sagemaker to populate imageName.
	// +kubebuilder:validation:Optional
	ImageNameRef *v1.Reference `json:"imageNameRef,omitempty" tf:"-"`

	// Selector for a Image in sagemaker to populate imageName.
	// +kubebuilder:validation:Optional
	ImageNameSelector *v1.Selector `json:"imageNameSelector,omitempty" tf:"-"`
}

type ImageVersionObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Image Version.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The registry path of the container image on which this image version is based.
	BaseImage *string `json:"baseImage,omitempty" tf:"base_image,omitempty"`

	// The registry path of the container image that contains this image version.
	ContainerImage *string `json:"containerImage,omitempty" tf:"container_image,omitempty"`

	// The name of the Image.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of the image the version is based on.
	ImageArn *string `json:"imageArn,omitempty" tf:"image_arn,omitempty"`

	// The name of the image. Must be unique to your account.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ImageVersionParameters struct {

	// The registry path of the container image on which this image version is based.
	// +kubebuilder:validation:Optional
	BaseImage *string `json:"baseImage,omitempty" tf:"base_image,omitempty"`

	// The name of the image. Must be unique to your account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.Image
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Reference to a Image in sagemaker to populate imageName.
	// +kubebuilder:validation:Optional
	ImageNameRef *v1.Reference `json:"imageNameRef,omitempty" tf:"-"`

	// Selector for a Image in sagemaker to populate imageName.
	// +kubebuilder:validation:Optional
	ImageNameSelector *v1.Selector `json:"imageNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ImageVersionSpec defines the desired state of ImageVersion
type ImageVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageVersionParameters `json:"forProvider"`
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
	InitProvider ImageVersionInitParameters `json:"initProvider,omitempty"`
}

// ImageVersionStatus defines the observed state of ImageVersion.
type ImageVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImageVersion is the Schema for the ImageVersions API. Provides a SageMaker Image Version resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type ImageVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.baseImage) || (has(self.initProvider) && has(self.initProvider.baseImage))",message="spec.forProvider.baseImage is a required parameter"
	Spec   ImageVersionSpec   `json:"spec"`
	Status ImageVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageVersionList contains a list of ImageVersions
type ImageVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageVersion `json:"items"`
}

// Repository type metadata.
var (
	ImageVersion_Kind             = "ImageVersion"
	ImageVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageVersion_Kind}.String()
	ImageVersion_KindAPIVersion   = ImageVersion_Kind + "." + CRDGroupVersion.String()
	ImageVersion_GroupVersionKind = CRDGroupVersion.WithKind(ImageVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageVersion{}, &ImageVersionList{})
}
