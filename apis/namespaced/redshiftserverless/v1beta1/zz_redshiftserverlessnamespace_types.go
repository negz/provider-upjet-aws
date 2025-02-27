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

type RedshiftServerlessNamespaceInitParameters struct {

	// ID of the KMS key used to encrypt the namespace's admin credentials secret.
	AdminPasswordSecretKMSKeyID *string `json:"adminPasswordSecretKmsKeyId,omitempty" tf:"admin_password_secret_kms_key_id,omitempty"`

	// The password of the administrator for the first database created in the namespace.
	// Conflicts with manage_admin_password.
	AdminUserPasswordSecretRef *v1.SecretKeySelector `json:"adminUserPasswordSecretRef,omitempty" tf:"-"`

	// The username of the administrator for the first database created in the namespace.
	AdminUsernameSecretRef *v1.SecretKeySelector `json:"adminUsernameSecretRef,omitempty" tf:"-"`

	// The name of the first database created in the namespace.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace. When specifying default_iam_role_arn, it also must be part of iam_roles.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	DefaultIAMRoleArn *string `json:"defaultIamRoleArn,omitempty" tf:"default_iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate defaultIamRoleArn.
	// +kubebuilder:validation:Optional
	DefaultIAMRoleArnRef *v1.Reference `json:"defaultIamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate defaultIamRoleArn.
	// +kubebuilder:validation:Optional
	DefaultIAMRoleArnSelector *v1.Selector `json:"defaultIamRoleArnSelector,omitempty" tf:"-"`

	// References to Role in iam to populate iamRoles.
	// +kubebuilder:validation:Optional
	IAMRoleRefs []v1.Reference `json:"iamRoleRefs,omitempty" tf:"-"`

	// Selector for a list of Role in iam to populate iamRoles.
	// +kubebuilder:validation:Optional
	IAMRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// A list of IAM roles to associate with the namespace.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:refFieldName=IAMRoleRefs
	// +crossplane:generate:reference:selectorFieldName=IAMRoleSelector
	// +listType=set
	IAMRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// The ARN of the Amazon Web Services Key Management Service key used to encrypt your data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The types of logs the namespace can export. Available export types are userlog, connectionlog, and useractivitylog.
	// +listType=set
	LogExports []*string `json:"logExports,omitempty" tf:"log_exports,omitempty"`

	// Whether to use AWS SecretManager to manage namespace's admin credentials.
	// Conflicts with admin_user_password.
	ManageAdminPassword *bool `json:"manageAdminPassword,omitempty" tf:"manage_admin_password,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RedshiftServerlessNamespaceObservation struct {

	// Amazon Resource Name (ARN) of the Redshift Serverless Namespace.
	AdminPasswordSecretArn *string `json:"adminPasswordSecretArn,omitempty" tf:"admin_password_secret_arn,omitempty"`

	// ID of the KMS key used to encrypt the namespace's admin credentials secret.
	AdminPasswordSecretKMSKeyID *string `json:"adminPasswordSecretKmsKeyId,omitempty" tf:"admin_password_secret_kms_key_id,omitempty"`

	// Amazon Resource Name (ARN) of the Redshift Serverless Namespace.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the first database created in the namespace.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace. When specifying default_iam_role_arn, it also must be part of iam_roles.
	DefaultIAMRoleArn *string `json:"defaultIamRoleArn,omitempty" tf:"default_iam_role_arn,omitempty"`

	// A list of IAM roles to associate with the namespace.
	// +listType=set
	IAMRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// The Redshift Namespace Name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the Amazon Web Services Key Management Service key used to encrypt your data.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The types of logs the namespace can export. Available export types are userlog, connectionlog, and useractivitylog.
	// +listType=set
	LogExports []*string `json:"logExports,omitempty" tf:"log_exports,omitempty"`

	// Whether to use AWS SecretManager to manage namespace's admin credentials.
	// Conflicts with admin_user_password.
	ManageAdminPassword *bool `json:"manageAdminPassword,omitempty" tf:"manage_admin_password,omitempty"`

	// The Redshift Namespace ID.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RedshiftServerlessNamespaceParameters struct {

	// ID of the KMS key used to encrypt the namespace's admin credentials secret.
	// +kubebuilder:validation:Optional
	AdminPasswordSecretKMSKeyID *string `json:"adminPasswordSecretKmsKeyId,omitempty" tf:"admin_password_secret_kms_key_id,omitempty"`

	// The password of the administrator for the first database created in the namespace.
	// Conflicts with manage_admin_password.
	// +kubebuilder:validation:Optional
	AdminUserPasswordSecretRef *v1.SecretKeySelector `json:"adminUserPasswordSecretRef,omitempty" tf:"-"`

	// The username of the administrator for the first database created in the namespace.
	// +kubebuilder:validation:Optional
	AdminUsernameSecretRef *v1.SecretKeySelector `json:"adminUsernameSecretRef,omitempty" tf:"-"`

	// The name of the first database created in the namespace.
	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace. When specifying default_iam_role_arn, it also must be part of iam_roles.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	DefaultIAMRoleArn *string `json:"defaultIamRoleArn,omitempty" tf:"default_iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate defaultIamRoleArn.
	// +kubebuilder:validation:Optional
	DefaultIAMRoleArnRef *v1.Reference `json:"defaultIamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate defaultIamRoleArn.
	// +kubebuilder:validation:Optional
	DefaultIAMRoleArnSelector *v1.Selector `json:"defaultIamRoleArnSelector,omitempty" tf:"-"`

	// References to Role in iam to populate iamRoles.
	// +kubebuilder:validation:Optional
	IAMRoleRefs []v1.Reference `json:"iamRoleRefs,omitempty" tf:"-"`

	// Selector for a list of Role in iam to populate iamRoles.
	// +kubebuilder:validation:Optional
	IAMRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// A list of IAM roles to associate with the namespace.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:refFieldName=IAMRoleRefs
	// +crossplane:generate:reference:selectorFieldName=IAMRoleSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	IAMRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// The ARN of the Amazon Web Services Key Management Service key used to encrypt your data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The types of logs the namespace can export. Available export types are userlog, connectionlog, and useractivitylog.
	// +kubebuilder:validation:Optional
	// +listType=set
	LogExports []*string `json:"logExports,omitempty" tf:"log_exports,omitempty"`

	// Whether to use AWS SecretManager to manage namespace's admin credentials.
	// Conflicts with admin_user_password.
	// +kubebuilder:validation:Optional
	ManageAdminPassword *bool `json:"manageAdminPassword,omitempty" tf:"manage_admin_password,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RedshiftServerlessNamespaceSpec defines the desired state of RedshiftServerlessNamespace
type RedshiftServerlessNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedshiftServerlessNamespaceParameters `json:"forProvider"`
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
	InitProvider RedshiftServerlessNamespaceInitParameters `json:"initProvider,omitempty"`
}

// RedshiftServerlessNamespaceStatus defines the observed state of RedshiftServerlessNamespace.
type RedshiftServerlessNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedshiftServerlessNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RedshiftServerlessNamespace is the Schema for the RedshiftServerlessNamespaces API. Provides a Redshift Serverless Namespace resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type RedshiftServerlessNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftServerlessNamespaceSpec   `json:"spec"`
	Status            RedshiftServerlessNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftServerlessNamespaceList contains a list of RedshiftServerlessNamespaces
type RedshiftServerlessNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftServerlessNamespace `json:"items"`
}

// Repository type metadata.
var (
	RedshiftServerlessNamespace_Kind             = "RedshiftServerlessNamespace"
	RedshiftServerlessNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedshiftServerlessNamespace_Kind}.String()
	RedshiftServerlessNamespace_KindAPIVersion   = RedshiftServerlessNamespace_Kind + "." + CRDGroupVersion.String()
	RedshiftServerlessNamespace_GroupVersionKind = CRDGroupVersion.WithKind(RedshiftServerlessNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&RedshiftServerlessNamespace{}, &RedshiftServerlessNamespaceList{})
}
