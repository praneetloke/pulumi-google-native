// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single UrlList.
func LookupUrlList(ctx *pulumi.Context, args *LookupUrlListArgs, opts ...pulumi.InvokeOption) (*LookupUrlListResult, error) {
	var rv LookupUrlListResult
	err := ctx.Invoke("google-native:networksecurity/v1beta1:getUrlList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUrlListArgs struct {
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	UrlListId string  `pulumi:"urlListId"`
}

type LookupUrlListResult struct {
	// Time when the security policy was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. Free-text description of the resource.
	Description string `pulumi:"description"`
	// Name of the resource provided by the user. Name is of the form projects/{project}/locations/{location}/urlLists/{url_list} url_list should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name string `pulumi:"name"`
	// Time when the security policy was updated.
	UpdateTime string `pulumi:"updateTime"`
	// FQDNs and URLs.
	Values []string `pulumi:"values"`
}

func LookupUrlListOutput(ctx *pulumi.Context, args LookupUrlListOutputArgs, opts ...pulumi.InvokeOption) LookupUrlListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUrlListResult, error) {
			args := v.(LookupUrlListArgs)
			r, err := LookupUrlList(ctx, &args, opts...)
			var s LookupUrlListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUrlListResultOutput)
}

type LookupUrlListOutputArgs struct {
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	UrlListId pulumi.StringInput    `pulumi:"urlListId"`
}

func (LookupUrlListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUrlListArgs)(nil)).Elem()
}

type LookupUrlListResultOutput struct{ *pulumi.OutputState }

func (LookupUrlListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUrlListResult)(nil)).Elem()
}

func (o LookupUrlListResultOutput) ToLookupUrlListResultOutput() LookupUrlListResultOutput {
	return o
}

func (o LookupUrlListResultOutput) ToLookupUrlListResultOutputWithContext(ctx context.Context) LookupUrlListResultOutput {
	return o
}

// Time when the security policy was created.
func (o LookupUrlListResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUrlListResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Free-text description of the resource.
func (o LookupUrlListResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUrlListResult) string { return v.Description }).(pulumi.StringOutput)
}

// Name of the resource provided by the user. Name is of the form projects/{project}/locations/{location}/urlLists/{url_list} url_list should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
func (o LookupUrlListResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUrlListResult) string { return v.Name }).(pulumi.StringOutput)
}

// Time when the security policy was updated.
func (o LookupUrlListResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUrlListResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// FQDNs and URLs.
func (o LookupUrlListResultOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUrlListResult) []string { return v.Values }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUrlListResultOutput{})
}
