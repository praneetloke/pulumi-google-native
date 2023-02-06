// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Instance.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:apigeeregistry/v1:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupInstanceResult struct {
	// Build info of the Instance if it's in `ACTIVE` state.
	Build BuildResponse `pulumi:"build"`
	// Config of the Instance.
	Config ConfigResponse `pulumi:"config"`
	// Creation timestamp.
	CreateTime string `pulumi:"createTime"`
	// Format: `projects/*/locations/*/instance`. Currently only `locations/global` is supported.
	Name string `pulumi:"name"`
	// The current state of the Instance.
	State string `pulumi:"state"`
	// Extra information of Instance.State if the state is `FAILED`.
	StateMessage string `pulumi:"stateMessage"`
	// Last update timestamp.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// Build info of the Instance if it's in `ACTIVE` state.
func (o LookupInstanceResultOutput) Build() BuildResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) BuildResponse { return v.Build }).(BuildResponseOutput)
}

// Config of the Instance.
func (o LookupInstanceResultOutput) Config() ConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) ConfigResponse { return v.Config }).(ConfigResponseOutput)
}

// Creation timestamp.
func (o LookupInstanceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Format: `projects/*/locations/*/instance`. Currently only `locations/global` is supported.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current state of the Instance.
func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

// Extra information of Instance.State if the state is `FAILED`.
func (o LookupInstanceResultOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.StateMessage }).(pulumi.StringOutput)
}

// Last update timestamp.
func (o LookupInstanceResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
