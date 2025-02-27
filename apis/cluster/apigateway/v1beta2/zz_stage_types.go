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

type AccessLogSettingsInitParameters struct {

	// ARN of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-. Automatically removes trailing :* if present.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// Formatting and values recorded in the logs.
	// For more information on configuring the log format rules visit the AWS documentation
	Format *string `json:"format,omitempty" tf:"format,omitempty"`
}

type AccessLogSettingsObservation struct {

	// ARN of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-. Automatically removes trailing :* if present.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// Formatting and values recorded in the logs.
	// For more information on configuring the log format rules visit the AWS documentation
	Format *string `json:"format,omitempty" tf:"format,omitempty"`
}

type AccessLogSettingsParameters struct {

	// ARN of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-. Automatically removes trailing :* if present.
	// +kubebuilder:validation:Optional
	DestinationArn *string `json:"destinationArn" tf:"destination_arn,omitempty"`

	// Formatting and values recorded in the logs.
	// For more information on configuring the log format rules visit the AWS documentation
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`
}

type CanarySettingsInitParameters struct {

	// ID of the deployment that the stage points to
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Percent 0.0 - 100.0 of traffic to divert to the canary deployment.
	PercentTraffic *float64 `json:"percentTraffic,omitempty" tf:"percent_traffic,omitempty"`

	// Map of overridden stage variables (including new variables) for the canary deployment.
	// +mapType=granular
	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty" tf:"stage_variable_overrides,omitempty"`

	// Whether the canary deployment uses the stage cache. Defaults to false.
	UseStageCache *bool `json:"useStageCache,omitempty" tf:"use_stage_cache,omitempty"`
}

type CanarySettingsObservation struct {

	// ID of the deployment that the stage points to
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Percent 0.0 - 100.0 of traffic to divert to the canary deployment.
	PercentTraffic *float64 `json:"percentTraffic,omitempty" tf:"percent_traffic,omitempty"`

	// Map of overridden stage variables (including new variables) for the canary deployment.
	// +mapType=granular
	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty" tf:"stage_variable_overrides,omitempty"`

	// Whether the canary deployment uses the stage cache. Defaults to false.
	UseStageCache *bool `json:"useStageCache,omitempty" tf:"use_stage_cache,omitempty"`
}

type CanarySettingsParameters struct {

	// ID of the deployment that the stage points to
	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// Percent 0.0 - 100.0 of traffic to divert to the canary deployment.
	// +kubebuilder:validation:Optional
	PercentTraffic *float64 `json:"percentTraffic,omitempty" tf:"percent_traffic,omitempty"`

	// Map of overridden stage variables (including new variables) for the canary deployment.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty" tf:"stage_variable_overrides,omitempty"`

	// Whether the canary deployment uses the stage cache. Defaults to false.
	// +kubebuilder:validation:Optional
	UseStageCache *bool `json:"useStageCache,omitempty" tf:"use_stage_cache,omitempty"`
}

type StageInitParameters struct {

	// Enables access logs for the API stage. See Access Log Settings below.
	AccessLogSettings *AccessLogSettingsInitParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// Whether a cache cluster is enabled for the stage
	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty" tf:"cache_cluster_enabled,omitempty"`

	// Size of the cache cluster for the stage, if enabled. Allowed values include 0.5, 1.6, 6.1, 13.5, 28.4, 58.2, 118 and 237.
	CacheClusterSize *string `json:"cacheClusterSize,omitempty" tf:"cache_cluster_size,omitempty"`

	// Configuration settings of a canary deployment. See Canary Settings below.
	CanarySettings *CanarySettingsInitParameters `json:"canarySettings,omitempty" tf:"canary_settings,omitempty"`

	// Identifier of a client certificate for the stage.
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// ID of the deployment that the stage points to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/apigateway/v1beta1.Deployment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Reference to a Deployment in apigateway to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDRef *v1.Reference `json:"deploymentIdRef,omitempty" tf:"-"`

	// Selector for a Deployment in apigateway to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDSelector *v1.Selector `json:"deploymentIdSelector,omitempty" tf:"-"`

	// Description of the stage.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Version of the associated API documentation
	DocumentationVersion *string `json:"documentationVersion,omitempty" tf:"documentation_version,omitempty"`

	// ID of the associated REST API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/apigateway/v1beta2.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// Name of the stage
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map that defines the stage variables
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`

	// Whether active tracing with X-ray is enabled. Defaults to false.
	XrayTracingEnabled *bool `json:"xrayTracingEnabled,omitempty" tf:"xray_tracing_enabled,omitempty"`
}

type StageObservation struct {

	// Enables access logs for the API stage. See Access Log Settings below.
	AccessLogSettings *AccessLogSettingsObservation `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Whether a cache cluster is enabled for the stage
	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty" tf:"cache_cluster_enabled,omitempty"`

	// Size of the cache cluster for the stage, if enabled. Allowed values include 0.5, 1.6, 6.1, 13.5, 28.4, 58.2, 118 and 237.
	CacheClusterSize *string `json:"cacheClusterSize,omitempty" tf:"cache_cluster_size,omitempty"`

	// Configuration settings of a canary deployment. See Canary Settings below.
	CanarySettings *CanarySettingsObservation `json:"canarySettings,omitempty" tf:"canary_settings,omitempty"`

	// Identifier of a client certificate for the stage.
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// ID of the deployment that the stage points to
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Description of the stage.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Version of the associated API documentation
	DocumentationVersion *string `json:"documentationVersion,omitempty" tf:"documentation_version,omitempty"`

	// Execution ARN to be used in lambda_permission's source_arn
	// when allowing API Gateway to invoke a Lambda function,
	// e.g., arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod
	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	// ID of the stage
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URL to invoke the API pointing to the stage,
	// e.g., https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod
	InvokeURL *string `json:"invokeUrl,omitempty" tf:"invoke_url,omitempty"`

	// ID of the associated REST API
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Name of the stage
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Map that defines the stage variables
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`

	// ARN of the WebAcl associated with the Stage.
	WebACLArn *string `json:"webAclArn,omitempty" tf:"web_acl_arn,omitempty"`

	// Whether active tracing with X-ray is enabled. Defaults to false.
	XrayTracingEnabled *bool `json:"xrayTracingEnabled,omitempty" tf:"xray_tracing_enabled,omitempty"`
}

type StageParameters struct {

	// Enables access logs for the API stage. See Access Log Settings below.
	// +kubebuilder:validation:Optional
	AccessLogSettings *AccessLogSettingsParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// Whether a cache cluster is enabled for the stage
	// +kubebuilder:validation:Optional
	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty" tf:"cache_cluster_enabled,omitempty"`

	// Size of the cache cluster for the stage, if enabled. Allowed values include 0.5, 1.6, 6.1, 13.5, 28.4, 58.2, 118 and 237.
	// +kubebuilder:validation:Optional
	CacheClusterSize *string `json:"cacheClusterSize,omitempty" tf:"cache_cluster_size,omitempty"`

	// Configuration settings of a canary deployment. See Canary Settings below.
	// +kubebuilder:validation:Optional
	CanarySettings *CanarySettingsParameters `json:"canarySettings,omitempty" tf:"canary_settings,omitempty"`

	// Identifier of a client certificate for the stage.
	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// ID of the deployment that the stage points to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/apigateway/v1beta1.Deployment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Reference to a Deployment in apigateway to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDRef *v1.Reference `json:"deploymentIdRef,omitempty" tf:"-"`

	// Selector for a Deployment in apigateway to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDSelector *v1.Selector `json:"deploymentIdSelector,omitempty" tf:"-"`

	// Description of the stage.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Version of the associated API documentation
	// +kubebuilder:validation:Optional
	DocumentationVersion *string `json:"documentationVersion,omitempty" tf:"documentation_version,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the associated REST API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/apigateway/v1beta2.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// Name of the stage
	// +kubebuilder:validation:Optional
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map that defines the stage variables
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`

	// Whether active tracing with X-ray is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	XrayTracingEnabled *bool `json:"xrayTracingEnabled,omitempty" tf:"xray_tracing_enabled,omitempty"`
}

// StageSpec defines the desired state of Stage
type StageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StageParameters `json:"forProvider"`
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
	InitProvider StageInitParameters `json:"initProvider,omitempty"`
}

// StageStatus defines the observed state of Stage.
type StageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Stage is the Schema for the Stages API. Manages an API Gateway Stage.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stageName) || (has(self.initProvider) && has(self.initProvider.stageName))",message="spec.forProvider.stageName is a required parameter"
	Spec   StageSpec   `json:"spec"`
	Status StageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StageList contains a list of Stages
type StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stage `json:"items"`
}

// Repository type metadata.
var (
	Stage_Kind             = "Stage"
	Stage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stage_Kind}.String()
	Stage_KindAPIVersion   = Stage_Kind + "." + CRDGroupVersion.String()
	Stage_GroupVersionKind = CRDGroupVersion.WithKind(Stage_Kind)
)

func init() {
	SchemeBuilder.Register(&Stage{}, &StageList{})
}
