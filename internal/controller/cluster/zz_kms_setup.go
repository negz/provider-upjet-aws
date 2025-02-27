// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alias "github.com/upbound/provider-aws/internal/controller/cluster/kms/alias"
	ciphertext "github.com/upbound/provider-aws/internal/controller/cluster/kms/ciphertext"
	externalkey "github.com/upbound/provider-aws/internal/controller/cluster/kms/externalkey"
	grant "github.com/upbound/provider-aws/internal/controller/cluster/kms/grant"
	key "github.com/upbound/provider-aws/internal/controller/cluster/kms/key"
	replicaexternalkey "github.com/upbound/provider-aws/internal/controller/cluster/kms/replicaexternalkey"
	replicakey "github.com/upbound/provider-aws/internal/controller/cluster/kms/replicakey"
)

// Setup_kms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		replicaexternalkey.Setup,
		replicakey.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
