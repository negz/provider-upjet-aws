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

type DiskIopsConfigurationInitParameters struct {

	// - The total number of SSD IOPS provisioned for the file system.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// - Specifies whether the number of IOPS for the file system is using the system. Valid values are AUTOMATIC and USER_PROVISIONED. Default value is AUTOMATIC.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type DiskIopsConfigurationObservation struct {

	// - The total number of SSD IOPS provisioned for the file system.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// - Specifies whether the number of IOPS for the file system is using the system. Valid values are AUTOMATIC and USER_PROVISIONED. Default value is AUTOMATIC.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type DiskIopsConfigurationParameters struct {

	// - The total number of SSD IOPS provisioned for the file system.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// - Specifies whether the number of IOPS for the file system is using the system. Valid values are AUTOMATIC and USER_PROVISIONED. Default value is AUTOMATIC.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type EndpointsInitParameters struct {
}

type EndpointsObservation struct {

	// An endpoint for managing your file system by setting up NetApp SnapMirror with other ONTAP systems. See Endpoint.
	Intercluster []InterclusterObservation `json:"intercluster,omitempty" tf:"intercluster,omitempty"`

	// An endpoint for managing your file system using the NetApp ONTAP CLI and NetApp ONTAP API. See Endpoint.
	Management []ManagementObservation `json:"management,omitempty" tf:"management,omitempty"`
}

type EndpointsParameters struct {
}

type InterclusterInitParameters struct {
}

type InterclusterObservation struct {

	// DNS name for the file system.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// IP addresses of the file system endpoint.
	// +listType=set
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`
}

type InterclusterParameters struct {
}

type ManagementInitParameters struct {
}

type ManagementObservation struct {

	// DNS name for the file system.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// IP addresses of the file system endpoint.
	// +listType=set
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`
}

type ManagementParameters struct {
}

type OntapFileSystemInitParameters struct {

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires automatic_backup_retention_days to be set.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// - The filesystem deployment type. Supports MULTI_AZ_1, MULTI_AZ_2, SINGLE_AZ_1, and SINGLE_AZ_2.
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration below.
	DiskIopsConfiguration *DiskIopsConfigurationInitParameters `json:"diskIopsConfiguration,omitempty" tf:"disk_iops_configuration,omitempty"`

	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	EndpointIPAddressRange *string `json:"endpointIpAddressRange,omitempty" tf:"endpoint_ip_address_range,omitempty"`

	// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
	FSXAdminPasswordSecretRef *v1.SecretKeySelector `json:"fsxAdminPasswordSecretRef,omitempty" tf:"-"`

	// - The number of ha_pairs to deploy for the file system. Valid value is 1 for SINGLE_AZ_1 or MULTI_AZ_1 and MULTI_AZ_2. Valid values are 1 through 12 for SINGLE_AZ_2.
	HaPairs *float64 `json:"haPairs,omitempty" tf:"ha_pairs,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PreferredSubnetID *string `json:"preferredSubnetId,omitempty" tf:"preferred_subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate preferredSubnetId.
	// +kubebuilder:validation:Optional
	PreferredSubnetIDRef *v1.Reference `json:"preferredSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate preferredSubnetId.
	// +kubebuilder:validation:Optional
	PreferredSubnetIDSelector *v1.Selector `json:"preferredSubnetIdSelector,omitempty" tf:"-"`

	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	// +listType=set
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The storage capacity (GiB) of the file system. Valid values between 1024 and 196608 for file systems with deployment_type SINGLE_AZ_1 and MULTI_AZ_1. Valid values are between 1024 and 524288 for MULTI_AZ_2. Valid values between 1024 (1024 per ha pair) and 1048576 for file systems with deployment_type SINGLE_AZ_2. For SINGLE_AZ_2, the 1048576 (1PB) maximum is only supported when using 2 or more ha_pairs, the maximum is 524288 (512TB) when using 1 ha_pair.
	StorageCapacity *float64 `json:"storageCapacity,omitempty" tf:"storage_capacity,omitempty"`

	// - The filesystem storage type. defaults to SSD.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of IDs for the subnets that the file system will be accessible from. Up to 2 subnets can be provided.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are 128, 256, 512, 1024, 2048, and 4096. This parameter is only supported when not using the ha_pairs parameter. Either throughput_capacity or throughput_capacity_per_ha_pair must be specified.
	ThroughputCapacity *float64 `json:"throughputCapacity,omitempty" tf:"throughput_capacity,omitempty"`

	// Sets the per-HA-pair throughput capacity (in MBps) for the file system that you're creating, as opposed to throughput_capacity which specifies the total throughput capacity for the file system. Valid value for MULTI_AZ_1 and SINGLE_AZ_1 are 128, 256, 512, 1024, 2048, and 4096. Valid values for deployment type MULTI_AZ_2 and SINGLE_AZ_2 are 384,768,1536,3072,6144 where ha_pairs is 1. Valid values for deployment type SINGLE_AZ_2 are 1536, 3072, and 6144 where ha_pairs is greater than 1. This parameter is only supported when specifying the ha_pairs parameter. Either throughput_capacity or throughput_capacity_per_ha_pair must be specified.
	ThroughputCapacityPerHaPair *float64 `json:"throughputCapacityPerHaPair,omitempty" tf:"throughput_capacity_per_ha_pair,omitempty"`

	// The preferred start time (in d:HH:MM format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

type OntapFileSystemObservation struct {

	// Amazon Resource Name of the file system.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// DNS name for the file system.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires automatic_backup_retention_days to be set.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// - The filesystem deployment type. Supports MULTI_AZ_1, MULTI_AZ_2, SINGLE_AZ_1, and SINGLE_AZ_2.
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration below.
	DiskIopsConfiguration *DiskIopsConfigurationObservation `json:"diskIopsConfiguration,omitempty" tf:"disk_iops_configuration,omitempty"`

	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	EndpointIPAddressRange *string `json:"endpointIpAddressRange,omitempty" tf:"endpoint_ip_address_range,omitempty"`

	// The endpoints that are used to access data or to manage the file system using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// - The number of ha_pairs to deploy for the file system. Valid value is 1 for SINGLE_AZ_1 or MULTI_AZ_1 and MULTI_AZ_2. Valid values are 1 through 12 for SINGLE_AZ_2.
	HaPairs *float64 `json:"haPairs,omitempty" tf:"ha_pairs,omitempty"`

	// Identifier of the file system, e.g., fs-12345678
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Set of Elastic Network Interface identifiers from which the file system is accessible The first network interface returned is the primary network interface.
	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	// AWS account identifier that created the file system.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	PreferredSubnetID *string `json:"preferredSubnetId,omitempty" tf:"preferred_subnet_id,omitempty"`

	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	// +listType=set
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The storage capacity (GiB) of the file system. Valid values between 1024 and 196608 for file systems with deployment_type SINGLE_AZ_1 and MULTI_AZ_1. Valid values are between 1024 and 524288 for MULTI_AZ_2. Valid values between 1024 (1024 per ha pair) and 1048576 for file systems with deployment_type SINGLE_AZ_2. For SINGLE_AZ_2, the 1048576 (1PB) maximum is only supported when using 2 or more ha_pairs, the maximum is 524288 (512TB) when using 1 ha_pair.
	StorageCapacity *float64 `json:"storageCapacity,omitempty" tf:"storage_capacity,omitempty"`

	// - The filesystem storage type. defaults to SSD.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A list of IDs for the subnets that the file system will be accessible from. Up to 2 subnets can be provided.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are 128, 256, 512, 1024, 2048, and 4096. This parameter is only supported when not using the ha_pairs parameter. Either throughput_capacity or throughput_capacity_per_ha_pair must be specified.
	ThroughputCapacity *float64 `json:"throughputCapacity,omitempty" tf:"throughput_capacity,omitempty"`

	// Sets the per-HA-pair throughput capacity (in MBps) for the file system that you're creating, as opposed to throughput_capacity which specifies the total throughput capacity for the file system. Valid value for MULTI_AZ_1 and SINGLE_AZ_1 are 128, 256, 512, 1024, 2048, and 4096. Valid values for deployment type MULTI_AZ_2 and SINGLE_AZ_2 are 384,768,1536,3072,6144 where ha_pairs is 1. Valid values for deployment type SINGLE_AZ_2 are 1536, 3072, and 6144 where ha_pairs is greater than 1. This parameter is only supported when specifying the ha_pairs parameter. Either throughput_capacity or throughput_capacity_per_ha_pair must be specified.
	ThroughputCapacityPerHaPair *float64 `json:"throughputCapacityPerHaPair,omitempty" tf:"throughput_capacity_per_ha_pair,omitempty"`

	// Identifier of the Virtual Private Cloud for the file system.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The preferred start time (in d:HH:MM format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

type OntapFileSystemParameters struct {

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.
	// +kubebuilder:validation:Optional
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. Requires automatic_backup_retention_days to be set.
	// +kubebuilder:validation:Optional
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// - The filesystem deployment type. Supports MULTI_AZ_1, MULTI_AZ_2, SINGLE_AZ_1, and SINGLE_AZ_2.
	// +kubebuilder:validation:Optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// The SSD IOPS configuration for the Amazon FSx for NetApp ONTAP file system. See Disk Iops Configuration below.
	// +kubebuilder:validation:Optional
	DiskIopsConfiguration *DiskIopsConfigurationParameters `json:"diskIopsConfiguration,omitempty" tf:"disk_iops_configuration,omitempty"`

	// Specifies the IP address range in which the endpoints to access your file system will be created. By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	// +kubebuilder:validation:Optional
	EndpointIPAddressRange *string `json:"endpointIpAddressRange,omitempty" tf:"endpoint_ip_address_range,omitempty"`

	// The ONTAP administrative password for the fsxadmin user that you can use to administer your file system using the ONTAP CLI and REST API.
	// +kubebuilder:validation:Optional
	FSXAdminPasswordSecretRef *v1.SecretKeySelector `json:"fsxAdminPasswordSecretRef,omitempty" tf:"-"`

	// - The number of ha_pairs to deploy for the file system. Valid value is 1 for SINGLE_AZ_1 or MULTI_AZ_1 and MULTI_AZ_2. Valid values are 1 through 12 for SINGLE_AZ_2.
	// +kubebuilder:validation:Optional
	HaPairs *float64 `json:"haPairs,omitempty" tf:"ha_pairs,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, Defaults to an AWS managed KMS Key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The ID for a subnet. A subnet is a range of IP addresses in your virtual private cloud (VPC).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PreferredSubnetID *string `json:"preferredSubnetId,omitempty" tf:"preferred_subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate preferredSubnetId.
	// +kubebuilder:validation:Optional
	PreferredSubnetIDRef *v1.Reference `json:"preferredSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate preferredSubnetId.
	// +kubebuilder:validation:Optional
	PreferredSubnetIDSelector *v1.Selector `json:"preferredSubnetIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the VPC route tables in which your file system's endpoints will be created. You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	// +kubebuilder:validation:Optional
	// +listType=set
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The storage capacity (GiB) of the file system. Valid values between 1024 and 196608 for file systems with deployment_type SINGLE_AZ_1 and MULTI_AZ_1. Valid values are between 1024 and 524288 for MULTI_AZ_2. Valid values between 1024 (1024 per ha pair) and 1048576 for file systems with deployment_type SINGLE_AZ_2. For SINGLE_AZ_2, the 1048576 (1PB) maximum is only supported when using 2 or more ha_pairs, the maximum is 524288 (512TB) when using 1 ha_pair.
	// +kubebuilder:validation:Optional
	StorageCapacity *float64 `json:"storageCapacity,omitempty" tf:"storage_capacity,omitempty"`

	// - The filesystem storage type. defaults to SSD.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of IDs for the subnets that the file system will be accessible from. Up to 2 subnets can be provided.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Sets the throughput capacity (in MBps) for the file system that you're creating. Valid values are 128, 256, 512, 1024, 2048, and 4096. This parameter is only supported when not using the ha_pairs parameter. Either throughput_capacity or throughput_capacity_per_ha_pair must be specified.
	// +kubebuilder:validation:Optional
	ThroughputCapacity *float64 `json:"throughputCapacity,omitempty" tf:"throughput_capacity,omitempty"`

	// Sets the per-HA-pair throughput capacity (in MBps) for the file system that you're creating, as opposed to throughput_capacity which specifies the total throughput capacity for the file system. Valid value for MULTI_AZ_1 and SINGLE_AZ_1 are 128, 256, 512, 1024, 2048, and 4096. Valid values for deployment type MULTI_AZ_2 and SINGLE_AZ_2 are 384,768,1536,3072,6144 where ha_pairs is 1. Valid values for deployment type SINGLE_AZ_2 are 1536, 3072, and 6144 where ha_pairs is greater than 1. This parameter is only supported when specifying the ha_pairs parameter. Either throughput_capacity or throughput_capacity_per_ha_pair must be specified.
	// +kubebuilder:validation:Optional
	ThroughputCapacityPerHaPair *float64 `json:"throughputCapacityPerHaPair,omitempty" tf:"throughput_capacity_per_ha_pair,omitempty"`

	// The preferred start time (in d:HH:MM format) to perform weekly maintenance, in the UTC time zone.
	// +kubebuilder:validation:Optional
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

// OntapFileSystemSpec defines the desired state of OntapFileSystem
type OntapFileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OntapFileSystemParameters `json:"forProvider"`
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
	InitProvider OntapFileSystemInitParameters `json:"initProvider,omitempty"`
}

// OntapFileSystemStatus defines the observed state of OntapFileSystem.
type OntapFileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OntapFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OntapFileSystem is the Schema for the OntapFileSystems API. Manages an Amazon FSx for NetApp ONTAP file system.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OntapFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.deploymentType) || (has(self.initProvider) && has(self.initProvider.deploymentType))",message="spec.forProvider.deploymentType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageCapacity) || (has(self.initProvider) && has(self.initProvider.storageCapacity))",message="spec.forProvider.storageCapacity is a required parameter"
	Spec   OntapFileSystemSpec   `json:"spec"`
	Status OntapFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OntapFileSystemList contains a list of OntapFileSystems
type OntapFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OntapFileSystem `json:"items"`
}

// Repository type metadata.
var (
	OntapFileSystem_Kind             = "OntapFileSystem"
	OntapFileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OntapFileSystem_Kind}.String()
	OntapFileSystem_KindAPIVersion   = OntapFileSystem_Kind + "." + CRDGroupVersion.String()
	OntapFileSystem_GroupVersionKind = CRDGroupVersion.WithKind(OntapFileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&OntapFileSystem{}, &OntapFileSystemList{})
}
