// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Optional. Describes the level of platform logging to apply to calls and call responses during executions of this workflow. If both the workflow and the execution specify a logging level, the execution level takes precedence.
type WorkflowCallLogLevel string

const (
	// No call logging level specified.
	WorkflowCallLogLevelCallLogLevelUnspecified = WorkflowCallLogLevel("CALL_LOG_LEVEL_UNSPECIFIED")
	// Log all call steps within workflows, all call returns, and all exceptions raised.
	WorkflowCallLogLevelLogAllCalls = WorkflowCallLogLevel("LOG_ALL_CALLS")
	// Log only exceptions that are raised from call steps within workflows.
	WorkflowCallLogLevelLogErrorsOnly = WorkflowCallLogLevel("LOG_ERRORS_ONLY")
	// Explicitly log nothing.
	WorkflowCallLogLevelLogNone = WorkflowCallLogLevel("LOG_NONE")
)

func (WorkflowCallLogLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowCallLogLevel)(nil)).Elem()
}

func (e WorkflowCallLogLevel) ToWorkflowCallLogLevelOutput() WorkflowCallLogLevelOutput {
	return pulumi.ToOutput(e).(WorkflowCallLogLevelOutput)
}

func (e WorkflowCallLogLevel) ToWorkflowCallLogLevelOutputWithContext(ctx context.Context) WorkflowCallLogLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkflowCallLogLevelOutput)
}

func (e WorkflowCallLogLevel) ToWorkflowCallLogLevelPtrOutput() WorkflowCallLogLevelPtrOutput {
	return e.ToWorkflowCallLogLevelPtrOutputWithContext(context.Background())
}

func (e WorkflowCallLogLevel) ToWorkflowCallLogLevelPtrOutputWithContext(ctx context.Context) WorkflowCallLogLevelPtrOutput {
	return WorkflowCallLogLevel(e).ToWorkflowCallLogLevelOutputWithContext(ctx).ToWorkflowCallLogLevelPtrOutputWithContext(ctx)
}

func (e WorkflowCallLogLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkflowCallLogLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkflowCallLogLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkflowCallLogLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkflowCallLogLevelOutput struct{ *pulumi.OutputState }

func (WorkflowCallLogLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowCallLogLevel)(nil)).Elem()
}

func (o WorkflowCallLogLevelOutput) ToWorkflowCallLogLevelOutput() WorkflowCallLogLevelOutput {
	return o
}

func (o WorkflowCallLogLevelOutput) ToWorkflowCallLogLevelOutputWithContext(ctx context.Context) WorkflowCallLogLevelOutput {
	return o
}

func (o WorkflowCallLogLevelOutput) ToWorkflowCallLogLevelPtrOutput() WorkflowCallLogLevelPtrOutput {
	return o.ToWorkflowCallLogLevelPtrOutputWithContext(context.Background())
}

func (o WorkflowCallLogLevelOutput) ToWorkflowCallLogLevelPtrOutputWithContext(ctx context.Context) WorkflowCallLogLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkflowCallLogLevel) *WorkflowCallLogLevel {
		return &v
	}).(WorkflowCallLogLevelPtrOutput)
}

func (o WorkflowCallLogLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkflowCallLogLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkflowCallLogLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkflowCallLogLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkflowCallLogLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkflowCallLogLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkflowCallLogLevelPtrOutput struct{ *pulumi.OutputState }

func (WorkflowCallLogLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowCallLogLevel)(nil)).Elem()
}

func (o WorkflowCallLogLevelPtrOutput) ToWorkflowCallLogLevelPtrOutput() WorkflowCallLogLevelPtrOutput {
	return o
}

func (o WorkflowCallLogLevelPtrOutput) ToWorkflowCallLogLevelPtrOutputWithContext(ctx context.Context) WorkflowCallLogLevelPtrOutput {
	return o
}

func (o WorkflowCallLogLevelPtrOutput) Elem() WorkflowCallLogLevelOutput {
	return o.ApplyT(func(v *WorkflowCallLogLevel) WorkflowCallLogLevel {
		if v != nil {
			return *v
		}
		var ret WorkflowCallLogLevel
		return ret
	}).(WorkflowCallLogLevelOutput)
}

func (o WorkflowCallLogLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkflowCallLogLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkflowCallLogLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkflowCallLogLevelInput is an input type that accepts WorkflowCallLogLevelArgs and WorkflowCallLogLevelOutput values.
// You can construct a concrete instance of `WorkflowCallLogLevelInput` via:
//
//	WorkflowCallLogLevelArgs{...}
type WorkflowCallLogLevelInput interface {
	pulumi.Input

	ToWorkflowCallLogLevelOutput() WorkflowCallLogLevelOutput
	ToWorkflowCallLogLevelOutputWithContext(context.Context) WorkflowCallLogLevelOutput
}

var workflowCallLogLevelPtrType = reflect.TypeOf((**WorkflowCallLogLevel)(nil)).Elem()

type WorkflowCallLogLevelPtrInput interface {
	pulumi.Input

	ToWorkflowCallLogLevelPtrOutput() WorkflowCallLogLevelPtrOutput
	ToWorkflowCallLogLevelPtrOutputWithContext(context.Context) WorkflowCallLogLevelPtrOutput
}

type workflowCallLogLevelPtr string

func WorkflowCallLogLevelPtr(v string) WorkflowCallLogLevelPtrInput {
	return (*workflowCallLogLevelPtr)(&v)
}

func (*workflowCallLogLevelPtr) ElementType() reflect.Type {
	return workflowCallLogLevelPtrType
}

func (in *workflowCallLogLevelPtr) ToWorkflowCallLogLevelPtrOutput() WorkflowCallLogLevelPtrOutput {
	return pulumi.ToOutput(in).(WorkflowCallLogLevelPtrOutput)
}

func (in *workflowCallLogLevelPtr) ToWorkflowCallLogLevelPtrOutputWithContext(ctx context.Context) WorkflowCallLogLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkflowCallLogLevelPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowCallLogLevelInput)(nil)).Elem(), WorkflowCallLogLevel("CALL_LOG_LEVEL_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowCallLogLevelPtrInput)(nil)).Elem(), WorkflowCallLogLevel("CALL_LOG_LEVEL_UNSPECIFIED"))
	pulumi.RegisterOutputType(WorkflowCallLogLevelOutput{})
	pulumi.RegisterOutputType(WorkflowCallLogLevelPtrOutput{})
}
