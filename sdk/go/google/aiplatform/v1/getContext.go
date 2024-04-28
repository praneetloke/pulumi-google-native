// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a specific Context.
func LookupContext(ctx *pulumi.Context, args *LookupContextArgs, opts ...pulumi.InvokeOption) (*LookupContextResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContextResult
	err := ctx.Invoke("google-native:aiplatform/v1:getContext", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContextArgs struct {
	ContextId       string  `pulumi:"contextId"`
	Location        string  `pulumi:"location"`
	MetadataStoreId string  `pulumi:"metadataStoreId"`
	Project         *string `pulumi:"project"`
}

type LookupContextResult struct {
	// Timestamp when this Context was created.
	CreateTime string `pulumi:"createTime"`
	// Description of the Context
	Description string `pulumi:"description"`
	// User provided display name of the Context. May be up to 128 Unicode characters.
	DisplayName string `pulumi:"displayName"`
	// An eTag used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag string `pulumi:"etag"`
	// The labels with user-defined metadata to organize your Contexts. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Context (System labels are excluded).
	Labels map[string]string `pulumi:"labels"`
	// Properties of the Context. Top level metadata keys' heading and trailing spaces will be trimmed. The size of this field should not exceed 200KB.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Immutable. The resource name of the Context.
	Name string `pulumi:"name"`
	// A list of resource names of Contexts that are parents of this Context. A Context may have at most 10 parent_contexts.
	ParentContexts []string `pulumi:"parentContexts"`
	// The title of the schema describing the metadata. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
	SchemaTitle string `pulumi:"schemaTitle"`
	// The version of the schema in schema_name to use. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
	SchemaVersion string `pulumi:"schemaVersion"`
	// Timestamp when this Context was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupContextOutput(ctx *pulumi.Context, args LookupContextOutputArgs, opts ...pulumi.InvokeOption) LookupContextResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContextResult, error) {
			args := v.(LookupContextArgs)
			r, err := LookupContext(ctx, &args, opts...)
			var s LookupContextResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContextResultOutput)
}

type LookupContextOutputArgs struct {
	ContextId       pulumi.StringInput    `pulumi:"contextId"`
	Location        pulumi.StringInput    `pulumi:"location"`
	MetadataStoreId pulumi.StringInput    `pulumi:"metadataStoreId"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupContextOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContextArgs)(nil)).Elem()
}

type LookupContextResultOutput struct{ *pulumi.OutputState }

func (LookupContextResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContextResult)(nil)).Elem()
}

func (o LookupContextResultOutput) ToLookupContextResultOutput() LookupContextResultOutput {
	return o
}

func (o LookupContextResultOutput) ToLookupContextResultOutputWithContext(ctx context.Context) LookupContextResultOutput {
	return o
}

// Timestamp when this Context was created.
func (o LookupContextResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the Context
func (o LookupContextResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextResult) string { return v.Description }).(pulumi.StringOutput)
}

// User provided display name of the Context. May be up to 128 Unicode characters.
func (o LookupContextResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// An eTag used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
func (o LookupContextResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The labels with user-defined metadata to organize your Contexts. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Context (System labels are excluded).
func (o LookupContextResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContextResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Properties of the Context. Top level metadata keys' heading and trailing spaces will be trimmed. The size of this field should not exceed 200KB.
func (o LookupContextResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupContextResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// Immutable. The resource name of the Context.
func (o LookupContextResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of resource names of Contexts that are parents of this Context. A Context may have at most 10 parent_contexts.
func (o LookupContextResultOutput) ParentContexts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContextResult) []string { return v.ParentContexts }).(pulumi.StringArrayOutput)
}

// The title of the schema describing the metadata. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
func (o LookupContextResultOutput) SchemaTitle() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextResult) string { return v.SchemaTitle }).(pulumi.StringOutput)
}

// The version of the schema in schema_name to use. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
func (o LookupContextResultOutput) SchemaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextResult) string { return v.SchemaVersion }).(pulumi.StringOutput)
}

// Timestamp when this Context was last updated.
func (o LookupContextResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContextResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContextResultOutput{})
}
