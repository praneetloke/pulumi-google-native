// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns an execution of the given name.
func LookupExecution(ctx *pulumi.Context, args *LookupExecutionArgs, opts ...pulumi.InvokeOption) (*LookupExecutionResult, error) {
	var rv LookupExecutionResult
	err := ctx.Invoke("google-native:workflowexecutions/v1:getExecution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExecutionArgs struct {
	ExecutionId string  `pulumi:"executionId"`
	Location    string  `pulumi:"location"`
	Project     *string `pulumi:"project"`
	View        *string `pulumi:"view"`
	WorkflowId  string  `pulumi:"workflowId"`
}

type LookupExecutionResult struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument string `pulumi:"argument"`
	// The call logging level associated to this execution.
	CallLogLevel string `pulumi:"callLogLevel"`
	// Marks the end of execution, successful or not.
	EndTime string `pulumi:"endTime"`
	// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
	Error ErrorResponse `pulumi:"error"`
	// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
	Name string `pulumi:"name"`
	// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
	Result string `pulumi:"result"`
	// Marks the beginning of execution.
	StartTime string `pulumi:"startTime"`
	// Current state of the execution.
	State string `pulumi:"state"`
	// Status tracks the current steps and progress data of this execution.
	Status StatusResponse `pulumi:"status"`
	// Revision of the workflow this execution is using.
	WorkflowRevisionId string `pulumi:"workflowRevisionId"`
}

func LookupExecutionOutput(ctx *pulumi.Context, args LookupExecutionOutputArgs, opts ...pulumi.InvokeOption) LookupExecutionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExecutionResult, error) {
			args := v.(LookupExecutionArgs)
			r, err := LookupExecution(ctx, &args, opts...)
			var s LookupExecutionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExecutionResultOutput)
}

type LookupExecutionOutputArgs struct {
	ExecutionId pulumi.StringInput    `pulumi:"executionId"`
	Location    pulumi.StringInput    `pulumi:"location"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
	View        pulumi.StringPtrInput `pulumi:"view"`
	WorkflowId  pulumi.StringInput    `pulumi:"workflowId"`
}

func (LookupExecutionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExecutionArgs)(nil)).Elem()
}

type LookupExecutionResultOutput struct{ *pulumi.OutputState }

func (LookupExecutionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExecutionResult)(nil)).Elem()
}

func (o LookupExecutionResultOutput) ToLookupExecutionResultOutput() LookupExecutionResultOutput {
	return o
}

func (o LookupExecutionResultOutput) ToLookupExecutionResultOutputWithContext(ctx context.Context) LookupExecutionResultOutput {
	return o
}

// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
func (o LookupExecutionResultOutput) Argument() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExecutionResult) string { return v.Argument }).(pulumi.StringOutput)
}

// The call logging level associated to this execution.
func (o LookupExecutionResultOutput) CallLogLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExecutionResult) string { return v.CallLogLevel }).(pulumi.StringOutput)
}

// Marks the end of execution, successful or not.
func (o LookupExecutionResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExecutionResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
func (o LookupExecutionResultOutput) Error() ErrorResponseOutput {
	return o.ApplyT(func(v LookupExecutionResult) ErrorResponse { return v.Error }).(ErrorResponseOutput)
}

// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
func (o LookupExecutionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExecutionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
func (o LookupExecutionResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExecutionResult) string { return v.Result }).(pulumi.StringOutput)
}

// Marks the beginning of execution.
func (o LookupExecutionResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExecutionResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// Current state of the execution.
func (o LookupExecutionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExecutionResult) string { return v.State }).(pulumi.StringOutput)
}

// Status tracks the current steps and progress data of this execution.
func (o LookupExecutionResultOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v LookupExecutionResult) StatusResponse { return v.Status }).(StatusResponseOutput)
}

// Revision of the workflow this execution is using.
func (o LookupExecutionResultOutput) WorkflowRevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExecutionResult) string { return v.WorkflowRevisionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExecutionResultOutput{})
}
