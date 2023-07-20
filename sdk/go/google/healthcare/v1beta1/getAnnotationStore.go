// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Annotation store or returns NOT_FOUND if it does not exist.
func LookupAnnotationStore(ctx *pulumi.Context, args *LookupAnnotationStoreArgs, opts ...pulumi.InvokeOption) (*LookupAnnotationStoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAnnotationStoreResult
	err := ctx.Invoke("google-native:healthcare/v1beta1:getAnnotationStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnnotationStoreArgs struct {
	AnnotationStoreId string  `pulumi:"annotationStoreId"`
	DatasetId         string  `pulumi:"datasetId"`
	Location          string  `pulumi:"location"`
	Project           *string `pulumi:"project"`
}

type LookupAnnotationStoreResult struct {
	// Optional. User-supplied key-value pairs used to organize Annotation stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels map[string]string `pulumi:"labels"`
	// Resource name of the Annotation store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}`.
	Name string `pulumi:"name"`
}

func LookupAnnotationStoreOutput(ctx *pulumi.Context, args LookupAnnotationStoreOutputArgs, opts ...pulumi.InvokeOption) LookupAnnotationStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnnotationStoreResult, error) {
			args := v.(LookupAnnotationStoreArgs)
			r, err := LookupAnnotationStore(ctx, &args, opts...)
			var s LookupAnnotationStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnnotationStoreResultOutput)
}

type LookupAnnotationStoreOutputArgs struct {
	AnnotationStoreId pulumi.StringInput    `pulumi:"annotationStoreId"`
	DatasetId         pulumi.StringInput    `pulumi:"datasetId"`
	Location          pulumi.StringInput    `pulumi:"location"`
	Project           pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupAnnotationStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnnotationStoreArgs)(nil)).Elem()
}

type LookupAnnotationStoreResultOutput struct{ *pulumi.OutputState }

func (LookupAnnotationStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnnotationStoreResult)(nil)).Elem()
}

func (o LookupAnnotationStoreResultOutput) ToLookupAnnotationStoreResultOutput() LookupAnnotationStoreResultOutput {
	return o
}

func (o LookupAnnotationStoreResultOutput) ToLookupAnnotationStoreResultOutputWithContext(ctx context.Context) LookupAnnotationStoreResultOutput {
	return o
}

// Optional. User-supplied key-value pairs used to organize Annotation stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
func (o LookupAnnotationStoreResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource name of the Annotation store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}`.
func (o LookupAnnotationStoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnnotationStoreResultOutput{})
}
