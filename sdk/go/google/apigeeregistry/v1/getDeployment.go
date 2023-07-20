// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a specified deployment.
func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeploymentResult
	err := ctx.Invoke("google-native:apigeeregistry/v1:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentArgs struct {
	ApiId        string  `pulumi:"apiId"`
	DeploymentId string  `pulumi:"deploymentId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupDeploymentResult struct {
	// Text briefly describing how to access the endpoint. Changes to this value will not affect the revision.
	AccessGuidance string `pulumi:"accessGuidance"`
	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations map[string]string `pulumi:"annotations"`
	// The full resource name (including revision ID) of the spec of the API being served by the deployment. Changes to this value will update the revision. Format: `projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec@revision}`
	ApiSpecRevision string `pulumi:"apiSpecRevision"`
	// Creation timestamp; when the deployment resource was created.
	CreateTime string `pulumi:"createTime"`
	// A detailed description.
	Description string `pulumi:"description"`
	// Human-meaningful name.
	DisplayName string `pulumi:"displayName"`
	// The address where the deployment is serving. Changes to this value will update the revision.
	EndpointUri string `pulumi:"endpointUri"`
	// The address of the external channel of the API (e.g., the Developer Portal). Changes to this value will not affect the revision.
	ExternalChannelUri string `pulumi:"externalChannelUri"`
	// Text briefly identifying the intended audience of the API. Changes to this value will not affect the revision.
	IntendedAudience string `pulumi:"intendedAudience"`
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
	Labels map[string]string `pulumi:"labels"`
	// Resource name.
	Name string `pulumi:"name"`
	// Revision creation timestamp; when the represented revision was created.
	RevisionCreateTime string `pulumi:"revisionCreateTime"`
	// Immutable. The revision ID of the deployment. A new revision is committed whenever the deployment contents are changed. The format is an 8-character hexadecimal string.
	RevisionId string `pulumi:"revisionId"`
	// Last update timestamp: when the represented revision was last modified.
	RevisionUpdateTime string `pulumi:"revisionUpdateTime"`
}

func LookupDeploymentOutput(ctx *pulumi.Context, args LookupDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentResult, error) {
			args := v.(LookupDeploymentArgs)
			r, err := LookupDeployment(ctx, &args, opts...)
			var s LookupDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentResultOutput)
}

type LookupDeploymentOutputArgs struct {
	ApiId        pulumi.StringInput    `pulumi:"apiId"`
	DeploymentId pulumi.StringInput    `pulumi:"deploymentId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentArgs)(nil)).Elem()
}

type LookupDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentResult)(nil)).Elem()
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutput() LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutputWithContext(ctx context.Context) LookupDeploymentResultOutput {
	return o
}

// Text briefly describing how to access the endpoint. Changes to this value will not affect the revision.
func (o LookupDeploymentResultOutput) AccessGuidance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.AccessGuidance }).(pulumi.StringOutput)
}

// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
func (o LookupDeploymentResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// The full resource name (including revision ID) of the spec of the API being served by the deployment. Changes to this value will update the revision. Format: `projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec@revision}`
func (o LookupDeploymentResultOutput) ApiSpecRevision() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.ApiSpecRevision }).(pulumi.StringOutput)
}

// Creation timestamp; when the deployment resource was created.
func (o LookupDeploymentResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A detailed description.
func (o LookupDeploymentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Description }).(pulumi.StringOutput)
}

// Human-meaningful name.
func (o LookupDeploymentResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The address where the deployment is serving. Changes to this value will update the revision.
func (o LookupDeploymentResultOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.EndpointUri }).(pulumi.StringOutput)
}

// The address of the external channel of the API (e.g., the Developer Portal). Changes to this value will not affect the revision.
func (o LookupDeploymentResultOutput) ExternalChannelUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.ExternalChannelUri }).(pulumi.StringOutput)
}

// Text briefly identifying the intended audience of the API. Changes to this value will not affect the revision.
func (o LookupDeploymentResultOutput) IntendedAudience() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.IntendedAudience }).(pulumi.StringOutput)
}

// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
func (o LookupDeploymentResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource name.
func (o LookupDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Revision creation timestamp; when the represented revision was created.
func (o LookupDeploymentResultOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// Immutable. The revision ID of the deployment. A new revision is committed whenever the deployment contents are changed. The format is an 8-character hexadecimal string.
func (o LookupDeploymentResultOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.RevisionId }).(pulumi.StringOutput)
}

// Last update timestamp: when the represented revision was last modified.
func (o LookupDeploymentResultOutput) RevisionUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.RevisionUpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentResultOutput{})
}
