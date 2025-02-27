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

type AutoscalingGroupsInitParameters struct {
}

type AutoscalingGroupsObservation struct {

	// Name of the AutoScaling Group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AutoscalingGroupsParameters struct {
}

type LaunchTemplateInitParameters struct {

	// Identifier of the EC2 Launch Template. Conflicts with name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the EC2 Launch Template. Conflicts with id.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// EC2 Launch Template version number. While the API accepts values like $Default and $Latest, the API will convert the value to the associated version number (e.g., 1). Using the default_version or latest_version attribute of the aws_launch_template resource or data source is recommended for this argument.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type LaunchTemplateObservation struct {

	// Identifier of the EC2 Launch Template. Conflicts with name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the EC2 Launch Template. Conflicts with id.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// EC2 Launch Template version number. While the API accepts values like $Default and $Latest, the API will convert the value to the associated version number (e.g., 1). Using the default_version or latest_version attribute of the aws_launch_template resource or data source is recommended for this argument.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type LaunchTemplateParameters struct {

	// Identifier of the EC2 Launch Template. Conflicts with name.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the EC2 Launch Template. Conflicts with id.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// EC2 Launch Template version number. While the API accepts values like $Default and $Latest, the API will convert the value to the associated version number (e.g., 1). Using the default_version or latest_version attribute of the aws_launch_template resource or data source is recommended for this argument.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type NodeGroupInitParameters struct {

	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the AWS documentation for valid values.
	AMIType *string `json:"amiType,omitempty" tf:"ami_type,omitempty"`

	// Type of capacity associated with the EKS Node Group. Valid values: ON_DEMAND, SPOT.
	CapacityType *string `json:"capacityType,omitempty" tf:"capacity_type,omitempty"`

	// Disk size in GiB for worker nodes. Defaults to 50 for Windows, 20 all other node groups.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateVersion *bool `json:"forceUpdateVersion,omitempty" tf:"force_update_version,omitempty"`

	// List of instance types associated with the EKS Node Group. Defaults to ["t3.medium"].
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Configuration block with Launch Template settings. See launch_template below for details. Conflicts with remote_access.
	LaunchTemplate []LaunchTemplateInitParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// –  Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	NodeRoleArn *string `json:"nodeRoleArn,omitempty" tf:"node_role_arn,omitempty"`

	// Reference to a Role in iam to populate nodeRoleArn.
	// +kubebuilder:validation:Optional
	NodeRoleArnRef *v1.Reference `json:"nodeRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate nodeRoleArn.
	// +kubebuilder:validation:Optional
	NodeRoleArnSelector *v1.Selector `json:"nodeRoleArnSelector,omitempty" tf:"-"`

	// –  AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`

	// Configuration block with remote access settings. See remote_access below for details. Conflicts with launch_template.
	RemoteAccess []RemoteAccessInitParameters `json:"remoteAccess,omitempty" tf:"remote_access,omitempty"`

	// Configuration block with scaling settings. See scaling_config below for details.
	ScalingConfig []ScalingConfigInitParameters `json:"scalingConfig,omitempty" tf:"scaling_config,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Identifiers of EC2 Subnets to associate with the EKS Node Group. Amazon EKS managed node groups can be launched in both public and private subnets. If you plan to deploy load balancers to a subnet, the private subnet must have tag kubernetes.io/role/internal-elb, the public subnet must have tag kubernetes.io/role/elb.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	Taint []TaintInitParameters `json:"taint,omitempty" tf:"taint,omitempty"`

	// Configuration block with update settings. See update_config below for details.
	UpdateConfig []UpdateConfigInitParameters `json:"updateConfig,omitempty" tf:"update_config,omitempty"`

	// –  Kubernetes version. Defaults to EKS Cluster Kubernetes version.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/eks/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("version",false)
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Reference to a Cluster in eks to populate version.
	// +kubebuilder:validation:Optional
	VersionRef *v1.Reference `json:"versionRef,omitempty" tf:"-"`

	// Selector for a Cluster in eks to populate version.
	// +kubebuilder:validation:Optional
	VersionSelector *v1.Selector `json:"versionSelector,omitempty" tf:"-"`
}

type NodeGroupObservation struct {

	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the AWS documentation for valid values.
	AMIType *string `json:"amiType,omitempty" tf:"ami_type,omitempty"`

	// Amazon Resource Name (ARN) of the EKS Node Group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Type of capacity associated with the EKS Node Group. Valid values: ON_DEMAND, SPOT.
	CapacityType *string `json:"capacityType,omitempty" tf:"capacity_type,omitempty"`

	// –  Name of the EKS Cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Disk size in GiB for worker nodes. Defaults to 50 for Windows, 20 all other node groups.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateVersion *bool `json:"forceUpdateVersion,omitempty" tf:"force_update_version,omitempty"`

	// EKS Cluster name and EKS Node Group name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of instance types associated with the EKS Node Group. Defaults to ["t3.medium"].
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Configuration block with Launch Template settings. See launch_template below for details. Conflicts with remote_access.
	LaunchTemplate []LaunchTemplateObservation `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// –  Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn *string `json:"nodeRoleArn,omitempty" tf:"node_role_arn,omitempty"`

	// –  AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`

	// Configuration block with remote access settings. See remote_access below for details. Conflicts with launch_template.
	RemoteAccess []RemoteAccessObservation `json:"remoteAccess,omitempty" tf:"remote_access,omitempty"`

	// List of objects containing information about underlying resources.
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// Configuration block with scaling settings. See scaling_config below for details.
	ScalingConfig []ScalingConfigObservation `json:"scalingConfig,omitempty" tf:"scaling_config,omitempty"`

	// Status of the EKS Node Group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Identifiers of EC2 Subnets to associate with the EKS Node Group. Amazon EKS managed node groups can be launched in both public and private subnets. If you plan to deploy load balancers to a subnet, the private subnet must have tag kubernetes.io/role/internal-elb, the public subnet must have tag kubernetes.io/role/elb.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	Taint []TaintObservation `json:"taint,omitempty" tf:"taint,omitempty"`

	// Configuration block with update settings. See update_config below for details.
	UpdateConfig []UpdateConfigObservation `json:"updateConfig,omitempty" tf:"update_config,omitempty"`

	// –  Kubernetes version. Defaults to EKS Cluster Kubernetes version.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NodeGroupParameters struct {

	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the AWS documentation for valid values.
	// +kubebuilder:validation:Optional
	AMIType *string `json:"amiType,omitempty" tf:"ami_type,omitempty"`

	// Type of capacity associated with the EKS Node Group. Valid values: ON_DEMAND, SPOT.
	// +kubebuilder:validation:Optional
	CapacityType *string `json:"capacityType,omitempty" tf:"capacity_type,omitempty"`

	// –  Name of the EKS Cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/eks/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=ExternalNameIfClusterActive()
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in eks to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in eks to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// Disk size in GiB for worker nodes. Defaults to 50 for Windows, 20 all other node groups.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	// +kubebuilder:validation:Optional
	ForceUpdateVersion *bool `json:"forceUpdateVersion,omitempty" tf:"force_update_version,omitempty"`

	// List of instance types associated with the EKS Node Group. Defaults to ["t3.medium"].
	// +kubebuilder:validation:Optional
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Configuration block with Launch Template settings. See launch_template below for details. Conflicts with remote_access.
	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// –  Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	NodeRoleArn *string `json:"nodeRoleArn,omitempty" tf:"node_role_arn,omitempty"`

	// Reference to a Role in iam to populate nodeRoleArn.
	// +kubebuilder:validation:Optional
	NodeRoleArnRef *v1.Reference `json:"nodeRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate nodeRoleArn.
	// +kubebuilder:validation:Optional
	NodeRoleArnSelector *v1.Selector `json:"nodeRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// –  AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	// +kubebuilder:validation:Optional
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`

	// Configuration block with remote access settings. See remote_access below for details. Conflicts with launch_template.
	// +kubebuilder:validation:Optional
	RemoteAccess []RemoteAccessParameters `json:"remoteAccess,omitempty" tf:"remote_access,omitempty"`

	// Configuration block with scaling settings. See scaling_config below for details.
	// +kubebuilder:validation:Optional
	ScalingConfig []ScalingConfigParameters `json:"scalingConfig,omitempty" tf:"scaling_config,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Identifiers of EC2 Subnets to associate with the EKS Node Group. Amazon EKS managed node groups can be launched in both public and private subnets. If you plan to deploy load balancers to a subnet, the private subnet must have tag kubernetes.io/role/internal-elb, the public subnet must have tag kubernetes.io/role/elb.
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

	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	// +kubebuilder:validation:Optional
	Taint []TaintParameters `json:"taint,omitempty" tf:"taint,omitempty"`

	// Configuration block with update settings. See update_config below for details.
	// +kubebuilder:validation:Optional
	UpdateConfig []UpdateConfigParameters `json:"updateConfig,omitempty" tf:"update_config,omitempty"`

	// –  Kubernetes version. Defaults to EKS Cluster Kubernetes version.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/eks/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("version",false)
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Reference to a Cluster in eks to populate version.
	// +kubebuilder:validation:Optional
	VersionRef *v1.Reference `json:"versionRef,omitempty" tf:"-"`

	// Selector for a Cluster in eks to populate version.
	// +kubebuilder:validation:Optional
	VersionSelector *v1.Selector `json:"versionSelector,omitempty" tf:"-"`
}

type RemoteAccessInitParameters struct {

	// EC2 Key Pair name that provides access for remote communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify source_security_group_ids when you create an EKS Node Group, either port 3389 for Windows, or port 22 for all other operating systems is opened on the worker nodes to the Internet (0.0.0.0/0). For Windows nodes, this will allow you to use RDP, for all others this allows you to SSH into the worker nodes.
	EC2SSHKey *string `json:"ec2SshKey,omitempty" tf:"ec2_ssh_key,omitempty"`

	// References to SecurityGroup in ec2 to populate sourceSecurityGroupIds.
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDRefs []v1.Reference `json:"sourceSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate sourceSecurityGroupIds.
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDSelector *v1.Selector `json:"sourceSecurityGroupIdSelector,omitempty" tf:"-"`

	// Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify ec2_ssh_key, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SourceSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SourceSecurityGroupIDSelector
	// +listType=set
	SourceSecurityGroupIds []*string `json:"sourceSecurityGroupIds,omitempty" tf:"source_security_group_ids,omitempty"`
}

type RemoteAccessObservation struct {

	// EC2 Key Pair name that provides access for remote communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify source_security_group_ids when you create an EKS Node Group, either port 3389 for Windows, or port 22 for all other operating systems is opened on the worker nodes to the Internet (0.0.0.0/0). For Windows nodes, this will allow you to use RDP, for all others this allows you to SSH into the worker nodes.
	EC2SSHKey *string `json:"ec2SshKey,omitempty" tf:"ec2_ssh_key,omitempty"`

	// Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify ec2_ssh_key, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
	// +listType=set
	SourceSecurityGroupIds []*string `json:"sourceSecurityGroupIds,omitempty" tf:"source_security_group_ids,omitempty"`
}

type RemoteAccessParameters struct {

	// EC2 Key Pair name that provides access for remote communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify source_security_group_ids when you create an EKS Node Group, either port 3389 for Windows, or port 22 for all other operating systems is opened on the worker nodes to the Internet (0.0.0.0/0). For Windows nodes, this will allow you to use RDP, for all others this allows you to SSH into the worker nodes.
	// +kubebuilder:validation:Optional
	EC2SSHKey *string `json:"ec2SshKey,omitempty" tf:"ec2_ssh_key,omitempty"`

	// References to SecurityGroup in ec2 to populate sourceSecurityGroupIds.
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDRefs []v1.Reference `json:"sourceSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate sourceSecurityGroupIds.
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDSelector *v1.Selector `json:"sourceSecurityGroupIdSelector,omitempty" tf:"-"`

	// Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify ec2_ssh_key, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SourceSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SourceSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceSecurityGroupIds []*string `json:"sourceSecurityGroupIds,omitempty" tf:"source_security_group_ids,omitempty"`
}

type ResourcesInitParameters struct {
}

type ResourcesObservation struct {

	// List of objects containing information about AutoScaling Groups.
	AutoscalingGroups []AutoscalingGroupsObservation `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups,omitempty"`

	// Identifier of the remote access EC2 Security Group.
	RemoteAccessSecurityGroupID *string `json:"remoteAccessSecurityGroupId,omitempty" tf:"remote_access_security_group_id,omitempty"`
}

type ResourcesParameters struct {
}

type ScalingConfigInitParameters struct {

	// Desired number of worker nodes.
	DesiredSize *float64 `json:"desiredSize,omitempty" tf:"desired_size,omitempty"`

	// Maximum number of worker nodes.
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Minimum number of worker nodes.
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`
}

type ScalingConfigObservation struct {

	// Desired number of worker nodes.
	DesiredSize *float64 `json:"desiredSize,omitempty" tf:"desired_size,omitempty"`

	// Maximum number of worker nodes.
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Minimum number of worker nodes.
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`
}

type ScalingConfigParameters struct {

	// Desired number of worker nodes.
	// +kubebuilder:validation:Optional
	DesiredSize *float64 `json:"desiredSize" tf:"desired_size,omitempty"`

	// Maximum number of worker nodes.
	// +kubebuilder:validation:Optional
	MaxSize *float64 `json:"maxSize" tf:"max_size,omitempty"`

	// Minimum number of worker nodes.
	// +kubebuilder:validation:Optional
	MinSize *float64 `json:"minSize" tf:"min_size,omitempty"`
}

type TaintInitParameters struct {

	// The effect of the taint. Valid values: NO_SCHEDULE, NO_EXECUTE, PREFER_NO_SCHEDULE.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// The key of the taint. Maximum length of 63.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value of the taint. Maximum length of 63.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TaintObservation struct {

	// The effect of the taint. Valid values: NO_SCHEDULE, NO_EXECUTE, PREFER_NO_SCHEDULE.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// The key of the taint. Maximum length of 63.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value of the taint. Maximum length of 63.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TaintParameters struct {

	// The effect of the taint. Valid values: NO_SCHEDULE, NO_EXECUTE, PREFER_NO_SCHEDULE.
	// +kubebuilder:validation:Optional
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// The key of the taint. Maximum length of 63.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The value of the taint. Maximum length of 63.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UpdateConfigInitParameters struct {

	// Desired max number of unavailable worker nodes during node group update.
	MaxUnavailable *float64 `json:"maxUnavailable,omitempty" tf:"max_unavailable,omitempty"`

	// Desired max percentage of unavailable worker nodes during node group update.
	MaxUnavailablePercentage *float64 `json:"maxUnavailablePercentage,omitempty" tf:"max_unavailable_percentage,omitempty"`
}

type UpdateConfigObservation struct {

	// Desired max number of unavailable worker nodes during node group update.
	MaxUnavailable *float64 `json:"maxUnavailable,omitempty" tf:"max_unavailable,omitempty"`

	// Desired max percentage of unavailable worker nodes during node group update.
	MaxUnavailablePercentage *float64 `json:"maxUnavailablePercentage,omitempty" tf:"max_unavailable_percentage,omitempty"`
}

type UpdateConfigParameters struct {

	// Desired max number of unavailable worker nodes during node group update.
	// +kubebuilder:validation:Optional
	MaxUnavailable *float64 `json:"maxUnavailable,omitempty" tf:"max_unavailable,omitempty"`

	// Desired max percentage of unavailable worker nodes during node group update.
	// +kubebuilder:validation:Optional
	MaxUnavailablePercentage *float64 `json:"maxUnavailablePercentage,omitempty" tf:"max_unavailable_percentage,omitempty"`
}

// NodeGroupSpec defines the desired state of NodeGroup
type NodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeGroupParameters `json:"forProvider"`
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
	InitProvider NodeGroupInitParameters `json:"initProvider,omitempty"`
}

// NodeGroupStatus defines the observed state of NodeGroup.
type NodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NodeGroup is the Schema for the NodeGroups API. Manages an EKS Node Group
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scalingConfig) || (has(self.initProvider) && has(self.initProvider.scalingConfig))",message="spec.forProvider.scalingConfig is a required parameter"
	Spec   NodeGroupSpec   `json:"spec"`
	Status NodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroupList contains a list of NodeGroups
type NodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeGroup `json:"items"`
}

// Repository type metadata.
var (
	NodeGroup_Kind             = "NodeGroup"
	NodeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeGroup_Kind}.String()
	NodeGroup_KindAPIVersion   = NodeGroup_Kind + "." + CRDGroupVersion.String()
	NodeGroup_GroupVersionKind = CRDGroupVersion.WithKind(NodeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeGroup{}, &NodeGroupList{})
}
