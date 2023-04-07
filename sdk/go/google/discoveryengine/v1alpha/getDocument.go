// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Document.
func LookupDocument(ctx *pulumi.Context, args *LookupDocumentArgs, opts ...pulumi.InvokeOption) (*LookupDocumentResult, error) {
	var rv LookupDocumentResult
	err := ctx.Invoke("google-native:discoveryengine/v1alpha:getDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDocumentArgs struct {
	BranchId     string  `pulumi:"branchId"`
	CollectionId string  `pulumi:"collectionId"`
	DataStoreId  string  `pulumi:"dataStoreId"`
	DocumentId   string  `pulumi:"documentId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupDocumentResult struct {
	// The JSON string representation of the document. It should conform to the registered schema or an INVALID_ARGUMENT error is thrown.
	JsonData string `pulumi:"jsonData"`
	// Immutable. The full resource name of the document. Format: `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
	Name string `pulumi:"name"`
	// The identifier of the parent document. Currently supports at most two level document hierarchy. Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
	ParentDocumentId string `pulumi:"parentDocumentId"`
	// The identifier of the schema located in the same data store.
	SchemaId string `pulumi:"schemaId"`
	// The structured JSON data for the document. It should conform to the registered schema or an INVALID_ARGUMENT error is thrown.
	StructData map[string]string `pulumi:"structData"`
}

func LookupDocumentOutput(ctx *pulumi.Context, args LookupDocumentOutputArgs, opts ...pulumi.InvokeOption) LookupDocumentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDocumentResult, error) {
			args := v.(LookupDocumentArgs)
			r, err := LookupDocument(ctx, &args, opts...)
			var s LookupDocumentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDocumentResultOutput)
}

type LookupDocumentOutputArgs struct {
	BranchId     pulumi.StringInput    `pulumi:"branchId"`
	CollectionId pulumi.StringInput    `pulumi:"collectionId"`
	DataStoreId  pulumi.StringInput    `pulumi:"dataStoreId"`
	DocumentId   pulumi.StringInput    `pulumi:"documentId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDocumentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentArgs)(nil)).Elem()
}

type LookupDocumentResultOutput struct{ *pulumi.OutputState }

func (LookupDocumentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentResult)(nil)).Elem()
}

func (o LookupDocumentResultOutput) ToLookupDocumentResultOutput() LookupDocumentResultOutput {
	return o
}

func (o LookupDocumentResultOutput) ToLookupDocumentResultOutputWithContext(ctx context.Context) LookupDocumentResultOutput {
	return o
}

// The JSON string representation of the document. It should conform to the registered schema or an INVALID_ARGUMENT error is thrown.
func (o LookupDocumentResultOutput) JsonData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.JsonData }).(pulumi.StringOutput)
}

// Immutable. The full resource name of the document. Format: `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
func (o LookupDocumentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The identifier of the parent document. Currently supports at most two level document hierarchy. Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
func (o LookupDocumentResultOutput) ParentDocumentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.ParentDocumentId }).(pulumi.StringOutput)
}

// The identifier of the schema located in the same data store.
func (o LookupDocumentResultOutput) SchemaId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.SchemaId }).(pulumi.StringOutput)
}

// The structured JSON data for the document. It should conform to the registered schema or an INVALID_ARGUMENT error is thrown.
func (o LookupDocumentResultOutput) StructData() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDocumentResult) map[string]string { return v.StructData }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDocumentResultOutput{})
}
