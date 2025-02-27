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

type LedgerInitParameters struct {

	// The deletion protection for the QLDB Ledger instance. By default it is true.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The key in AWS Key Management Service (AWS KMS) to use for encryption of data at rest in the ledger. For more information, see the AWS documentation. Valid values are "AWS_OWNED_KMS_KEY" to use an AWS KMS key that is owned and managed by AWS on your behalf, or the ARN of a valid symmetric customer managed KMS key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kms/v1beta1.Key
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Reference to a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`

	// The permissions mode for the QLDB ledger instance. Specify either ALLOW_ALL or STANDARD.
	PermissionsMode *string `json:"permissionsMode,omitempty" tf:"permissions_mode,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LedgerObservation struct {

	// The ARN of the QLDB Ledger
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The deletion protection for the QLDB Ledger instance. By default it is true.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The Name of the QLDB Ledger
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The key in AWS Key Management Service (AWS KMS) to use for encryption of data at rest in the ledger. For more information, see the AWS documentation. Valid values are "AWS_OWNED_KMS_KEY" to use an AWS KMS key that is owned and managed by AWS on your behalf, or the ARN of a valid symmetric customer managed KMS key.
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// The permissions mode for the QLDB ledger instance. Specify either ALLOW_ALL or STANDARD.
	PermissionsMode *string `json:"permissionsMode,omitempty" tf:"permissions_mode,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LedgerParameters struct {

	// The deletion protection for the QLDB Ledger instance. By default it is true.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The key in AWS Key Management Service (AWS KMS) to use for encryption of data at rest in the ledger. For more information, see the AWS documentation. Valid values are "AWS_OWNED_KMS_KEY" to use an AWS KMS key that is owned and managed by AWS on your behalf, or the ARN of a valid symmetric customer managed KMS key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Reference to a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`

	// The permissions mode for the QLDB ledger instance. Specify either ALLOW_ALL or STANDARD.
	// +kubebuilder:validation:Optional
	PermissionsMode *string `json:"permissionsMode,omitempty" tf:"permissions_mode,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LedgerSpec defines the desired state of Ledger
type LedgerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LedgerParameters `json:"forProvider"`
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
	InitProvider LedgerInitParameters `json:"initProvider,omitempty"`
}

// LedgerStatus defines the observed state of Ledger.
type LedgerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LedgerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Ledger is the Schema for the Ledgers API. Provides an QLDB Resource resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Ledger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissionsMode) || (has(self.initProvider) && has(self.initProvider.permissionsMode))",message="spec.forProvider.permissionsMode is a required parameter"
	Spec   LedgerSpec   `json:"spec"`
	Status LedgerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LedgerList contains a list of Ledgers
type LedgerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ledger `json:"items"`
}

// Repository type metadata.
var (
	Ledger_Kind             = "Ledger"
	Ledger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ledger_Kind}.String()
	Ledger_KindAPIVersion   = Ledger_Kind + "." + CRDGroupVersion.String()
	Ledger_GroupVersionKind = CRDGroupVersion.WithKind(Ledger_Kind)
)

func init() {
	SchemeBuilder.Register(&Ledger{}, &LedgerList{})
}
