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

type AccessEndpointsInitParameters struct {

	// Type of the interface endpoint.
	// See the AccessEndpoint AWS API documentation for valid values.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// ID of the VPC in which the interface endpoint is used.
	VpceID *string `json:"vpceId,omitempty" tf:"vpce_id,omitempty"`
}

type AccessEndpointsObservation struct {

	// Type of the interface endpoint.
	// See the AccessEndpoint AWS API documentation for valid values.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// ID of the VPC in which the interface endpoint is used.
	VpceID *string `json:"vpceId,omitempty" tf:"vpce_id,omitempty"`
}

type AccessEndpointsParameters struct {

	// Type of the interface endpoint.
	// See the AccessEndpoint AWS API documentation for valid values.
	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType" tf:"endpoint_type,omitempty"`

	// ID of the VPC in which the interface endpoint is used.
	// +kubebuilder:validation:Optional
	VpceID *string `json:"vpceId,omitempty" tf:"vpce_id,omitempty"`
}

type ApplicationSettingsInitParameters struct {

	// Whether application settings should be persisted.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Name of the settings group.
	// Required when enabled is true.
	// Can be up to 100 characters.
	SettingsGroup *string `json:"settingsGroup,omitempty" tf:"settings_group,omitempty"`
}

type ApplicationSettingsObservation struct {

	// Whether application settings should be persisted.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Name of the settings group.
	// Required when enabled is true.
	// Can be up to 100 characters.
	SettingsGroup *string `json:"settingsGroup,omitempty" tf:"settings_group,omitempty"`
}

type ApplicationSettingsParameters struct {

	// Whether application settings should be persisted.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Name of the settings group.
	// Required when enabled is true.
	// Can be up to 100 characters.
	// +kubebuilder:validation:Optional
	SettingsGroup *string `json:"settingsGroup,omitempty" tf:"settings_group,omitempty"`
}

type StackInitParameters struct {

	// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	// See access_endpoints below.
	AccessEndpoints []AccessEndpointsInitParameters `json:"accessEndpoints,omitempty" tf:"access_endpoints,omitempty"`

	// Settings for application settings persistence.
	// See application_settings below.
	ApplicationSettings []ApplicationSettingsInitParameters `json:"applicationSettings,omitempty" tf:"application_settings,omitempty"`

	// Description for the AppStream stack.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Stack name to display.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	// +listType=set
	EmbedHostDomains []*string `json:"embedHostDomains,omitempty" tf:"embed_host_domains,omitempty"`

	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	FeedbackURL *string `json:"feedbackUrl,omitempty" tf:"feedback_url,omitempty"`

	// Unique name for the AppStream stack.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// URL that users are redirected to after their streaming session ends.
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// Configuration block for the storage connectors to enable.
	// See storage_connectors below.
	StorageConnectors []StorageConnectorsInitParameters `json:"storageConnectors,omitempty" tf:"storage_connectors,omitempty"`

	// The streaming protocol you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
	// See streaming_experience_settings below.
	StreamingExperienceSettings []StreamingExperienceSettingsInitParameters `json:"streamingExperienceSettings,omitempty" tf:"streaming_experience_settings,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. If not provided, these settings are configured automatically by AWS.
	// See user_settings below.
	UserSettings []UserSettingsInitParameters `json:"userSettings,omitempty" tf:"user_settings,omitempty"`
}

type StackObservation struct {

	// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	// See access_endpoints below.
	AccessEndpoints []AccessEndpointsObservation `json:"accessEndpoints,omitempty" tf:"access_endpoints,omitempty"`

	// Settings for application settings persistence.
	// See application_settings below.
	ApplicationSettings []ApplicationSettingsObservation `json:"applicationSettings,omitempty" tf:"application_settings,omitempty"`

	// ARN of the appstream stack.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Date and time, in UTC and extended RFC 3339 format, when the stack was created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Description for the AppStream stack.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Stack name to display.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	// +listType=set
	EmbedHostDomains []*string `json:"embedHostDomains,omitempty" tf:"embed_host_domains,omitempty"`

	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	FeedbackURL *string `json:"feedbackUrl,omitempty" tf:"feedback_url,omitempty"`

	// Unique ID of the appstream stack.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique name for the AppStream stack.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// URL that users are redirected to after their streaming session ends.
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// Configuration block for the storage connectors to enable.
	// See storage_connectors below.
	StorageConnectors []StorageConnectorsObservation `json:"storageConnectors,omitempty" tf:"storage_connectors,omitempty"`

	// The streaming protocol you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
	// See streaming_experience_settings below.
	StreamingExperienceSettings []StreamingExperienceSettingsObservation `json:"streamingExperienceSettings,omitempty" tf:"streaming_experience_settings,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. If not provided, these settings are configured automatically by AWS.
	// See user_settings below.
	UserSettings []UserSettingsObservation `json:"userSettings,omitempty" tf:"user_settings,omitempty"`
}

type StackParameters struct {

	// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	// See access_endpoints below.
	// +kubebuilder:validation:Optional
	AccessEndpoints []AccessEndpointsParameters `json:"accessEndpoints,omitempty" tf:"access_endpoints,omitempty"`

	// Settings for application settings persistence.
	// See application_settings below.
	// +kubebuilder:validation:Optional
	ApplicationSettings []ApplicationSettingsParameters `json:"applicationSettings,omitempty" tf:"application_settings,omitempty"`

	// Description for the AppStream stack.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Stack name to display.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	// +kubebuilder:validation:Optional
	// +listType=set
	EmbedHostDomains []*string `json:"embedHostDomains,omitempty" tf:"embed_host_domains,omitempty"`

	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	// +kubebuilder:validation:Optional
	FeedbackURL *string `json:"feedbackUrl,omitempty" tf:"feedback_url,omitempty"`

	// Unique name for the AppStream stack.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// URL that users are redirected to after their streaming session ends.
	// +kubebuilder:validation:Optional
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for the storage connectors to enable.
	// See storage_connectors below.
	// +kubebuilder:validation:Optional
	StorageConnectors []StorageConnectorsParameters `json:"storageConnectors,omitempty" tf:"storage_connectors,omitempty"`

	// The streaming protocol you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
	// See streaming_experience_settings below.
	// +kubebuilder:validation:Optional
	StreamingExperienceSettings []StreamingExperienceSettingsParameters `json:"streamingExperienceSettings,omitempty" tf:"streaming_experience_settings,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. If not provided, these settings are configured automatically by AWS.
	// See user_settings below.
	// +kubebuilder:validation:Optional
	UserSettings []UserSettingsParameters `json:"userSettings,omitempty" tf:"user_settings,omitempty"`
}

type StorageConnectorsInitParameters struct {

	// Type of storage connector.
	// Valid values are HOMEFOLDERS, GOOGLE_DRIVE, or ONE_DRIVE.
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type,omitempty"`

	// Names of the domains for the account.
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// ARN of the storage connector.
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" tf:"resource_identifier,omitempty"`
}

type StorageConnectorsObservation struct {

	// Type of storage connector.
	// Valid values are HOMEFOLDERS, GOOGLE_DRIVE, or ONE_DRIVE.
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type,omitempty"`

	// Names of the domains for the account.
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// ARN of the storage connector.
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" tf:"resource_identifier,omitempty"`
}

type StorageConnectorsParameters struct {

	// Type of storage connector.
	// Valid values are HOMEFOLDERS, GOOGLE_DRIVE, or ONE_DRIVE.
	// +kubebuilder:validation:Optional
	ConnectorType *string `json:"connectorType" tf:"connector_type,omitempty"`

	// Names of the domains for the account.
	// +kubebuilder:validation:Optional
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// ARN of the storage connector.
	// +kubebuilder:validation:Optional
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" tf:"resource_identifier,omitempty"`
}

type StreamingExperienceSettingsInitParameters struct {

	// The preferred protocol that you want to use while streaming your application.
	// Valid values are TCP and UDP.
	PreferredProtocol *string `json:"preferredProtocol,omitempty" tf:"preferred_protocol,omitempty"`
}

type StreamingExperienceSettingsObservation struct {

	// The preferred protocol that you want to use while streaming your application.
	// Valid values are TCP and UDP.
	PreferredProtocol *string `json:"preferredProtocol,omitempty" tf:"preferred_protocol,omitempty"`
}

type StreamingExperienceSettingsParameters struct {

	// The preferred protocol that you want to use while streaming your application.
	// Valid values are TCP and UDP.
	// +kubebuilder:validation:Optional
	PreferredProtocol *string `json:"preferredProtocol,omitempty" tf:"preferred_protocol,omitempty"`
}

type UserSettingsInitParameters struct {

	// Action that is enabled or disabled.
	// Valid values are CLIPBOARD_COPY_FROM_LOCAL_DEVICE,  CLIPBOARD_COPY_TO_LOCAL_DEVICE, FILE_UPLOAD, FILE_DOWNLOAD, PRINTING_TO_LOCAL_DEVICE, DOMAIN_PASSWORD_SIGNIN, or DOMAIN_SMART_CARD_SIGNIN.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Whether the action is enabled or disabled.
	// Valid values are ENABLED or DISABLED.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type UserSettingsObservation struct {

	// Action that is enabled or disabled.
	// Valid values are CLIPBOARD_COPY_FROM_LOCAL_DEVICE,  CLIPBOARD_COPY_TO_LOCAL_DEVICE, FILE_UPLOAD, FILE_DOWNLOAD, PRINTING_TO_LOCAL_DEVICE, DOMAIN_PASSWORD_SIGNIN, or DOMAIN_SMART_CARD_SIGNIN.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Whether the action is enabled or disabled.
	// Valid values are ENABLED or DISABLED.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type UserSettingsParameters struct {

	// Action that is enabled or disabled.
	// Valid values are CLIPBOARD_COPY_FROM_LOCAL_DEVICE,  CLIPBOARD_COPY_TO_LOCAL_DEVICE, FILE_UPLOAD, FILE_DOWNLOAD, PRINTING_TO_LOCAL_DEVICE, DOMAIN_PASSWORD_SIGNIN, or DOMAIN_SMART_CARD_SIGNIN.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// Whether the action is enabled or disabled.
	// Valid values are ENABLED or DISABLED.
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission" tf:"permission,omitempty"`
}

// StackSpec defines the desired state of Stack
type StackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackParameters `json:"forProvider"`
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
	InitProvider StackInitParameters `json:"initProvider,omitempty"`
}

// StackStatus defines the observed state of Stack.
type StackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Stack is the Schema for the Stacks API. Provides an AppStream stack
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   StackSpec   `json:"spec"`
	Status StackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackList contains a list of Stacks
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stack `json:"items"`
}

// Repository type metadata.
var (
	Stack_Kind             = "Stack"
	Stack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stack_Kind}.String()
	Stack_KindAPIVersion   = Stack_Kind + "." + CRDGroupVersion.String()
	Stack_GroupVersionKind = CRDGroupVersion.WithKind(Stack_Kind)
)

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
