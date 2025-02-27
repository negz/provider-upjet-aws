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

type EBSSnapshotCopyInitParameters struct {

	// Specifies a completion duration to initiate a time-based snapshot copy. Time-based snapshot copy operations complete within the specified duration.  Value must be between 15 and 2880 minutes, in 15 minute increments only.
	CompletionDurationMinutes *float64 `json:"completionDurationMinutes,omitempty" tf:"completion_duration_minutes,omitempty"`

	// A description of what the snapshot is.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the snapshot is encrypted.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The ARN for the KMS encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore *bool `json:"permanentRestore,omitempty" tf:"permanent_restore,omitempty"`

	// The region of the source snapshot.
	SourceRegion *string `json:"sourceRegion,omitempty" tf:"source_region,omitempty"`

	// The ARN for the snapshot to be copied.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EBSSnapshot
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SourceSnapshotID *string `json:"sourceSnapshotId,omitempty" tf:"source_snapshot_id,omitempty"`

	// Reference to a EBSSnapshot in ec2 to populate sourceSnapshotId.
	// +kubebuilder:validation:Optional
	SourceSnapshotIDRef *v1.Reference `json:"sourceSnapshotIdRef,omitempty" tf:"-"`

	// Selector for a EBSSnapshot in ec2 to populate sourceSnapshotId.
	// +kubebuilder:validation:Optional
	SourceSnapshotIDSelector *v1.Selector `json:"sourceSnapshotIdSelector,omitempty" tf:"-"`

	// The name of the storage tier. Valid values are archive and standard. Default value is standard.
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays *float64 `json:"temporaryRestoreDays,omitempty" tf:"temporary_restore_days,omitempty"`
}

type EBSSnapshotCopyObservation struct {

	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies a completion duration to initiate a time-based snapshot copy. Time-based snapshot copy operations complete within the specified duration.  Value must be between 15 and 2880 minutes, in 15 minute increments only.
	CompletionDurationMinutes *float64 `json:"completionDurationMinutes,omitempty" tf:"completion_duration_minutes,omitempty"`

	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyID *string `json:"dataEncryptionKeyId,omitempty" tf:"data_encryption_key_id,omitempty"`

	// A description of what the snapshot is.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the snapshot is encrypted.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The snapshot ID (e.g., snap-59fcb34e).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for the KMS encryption key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Amazon Resource Name (ARN) of the EBS Snapshot.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// Value from an Amazon-maintained list (amazon, aws-marketplace, microsoft) of snapshot owners.
	OwnerAlias *string `json:"ownerAlias,omitempty" tf:"owner_alias,omitempty"`

	// The AWS account ID of the snapshot owner.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Indicates whether to permanently restore an archived snapshot.
	PermanentRestore *bool `json:"permanentRestore,omitempty" tf:"permanent_restore,omitempty"`

	// The region of the source snapshot.
	SourceRegion *string `json:"sourceRegion,omitempty" tf:"source_region,omitempty"`

	// The ARN for the snapshot to be copied.
	SourceSnapshotID *string `json:"sourceSnapshotId,omitempty" tf:"source_snapshot_id,omitempty"`

	// The name of the storage tier. Valid values are archive and standard. Default value is standard.
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	TemporaryRestoreDays *float64 `json:"temporaryRestoreDays,omitempty" tf:"temporary_restore_days,omitempty"`

	// The snapshot ID (e.g., snap-59fcb34e).
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// The size of the drive in GiBs.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
}

type EBSSnapshotCopyParameters struct {

	// Specifies a completion duration to initiate a time-based snapshot copy. Time-based snapshot copy operations complete within the specified duration.  Value must be between 15 and 2880 minutes, in 15 minute increments only.
	// +kubebuilder:validation:Optional
	CompletionDurationMinutes *float64 `json:"completionDurationMinutes,omitempty" tf:"completion_duration_minutes,omitempty"`

	// A description of what the snapshot is.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the snapshot is encrypted.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The ARN for the KMS encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Indicates whether to permanently restore an archived snapshot.
	// +kubebuilder:validation:Optional
	PermanentRestore *bool `json:"permanentRestore,omitempty" tf:"permanent_restore,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The region of the source snapshot.
	// +kubebuilder:validation:Optional
	SourceRegion *string `json:"sourceRegion,omitempty" tf:"source_region,omitempty"`

	// The ARN for the snapshot to be copied.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EBSSnapshot
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceSnapshotID *string `json:"sourceSnapshotId,omitempty" tf:"source_snapshot_id,omitempty"`

	// Reference to a EBSSnapshot in ec2 to populate sourceSnapshotId.
	// +kubebuilder:validation:Optional
	SourceSnapshotIDRef *v1.Reference `json:"sourceSnapshotIdRef,omitempty" tf:"-"`

	// Selector for a EBSSnapshot in ec2 to populate sourceSnapshotId.
	// +kubebuilder:validation:Optional
	SourceSnapshotIDSelector *v1.Selector `json:"sourceSnapshotIdSelector,omitempty" tf:"-"`

	// The name of the storage tier. Valid values are archive and standard. Default value is standard.
	// +kubebuilder:validation:Optional
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	// +kubebuilder:validation:Optional
	TemporaryRestoreDays *float64 `json:"temporaryRestoreDays,omitempty" tf:"temporary_restore_days,omitempty"`
}

// EBSSnapshotCopySpec defines the desired state of EBSSnapshotCopy
type EBSSnapshotCopySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EBSSnapshotCopyParameters `json:"forProvider"`
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
	InitProvider EBSSnapshotCopyInitParameters `json:"initProvider,omitempty"`
}

// EBSSnapshotCopyStatus defines the observed state of EBSSnapshotCopy.
type EBSSnapshotCopyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EBSSnapshotCopyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EBSSnapshotCopy is the Schema for the EBSSnapshotCopys API. Duplicates an existing Amazon snapshot
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type EBSSnapshotCopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceRegion) || (has(self.initProvider) && has(self.initProvider.sourceRegion))",message="spec.forProvider.sourceRegion is a required parameter"
	Spec   EBSSnapshotCopySpec   `json:"spec"`
	Status EBSSnapshotCopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EBSSnapshotCopyList contains a list of EBSSnapshotCopys
type EBSSnapshotCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EBSSnapshotCopy `json:"items"`
}

// Repository type metadata.
var (
	EBSSnapshotCopy_Kind             = "EBSSnapshotCopy"
	EBSSnapshotCopy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EBSSnapshotCopy_Kind}.String()
	EBSSnapshotCopy_KindAPIVersion   = EBSSnapshotCopy_Kind + "." + CRDGroupVersion.String()
	EBSSnapshotCopy_GroupVersionKind = CRDGroupVersion.WithKind(EBSSnapshotCopy_Kind)
)

func init() {
	SchemeBuilder.Register(&EBSSnapshotCopy{}, &EBSSnapshotCopyList{})
}
