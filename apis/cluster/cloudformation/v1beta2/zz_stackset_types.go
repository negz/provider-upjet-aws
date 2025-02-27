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

type AutoDeploymentInitParameters struct {

	// Whether or not auto-deployment is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether or not to retain stacks when the account is removed.
	RetainStacksOnAccountRemoval *bool `json:"retainStacksOnAccountRemoval,omitempty" tf:"retain_stacks_on_account_removal,omitempty"`
}

type AutoDeploymentObservation struct {

	// Whether or not auto-deployment is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether or not to retain stacks when the account is removed.
	RetainStacksOnAccountRemoval *bool `json:"retainStacksOnAccountRemoval,omitempty" tf:"retain_stacks_on_account_removal,omitempty"`
}

type AutoDeploymentParameters struct {

	// Whether or not auto-deployment is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether or not to retain stacks when the account is removed.
	// +kubebuilder:validation:Optional
	RetainStacksOnAccountRemoval *bool `json:"retainStacksOnAccountRemoval,omitempty" tf:"retain_stacks_on_account_removal,omitempty"`
}

type ManagedExecutionInitParameters struct {

	// When set to true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order. Default is false.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`
}

type ManagedExecutionObservation struct {

	// When set to true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order. Default is false.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`
}

type ManagedExecutionParameters struct {

	// When set to true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order. Default is false.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`
}

type OperationPreferencesInitParameters struct {

	// The number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureToleranceCount *float64 `json:"failureToleranceCount,omitempty" tf:"failure_tolerance_count,omitempty"`

	// The percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureTolerancePercentage *float64 `json:"failureTolerancePercentage,omitempty" tf:"failure_tolerance_percentage,omitempty"`

	// The maximum number of accounts in which to perform this operation at one time.
	MaxConcurrentCount *float64 `json:"maxConcurrentCount,omitempty" tf:"max_concurrent_count,omitempty"`

	// The maximum percentage of accounts in which to perform this operation at one time.
	MaxConcurrentPercentage *float64 `json:"maxConcurrentPercentage,omitempty" tf:"max_concurrent_percentage,omitempty"`

	// The concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time.
	RegionConcurrencyType *string `json:"regionConcurrencyType,omitempty" tf:"region_concurrency_type,omitempty"`

	// The order of the Regions in where you want to perform the stack operation.
	RegionOrder []*string `json:"regionOrder,omitempty" tf:"region_order,omitempty"`
}

type OperationPreferencesObservation struct {

	// The number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureToleranceCount *float64 `json:"failureToleranceCount,omitempty" tf:"failure_tolerance_count,omitempty"`

	// The percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureTolerancePercentage *float64 `json:"failureTolerancePercentage,omitempty" tf:"failure_tolerance_percentage,omitempty"`

	// The maximum number of accounts in which to perform this operation at one time.
	MaxConcurrentCount *float64 `json:"maxConcurrentCount,omitempty" tf:"max_concurrent_count,omitempty"`

	// The maximum percentage of accounts in which to perform this operation at one time.
	MaxConcurrentPercentage *float64 `json:"maxConcurrentPercentage,omitempty" tf:"max_concurrent_percentage,omitempty"`

	// The concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time.
	RegionConcurrencyType *string `json:"regionConcurrencyType,omitempty" tf:"region_concurrency_type,omitempty"`

	// The order of the Regions in where you want to perform the stack operation.
	RegionOrder []*string `json:"regionOrder,omitempty" tf:"region_order,omitempty"`
}

type OperationPreferencesParameters struct {

	// The number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	// +kubebuilder:validation:Optional
	FailureToleranceCount *float64 `json:"failureToleranceCount,omitempty" tf:"failure_tolerance_count,omitempty"`

	// The percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	// +kubebuilder:validation:Optional
	FailureTolerancePercentage *float64 `json:"failureTolerancePercentage,omitempty" tf:"failure_tolerance_percentage,omitempty"`

	// The maximum number of accounts in which to perform this operation at one time.
	// +kubebuilder:validation:Optional
	MaxConcurrentCount *float64 `json:"maxConcurrentCount,omitempty" tf:"max_concurrent_count,omitempty"`

	// The maximum percentage of accounts in which to perform this operation at one time.
	// +kubebuilder:validation:Optional
	MaxConcurrentPercentage *float64 `json:"maxConcurrentPercentage,omitempty" tf:"max_concurrent_percentage,omitempty"`

	// The concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time.
	// +kubebuilder:validation:Optional
	RegionConcurrencyType *string `json:"regionConcurrencyType,omitempty" tf:"region_concurrency_type,omitempty"`

	// The order of the Regions in where you want to perform the stack operation.
	// +kubebuilder:validation:Optional
	RegionOrder []*string `json:"regionOrder,omitempty" tf:"region_order,omitempty"`
}

type StackSetInitParameters struct {

	// Amazon Resource Number (ARN) of the IAM Role in the administrator account. This must be defined when using the SELF_MANAGED permission model.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	AdministrationRoleArn *string `json:"administrationRoleArn,omitempty" tf:"administration_role_arn,omitempty"`

	// Reference to a Role in iam to populate administrationRoleArn.
	// +kubebuilder:validation:Optional
	AdministrationRoleArnRef *v1.Reference `json:"administrationRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate administrationRoleArn.
	// +kubebuilder:validation:Optional
	AdministrationRoleArnSelector *v1.Selector `json:"administrationRoleArnSelector,omitempty" tf:"-"`

	// Configuration block containing the auto-deployment model for your StackSet. This can only be defined when using the SERVICE_MANAGED permission model.
	AutoDeployment *AutoDeploymentInitParameters `json:"autoDeployment,omitempty" tf:"auto_deployment,omitempty"`

	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account. Valid values: SELF (default), DELEGATED_ADMIN.
	CallAs *string `json:"callAs,omitempty" tf:"call_as,omitempty"`

	// A list of capabilities. Valid values: CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_AUTO_EXPAND.
	// +listType=set
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// Description of the StackSet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to AWSCloudFormationStackSetExecutionRole when using the SELF_MANAGED permission model. This should not be defined when using the SERVICE_MANAGED permission model.
	ExecutionRoleName *string `json:"executionRoleName,omitempty" tf:"execution_role_name,omitempty"`

	// Configuration block to allow StackSets to perform non-conflicting operations concurrently and queues conflicting operations.
	ManagedExecution *ManagedExecutionInitParameters `json:"managedExecution,omitempty" tf:"managed_execution,omitempty"`

	// Preferences for how AWS CloudFormation performs a stack set update.
	OperationPreferences *OperationPreferencesInitParameters `json:"operationPreferences,omitempty" tf:"operation_preferences,omitempty"`

	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a Default, must be configured or ignored with lifecycle configuration block ignore_changes argument. All NoEcho template parameters must be ignored with the lifecycle configuration block ignore_changes argument.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Describes how the IAM roles required for your StackSet are created. Valid values: SELF_MANAGED (default), SERVICE_MANAGED.
	PermissionModel *string `json:"permissionModel,omitempty" tf:"permission_model,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with template_url.
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with template_body.
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`
}

type StackSetObservation struct {

	// Amazon Resource Number (ARN) of the IAM Role in the administrator account. This must be defined when using the SELF_MANAGED permission model.
	AdministrationRoleArn *string `json:"administrationRoleArn,omitempty" tf:"administration_role_arn,omitempty"`

	// Amazon Resource Name (ARN) of the StackSet.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block containing the auto-deployment model for your StackSet. This can only be defined when using the SERVICE_MANAGED permission model.
	AutoDeployment *AutoDeploymentObservation `json:"autoDeployment,omitempty" tf:"auto_deployment,omitempty"`

	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account. Valid values: SELF (default), DELEGATED_ADMIN.
	CallAs *string `json:"callAs,omitempty" tf:"call_as,omitempty"`

	// A list of capabilities. Valid values: CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_AUTO_EXPAND.
	// +listType=set
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// Description of the StackSet.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to AWSCloudFormationStackSetExecutionRole when using the SELF_MANAGED permission model. This should not be defined when using the SERVICE_MANAGED permission model.
	ExecutionRoleName *string `json:"executionRoleName,omitempty" tf:"execution_role_name,omitempty"`

	// Name of the StackSet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block to allow StackSets to perform non-conflicting operations concurrently and queues conflicting operations.
	ManagedExecution *ManagedExecutionObservation `json:"managedExecution,omitempty" tf:"managed_execution,omitempty"`

	// Preferences for how AWS CloudFormation performs a stack set update.
	OperationPreferences *OperationPreferencesObservation `json:"operationPreferences,omitempty" tf:"operation_preferences,omitempty"`

	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a Default, must be configured or ignored with lifecycle configuration block ignore_changes argument. All NoEcho template parameters must be ignored with the lifecycle configuration block ignore_changes argument.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Describes how the IAM roles required for your StackSet are created. Valid values: SELF_MANAGED (default), SERVICE_MANAGED.
	PermissionModel *string `json:"permissionModel,omitempty" tf:"permission_model,omitempty"`

	// Unique identifier of the StackSet.
	StackSetID *string `json:"stackSetId,omitempty" tf:"stack_set_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with template_url.
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with template_body.
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`
}

type StackSetParameters struct {

	// Amazon Resource Number (ARN) of the IAM Role in the administrator account. This must be defined when using the SELF_MANAGED permission model.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	AdministrationRoleArn *string `json:"administrationRoleArn,omitempty" tf:"administration_role_arn,omitempty"`

	// Reference to a Role in iam to populate administrationRoleArn.
	// +kubebuilder:validation:Optional
	AdministrationRoleArnRef *v1.Reference `json:"administrationRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate administrationRoleArn.
	// +kubebuilder:validation:Optional
	AdministrationRoleArnSelector *v1.Selector `json:"administrationRoleArnSelector,omitempty" tf:"-"`

	// Configuration block containing the auto-deployment model for your StackSet. This can only be defined when using the SERVICE_MANAGED permission model.
	// +kubebuilder:validation:Optional
	AutoDeployment *AutoDeploymentParameters `json:"autoDeployment,omitempty" tf:"auto_deployment,omitempty"`

	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account. Valid values: SELF (default), DELEGATED_ADMIN.
	// +kubebuilder:validation:Optional
	CallAs *string `json:"callAs,omitempty" tf:"call_as,omitempty"`

	// A list of capabilities. Valid values: CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_AUTO_EXPAND.
	// +kubebuilder:validation:Optional
	// +listType=set
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// Description of the StackSet.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to AWSCloudFormationStackSetExecutionRole when using the SELF_MANAGED permission model. This should not be defined when using the SERVICE_MANAGED permission model.
	// +kubebuilder:validation:Optional
	ExecutionRoleName *string `json:"executionRoleName,omitempty" tf:"execution_role_name,omitempty"`

	// Configuration block to allow StackSets to perform non-conflicting operations concurrently and queues conflicting operations.
	// +kubebuilder:validation:Optional
	ManagedExecution *ManagedExecutionParameters `json:"managedExecution,omitempty" tf:"managed_execution,omitempty"`

	// Preferences for how AWS CloudFormation performs a stack set update.
	// +kubebuilder:validation:Optional
	OperationPreferences *OperationPreferencesParameters `json:"operationPreferences,omitempty" tf:"operation_preferences,omitempty"`

	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a Default, must be configured or ignored with lifecycle configuration block ignore_changes argument. All NoEcho template parameters must be ignored with the lifecycle configuration block ignore_changes argument.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Describes how the IAM roles required for your StackSet are created. Valid values: SELF_MANAGED (default), SERVICE_MANAGED.
	// +kubebuilder:validation:Optional
	PermissionModel *string `json:"permissionModel,omitempty" tf:"permission_model,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with template_url.
	// +kubebuilder:validation:Optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with template_body.
	// +kubebuilder:validation:Optional
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`
}

// StackSetSpec defines the desired state of StackSet
type StackSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackSetParameters `json:"forProvider"`
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
	InitProvider StackSetInitParameters `json:"initProvider,omitempty"`
}

// StackSetStatus defines the observed state of StackSet.
type StackSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StackSet is the Schema for the StackSets API. Manages a CloudFormation StackSet.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StackSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackSetSpec   `json:"spec"`
	Status            StackSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackSetList contains a list of StackSets
type StackSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StackSet `json:"items"`
}

// Repository type metadata.
var (
	StackSet_Kind             = "StackSet"
	StackSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StackSet_Kind}.String()
	StackSet_KindAPIVersion   = StackSet_Kind + "." + CRDGroupVersion.String()
	StackSet_GroupVersionKind = CRDGroupVersion.WithKind(StackSet_Kind)
)

func init() {
	SchemeBuilder.Register(&StackSet{}, &StackSetList{})
}
