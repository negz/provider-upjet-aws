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

type UploadInitParameters struct {

	// The upload's content type (for example, application/octet-stream).
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the project for the upload.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/devicefarm/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ProjectArn *string `json:"projectArn,omitempty" tf:"project_arn,omitempty"`

	// Reference to a Project in devicefarm to populate projectArn.
	// +kubebuilder:validation:Optional
	ProjectArnRef *v1.Reference `json:"projectArnRef,omitempty" tf:"-"`

	// Selector for a Project in devicefarm to populate projectArn.
	// +kubebuilder:validation:Optional
	ProjectArnSelector *v1.Selector `json:"projectArnSelector,omitempty" tf:"-"`

	// The upload's upload type. See AWS Docs for valid list of values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type UploadObservation struct {

	// The Amazon Resource Name of this upload.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The upload's category.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The upload's content type (for example, application/octet-stream).
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The upload's metadata. For example, for Android, this contains information that is parsed from the manifest and is displayed in the AWS Device Farm console after the associated app is uploaded.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the project for the upload.
	ProjectArn *string `json:"projectArn,omitempty" tf:"project_arn,omitempty"`

	// The upload's upload type. See AWS Docs for valid list of values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The presigned Amazon S3 URL that was used to store a file using a PUT request.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type UploadParameters struct {

	// The upload's content type (for example, application/octet-stream).
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The upload's file name. The name should not contain any forward slashes (/). If you are uploading an iOS app, the file name must end with the .ipa extension. If you are uploading an Android app, the file name must end with the .apk extension. For all others, the file name must end with the .zip file extension.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the project for the upload.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/devicefarm/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ProjectArn *string `json:"projectArn,omitempty" tf:"project_arn,omitempty"`

	// Reference to a Project in devicefarm to populate projectArn.
	// +kubebuilder:validation:Optional
	ProjectArnRef *v1.Reference `json:"projectArnRef,omitempty" tf:"-"`

	// Selector for a Project in devicefarm to populate projectArn.
	// +kubebuilder:validation:Optional
	ProjectArnSelector *v1.Selector `json:"projectArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The upload's upload type. See AWS Docs for valid list of values.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// UploadSpec defines the desired state of Upload
type UploadSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UploadParameters `json:"forProvider"`
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
	InitProvider UploadInitParameters `json:"initProvider,omitempty"`
}

// UploadStatus defines the observed state of Upload.
type UploadStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UploadObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Upload is the Schema for the Uploads API. Provides a Devicefarm upload
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Upload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   UploadSpec   `json:"spec"`
	Status UploadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UploadList contains a list of Uploads
type UploadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Upload `json:"items"`
}

// Repository type metadata.
var (
	Upload_Kind             = "Upload"
	Upload_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Upload_Kind}.String()
	Upload_KindAPIVersion   = Upload_Kind + "." + CRDGroupVersion.String()
	Upload_GroupVersionKind = CRDGroupVersion.WithKind(Upload_Kind)
)

func init() {
	SchemeBuilder.Register(&Upload{}, &UploadList{})
}
