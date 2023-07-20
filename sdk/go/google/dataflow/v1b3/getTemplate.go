// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1b3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the template associated with a template.
func LookupTemplate(ctx *pulumi.Context, args *LookupTemplateArgs, opts ...pulumi.InvokeOption) (*LookupTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTemplateResult
	err := ctx.Invoke("google-native:dataflow/v1b3:getTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateArgs struct {
	GcsPath  string  `pulumi:"gcsPath"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	View     *string `pulumi:"view"`
}

type LookupTemplateResult struct {
	// The template metadata describing the template name, available parameters, etc.
	Metadata TemplateMetadataResponse `pulumi:"metadata"`
	// Describes the runtime metadata with SDKInfo and available parameters.
	RuntimeMetadata RuntimeMetadataResponse `pulumi:"runtimeMetadata"`
	// The status of the get template request. Any problems with the request will be indicated in the error_details.
	Status StatusResponse `pulumi:"status"`
	// Template Type.
	TemplateType string `pulumi:"templateType"`
}

func LookupTemplateOutput(ctx *pulumi.Context, args LookupTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplateResult, error) {
			args := v.(LookupTemplateArgs)
			r, err := LookupTemplate(ctx, &args, opts...)
			var s LookupTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplateResultOutput)
}

type LookupTemplateOutputArgs struct {
	GcsPath  pulumi.StringInput    `pulumi:"gcsPath"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	View     pulumi.StringPtrInput `pulumi:"view"`
}

func (LookupTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateArgs)(nil)).Elem()
}

type LookupTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateResult)(nil)).Elem()
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutput() LookupTemplateResultOutput {
	return o
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutputWithContext(ctx context.Context) LookupTemplateResultOutput {
	return o
}

// The template metadata describing the template name, available parameters, etc.
func (o LookupTemplateResultOutput) Metadata() TemplateMetadataResponseOutput {
	return o.ApplyT(func(v LookupTemplateResult) TemplateMetadataResponse { return v.Metadata }).(TemplateMetadataResponseOutput)
}

// Describes the runtime metadata with SDKInfo and available parameters.
func (o LookupTemplateResultOutput) RuntimeMetadata() RuntimeMetadataResponseOutput {
	return o.ApplyT(func(v LookupTemplateResult) RuntimeMetadataResponse { return v.RuntimeMetadata }).(RuntimeMetadataResponseOutput)
}

// The status of the get template request. Any problems with the request will be indicated in the error_details.
func (o LookupTemplateResultOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v LookupTemplateResult) StatusResponse { return v.Status }).(StatusResponseOutput)
}

// Template Type.
func (o LookupTemplateResultOutput) TemplateType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateResult) string { return v.TemplateType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateResultOutput{})
}
