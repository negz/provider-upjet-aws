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

type CapacityUnitsInitParameters struct {

	// The amount of extra query capacity for an index and GetQuerySuggestions capacity. For more information, refer to QueryCapacityUnits.
	QueryCapacityUnits *float64 `json:"queryCapacityUnits,omitempty" tf:"query_capacity_units,omitempty"`

	// The amount of extra storage capacity for an index. A single capacity unit provides 30 GB of storage space or 100,000 documents, whichever is reached first. Minimum value of 0.
	StorageCapacityUnits *float64 `json:"storageCapacityUnits,omitempty" tf:"storage_capacity_units,omitempty"`
}

type CapacityUnitsObservation struct {

	// The amount of extra query capacity for an index and GetQuerySuggestions capacity. For more information, refer to QueryCapacityUnits.
	QueryCapacityUnits *float64 `json:"queryCapacityUnits,omitempty" tf:"query_capacity_units,omitempty"`

	// The amount of extra storage capacity for an index. A single capacity unit provides 30 GB of storage space or 100,000 documents, whichever is reached first. Minimum value of 0.
	StorageCapacityUnits *float64 `json:"storageCapacityUnits,omitempty" tf:"storage_capacity_units,omitempty"`
}

type CapacityUnitsParameters struct {

	// The amount of extra query capacity for an index and GetQuerySuggestions capacity. For more information, refer to QueryCapacityUnits.
	// +kubebuilder:validation:Optional
	QueryCapacityUnits *float64 `json:"queryCapacityUnits,omitempty" tf:"query_capacity_units,omitempty"`

	// The amount of extra storage capacity for an index. A single capacity unit provides 30 GB of storage space or 100,000 documents, whichever is reached first. Minimum value of 0.
	// +kubebuilder:validation:Optional
	StorageCapacityUnits *float64 `json:"storageCapacityUnits,omitempty" tf:"storage_capacity_units,omitempty"`
}

type DocumentMetadataConfigurationUpdatesInitParameters struct {

	// The name of the index field. Minimum length of 1. Maximum length of 30.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that provides manual tuning parameters to determine how the field affects the search results. Detailed below
	Relevance *RelevanceInitParameters `json:"relevance,omitempty" tf:"relevance,omitempty"`

	// A block that provides information about how the field is used during a search. Documented below. Detailed below
	Search *SearchInitParameters `json:"search,omitempty" tf:"search,omitempty"`

	// The data type of the index field. Valid values are STRING_VALUE, STRING_LIST_VALUE, LONG_VALUE, DATE_VALUE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DocumentMetadataConfigurationUpdatesObservation struct {

	// The name of the index field. Minimum length of 1. Maximum length of 30.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that provides manual tuning parameters to determine how the field affects the search results. Detailed below
	Relevance *RelevanceObservation `json:"relevance,omitempty" tf:"relevance,omitempty"`

	// A block that provides information about how the field is used during a search. Documented below. Detailed below
	Search *SearchObservation `json:"search,omitempty" tf:"search,omitempty"`

	// The data type of the index field. Valid values are STRING_VALUE, STRING_LIST_VALUE, LONG_VALUE, DATE_VALUE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DocumentMetadataConfigurationUpdatesParameters struct {

	// The name of the index field. Minimum length of 1. Maximum length of 30.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A block that provides manual tuning parameters to determine how the field affects the search results. Detailed below
	// +kubebuilder:validation:Optional
	Relevance *RelevanceParameters `json:"relevance,omitempty" tf:"relevance,omitempty"`

	// A block that provides information about how the field is used during a search. Documented below. Detailed below
	// +kubebuilder:validation:Optional
	Search *SearchParameters `json:"search,omitempty" tf:"search,omitempty"`

	// The data type of the index field. Valid values are STRING_VALUE, STRING_LIST_VALUE, LONG_VALUE, DATE_VALUE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type FaqStatisticsInitParameters struct {
}

type FaqStatisticsObservation struct {

	// The total number of FAQ questions and answers contained in the index.
	IndexedQuestionAnswersCount *float64 `json:"indexedQuestionAnswersCount,omitempty" tf:"indexed_question_answers_count,omitempty"`
}

type FaqStatisticsParameters struct {
}

type IndexInitParameters struct {

	// A block that sets the number of additional document storage and query capacity units that should be used by the index. Detailed below.
	CapacityUnits *CapacityUnitsInitParameters `json:"capacityUnits,omitempty" tf:"capacity_units,omitempty"`

	// The description of the Index.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more blocks that specify the configuration settings for any metadata applied to the documents in the index. Minimum number of 0 items. Maximum number of 500 items. If specified, you must define all elements, including those that are provided by default. These index fields are documented at Amazon Kendra Index documentation. For an example resource that defines these default index fields, refer to the default example above. For an example resource that appends additional index fields, refer to the append example above. All arguments for each block must be specified. Note that blocks cannot be removed since index fields cannot be deleted. This argument is detailed below.
	DocumentMetadataConfigurationUpdates []DocumentMetadataConfigurationUpdatesInitParameters `json:"documentMetadataConfigurationUpdates,omitempty" tf:"document_metadata_configuration_updates,omitempty"`

	// The Amazon Kendra edition to use for the index. Choose DEVELOPER_EDITION for indexes intended for development, testing, or proof of concept. Use ENTERPRISE_EDITION for your production databases. Once you set the edition for an index, it can't be changed. Defaults to ENTERPRISE_EDITION
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// Specifies the name of the Index.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An AWS Identity and Access Management (IAM) role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics. This is also the role you use when you call the BatchPutDocument API to index documents from an Amazon S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// A block that specifies the identifier of the AWS KMS customer managed key (CMK) that's used to encrypt data indexed by Amazon Kendra. Amazon Kendra doesn't support asymmetric CMKs. Detailed below.
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfigurationInitParameters `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user context policy. Valid values are ATTRIBUTE_FILTER or USER_TOKEN. For more information, refer to UserContextPolicy. Defaults to ATTRIBUTE_FILTER.
	UserContextPolicy *string `json:"userContextPolicy,omitempty" tf:"user_context_policy,omitempty"`

	// A block that enables fetching access levels of groups and users from an AWS Single Sign-On identity source. To configure this, see UserGroupResolutionConfiguration. Detailed below.
	UserGroupResolutionConfiguration *UserGroupResolutionConfigurationInitParameters `json:"userGroupResolutionConfiguration,omitempty" tf:"user_group_resolution_configuration,omitempty"`

	// A block that specifies the user token configuration. Detailed below.
	UserTokenConfigurations *UserTokenConfigurationsInitParameters `json:"userTokenConfigurations,omitempty" tf:"user_token_configurations,omitempty"`
}

type IndexObservation struct {

	// The Amazon Resource Name (ARN) of the Index.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A block that sets the number of additional document storage and query capacity units that should be used by the index. Detailed below.
	CapacityUnits *CapacityUnitsObservation `json:"capacityUnits,omitempty" tf:"capacity_units,omitempty"`

	// The Unix datetime that the index was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The description of the Index.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more blocks that specify the configuration settings for any metadata applied to the documents in the index. Minimum number of 0 items. Maximum number of 500 items. If specified, you must define all elements, including those that are provided by default. These index fields are documented at Amazon Kendra Index documentation. For an example resource that defines these default index fields, refer to the default example above. For an example resource that appends additional index fields, refer to the append example above. All arguments for each block must be specified. Note that blocks cannot be removed since index fields cannot be deleted. This argument is detailed below.
	DocumentMetadataConfigurationUpdates []DocumentMetadataConfigurationUpdatesObservation `json:"documentMetadataConfigurationUpdates,omitempty" tf:"document_metadata_configuration_updates,omitempty"`

	// The Amazon Kendra edition to use for the index. Choose DEVELOPER_EDITION for indexes intended for development, testing, or proof of concept. Use ENTERPRISE_EDITION for your production databases. Once you set the edition for an index, it can't be changed. Defaults to ENTERPRISE_EDITION
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// When the Status field value is FAILED, this contains a message that explains why.
	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message,omitempty"`

	// The identifier of the Index.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A block that provides information about the number of FAQ questions and answers and the number of text documents indexed. Detailed below.
	IndexStatistics []IndexStatisticsObservation `json:"indexStatistics,omitempty" tf:"index_statistics,omitempty"`

	// Specifies the name of the Index.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An AWS Identity and Access Management (IAM) role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics. This is also the role you use when you call the BatchPutDocument API to index documents from an Amazon S3 bucket.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// A block that specifies the identifier of the AWS KMS customer managed key (CMK) that's used to encrypt data indexed by Amazon Kendra. Amazon Kendra doesn't support asymmetric CMKs. Detailed below.
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfigurationObservation `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration,omitempty"`

	// The current status of the index. When the value is ACTIVE, the index is ready for use. If the Status field value is FAILED, the error_message field contains a message that explains why.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Unix datetime that the index was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// The user context policy. Valid values are ATTRIBUTE_FILTER or USER_TOKEN. For more information, refer to UserContextPolicy. Defaults to ATTRIBUTE_FILTER.
	UserContextPolicy *string `json:"userContextPolicy,omitempty" tf:"user_context_policy,omitempty"`

	// A block that enables fetching access levels of groups and users from an AWS Single Sign-On identity source. To configure this, see UserGroupResolutionConfiguration. Detailed below.
	UserGroupResolutionConfiguration *UserGroupResolutionConfigurationObservation `json:"userGroupResolutionConfiguration,omitempty" tf:"user_group_resolution_configuration,omitempty"`

	// A block that specifies the user token configuration. Detailed below.
	UserTokenConfigurations *UserTokenConfigurationsObservation `json:"userTokenConfigurations,omitempty" tf:"user_token_configurations,omitempty"`
}

type IndexParameters struct {

	// A block that sets the number of additional document storage and query capacity units that should be used by the index. Detailed below.
	// +kubebuilder:validation:Optional
	CapacityUnits *CapacityUnitsParameters `json:"capacityUnits,omitempty" tf:"capacity_units,omitempty"`

	// The description of the Index.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more blocks that specify the configuration settings for any metadata applied to the documents in the index. Minimum number of 0 items. Maximum number of 500 items. If specified, you must define all elements, including those that are provided by default. These index fields are documented at Amazon Kendra Index documentation. For an example resource that defines these default index fields, refer to the default example above. For an example resource that appends additional index fields, refer to the append example above. All arguments for each block must be specified. Note that blocks cannot be removed since index fields cannot be deleted. This argument is detailed below.
	// +kubebuilder:validation:Optional
	DocumentMetadataConfigurationUpdates []DocumentMetadataConfigurationUpdatesParameters `json:"documentMetadataConfigurationUpdates,omitempty" tf:"document_metadata_configuration_updates,omitempty"`

	// The Amazon Kendra edition to use for the index. Choose DEVELOPER_EDITION for indexes intended for development, testing, or proof of concept. Use ENTERPRISE_EDITION for your production databases. Once you set the edition for an index, it can't be changed. Defaults to ENTERPRISE_EDITION
	// +kubebuilder:validation:Optional
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// Specifies the name of the Index.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// An AWS Identity and Access Management (IAM) role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics. This is also the role you use when you call the BatchPutDocument API to index documents from an Amazon S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// A block that specifies the identifier of the AWS KMS customer managed key (CMK) that's used to encrypt data indexed by Amazon Kendra. Amazon Kendra doesn't support asymmetric CMKs. Detailed below.
	// +kubebuilder:validation:Optional
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfigurationParameters `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user context policy. Valid values are ATTRIBUTE_FILTER or USER_TOKEN. For more information, refer to UserContextPolicy. Defaults to ATTRIBUTE_FILTER.
	// +kubebuilder:validation:Optional
	UserContextPolicy *string `json:"userContextPolicy,omitempty" tf:"user_context_policy,omitempty"`

	// A block that enables fetching access levels of groups and users from an AWS Single Sign-On identity source. To configure this, see UserGroupResolutionConfiguration. Detailed below.
	// +kubebuilder:validation:Optional
	UserGroupResolutionConfiguration *UserGroupResolutionConfigurationParameters `json:"userGroupResolutionConfiguration,omitempty" tf:"user_group_resolution_configuration,omitempty"`

	// A block that specifies the user token configuration. Detailed below.
	// +kubebuilder:validation:Optional
	UserTokenConfigurations *UserTokenConfigurationsParameters `json:"userTokenConfigurations,omitempty" tf:"user_token_configurations,omitempty"`
}

type IndexStatisticsInitParameters struct {
}

type IndexStatisticsObservation struct {

	// A block that specifies the number of question and answer topics in the index. Detailed below.
	FaqStatistics []FaqStatisticsObservation `json:"faqStatistics,omitempty" tf:"faq_statistics,omitempty"`

	// A block that specifies the number of text documents indexed. Detailed below.
	TextDocumentStatistics []TextDocumentStatisticsObservation `json:"textDocumentStatistics,omitempty" tf:"text_document_statistics,omitempty"`
}

type IndexStatisticsParameters struct {
}

type JSONTokenTypeConfigurationInitParameters struct {

	// The group attribute field. Minimum length of 1. Maximum length of 2048.
	GroupAttributeField *string `json:"groupAttributeField,omitempty" tf:"group_attribute_field,omitempty"`

	// The user name attribute field. Minimum length of 1. Maximum length of 2048.
	UserNameAttributeField *string `json:"userNameAttributeField,omitempty" tf:"user_name_attribute_field,omitempty"`
}

type JSONTokenTypeConfigurationObservation struct {

	// The group attribute field. Minimum length of 1. Maximum length of 2048.
	GroupAttributeField *string `json:"groupAttributeField,omitempty" tf:"group_attribute_field,omitempty"`

	// The user name attribute field. Minimum length of 1. Maximum length of 2048.
	UserNameAttributeField *string `json:"userNameAttributeField,omitempty" tf:"user_name_attribute_field,omitempty"`
}

type JSONTokenTypeConfigurationParameters struct {

	// The group attribute field. Minimum length of 1. Maximum length of 2048.
	// +kubebuilder:validation:Optional
	GroupAttributeField *string `json:"groupAttributeField" tf:"group_attribute_field,omitempty"`

	// The user name attribute field. Minimum length of 1. Maximum length of 2048.
	// +kubebuilder:validation:Optional
	UserNameAttributeField *string `json:"userNameAttributeField" tf:"user_name_attribute_field,omitempty"`
}

type JwtTokenTypeConfigurationInitParameters struct {

	// The regular expression that identifies the claim. Minimum length of 1. Maximum length of 100.
	ClaimRegex *string `json:"claimRegex,omitempty" tf:"claim_regex,omitempty"`

	// The group attribute field. Minimum length of 1. Maximum length of 2048.
	GroupAttributeField *string `json:"groupAttributeField,omitempty" tf:"group_attribute_field,omitempty"`

	// The issuer of the token. Minimum length of 1. Maximum length of 65.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The location of the key. Valid values are URL or SECRET_MANAGER
	KeyLocation *string `json:"keyLocation,omitempty" tf:"key_location,omitempty"`

	// The Amazon Resource Name (ARN) of the secret.
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn,omitempty"`

	// The signing key URL. Valid pattern is ^(https?|ftp|file):\/\/([^\s]*)
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The user name attribute field. Minimum length of 1. Maximum length of 2048.
	UserNameAttributeField *string `json:"userNameAttributeField,omitempty" tf:"user_name_attribute_field,omitempty"`
}

type JwtTokenTypeConfigurationObservation struct {

	// The regular expression that identifies the claim. Minimum length of 1. Maximum length of 100.
	ClaimRegex *string `json:"claimRegex,omitempty" tf:"claim_regex,omitempty"`

	// The group attribute field. Minimum length of 1. Maximum length of 2048.
	GroupAttributeField *string `json:"groupAttributeField,omitempty" tf:"group_attribute_field,omitempty"`

	// The issuer of the token. Minimum length of 1. Maximum length of 65.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The location of the key. Valid values are URL or SECRET_MANAGER
	KeyLocation *string `json:"keyLocation,omitempty" tf:"key_location,omitempty"`

	// The Amazon Resource Name (ARN) of the secret.
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn,omitempty"`

	// The signing key URL. Valid pattern is ^(https?|ftp|file):\/\/([^\s]*)
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The user name attribute field. Minimum length of 1. Maximum length of 2048.
	UserNameAttributeField *string `json:"userNameAttributeField,omitempty" tf:"user_name_attribute_field,omitempty"`
}

type JwtTokenTypeConfigurationParameters struct {

	// The regular expression that identifies the claim. Minimum length of 1. Maximum length of 100.
	// +kubebuilder:validation:Optional
	ClaimRegex *string `json:"claimRegex,omitempty" tf:"claim_regex,omitempty"`

	// The group attribute field. Minimum length of 1. Maximum length of 2048.
	// +kubebuilder:validation:Optional
	GroupAttributeField *string `json:"groupAttributeField,omitempty" tf:"group_attribute_field,omitempty"`

	// The issuer of the token. Minimum length of 1. Maximum length of 65.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The location of the key. Valid values are URL or SECRET_MANAGER
	// +kubebuilder:validation:Optional
	KeyLocation *string `json:"keyLocation" tf:"key_location,omitempty"`

	// The Amazon Resource Name (ARN) of the secret.
	// +kubebuilder:validation:Optional
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn,omitempty"`

	// The signing key URL. Valid pattern is ^(https?|ftp|file):\/\/([^\s]*)
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The user name attribute field. Minimum length of 1. Maximum length of 2048.
	// +kubebuilder:validation:Optional
	UserNameAttributeField *string `json:"userNameAttributeField,omitempty" tf:"user_name_attribute_field,omitempty"`
}

type RelevanceInitParameters struct {

	// Specifies the time period that the boost applies to. For more information, refer to Duration.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Indicates that this field determines how "fresh" a document is. For more information, refer to Freshness.
	Freshness *bool `json:"freshness,omitempty" tf:"freshness,omitempty"`

	// The relative importance of the field in the search. Larger numbers provide more of a boost than smaller numbers. Minimum value of 1. Maximum value of 10.
	Importance *float64 `json:"importance,omitempty" tf:"importance,omitempty"`

	// Determines how values should be interpreted. For more information, refer to RankOrder.
	RankOrder *string `json:"rankOrder,omitempty" tf:"rank_order,omitempty"`

	// A list of values that should be given a different boost when they appear in the result list. For more information, refer to ValueImportanceMap.
	// +mapType=granular
	ValuesImportanceMap map[string]*float64 `json:"valuesImportanceMap,omitempty" tf:"values_importance_map,omitempty"`
}

type RelevanceObservation struct {

	// Specifies the time period that the boost applies to. For more information, refer to Duration.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Indicates that this field determines how "fresh" a document is. For more information, refer to Freshness.
	Freshness *bool `json:"freshness,omitempty" tf:"freshness,omitempty"`

	// The relative importance of the field in the search. Larger numbers provide more of a boost than smaller numbers. Minimum value of 1. Maximum value of 10.
	Importance *float64 `json:"importance,omitempty" tf:"importance,omitempty"`

	// Determines how values should be interpreted. For more information, refer to RankOrder.
	RankOrder *string `json:"rankOrder,omitempty" tf:"rank_order,omitempty"`

	// A list of values that should be given a different boost when they appear in the result list. For more information, refer to ValueImportanceMap.
	// +mapType=granular
	ValuesImportanceMap map[string]*float64 `json:"valuesImportanceMap,omitempty" tf:"values_importance_map,omitempty"`
}

type RelevanceParameters struct {

	// Specifies the time period that the boost applies to. For more information, refer to Duration.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Indicates that this field determines how "fresh" a document is. For more information, refer to Freshness.
	// +kubebuilder:validation:Optional
	Freshness *bool `json:"freshness,omitempty" tf:"freshness,omitempty"`

	// The relative importance of the field in the search. Larger numbers provide more of a boost than smaller numbers. Minimum value of 1. Maximum value of 10.
	// +kubebuilder:validation:Optional
	Importance *float64 `json:"importance,omitempty" tf:"importance,omitempty"`

	// Determines how values should be interpreted. For more information, refer to RankOrder.
	// +kubebuilder:validation:Optional
	RankOrder *string `json:"rankOrder,omitempty" tf:"rank_order,omitempty"`

	// A list of values that should be given a different boost when they appear in the result list. For more information, refer to ValueImportanceMap.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ValuesImportanceMap map[string]*float64 `json:"valuesImportanceMap,omitempty" tf:"values_importance_map,omitempty"`
}

type SearchInitParameters struct {

	// Determines whether the field is returned in the query response. The default is true.
	Displayable *bool `json:"displayable,omitempty" tf:"displayable,omitempty"`

	// Indicates that the field can be used to create search facets, a count of results for each value in the field. The default is false.
	Facetable *bool `json:"facetable,omitempty" tf:"facetable,omitempty"`

	// Determines whether the field is used in the search. If the Searchable field is true, you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is true for string fields and false for number and date fields.
	Searchable *bool `json:"searchable,omitempty" tf:"searchable,omitempty"`

	// Determines whether the field can be used to sort the results of a query. If you specify sorting on a field that does not have Sortable set to true, Amazon Kendra returns an exception. The default is false.
	Sortable *bool `json:"sortable,omitempty" tf:"sortable,omitempty"`
}

type SearchObservation struct {

	// Determines whether the field is returned in the query response. The default is true.
	Displayable *bool `json:"displayable,omitempty" tf:"displayable,omitempty"`

	// Indicates that the field can be used to create search facets, a count of results for each value in the field. The default is false.
	Facetable *bool `json:"facetable,omitempty" tf:"facetable,omitempty"`

	// Determines whether the field is used in the search. If the Searchable field is true, you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is true for string fields and false for number and date fields.
	Searchable *bool `json:"searchable,omitempty" tf:"searchable,omitempty"`

	// Determines whether the field can be used to sort the results of a query. If you specify sorting on a field that does not have Sortable set to true, Amazon Kendra returns an exception. The default is false.
	Sortable *bool `json:"sortable,omitempty" tf:"sortable,omitempty"`
}

type SearchParameters struct {

	// Determines whether the field is returned in the query response. The default is true.
	// +kubebuilder:validation:Optional
	Displayable *bool `json:"displayable,omitempty" tf:"displayable,omitempty"`

	// Indicates that the field can be used to create search facets, a count of results for each value in the field. The default is false.
	// +kubebuilder:validation:Optional
	Facetable *bool `json:"facetable,omitempty" tf:"facetable,omitempty"`

	// Determines whether the field is used in the search. If the Searchable field is true, you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is true for string fields and false for number and date fields.
	// +kubebuilder:validation:Optional
	Searchable *bool `json:"searchable,omitempty" tf:"searchable,omitempty"`

	// Determines whether the field can be used to sort the results of a query. If you specify sorting on a field that does not have Sortable set to true, Amazon Kendra returns an exception. The default is false.
	// +kubebuilder:validation:Optional
	Sortable *bool `json:"sortable,omitempty" tf:"sortable,omitempty"`
}

type ServerSideEncryptionConfigurationInitParameters struct {

	// The identifier of the AWS KMScustomer master key (CMK). Amazon Kendra doesn't support asymmetric CMKs.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type ServerSideEncryptionConfigurationObservation struct {

	// The identifier of the AWS KMScustomer master key (CMK). Amazon Kendra doesn't support asymmetric CMKs.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type ServerSideEncryptionConfigurationParameters struct {

	// The identifier of the AWS KMScustomer master key (CMK). Amazon Kendra doesn't support asymmetric CMKs.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type TextDocumentStatisticsInitParameters struct {
}

type TextDocumentStatisticsObservation struct {

	// The total size, in bytes, of the indexed documents.
	IndexedTextBytes *float64 `json:"indexedTextBytes,omitempty" tf:"indexed_text_bytes,omitempty"`

	// The number of text documents indexed.
	IndexedTextDocumentsCount *float64 `json:"indexedTextDocumentsCount,omitempty" tf:"indexed_text_documents_count,omitempty"`
}

type TextDocumentStatisticsParameters struct {
}

type UserGroupResolutionConfigurationInitParameters struct {

	// The identity store provider (mode) you want to use to fetch access levels of groups and users. AWS Single Sign-On is currently the only available mode. Your users and groups must exist in an AWS SSO identity source in order to use this mode. Valid Values are AWS_SSO or NONE.
	UserGroupResolutionMode *string `json:"userGroupResolutionMode,omitempty" tf:"user_group_resolution_mode,omitempty"`
}

type UserGroupResolutionConfigurationObservation struct {

	// The identity store provider (mode) you want to use to fetch access levels of groups and users. AWS Single Sign-On is currently the only available mode. Your users and groups must exist in an AWS SSO identity source in order to use this mode. Valid Values are AWS_SSO or NONE.
	UserGroupResolutionMode *string `json:"userGroupResolutionMode,omitempty" tf:"user_group_resolution_mode,omitempty"`
}

type UserGroupResolutionConfigurationParameters struct {

	// The identity store provider (mode) you want to use to fetch access levels of groups and users. AWS Single Sign-On is currently the only available mode. Your users and groups must exist in an AWS SSO identity source in order to use this mode. Valid Values are AWS_SSO or NONE.
	// +kubebuilder:validation:Optional
	UserGroupResolutionMode *string `json:"userGroupResolutionMode" tf:"user_group_resolution_mode,omitempty"`
}

type UserTokenConfigurationsInitParameters struct {

	// A block that specifies the information about the JSON token type configuration. Detailed below.
	JSONTokenTypeConfiguration *JSONTokenTypeConfigurationInitParameters `json:"jsonTokenTypeConfiguration,omitempty" tf:"json_token_type_configuration,omitempty"`

	// A block that specifies the information about the JWT token type configuration. Detailed below.
	JwtTokenTypeConfiguration *JwtTokenTypeConfigurationInitParameters `json:"jwtTokenTypeConfiguration,omitempty" tf:"jwt_token_type_configuration,omitempty"`
}

type UserTokenConfigurationsObservation struct {

	// A block that specifies the information about the JSON token type configuration. Detailed below.
	JSONTokenTypeConfiguration *JSONTokenTypeConfigurationObservation `json:"jsonTokenTypeConfiguration,omitempty" tf:"json_token_type_configuration,omitempty"`

	// A block that specifies the information about the JWT token type configuration. Detailed below.
	JwtTokenTypeConfiguration *JwtTokenTypeConfigurationObservation `json:"jwtTokenTypeConfiguration,omitempty" tf:"jwt_token_type_configuration,omitempty"`
}

type UserTokenConfigurationsParameters struct {

	// A block that specifies the information about the JSON token type configuration. Detailed below.
	// +kubebuilder:validation:Optional
	JSONTokenTypeConfiguration *JSONTokenTypeConfigurationParameters `json:"jsonTokenTypeConfiguration,omitempty" tf:"json_token_type_configuration,omitempty"`

	// A block that specifies the information about the JWT token type configuration. Detailed below.
	// +kubebuilder:validation:Optional
	JwtTokenTypeConfiguration *JwtTokenTypeConfigurationParameters `json:"jwtTokenTypeConfiguration,omitempty" tf:"jwt_token_type_configuration,omitempty"`
}

// IndexSpec defines the desired state of Index
type IndexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IndexParameters `json:"forProvider"`
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
	InitProvider IndexInitParameters `json:"initProvider,omitempty"`
}

// IndexStatus defines the observed state of Index.
type IndexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IndexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Index is the Schema for the Indexs API. Provides an Amazon Kendra Index resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Index struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IndexSpec   `json:"spec"`
	Status IndexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IndexList contains a list of Indexs
type IndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Index `json:"items"`
}

// Repository type metadata.
var (
	Index_Kind             = "Index"
	Index_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Index_Kind}.String()
	Index_KindAPIVersion   = Index_Kind + "." + CRDGroupVersion.String()
	Index_GroupVersionKind = CRDGroupVersion.WithKind(Index_Kind)
)

func init() {
	SchemeBuilder.Register(&Index{}, &IndexList{})
}
