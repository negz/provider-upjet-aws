// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CachingConfigInitParameters struct {

	// The caching keys for a resolver that has caching activated. Valid values are entries from the $context.arguments, $context.source, and $context.identity maps.
	// +listType=set
	CachingKeys []*string `json:"cachingKeys,omitempty" tf:"caching_keys,omitempty"`

	// The TTL in seconds for a resolver that has caching activated. Valid values are between 1 and 3600 seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type CachingConfigObservation struct {

	// The caching keys for a resolver that has caching activated. Valid values are entries from the $context.arguments, $context.source, and $context.identity maps.
	// +listType=set
	CachingKeys []*string `json:"cachingKeys,omitempty" tf:"caching_keys,omitempty"`

	// The TTL in seconds for a resolver that has caching activated. Valid values are between 1 and 3600 seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type CachingConfigParameters struct {

	// The caching keys for a resolver that has caching activated. Valid values are entries from the $context.arguments, $context.source, and $context.identity maps.
	// +kubebuilder:validation:Optional
	// +listType=set
	CachingKeys []*string `json:"cachingKeys,omitempty" tf:"caching_keys,omitempty"`

	// The TTL in seconds for a resolver that has caching activated. Valid values are between 1 and 3600 seconds.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type PipelineConfigInitParameters struct {

	// A list of Function objects.
	Functions []*string `json:"functions,omitempty" tf:"functions,omitempty"`
}

type PipelineConfigObservation struct {

	// A list of Function objects.
	Functions []*string `json:"functions,omitempty" tf:"functions,omitempty"`
}

type PipelineConfigParameters struct {

	// A list of Function objects.
	// +kubebuilder:validation:Optional
	Functions []*string `json:"functions,omitempty" tf:"functions,omitempty"`
}

type ResolverInitParameters struct {

	// The Caching Config. See Caching Config.
	CachingConfig []CachingConfigInitParameters `json:"cachingConfig,omitempty" tf:"caching_config,omitempty"`

	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Data source name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta1.Datasource
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Reference to a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceRef *v1.Reference `json:"dataSourceRef,omitempty" tf:"-"`

	// Selector for a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceSelector *v1.Selector `json:"dataSourceSelector,omitempty" tf:"-"`

	// Resolver type. Valid values are UNIT and PIPELINE.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Maximum batching size for a resolver. Valid values are between 0 and 2000.
	MaxBatchSize *float64 `json:"maxBatchSize,omitempty" tf:"max_batch_size,omitempty"`

	// The caching configuration for the resolver. See Pipeline Config.
	PipelineConfig []PipelineConfigInitParameters `json:"pipelineConfig,omitempty" tf:"pipeline_config,omitempty"`

	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	RequestTemplate *string `json:"requestTemplate,omitempty" tf:"request_template,omitempty"`

	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	ResponseTemplate *string `json:"responseTemplate,omitempty" tf:"response_template,omitempty"`

	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime []ResolverRuntimeInitParameters `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig []ResolverSyncConfigInitParameters `json:"syncConfig,omitempty" tf:"sync_config,omitempty"`
}

type ResolverObservation struct {

	// API ID for the GraphQL API.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Caching Config. See Caching Config.
	CachingConfig []CachingConfigObservation `json:"cachingConfig,omitempty" tf:"caching_config,omitempty"`

	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Data source name.
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Field name from the schema defined in the GraphQL API.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resolver type. Valid values are UNIT and PIPELINE.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Maximum batching size for a resolver. Valid values are between 0 and 2000.
	MaxBatchSize *float64 `json:"maxBatchSize,omitempty" tf:"max_batch_size,omitempty"`

	// The caching configuration for the resolver. See Pipeline Config.
	PipelineConfig []PipelineConfigObservation `json:"pipelineConfig,omitempty" tf:"pipeline_config,omitempty"`

	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	RequestTemplate *string `json:"requestTemplate,omitempty" tf:"request_template,omitempty"`

	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	ResponseTemplate *string `json:"responseTemplate,omitempty" tf:"response_template,omitempty"`

	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime []ResolverRuntimeObservation `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig []ResolverSyncConfigObservation `json:"syncConfig,omitempty" tf:"sync_config,omitempty"`

	// Type name from the schema defined in the GraphQL API.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResolverParameters struct {

	// API ID for the GraphQL API.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta1.GraphQLAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// The Caching Config. See Caching Config.
	// +kubebuilder:validation:Optional
	CachingConfig []CachingConfigParameters `json:"cachingConfig,omitempty" tf:"caching_config,omitempty"`

	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	// +kubebuilder:validation:Optional
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Data source name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta1.Datasource
	// +kubebuilder:validation:Optional
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Reference to a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceRef *v1.Reference `json:"dataSourceRef,omitempty" tf:"-"`

	// Selector for a Datasource in appsync to populate dataSource.
	// +kubebuilder:validation:Optional
	DataSourceSelector *v1.Selector `json:"dataSourceSelector,omitempty" tf:"-"`

	// Field name from the schema defined in the GraphQL API.
	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// Resolver type. Valid values are UNIT and PIPELINE.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Maximum batching size for a resolver. Valid values are between 0 and 2000.
	// +kubebuilder:validation:Optional
	MaxBatchSize *float64 `json:"maxBatchSize,omitempty" tf:"max_batch_size,omitempty"`

	// The caching configuration for the resolver. See Pipeline Config.
	// +kubebuilder:validation:Optional
	PipelineConfig []PipelineConfigParameters `json:"pipelineConfig,omitempty" tf:"pipeline_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	// +kubebuilder:validation:Optional
	RequestTemplate *string `json:"requestTemplate,omitempty" tf:"request_template,omitempty"`

	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	// +kubebuilder:validation:Optional
	ResponseTemplate *string `json:"responseTemplate,omitempty" tf:"response_template,omitempty"`

	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	// +kubebuilder:validation:Optional
	Runtime []ResolverRuntimeParameters `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Describes a Sync configuration for a resolver. See Sync Config.
	// +kubebuilder:validation:Optional
	SyncConfig []ResolverSyncConfigParameters `json:"syncConfig,omitempty" tf:"sync_config,omitempty"`

	// Type name from the schema defined in the GraphQL API.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ResolverRuntimeInitParameters struct {

	// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The version of the runtime to use. Currently, the only allowed version is 1.0.0.
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
}

type ResolverRuntimeObservation struct {

	// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The version of the runtime to use. Currently, the only allowed version is 1.0.0.
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
}

type ResolverRuntimeParameters struct {

	// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The version of the runtime to use. Currently, the only allowed version is 1.0.0.
	// +kubebuilder:validation:Optional
	RuntimeVersion *string `json:"runtimeVersion" tf:"runtime_version,omitempty"`
}

type ResolverSyncConfigInitParameters struct {

	// Conflict Detection strategy to use. Valid values are NONE and VERSION.
	ConflictDetection *string `json:"conflictDetection,omitempty" tf:"conflict_detection,omitempty"`

	// Conflict Resolution strategy to perform in the event of a conflict. Valid values are NONE, OPTIMISTIC_CONCURRENCY, AUTOMERGE, and LAMBDA.
	ConflictHandler *string `json:"conflictHandler,omitempty" tf:"conflict_handler,omitempty"`

	// Lambda Conflict Handler Config when configuring LAMBDA as the Conflict Handler. See Lambda Conflict Handler Config.
	LambdaConflictHandlerConfig []SyncConfigLambdaConflictHandlerConfigInitParameters `json:"lambdaConflictHandlerConfig,omitempty" tf:"lambda_conflict_handler_config,omitempty"`
}

type ResolverSyncConfigObservation struct {

	// Conflict Detection strategy to use. Valid values are NONE and VERSION.
	ConflictDetection *string `json:"conflictDetection,omitempty" tf:"conflict_detection,omitempty"`

	// Conflict Resolution strategy to perform in the event of a conflict. Valid values are NONE, OPTIMISTIC_CONCURRENCY, AUTOMERGE, and LAMBDA.
	ConflictHandler *string `json:"conflictHandler,omitempty" tf:"conflict_handler,omitempty"`

	// Lambda Conflict Handler Config when configuring LAMBDA as the Conflict Handler. See Lambda Conflict Handler Config.
	LambdaConflictHandlerConfig []SyncConfigLambdaConflictHandlerConfigObservation `json:"lambdaConflictHandlerConfig,omitempty" tf:"lambda_conflict_handler_config,omitempty"`
}

type ResolverSyncConfigParameters struct {

	// Conflict Detection strategy to use. Valid values are NONE and VERSION.
	// +kubebuilder:validation:Optional
	ConflictDetection *string `json:"conflictDetection,omitempty" tf:"conflict_detection,omitempty"`

	// Conflict Resolution strategy to perform in the event of a conflict. Valid values are NONE, OPTIMISTIC_CONCURRENCY, AUTOMERGE, and LAMBDA.
	// +kubebuilder:validation:Optional
	ConflictHandler *string `json:"conflictHandler,omitempty" tf:"conflict_handler,omitempty"`

	// Lambda Conflict Handler Config when configuring LAMBDA as the Conflict Handler. See Lambda Conflict Handler Config.
	// +kubebuilder:validation:Optional
	LambdaConflictHandlerConfig []SyncConfigLambdaConflictHandlerConfigParameters `json:"lambdaConflictHandlerConfig,omitempty" tf:"lambda_conflict_handler_config,omitempty"`
}

type SyncConfigLambdaConflictHandlerConfigInitParameters struct {

	// ARN for the Lambda function to use as the Conflict Handler.
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn,omitempty" tf:"lambda_conflict_handler_arn,omitempty"`
}

type SyncConfigLambdaConflictHandlerConfigObservation struct {

	// ARN for the Lambda function to use as the Conflict Handler.
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn,omitempty" tf:"lambda_conflict_handler_arn,omitempty"`
}

type SyncConfigLambdaConflictHandlerConfigParameters struct {

	// ARN for the Lambda function to use as the Conflict Handler.
	// +kubebuilder:validation:Optional
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn,omitempty" tf:"lambda_conflict_handler_arn,omitempty"`
}

// ResolverSpec defines the desired state of Resolver
type ResolverSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResolverParameters `json:"forProvider"`
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
	InitProvider ResolverInitParameters `json:"initProvider,omitempty"`
}

// ResolverStatus defines the observed state of Resolver.
type ResolverStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResolverObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Resolver is the Schema for the Resolvers API. Provides an AppSync Resolver.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Resolver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverSpec   `json:"spec"`
	Status            ResolverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverList contains a list of Resolvers
type ResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resolver `json:"items"`
}

// Repository type metadata.
var (
	Resolver_Kind             = "Resolver"
	Resolver_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Resolver_Kind}.String()
	Resolver_KindAPIVersion   = Resolver_Kind + "." + CRDGroupVersion.String()
	Resolver_GroupVersionKind = CRDGroupVersion.WithKind(Resolver_Kind)
)

func init() {
	SchemeBuilder.Register(&Resolver{}, &ResolverList{})
}
