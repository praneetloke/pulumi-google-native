// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single address group.
func LookupAddressGroup(ctx *pulumi.Context, args *LookupAddressGroupArgs, opts ...pulumi.InvokeOption) (*LookupAddressGroupResult, error) {
	var rv LookupAddressGroupResult
	err := ctx.Invoke("google-native:networksecurity/v1beta1:getAddressGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAddressGroupArgs struct {
	AddressGroupId string  `pulumi:"addressGroupId"`
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
}

type LookupAddressGroupResult struct {
	// Capacity of the Address Group
	Capacity int `pulumi:"capacity"`
	// The timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. Free-text description of the resource.
	Description string `pulumi:"description"`
	// Optional. List of items.
	Items []string `pulumi:"items"`
	// Optional. Set of label tags associated with the AddressGroup resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the AddressGroup resource. It matches pattern `projects/*/locations/{location}/addressGroups/`.
	Name string `pulumi:"name"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink string `pulumi:"selfLink"`
	// The type of the Address Group. Possible values are "IPv4" or "IPV6".
	Type string `pulumi:"type"`
	// The timestamp when the resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupAddressGroupOutput(ctx *pulumi.Context, args LookupAddressGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAddressGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAddressGroupResult, error) {
			args := v.(LookupAddressGroupArgs)
			r, err := LookupAddressGroup(ctx, &args, opts...)
			var s LookupAddressGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAddressGroupResultOutput)
}

type LookupAddressGroupOutputArgs struct {
	AddressGroupId pulumi.StringInput    `pulumi:"addressGroupId"`
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupAddressGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddressGroupArgs)(nil)).Elem()
}

type LookupAddressGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAddressGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddressGroupResult)(nil)).Elem()
}

func (o LookupAddressGroupResultOutput) ToLookupAddressGroupResultOutput() LookupAddressGroupResultOutput {
	return o
}

func (o LookupAddressGroupResultOutput) ToLookupAddressGroupResultOutputWithContext(ctx context.Context) LookupAddressGroupResultOutput {
	return o
}

// Capacity of the Address Group
func (o LookupAddressGroupResultOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) int { return v.Capacity }).(pulumi.IntOutput)
}

// The timestamp when the resource was created.
func (o LookupAddressGroupResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Free-text description of the resource.
func (o LookupAddressGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. List of items.
func (o LookupAddressGroupResultOutput) Items() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) []string { return v.Items }).(pulumi.StringArrayOutput)
}

// Optional. Set of label tags associated with the AddressGroup resource.
func (o LookupAddressGroupResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the AddressGroup resource. It matches pattern `projects/*/locations/{location}/addressGroups/`.
func (o LookupAddressGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Server-defined fully-qualified URL for this resource.
func (o LookupAddressGroupResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The type of the Address Group. Possible values are "IPv4" or "IPV6".
func (o LookupAddressGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o LookupAddressGroupResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressGroupResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAddressGroupResultOutput{})
}
