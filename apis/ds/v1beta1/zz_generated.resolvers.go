/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ConditionalForwarder) ResolveReferences( // ResolveReferences of this ConditionalForwarder.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta1", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DirectoryIDRef,
			Selector:     mg.Spec.ForProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectoryID")
	}
	mg.Spec.ForProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectoryIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Directory.
func (mg *Directory) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ConnectSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.ConnectSettings[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.ConnectSettings[i3].SubnetIdsRefs,
				Selector:      mg.Spec.ForProvider.ConnectSettings[i3].SubnetIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ConnectSettings[i3].SubnetIds")
		}
		mg.Spec.ForProvider.ConnectSettings[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.ConnectSettings[i3].SubnetIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ConnectSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectSettings[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.ConnectSettings[i3].VPCIDRef,
				Selector:     mg.Spec.ForProvider.ConnectSettings[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ConnectSettings[i3].VPCID")
		}
		mg.Spec.ForProvider.ConnectSettings[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ConnectSettings[i3].VPCIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSettings[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCSettings[i3].SubnetIdsRefs,
				Selector:      mg.Spec.ForProvider.VPCSettings[i3].SubnetIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCSettings[i3].SubnetIds")
		}
		mg.Spec.ForProvider.VPCSettings[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCSettings[i3].SubnetIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCSettings[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.VPCSettings[i3].VPCIDRef,
				Selector:     mg.Spec.ForProvider.VPCSettings[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCSettings[i3].VPCID")
		}
		mg.Spec.ForProvider.VPCSettings[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPCSettings[i3].VPCIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ConnectSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.ConnectSettings[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.ConnectSettings[i3].SubnetIdsRefs,
				Selector:      mg.Spec.InitProvider.ConnectSettings[i3].SubnetIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ConnectSettings[i3].SubnetIds")
		}
		mg.Spec.InitProvider.ConnectSettings[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.ConnectSettings[i3].SubnetIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ConnectSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConnectSettings[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.ConnectSettings[i3].VPCIDRef,
				Selector:     mg.Spec.InitProvider.ConnectSettings[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ConnectSettings[i3].VPCID")
		}
		mg.Spec.InitProvider.ConnectSettings[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ConnectSettings[i3].VPCIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCSettings[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCSettings[i3].SubnetIdsRefs,
				Selector:      mg.Spec.InitProvider.VPCSettings[i3].SubnetIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCSettings[i3].SubnetIds")
		}
		mg.Spec.InitProvider.VPCSettings[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCSettings[i3].SubnetIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCSettings[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.VPCSettings[i3].VPCIDRef,
				Selector:     mg.Spec.InitProvider.VPCSettings[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCSettings[i3].VPCID")
		}
		mg.Spec.InitProvider.VPCSettings[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.VPCSettings[i3].VPCIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this SharedDirectory.
func (mg *SharedDirectory) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta1", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DirectoryIDRef,
			Selector:     mg.Spec.ForProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectoryID")
	}
	mg.Spec.ForProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectoryIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ds.aws.upbound.io", "v1beta1", "Directory", "DirectoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DirectoryID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DirectoryIDRef,
			Selector:     mg.Spec.InitProvider.DirectoryIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DirectoryID")
	}
	mg.Spec.InitProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DirectoryIDRef = rsp.ResolvedReference

	return nil
}
