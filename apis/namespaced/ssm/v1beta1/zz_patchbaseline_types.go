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

type ApprovalRuleInitParameters struct {

	// Number of days after the release date of each patch matched by the rule the patch is marked as approved in the patch baseline. Valid Range: 0 to 360. Conflicts with approve_until_date.
	ApproveAfterDays *float64 `json:"approveAfterDays,omitempty" tf:"approve_after_days,omitempty"`

	// Cutoff date for auto approval of released patches. Any patches released on or before this date are installed automatically. Date is formatted as YYYY-MM-DD. Conflicts with approve_after_days
	ApproveUntilDate *string `json:"approveUntilDate,omitempty" tf:"approve_until_date,omitempty"`

	// Compliance level for patches approved by this rule. Valid values are CRITICAL, HIGH, MEDIUM, LOW, INFORMATIONAL, and UNSPECIFIED. The default value is UNSPECIFIED.
	ComplianceLevel *string `json:"complianceLevel,omitempty" tf:"compliance_level,omitempty"`

	// Boolean enabling the application of non-security updates. The default value is false. Valid for Linux instances only.
	EnableNonSecurity *bool `json:"enableNonSecurity,omitempty" tf:"enable_non_security,omitempty"`

	// Patch filter group that defines the criteria for the rule. Up to 5 patch filters can be specified per approval rule using Key/Value pairs. Valid combinations of these Keys and the operating_system value can be found in the SSM DescribePatchProperties API Reference. Valid Values are exact values for the patch property given as the key, or a wildcard *, which matches all values. PATCH_SET defaults to OS if unspecified
	PatchFilter []PatchFilterInitParameters `json:"patchFilter,omitempty" tf:"patch_filter,omitempty"`
}

type ApprovalRuleObservation struct {

	// Number of days after the release date of each patch matched by the rule the patch is marked as approved in the patch baseline. Valid Range: 0 to 360. Conflicts with approve_until_date.
	ApproveAfterDays *float64 `json:"approveAfterDays,omitempty" tf:"approve_after_days,omitempty"`

	// Cutoff date for auto approval of released patches. Any patches released on or before this date are installed automatically. Date is formatted as YYYY-MM-DD. Conflicts with approve_after_days
	ApproveUntilDate *string `json:"approveUntilDate,omitempty" tf:"approve_until_date,omitempty"`

	// Compliance level for patches approved by this rule. Valid values are CRITICAL, HIGH, MEDIUM, LOW, INFORMATIONAL, and UNSPECIFIED. The default value is UNSPECIFIED.
	ComplianceLevel *string `json:"complianceLevel,omitempty" tf:"compliance_level,omitempty"`

	// Boolean enabling the application of non-security updates. The default value is false. Valid for Linux instances only.
	EnableNonSecurity *bool `json:"enableNonSecurity,omitempty" tf:"enable_non_security,omitempty"`

	// Patch filter group that defines the criteria for the rule. Up to 5 patch filters can be specified per approval rule using Key/Value pairs. Valid combinations of these Keys and the operating_system value can be found in the SSM DescribePatchProperties API Reference. Valid Values are exact values for the patch property given as the key, or a wildcard *, which matches all values. PATCH_SET defaults to OS if unspecified
	PatchFilter []PatchFilterObservation `json:"patchFilter,omitempty" tf:"patch_filter,omitempty"`
}

type ApprovalRuleParameters struct {

	// Number of days after the release date of each patch matched by the rule the patch is marked as approved in the patch baseline. Valid Range: 0 to 360. Conflicts with approve_until_date.
	// +kubebuilder:validation:Optional
	ApproveAfterDays *float64 `json:"approveAfterDays,omitempty" tf:"approve_after_days,omitempty"`

	// Cutoff date for auto approval of released patches. Any patches released on or before this date are installed automatically. Date is formatted as YYYY-MM-DD. Conflicts with approve_after_days
	// +kubebuilder:validation:Optional
	ApproveUntilDate *string `json:"approveUntilDate,omitempty" tf:"approve_until_date,omitempty"`

	// Compliance level for patches approved by this rule. Valid values are CRITICAL, HIGH, MEDIUM, LOW, INFORMATIONAL, and UNSPECIFIED. The default value is UNSPECIFIED.
	// +kubebuilder:validation:Optional
	ComplianceLevel *string `json:"complianceLevel,omitempty" tf:"compliance_level,omitempty"`

	// Boolean enabling the application of non-security updates. The default value is false. Valid for Linux instances only.
	// +kubebuilder:validation:Optional
	EnableNonSecurity *bool `json:"enableNonSecurity,omitempty" tf:"enable_non_security,omitempty"`

	// Patch filter group that defines the criteria for the rule. Up to 5 patch filters can be specified per approval rule using Key/Value pairs. Valid combinations of these Keys and the operating_system value can be found in the SSM DescribePatchProperties API Reference. Valid Values are exact values for the patch property given as the key, or a wildcard *, which matches all values. PATCH_SET defaults to OS if unspecified
	// +kubebuilder:validation:Optional
	PatchFilter []PatchFilterParameters `json:"patchFilter" tf:"patch_filter,omitempty"`
}

type GlobalFilterInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type GlobalFilterObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type GlobalFilterParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type PatchBaselineInitParameters struct {

	// Set of rules used to include patches in the baseline. Up to 10 approval rules can be specified. See approval_rule below.
	ApprovalRule []ApprovalRuleInitParameters `json:"approvalRule,omitempty" tf:"approval_rule,omitempty"`

	// List of explicitly approved patches for the baseline. Cannot be specified with approval_rule.
	// +listType=set
	ApprovedPatches []*string `json:"approvedPatches,omitempty" tf:"approved_patches,omitempty"`

	// Compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid values are CRITICAL, HIGH, MEDIUM, LOW, INFORMATIONAL, UNSPECIFIED. The default value is UNSPECIFIED.
	ApprovedPatchesComplianceLevel *string `json:"approvedPatchesComplianceLevel,omitempty" tf:"approved_patches_compliance_level,omitempty"`

	// Whether the list of approved patches includes non-security updates that should be applied to the instances. Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool `json:"approvedPatchesEnableNonSecurity,omitempty" tf:"approved_patches_enable_non_security,omitempty"`

	// Description of the patch baseline.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are PRODUCT, CLASSIFICATION, MSRC_SEVERITY, and PATCH_ID.
	GlobalFilter []GlobalFilterInitParameters `json:"globalFilter,omitempty" tf:"global_filter,omitempty"`

	// Name of the patch baseline.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system the patch baseline applies to. Valid values are ALMA_LINUX, AMAZON_LINUX, AMAZON_LINUX_2, AMAZON_LINUX_2022, AMAZON_LINUX_2023, CENTOS, DEBIAN, MACOS, ORACLE_LINUX, RASPBIAN, REDHAT_ENTERPRISE_LINUX, ROCKY_LINUX, SUSE, UBUNTU, and WINDOWS. The default value is WINDOWS.
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// List of rejected patches.
	// +listType=set
	RejectedPatches []*string `json:"rejectedPatches,omitempty" tf:"rejected_patches,omitempty"`

	// Action for Patch Manager to take on patches included in the rejected_patches list. Valid values are ALLOW_AS_DEPENDENCY and BLOCK.
	RejectedPatchesAction *string `json:"rejectedPatchesAction,omitempty" tf:"rejected_patches_action,omitempty"`

	// Configuration block with alternate sources for patches. Applies to Linux instances only. See source below.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PatchBaselineObservation struct {

	// Set of rules used to include patches in the baseline. Up to 10 approval rules can be specified. See approval_rule below.
	ApprovalRule []ApprovalRuleObservation `json:"approvalRule,omitempty" tf:"approval_rule,omitempty"`

	// List of explicitly approved patches for the baseline. Cannot be specified with approval_rule.
	// +listType=set
	ApprovedPatches []*string `json:"approvedPatches,omitempty" tf:"approved_patches,omitempty"`

	// Compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid values are CRITICAL, HIGH, MEDIUM, LOW, INFORMATIONAL, UNSPECIFIED. The default value is UNSPECIFIED.
	ApprovedPatchesComplianceLevel *string `json:"approvedPatchesComplianceLevel,omitempty" tf:"approved_patches_compliance_level,omitempty"`

	// Whether the list of approved patches includes non-security updates that should be applied to the instances. Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool `json:"approvedPatchesEnableNonSecurity,omitempty" tf:"approved_patches_enable_non_security,omitempty"`

	// ARN of the baseline.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the patch baseline.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are PRODUCT, CLASSIFICATION, MSRC_SEVERITY, and PATCH_ID.
	GlobalFilter []GlobalFilterObservation `json:"globalFilter,omitempty" tf:"global_filter,omitempty"`

	// ID of the baseline.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// JSON definition of the baseline.
	JSON *string `json:"json,omitempty" tf:"json,omitempty"`

	// Name of the patch baseline.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system the patch baseline applies to. Valid values are ALMA_LINUX, AMAZON_LINUX, AMAZON_LINUX_2, AMAZON_LINUX_2022, AMAZON_LINUX_2023, CENTOS, DEBIAN, MACOS, ORACLE_LINUX, RASPBIAN, REDHAT_ENTERPRISE_LINUX, ROCKY_LINUX, SUSE, UBUNTU, and WINDOWS. The default value is WINDOWS.
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// List of rejected patches.
	// +listType=set
	RejectedPatches []*string `json:"rejectedPatches,omitempty" tf:"rejected_patches,omitempty"`

	// Action for Patch Manager to take on patches included in the rejected_patches list. Valid values are ALLOW_AS_DEPENDENCY and BLOCK.
	RejectedPatchesAction *string `json:"rejectedPatchesAction,omitempty" tf:"rejected_patches_action,omitempty"`

	// Configuration block with alternate sources for patches. Applies to Linux instances only. See source below.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PatchBaselineParameters struct {

	// Set of rules used to include patches in the baseline. Up to 10 approval rules can be specified. See approval_rule below.
	// +kubebuilder:validation:Optional
	ApprovalRule []ApprovalRuleParameters `json:"approvalRule,omitempty" tf:"approval_rule,omitempty"`

	// List of explicitly approved patches for the baseline. Cannot be specified with approval_rule.
	// +kubebuilder:validation:Optional
	// +listType=set
	ApprovedPatches []*string `json:"approvedPatches,omitempty" tf:"approved_patches,omitempty"`

	// Compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid values are CRITICAL, HIGH, MEDIUM, LOW, INFORMATIONAL, UNSPECIFIED. The default value is UNSPECIFIED.
	// +kubebuilder:validation:Optional
	ApprovedPatchesComplianceLevel *string `json:"approvedPatchesComplianceLevel,omitempty" tf:"approved_patches_compliance_level,omitempty"`

	// Whether the list of approved patches includes non-security updates that should be applied to the instances. Applies to Linux instances only.
	// +kubebuilder:validation:Optional
	ApprovedPatchesEnableNonSecurity *bool `json:"approvedPatchesEnableNonSecurity,omitempty" tf:"approved_patches_enable_non_security,omitempty"`

	// Description of the patch baseline.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are PRODUCT, CLASSIFICATION, MSRC_SEVERITY, and PATCH_ID.
	// +kubebuilder:validation:Optional
	GlobalFilter []GlobalFilterParameters `json:"globalFilter,omitempty" tf:"global_filter,omitempty"`

	// Name of the patch baseline.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system the patch baseline applies to. Valid values are ALMA_LINUX, AMAZON_LINUX, AMAZON_LINUX_2, AMAZON_LINUX_2022, AMAZON_LINUX_2023, CENTOS, DEBIAN, MACOS, ORACLE_LINUX, RASPBIAN, REDHAT_ENTERPRISE_LINUX, ROCKY_LINUX, SUSE, UBUNTU, and WINDOWS. The default value is WINDOWS.
	// +kubebuilder:validation:Optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of rejected patches.
	// +kubebuilder:validation:Optional
	// +listType=set
	RejectedPatches []*string `json:"rejectedPatches,omitempty" tf:"rejected_patches,omitempty"`

	// Action for Patch Manager to take on patches included in the rejected_patches list. Valid values are ALLOW_AS_DEPENDENCY and BLOCK.
	// +kubebuilder:validation:Optional
	RejectedPatchesAction *string `json:"rejectedPatchesAction,omitempty" tf:"rejected_patches_action,omitempty"`

	// Configuration block with alternate sources for patches. Applies to Linux instances only. See source below.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PatchFilterInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type PatchFilterObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type PatchFilterParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type SourceInitParameters struct {

	// Value of the yum repo configuration. For information about other options available for your yum repository configuration, see the dnf.conf documentation
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Name specified to identify the patch source.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specific operating system versions a patch repository applies to, such as "Ubuntu16.04", "AmazonLinux2016.09", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product values, see PatchFilter.
	Products []*string `json:"products,omitempty" tf:"products,omitempty"`
}

type SourceObservation struct {

	// Value of the yum repo configuration. For information about other options available for your yum repository configuration, see the dnf.conf documentation
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Name specified to identify the patch source.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specific operating system versions a patch repository applies to, such as "Ubuntu16.04", "AmazonLinux2016.09", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product values, see PatchFilter.
	Products []*string `json:"products,omitempty" tf:"products,omitempty"`
}

type SourceParameters struct {

	// Value of the yum repo configuration. For information about other options available for your yum repository configuration, see the dnf.conf documentation
	// +kubebuilder:validation:Optional
	Configuration *string `json:"configuration" tf:"configuration,omitempty"`

	// Name specified to identify the patch source.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specific operating system versions a patch repository applies to, such as "Ubuntu16.04", "AmazonLinux2016.09", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product values, see PatchFilter.
	// +kubebuilder:validation:Optional
	Products []*string `json:"products" tf:"products,omitempty"`
}

// PatchBaselineSpec defines the desired state of PatchBaseline
type PatchBaselineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PatchBaselineParameters `json:"forProvider"`
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
	InitProvider PatchBaselineInitParameters `json:"initProvider,omitempty"`
}

// PatchBaselineStatus defines the observed state of PatchBaseline.
type PatchBaselineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PatchBaselineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PatchBaseline is the Schema for the PatchBaselines API. Provides an SSM Patch Baseline resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type PatchBaseline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PatchBaselineSpec   `json:"spec"`
	Status PatchBaselineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PatchBaselineList contains a list of PatchBaselines
type PatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PatchBaseline `json:"items"`
}

// Repository type metadata.
var (
	PatchBaseline_Kind             = "PatchBaseline"
	PatchBaseline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PatchBaseline_Kind}.String()
	PatchBaseline_KindAPIVersion   = PatchBaseline_Kind + "." + CRDGroupVersion.String()
	PatchBaseline_GroupVersionKind = CRDGroupVersion.WithKind(PatchBaseline_Kind)
)

func init() {
	SchemeBuilder.Register(&PatchBaseline{}, &PatchBaselineList{})
}
