// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single ServiceConnectionToken.
func LookupServiceConnectionToken(ctx *pulumi.Context, args *LookupServiceConnectionTokenArgs, opts ...pulumi.InvokeOption) (*LookupServiceConnectionTokenResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceConnectionTokenResult
	err := ctx.Invoke("google-native:networkconnectivity/v1:getServiceConnectionToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceConnectionTokenArgs struct {
	Location                 string  `pulumi:"location"`
	Project                  *string `pulumi:"project"`
	ServiceConnectionTokenId string  `pulumi:"serviceConnectionTokenId"`
}

type LookupServiceConnectionTokenResult struct {
	// Time when the ServiceConnectionToken was created.
	CreateTime string `pulumi:"createTime"`
	// A description of this resource.
	Description string `pulumi:"description"`
	// The time to which this token is valid.
	ExpireTime string `pulumi:"expireTime"`
	// User-defined labels.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of a ServiceConnectionToken. Format: projects/{project}/locations/{location}/ServiceConnectionTokens/{service_connection_token} See: https://google.aip.dev/122#fields-representing-resource-names
	Name string `pulumi:"name"`
	// The resource path of the network associated with this token. Example: projects/{projectNumOrId}/global/networks/{resourceId}.
	Network string `pulumi:"network"`
	// The token generated by Automation.
	Token string `pulumi:"token"`
	// Time when the ServiceConnectionToken was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupServiceConnectionTokenOutput(ctx *pulumi.Context, args LookupServiceConnectionTokenOutputArgs, opts ...pulumi.InvokeOption) LookupServiceConnectionTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceConnectionTokenResult, error) {
			args := v.(LookupServiceConnectionTokenArgs)
			r, err := LookupServiceConnectionToken(ctx, &args, opts...)
			var s LookupServiceConnectionTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceConnectionTokenResultOutput)
}

type LookupServiceConnectionTokenOutputArgs struct {
	Location                 pulumi.StringInput    `pulumi:"location"`
	Project                  pulumi.StringPtrInput `pulumi:"project"`
	ServiceConnectionTokenId pulumi.StringInput    `pulumi:"serviceConnectionTokenId"`
}

func (LookupServiceConnectionTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceConnectionTokenArgs)(nil)).Elem()
}

type LookupServiceConnectionTokenResultOutput struct{ *pulumi.OutputState }

func (LookupServiceConnectionTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceConnectionTokenResult)(nil)).Elem()
}

func (o LookupServiceConnectionTokenResultOutput) ToLookupServiceConnectionTokenResultOutput() LookupServiceConnectionTokenResultOutput {
	return o
}

func (o LookupServiceConnectionTokenResultOutput) ToLookupServiceConnectionTokenResultOutputWithContext(ctx context.Context) LookupServiceConnectionTokenResultOutput {
	return o
}

// Time when the ServiceConnectionToken was created.
func (o LookupServiceConnectionTokenResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConnectionTokenResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A description of this resource.
func (o LookupServiceConnectionTokenResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConnectionTokenResult) string { return v.Description }).(pulumi.StringOutput)
}

// The time to which this token is valid.
func (o LookupServiceConnectionTokenResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConnectionTokenResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// User-defined labels.
func (o LookupServiceConnectionTokenResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceConnectionTokenResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Immutable. The name of a ServiceConnectionToken. Format: projects/{project}/locations/{location}/ServiceConnectionTokens/{service_connection_token} See: https://google.aip.dev/122#fields-representing-resource-names
func (o LookupServiceConnectionTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConnectionTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource path of the network associated with this token. Example: projects/{projectNumOrId}/global/networks/{resourceId}.
func (o LookupServiceConnectionTokenResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConnectionTokenResult) string { return v.Network }).(pulumi.StringOutput)
}

// The token generated by Automation.
func (o LookupServiceConnectionTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConnectionTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

// Time when the ServiceConnectionToken was updated.
func (o LookupServiceConnectionTokenResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceConnectionTokenResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceConnectionTokenResultOutput{})
}
