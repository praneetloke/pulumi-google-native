// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a specified artifact.
func LookupDeploymentArtifact(ctx *pulumi.Context, args *LookupDeploymentArtifactArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentArtifactResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeploymentArtifactResult
	err := ctx.Invoke("google-native:apigeeregistry/v1:getDeploymentArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentArtifactArgs struct {
	ApiId        string  `pulumi:"apiId"`
	ArtifactId   string  `pulumi:"artifactId"`
	DeploymentId string  `pulumi:"deploymentId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupDeploymentArtifactResult struct {
	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations map[string]string `pulumi:"annotations"`
	// Input only. The contents of the artifact. Provided by API callers when artifacts are created or replaced. To access the contents of an artifact, use GetArtifactContents.
	Contents string `pulumi:"contents"`
	// Creation timestamp.
	CreateTime string `pulumi:"createTime"`
	// A SHA-256 hash of the artifact's contents. If the artifact is gzipped, this is the hash of the uncompressed artifact.
	Hash string `pulumi:"hash"`
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "registry.googleapis.com/" and cannot be changed.
	Labels map[string]string `pulumi:"labels"`
	// A content type specifier for the artifact. Content type specifiers are Media Types (https://en.wikipedia.org/wiki/Media_type) with a possible "schema" parameter that specifies a schema for the stored information. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
	MimeType string `pulumi:"mimeType"`
	// Resource name.
	Name string `pulumi:"name"`
	// The size of the artifact in bytes. If the artifact is gzipped, this is the size of the uncompressed artifact.
	SizeBytes int `pulumi:"sizeBytes"`
	// Last update timestamp.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupDeploymentArtifactOutput(ctx *pulumi.Context, args LookupDeploymentArtifactOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentArtifactResult, error) {
			args := v.(LookupDeploymentArtifactArgs)
			r, err := LookupDeploymentArtifact(ctx, &args, opts...)
			var s LookupDeploymentArtifactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentArtifactResultOutput)
}

type LookupDeploymentArtifactOutputArgs struct {
	ApiId        pulumi.StringInput    `pulumi:"apiId"`
	ArtifactId   pulumi.StringInput    `pulumi:"artifactId"`
	DeploymentId pulumi.StringInput    `pulumi:"deploymentId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDeploymentArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentArtifactArgs)(nil)).Elem()
}

type LookupDeploymentArtifactResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentArtifactResult)(nil)).Elem()
}

func (o LookupDeploymentArtifactResultOutput) ToLookupDeploymentArtifactResultOutput() LookupDeploymentArtifactResultOutput {
	return o
}

func (o LookupDeploymentArtifactResultOutput) ToLookupDeploymentArtifactResultOutputWithContext(ctx context.Context) LookupDeploymentArtifactResultOutput {
	return o
}

// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
func (o LookupDeploymentArtifactResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Input only. The contents of the artifact. Provided by API callers when artifacts are created or replaced. To access the contents of an artifact, use GetArtifactContents.
func (o LookupDeploymentArtifactResultOutput) Contents() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) string { return v.Contents }).(pulumi.StringOutput)
}

// Creation timestamp.
func (o LookupDeploymentArtifactResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A SHA-256 hash of the artifact's contents. If the artifact is gzipped, this is the hash of the uncompressed artifact.
func (o LookupDeploymentArtifactResultOutput) Hash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) string { return v.Hash }).(pulumi.StringOutput)
}

// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "registry.googleapis.com/" and cannot be changed.
func (o LookupDeploymentArtifactResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// A content type specifier for the artifact. Content type specifiers are Media Types (https://en.wikipedia.org/wiki/Media_type) with a possible "schema" parameter that specifies a schema for the stored information. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
func (o LookupDeploymentArtifactResultOutput) MimeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) string { return v.MimeType }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupDeploymentArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

// The size of the artifact in bytes. If the artifact is gzipped, this is the size of the uncompressed artifact.
func (o LookupDeploymentArtifactResultOutput) SizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) int { return v.SizeBytes }).(pulumi.IntOutput)
}

// Last update timestamp.
func (o LookupDeploymentArtifactResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentArtifactResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentArtifactResultOutput{})
}
