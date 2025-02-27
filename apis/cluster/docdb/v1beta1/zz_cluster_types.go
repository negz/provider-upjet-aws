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

type ClusterInitParameters struct {

	// A value that indicates whether major version upgrades are allowed. Constraints: You must allow major version upgrades when specifying a value for the EngineVersion parameter that is a different major version than the DB cluster's current version.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications
	// are applied immediately, or during the next maintenance window. Default is
	// false.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// A list of EC2 Availability Zones that
	// instances in the DB cluster can be created in.
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The days to retain backups for. Default 1
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// A cluster parameter group to associate with the cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/docdb/v1beta1.ClusterParameterGroup
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty" tf:"db_cluster_parameter_group_name,omitempty"`

	// Reference to a ClusterParameterGroup in docdb to populate dbClusterParameterGroupName.
	// +kubebuilder:validation:Optional
	DBClusterParameterGroupNameRef *v1.Reference `json:"dbClusterParameterGroupNameRef,omitempty" tf:"-"`

	// Selector for a ClusterParameterGroup in docdb to populate dbClusterParameterGroupName.
	// +kubebuilder:validation:Optional
	DBClusterParameterGroupNameSelector *v1.Selector `json:"dbClusterParameterGroupNameSelector,omitempty" tf:"-"`

	// A DB subnet group to associate with this DB instance.
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// A boolean value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. Defaults to false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// List of log types to export to cloudwatch. If omitted, no logs will be exported.
	// The following log types are supported: audit, profiler.
	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`

	// The name of the database engine to be used for this DB cluster. Defaults to docdb. Valid values: docdb.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The database engine version. Updating this argument results in an outage.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The name of your final DB snapshot
	// when this DB cluster is deleted. If omitted, no final snapshot will be
	// made.
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// The global cluster identifier specified on aws_docdb_global_cluster.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_id, storage_encrypted needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Password for the master DB user. Note that this may
	// show up in logs, and it will be stored in the state file. Please refer to the DocumentDB Naming Constraints.
	// Password for the master DB user. If you set autoGeneratePassword to true, the Secret referenced here will be created or updated with generated password if it does not already contain one.
	MasterPasswordSecretRef *v1.SecretKeySelector `json:"masterPasswordSecretRef,omitempty" tf:"-"`

	// Username for the master DB user.
	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`

	// The port on which the DB accepts connections
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
	// Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// A configuration block for restoring a DB instance to an arbitrary point in time. Requires the identifier argument to be set with the name of the new DB instance to be created. See Restore To Point In Time below for details.
	RestoreToPointInTime *RestoreToPointInTimeInitParameters `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time,omitempty"`

	// Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from final_snapshot_identifier. Default is false.
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot. Automated snapshots should not be used for this attribute, unless from a different cluster. Automated snapshots are deleted as part of cluster destruction when the resource is replaced.
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is false.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// The storage type to associate with the DB cluster. Valid values: standard, iopt1.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// List of VPC security groups to associate
	// with the Cluster
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +listType=set
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type ClusterObservation struct {

	// A value that indicates whether major version upgrades are allowed. Constraints: You must allow major version upgrades when specifying a value for the EngineVersion parameter that is a different major version than the DB cluster's current version.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications
	// are applied immediately, or during the next maintenance window. Default is
	// false.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Amazon Resource Name (ARN) of cluster
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A list of EC2 Availability Zones that
	// instances in the DB cluster can be created in.
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The days to retain backups for. Default 1
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// – List of DocumentDB Instances that are a part of this cluster
	// +listType=set
	ClusterMembers []*string `json:"clusterMembers,omitempty" tf:"cluster_members,omitempty"`

	// The DocumentDB Cluster Resource ID
	ClusterResourceID *string `json:"clusterResourceId,omitempty" tf:"cluster_resource_id,omitempty"`

	// A cluster parameter group to associate with the cluster.
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty" tf:"db_cluster_parameter_group_name,omitempty"`

	// A DB subnet group to associate with this DB instance.
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// A boolean value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. Defaults to false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// List of log types to export to cloudwatch. If omitted, no logs will be exported.
	// The following log types are supported: audit, profiler.
	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`

	// The DNS address of the DocumentDB instance
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The name of the database engine to be used for this DB cluster. Defaults to docdb. Valid values: docdb.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The database engine version. Updating this argument results in an outage.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The name of your final DB snapshot
	// when this DB cluster is deleted. If omitted, no final snapshot will be
	// made.
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// The global cluster identifier specified on aws_docdb_global_cluster.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// The Route53 Hosted Zone ID of the endpoint
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// The DocumentDB Cluster Identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_id, storage_encrypted needs to be set to true.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Username for the master DB user.
	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`

	// The port on which the DB accepts connections
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
	// Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// A read-only endpoint for the DocumentDB cluster, automatically load-balanced across replicas
	ReaderEndpoint *string `json:"readerEndpoint,omitempty" tf:"reader_endpoint,omitempty"`

	// A configuration block for restoring a DB instance to an arbitrary point in time. Requires the identifier argument to be set with the name of the new DB instance to be created. See Restore To Point In Time below for details.
	RestoreToPointInTime *RestoreToPointInTimeObservation `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time,omitempty"`

	// Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from final_snapshot_identifier. Default is false.
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot. Automated snapshots should not be used for this attribute, unless from a different cluster. Automated snapshots are deleted as part of cluster destruction when the resource is replaced.
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is false.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// The storage type to associate with the DB cluster. Valid values: standard, iopt1.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// List of VPC security groups to associate
	// with the Cluster
	// +listType=set
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type ClusterParameters struct {

	// A value that indicates whether major version upgrades are allowed. Constraints: You must allow major version upgrades when specifying a value for the EngineVersion parameter that is a different major version than the DB cluster's current version.
	// +kubebuilder:validation:Optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications
	// are applied immediately, or during the next maintenance window. Default is
	// false.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// If true, the password will be auto-generated and stored in the Secret referenced by the masterPasswordSecretRef field.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	AutoGeneratePassword *bool `json:"autoGeneratePassword,omitempty" tf:"-"`

	// A list of EC2 Availability Zones that
	// instances in the DB cluster can be created in.
	// +kubebuilder:validation:Optional
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The days to retain backups for. Default 1
	// +kubebuilder:validation:Optional
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// A cluster parameter group to associate with the cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/docdb/v1beta1.ClusterParameterGroup
	// +kubebuilder:validation:Optional
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty" tf:"db_cluster_parameter_group_name,omitempty"`

	// Reference to a ClusterParameterGroup in docdb to populate dbClusterParameterGroupName.
	// +kubebuilder:validation:Optional
	DBClusterParameterGroupNameRef *v1.Reference `json:"dbClusterParameterGroupNameRef,omitempty" tf:"-"`

	// Selector for a ClusterParameterGroup in docdb to populate dbClusterParameterGroupName.
	// +kubebuilder:validation:Optional
	DBClusterParameterGroupNameSelector *v1.Selector `json:"dbClusterParameterGroupNameSelector,omitempty" tf:"-"`

	// A DB subnet group to associate with this DB instance.
	// +kubebuilder:validation:Optional
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// A boolean value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// List of log types to export to cloudwatch. If omitted, no logs will be exported.
	// The following log types are supported: audit, profiler.
	// +kubebuilder:validation:Optional
	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`

	// The name of the database engine to be used for this DB cluster. Defaults to docdb. Valid values: docdb.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The database engine version. Updating this argument results in an outage.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The name of your final DB snapshot
	// when this DB cluster is deleted. If omitted, no final snapshot will be
	// made.
	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// The global cluster identifier specified on aws_docdb_global_cluster.
	// +kubebuilder:validation:Optional
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_id, storage_encrypted needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Password for the master DB user. Note that this may
	// show up in logs, and it will be stored in the state file. Please refer to the DocumentDB Naming Constraints.
	// Password for the master DB user. If you set autoGeneratePassword to true, the Secret referenced here will be created or updated with generated password if it does not already contain one.
	// +kubebuilder:validation:Optional
	MasterPasswordSecretRef *v1.SecretKeySelector `json:"masterPasswordSecretRef,omitempty" tf:"-"`

	// Username for the master DB user.
	// +kubebuilder:validation:Optional
	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`

	// The port on which the DB accepts connections
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
	// Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	// +kubebuilder:validation:Optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A configuration block for restoring a DB instance to an arbitrary point in time. Requires the identifier argument to be set with the name of the new DB instance to be created. See Restore To Point In Time below for details.
	// +kubebuilder:validation:Optional
	RestoreToPointInTime *RestoreToPointInTimeParameters `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time,omitempty"`

	// Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from final_snapshot_identifier. Default is false.
	// +kubebuilder:validation:Optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot. Automated snapshots should not be used for this attribute, unless from a different cluster. Automated snapshots are deleted as part of cluster destruction when the resource is replaced.
	// +kubebuilder:validation:Optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is false.
	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// The storage type to associate with the DB cluster. Valid values: standard, iopt1.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// List of VPC security groups to associate
	// with the Cluster
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type RestoreToPointInTimeInitParameters struct {

	// The date and time to restore from. Value must be a time in Universal Coordinated Time (UTC) format and must be before the latest restorable time for the DB instance. Cannot be specified with use_latest_restorable_time.
	RestoreToTime *string `json:"restoreToTime,omitempty" tf:"restore_to_time,omitempty"`

	// The type of restore to be performed. Valid values are full-copy, copy-on-write.
	RestoreType *string `json:"restoreType,omitempty" tf:"restore_type,omitempty"`

	// The identifier of the source DB cluster from which to restore. Must match the identifier of an existing DB cluster.
	SourceClusterIdentifier *string `json:"sourceClusterIdentifier,omitempty" tf:"source_cluster_identifier,omitempty"`

	// A boolean value that indicates whether the DB cluster is restored from the latest backup time. Defaults to false. Cannot be specified with restore_to_time.
	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time,omitempty"`
}

type RestoreToPointInTimeObservation struct {

	// The date and time to restore from. Value must be a time in Universal Coordinated Time (UTC) format and must be before the latest restorable time for the DB instance. Cannot be specified with use_latest_restorable_time.
	RestoreToTime *string `json:"restoreToTime,omitempty" tf:"restore_to_time,omitempty"`

	// The type of restore to be performed. Valid values are full-copy, copy-on-write.
	RestoreType *string `json:"restoreType,omitempty" tf:"restore_type,omitempty"`

	// The identifier of the source DB cluster from which to restore. Must match the identifier of an existing DB cluster.
	SourceClusterIdentifier *string `json:"sourceClusterIdentifier,omitempty" tf:"source_cluster_identifier,omitempty"`

	// A boolean value that indicates whether the DB cluster is restored from the latest backup time. Defaults to false. Cannot be specified with restore_to_time.
	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time,omitempty"`
}

type RestoreToPointInTimeParameters struct {

	// The date and time to restore from. Value must be a time in Universal Coordinated Time (UTC) format and must be before the latest restorable time for the DB instance. Cannot be specified with use_latest_restorable_time.
	// +kubebuilder:validation:Optional
	RestoreToTime *string `json:"restoreToTime,omitempty" tf:"restore_to_time,omitempty"`

	// The type of restore to be performed. Valid values are full-copy, copy-on-write.
	// +kubebuilder:validation:Optional
	RestoreType *string `json:"restoreType,omitempty" tf:"restore_type,omitempty"`

	// The identifier of the source DB cluster from which to restore. Must match the identifier of an existing DB cluster.
	// +kubebuilder:validation:Optional
	SourceClusterIdentifier *string `json:"sourceClusterIdentifier" tf:"source_cluster_identifier,omitempty"`

	// A boolean value that indicates whether the DB cluster is restored from the latest backup time. Defaults to false. Cannot be specified with restore_to_time.
	// +kubebuilder:validation:Optional
	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Cluster is the Schema for the Clusters API. Manages a DocumentDB Aurora Cluster
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
