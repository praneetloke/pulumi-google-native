// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the details of a Scope.
func LookupScope(ctx *pulumi.Context, args *LookupScopeArgs, opts ...pulumi.InvokeOption) (*LookupScopeResult, error) {
	var rv LookupScopeResult
	err := ctx.Invoke("google-native:gkehub/v1beta:getScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeArgs struct {
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	ScopeId  string  `pulumi:"scopeId"`
}

type LookupScopeResult struct {
	// When the scope was created.
	CreateTime string `pulumi:"createTime"`
	// When the scope was deleted.
	DeleteTime string `pulumi:"deleteTime"`
	// The resource name for the scope `projects/{project}/locations/{location}/scopes/{scope}`
	Name string `pulumi:"name"`
	// State of the scope resource.
	State ScopeLifecycleStateResponse `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all scope resources. If a scope resource is deleted and another resource with the same name is created, it gets a different uid.
	Uid string `pulumi:"uid"`
	// When the scope was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupScopeOutput(ctx *pulumi.Context, args LookupScopeOutputArgs, opts ...pulumi.InvokeOption) LookupScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScopeResult, error) {
			args := v.(LookupScopeArgs)
			r, err := LookupScope(ctx, &args, opts...)
			var s LookupScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScopeResultOutput)
}

type LookupScopeOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	ScopeId  pulumi.StringInput    `pulumi:"scopeId"`
}

func (LookupScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeArgs)(nil)).Elem()
}

type LookupScopeResultOutput struct{ *pulumi.OutputState }

func (LookupScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeResult)(nil)).Elem()
}

func (o LookupScopeResultOutput) ToLookupScopeResultOutput() LookupScopeResultOutput {
	return o
}

func (o LookupScopeResultOutput) ToLookupScopeResultOutputWithContext(ctx context.Context) LookupScopeResultOutput {
	return o
}

// When the scope was created.
func (o LookupScopeResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// When the scope was deleted.
func (o LookupScopeResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// The resource name for the scope `projects/{project}/locations/{location}/scopes/{scope}`
func (o LookupScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of the scope resource.
func (o LookupScopeResultOutput) State() ScopeLifecycleStateResponseOutput {
	return o.ApplyT(func(v LookupScopeResult) ScopeLifecycleStateResponse { return v.State }).(ScopeLifecycleStateResponseOutput)
}

// Google-generated UUID for this resource. This is unique across all scope resources. If a scope resource is deleted and another resource with the same name is created, it gets a different uid.
func (o LookupScopeResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeResult) string { return v.Uid }).(pulumi.StringOutput)
}

// When the scope was last updated.
func (o LookupScopeResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScopeResultOutput{})
}
