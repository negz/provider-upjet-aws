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

type AdvancedBackupSettingInitParameters struct {

	// Specifies the backup option for a selected resource. This option is only available for Windows VSS backup jobs. Set to { WindowsVSS = "enabled" } to enable Windows VSS backup option and create a VSS Windows backup.
	// +mapType=granular
	BackupOptions map[string]*string `json:"backupOptions,omitempty" tf:"backup_options,omitempty"`

	// The type of AWS resource to be backed up. For VSS Windows backups, the only supported resource type is Amazon EC2. Valid values: EC2.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type AdvancedBackupSettingObservation struct {

	// Specifies the backup option for a selected resource. This option is only available for Windows VSS backup jobs. Set to { WindowsVSS = "enabled" } to enable Windows VSS backup option and create a VSS Windows backup.
	// +mapType=granular
	BackupOptions map[string]*string `json:"backupOptions,omitempty" tf:"backup_options,omitempty"`

	// The type of AWS resource to be backed up. For VSS Windows backups, the only supported resource type is Amazon EC2. Valid values: EC2.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type AdvancedBackupSettingParameters struct {

	// Specifies the backup option for a selected resource. This option is only available for Windows VSS backup jobs. Set to { WindowsVSS = "enabled" } to enable Windows VSS backup option and create a VSS Windows backup.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	BackupOptions map[string]*string `json:"backupOptions" tf:"backup_options,omitempty"`

	// The type of AWS resource to be backed up. For VSS Windows backups, the only supported resource type is Amazon EC2. Valid values: EC2.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`
}

type CopyActionInitParameters struct {

	// An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
	DestinationVaultArn *string `json:"destinationVaultArn,omitempty" tf:"destination_vault_arn,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	Lifecycle *LifecycleInitParameters `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`
}

type CopyActionObservation struct {

	// An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
	DestinationVaultArn *string `json:"destinationVaultArn,omitempty" tf:"destination_vault_arn,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	Lifecycle *LifecycleObservation `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`
}

type CopyActionParameters struct {

	// An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
	// +kubebuilder:validation:Optional
	DestinationVaultArn *string `json:"destinationVaultArn" tf:"destination_vault_arn,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	// +kubebuilder:validation:Optional
	Lifecycle *LifecycleParameters `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`
}

type LifecycleInitParameters struct {

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	ColdStorageAfter *float64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`

	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than cold_storage_after.
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// This setting will instruct your backup plan to transition supported resources to archive (cold) storage tier in accordance with your lifecycle settings.
	OptInToArchiveForSupportedResources *bool `json:"optInToArchiveForSupportedResources,omitempty" tf:"opt_in_to_archive_for_supported_resources,omitempty"`
}

type LifecycleObservation struct {

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	ColdStorageAfter *float64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`

	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than cold_storage_after.
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// This setting will instruct your backup plan to transition supported resources to archive (cold) storage tier in accordance with your lifecycle settings.
	OptInToArchiveForSupportedResources *bool `json:"optInToArchiveForSupportedResources,omitempty" tf:"opt_in_to_archive_for_supported_resources,omitempty"`
}

type LifecycleParameters struct {

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	// +kubebuilder:validation:Optional
	ColdStorageAfter *float64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`

	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than cold_storage_after.
	// +kubebuilder:validation:Optional
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// This setting will instruct your backup plan to transition supported resources to archive (cold) storage tier in accordance with your lifecycle settings.
	// +kubebuilder:validation:Optional
	OptInToArchiveForSupportedResources *bool `json:"optInToArchiveForSupportedResources,omitempty" tf:"opt_in_to_archive_for_supported_resources,omitempty"`
}

type PlanInitParameters struct {

	// An object that specifies backup options for each resource type.
	AdvancedBackupSetting []AdvancedBackupSettingInitParameters `json:"advancedBackupSetting,omitempty" tf:"advanced_backup_setting,omitempty"`

	// The display name of a backup plan.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PlanObservation struct {

	// An object that specifies backup options for each resource type.
	AdvancedBackupSetting []AdvancedBackupSettingObservation `json:"advancedBackupSetting,omitempty" tf:"advanced_backup_setting,omitempty"`

	// The ARN of the backup plan.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The id of the backup plan.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The display name of a backup plan.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PlanParameters struct {

	// An object that specifies backup options for each resource type.
	// +kubebuilder:validation:Optional
	AdvancedBackupSetting []AdvancedBackupSettingParameters `json:"advancedBackupSetting,omitempty" tf:"advanced_backup_setting,omitempty"`

	// The display name of a backup plan.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RuleInitParameters struct {

	// The amount of time in minutes AWS Backup attempts a backup before canceling the job and returning an error.
	CompletionWindow *float64 `json:"completionWindow,omitempty" tf:"completion_window,omitempty"`

	// Configuration block(s) with copy operation settings. Detailed below.
	CopyAction []CopyActionInitParameters `json:"copyAction,omitempty" tf:"copy_action,omitempty"`

	// Enable continuous backups for supported resources.
	EnableContinuousBackup *bool `json:"enableContinuousBackup,omitempty" tf:"enable_continuous_backup,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	Lifecycle *RuleLifecycleInitParameters `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`

	// Metadata that you can assign to help organize the resources that you create.
	// +mapType=granular
	RecoveryPointTags map[string]*string `json:"recoveryPointTags,omitempty" tf:"recovery_point_tags,omitempty"`

	// An display name for a backup rule.
	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	// A CRON expression specifying when AWS Backup initiates a backup job.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The timezone in which the schedule expression is set. Default value: "Etc/UTC".
	ScheduleExpressionTimezone *string `json:"scheduleExpressionTimezone,omitempty" tf:"schedule_expression_timezone,omitempty"`

	// The amount of time in minutes before beginning a backup.
	StartWindow *float64 `json:"startWindow,omitempty" tf:"start_window,omitempty"`

	// The name of a logical container where backups are stored.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/backup/v1beta1.Vault
	TargetVaultName *string `json:"targetVaultName,omitempty" tf:"target_vault_name,omitempty"`

	// Reference to a Vault in backup to populate targetVaultName.
	// +kubebuilder:validation:Optional
	TargetVaultNameRef *v1.Reference `json:"targetVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in backup to populate targetVaultName.
	// +kubebuilder:validation:Optional
	TargetVaultNameSelector *v1.Selector `json:"targetVaultNameSelector,omitempty" tf:"-"`
}

type RuleLifecycleInitParameters struct {

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	ColdStorageAfter *float64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`

	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than cold_storage_after.
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// This setting will instruct your backup plan to transition supported resources to archive (cold) storage tier in accordance with your lifecycle settings.
	OptInToArchiveForSupportedResources *bool `json:"optInToArchiveForSupportedResources,omitempty" tf:"opt_in_to_archive_for_supported_resources,omitempty"`
}

type RuleLifecycleObservation struct {

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	ColdStorageAfter *float64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`

	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than cold_storage_after.
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// This setting will instruct your backup plan to transition supported resources to archive (cold) storage tier in accordance with your lifecycle settings.
	OptInToArchiveForSupportedResources *bool `json:"optInToArchiveForSupportedResources,omitempty" tf:"opt_in_to_archive_for_supported_resources,omitempty"`
}

type RuleLifecycleParameters struct {

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	// +kubebuilder:validation:Optional
	ColdStorageAfter *float64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`

	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than cold_storage_after.
	// +kubebuilder:validation:Optional
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`

	// This setting will instruct your backup plan to transition supported resources to archive (cold) storage tier in accordance with your lifecycle settings.
	// +kubebuilder:validation:Optional
	OptInToArchiveForSupportedResources *bool `json:"optInToArchiveForSupportedResources,omitempty" tf:"opt_in_to_archive_for_supported_resources,omitempty"`
}

type RuleObservation struct {

	// The amount of time in minutes AWS Backup attempts a backup before canceling the job and returning an error.
	CompletionWindow *float64 `json:"completionWindow,omitempty" tf:"completion_window,omitempty"`

	// Configuration block(s) with copy operation settings. Detailed below.
	CopyAction []CopyActionObservation `json:"copyAction,omitempty" tf:"copy_action,omitempty"`

	// Enable continuous backups for supported resources.
	EnableContinuousBackup *bool `json:"enableContinuousBackup,omitempty" tf:"enable_continuous_backup,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	Lifecycle *RuleLifecycleObservation `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`

	// Metadata that you can assign to help organize the resources that you create.
	// +mapType=granular
	RecoveryPointTags map[string]*string `json:"recoveryPointTags,omitempty" tf:"recovery_point_tags,omitempty"`

	// An display name for a backup rule.
	RuleName *string `json:"ruleName,omitempty" tf:"rule_name,omitempty"`

	// A CRON expression specifying when AWS Backup initiates a backup job.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The timezone in which the schedule expression is set. Default value: "Etc/UTC".
	ScheduleExpressionTimezone *string `json:"scheduleExpressionTimezone,omitempty" tf:"schedule_expression_timezone,omitempty"`

	// The amount of time in minutes before beginning a backup.
	StartWindow *float64 `json:"startWindow,omitempty" tf:"start_window,omitempty"`

	// The name of a logical container where backups are stored.
	TargetVaultName *string `json:"targetVaultName,omitempty" tf:"target_vault_name,omitempty"`
}

type RuleParameters struct {

	// The amount of time in minutes AWS Backup attempts a backup before canceling the job and returning an error.
	// +kubebuilder:validation:Optional
	CompletionWindow *float64 `json:"completionWindow,omitempty" tf:"completion_window,omitempty"`

	// Configuration block(s) with copy operation settings. Detailed below.
	// +kubebuilder:validation:Optional
	CopyAction []CopyActionParameters `json:"copyAction,omitempty" tf:"copy_action,omitempty"`

	// Enable continuous backups for supported resources.
	// +kubebuilder:validation:Optional
	EnableContinuousBackup *bool `json:"enableContinuousBackup,omitempty" tf:"enable_continuous_backup,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	// +kubebuilder:validation:Optional
	Lifecycle *RuleLifecycleParameters `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`

	// Metadata that you can assign to help organize the resources that you create.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	RecoveryPointTags map[string]*string `json:"recoveryPointTags,omitempty" tf:"recovery_point_tags,omitempty"`

	// An display name for a backup rule.
	// +kubebuilder:validation:Optional
	RuleName *string `json:"ruleName" tf:"rule_name,omitempty"`

	// A CRON expression specifying when AWS Backup initiates a backup job.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The timezone in which the schedule expression is set. Default value: "Etc/UTC".
	// +kubebuilder:validation:Optional
	ScheduleExpressionTimezone *string `json:"scheduleExpressionTimezone,omitempty" tf:"schedule_expression_timezone,omitempty"`

	// The amount of time in minutes before beginning a backup.
	// +kubebuilder:validation:Optional
	StartWindow *float64 `json:"startWindow,omitempty" tf:"start_window,omitempty"`

	// The name of a logical container where backups are stored.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/backup/v1beta1.Vault
	// +kubebuilder:validation:Optional
	TargetVaultName *string `json:"targetVaultName,omitempty" tf:"target_vault_name,omitempty"`

	// Reference to a Vault in backup to populate targetVaultName.
	// +kubebuilder:validation:Optional
	TargetVaultNameRef *v1.Reference `json:"targetVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in backup to populate targetVaultName.
	// +kubebuilder:validation:Optional
	TargetVaultNameSelector *v1.Selector `json:"targetVaultNameSelector,omitempty" tf:"-"`
}

// PlanSpec defines the desired state of Plan
type PlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlanParameters `json:"forProvider"`
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
	InitProvider PlanInitParameters `json:"initProvider,omitempty"`
}

// PlanStatus defines the observed state of Plan.
type PlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Plan is the Schema for the Plans API. Provides an AWS Backup plan resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Plan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule) || (has(self.initProvider) && has(self.initProvider.rule))",message="spec.forProvider.rule is a required parameter"
	Spec   PlanSpec   `json:"spec"`
	Status PlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlanList contains a list of Plans
type PlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Plan `json:"items"`
}

// Repository type metadata.
var (
	Plan_Kind             = "Plan"
	Plan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Plan_Kind}.String()
	Plan_KindAPIVersion   = Plan_Kind + "." + CRDGroupVersion.String()
	Plan_GroupVersionKind = CRDGroupVersion.WithKind(Plan_Kind)
)

func init() {
	SchemeBuilder.Register(&Plan{}, &PlanList{})
}
