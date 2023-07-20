// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single GatewaySecurityPolicy.
func LookupGatewaySecurityPolicy(ctx *pulumi.Context, args *LookupGatewaySecurityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupGatewaySecurityPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewaySecurityPolicyResult
	err := ctx.Invoke("google-native:networksecurity/v1beta1:getGatewaySecurityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewaySecurityPolicyArgs struct {
	GatewaySecurityPolicyId string  `pulumi:"gatewaySecurityPolicyId"`
	Location                string  `pulumi:"location"`
	Project                 *string `pulumi:"project"`
}

type LookupGatewaySecurityPolicyResult struct {
	// The timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. Free-text description of the resource.
	Description string `pulumi:"description"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy} gateway_security_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name string `pulumi:"name"`
	// Optional. Name of a TLS Inspection Policy resource that defines how TLS inspection will be performed for any rule(s) which enables it.
	TlsInspectionPolicy string `pulumi:"tlsInspectionPolicy"`
	// The timestamp when the resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupGatewaySecurityPolicyOutput(ctx *pulumi.Context, args LookupGatewaySecurityPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupGatewaySecurityPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewaySecurityPolicyResult, error) {
			args := v.(LookupGatewaySecurityPolicyArgs)
			r, err := LookupGatewaySecurityPolicy(ctx, &args, opts...)
			var s LookupGatewaySecurityPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewaySecurityPolicyResultOutput)
}

type LookupGatewaySecurityPolicyOutputArgs struct {
	GatewaySecurityPolicyId pulumi.StringInput    `pulumi:"gatewaySecurityPolicyId"`
	Location                pulumi.StringInput    `pulumi:"location"`
	Project                 pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGatewaySecurityPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewaySecurityPolicyArgs)(nil)).Elem()
}

type LookupGatewaySecurityPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupGatewaySecurityPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewaySecurityPolicyResult)(nil)).Elem()
}

func (o LookupGatewaySecurityPolicyResultOutput) ToLookupGatewaySecurityPolicyResultOutput() LookupGatewaySecurityPolicyResultOutput {
	return o
}

func (o LookupGatewaySecurityPolicyResultOutput) ToLookupGatewaySecurityPolicyResultOutputWithContext(ctx context.Context) LookupGatewaySecurityPolicyResultOutput {
	return o
}

// The timestamp when the resource was created.
func (o LookupGatewaySecurityPolicyResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewaySecurityPolicyResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Free-text description of the resource.
func (o LookupGatewaySecurityPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewaySecurityPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy} gateway_security_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
func (o LookupGatewaySecurityPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewaySecurityPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Name of a TLS Inspection Policy resource that defines how TLS inspection will be performed for any rule(s) which enables it.
func (o LookupGatewaySecurityPolicyResultOutput) TlsInspectionPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewaySecurityPolicyResult) string { return v.TlsInspectionPolicy }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o LookupGatewaySecurityPolicyResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewaySecurityPolicyResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewaySecurityPolicyResultOutput{})
}
