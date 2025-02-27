//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Association) DeepCopyInto(out *Association) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Association.
func (in *Association) DeepCopy() *Association {
	if in == nil {
		return nil
	}
	out := new(Association)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Association) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationInitParameters) DeepCopyInto(out *AssociationInitParameters) {
	*out = *in
	if in.LicenseConfigurationArn != nil {
		in, out := &in.LicenseConfigurationArn, &out.LicenseConfigurationArn
		*out = new(string)
		**out = **in
	}
	if in.LicenseConfigurationArnRef != nil {
		in, out := &in.LicenseConfigurationArnRef, &out.LicenseConfigurationArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LicenseConfigurationArnSelector != nil {
		in, out := &in.LicenseConfigurationArnSelector, &out.LicenseConfigurationArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceArn != nil {
		in, out := &in.ResourceArn, &out.ResourceArn
		*out = new(string)
		**out = **in
	}
	if in.ResourceArnRef != nil {
		in, out := &in.ResourceArnRef, &out.ResourceArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceArnSelector != nil {
		in, out := &in.ResourceArnSelector, &out.ResourceArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationInitParameters.
func (in *AssociationInitParameters) DeepCopy() *AssociationInitParameters {
	if in == nil {
		return nil
	}
	out := new(AssociationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationList) DeepCopyInto(out *AssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Association, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationList.
func (in *AssociationList) DeepCopy() *AssociationList {
	if in == nil {
		return nil
	}
	out := new(AssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationObservation) DeepCopyInto(out *AssociationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LicenseConfigurationArn != nil {
		in, out := &in.LicenseConfigurationArn, &out.LicenseConfigurationArn
		*out = new(string)
		**out = **in
	}
	if in.ResourceArn != nil {
		in, out := &in.ResourceArn, &out.ResourceArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationObservation.
func (in *AssociationObservation) DeepCopy() *AssociationObservation {
	if in == nil {
		return nil
	}
	out := new(AssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationParameters) DeepCopyInto(out *AssociationParameters) {
	*out = *in
	if in.LicenseConfigurationArn != nil {
		in, out := &in.LicenseConfigurationArn, &out.LicenseConfigurationArn
		*out = new(string)
		**out = **in
	}
	if in.LicenseConfigurationArnRef != nil {
		in, out := &in.LicenseConfigurationArnRef, &out.LicenseConfigurationArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LicenseConfigurationArnSelector != nil {
		in, out := &in.LicenseConfigurationArnSelector, &out.LicenseConfigurationArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResourceArn != nil {
		in, out := &in.ResourceArn, &out.ResourceArn
		*out = new(string)
		**out = **in
	}
	if in.ResourceArnRef != nil {
		in, out := &in.ResourceArnRef, &out.ResourceArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceArnSelector != nil {
		in, out := &in.ResourceArnSelector, &out.ResourceArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationParameters.
func (in *AssociationParameters) DeepCopy() *AssociationParameters {
	if in == nil {
		return nil
	}
	out := new(AssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationSpec) DeepCopyInto(out *AssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationSpec.
func (in *AssociationSpec) DeepCopy() *AssociationSpec {
	if in == nil {
		return nil
	}
	out := new(AssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociationStatus) DeepCopyInto(out *AssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociationStatus.
func (in *AssociationStatus) DeepCopy() *AssociationStatus {
	if in == nil {
		return nil
	}
	out := new(AssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseConfiguration) DeepCopyInto(out *LicenseConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseConfiguration.
func (in *LicenseConfiguration) DeepCopy() *LicenseConfiguration {
	if in == nil {
		return nil
	}
	out := new(LicenseConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseConfigurationInitParameters) DeepCopyInto(out *LicenseConfigurationInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LicenseCount != nil {
		in, out := &in.LicenseCount, &out.LicenseCount
		*out = new(float64)
		**out = **in
	}
	if in.LicenseCountHardLimit != nil {
		in, out := &in.LicenseCountHardLimit, &out.LicenseCountHardLimit
		*out = new(bool)
		**out = **in
	}
	if in.LicenseCountingType != nil {
		in, out := &in.LicenseCountingType, &out.LicenseCountingType
		*out = new(string)
		**out = **in
	}
	if in.LicenseRules != nil {
		in, out := &in.LicenseRules, &out.LicenseRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseConfigurationInitParameters.
func (in *LicenseConfigurationInitParameters) DeepCopy() *LicenseConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(LicenseConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseConfigurationList) DeepCopyInto(out *LicenseConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseConfigurationList.
func (in *LicenseConfigurationList) DeepCopy() *LicenseConfigurationList {
	if in == nil {
		return nil
	}
	out := new(LicenseConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseConfigurationObservation) DeepCopyInto(out *LicenseConfigurationObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LicenseCount != nil {
		in, out := &in.LicenseCount, &out.LicenseCount
		*out = new(float64)
		**out = **in
	}
	if in.LicenseCountHardLimit != nil {
		in, out := &in.LicenseCountHardLimit, &out.LicenseCountHardLimit
		*out = new(bool)
		**out = **in
	}
	if in.LicenseCountingType != nil {
		in, out := &in.LicenseCountingType, &out.LicenseCountingType
		*out = new(string)
		**out = **in
	}
	if in.LicenseRules != nil {
		in, out := &in.LicenseRules, &out.LicenseRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerAccountID != nil {
		in, out := &in.OwnerAccountID, &out.OwnerAccountID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseConfigurationObservation.
func (in *LicenseConfigurationObservation) DeepCopy() *LicenseConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(LicenseConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseConfigurationParameters) DeepCopyInto(out *LicenseConfigurationParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LicenseCount != nil {
		in, out := &in.LicenseCount, &out.LicenseCount
		*out = new(float64)
		**out = **in
	}
	if in.LicenseCountHardLimit != nil {
		in, out := &in.LicenseCountHardLimit, &out.LicenseCountHardLimit
		*out = new(bool)
		**out = **in
	}
	if in.LicenseCountingType != nil {
		in, out := &in.LicenseCountingType, &out.LicenseCountingType
		*out = new(string)
		**out = **in
	}
	if in.LicenseRules != nil {
		in, out := &in.LicenseRules, &out.LicenseRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseConfigurationParameters.
func (in *LicenseConfigurationParameters) DeepCopy() *LicenseConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(LicenseConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseConfigurationSpec) DeepCopyInto(out *LicenseConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseConfigurationSpec.
func (in *LicenseConfigurationSpec) DeepCopy() *LicenseConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseConfigurationStatus) DeepCopyInto(out *LicenseConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseConfigurationStatus.
func (in *LicenseConfigurationStatus) DeepCopy() *LicenseConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}
