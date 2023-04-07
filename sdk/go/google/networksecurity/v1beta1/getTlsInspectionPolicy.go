// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single TlsInspectionPolicy.
func LookupTlsInspectionPolicy(ctx *pulumi.Context, args *LookupTlsInspectionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupTlsInspectionPolicyResult, error) {
	var rv LookupTlsInspectionPolicyResult
	err := ctx.Invoke("google-native:networksecurity/v1beta1:getTlsInspectionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTlsInspectionPolicyArgs struct {
	Location              string  `pulumi:"location"`
	Project               *string `pulumi:"project"`
	TlsInspectionPolicyId string  `pulumi:"tlsInspectionPolicyId"`
}

type LookupTlsInspectionPolicyResult struct {
	// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
	CaPool string `pulumi:"caPool"`
	// The timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. Free-text description of the resource.
	Description string `pulumi:"description"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name string `pulumi:"name"`
	// The timestamp when the resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupTlsInspectionPolicyOutput(ctx *pulumi.Context, args LookupTlsInspectionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupTlsInspectionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTlsInspectionPolicyResult, error) {
			args := v.(LookupTlsInspectionPolicyArgs)
			r, err := LookupTlsInspectionPolicy(ctx, &args, opts...)
			var s LookupTlsInspectionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTlsInspectionPolicyResultOutput)
}

type LookupTlsInspectionPolicyOutputArgs struct {
	Location              pulumi.StringInput    `pulumi:"location"`
	Project               pulumi.StringPtrInput `pulumi:"project"`
	TlsInspectionPolicyId pulumi.StringInput    `pulumi:"tlsInspectionPolicyId"`
}

func (LookupTlsInspectionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTlsInspectionPolicyArgs)(nil)).Elem()
}

type LookupTlsInspectionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupTlsInspectionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTlsInspectionPolicyResult)(nil)).Elem()
}

func (o LookupTlsInspectionPolicyResultOutput) ToLookupTlsInspectionPolicyResultOutput() LookupTlsInspectionPolicyResultOutput {
	return o
}

func (o LookupTlsInspectionPolicyResultOutput) ToLookupTlsInspectionPolicyResultOutputWithContext(ctx context.Context) LookupTlsInspectionPolicyResultOutput {
	return o
}

// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
func (o LookupTlsInspectionPolicyResultOutput) CaPool() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTlsInspectionPolicyResult) string { return v.CaPool }).(pulumi.StringOutput)
}

// The timestamp when the resource was created.
func (o LookupTlsInspectionPolicyResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTlsInspectionPolicyResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Free-text description of the resource.
func (o LookupTlsInspectionPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTlsInspectionPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
func (o LookupTlsInspectionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTlsInspectionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o LookupTlsInspectionPolicyResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTlsInspectionPolicyResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTlsInspectionPolicyResultOutput{})
}
