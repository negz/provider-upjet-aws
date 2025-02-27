// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *EventDestination) ResolveReferences( // ResolveReferences of this EventDestination.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ses.aws.upbound.io", "v1beta2", "ConfigurationSet", "ConfigurationSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationSetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ConfigurationSetNameRef,
			Selector:     mg.Spec.ForProvider.ConfigurationSetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationSetName")
	}
	mg.Spec.ForProvider.ConfigurationSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationSetNameRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.KinesisDestination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KinesisDestination.RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.KinesisDestination.RoleArnRef,
				Selector:     mg.Spec.ForProvider.KinesisDestination.RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KinesisDestination.RoleArn")
		}
		mg.Spec.ForProvider.KinesisDestination.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KinesisDestination.RoleArnRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.KinesisDestination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta2", "DeliveryStream", "DeliveryStreamList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KinesisDestination.StreamArn),
				Extract:      resource.ExtractParamPath("arn", false),
				Reference:    mg.Spec.ForProvider.KinesisDestination.StreamArnRef,
				Selector:     mg.Spec.ForProvider.KinesisDestination.StreamArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KinesisDestination.StreamArn")
		}
		mg.Spec.ForProvider.KinesisDestination.StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KinesisDestination.StreamArnRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.SnsDestination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnsDestination.TopicArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.SnsDestination.TopicArnRef,
				Selector:     mg.Spec.ForProvider.SnsDestination.TopicArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SnsDestination.TopicArn")
		}
		mg.Spec.ForProvider.SnsDestination.TopicArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SnsDestination.TopicArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ses.aws.upbound.io", "v1beta2", "ConfigurationSet", "ConfigurationSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationSetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ConfigurationSetNameRef,
			Selector:     mg.Spec.InitProvider.ConfigurationSetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationSetName")
	}
	mg.Spec.InitProvider.ConfigurationSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConfigurationSetNameRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.KinesisDestination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KinesisDestination.RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.KinesisDestination.RoleArnRef,
				Selector:     mg.Spec.InitProvider.KinesisDestination.RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.KinesisDestination.RoleArn")
		}
		mg.Spec.InitProvider.KinesisDestination.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.KinesisDestination.RoleArnRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.KinesisDestination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta2", "DeliveryStream", "DeliveryStreamList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KinesisDestination.StreamArn),
				Extract:      resource.ExtractParamPath("arn", false),
				Reference:    mg.Spec.InitProvider.KinesisDestination.StreamArnRef,
				Selector:     mg.Spec.InitProvider.KinesisDestination.StreamArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.KinesisDestination.StreamArn")
		}
		mg.Spec.InitProvider.KinesisDestination.StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.KinesisDestination.StreamArnRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.SnsDestination != nil {
		{
			m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SnsDestination.TopicArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.SnsDestination.TopicArnRef,
				Selector:     mg.Spec.InitProvider.SnsDestination.TopicArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SnsDestination.TopicArn")
		}
		mg.Spec.InitProvider.SnsDestination.TopicArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SnsDestination.TopicArnRef = rsp.ResolvedReference

	}

	return nil
}
