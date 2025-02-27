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

type SnapshotInitParameters struct {

	// The namespace to create a snapshot for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshiftserverless/v1beta1.Workgroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("namespace_name",false)
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a Workgroup in redshiftserverless to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a Workgroup in redshiftserverless to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// How long to retain the created snapshot. Default value is -1.
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type SnapshotObservation struct {

	// All of the Amazon Web Services accounts that have access to restore a snapshot to a provisioned cluster.
	// +listType=set
	AccountsWithProvisionedRestoreAccess []*string `json:"accountsWithProvisionedRestoreAccess,omitempty" tf:"accounts_with_provisioned_restore_access,omitempty"`

	// All of the Amazon Web Services accounts that have access to restore a snapshot to a namespace.
	// +listType=set
	AccountsWithRestoreAccess []*string `json:"accountsWithRestoreAccess,omitempty" tf:"accounts_with_restore_access,omitempty"`

	// The username of the database within a snapshot.
	AdminUsername *string `json:"adminUsername,omitempty" tf:"admin_username,omitempty"`

	// The Amazon Resource Name (ARN) of the snapshot.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier of the KMS key used to encrypt the snapshot.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The Amazon Resource Name (ARN) of the namespace the snapshot was created from.
	NamespaceArn *string `json:"namespaceArn,omitempty" tf:"namespace_arn,omitempty"`

	// The namespace to create a snapshot for.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// The owner Amazon Web Services; account of the snapshot.
	OwnerAccount *string `json:"ownerAccount,omitempty" tf:"owner_account,omitempty"`

	// How long to retain the created snapshot. Default value is -1.
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type SnapshotParameters struct {

	// The namespace to create a snapshot for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshiftserverless/v1beta1.Workgroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("namespace_name",false)
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a Workgroup in redshiftserverless to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a Workgroup in redshiftserverless to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// How long to retain the created snapshot. Default value is -1.
	// +kubebuilder:validation:Optional
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
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
	InitProvider SnapshotInitParameters `json:"initProvider,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Snapshot is the Schema for the Snapshots API. Provides a Redshift Serverless Snapshot resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
