// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the key value entry value for a key value map scoped to an organization, environment, or API proxy. **Note**: Supported for Apigee hybrid 1.8.x and higher.
func LookupEntry(ctx *pulumi.Context, args *LookupEntryArgs, opts ...pulumi.InvokeOption) (*LookupEntryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEntryResult
	err := ctx.Invoke("google-native:apigee/v1:getEntry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntryArgs struct {
	ApiId          string `pulumi:"apiId"`
	EntryId        string `pulumi:"entryId"`
	KeyvaluemapId  string `pulumi:"keyvaluemapId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupEntryResult struct {
	// Resource URI that can be used to identify the scope of the key value map entries.
	Name string `pulumi:"name"`
	// Data or payload that is being retrieved and associated with the unique key.
	Value string `pulumi:"value"`
}

func LookupEntryOutput(ctx *pulumi.Context, args LookupEntryOutputArgs, opts ...pulumi.InvokeOption) LookupEntryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEntryResult, error) {
			args := v.(LookupEntryArgs)
			r, err := LookupEntry(ctx, &args, opts...)
			var s LookupEntryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEntryResultOutput)
}

type LookupEntryOutputArgs struct {
	ApiId          pulumi.StringInput `pulumi:"apiId"`
	EntryId        pulumi.StringInput `pulumi:"entryId"`
	KeyvaluemapId  pulumi.StringInput `pulumi:"keyvaluemapId"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupEntryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntryArgs)(nil)).Elem()
}

type LookupEntryResultOutput struct{ *pulumi.OutputState }

func (LookupEntryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntryResult)(nil)).Elem()
}

func (o LookupEntryResultOutput) ToLookupEntryResultOutput() LookupEntryResultOutput {
	return o
}

func (o LookupEntryResultOutput) ToLookupEntryResultOutputWithContext(ctx context.Context) LookupEntryResultOutput {
	return o
}

// Resource URI that can be used to identify the scope of the key value map entries.
func (o LookupEntryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Data or payload that is being retrieved and associated with the unique key.
func (o LookupEntryResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEntryResultOutput{})
}
