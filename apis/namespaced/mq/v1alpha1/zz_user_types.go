// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserInitParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/mq/v1beta2.Broker
	BrokerID *string `json:"brokerId,omitempty" tf:"broker_id,omitempty"`

	// Reference to a Broker in mq to populate brokerId.
	// +kubebuilder:validation:Optional
	BrokerIDRef *v1.Reference `json:"brokerIdRef,omitempty" tf:"-"`

	// Selector for a Broker in mq to populate brokerId.
	// +kubebuilder:validation:Optional
	BrokerIDSelector *v1.Selector `json:"brokerIdSelector,omitempty" tf:"-"`

	// Setting consoleAccess will result in an update loop till the MQ Broker to which this user belongs is restarted.
	ConsoleAccess *bool `json:"consoleAccess,omitempty" tf:"console_access,omitempty"`

	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	ReplicationUser *bool `json:"replicationUser,omitempty" tf:"replication_user,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserObservation struct {
	BrokerID *string `json:"brokerId,omitempty" tf:"broker_id,omitempty"`

	// Setting consoleAccess will result in an update loop till the MQ Broker to which this user belongs is restarted.
	ConsoleAccess *bool `json:"consoleAccess,omitempty" tf:"console_access,omitempty"`

	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ReplicationUser *bool `json:"replicationUser,omitempty" tf:"replication_user,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cluster/mq/v1beta2.Broker
	// +kubebuilder:validation:Optional
	BrokerID *string `json:"brokerId,omitempty" tf:"broker_id,omitempty"`

	// Reference to a Broker in mq to populate brokerId.
	// +kubebuilder:validation:Optional
	BrokerIDRef *v1.Reference `json:"brokerIdRef,omitempty" tf:"-"`

	// Selector for a Broker in mq to populate brokerId.
	// +kubebuilder:validation:Optional
	BrokerIDSelector *v1.Selector `json:"brokerIdSelector,omitempty" tf:"-"`

	// Setting consoleAccess will result in an update loop till the MQ Broker to which this user belongs is restarted.
	// +kubebuilder:validation:Optional
	ConsoleAccess *bool `json:"consoleAccess,omitempty" tf:"console_access,omitempty"`

	// +kubebuilder:validation:Optional
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplicationUser *bool `json:"replicationUser,omitempty" tf:"replication_user,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// User is the Schema for the Users API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
