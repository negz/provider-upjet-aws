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

type CognitoMemberDefinitionInitParameters struct {

	// An identifier for an application client. You must create the app client ID using Amazon Cognito.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolClient
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// An identifier for a user group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	UserGroup *string `json:"userGroup,omitempty" tf:"user_group,omitempty"`

	// Reference to a UserGroup in cognitoidp to populate userGroup.
	// +kubebuilder:validation:Optional
	UserGroupRef *v1.Reference `json:"userGroupRef,omitempty" tf:"-"`

	// Selector for a UserGroup in cognitoidp to populate userGroup.
	// +kubebuilder:validation:Optional
	UserGroupSelector *v1.Selector `json:"userGroupSelector,omitempty" tf:"-"`

	// An identifier for a user pool. The user pool must be in the same region as the service that you are calling.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("user_pool_id",false)
	UserPool *string `json:"userPool,omitempty" tf:"user_pool,omitempty"`

	// Reference to a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolRef *v1.Reference `json:"userPoolRef,omitempty" tf:"-"`

	// Selector for a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolSelector *v1.Selector `json:"userPoolSelector,omitempty" tf:"-"`
}

type CognitoMemberDefinitionObservation struct {

	// An identifier for an application client. You must create the app client ID using Amazon Cognito.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// An identifier for a user group.
	UserGroup *string `json:"userGroup,omitempty" tf:"user_group,omitempty"`

	// An identifier for a user pool. The user pool must be in the same region as the service that you are calling.
	UserPool *string `json:"userPool,omitempty" tf:"user_pool,omitempty"`
}

type CognitoMemberDefinitionParameters struct {

	// An identifier for an application client. You must create the app client ID using Amazon Cognito.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolClient
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a UserPoolClient in cognitoidp to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// An identifier for a user group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	UserGroup *string `json:"userGroup,omitempty" tf:"user_group,omitempty"`

	// Reference to a UserGroup in cognitoidp to populate userGroup.
	// +kubebuilder:validation:Optional
	UserGroupRef *v1.Reference `json:"userGroupRef,omitempty" tf:"-"`

	// Selector for a UserGroup in cognitoidp to populate userGroup.
	// +kubebuilder:validation:Optional
	UserGroupSelector *v1.Selector `json:"userGroupSelector,omitempty" tf:"-"`

	// An identifier for a user pool. The user pool must be in the same region as the service that you are calling.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("user_pool_id",false)
	// +kubebuilder:validation:Optional
	UserPool *string `json:"userPool,omitempty" tf:"user_pool,omitempty"`

	// Reference to a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolRef *v1.Reference `json:"userPoolRef,omitempty" tf:"-"`

	// Selector for a UserPoolDomain in cognitoidp to populate userPool.
	// +kubebuilder:validation:Optional
	UserPoolSelector *v1.Selector `json:"userPoolSelector,omitempty" tf:"-"`
}

type IAMPolicyConstraintsInitParameters struct {

	// When SourceIp is Enabled the worker's IP address when a task is rendered in the worker portal is added to the IAM policy as a Condition used to generate the Amazon S3 presigned URL. This IP address is checked by Amazon S3 and must match in order for the Amazon S3 resource to be rendered in the worker portal. Valid values are Enabled or Disabled
	SourceIP *string `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// When VpcSourceIp is Enabled the worker's IP address when a task is rendered in private worker portal inside the VPC is added to the IAM policy as a Condition used to generate the Amazon S3 presigned URL. To render the task successfully Amazon S3 checks that the presigned URL is being accessed over an Amazon S3 VPC Endpoint, and that the worker's IP address matches the IP address in the IAM policy. To learn more about configuring private worker portal, see Use Amazon VPC mode from a private worker portal. Valid values are Enabled or Disabled
	VPCSourceIP *string `json:"vpcSourceIp,omitempty" tf:"vpc_source_ip,omitempty"`
}

type IAMPolicyConstraintsObservation struct {

	// When SourceIp is Enabled the worker's IP address when a task is rendered in the worker portal is added to the IAM policy as a Condition used to generate the Amazon S3 presigned URL. This IP address is checked by Amazon S3 and must match in order for the Amazon S3 resource to be rendered in the worker portal. Valid values are Enabled or Disabled
	SourceIP *string `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// When VpcSourceIp is Enabled the worker's IP address when a task is rendered in private worker portal inside the VPC is added to the IAM policy as a Condition used to generate the Amazon S3 presigned URL. To render the task successfully Amazon S3 checks that the presigned URL is being accessed over an Amazon S3 VPC Endpoint, and that the worker's IP address matches the IP address in the IAM policy. To learn more about configuring private worker portal, see Use Amazon VPC mode from a private worker portal. Valid values are Enabled or Disabled
	VPCSourceIP *string `json:"vpcSourceIp,omitempty" tf:"vpc_source_ip,omitempty"`
}

type IAMPolicyConstraintsParameters struct {

	// When SourceIp is Enabled the worker's IP address when a task is rendered in the worker portal is added to the IAM policy as a Condition used to generate the Amazon S3 presigned URL. This IP address is checked by Amazon S3 and must match in order for the Amazon S3 resource to be rendered in the worker portal. Valid values are Enabled or Disabled
	// +kubebuilder:validation:Optional
	SourceIP *string `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// When VpcSourceIp is Enabled the worker's IP address when a task is rendered in private worker portal inside the VPC is added to the IAM policy as a Condition used to generate the Amazon S3 presigned URL. To render the task successfully Amazon S3 checks that the presigned URL is being accessed over an Amazon S3 VPC Endpoint, and that the worker's IP address matches the IP address in the IAM policy. To learn more about configuring private worker portal, see Use Amazon VPC mode from a private worker portal. Valid values are Enabled or Disabled
	// +kubebuilder:validation:Optional
	VPCSourceIP *string `json:"vpcSourceIp,omitempty" tf:"vpc_source_ip,omitempty"`
}

type MemberDefinitionInitParameters struct {

	// The Amazon Cognito user group that is part of the work team. See Cognito Member Definition details below.
	CognitoMemberDefinition []CognitoMemberDefinitionInitParameters `json:"cognitoMemberDefinition,omitempty" tf:"cognito_member_definition,omitempty"`

	// A list user groups that exist in your OIDC Identity Provider (IdP). One to ten groups can be used to create a single private work team. See Cognito Member Definition details below.
	OidcMemberDefinition []OidcMemberDefinitionInitParameters `json:"oidcMemberDefinition,omitempty" tf:"oidc_member_definition,omitempty"`
}

type MemberDefinitionObservation struct {

	// The Amazon Cognito user group that is part of the work team. See Cognito Member Definition details below.
	CognitoMemberDefinition []CognitoMemberDefinitionObservation `json:"cognitoMemberDefinition,omitempty" tf:"cognito_member_definition,omitempty"`

	// A list user groups that exist in your OIDC Identity Provider (IdP). One to ten groups can be used to create a single private work team. See Cognito Member Definition details below.
	OidcMemberDefinition []OidcMemberDefinitionObservation `json:"oidcMemberDefinition,omitempty" tf:"oidc_member_definition,omitempty"`
}

type MemberDefinitionParameters struct {

	// The Amazon Cognito user group that is part of the work team. See Cognito Member Definition details below.
	// +kubebuilder:validation:Optional
	CognitoMemberDefinition []CognitoMemberDefinitionParameters `json:"cognitoMemberDefinition,omitempty" tf:"cognito_member_definition,omitempty"`

	// A list user groups that exist in your OIDC Identity Provider (IdP). One to ten groups can be used to create a single private work team. See Cognito Member Definition details below.
	// +kubebuilder:validation:Optional
	OidcMemberDefinition []OidcMemberDefinitionParameters `json:"oidcMemberDefinition,omitempty" tf:"oidc_member_definition,omitempty"`
}

type NotificationConfigurationInitParameters struct {

	// The ARN for the SNS topic to which notifications should be published.
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
}

type NotificationConfigurationObservation struct {

	// The ARN for the SNS topic to which notifications should be published.
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
}

type NotificationConfigurationParameters struct {

	// The ARN for the SNS topic to which notifications should be published.
	// +kubebuilder:validation:Optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
}

type OidcMemberDefinitionInitParameters struct {

	// A list of comma separated strings that identifies user groups in your OIDC IdP. Each user group is made up of a group of private workers.
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`
}

type OidcMemberDefinitionObservation struct {

	// A list of comma separated strings that identifies user groups in your OIDC IdP. Each user group is made up of a group of private workers.
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`
}

type OidcMemberDefinitionParameters struct {

	// A list of comma separated strings that identifies user groups in your OIDC IdP. Each user group is made up of a group of private workers.
	// +kubebuilder:validation:Optional
	// +listType=set
	Groups []*string `json:"groups" tf:"groups,omitempty"`
}

type S3PresignInitParameters struct {

	// Use this parameter to specify the allowed request source. Possible sources are either SourceIp or VpcSourceIp. see IAM Policy Constraints details below.
	IAMPolicyConstraints *IAMPolicyConstraintsInitParameters `json:"iamPolicyConstraints,omitempty" tf:"iam_policy_constraints,omitempty"`
}

type S3PresignObservation struct {

	// Use this parameter to specify the allowed request source. Possible sources are either SourceIp or VpcSourceIp. see IAM Policy Constraints details below.
	IAMPolicyConstraints *IAMPolicyConstraintsObservation `json:"iamPolicyConstraints,omitempty" tf:"iam_policy_constraints,omitempty"`
}

type S3PresignParameters struct {

	// Use this parameter to specify the allowed request source. Possible sources are either SourceIp or VpcSourceIp. see IAM Policy Constraints details below.
	// +kubebuilder:validation:Optional
	IAMPolicyConstraints *IAMPolicyConstraintsParameters `json:"iamPolicyConstraints,omitempty" tf:"iam_policy_constraints,omitempty"`
}

type WorkerAccessConfigurationInitParameters struct {

	// Defines any Amazon S3 resource constraints. see S3 Presign details below.
	S3Presign *S3PresignInitParameters `json:"s3Presign,omitempty" tf:"s3_presign,omitempty"`
}

type WorkerAccessConfigurationObservation struct {

	// Defines any Amazon S3 resource constraints. see S3 Presign details below.
	S3Presign *S3PresignObservation `json:"s3Presign,omitempty" tf:"s3_presign,omitempty"`
}

type WorkerAccessConfigurationParameters struct {

	// Defines any Amazon S3 resource constraints. see S3 Presign details below.
	// +kubebuilder:validation:Optional
	S3Presign *S3PresignParameters `json:"s3Presign,omitempty" tf:"s3_presign,omitempty"`
}

type WorkteamInitParameters struct {

	// A description of the work team.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use cognito_member_definition. For workforces created using your own OIDC identity provider (IdP) use oidc_member_definition. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	MemberDefinition []MemberDefinitionInitParameters `json:"memberDefinition,omitempty" tf:"member_definition,omitempty"`

	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	NotificationConfiguration []NotificationConfigurationInitParameters `json:"notificationConfiguration,omitempty" tf:"notification_configuration,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Use this optional parameter to constrain access to an Amazon S3 resource based on the IP address using supported IAM global condition keys. The Amazon S3 resource is accessed in the worker portal using a Amazon S3 presigned URL. see Worker Access Configuration details below.
	WorkerAccessConfiguration *WorkerAccessConfigurationInitParameters `json:"workerAccessConfiguration,omitempty" tf:"worker_access_configuration,omitempty"`

	// The name of the workforce.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.Workforce
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	WorkforceName *string `json:"workforceName,omitempty" tf:"workforce_name,omitempty"`

	// Reference to a Workforce in sagemaker to populate workforceName.
	// +kubebuilder:validation:Optional
	WorkforceNameRef *v1.Reference `json:"workforceNameRef,omitempty" tf:"-"`

	// Selector for a Workforce in sagemaker to populate workforceName.
	// +kubebuilder:validation:Optional
	WorkforceNameSelector *v1.Selector `json:"workforceNameSelector,omitempty" tf:"-"`
}

type WorkteamObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A description of the work team.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the Workteam.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use cognito_member_definition. For workforces created using your own OIDC identity provider (IdP) use oidc_member_definition. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	MemberDefinition []MemberDefinitionObservation `json:"memberDefinition,omitempty" tf:"member_definition,omitempty"`

	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	NotificationConfiguration []NotificationConfigurationObservation `json:"notificationConfiguration,omitempty" tf:"notification_configuration,omitempty"`

	// The subdomain for your OIDC Identity Provider.
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Use this optional parameter to constrain access to an Amazon S3 resource based on the IP address using supported IAM global condition keys. The Amazon S3 resource is accessed in the worker portal using a Amazon S3 presigned URL. see Worker Access Configuration details below.
	WorkerAccessConfiguration *WorkerAccessConfigurationObservation `json:"workerAccessConfiguration,omitempty" tf:"worker_access_configuration,omitempty"`

	// The name of the workforce.
	WorkforceName *string `json:"workforceName,omitempty" tf:"workforce_name,omitempty"`
}

type WorkteamParameters struct {

	// A description of the work team.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use cognito_member_definition. For workforces created using your own OIDC identity provider (IdP) use oidc_member_definition. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	// +kubebuilder:validation:Optional
	MemberDefinition []MemberDefinitionParameters `json:"memberDefinition,omitempty" tf:"member_definition,omitempty"`

	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	// +kubebuilder:validation:Optional
	NotificationConfiguration []NotificationConfigurationParameters `json:"notificationConfiguration,omitempty" tf:"notification_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Use this optional parameter to constrain access to an Amazon S3 resource based on the IP address using supported IAM global condition keys. The Amazon S3 resource is accessed in the worker portal using a Amazon S3 presigned URL. see Worker Access Configuration details below.
	// +kubebuilder:validation:Optional
	WorkerAccessConfiguration *WorkerAccessConfigurationParameters `json:"workerAccessConfiguration,omitempty" tf:"worker_access_configuration,omitempty"`

	// The name of the workforce.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.Workforce
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkforceName *string `json:"workforceName,omitempty" tf:"workforce_name,omitempty"`

	// Reference to a Workforce in sagemaker to populate workforceName.
	// +kubebuilder:validation:Optional
	WorkforceNameRef *v1.Reference `json:"workforceNameRef,omitempty" tf:"-"`

	// Selector for a Workforce in sagemaker to populate workforceName.
	// +kubebuilder:validation:Optional
	WorkforceNameSelector *v1.Selector `json:"workforceNameSelector,omitempty" tf:"-"`
}

// WorkteamSpec defines the desired state of Workteam
type WorkteamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkteamParameters `json:"forProvider"`
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
	InitProvider WorkteamInitParameters `json:"initProvider,omitempty"`
}

// WorkteamStatus defines the observed state of Workteam.
type WorkteamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkteamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Workteam is the Schema for the Workteams API. Provides a SageMaker Workteam resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workteam struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.memberDefinition) || (has(self.initProvider) && has(self.initProvider.memberDefinition))",message="spec.forProvider.memberDefinition is a required parameter"
	Spec   WorkteamSpec   `json:"spec"`
	Status WorkteamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkteamList contains a list of Workteams
type WorkteamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workteam `json:"items"`
}

// Repository type metadata.
var (
	Workteam_Kind             = "Workteam"
	Workteam_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workteam_Kind}.String()
	Workteam_KindAPIVersion   = Workteam_Kind + "." + CRDGroupVersion.String()
	Workteam_GroupVersionKind = CRDGroupVersion.WithKind(Workteam_Kind)
)

func init() {
	SchemeBuilder.Register(&Workteam{}, &WorkteamList{})
}
