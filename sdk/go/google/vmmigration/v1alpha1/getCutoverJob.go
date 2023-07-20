// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single CutoverJob.
func LookupCutoverJob(ctx *pulumi.Context, args *LookupCutoverJobArgs, opts ...pulumi.InvokeOption) (*LookupCutoverJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCutoverJobResult
	err := ctx.Invoke("google-native:vmmigration/v1alpha1:getCutoverJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCutoverJobArgs struct {
	CutoverJobId  string  `pulumi:"cutoverJobId"`
	Location      string  `pulumi:"location"`
	MigratingVmId string  `pulumi:"migratingVmId"`
	Project       *string `pulumi:"project"`
	SourceId      string  `pulumi:"sourceId"`
}

type LookupCutoverJobResult struct {
	// Details of the target VM in Compute Engine.
	ComputeEngineTargetDetails ComputeEngineTargetDetailsResponse `pulumi:"computeEngineTargetDetails"`
	// Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_details instead.
	//
	// Deprecated: Output only. Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_details instead.
	ComputeEngineVmDetails TargetVMDetailsResponse `pulumi:"computeEngineVmDetails"`
	// The time the cutover job was created (as an API call, not when it was actually created in the target).
	CreateTime string `pulumi:"createTime"`
	// The time the cutover job had finished.
	EndTime string `pulumi:"endTime"`
	// Provides details for the errors that led to the Cutover Job's state.
	Error StatusResponse `pulumi:"error"`
	// The name of the cutover job.
	Name string `pulumi:"name"`
	// The current progress in percentage of the cutover job.
	Progress int `pulumi:"progress"`
	// The current progress in percentage of the cutover job.
	ProgressPercent int `pulumi:"progressPercent"`
	// State of the cutover job.
	State string `pulumi:"state"`
	// A message providing possible extra details about the current state.
	StateMessage string `pulumi:"stateMessage"`
	// The time the state was last updated.
	StateTime string `pulumi:"stateTime"`
	// The cutover steps list representing its progress.
	Steps []CutoverStepResponse `pulumi:"steps"`
	// Details of the VM to create as the target of this cutover job. Deprecated: Use compute_engine_target_details instead.
	//
	// Deprecated: Output only. Details of the VM to create as the target of this cutover job. Deprecated: Use compute_engine_target_details instead.
	TargetDetails TargetVMDetailsResponse `pulumi:"targetDetails"`
}

func LookupCutoverJobOutput(ctx *pulumi.Context, args LookupCutoverJobOutputArgs, opts ...pulumi.InvokeOption) LookupCutoverJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCutoverJobResult, error) {
			args := v.(LookupCutoverJobArgs)
			r, err := LookupCutoverJob(ctx, &args, opts...)
			var s LookupCutoverJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCutoverJobResultOutput)
}

type LookupCutoverJobOutputArgs struct {
	CutoverJobId  pulumi.StringInput    `pulumi:"cutoverJobId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	MigratingVmId pulumi.StringInput    `pulumi:"migratingVmId"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
	SourceId      pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupCutoverJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCutoverJobArgs)(nil)).Elem()
}

type LookupCutoverJobResultOutput struct{ *pulumi.OutputState }

func (LookupCutoverJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCutoverJobResult)(nil)).Elem()
}

func (o LookupCutoverJobResultOutput) ToLookupCutoverJobResultOutput() LookupCutoverJobResultOutput {
	return o
}

func (o LookupCutoverJobResultOutput) ToLookupCutoverJobResultOutputWithContext(ctx context.Context) LookupCutoverJobResultOutput {
	return o
}

// Details of the target VM in Compute Engine.
func (o LookupCutoverJobResultOutput) ComputeEngineTargetDetails() ComputeEngineTargetDetailsResponseOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) ComputeEngineTargetDetailsResponse { return v.ComputeEngineTargetDetails }).(ComputeEngineTargetDetailsResponseOutput)
}

// Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_details instead.
//
// Deprecated: Output only. Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_details instead.
func (o LookupCutoverJobResultOutput) ComputeEngineVmDetails() TargetVMDetailsResponseOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) TargetVMDetailsResponse { return v.ComputeEngineVmDetails }).(TargetVMDetailsResponseOutput)
}

// The time the cutover job was created (as an API call, not when it was actually created in the target).
func (o LookupCutoverJobResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The time the cutover job had finished.
func (o LookupCutoverJobResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// Provides details for the errors that led to the Cutover Job's state.
func (o LookupCutoverJobResultOutput) Error() StatusResponseOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) StatusResponse { return v.Error }).(StatusResponseOutput)
}

// The name of the cutover job.
func (o LookupCutoverJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current progress in percentage of the cutover job.
func (o LookupCutoverJobResultOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) int { return v.Progress }).(pulumi.IntOutput)
}

// The current progress in percentage of the cutover job.
func (o LookupCutoverJobResultOutput) ProgressPercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) int { return v.ProgressPercent }).(pulumi.IntOutput)
}

// State of the cutover job.
func (o LookupCutoverJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) string { return v.State }).(pulumi.StringOutput)
}

// A message providing possible extra details about the current state.
func (o LookupCutoverJobResultOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) string { return v.StateMessage }).(pulumi.StringOutput)
}

// The time the state was last updated.
func (o LookupCutoverJobResultOutput) StateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) string { return v.StateTime }).(pulumi.StringOutput)
}

// The cutover steps list representing its progress.
func (o LookupCutoverJobResultOutput) Steps() CutoverStepResponseArrayOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) []CutoverStepResponse { return v.Steps }).(CutoverStepResponseArrayOutput)
}

// Details of the VM to create as the target of this cutover job. Deprecated: Use compute_engine_target_details instead.
//
// Deprecated: Output only. Details of the VM to create as the target of this cutover job. Deprecated: Use compute_engine_target_details instead.
func (o LookupCutoverJobResultOutput) TargetDetails() TargetVMDetailsResponseOutput {
	return o.ApplyT(func(v LookupCutoverJobResult) TargetVMDetailsResponse { return v.TargetDetails }).(TargetVMDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCutoverJobResultOutput{})
}
