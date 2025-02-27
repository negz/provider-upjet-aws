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

type CompositePartitionKeyInitParameters struct {

	// The level of enforcement for the specification of a dimension key in ingested records. Valid values: REQUIRED, OPTIONAL.
	EnforcementInRecord *string `json:"enforcementInRecord,omitempty" tf:"enforcement_in_record,omitempty"`

	// The name of the attribute used for a dimension key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of the partition key. Valid values: DIMENSION, MEASURE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CompositePartitionKeyObservation struct {

	// The level of enforcement for the specification of a dimension key in ingested records. Valid values: REQUIRED, OPTIONAL.
	EnforcementInRecord *string `json:"enforcementInRecord,omitempty" tf:"enforcement_in_record,omitempty"`

	// The name of the attribute used for a dimension key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of the partition key. Valid values: DIMENSION, MEASURE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CompositePartitionKeyParameters struct {

	// The level of enforcement for the specification of a dimension key in ingested records. Valid values: REQUIRED, OPTIONAL.
	// +kubebuilder:validation:Optional
	EnforcementInRecord *string `json:"enforcementInRecord,omitempty" tf:"enforcement_in_record,omitempty"`

	// The name of the attribute used for a dimension key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of the partition key. Valid values: DIMENSION, MEASURE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type MagneticStoreRejectedDataLocationInitParameters struct {

	// Configuration of an S3 location to write error reports for records rejected, asynchronously, during magnetic store writes. See S3 Configuration below for more details.
	S3Configuration *S3ConfigurationInitParameters `json:"s3Configuration,omitempty" tf:"s3_configuration,omitempty"`
}

type MagneticStoreRejectedDataLocationObservation struct {

	// Configuration of an S3 location to write error reports for records rejected, asynchronously, during magnetic store writes. See S3 Configuration below for more details.
	S3Configuration *S3ConfigurationObservation `json:"s3Configuration,omitempty" tf:"s3_configuration,omitempty"`
}

type MagneticStoreRejectedDataLocationParameters struct {

	// Configuration of an S3 location to write error reports for records rejected, asynchronously, during magnetic store writes. See S3 Configuration below for more details.
	// +kubebuilder:validation:Optional
	S3Configuration *S3ConfigurationParameters `json:"s3Configuration,omitempty" tf:"s3_configuration,omitempty"`
}

type MagneticStoreWritePropertiesInitParameters struct {

	// A flag to enable magnetic store writes.
	EnableMagneticStoreWrites *bool `json:"enableMagneticStoreWrites,omitempty" tf:"enable_magnetic_store_writes,omitempty"`

	// The location to write error reports for records rejected asynchronously during magnetic store writes. See Magnetic Store Rejected Data Location below for more details.
	MagneticStoreRejectedDataLocation *MagneticStoreRejectedDataLocationInitParameters `json:"magneticStoreRejectedDataLocation,omitempty" tf:"magnetic_store_rejected_data_location,omitempty"`
}

type MagneticStoreWritePropertiesObservation struct {

	// A flag to enable magnetic store writes.
	EnableMagneticStoreWrites *bool `json:"enableMagneticStoreWrites,omitempty" tf:"enable_magnetic_store_writes,omitempty"`

	// The location to write error reports for records rejected asynchronously during magnetic store writes. See Magnetic Store Rejected Data Location below for more details.
	MagneticStoreRejectedDataLocation *MagneticStoreRejectedDataLocationObservation `json:"magneticStoreRejectedDataLocation,omitempty" tf:"magnetic_store_rejected_data_location,omitempty"`
}

type MagneticStoreWritePropertiesParameters struct {

	// A flag to enable magnetic store writes.
	// +kubebuilder:validation:Optional
	EnableMagneticStoreWrites *bool `json:"enableMagneticStoreWrites,omitempty" tf:"enable_magnetic_store_writes,omitempty"`

	// The location to write error reports for records rejected asynchronously during magnetic store writes. See Magnetic Store Rejected Data Location below for more details.
	// +kubebuilder:validation:Optional
	MagneticStoreRejectedDataLocation *MagneticStoreRejectedDataLocationParameters `json:"magneticStoreRejectedDataLocation,omitempty" tf:"magnetic_store_rejected_data_location,omitempty"`
}

type RetentionPropertiesInitParameters struct {

	// The duration for which data must be stored in the magnetic store. Minimum value of 1. Maximum value of 73000.
	MagneticStoreRetentionPeriodInDays *float64 `json:"magneticStoreRetentionPeriodInDays,omitempty" tf:"magnetic_store_retention_period_in_days,omitempty"`

	// The duration for which data must be stored in the memory store. Minimum value of 1. Maximum value of 8766.
	MemoryStoreRetentionPeriodInHours *float64 `json:"memoryStoreRetentionPeriodInHours,omitempty" tf:"memory_store_retention_period_in_hours,omitempty"`
}

type RetentionPropertiesObservation struct {

	// The duration for which data must be stored in the magnetic store. Minimum value of 1. Maximum value of 73000.
	MagneticStoreRetentionPeriodInDays *float64 `json:"magneticStoreRetentionPeriodInDays,omitempty" tf:"magnetic_store_retention_period_in_days,omitempty"`

	// The duration for which data must be stored in the memory store. Minimum value of 1. Maximum value of 8766.
	MemoryStoreRetentionPeriodInHours *float64 `json:"memoryStoreRetentionPeriodInHours,omitempty" tf:"memory_store_retention_period_in_hours,omitempty"`
}

type RetentionPropertiesParameters struct {

	// The duration for which data must be stored in the magnetic store. Minimum value of 1. Maximum value of 73000.
	// +kubebuilder:validation:Optional
	MagneticStoreRetentionPeriodInDays *float64 `json:"magneticStoreRetentionPeriodInDays" tf:"magnetic_store_retention_period_in_days,omitempty"`

	// The duration for which data must be stored in the memory store. Minimum value of 1. Maximum value of 8766.
	// +kubebuilder:validation:Optional
	MemoryStoreRetentionPeriodInHours *float64 `json:"memoryStoreRetentionPeriodInHours" tf:"memory_store_retention_period_in_hours,omitempty"`
}

type S3ConfigurationInitParameters struct {

	// Bucket name of the customer S3 bucket.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Encryption option for the customer s3 location. Options are S3 server side encryption with an S3-managed key or KMS managed key. Valid values are SSE_KMS and SSE_S3.
	EncryptionOption *string `json:"encryptionOption,omitempty" tf:"encryption_option,omitempty"`

	// KMS key arn for the customer s3 location when encrypting with a KMS managed key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Object key prefix for the customer S3 location.
	ObjectKeyPrefix *string `json:"objectKeyPrefix,omitempty" tf:"object_key_prefix,omitempty"`
}

type S3ConfigurationObservation struct {

	// Bucket name of the customer S3 bucket.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Encryption option for the customer s3 location. Options are S3 server side encryption with an S3-managed key or KMS managed key. Valid values are SSE_KMS and SSE_S3.
	EncryptionOption *string `json:"encryptionOption,omitempty" tf:"encryption_option,omitempty"`

	// KMS key arn for the customer s3 location when encrypting with a KMS managed key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Object key prefix for the customer S3 location.
	ObjectKeyPrefix *string `json:"objectKeyPrefix,omitempty" tf:"object_key_prefix,omitempty"`
}

type S3ConfigurationParameters struct {

	// Bucket name of the customer S3 bucket.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Encryption option for the customer s3 location. Options are S3 server side encryption with an S3-managed key or KMS managed key. Valid values are SSE_KMS and SSE_S3.
	// +kubebuilder:validation:Optional
	EncryptionOption *string `json:"encryptionOption,omitempty" tf:"encryption_option,omitempty"`

	// KMS key arn for the customer s3 location when encrypting with a KMS managed key.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Object key prefix for the customer S3 location.
	// +kubebuilder:validation:Optional
	ObjectKeyPrefix *string `json:"objectKeyPrefix,omitempty" tf:"object_key_prefix,omitempty"`
}

type SchemaInitParameters struct {

	// A non-empty list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed. See Composite Partition Key below for more details.
	CompositePartitionKey *CompositePartitionKeyInitParameters `json:"compositePartitionKey,omitempty" tf:"composite_partition_key,omitempty"`
}

type SchemaObservation struct {

	// A non-empty list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed. See Composite Partition Key below for more details.
	CompositePartitionKey *CompositePartitionKeyObservation `json:"compositePartitionKey,omitempty" tf:"composite_partition_key,omitempty"`
}

type SchemaParameters struct {

	// A non-empty list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed. See Composite Partition Key below for more details.
	// +kubebuilder:validation:Optional
	CompositePartitionKey *CompositePartitionKeyParameters `json:"compositePartitionKey,omitempty" tf:"composite_partition_key,omitempty"`
}

type TableInitParameters struct {

	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	MagneticStoreWriteProperties *MagneticStoreWritePropertiesInitParameters `json:"magneticStoreWriteProperties,omitempty" tf:"magnetic_store_write_properties,omitempty"`

	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, magnetic_store_retention_period_in_days default to 73000 and memory_store_retention_period_in_hours defaults to 6.
	RetentionProperties *RetentionPropertiesInitParameters `json:"retentionProperties,omitempty" tf:"retention_properties,omitempty"`

	// The schema of the table. See Schema below for more details.
	Schema *SchemaInitParameters `json:"schema,omitempty" tf:"schema,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TableObservation struct {

	// The ARN that uniquely identifies this table.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// –  The name of the Timestream database.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The table_name and database_name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	MagneticStoreWriteProperties *MagneticStoreWritePropertiesObservation `json:"magneticStoreWriteProperties,omitempty" tf:"magnetic_store_write_properties,omitempty"`

	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, magnetic_store_retention_period_in_days default to 73000 and memory_store_retention_period_in_hours defaults to 6.
	RetentionProperties *RetentionPropertiesObservation `json:"retentionProperties,omitempty" tf:"retention_properties,omitempty"`

	// The schema of the table. See Schema below for more details.
	Schema *SchemaObservation `json:"schema,omitempty" tf:"schema,omitempty"`

	// The name of the Timestream table.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TableParameters struct {

	// –  The name of the Timestream database.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/timestreamwrite/v1beta1.Database
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a Database in timestreamwrite to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a Database in timestreamwrite to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	// +kubebuilder:validation:Optional
	MagneticStoreWriteProperties *MagneticStoreWritePropertiesParameters `json:"magneticStoreWriteProperties,omitempty" tf:"magnetic_store_write_properties,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, magnetic_store_retention_period_in_days default to 73000 and memory_store_retention_period_in_hours defaults to 6.
	// +kubebuilder:validation:Optional
	RetentionProperties *RetentionPropertiesParameters `json:"retentionProperties,omitempty" tf:"retention_properties,omitempty"`

	// The schema of the table. See Schema below for more details.
	// +kubebuilder:validation:Optional
	Schema *SchemaParameters `json:"schema,omitempty" tf:"schema,omitempty"`

	// The name of the Timestream table.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
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
	InitProvider TableInitParameters `json:"initProvider,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Table is the Schema for the Tables API. Provides a Timestream table resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
