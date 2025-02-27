// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package directorybucket

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/controller/handler"
	"github.com/crossplane/upjet/pkg/metrics"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"

	v1beta1 "github.com/upbound/provider-aws/apis/cluster/s3/v1beta1"
	features "github.com/upbound/provider-aws/internal/features"
)

// Setup adds a controller that reconciles DirectoryBucket managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1beta1.DirectoryBucket_GroupVersionKind.String())
	var initializers managed.InitializerChain
	initializers = append(initializers, managed.NewNameAsExternalName(mgr.GetClient()))
	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.SecretStoreConfigGVK != nil {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), *o.SecretStoreConfigGVK, connection.WithTLSConfig(o.ESSOptions.TLSConfig)))
	}
	eventHandler := handler.NewEventHandler(handler.WithLogger(o.Logger.WithValues("gvk", v1beta1.DirectoryBucket_GroupVersionKind)))
	ac := tjcontroller.NewAPICallbacks(mgr, xpresource.ManagedKind(v1beta1.DirectoryBucket_GroupVersionKind), tjcontroller.WithEventHandler(eventHandler), tjcontroller.WithStatusUpdates(false))
	opts := []managed.ReconcilerOption{
		managed.WithExternalConnecter(
			tjcontroller.NewTerraformPluginFrameworkAsyncConnector(mgr.GetClient(), o.OperationTrackerStore, o.SetupFn, o.Provider.Resources["aws_s3_directory_bucket"],
				tjcontroller.WithTerraformPluginFrameworkAsyncLogger(o.Logger),
				tjcontroller.WithTerraformPluginFrameworkAsyncConnectorEventHandler(eventHandler),
				tjcontroller.WithTerraformPluginFrameworkAsyncCallbackProvider(ac),
				tjcontroller.WithTerraformPluginFrameworkAsyncMetricRecorder(metrics.NewMetricRecorder(v1beta1.DirectoryBucket_GroupVersionKind, mgr, o.PollInterval)),
				tjcontroller.WithTerraformPluginFrameworkAsyncManagementPolicies(o.Features.Enabled(features.EnableBetaManagementPolicies)))),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(tjcontroller.NewOperationTrackerFinalizer(o.OperationTrackerStore, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3 * time.Minute),
		managed.WithInitializers(initializers),
		managed.WithConnectionPublishers(cps...),
		managed.WithPollInterval(o.PollInterval),
	}
	if o.PollJitter != 0 {
		opts = append(opts, managed.WithPollJitterHook(o.PollJitter))
	}
	if o.Features.Enabled(features.EnableBetaManagementPolicies) {
		opts = append(opts, managed.WithManagementPolicies())
	}
	if o.MetricOptions != nil {
		opts = append(opts, managed.WithMetricRecorder(o.MetricOptions.MRMetrics))
	}

	// register webhooks for the kind v1beta1.DirectoryBucket
	// if they're enabled.
	if o.StartWebhooks {
		if err := ctrl.NewWebhookManagedBy(mgr).
			For(&v1beta1.DirectoryBucket{}).
			Complete(); err != nil {
			return errors.Wrap(err, "cannot register webhook for the kind v1beta1.DirectoryBucket")
		}
	}

	if o.MetricOptions != nil && o.MetricOptions.MRStateMetrics != nil {
		stateMetricsRecorder := statemetrics.NewMRStateRecorder(
			mgr.GetClient(), o.Logger, o.MetricOptions.MRStateMetrics, &v1beta1.DirectoryBucketList{}, o.MetricOptions.PollStateMetricInterval,
		)
		if err := mgr.Add(stateMetricsRecorder); err != nil {
			return errors.Wrap(err, "cannot register MR state metrics recorder for kind v1beta1.DirectoryBucketList")
		}
	}

	r := managed.NewReconciler(mgr, xpresource.ManagedKind(v1beta1.DirectoryBucket_GroupVersionKind), opts...)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(xpresource.DesiredStateChanged()).
		Watches(&v1beta1.DirectoryBucket{}, eventHandler).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}
