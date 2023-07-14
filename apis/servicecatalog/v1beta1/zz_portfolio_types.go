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

type PortfolioObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Description of the portfolio
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Service Catalog Portfolio.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the portfolio.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the person or organization who owns the portfolio.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PortfolioParameters struct {

	// Description of the portfolio
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the portfolio.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the person or organization who owns the portfolio.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PortfolioSpec defines the desired state of Portfolio
type PortfolioSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortfolioParameters `json:"forProvider"`
}

// PortfolioStatus defines the observed state of Portfolio.
type PortfolioStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortfolioObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Portfolio is the Schema for the Portfolios API. Provides a resource to create a Service Catalog portfolio
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Portfolio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerName)",message="providerName is a required parameter"
	Spec   PortfolioSpec   `json:"spec"`
	Status PortfolioStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortfolioList contains a list of Portfolios
type PortfolioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Portfolio `json:"items"`
}

// Repository type metadata.
var (
	Portfolio_Kind             = "Portfolio"
	Portfolio_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Portfolio_Kind}.String()
	Portfolio_KindAPIVersion   = Portfolio_Kind + "." + CRDGroupVersion.String()
	Portfolio_GroupVersionKind = CRDGroupVersion.WithKind(Portfolio_Kind)
)

func init() {
	SchemeBuilder.Register(&Portfolio{}, &PortfolioList{})
}
