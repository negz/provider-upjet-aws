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

type ResourceUrisInitParameters struct {

	// The type of the resource. can be one of JAR, FILE, and ARCHIVE.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// The URI for accessing the resource.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ResourceUrisObservation struct {

	// The type of the resource. can be one of JAR, FILE, and ARCHIVE.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// The URI for accessing the resource.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ResourceUrisParameters struct {

	// The type of the resource. can be one of JAR, FILE, and ARCHIVE.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`

	// The URI for accessing the resource.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type UserDefinedFunctionInitParameters struct {

	// The Java class that contains the function code.
	ClassName *string `json:"className,omitempty" tf:"class_name,omitempty"`

	// The owner of the function.
	OwnerName *string `json:"ownerName,omitempty" tf:"owner_name,omitempty"`

	// The owner type. can be one of USER, ROLE, and GROUP.
	OwnerType *string `json:"ownerType,omitempty" tf:"owner_type,omitempty"`

	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris []ResourceUrisInitParameters `json:"resourceUris,omitempty" tf:"resource_uris,omitempty"`
}

type UserDefinedFunctionObservation struct {

	// The ARN of the Glue User Defined Function.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// The Java class that contains the function code.
	ClassName *string `json:"className,omitempty" tf:"class_name,omitempty"`

	// The time at which the function was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The name of the Database to create the Function.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The id of the Glue User Defined Function.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The owner of the function.
	OwnerName *string `json:"ownerName,omitempty" tf:"owner_name,omitempty"`

	// The owner type. can be one of USER, ROLE, and GROUP.
	OwnerType *string `json:"ownerType,omitempty" tf:"owner_type,omitempty"`

	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris []ResourceUrisObservation `json:"resourceUris,omitempty" tf:"resource_uris,omitempty"`
}

type UserDefinedFunctionParameters struct {

	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// The Java class that contains the function code.
	// +kubebuilder:validation:Optional
	ClassName *string `json:"className,omitempty" tf:"class_name,omitempty"`

	// The name of the Database to create the Function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta2.CatalogDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a CatalogDatabase in glue to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a CatalogDatabase in glue to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// The owner of the function.
	// +kubebuilder:validation:Optional
	OwnerName *string `json:"ownerName,omitempty" tf:"owner_name,omitempty"`

	// The owner type. can be one of USER, ROLE, and GROUP.
	// +kubebuilder:validation:Optional
	OwnerType *string `json:"ownerType,omitempty" tf:"owner_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The configuration block for Resource URIs. See resource uris below for more details.
	// +kubebuilder:validation:Optional
	ResourceUris []ResourceUrisParameters `json:"resourceUris,omitempty" tf:"resource_uris,omitempty"`
}

// UserDefinedFunctionSpec defines the desired state of UserDefinedFunction
type UserDefinedFunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserDefinedFunctionParameters `json:"forProvider"`
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
	InitProvider UserDefinedFunctionInitParameters `json:"initProvider,omitempty"`
}

// UserDefinedFunctionStatus defines the observed state of UserDefinedFunction.
type UserDefinedFunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserDefinedFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserDefinedFunction is the Schema for the UserDefinedFunctions API. Provides a Glue User Defined Function.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type UserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.className) || (has(self.initProvider) && has(self.initProvider.className))",message="spec.forProvider.className is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ownerName) || (has(self.initProvider) && has(self.initProvider.ownerName))",message="spec.forProvider.ownerName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ownerType) || (has(self.initProvider) && has(self.initProvider.ownerType))",message="spec.forProvider.ownerType is a required parameter"
	Spec   UserDefinedFunctionSpec   `json:"spec"`
	Status UserDefinedFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserDefinedFunctionList contains a list of UserDefinedFunctions
type UserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserDefinedFunction `json:"items"`
}

// Repository type metadata.
var (
	UserDefinedFunction_Kind             = "UserDefinedFunction"
	UserDefinedFunction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserDefinedFunction_Kind}.String()
	UserDefinedFunction_KindAPIVersion   = UserDefinedFunction_Kind + "." + CRDGroupVersion.String()
	UserDefinedFunction_GroupVersionKind = CRDGroupVersion.WithKind(UserDefinedFunction_Kind)
)

func init() {
	SchemeBuilder.Register(&UserDefinedFunction{}, &UserDefinedFunctionList{})
}
