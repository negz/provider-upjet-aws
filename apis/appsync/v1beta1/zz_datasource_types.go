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

type AuthorizationConfigInitParameters struct {

	// Authorization type that the HTTP endpoint requires. Default values is AWS_IAM.
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Identity and Access Management (IAM) settings. See AWS IAM Config.
	AwsIAMConfig []AwsIAMConfigInitParameters `json:"awsIamConfig,omitempty" tf:"aws_iam_config,omitempty"`
}

type AuthorizationConfigObservation struct {

	// Authorization type that the HTTP endpoint requires. Default values is AWS_IAM.
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Identity and Access Management (IAM) settings. See AWS IAM Config.
	AwsIAMConfig []AwsIAMConfigObservation `json:"awsIamConfig,omitempty" tf:"aws_iam_config,omitempty"`
}

type AuthorizationConfigParameters struct {

	// Authorization type that the HTTP endpoint requires. Default values is AWS_IAM.
	// +kubebuilder:validation:Optional
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Identity and Access Management (IAM) settings. See AWS IAM Config.
	// +kubebuilder:validation:Optional
	AwsIAMConfig []AwsIAMConfigParameters `json:"awsIamConfig,omitempty" tf:"aws_iam_config,omitempty"`
}

type AwsIAMConfigInitParameters struct {

	// Signing Amazon Web Services Region for IAM authorization.
	SigningRegion *string `json:"signingRegion,omitempty" tf:"signing_region,omitempty"`

	// Signing service name for IAM authorization.
	SigningServiceName *string `json:"signingServiceName,omitempty" tf:"signing_service_name,omitempty"`
}

type AwsIAMConfigObservation struct {

	// Signing Amazon Web Services Region for IAM authorization.
	SigningRegion *string `json:"signingRegion,omitempty" tf:"signing_region,omitempty"`

	// Signing service name for IAM authorization.
	SigningServiceName *string `json:"signingServiceName,omitempty" tf:"signing_service_name,omitempty"`
}

type AwsIAMConfigParameters struct {

	// Signing Amazon Web Services Region for IAM authorization.
	// +kubebuilder:validation:Optional
	SigningRegion *string `json:"signingRegion,omitempty" tf:"signing_region,omitempty"`

	// Signing service name for IAM authorization.
	// +kubebuilder:validation:Optional
	SigningServiceName *string `json:"signingServiceName,omitempty" tf:"signing_service_name,omitempty"`
}

type DatasourceInitParameters struct {

	// Description of the data source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// DynamoDB settings. See DynamoDB Config
	DynamodbConfig []DynamodbConfigInitParameters `json:"dynamodbConfig,omitempty" tf:"dynamodb_config,omitempty"`

	// Amazon Elasticsearch settings. See ElasticSearch Config
	ElasticsearchConfig []ElasticsearchConfigInitParameters `json:"elasticsearchConfig,omitempty" tf:"elasticsearch_config,omitempty"`

	// AWS EventBridge settings. See Event Bridge Config
	EventBridgeConfig []EventBridgeConfigInitParameters `json:"eventBridgeConfig,omitempty" tf:"event_bridge_config,omitempty"`

	// HTTP settings. See HTTP Config
	HTTPConfig []HTTPConfigInitParameters `json:"httpConfig,omitempty" tf:"http_config,omitempty"`

	// AWS Lambda settings. See Lambda Config
	LambdaConfig []LambdaConfigInitParameters `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`

	// Amazon OpenSearch Service settings. See OpenSearch Service Config
	OpensearchserviceConfig []OpensearchserviceConfigInitParameters `json:"opensearchserviceConfig,omitempty" tf:"opensearchservice_config,omitempty"`

	// AWS RDS settings. See Relational Database Config
	RelationalDatabaseConfig []RelationalDatabaseConfigInitParameters `json:"relationalDatabaseConfig,omitempty" tf:"relational_database_config,omitempty"`

	// IAM service role ARN for the data source. Required if type is specified as AWS_LAMBDA, AMAZON_DYNAMODB, AMAZON_ELASTICSEARCH, AMAZON_EVENTBRIDGE, or AMAZON_OPENSEARCH_SERVICE.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// Type of the Data Source. Valid values: AWS_LAMBDA, AMAZON_DYNAMODB, AMAZON_ELASTICSEARCH, HTTP, NONE, RELATIONAL_DATABASE, AMAZON_EVENTBRIDGE, AMAZON_OPENSEARCH_SERVICE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DatasourceObservation struct {

	// API ID for the GraphQL API for the data source.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the data source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// DynamoDB settings. See DynamoDB Config
	DynamodbConfig []DynamodbConfigObservation `json:"dynamodbConfig,omitempty" tf:"dynamodb_config,omitempty"`

	// Amazon Elasticsearch settings. See ElasticSearch Config
	ElasticsearchConfig []ElasticsearchConfigObservation `json:"elasticsearchConfig,omitempty" tf:"elasticsearch_config,omitempty"`

	// AWS EventBridge settings. See Event Bridge Config
	EventBridgeConfig []EventBridgeConfigObservation `json:"eventBridgeConfig,omitempty" tf:"event_bridge_config,omitempty"`

	// HTTP settings. See HTTP Config
	HTTPConfig []HTTPConfigObservation `json:"httpConfig,omitempty" tf:"http_config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// AWS Lambda settings. See Lambda Config
	LambdaConfig []LambdaConfigObservation `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`

	// Amazon OpenSearch Service settings. See OpenSearch Service Config
	OpensearchserviceConfig []OpensearchserviceConfigObservation `json:"opensearchserviceConfig,omitempty" tf:"opensearchservice_config,omitempty"`

	// AWS RDS settings. See Relational Database Config
	RelationalDatabaseConfig []RelationalDatabaseConfigObservation `json:"relationalDatabaseConfig,omitempty" tf:"relational_database_config,omitempty"`

	// IAM service role ARN for the data source. Required if type is specified as AWS_LAMBDA, AMAZON_DYNAMODB, AMAZON_ELASTICSEARCH, AMAZON_EVENTBRIDGE, or AMAZON_OPENSEARCH_SERVICE.
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Type of the Data Source. Valid values: AWS_LAMBDA, AMAZON_DYNAMODB, AMAZON_ELASTICSEARCH, HTTP, NONE, RELATIONAL_DATABASE, AMAZON_EVENTBRIDGE, AMAZON_OPENSEARCH_SERVICE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DatasourceParameters struct {

	// API ID for the GraphQL API for the data source.
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

	// Description of the data source.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// DynamoDB settings. See DynamoDB Config
	// +kubebuilder:validation:Optional
	DynamodbConfig []DynamodbConfigParameters `json:"dynamodbConfig,omitempty" tf:"dynamodb_config,omitempty"`

	// Amazon Elasticsearch settings. See ElasticSearch Config
	// +kubebuilder:validation:Optional
	ElasticsearchConfig []ElasticsearchConfigParameters `json:"elasticsearchConfig,omitempty" tf:"elasticsearch_config,omitempty"`

	// AWS EventBridge settings. See Event Bridge Config
	// +kubebuilder:validation:Optional
	EventBridgeConfig []EventBridgeConfigParameters `json:"eventBridgeConfig,omitempty" tf:"event_bridge_config,omitempty"`

	// HTTP settings. See HTTP Config
	// +kubebuilder:validation:Optional
	HTTPConfig []HTTPConfigParameters `json:"httpConfig,omitempty" tf:"http_config,omitempty"`

	// AWS Lambda settings. See Lambda Config
	// +kubebuilder:validation:Optional
	LambdaConfig []LambdaConfigParameters `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`

	// Amazon OpenSearch Service settings. See OpenSearch Service Config
	// +kubebuilder:validation:Optional
	OpensearchserviceConfig []OpensearchserviceConfigParameters `json:"opensearchserviceConfig,omitempty" tf:"opensearchservice_config,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// AWS RDS settings. See Relational Database Config
	// +kubebuilder:validation:Optional
	RelationalDatabaseConfig []RelationalDatabaseConfigParameters `json:"relationalDatabaseConfig,omitempty" tf:"relational_database_config,omitempty"`

	// IAM service role ARN for the data source. Required if type is specified as AWS_LAMBDA, AMAZON_DYNAMODB, AMAZON_ELASTICSEARCH, AMAZON_EVENTBRIDGE, or AMAZON_OPENSEARCH_SERVICE.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// Type of the Data Source. Valid values: AWS_LAMBDA, AMAZON_DYNAMODB, AMAZON_ELASTICSEARCH, HTTP, NONE, RELATIONAL_DATABASE, AMAZON_EVENTBRIDGE, AMAZON_OPENSEARCH_SERVICE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DeltaSyncConfigInitParameters struct {

	// The number of minutes that an Item is stored in the data source.
	BaseTableTTL *float64 `json:"baseTableTtl,omitempty" tf:"base_table_ttl,omitempty"`

	// The table name.
	DeltaSyncTableName *string `json:"deltaSyncTableName,omitempty" tf:"delta_sync_table_name,omitempty"`

	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.
	DeltaSyncTableTTL *float64 `json:"deltaSyncTableTtl,omitempty" tf:"delta_sync_table_ttl,omitempty"`
}

type DeltaSyncConfigObservation struct {

	// The number of minutes that an Item is stored in the data source.
	BaseTableTTL *float64 `json:"baseTableTtl,omitempty" tf:"base_table_ttl,omitempty"`

	// The table name.
	DeltaSyncTableName *string `json:"deltaSyncTableName,omitempty" tf:"delta_sync_table_name,omitempty"`

	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.
	DeltaSyncTableTTL *float64 `json:"deltaSyncTableTtl,omitempty" tf:"delta_sync_table_ttl,omitempty"`
}

type DeltaSyncConfigParameters struct {

	// The number of minutes that an Item is stored in the data source.
	// +kubebuilder:validation:Optional
	BaseTableTTL *float64 `json:"baseTableTtl,omitempty" tf:"base_table_ttl,omitempty"`

	// The table name.
	// +kubebuilder:validation:Optional
	DeltaSyncTableName *string `json:"deltaSyncTableName" tf:"delta_sync_table_name,omitempty"`

	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.
	// +kubebuilder:validation:Optional
	DeltaSyncTableTTL *float64 `json:"deltaSyncTableTtl,omitempty" tf:"delta_sync_table_ttl,omitempty"`
}

type DynamodbConfigInitParameters struct {

	// The DeltaSyncConfig for a versioned data source. See Delta Sync Config
	DeltaSyncConfig []DeltaSyncConfigInitParameters `json:"deltaSyncConfig,omitempty" tf:"delta_sync_config,omitempty"`

	// Name of the DynamoDB table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dynamodb/v1beta1.Table
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table in dynamodb to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table in dynamodb to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`

	// Set to true to use Amazon Cognito credentials with this data source.
	UseCallerCredentials *bool `json:"useCallerCredentials,omitempty" tf:"use_caller_credentials,omitempty"`

	// Detects Conflict Detection and Resolution with this data source.
	Versioned *bool `json:"versioned,omitempty" tf:"versioned,omitempty"`
}

type DynamodbConfigObservation struct {

	// The DeltaSyncConfig for a versioned data source. See Delta Sync Config
	DeltaSyncConfig []DeltaSyncConfigObservation `json:"deltaSyncConfig,omitempty" tf:"delta_sync_config,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Name of the DynamoDB table.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Set to true to use Amazon Cognito credentials with this data source.
	UseCallerCredentials *bool `json:"useCallerCredentials,omitempty" tf:"use_caller_credentials,omitempty"`

	// Detects Conflict Detection and Resolution with this data source.
	Versioned *bool `json:"versioned,omitempty" tf:"versioned,omitempty"`
}

type DynamodbConfigParameters struct {

	// The DeltaSyncConfig for a versioned data source. See Delta Sync Config
	// +kubebuilder:validation:Optional
	DeltaSyncConfig []DeltaSyncConfigParameters `json:"deltaSyncConfig,omitempty" tf:"delta_sync_config,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Name of the DynamoDB table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dynamodb/v1beta1.Table
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table in dynamodb to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table in dynamodb to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`

	// Set to true to use Amazon Cognito credentials with this data source.
	// +kubebuilder:validation:Optional
	UseCallerCredentials *bool `json:"useCallerCredentials,omitempty" tf:"use_caller_credentials,omitempty"`

	// Detects Conflict Detection and Resolution with this data source.
	// +kubebuilder:validation:Optional
	Versioned *bool `json:"versioned,omitempty" tf:"versioned,omitempty"`
}

type ElasticsearchConfigInitParameters struct {

	// HTTP endpoint of the Elasticsearch domain.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
}

type ElasticsearchConfigObservation struct {

	// HTTP endpoint of the Elasticsearch domain.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ElasticsearchConfigParameters struct {

	// HTTP endpoint of the Elasticsearch domain.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EventBridgeConfigInitParameters struct {

	// ARN for the EventBridge bus.
	EventBusArn *string `json:"eventBusArn,omitempty" tf:"event_bus_arn,omitempty"`
}

type EventBridgeConfigObservation struct {

	// ARN for the EventBridge bus.
	EventBusArn *string `json:"eventBusArn,omitempty" tf:"event_bus_arn,omitempty"`
}

type EventBridgeConfigParameters struct {

	// ARN for the EventBridge bus.
	// +kubebuilder:validation:Optional
	EventBusArn *string `json:"eventBusArn" tf:"event_bus_arn,omitempty"`
}

type HTTPConfigInitParameters struct {

	// Authorization configuration in case the HTTP endpoint requires authorization. See Authorization Config.
	AuthorizationConfig []AuthorizationConfigInitParameters `json:"authorizationConfig,omitempty" tf:"authorization_config,omitempty"`

	// HTTP endpoint of the Elasticsearch domain.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
}

type HTTPConfigObservation struct {

	// Authorization configuration in case the HTTP endpoint requires authorization. See Authorization Config.
	AuthorizationConfig []AuthorizationConfigObservation `json:"authorizationConfig,omitempty" tf:"authorization_config,omitempty"`

	// HTTP endpoint of the Elasticsearch domain.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
}

type HTTPConfigParameters struct {

	// Authorization configuration in case the HTTP endpoint requires authorization. See Authorization Config.
	// +kubebuilder:validation:Optional
	AuthorizationConfig []AuthorizationConfigParameters `json:"authorizationConfig,omitempty" tf:"authorization_config,omitempty"`

	// HTTP endpoint of the Elasticsearch domain.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`
}

type HTTPEndpointConfigInitParameters struct {

	// AWS secret store ARN for database credentials.
	AwsSecretStoreArn *string `json:"awsSecretStoreArn,omitempty" tf:"aws_secret_store_arn,omitempty"`

	// Amazon RDS cluster identifier.
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier,omitempty"`

	// Logical database name.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Logical schema name.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

type HTTPEndpointConfigObservation struct {

	// AWS secret store ARN for database credentials.
	AwsSecretStoreArn *string `json:"awsSecretStoreArn,omitempty" tf:"aws_secret_store_arn,omitempty"`

	// Amazon RDS cluster identifier.
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier,omitempty"`

	// Logical database name.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Logical schema name.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

type HTTPEndpointConfigParameters struct {

	// AWS secret store ARN for database credentials.
	// +kubebuilder:validation:Optional
	AwsSecretStoreArn *string `json:"awsSecretStoreArn" tf:"aws_secret_store_arn,omitempty"`

	// Amazon RDS cluster identifier.
	// +kubebuilder:validation:Optional
	DBClusterIdentifier *string `json:"dbClusterIdentifier" tf:"db_cluster_identifier,omitempty"`

	// Logical database name.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Logical schema name.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

type LambdaConfigInitParameters struct {

	// ARN for the Lambda function.
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`
}

type LambdaConfigObservation struct {

	// ARN for the Lambda function.
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`
}

type LambdaConfigParameters struct {

	// ARN for the Lambda function.
	// +kubebuilder:validation:Optional
	FunctionArn *string `json:"functionArn" tf:"function_arn,omitempty"`
}

type OpensearchserviceConfigInitParameters struct {

	// HTTP endpoint of the Elasticsearch domain.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
}

type OpensearchserviceConfigObservation struct {

	// HTTP endpoint of the Elasticsearch domain.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type OpensearchserviceConfigParameters struct {

	// HTTP endpoint of the Elasticsearch domain.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type RelationalDatabaseConfigInitParameters struct {

	// Amazon RDS HTTP endpoint configuration. See HTTP Endpoint Config.
	HTTPEndpointConfig []HTTPEndpointConfigInitParameters `json:"httpEndpointConfig,omitempty" tf:"http_endpoint_config,omitempty"`

	// Source type for the relational database. Valid values: RDS_HTTP_ENDPOINT.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type RelationalDatabaseConfigObservation struct {

	// Amazon RDS HTTP endpoint configuration. See HTTP Endpoint Config.
	HTTPEndpointConfig []HTTPEndpointConfigObservation `json:"httpEndpointConfig,omitempty" tf:"http_endpoint_config,omitempty"`

	// Source type for the relational database. Valid values: RDS_HTTP_ENDPOINT.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

type RelationalDatabaseConfigParameters struct {

	// Amazon RDS HTTP endpoint configuration. See HTTP Endpoint Config.
	// +kubebuilder:validation:Optional
	HTTPEndpointConfig []HTTPEndpointConfigParameters `json:"httpEndpointConfig,omitempty" tf:"http_endpoint_config,omitempty"`

	// Source type for the relational database. Valid values: RDS_HTTP_ENDPOINT.
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

// DatasourceSpec defines the desired state of Datasource
type DatasourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasourceParameters `json:"forProvider"`
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
	InitProvider DatasourceInitParameters `json:"initProvider,omitempty"`
}

// DatasourceStatus defines the observed state of Datasource.
type DatasourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Datasource is the Schema for the Datasources API. Provides an AppSync Data Source.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Datasource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   DatasourceSpec   `json:"spec"`
	Status DatasourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasourceList contains a list of Datasources
type DatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datasource `json:"items"`
}

// Repository type metadata.
var (
	Datasource_Kind             = "Datasource"
	Datasource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Datasource_Kind}.String()
	Datasource_KindAPIVersion   = Datasource_Kind + "." + CRDGroupVersion.String()
	Datasource_GroupVersionKind = CRDGroupVersion.WithKind(Datasource_Kind)
)

func init() {
	SchemeBuilder.Register(&Datasource{}, &DatasourceList{})
}
