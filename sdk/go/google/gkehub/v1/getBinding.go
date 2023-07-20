// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the details of a MembershipBinding.
func LookupBinding(ctx *pulumi.Context, args *LookupBindingArgs, opts ...pulumi.InvokeOption) (*LookupBindingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBindingResult
	err := ctx.Invoke("google-native:gkehub/v1:getBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBindingArgs struct {
	BindingId    string  `pulumi:"bindingId"`
	Location     string  `pulumi:"location"`
	MembershipId string  `pulumi:"membershipId"`
	Project      *string `pulumi:"project"`
}

type LookupBindingResult struct {
	// When the membership binding was created.
	CreateTime string `pulumi:"createTime"`
	// When the membership binding was deleted.
	DeleteTime string `pulumi:"deleteTime"`
	// Whether the membershipbinding is Fleet-wide; true means that this Membership should be bound to all Namespaces in this entire Fleet.
	Fleet bool `pulumi:"fleet"`
	// The resource name for the membershipbinding itself `projects/{project}/locations/{location}/memberships/{membership}/bindings/{membershipbinding}`
	Name string `pulumi:"name"`
	// A Workspace resource name in the format `projects/*/locations/*/scopes/*`.
	Scope string `pulumi:"scope"`
	// State of the membership binding resource.
	State MembershipBindingLifecycleStateResponse `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all membershipbinding resources. If a membershipbinding resource is deleted and another resource with the same name is created, it gets a different uid.
	Uid string `pulumi:"uid"`
	// When the membership binding was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupBindingOutput(ctx *pulumi.Context, args LookupBindingOutputArgs, opts ...pulumi.InvokeOption) LookupBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBindingResult, error) {
			args := v.(LookupBindingArgs)
			r, err := LookupBinding(ctx, &args, opts...)
			var s LookupBindingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBindingResultOutput)
}

type LookupBindingOutputArgs struct {
	BindingId    pulumi.StringInput    `pulumi:"bindingId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	MembershipId pulumi.StringInput    `pulumi:"membershipId"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBindingArgs)(nil)).Elem()
}

type LookupBindingResultOutput struct{ *pulumi.OutputState }

func (LookupBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBindingResult)(nil)).Elem()
}

func (o LookupBindingResultOutput) ToLookupBindingResultOutput() LookupBindingResultOutput {
	return o
}

func (o LookupBindingResultOutput) ToLookupBindingResultOutputWithContext(ctx context.Context) LookupBindingResultOutput {
	return o
}

// When the membership binding was created.
func (o LookupBindingResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// When the membership binding was deleted.
func (o LookupBindingResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// Whether the membershipbinding is Fleet-wide; true means that this Membership should be bound to all Namespaces in this entire Fleet.
func (o LookupBindingResultOutput) Fleet() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBindingResult) bool { return v.Fleet }).(pulumi.BoolOutput)
}

// The resource name for the membershipbinding itself `projects/{project}/locations/{location}/memberships/{membership}/bindings/{membershipbinding}`
func (o LookupBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

// A Workspace resource name in the format `projects/*/locations/*/scopes/*`.
func (o LookupBindingResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Scope }).(pulumi.StringOutput)
}

// State of the membership binding resource.
func (o LookupBindingResultOutput) State() MembershipBindingLifecycleStateResponseOutput {
	return o.ApplyT(func(v LookupBindingResult) MembershipBindingLifecycleStateResponse { return v.State }).(MembershipBindingLifecycleStateResponseOutput)
}

// Google-generated UUID for this resource. This is unique across all membershipbinding resources. If a membershipbinding resource is deleted and another resource with the same name is created, it gets a different uid.
func (o LookupBindingResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Uid }).(pulumi.StringOutput)
}

// When the membership binding was last updated.
func (o LookupBindingResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBindingResultOutput{})
}
