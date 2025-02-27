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

func (mg *Application) ResolveReferences( // ResolveReferences of this Application.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.AppversionLifecycle != nil {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppversionLifecycle.ServiceRole),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.AppversionLifecycle.ServiceRoleRef,
				Selector:     mg.Spec.ForProvider.AppversionLifecycle.ServiceRoleSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AppversionLifecycle.ServiceRole")
		}
		mg.Spec.ForProvider.AppversionLifecycle.ServiceRole = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AppversionLifecycle.ServiceRoleRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.AppversionLifecycle != nil {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AppversionLifecycle.ServiceRole),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.AppversionLifecycle.ServiceRoleRef,
				Selector:     mg.Spec.InitProvider.AppversionLifecycle.ServiceRoleSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.AppversionLifecycle.ServiceRole")
		}
		mg.Spec.InitProvider.AppversionLifecycle.ServiceRole = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.AppversionLifecycle.ServiceRoleRef = rsp.ResolvedReference

	}

	return nil
}
