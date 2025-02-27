// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/upbound/provider-aws/internal/controller/namespaced/appconfig/application"
	configurationprofile "github.com/upbound/provider-aws/internal/controller/namespaced/appconfig/configurationprofile"
	deployment "github.com/upbound/provider-aws/internal/controller/namespaced/appconfig/deployment"
	deploymentstrategy "github.com/upbound/provider-aws/internal/controller/namespaced/appconfig/deploymentstrategy"
	environment "github.com/upbound/provider-aws/internal/controller/namespaced/appconfig/environment"
	extension "github.com/upbound/provider-aws/internal/controller/namespaced/appconfig/extension"
	extensionassociation "github.com/upbound/provider-aws/internal/controller/namespaced/appconfig/extensionassociation"
	hostedconfigurationversion "github.com/upbound/provider-aws/internal/controller/namespaced/appconfig/hostedconfigurationversion"
)

// Setup_appconfig creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appconfig(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		configurationprofile.Setup,
		deployment.Setup,
		deploymentstrategy.Setup,
		environment.Setup,
		extension.Setup,
		extensionassociation.Setup,
		hostedconfigurationversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
