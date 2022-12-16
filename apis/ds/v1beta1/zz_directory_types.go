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

type ConnectSettingsObservation struct {
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The IP addresses of the AD Connector servers.
	ConnectIps []*string `json:"connectIps,omitempty" tf:"connect_ips,omitempty"`
}

type ConnectSettingsParameters struct {

	// The DNS IP addresses of the domain to connect to.
	// +kubebuilder:validation:Required
	CustomerDNSIps []*string `json:"customerDnsIps" tf:"customer_dns_ips,omitempty"`

	// The username corresponding to the password provided.
	// +kubebuilder:validation:Required
	CustomerUsername *string `json:"customerUsername" tf:"customer_username,omitempty"`

	// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`

	// The identifier of the VPC that the directory is in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type DirectoryObservation struct {

	// The access URL for the directory, such as http://alias.awsapps.com.
	AccessURL *string `json:"accessUrl,omitempty" tf:"access_url,omitempty"`

	// Connector related information about the directory. Fields documented below.
	// +kubebuilder:validation:Optional
	ConnectSettings []ConnectSettingsObservation `json:"connectSettings,omitempty" tf:"connect_settings,omitempty"`

	// A list of IP addresses of the DNS servers for the directory or connector.
	DNSIPAddresses []*string `json:"dnsIpAddresses,omitempty" tf:"dns_ip_addresses,omitempty"`

	// The directory identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the security group created by the directory.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// VPC related information about the directory. Fields documented below.
	// +kubebuilder:validation:Optional
	VPCSettings []VPCSettingsObservation `json:"vpcSettings,omitempty" tf:"vpc_settings,omitempty"`
}

type DirectoryParameters struct {

	// The alias for the directory (must be unique amongst all aliases in AWS). Required for enable_sso.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Connector related information about the directory. Fields documented below.
	// +kubebuilder:validation:Optional
	ConnectSettings []ConnectSettingsParameters `json:"connectSettings,omitempty" tf:"connect_settings,omitempty"`

	// A textual description for the directory.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The MicrosoftAD edition (Standard or Enterprise). Defaults to Enterprise (applies to MicrosoftAD type only).
	// +kubebuilder:validation:Optional
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// Whether to enable single-sign on for the directory. Requires alias. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableSso *bool `json:"enableSso,omitempty" tf:"enable_sso,omitempty"`

	// The fully qualified name for the directory, such as corp.example.com
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The password for the directory administrator or connector user.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The short name of the directory, such as CORP.
	// +kubebuilder:validation:Optional
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`

	// The size of the directory (Small or Large are accepted values).
	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The directory type (SimpleAD, ADConnector or MicrosoftAD are accepted values). Defaults to SimpleAD.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// VPC related information about the directory. Fields documented below.
	// +kubebuilder:validation:Optional
	VPCSettings []VPCSettingsParameters `json:"vpcSettings,omitempty" tf:"vpc_settings,omitempty"`
}

type VPCSettingsObservation struct {
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
}

type VPCSettingsParameters struct {

	// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`

	// The identifier of the VPC that the directory is in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// DirectorySpec defines the desired state of Directory
type DirectorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DirectoryParameters `json:"forProvider"`
}

// DirectoryStatus defines the observed state of Directory.
type DirectoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DirectoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Directory is the Schema for the Directorys API. Provides a directory in AWS Directory Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Directory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectorySpec   `json:"spec"`
	Status            DirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryList contains a list of Directorys
type DirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Directory `json:"items"`
}

// Repository type metadata.
var (
	Directory_Kind             = "Directory"
	Directory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Directory_Kind}.String()
	Directory_KindAPIVersion   = Directory_Kind + "." + CRDGroupVersion.String()
	Directory_GroupVersionKind = CRDGroupVersion.WithKind(Directory_Kind)
)

func init() {
	SchemeBuilder.Register(&Directory{}, &DirectoryList{})
}
