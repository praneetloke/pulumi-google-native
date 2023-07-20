// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a SynonymSet for a particular context. Throws a NOT_FOUND exception if the Synonymset does not exist
func LookupSynonymSet(ctx *pulumi.Context, args *LookupSynonymSetArgs, opts ...pulumi.InvokeOption) (*LookupSynonymSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSynonymSetResult
	err := ctx.Invoke("google-native:contentwarehouse/v1:getSynonymSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSynonymSetArgs struct {
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
	SynonymSetId string  `pulumi:"synonymSetId"`
}

type LookupSynonymSetResult struct {
	// This is a freeform field. Example contexts can be "sales," "engineering," "real estate," "accounting," etc. The context can be supplied during search requests.
	Context string `pulumi:"context"`
	// The resource name of the SynonymSet This is mandatory for google.api.resource. Format: projects/{project_number}/locations/{location}/synonymSets/{context}.
	Name string `pulumi:"name"`
	// List of Synonyms for the context.
	Synonyms []GoogleCloudContentwarehouseV1SynonymSetSynonymResponse `pulumi:"synonyms"`
}

func LookupSynonymSetOutput(ctx *pulumi.Context, args LookupSynonymSetOutputArgs, opts ...pulumi.InvokeOption) LookupSynonymSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSynonymSetResult, error) {
			args := v.(LookupSynonymSetArgs)
			r, err := LookupSynonymSet(ctx, &args, opts...)
			var s LookupSynonymSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSynonymSetResultOutput)
}

type LookupSynonymSetOutputArgs struct {
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
	SynonymSetId pulumi.StringInput    `pulumi:"synonymSetId"`
}

func (LookupSynonymSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynonymSetArgs)(nil)).Elem()
}

type LookupSynonymSetResultOutput struct{ *pulumi.OutputState }

func (LookupSynonymSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynonymSetResult)(nil)).Elem()
}

func (o LookupSynonymSetResultOutput) ToLookupSynonymSetResultOutput() LookupSynonymSetResultOutput {
	return o
}

func (o LookupSynonymSetResultOutput) ToLookupSynonymSetResultOutputWithContext(ctx context.Context) LookupSynonymSetResultOutput {
	return o
}

// This is a freeform field. Example contexts can be "sales," "engineering," "real estate," "accounting," etc. The context can be supplied during search requests.
func (o LookupSynonymSetResultOutput) Context() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynonymSetResult) string { return v.Context }).(pulumi.StringOutput)
}

// The resource name of the SynonymSet This is mandatory for google.api.resource. Format: projects/{project_number}/locations/{location}/synonymSets/{context}.
func (o LookupSynonymSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynonymSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of Synonyms for the context.
func (o LookupSynonymSetResultOutput) Synonyms() GoogleCloudContentwarehouseV1SynonymSetSynonymResponseArrayOutput {
	return o.ApplyT(func(v LookupSynonymSetResult) []GoogleCloudContentwarehouseV1SynonymSetSynonymResponse {
		return v.Synonyms
	}).(GoogleCloudContentwarehouseV1SynonymSetSynonymResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSynonymSetResultOutput{})
}
