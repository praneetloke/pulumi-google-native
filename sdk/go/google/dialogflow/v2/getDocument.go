// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified document.
func LookupDocument(ctx *pulumi.Context, args *LookupDocumentArgs, opts ...pulumi.InvokeOption) (*LookupDocumentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDocumentResult
	err := ctx.Invoke("google-native:dialogflow/v2:getDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDocumentArgs struct {
	DocumentId      string  `pulumi:"documentId"`
	KnowledgeBaseId string  `pulumi:"knowledgeBaseId"`
	Location        string  `pulumi:"location"`
	Project         *string `pulumi:"project"`
}

type LookupDocumentResult struct {
	// The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
	ContentUri string `pulumi:"contentUri"`
	// The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
	DisplayName string `pulumi:"displayName"`
	// Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
	EnableAutoReload bool `pulumi:"enableAutoReload"`
	// The knowledge type of document content.
	KnowledgeTypes []string `pulumi:"knowledgeTypes"`
	// The time and status of the latest reload. This reload may have been triggered automatically or manually and may not have succeeded.
	LatestReloadStatus GoogleCloudDialogflowV2DocumentReloadStatusResponse `pulumi:"latestReloadStatus"`
	// Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
	Metadata map[string]string `pulumi:"metadata"`
	// The MIME type of this document.
	MimeType string `pulumi:"mimeType"`
	// Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
	Name string `pulumi:"name"`
	// The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
	RawContent string `pulumi:"rawContent"`
	// The current state of the document.
	State string `pulumi:"state"`
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
	DocumentId      pulumi.StringInput    `pulumi:"documentId"`
	KnowledgeBaseId pulumi.StringInput    `pulumi:"knowledgeBaseId"`
	Location        pulumi.StringInput    `pulumi:"location"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
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

// The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
func (o LookupDocumentResultOutput) ContentUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.ContentUri }).(pulumi.StringOutput)
}

// The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
func (o LookupDocumentResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
func (o LookupDocumentResultOutput) EnableAutoReload() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDocumentResult) bool { return v.EnableAutoReload }).(pulumi.BoolOutput)
}

// The knowledge type of document content.
func (o LookupDocumentResultOutput) KnowledgeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDocumentResult) []string { return v.KnowledgeTypes }).(pulumi.StringArrayOutput)
}

// The time and status of the latest reload. This reload may have been triggered automatically or manually and may not have succeeded.
func (o LookupDocumentResultOutput) LatestReloadStatus() GoogleCloudDialogflowV2DocumentReloadStatusResponseOutput {
	return o.ApplyT(func(v LookupDocumentResult) GoogleCloudDialogflowV2DocumentReloadStatusResponse {
		return v.LatestReloadStatus
	}).(GoogleCloudDialogflowV2DocumentReloadStatusResponseOutput)
}

// Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
func (o LookupDocumentResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDocumentResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The MIME type of this document.
func (o LookupDocumentResultOutput) MimeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.MimeType }).(pulumi.StringOutput)
}

// Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
func (o LookupDocumentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
func (o LookupDocumentResultOutput) RawContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.RawContent }).(pulumi.StringOutput)
}

// The current state of the document.
func (o LookupDocumentResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDocumentResultOutput{})
}
