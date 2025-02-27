// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	attachmentaccepter "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/attachmentaccepter"
	connectattachment "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/connectattachment"
	connection "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/connection"
	corenetwork "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/corenetwork"
	customergatewayassociation "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/customergatewayassociation"
	device "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/device"
	globalnetwork "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/globalnetwork"
	link "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/link"
	linkassociation "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/linkassociation"
	site "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/site"
	transitgatewayconnectpeerassociation "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/transitgatewayconnectpeerassociation"
	transitgatewayregistration "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/transitgatewayregistration"
	vpcattachment "github.com/upbound/provider-aws/internal/controller/cluster/networkmanager/vpcattachment"
)

// Setup_networkmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachmentaccepter.Setup,
		connectattachment.Setup,
		connection.Setup,
		corenetwork.Setup,
		customergatewayassociation.Setup,
		device.Setup,
		globalnetwork.Setup,
		link.Setup,
		linkassociation.Setup,
		site.Setup,
		transitgatewayconnectpeerassociation.Setup,
		transitgatewayregistration.Setup,
		vpcattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
