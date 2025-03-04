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
	common "github.com/upbound/provider-aws/config/common"
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn),
							Extract:      common.ARNExtractor(),
							Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnRef,
							Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Object", "ObjectList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey),
							Extract:      resource.ExtractParamPath("key", false),
							Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeyRef,
							Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeySelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeyRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("kinesis.aws.upbound.io", "v1beta1", "Stream", "StreamList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn),
							Extract:      common.TerraformID(),
							Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnRef,
							Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta1", "DeliveryStream", "DeliveryStreamList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn),
							Extract:      resource.ExtractParamPath("arn", false),
							Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnRef,
							Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn),
							Extract:      resource.ExtractParamPath("arn", true),
							Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnRef,
							Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn),
							Extract:      resource.ExtractParamPath("arn", true),
							Reference:    mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnRef,
							Selector:     mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn")
					}
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.CloudwatchLoggingOptions); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudwatchlogs.aws.upbound.io", "v1beta1", "Stream", "StreamList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnRef,
				Selector:     mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn")
		}
		mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CloudwatchLoggingOptions[i3].LogStreamArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceExecutionRole),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ServiceExecutionRoleRef,
			Selector:     mg.Spec.ForProvider.ServiceExecutionRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceExecutionRole")
	}
	mg.Spec.ForProvider.ServiceExecutionRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceExecutionRoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn),
							Extract:      common.ARNExtractor(),
							Reference:    mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnRef,
							Selector:     mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn")
					}
					mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].BucketArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Object", "ObjectList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey),
							Extract:      resource.ExtractParamPath("key", false),
							Reference:    mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeyRef,
							Selector:     mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeySelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey")
					}
					mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKey = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.ApplicationConfiguration[i3].ApplicationCodeConfiguration[i4].CodeContent[i5].S3ContentLocation[i6].FileKeyRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("kinesis.aws.upbound.io", "v1beta1", "Stream", "StreamList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn),
							Extract:      common.TerraformID(),
							Reference:    mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnRef,
							Selector:     mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn")
					}
					mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Input[i5].KinesisStreamsInput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta1", "DeliveryStream", "DeliveryStreamList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn),
							Extract:      resource.ExtractParamPath("arn", false),
							Reference:    mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnRef,
							Selector:     mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn")
					}
					mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].KinesisFirehoseOutput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn),
							Extract:      resource.ExtractParamPath("arn", true),
							Reference:    mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnRef,
							Selector:     mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn")
					}
					mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].Output[i5].LambdaOutput[i6].ResourceArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ApplicationConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn),
							Extract:      resource.ExtractParamPath("arn", true),
							Reference:    mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnRef,
							Selector:     mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn")
					}
					mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.ApplicationConfiguration[i3].SQLApplicationConfiguration[i4].ReferenceDataSource[i5].S3ReferenceDataSource[i6].BucketArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.CloudwatchLoggingOptions); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudwatchlogs.aws.upbound.io", "v1beta1", "Stream", "StreamList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CloudwatchLoggingOptions[i3].LogStreamArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.CloudwatchLoggingOptions[i3].LogStreamArnRef,
				Selector:     mg.Spec.InitProvider.CloudwatchLoggingOptions[i3].LogStreamArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.CloudwatchLoggingOptions[i3].LogStreamArn")
		}
		mg.Spec.InitProvider.CloudwatchLoggingOptions[i3].LogStreamArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.CloudwatchLoggingOptions[i3].LogStreamArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceExecutionRole),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ServiceExecutionRoleRef,
			Selector:     mg.Spec.InitProvider.ServiceExecutionRoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceExecutionRole")
	}
	mg.Spec.InitProvider.ServiceExecutionRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceExecutionRoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ApplicationSnapshot.
func (mg *ApplicationSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kinesisanalyticsv2.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ApplicationNameRef,
			Selector:     mg.Spec.ForProvider.ApplicationNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationName")
	}
	mg.Spec.ForProvider.ApplicationName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kinesisanalyticsv2.aws.upbound.io", "v1beta1", "Application", "ApplicationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ApplicationNameRef,
			Selector:     mg.Spec.InitProvider.ApplicationNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationName")
	}
	mg.Spec.InitProvider.ApplicationName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationNameRef = rsp.ResolvedReference

	return nil
}
