// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an issue model.
func LookupIssueModel(ctx *pulumi.Context, args *LookupIssueModelArgs, opts ...pulumi.InvokeOption) (*LookupIssueModelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIssueModelResult
	err := ctx.Invoke("google-native:contactcenterinsights/v1:getIssueModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIssueModelArgs struct {
	IssueModelId string  `pulumi:"issueModelId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupIssueModelResult struct {
	// The time at which this issue model was created.
	CreateTime string `pulumi:"createTime"`
	// The representative name for the issue model.
	DisplayName string `pulumi:"displayName"`
	// Configs for the input data that used to create the issue model.
	InputDataConfig GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigResponse `pulumi:"inputDataConfig"`
	// Number of issues in this issue model.
	IssueCount string `pulumi:"issueCount"`
	// Immutable. The resource name of the issue model. Format: projects/{project}/locations/{location}/issueModels/{issue_model}
	Name string `pulumi:"name"`
	// State of the model.
	State string `pulumi:"state"`
	// Immutable. The issue model's label statistics on its training data.
	TrainingStats GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponse `pulumi:"trainingStats"`
	// The most recent time at which the issue model was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupIssueModelOutput(ctx *pulumi.Context, args LookupIssueModelOutputArgs, opts ...pulumi.InvokeOption) LookupIssueModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIssueModelResult, error) {
			args := v.(LookupIssueModelArgs)
			r, err := LookupIssueModel(ctx, &args, opts...)
			var s LookupIssueModelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIssueModelResultOutput)
}

type LookupIssueModelOutputArgs struct {
	IssueModelId pulumi.StringInput    `pulumi:"issueModelId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupIssueModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIssueModelArgs)(nil)).Elem()
}

type LookupIssueModelResultOutput struct{ *pulumi.OutputState }

func (LookupIssueModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIssueModelResult)(nil)).Elem()
}

func (o LookupIssueModelResultOutput) ToLookupIssueModelResultOutput() LookupIssueModelResultOutput {
	return o
}

func (o LookupIssueModelResultOutput) ToLookupIssueModelResultOutputWithContext(ctx context.Context) LookupIssueModelResultOutput {
	return o
}

// The time at which this issue model was created.
func (o LookupIssueModelResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueModelResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The representative name for the issue model.
func (o LookupIssueModelResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueModelResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Configs for the input data that used to create the issue model.
func (o LookupIssueModelResultOutput) InputDataConfig() GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigResponseOutput {
	return o.ApplyT(func(v LookupIssueModelResult) GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigResponse {
		return v.InputDataConfig
	}).(GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigResponseOutput)
}

// Number of issues in this issue model.
func (o LookupIssueModelResultOutput) IssueCount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueModelResult) string { return v.IssueCount }).(pulumi.StringOutput)
}

// Immutable. The resource name of the issue model. Format: projects/{project}/locations/{location}/issueModels/{issue_model}
func (o LookupIssueModelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueModelResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of the model.
func (o LookupIssueModelResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueModelResult) string { return v.State }).(pulumi.StringOutput)
}

// Immutable. The issue model's label statistics on its training data.
func (o LookupIssueModelResultOutput) TrainingStats() GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponseOutput {
	return o.ApplyT(func(v LookupIssueModelResult) GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponse {
		return v.TrainingStats
	}).(GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponseOutput)
}

// The most recent time at which the issue model was updated.
func (o LookupIssueModelResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIssueModelResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIssueModelResultOutput{})
}
