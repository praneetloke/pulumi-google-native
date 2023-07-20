// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a job trigger. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.
func LookupJobTrigger(ctx *pulumi.Context, args *LookupJobTriggerArgs, opts ...pulumi.InvokeOption) (*LookupJobTriggerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupJobTriggerResult
	err := ctx.Invoke("google-native:dlp/v2:getJobTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobTriggerArgs struct {
	JobTriggerId string  `pulumi:"jobTriggerId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupJobTriggerResult struct {
	// The creation timestamp of a triggeredJob.
	CreateTime string `pulumi:"createTime"`
	// User provided description (max 256 chars)
	Description string `pulumi:"description"`
	// Display name (max 100 chars)
	DisplayName string `pulumi:"displayName"`
	// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
	Errors []GooglePrivacyDlpV2ErrorResponse `pulumi:"errors"`
	// For inspect jobs, a snapshot of the configuration.
	InspectJob GooglePrivacyDlpV2InspectJobConfigResponse `pulumi:"inspectJob"`
	// The timestamp of the last time this trigger executed.
	LastRunTime string `pulumi:"lastRunTime"`
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name string `pulumi:"name"`
	// A status for this trigger.
	Status string `pulumi:"status"`
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers []GooglePrivacyDlpV2TriggerResponse `pulumi:"triggers"`
	// The last update timestamp of a triggeredJob.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupJobTriggerOutput(ctx *pulumi.Context, args LookupJobTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupJobTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobTriggerResult, error) {
			args := v.(LookupJobTriggerArgs)
			r, err := LookupJobTrigger(ctx, &args, opts...)
			var s LookupJobTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobTriggerResultOutput)
}

type LookupJobTriggerOutputArgs struct {
	JobTriggerId pulumi.StringInput    `pulumi:"jobTriggerId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupJobTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobTriggerArgs)(nil)).Elem()
}

type LookupJobTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupJobTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobTriggerResult)(nil)).Elem()
}

func (o LookupJobTriggerResultOutput) ToLookupJobTriggerResultOutput() LookupJobTriggerResultOutput {
	return o
}

func (o LookupJobTriggerResultOutput) ToLookupJobTriggerResultOutputWithContext(ctx context.Context) LookupJobTriggerResultOutput {
	return o
}

// The creation timestamp of a triggeredJob.
func (o LookupJobTriggerResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// User provided description (max 256 chars)
func (o LookupJobTriggerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) string { return v.Description }).(pulumi.StringOutput)
}

// Display name (max 100 chars)
func (o LookupJobTriggerResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
func (o LookupJobTriggerResultOutput) Errors() GooglePrivacyDlpV2ErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) []GooglePrivacyDlpV2ErrorResponse { return v.Errors }).(GooglePrivacyDlpV2ErrorResponseArrayOutput)
}

// For inspect jobs, a snapshot of the configuration.
func (o LookupJobTriggerResultOutput) InspectJob() GooglePrivacyDlpV2InspectJobConfigResponseOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) GooglePrivacyDlpV2InspectJobConfigResponse { return v.InspectJob }).(GooglePrivacyDlpV2InspectJobConfigResponseOutput)
}

// The timestamp of the last time this trigger executed.
func (o LookupJobTriggerResultOutput) LastRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) string { return v.LastRunTime }).(pulumi.StringOutput)
}

// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
func (o LookupJobTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

// A status for this trigger.
func (o LookupJobTriggerResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) string { return v.Status }).(pulumi.StringOutput)
}

// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
func (o LookupJobTriggerResultOutput) Triggers() GooglePrivacyDlpV2TriggerResponseArrayOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) []GooglePrivacyDlpV2TriggerResponse { return v.Triggers }).(GooglePrivacyDlpV2TriggerResponseArrayOutput)
}

// The last update timestamp of a triggeredJob.
func (o LookupJobTriggerResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobTriggerResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobTriggerResultOutput{})
}
