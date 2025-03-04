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

type BrokerInitParameters struct {

	// Specifies whether any broker modifications are applied immediately, or during the next maintenance window. Default is false.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Authentication strategy used to secure the broker. Valid values are simple and ldap. ldap is not supported for engine_type RabbitMQ.
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty" tf:"authentication_strategy,omitempty"`

	// Whether to automatically upgrade to new minor versions of brokers as Amazon MQ makes releases available.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// Name of the broker.
	BrokerName *string `json:"brokerName,omitempty" tf:"broker_name,omitempty"`

	// Configuration block for broker configuration. Applies to engine_type of ActiveMQ and RabbitMQ only. Detailed below.
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Deployment mode of the broker. Valid values are SINGLE_INSTANCE, ACTIVE_STANDBY_MULTI_AZ, and CLUSTER_MULTI_AZ. Default is SINGLE_INSTANCE.
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// Configuration block containing encryption options. Detailed below.
	EncryptionOptions []EncryptionOptionsInitParameters `json:"encryptionOptions,omitempty" tf:"encryption_options,omitempty"`

	// Type of broker engine. Valid values are ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty" tf:"engine_type,omitempty"`

	// Version of the broker engine. See the AmazonMQ Broker Engine docs for supported versions. For example, 5.17.6.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Broker's instance type. For example, mq.t3.micro, mq.m5.large.
	HostInstanceType *string `json:"hostInstanceType,omitempty" tf:"host_instance_type,omitempty"`

	// Configuration block for the LDAP server used to authenticate and authorize connections to the broker. Not supported for engine_type RabbitMQ. Detailed below. (Currently, AWS may not process changes to LDAP server metadata.)
	LdapServerMetadata []LdapServerMetadataInitParameters `json:"ldapServerMetadata,omitempty" tf:"ldap_server_metadata,omitempty"`

	// Configuration block for the logging configuration of the broker. Detailed below.
	Logs []LogsInitParameters `json:"logs,omitempty" tf:"logs,omitempty"`

	// Configuration block for the maintenance window start time. Detailed below.
	MaintenanceWindowStartTime []MaintenanceWindowStartTimeInitParameters `json:"maintenanceWindowStartTime,omitempty" tf:"maintenance_window_start_time,omitempty"`

	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// List of security group IDs assigned to the broker.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Storage type of the broker. For engine_type ActiveMQ, the valid values are efs and ebs, and the AWS-default is efs. For engine_type RabbitMQ, only ebs is supported. When using ebs, only the mq.m5 broker instance type family is supported.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// List of subnet IDs in which to launch the broker. A SINGLE_INSTANCE deployment requires one subnet. An ACTIVE_STANDBY_MULTI_AZ deployment requires multiple subnets.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for broker users. For engine_type of RabbitMQ, Amazon MQ does not return broker users preventing this resource from making user updates and drift detection. Detailed below.
	User []UserInitParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type BrokerObservation struct {

	// Specifies whether any broker modifications are applied immediately, or during the next maintenance window. Default is false.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// ARN of the broker.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Authentication strategy used to secure the broker. Valid values are simple and ldap. ldap is not supported for engine_type RabbitMQ.
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty" tf:"authentication_strategy,omitempty"`

	// Whether to automatically upgrade to new minor versions of brokers as Amazon MQ makes releases available.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// Name of the broker.
	BrokerName *string `json:"brokerName,omitempty" tf:"broker_name,omitempty"`

	// Configuration block for broker configuration. Applies to engine_type of ActiveMQ and RabbitMQ only. Detailed below.
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Deployment mode of the broker. Valid values are SINGLE_INSTANCE, ACTIVE_STANDBY_MULTI_AZ, and CLUSTER_MULTI_AZ. Default is SINGLE_INSTANCE.
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// Configuration block containing encryption options. Detailed below.
	EncryptionOptions []EncryptionOptionsObservation `json:"encryptionOptions,omitempty" tf:"encryption_options,omitempty"`

	// Type of broker engine. Valid values are ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty" tf:"engine_type,omitempty"`

	// Version of the broker engine. See the AmazonMQ Broker Engine docs for supported versions. For example, 5.17.6.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Broker's instance type. For example, mq.t3.micro, mq.m5.large.
	HostInstanceType *string `json:"hostInstanceType,omitempty" tf:"host_instance_type,omitempty"`

	// Unique ID that Amazon MQ generates for the broker.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of information about allocated brokers (both active & standby).
	Instances []InstancesObservation `json:"instances,omitempty" tf:"instances,omitempty"`

	// Configuration block for the LDAP server used to authenticate and authorize connections to the broker. Not supported for engine_type RabbitMQ. Detailed below. (Currently, AWS may not process changes to LDAP server metadata.)
	LdapServerMetadata []LdapServerMetadataObservation `json:"ldapServerMetadata,omitempty" tf:"ldap_server_metadata,omitempty"`

	// Configuration block for the logging configuration of the broker. Detailed below.
	Logs []LogsObservation `json:"logs,omitempty" tf:"logs,omitempty"`

	// Configuration block for the maintenance window start time. Detailed below.
	MaintenanceWindowStartTime []MaintenanceWindowStartTimeObservation `json:"maintenanceWindowStartTime,omitempty" tf:"maintenance_window_start_time,omitempty"`

	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// List of security group IDs assigned to the broker.
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Storage type of the broker. For engine_type ActiveMQ, the valid values are efs and ebs, and the AWS-default is efs. For engine_type RabbitMQ, only ebs is supported. When using ebs, only the mq.m5 broker instance type family is supported.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// List of subnet IDs in which to launch the broker. A SINGLE_INSTANCE deployment requires one subnet. An ACTIVE_STANDBY_MULTI_AZ deployment requires multiple subnets.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration block for broker users. For engine_type of RabbitMQ, Amazon MQ does not return broker users preventing this resource from making user updates and drift detection. Detailed below.
	User []UserObservation `json:"user,omitempty" tf:"user,omitempty"`
}

type BrokerParameters struct {

	// Specifies whether any broker modifications are applied immediately, or during the next maintenance window. Default is false.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Authentication strategy used to secure the broker. Valid values are simple and ldap. ldap is not supported for engine_type RabbitMQ.
	// +kubebuilder:validation:Optional
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty" tf:"authentication_strategy,omitempty"`

	// Whether to automatically upgrade to new minor versions of brokers as Amazon MQ makes releases available.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// Name of the broker.
	// +kubebuilder:validation:Optional
	BrokerName *string `json:"brokerName,omitempty" tf:"broker_name,omitempty"`

	// Configuration block for broker configuration. Applies to engine_type of ActiveMQ and RabbitMQ only. Detailed below.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Deployment mode of the broker. Valid values are SINGLE_INSTANCE, ACTIVE_STANDBY_MULTI_AZ, and CLUSTER_MULTI_AZ. Default is SINGLE_INSTANCE.
	// +kubebuilder:validation:Optional
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// Configuration block containing encryption options. Detailed below.
	// +kubebuilder:validation:Optional
	EncryptionOptions []EncryptionOptionsParameters `json:"encryptionOptions,omitempty" tf:"encryption_options,omitempty"`

	// Type of broker engine. Valid values are ActiveMQ and RabbitMQ.
	// +kubebuilder:validation:Optional
	EngineType *string `json:"engineType,omitempty" tf:"engine_type,omitempty"`

	// Version of the broker engine. See the AmazonMQ Broker Engine docs for supported versions. For example, 5.17.6.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Broker's instance type. For example, mq.t3.micro, mq.m5.large.
	// +kubebuilder:validation:Optional
	HostInstanceType *string `json:"hostInstanceType,omitempty" tf:"host_instance_type,omitempty"`

	// Configuration block for the LDAP server used to authenticate and authorize connections to the broker. Not supported for engine_type RabbitMQ. Detailed below. (Currently, AWS may not process changes to LDAP server metadata.)
	// +kubebuilder:validation:Optional
	LdapServerMetadata []LdapServerMetadataParameters `json:"ldapServerMetadata,omitempty" tf:"ldap_server_metadata,omitempty"`

	// Configuration block for the logging configuration of the broker. Detailed below.
	// +kubebuilder:validation:Optional
	Logs []LogsParameters `json:"logs,omitempty" tf:"logs,omitempty"`

	// Configuration block for the maintenance window start time. Detailed below.
	// +kubebuilder:validation:Optional
	MaintenanceWindowStartTime []MaintenanceWindowStartTimeParameters `json:"maintenanceWindowStartTime,omitempty" tf:"maintenance_window_start_time,omitempty"`

	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// List of security group IDs assigned to the broker.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Storage type of the broker. For engine_type ActiveMQ, the valid values are efs and ebs, and the AWS-default is efs. For engine_type RabbitMQ, only ebs is supported. When using ebs, only the mq.m5 broker instance type family is supported.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// List of subnet IDs in which to launch the broker. A SINGLE_INSTANCE deployment requires one subnet. An ACTIVE_STANDBY_MULTI_AZ deployment requires multiple subnets.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for broker users. For engine_type of RabbitMQ, Amazon MQ does not return broker users preventing this resource from making user updates and drift detection. Detailed below.
	// +kubebuilder:validation:Optional
	User []UserParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type ConfigurationInitParameters struct {

	// The Configuration ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/mq/v1beta1.Configuration
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Configuration in mq to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Configuration in mq to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Revision of the Configuration.
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`
}

type ConfigurationObservation struct {

	// The Configuration ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Revision of the Configuration.
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`
}

type ConfigurationParameters struct {

	// The Configuration ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/mq/v1beta1.Configuration
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Configuration in mq to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Configuration in mq to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Revision of the Configuration.
	// +kubebuilder:validation:Optional
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`
}

type EncryptionOptionsInitParameters struct {

	// Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting use_aws_owned_key to false. To perform drift detection when AWS-managed CMKs or customer-managed CMKs are in use, this value must be configured.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Whether to enable an AWS-owned KMS CMK that is not in your account. Defaults to true. Setting to false without configuring kms_key_id will create an AWS-managed CMK aliased to aws/mq in your account.
	UseAwsOwnedKey *bool `json:"useAwsOwnedKey,omitempty" tf:"use_aws_owned_key,omitempty"`
}

type EncryptionOptionsObservation struct {

	// Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting use_aws_owned_key to false. To perform drift detection when AWS-managed CMKs or customer-managed CMKs are in use, this value must be configured.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Whether to enable an AWS-owned KMS CMK that is not in your account. Defaults to true. Setting to false without configuring kms_key_id will create an AWS-managed CMK aliased to aws/mq in your account.
	UseAwsOwnedKey *bool `json:"useAwsOwnedKey,omitempty" tf:"use_aws_owned_key,omitempty"`
}

type EncryptionOptionsParameters struct {

	// Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting use_aws_owned_key to false. To perform drift detection when AWS-managed CMKs or customer-managed CMKs are in use, this value must be configured.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Whether to enable an AWS-owned KMS CMK that is not in your account. Defaults to true. Setting to false without configuring kms_key_id will create an AWS-managed CMK aliased to aws/mq in your account.
	// +kubebuilder:validation:Optional
	UseAwsOwnedKey *bool `json:"useAwsOwnedKey,omitempty" tf:"use_aws_owned_key,omitempty"`
}

type InstancesInitParameters struct {
}

type InstancesObservation struct {

	// The URL of the ActiveMQ Web Console or the RabbitMQ Management UI depending on engine_type.
	ConsoleURL *string `json:"consoleUrl,omitempty" tf:"console_url,omitempty"`

	// Broker's wire-level protocol endpoints in the following order & format referenceable e.g., as instances.0.endpoints.0 (SSL):
	Endpoints []*string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// IP Address of the broker.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

type InstancesParameters struct {
}

type LdapServerMetadataInitParameters struct {

	// List of a fully qualified domain name of the LDAP server and an optional failover server.
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// Fully qualified name of the directory to search for a user’s groups.
	RoleBase *string `json:"roleBase,omitempty" tf:"role_base,omitempty"`

	// Specifies the LDAP attribute that identifies the group name attribute in the object returned from the group membership query.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// Search criteria for groups.
	RoleSearchMatching *string `json:"roleSearchMatching,omitempty" tf:"role_search_matching,omitempty"`

	// Whether the directory search scope is the entire sub-tree.
	RoleSearchSubtree *bool `json:"roleSearchSubtree,omitempty" tf:"role_search_subtree,omitempty"`

	// Service account username.
	ServiceAccountUsername *string `json:"serviceAccountUsername,omitempty" tf:"service_account_username,omitempty"`

	// Fully qualified name of the directory where you want to search for users.
	UserBase *string `json:"userBase,omitempty" tf:"user_base,omitempty"`

	// Specifies the name of the LDAP attribute for the user group membership.
	UserRoleName *string `json:"userRoleName,omitempty" tf:"user_role_name,omitempty"`

	// Search criteria for users.
	UserSearchMatching *string `json:"userSearchMatching,omitempty" tf:"user_search_matching,omitempty"`

	// Whether the directory search scope is the entire sub-tree.
	UserSearchSubtree *bool `json:"userSearchSubtree,omitempty" tf:"user_search_subtree,omitempty"`
}

type LdapServerMetadataObservation struct {

	// List of a fully qualified domain name of the LDAP server and an optional failover server.
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// Fully qualified name of the directory to search for a user’s groups.
	RoleBase *string `json:"roleBase,omitempty" tf:"role_base,omitempty"`

	// Specifies the LDAP attribute that identifies the group name attribute in the object returned from the group membership query.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// Search criteria for groups.
	RoleSearchMatching *string `json:"roleSearchMatching,omitempty" tf:"role_search_matching,omitempty"`

	// Whether the directory search scope is the entire sub-tree.
	RoleSearchSubtree *bool `json:"roleSearchSubtree,omitempty" tf:"role_search_subtree,omitempty"`

	// Service account username.
	ServiceAccountUsername *string `json:"serviceAccountUsername,omitempty" tf:"service_account_username,omitempty"`

	// Fully qualified name of the directory where you want to search for users.
	UserBase *string `json:"userBase,omitempty" tf:"user_base,omitempty"`

	// Specifies the name of the LDAP attribute for the user group membership.
	UserRoleName *string `json:"userRoleName,omitempty" tf:"user_role_name,omitempty"`

	// Search criteria for users.
	UserSearchMatching *string `json:"userSearchMatching,omitempty" tf:"user_search_matching,omitempty"`

	// Whether the directory search scope is the entire sub-tree.
	UserSearchSubtree *bool `json:"userSearchSubtree,omitempty" tf:"user_search_subtree,omitempty"`
}

type LdapServerMetadataParameters struct {

	// List of a fully qualified domain name of the LDAP server and an optional failover server.
	// +kubebuilder:validation:Optional
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// Fully qualified name of the directory to search for a user’s groups.
	// +kubebuilder:validation:Optional
	RoleBase *string `json:"roleBase,omitempty" tf:"role_base,omitempty"`

	// Specifies the LDAP attribute that identifies the group name attribute in the object returned from the group membership query.
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// Search criteria for groups.
	// +kubebuilder:validation:Optional
	RoleSearchMatching *string `json:"roleSearchMatching,omitempty" tf:"role_search_matching,omitempty"`

	// Whether the directory search scope is the entire sub-tree.
	// +kubebuilder:validation:Optional
	RoleSearchSubtree *bool `json:"roleSearchSubtree,omitempty" tf:"role_search_subtree,omitempty"`

	// Service account password.
	// +kubebuilder:validation:Optional
	ServiceAccountPasswordSecretRef *v1.SecretKeySelector `json:"serviceAccountPasswordSecretRef,omitempty" tf:"-"`

	// Service account username.
	// +kubebuilder:validation:Optional
	ServiceAccountUsername *string `json:"serviceAccountUsername,omitempty" tf:"service_account_username,omitempty"`

	// Fully qualified name of the directory where you want to search for users.
	// +kubebuilder:validation:Optional
	UserBase *string `json:"userBase,omitempty" tf:"user_base,omitempty"`

	// Specifies the name of the LDAP attribute for the user group membership.
	// +kubebuilder:validation:Optional
	UserRoleName *string `json:"userRoleName,omitempty" tf:"user_role_name,omitempty"`

	// Search criteria for users.
	// +kubebuilder:validation:Optional
	UserSearchMatching *string `json:"userSearchMatching,omitempty" tf:"user_search_matching,omitempty"`

	// Whether the directory search scope is the entire sub-tree.
	// +kubebuilder:validation:Optional
	UserSearchSubtree *bool `json:"userSearchSubtree,omitempty" tf:"user_search_subtree,omitempty"`
}

type LogsInitParameters struct {

	// Enables audit logging. Auditing is only possible for engine_type of ActiveMQ. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to false.
	Audit *string `json:"audit,omitempty" tf:"audit,omitempty"`

	// Enables general logging via CloudWatch. Defaults to false.
	General *bool `json:"general,omitempty" tf:"general,omitempty"`
}

type LogsObservation struct {

	// Enables audit logging. Auditing is only possible for engine_type of ActiveMQ. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to false.
	Audit *string `json:"audit,omitempty" tf:"audit,omitempty"`

	// Enables general logging via CloudWatch. Defaults to false.
	General *bool `json:"general,omitempty" tf:"general,omitempty"`
}

type LogsParameters struct {

	// Enables audit logging. Auditing is only possible for engine_type of ActiveMQ. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to false.
	// +kubebuilder:validation:Optional
	Audit *string `json:"audit,omitempty" tf:"audit,omitempty"`

	// Enables general logging via CloudWatch. Defaults to false.
	// +kubebuilder:validation:Optional
	General *bool `json:"general,omitempty" tf:"general,omitempty"`
}

type MaintenanceWindowStartTimeInitParameters struct {

	// Day of the week, e.g., MONDAY, TUESDAY, or WEDNESDAY.
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// Time, in 24-hour format, e.g., 02:00.
	TimeOfDay *string `json:"timeOfDay,omitempty" tf:"time_of_day,omitempty"`

	// Time zone in either the Country/City format or the UTC offset format, e.g., CET.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type MaintenanceWindowStartTimeObservation struct {

	// Day of the week, e.g., MONDAY, TUESDAY, or WEDNESDAY.
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// Time, in 24-hour format, e.g., 02:00.
	TimeOfDay *string `json:"timeOfDay,omitempty" tf:"time_of_day,omitempty"`

	// Time zone in either the Country/City format or the UTC offset format, e.g., CET.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type MaintenanceWindowStartTimeParameters struct {

	// Day of the week, e.g., MONDAY, TUESDAY, or WEDNESDAY.
	// +kubebuilder:validation:Optional
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`

	// Time, in 24-hour format, e.g., 02:00.
	// +kubebuilder:validation:Optional
	TimeOfDay *string `json:"timeOfDay" tf:"time_of_day,omitempty"`

	// Time zone in either the Country/City format or the UTC offset format, e.g., CET.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone" tf:"time_zone,omitempty"`
}

type UserInitParameters struct {

	// Whether to enable access to the ActiveMQ Web Console for the user. Applies to engine_type of ActiveMQ only.
	ConsoleAccess *bool `json:"consoleAccess,omitempty" tf:"console_access,omitempty"`

	// List of groups (20 maximum) to which the ActiveMQ user belongs. Applies to engine_type of ActiveMQ only.
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// Whether to set set replication user. Defaults to false.
	ReplicationUser *bool `json:"replicationUser,omitempty" tf:"replication_user,omitempty"`

	// Username of the user.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserObservation struct {

	// Whether to enable access to the ActiveMQ Web Console for the user. Applies to engine_type of ActiveMQ only.
	ConsoleAccess *bool `json:"consoleAccess,omitempty" tf:"console_access,omitempty"`

	// List of groups (20 maximum) to which the ActiveMQ user belongs. Applies to engine_type of ActiveMQ only.
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// Whether to set set replication user. Defaults to false.
	ReplicationUser *bool `json:"replicationUser,omitempty" tf:"replication_user,omitempty"`

	// Username of the user.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserParameters struct {

	// Whether to enable access to the ActiveMQ Web Console for the user. Applies to engine_type of ActiveMQ only.
	// +kubebuilder:validation:Optional
	ConsoleAccess *bool `json:"consoleAccess,omitempty" tf:"console_access,omitempty"`

	// List of groups (20 maximum) to which the ActiveMQ user belongs. Applies to engine_type of ActiveMQ only.
	// +kubebuilder:validation:Optional
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// Password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Whether to set set replication user. Defaults to false.
	// +kubebuilder:validation:Optional
	ReplicationUser *bool `json:"replicationUser,omitempty" tf:"replication_user,omitempty"`

	// Username of the user.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

// BrokerSpec defines the desired state of Broker
type BrokerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BrokerParameters `json:"forProvider"`
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
	InitProvider BrokerInitParameters `json:"initProvider,omitempty"`
}

// BrokerStatus defines the observed state of Broker.
type BrokerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BrokerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Broker is the Schema for the Brokers API. Provides an MQ Broker Resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Broker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.brokerName) || (has(self.initProvider) && has(self.initProvider.brokerName))",message="spec.forProvider.brokerName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engineType) || (has(self.initProvider) && has(self.initProvider.engineType))",message="spec.forProvider.engineType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engineVersion) || (has(self.initProvider) && has(self.initProvider.engineVersion))",message="spec.forProvider.engineVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostInstanceType) || (has(self.initProvider) && has(self.initProvider.hostInstanceType))",message="spec.forProvider.hostInstanceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.user) || (has(self.initProvider) && has(self.initProvider.user))",message="spec.forProvider.user is a required parameter"
	Spec   BrokerSpec   `json:"spec"`
	Status BrokerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BrokerList contains a list of Brokers
type BrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Broker `json:"items"`
}

// Repository type metadata.
var (
	Broker_Kind             = "Broker"
	Broker_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Broker_Kind}.String()
	Broker_KindAPIVersion   = Broker_Kind + "." + CRDGroupVersion.String()
	Broker_GroupVersionKind = CRDGroupVersion.WithKind(Broker_Kind)
)

func init() {
	SchemeBuilder.Register(&Broker{}, &BrokerList{})
}
