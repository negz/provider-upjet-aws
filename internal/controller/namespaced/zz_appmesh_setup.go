// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	gatewayroute "github.com/upbound/provider-aws/internal/controller/namespaced/appmesh/gatewayroute"
	mesh "github.com/upbound/provider-aws/internal/controller/namespaced/appmesh/mesh"
	route "github.com/upbound/provider-aws/internal/controller/namespaced/appmesh/route"
	virtualgateway "github.com/upbound/provider-aws/internal/controller/namespaced/appmesh/virtualgateway"
	virtualnode "github.com/upbound/provider-aws/internal/controller/namespaced/appmesh/virtualnode"
	virtualrouter "github.com/upbound/provider-aws/internal/controller/namespaced/appmesh/virtualrouter"
	virtualservice "github.com/upbound/provider-aws/internal/controller/namespaced/appmesh/virtualservice"
)

// Setup_appmesh creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appmesh(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		gatewayroute.Setup,
		mesh.Setup,
		route.Setup,
		virtualgateway.Setup,
		virtualnode.Setup,
		virtualrouter.Setup,
		virtualservice.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
