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

type AlarmsInitParameters struct {

	// The name of a CloudWatch alarm in your account.
	AlarmName *string `json:"alarmName,omitempty" tf:"alarm_name,omitempty"`
}

type AlarmsObservation struct {

	// The name of a CloudWatch alarm in your account.
	AlarmName *string `json:"alarmName,omitempty" tf:"alarm_name,omitempty"`
}

type AlarmsParameters struct {

	// The name of a CloudWatch alarm in your account.
	// +kubebuilder:validation:Optional
	AlarmName *string `json:"alarmName" tf:"alarm_name,omitempty"`
}

type AutoRollbackConfigurationInitParameters struct {

	// List of CloudWatch alarms in your account that are configured to monitor metrics on an endpoint. If any alarms are tripped during a deployment, SageMaker rolls back the deployment. See Alarms.
	Alarms []AlarmsInitParameters `json:"alarms,omitempty" tf:"alarms,omitempty"`
}

type AutoRollbackConfigurationObservation struct {

	// List of CloudWatch alarms in your account that are configured to monitor metrics on an endpoint. If any alarms are tripped during a deployment, SageMaker rolls back the deployment. See Alarms.
	Alarms []AlarmsObservation `json:"alarms,omitempty" tf:"alarms,omitempty"`
}

type AutoRollbackConfigurationParameters struct {

	// List of CloudWatch alarms in your account that are configured to monitor metrics on an endpoint. If any alarms are tripped during a deployment, SageMaker rolls back the deployment. See Alarms.
	// +kubebuilder:validation:Optional
	Alarms []AlarmsParameters `json:"alarms,omitempty" tf:"alarms,omitempty"`
}

type BlueGreenUpdatePolicyInitParameters struct {

	// Maximum execution timeout for the deployment. Note that the timeout value should be larger than the total waiting time specified in termination_wait_in_seconds and wait_interval_in_seconds. Valid values are between 600 and 14400.
	MaximumExecutionTimeoutInSeconds *float64 `json:"maximumExecutionTimeoutInSeconds,omitempty" tf:"maximum_execution_timeout_in_seconds,omitempty"`

	// Additional waiting time in seconds after the completion of an endpoint deployment before terminating the old endpoint fleet. Default is 0. Valid values are between 0 and 3600.
	TerminationWaitInSeconds *float64 `json:"terminationWaitInSeconds,omitempty" tf:"termination_wait_in_seconds,omitempty"`

	// Defines the traffic routing strategy to shift traffic from the old fleet to the new fleet during an endpoint deployment. See Traffic Routing Configuration.
	TrafficRoutingConfiguration *TrafficRoutingConfigurationInitParameters `json:"trafficRoutingConfiguration,omitempty" tf:"traffic_routing_configuration,omitempty"`
}

type BlueGreenUpdatePolicyObservation struct {

	// Maximum execution timeout for the deployment. Note that the timeout value should be larger than the total waiting time specified in termination_wait_in_seconds and wait_interval_in_seconds. Valid values are between 600 and 14400.
	MaximumExecutionTimeoutInSeconds *float64 `json:"maximumExecutionTimeoutInSeconds,omitempty" tf:"maximum_execution_timeout_in_seconds,omitempty"`

	// Additional waiting time in seconds after the completion of an endpoint deployment before terminating the old endpoint fleet. Default is 0. Valid values are between 0 and 3600.
	TerminationWaitInSeconds *float64 `json:"terminationWaitInSeconds,omitempty" tf:"termination_wait_in_seconds,omitempty"`

	// Defines the traffic routing strategy to shift traffic from the old fleet to the new fleet during an endpoint deployment. See Traffic Routing Configuration.
	TrafficRoutingConfiguration *TrafficRoutingConfigurationObservation `json:"trafficRoutingConfiguration,omitempty" tf:"traffic_routing_configuration,omitempty"`
}

type BlueGreenUpdatePolicyParameters struct {

	// Maximum execution timeout for the deployment. Note that the timeout value should be larger than the total waiting time specified in termination_wait_in_seconds and wait_interval_in_seconds. Valid values are between 600 and 14400.
	// +kubebuilder:validation:Optional
	MaximumExecutionTimeoutInSeconds *float64 `json:"maximumExecutionTimeoutInSeconds,omitempty" tf:"maximum_execution_timeout_in_seconds,omitempty"`

	// Additional waiting time in seconds after the completion of an endpoint deployment before terminating the old endpoint fleet. Default is 0. Valid values are between 0 and 3600.
	// +kubebuilder:validation:Optional
	TerminationWaitInSeconds *float64 `json:"terminationWaitInSeconds,omitempty" tf:"termination_wait_in_seconds,omitempty"`

	// Defines the traffic routing strategy to shift traffic from the old fleet to the new fleet during an endpoint deployment. See Traffic Routing Configuration.
	// +kubebuilder:validation:Optional
	TrafficRoutingConfiguration *TrafficRoutingConfigurationParameters `json:"trafficRoutingConfiguration" tf:"traffic_routing_configuration,omitempty"`
}

type CanarySizeInitParameters struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type CanarySizeObservation struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type CanarySizeParameters struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type DeploymentConfigInitParameters struct {

	// Automatic rollback configuration for handling endpoint deployment failures and recovery. See Auto Rollback Configuration.
	AutoRollbackConfiguration *AutoRollbackConfigurationInitParameters `json:"autoRollbackConfiguration,omitempty" tf:"auto_rollback_configuration,omitempty"`

	// Update policy for a blue/green deployment. If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default. See Blue Green Update Config.
	BlueGreenUpdatePolicy *BlueGreenUpdatePolicyInitParameters `json:"blueGreenUpdatePolicy,omitempty" tf:"blue_green_update_policy,omitempty"`

	// Specifies a rolling deployment strategy for updating a SageMaker endpoint. See Rolling Update Policy.
	RollingUpdatePolicy *RollingUpdatePolicyInitParameters `json:"rollingUpdatePolicy,omitempty" tf:"rolling_update_policy,omitempty"`
}

type DeploymentConfigObservation struct {

	// Automatic rollback configuration for handling endpoint deployment failures and recovery. See Auto Rollback Configuration.
	AutoRollbackConfiguration *AutoRollbackConfigurationObservation `json:"autoRollbackConfiguration,omitempty" tf:"auto_rollback_configuration,omitempty"`

	// Update policy for a blue/green deployment. If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default. See Blue Green Update Config.
	BlueGreenUpdatePolicy *BlueGreenUpdatePolicyObservation `json:"blueGreenUpdatePolicy,omitempty" tf:"blue_green_update_policy,omitempty"`

	// Specifies a rolling deployment strategy for updating a SageMaker endpoint. See Rolling Update Policy.
	RollingUpdatePolicy *RollingUpdatePolicyObservation `json:"rollingUpdatePolicy,omitempty" tf:"rolling_update_policy,omitempty"`
}

type DeploymentConfigParameters struct {

	// Automatic rollback configuration for handling endpoint deployment failures and recovery. See Auto Rollback Configuration.
	// +kubebuilder:validation:Optional
	AutoRollbackConfiguration *AutoRollbackConfigurationParameters `json:"autoRollbackConfiguration,omitempty" tf:"auto_rollback_configuration,omitempty"`

	// Update policy for a blue/green deployment. If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default. See Blue Green Update Config.
	// +kubebuilder:validation:Optional
	BlueGreenUpdatePolicy *BlueGreenUpdatePolicyParameters `json:"blueGreenUpdatePolicy,omitempty" tf:"blue_green_update_policy,omitempty"`

	// Specifies a rolling deployment strategy for updating a SageMaker endpoint. See Rolling Update Policy.
	// +kubebuilder:validation:Optional
	RollingUpdatePolicy *RollingUpdatePolicyParameters `json:"rollingUpdatePolicy,omitempty" tf:"rolling_update_policy,omitempty"`
}

type EndpointInitParameters struct {

	// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations. See Deployment Config.
	DeploymentConfig *DeploymentConfigInitParameters `json:"deploymentConfig,omitempty" tf:"deployment_config,omitempty"`

	// The name of the endpoint configuration to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/sagemaker/v1beta2.EndpointConfiguration
	EndpointConfigName *string `json:"endpointConfigName,omitempty" tf:"endpoint_config_name,omitempty"`

	// Reference to a EndpointConfiguration in sagemaker to populate endpointConfigName.
	// +kubebuilder:validation:Optional
	EndpointConfigNameRef *v1.Reference `json:"endpointConfigNameRef,omitempty" tf:"-"`

	// Selector for a EndpointConfiguration in sagemaker to populate endpointConfigName.
	// +kubebuilder:validation:Optional
	EndpointConfigNameSelector *v1.Selector `json:"endpointConfigNameSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EndpointObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this endpoint.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations. See Deployment Config.
	DeploymentConfig *DeploymentConfigObservation `json:"deploymentConfig,omitempty" tf:"deployment_config,omitempty"`

	// The name of the endpoint configuration to use.
	EndpointConfigName *string `json:"endpointConfigName,omitempty" tf:"endpoint_config_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EndpointParameters struct {

	// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations. See Deployment Config.
	// +kubebuilder:validation:Optional
	DeploymentConfig *DeploymentConfigParameters `json:"deploymentConfig,omitempty" tf:"deployment_config,omitempty"`

	// The name of the endpoint configuration to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/sagemaker/v1beta2.EndpointConfiguration
	// +kubebuilder:validation:Optional
	EndpointConfigName *string `json:"endpointConfigName,omitempty" tf:"endpoint_config_name,omitempty"`

	// Reference to a EndpointConfiguration in sagemaker to populate endpointConfigName.
	// +kubebuilder:validation:Optional
	EndpointConfigNameRef *v1.Reference `json:"endpointConfigNameRef,omitempty" tf:"-"`

	// Selector for a EndpointConfiguration in sagemaker to populate endpointConfigName.
	// +kubebuilder:validation:Optional
	EndpointConfigNameSelector *v1.Selector `json:"endpointConfigNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LinearStepSizeInitParameters struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type LinearStepSizeObservation struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type LinearStepSizeParameters struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type MaximumBatchSizeInitParameters struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type MaximumBatchSizeObservation struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type MaximumBatchSizeParameters struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type RollbackMaximumBatchSizeInitParameters struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type RollbackMaximumBatchSizeObservation struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type RollbackMaximumBatchSizeParameters struct {

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Defines the capacity size, either as a number of instances or a capacity percentage.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type RollingUpdatePolicyInitParameters struct {

	// Batch size for each rolling step to provision capacity and turn on traffic on the new endpoint fleet, and terminate capacity on the old endpoint fleet. Value must be between 5% to 50% of the variant's total instance count. See Maximum Batch Size.
	MaximumBatchSize *MaximumBatchSizeInitParameters `json:"maximumBatchSize,omitempty" tf:"maximum_batch_size,omitempty"`

	// Maximum execution timeout for the deployment. Note that the timeout value should be larger than the total waiting time specified in termination_wait_in_seconds and wait_interval_in_seconds. Valid values are between 600 and 14400.
	MaximumExecutionTimeoutInSeconds *float64 `json:"maximumExecutionTimeoutInSeconds,omitempty" tf:"maximum_execution_timeout_in_seconds,omitempty"`

	// Batch size for rollback to the old endpoint fleet. Each rolling step to provision capacity and turn on traffic on the old endpoint fleet, and terminate capacity on the new endpoint fleet. If this field is absent, the default value will be set to 100% of total capacity which means to bring up the whole capacity of the old fleet at once during rollback. See Rollback Maximum Batch Size.
	RollbackMaximumBatchSize *RollbackMaximumBatchSizeInitParameters `json:"rollbackMaximumBatchSize,omitempty" tf:"rollback_maximum_batch_size,omitempty"`

	// The length of the baking period, during which SageMaker monitors alarms for each batch on the new fleet. Valid values are between 0 and 3600.
	WaitIntervalInSeconds *float64 `json:"waitIntervalInSeconds,omitempty" tf:"wait_interval_in_seconds,omitempty"`
}

type RollingUpdatePolicyObservation struct {

	// Batch size for each rolling step to provision capacity and turn on traffic on the new endpoint fleet, and terminate capacity on the old endpoint fleet. Value must be between 5% to 50% of the variant's total instance count. See Maximum Batch Size.
	MaximumBatchSize *MaximumBatchSizeObservation `json:"maximumBatchSize,omitempty" tf:"maximum_batch_size,omitempty"`

	// Maximum execution timeout for the deployment. Note that the timeout value should be larger than the total waiting time specified in termination_wait_in_seconds and wait_interval_in_seconds. Valid values are between 600 and 14400.
	MaximumExecutionTimeoutInSeconds *float64 `json:"maximumExecutionTimeoutInSeconds,omitempty" tf:"maximum_execution_timeout_in_seconds,omitempty"`

	// Batch size for rollback to the old endpoint fleet. Each rolling step to provision capacity and turn on traffic on the old endpoint fleet, and terminate capacity on the new endpoint fleet. If this field is absent, the default value will be set to 100% of total capacity which means to bring up the whole capacity of the old fleet at once during rollback. See Rollback Maximum Batch Size.
	RollbackMaximumBatchSize *RollbackMaximumBatchSizeObservation `json:"rollbackMaximumBatchSize,omitempty" tf:"rollback_maximum_batch_size,omitempty"`

	// The length of the baking period, during which SageMaker monitors alarms for each batch on the new fleet. Valid values are between 0 and 3600.
	WaitIntervalInSeconds *float64 `json:"waitIntervalInSeconds,omitempty" tf:"wait_interval_in_seconds,omitempty"`
}

type RollingUpdatePolicyParameters struct {

	// Batch size for each rolling step to provision capacity and turn on traffic on the new endpoint fleet, and terminate capacity on the old endpoint fleet. Value must be between 5% to 50% of the variant's total instance count. See Maximum Batch Size.
	// +kubebuilder:validation:Optional
	MaximumBatchSize *MaximumBatchSizeParameters `json:"maximumBatchSize" tf:"maximum_batch_size,omitempty"`

	// Maximum execution timeout for the deployment. Note that the timeout value should be larger than the total waiting time specified in termination_wait_in_seconds and wait_interval_in_seconds. Valid values are between 600 and 14400.
	// +kubebuilder:validation:Optional
	MaximumExecutionTimeoutInSeconds *float64 `json:"maximumExecutionTimeoutInSeconds,omitempty" tf:"maximum_execution_timeout_in_seconds,omitempty"`

	// Batch size for rollback to the old endpoint fleet. Each rolling step to provision capacity and turn on traffic on the old endpoint fleet, and terminate capacity on the new endpoint fleet. If this field is absent, the default value will be set to 100% of total capacity which means to bring up the whole capacity of the old fleet at once during rollback. See Rollback Maximum Batch Size.
	// +kubebuilder:validation:Optional
	RollbackMaximumBatchSize *RollbackMaximumBatchSizeParameters `json:"rollbackMaximumBatchSize,omitempty" tf:"rollback_maximum_batch_size,omitempty"`

	// The length of the baking period, during which SageMaker monitors alarms for each batch on the new fleet. Valid values are between 0 and 3600.
	// +kubebuilder:validation:Optional
	WaitIntervalInSeconds *float64 `json:"waitIntervalInSeconds" tf:"wait_interval_in_seconds,omitempty"`
}

type TrafficRoutingConfigurationInitParameters struct {

	// Batch size for the first step to turn on traffic on the new endpoint fleet. Value must be less than or equal to 50% of the variant's total instance count. See Canary Size.
	CanarySize *CanarySizeInitParameters `json:"canarySize,omitempty" tf:"canary_size,omitempty"`

	// Batch size for each step to turn on traffic on the new endpoint fleet. Value must be 10-50% of the variant's total instance count. See Linear Step Size.
	LinearStepSize *LinearStepSizeInitParameters `json:"linearStepSize,omitempty" tf:"linear_step_size,omitempty"`

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The length of the baking period, during which SageMaker monitors alarms for each batch on the new fleet. Valid values are between 0 and 3600.
	WaitIntervalInSeconds *float64 `json:"waitIntervalInSeconds,omitempty" tf:"wait_interval_in_seconds,omitempty"`
}

type TrafficRoutingConfigurationObservation struct {

	// Batch size for the first step to turn on traffic on the new endpoint fleet. Value must be less than or equal to 50% of the variant's total instance count. See Canary Size.
	CanarySize *CanarySizeObservation `json:"canarySize,omitempty" tf:"canary_size,omitempty"`

	// Batch size for each step to turn on traffic on the new endpoint fleet. Value must be 10-50% of the variant's total instance count. See Linear Step Size.
	LinearStepSize *LinearStepSizeObservation `json:"linearStepSize,omitempty" tf:"linear_step_size,omitempty"`

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The length of the baking period, during which SageMaker monitors alarms for each batch on the new fleet. Valid values are between 0 and 3600.
	WaitIntervalInSeconds *float64 `json:"waitIntervalInSeconds,omitempty" tf:"wait_interval_in_seconds,omitempty"`
}

type TrafficRoutingConfigurationParameters struct {

	// Batch size for the first step to turn on traffic on the new endpoint fleet. Value must be less than or equal to 50% of the variant's total instance count. See Canary Size.
	// +kubebuilder:validation:Optional
	CanarySize *CanarySizeParameters `json:"canarySize,omitempty" tf:"canary_size,omitempty"`

	// Batch size for each step to turn on traffic on the new endpoint fleet. Value must be 10-50% of the variant's total instance count. See Linear Step Size.
	// +kubebuilder:validation:Optional
	LinearStepSize *LinearStepSizeParameters `json:"linearStepSize,omitempty" tf:"linear_step_size,omitempty"`

	// Traffic routing strategy type. Valid values are: ALL_AT_ONCE, CANARY, and LINEAR.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// The length of the baking period, during which SageMaker monitors alarms for each batch on the new fleet. Valid values are between 0 and 3600.
	// +kubebuilder:validation:Optional
	WaitIntervalInSeconds *float64 `json:"waitIntervalInSeconds" tf:"wait_interval_in_seconds,omitempty"`
}

// EndpointSpec defines the desired state of Endpoint
type EndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointParameters `json:"forProvider"`
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
	InitProvider EndpointInitParameters `json:"initProvider,omitempty"`
}

// EndpointStatus defines the observed state of Endpoint.
type EndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Endpoint is the Schema for the Endpoints API. Provides a SageMaker Endpoint resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec"`
	Status            EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

// Repository type metadata.
var (
	Endpoint_Kind             = "Endpoint"
	Endpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Endpoint_Kind}.String()
	Endpoint_KindAPIVersion   = Endpoint_Kind + "." + CRDGroupVersion.String()
	Endpoint_GroupVersionKind = CRDGroupVersion.WithKind(Endpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
