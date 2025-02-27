// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/upbound/provider-aws/internal/controller/cluster/rds/cluster"
	clusteractivitystream "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusteractivitystream"
	clusterendpoint "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusterendpoint"
	clusterinstance "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusterinstance"
	clusterparametergroup "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusterparametergroup"
	clusterroleassociation "github.com/upbound/provider-aws/internal/controller/cluster/rds/clusterroleassociation"
	clustersnapshot "github.com/upbound/provider-aws/internal/controller/cluster/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "github.com/upbound/provider-aws/internal/controller/cluster/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "github.com/upbound/provider-aws/internal/controller/cluster/rds/dbsnapshotcopy"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/cluster/rds/eventsubscription"
	globalcluster "github.com/upbound/provider-aws/internal/controller/cluster/rds/globalcluster"
	instance "github.com/upbound/provider-aws/internal/controller/cluster/rds/instance"
	instanceroleassociation "github.com/upbound/provider-aws/internal/controller/cluster/rds/instanceroleassociation"
	optiongroup "github.com/upbound/provider-aws/internal/controller/cluster/rds/optiongroup"
	parametergroup "github.com/upbound/provider-aws/internal/controller/cluster/rds/parametergroup"
	proxy "github.com/upbound/provider-aws/internal/controller/cluster/rds/proxy"
	proxydefaulttargetgroup "github.com/upbound/provider-aws/internal/controller/cluster/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/upbound/provider-aws/internal/controller/cluster/rds/proxyendpoint"
	proxytarget "github.com/upbound/provider-aws/internal/controller/cluster/rds/proxytarget"
	snapshot "github.com/upbound/provider-aws/internal/controller/cluster/rds/snapshot"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/cluster/rds/subnetgroup"
)

// Setup_rds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_rds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusteractivitystream.Setup,
		clusterendpoint.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clusterroleassociation.Setup,
		clustersnapshot.Setup,
		dbinstanceautomatedbackupsreplication.Setup,
		dbsnapshotcopy.Setup,
		eventsubscription.Setup,
		globalcluster.Setup,
		instance.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergroup.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		snapshot.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
