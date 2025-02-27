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

type ReportDefinitionInitParameters struct {

	// A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT, ATHENA. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be OVERWRITE_REPORT.
	// +listType=set
	AdditionalArtifacts []*string `json:"additionalArtifacts,omitempty" tf:"additional_artifacts,omitempty"`

	// A list of schema elements. Valid values are: RESOURCES, SPLIT_COST_ALLOCATION_DATA.
	// +listType=set
	AdditionalSchemaElements []*string `json:"additionalSchemaElements,omitempty" tf:"additional_schema_elements,omitempty"`

	// Compression format for report. Valid values are: GZIP, ZIP, Parquet. If Parquet is used, then format must also be Parquet.
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// Format for report. Valid values are: textORcsv, Parquet. If Parquet is used, then Compression must also be Parquet.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
	RefreshClosedReports *bool `json:"refreshClosedReports,omitempty" tf:"refresh_closed_reports,omitempty"`

	// Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: CREATE_NEW_REPORT and OVERWRITE_REPORT.
	ReportVersioning *string `json:"reportVersioning,omitempty" tf:"report_versioning,omitempty"`

	// Name of the existing S3 bucket to hold generated reports.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// Reference to a Bucket in s3 to populate s3Bucket.
	// +kubebuilder:validation:Optional
	S3BucketRef *v1.Reference `json:"s3BucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3Bucket.
	// +kubebuilder:validation:Optional
	S3BucketSelector *v1.Selector `json:"s3BucketSelector,omitempty" tf:"-"`

	// Report path prefix. Limited to 256 characters.
	S3Prefix *string `json:"s3Prefix,omitempty" tf:"s3_prefix,omitempty"`

	// Region of the existing S3 bucket to hold generated reports.
	S3Region *string `json:"s3Region,omitempty" tf:"s3_region,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The frequency on which report data are measured and displayed.  Valid values are: DAILY, HOURLY, MONTHLY.
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type ReportDefinitionObservation struct {

	// A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT, ATHENA. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be OVERWRITE_REPORT.
	// +listType=set
	AdditionalArtifacts []*string `json:"additionalArtifacts,omitempty" tf:"additional_artifacts,omitempty"`

	// A list of schema elements. Valid values are: RESOURCES, SPLIT_COST_ALLOCATION_DATA.
	// +listType=set
	AdditionalSchemaElements []*string `json:"additionalSchemaElements,omitempty" tf:"additional_schema_elements,omitempty"`

	// The Amazon Resource Name (ARN) specifying the cur report.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Compression format for report. Valid values are: GZIP, ZIP, Parquet. If Parquet is used, then format must also be Parquet.
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// Format for report. Valid values are: textORcsv, Parquet. If Parquet is used, then Compression must also be Parquet.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
	RefreshClosedReports *bool `json:"refreshClosedReports,omitempty" tf:"refresh_closed_reports,omitempty"`

	// Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: CREATE_NEW_REPORT and OVERWRITE_REPORT.
	ReportVersioning *string `json:"reportVersioning,omitempty" tf:"report_versioning,omitempty"`

	// Name of the existing S3 bucket to hold generated reports.
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// Report path prefix. Limited to 256 characters.
	S3Prefix *string `json:"s3Prefix,omitempty" tf:"s3_prefix,omitempty"`

	// Region of the existing S3 bucket to hold generated reports.
	S3Region *string `json:"s3Region,omitempty" tf:"s3_region,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The frequency on which report data are measured and displayed.  Valid values are: DAILY, HOURLY, MONTHLY.
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type ReportDefinitionParameters struct {

	// A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT, ATHENA. When ATHENA exists within additional_artifacts, no other artifact type can be declared and report_versioning must be OVERWRITE_REPORT.
	// +kubebuilder:validation:Optional
	// +listType=set
	AdditionalArtifacts []*string `json:"additionalArtifacts,omitempty" tf:"additional_artifacts,omitempty"`

	// A list of schema elements. Valid values are: RESOURCES, SPLIT_COST_ALLOCATION_DATA.
	// +kubebuilder:validation:Optional
	// +listType=set
	AdditionalSchemaElements []*string `json:"additionalSchemaElements,omitempty" tf:"additional_schema_elements,omitempty"`

	// Compression format for report. Valid values are: GZIP, ZIP, Parquet. If Parquet is used, then format must also be Parquet.
	// +kubebuilder:validation:Optional
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// Format for report. Valid values are: textORcsv, Parquet. If Parquet is used, then Compression must also be Parquet.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
	// +kubebuilder:validation:Optional
	RefreshClosedReports *bool `json:"refreshClosedReports,omitempty" tf:"refresh_closed_reports,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: CREATE_NEW_REPORT and OVERWRITE_REPORT.
	// +kubebuilder:validation:Optional
	ReportVersioning *string `json:"reportVersioning,omitempty" tf:"report_versioning,omitempty"`

	// Name of the existing S3 bucket to hold generated reports.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +kubebuilder:validation:Optional
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// Reference to a Bucket in s3 to populate s3Bucket.
	// +kubebuilder:validation:Optional
	S3BucketRef *v1.Reference `json:"s3BucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3Bucket.
	// +kubebuilder:validation:Optional
	S3BucketSelector *v1.Selector `json:"s3BucketSelector,omitempty" tf:"-"`

	// Report path prefix. Limited to 256 characters.
	// +kubebuilder:validation:Optional
	S3Prefix *string `json:"s3Prefix,omitempty" tf:"s3_prefix,omitempty"`

	// Region of the existing S3 bucket to hold generated reports.
	// +kubebuilder:validation:Optional
	S3Region *string `json:"s3Region,omitempty" tf:"s3_region,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The frequency on which report data are measured and displayed.  Valid values are: DAILY, HOURLY, MONTHLY.
	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

// ReportDefinitionSpec defines the desired state of ReportDefinition
type ReportDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReportDefinitionParameters `json:"forProvider"`
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
	InitProvider ReportDefinitionInitParameters `json:"initProvider,omitempty"`
}

// ReportDefinitionStatus defines the observed state of ReportDefinition.
type ReportDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReportDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ReportDefinition is the Schema for the ReportDefinitions API. Provides a Cost and Usage Report Definition.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type ReportDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.additionalSchemaElements) || (has(self.initProvider) && has(self.initProvider.additionalSchemaElements))",message="spec.forProvider.additionalSchemaElements is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.compression) || (has(self.initProvider) && has(self.initProvider.compression))",message="spec.forProvider.compression is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.format) || (has(self.initProvider) && has(self.initProvider.format))",message="spec.forProvider.format is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.s3Region) || (has(self.initProvider) && has(self.initProvider.s3Region))",message="spec.forProvider.s3Region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeUnit) || (has(self.initProvider) && has(self.initProvider.timeUnit))",message="spec.forProvider.timeUnit is a required parameter"
	Spec   ReportDefinitionSpec   `json:"spec"`
	Status ReportDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReportDefinitionList contains a list of ReportDefinitions
type ReportDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReportDefinition `json:"items"`
}

// Repository type metadata.
var (
	ReportDefinition_Kind             = "ReportDefinition"
	ReportDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReportDefinition_Kind}.String()
	ReportDefinition_KindAPIVersion   = ReportDefinition_Kind + "." + CRDGroupVersion.String()
	ReportDefinition_GroupVersionKind = CRDGroupVersion.WithKind(ReportDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&ReportDefinition{}, &ReportDefinitionList{})
}
