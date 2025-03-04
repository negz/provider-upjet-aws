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
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Service.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "VPCConnector", "VPCConnectorList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnRef,
					Selector:     mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn")
			}
			mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ObservabilityConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "ObservabilityConfiguration", "ObservabilityConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnRef,
				Selector:     mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn")
		}
		mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SourceConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "Connection", "ConnectionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnRef,
					Selector:     mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn")
			}
			mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.NetworkConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.NetworkConfiguration[i3].EgressConfiguration); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "VPCConnector", "VPCConnectorList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnRef,
					Selector:     mg.Spec.InitProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn")
			}
			mg.Spec.InitProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ObservabilityConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "ObservabilityConfiguration", "ObservabilityConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnRef,
				Selector:     mg.Spec.InitProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn")
		}
		mg.Spec.InitProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.SourceConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.SourceConfiguration[i3].AuthenticationConfiguration); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("apprunner.aws.upbound.io", "v1beta1", "Connection", "ConnectionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnRef,
					Selector:     mg.Spec.InitProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn")
			}
			mg.Spec.InitProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this VPCConnector.
func (mg *VPCConnector) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SecurityGroupRefs,
			Selector:      mg.Spec.ForProvider.SecurityGroupSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SubnetRefs,
			Selector:      mg.Spec.ForProvider.SubnetSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SecurityGroupRefs,
			Selector:      mg.Spec.InitProvider.SecurityGroupSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroups")
	}
	mg.Spec.InitProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SecurityGroupRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Subnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SubnetRefs,
			Selector:      mg.Spec.InitProvider.SubnetSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Subnets")
	}
	mg.Spec.InitProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetRefs = mrsp.ResolvedReferences

	return nil
}
