// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a WorkforcePoolProviderKey.
func LookupWorkforcePoolKey(ctx *pulumi.Context, args *LookupWorkforcePoolKeyArgs, opts ...pulumi.InvokeOption) (*LookupWorkforcePoolKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkforcePoolKeyResult
	err := ctx.Invoke("google-native:iam/v1:getWorkforcePoolKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkforcePoolKeyArgs struct {
	KeyId           string `pulumi:"keyId"`
	Location        string `pulumi:"location"`
	ProviderId      string `pulumi:"providerId"`
	WorkforcePoolId string `pulumi:"workforcePoolId"`
}

type LookupWorkforcePoolKeyResult struct {
	// The time after which the key will be permanently deleted and cannot be recovered. Note that the key may get purged before this time if the total limit of keys per provider is exceeded.
	ExpireTime string `pulumi:"expireTime"`
	// Immutable. Public half of the asymmetric key.
	KeyData KeyDataResponse `pulumi:"keyData"`
	// The resource name of the key.
	Name string `pulumi:"name"`
	// The state of the key.
	State string `pulumi:"state"`
	// The purpose of the key.
	Use string `pulumi:"use"`
}

func LookupWorkforcePoolKeyOutput(ctx *pulumi.Context, args LookupWorkforcePoolKeyOutputArgs, opts ...pulumi.InvokeOption) LookupWorkforcePoolKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkforcePoolKeyResult, error) {
			args := v.(LookupWorkforcePoolKeyArgs)
			r, err := LookupWorkforcePoolKey(ctx, &args, opts...)
			var s LookupWorkforcePoolKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkforcePoolKeyResultOutput)
}

type LookupWorkforcePoolKeyOutputArgs struct {
	KeyId           pulumi.StringInput `pulumi:"keyId"`
	Location        pulumi.StringInput `pulumi:"location"`
	ProviderId      pulumi.StringInput `pulumi:"providerId"`
	WorkforcePoolId pulumi.StringInput `pulumi:"workforcePoolId"`
}

func (LookupWorkforcePoolKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkforcePoolKeyArgs)(nil)).Elem()
}

type LookupWorkforcePoolKeyResultOutput struct{ *pulumi.OutputState }

func (LookupWorkforcePoolKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkforcePoolKeyResult)(nil)).Elem()
}

func (o LookupWorkforcePoolKeyResultOutput) ToLookupWorkforcePoolKeyResultOutput() LookupWorkforcePoolKeyResultOutput {
	return o
}

func (o LookupWorkforcePoolKeyResultOutput) ToLookupWorkforcePoolKeyResultOutputWithContext(ctx context.Context) LookupWorkforcePoolKeyResultOutput {
	return o
}

// The time after which the key will be permanently deleted and cannot be recovered. Note that the key may get purged before this time if the total limit of keys per provider is exceeded.
func (o LookupWorkforcePoolKeyResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkforcePoolKeyResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

// Immutable. Public half of the asymmetric key.
func (o LookupWorkforcePoolKeyResultOutput) KeyData() KeyDataResponseOutput {
	return o.ApplyT(func(v LookupWorkforcePoolKeyResult) KeyDataResponse { return v.KeyData }).(KeyDataResponseOutput)
}

// The resource name of the key.
func (o LookupWorkforcePoolKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkforcePoolKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the key.
func (o LookupWorkforcePoolKeyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkforcePoolKeyResult) string { return v.State }).(pulumi.StringOutput)
}

// The purpose of the key.
func (o LookupWorkforcePoolKeyResultOutput) Use() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkforcePoolKeyResult) string { return v.Use }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkforcePoolKeyResultOutput{})
}
