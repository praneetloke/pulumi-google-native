// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified instance template.
func LookupRegionInstanceTemplate(ctx *pulumi.Context, args *LookupRegionInstanceTemplateArgs, opts ...pulumi.InvokeOption) (*LookupRegionInstanceTemplateResult, error) {
	var rv LookupRegionInstanceTemplateResult
	err := ctx.Invoke("google-native:compute/alpha:getRegionInstanceTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionInstanceTemplateArgs struct {
	InstanceTemplate string  `pulumi:"instanceTemplate"`
	Project          *string `pulumi:"project"`
	Region           string  `pulumi:"region"`
}

type LookupRegionInstanceTemplateResult struct {
	// The creation timestamp for this instance template in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// The resource type, which is always compute#instanceTemplate for instance templates.
	Kind string `pulumi:"kind"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The instance properties for this instance template.
	Properties InstancePropertiesResponse `pulumi:"properties"`
	// URL of the region where the instance template resides. Only applicable for regional resources.
	Region string `pulumi:"region"`
	// The URL for this instance template. The server defines this URL.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
	SourceInstance string `pulumi:"sourceInstance"`
	// The source instance params to use to create this instance template.
	SourceInstanceParams SourceInstanceParamsResponse `pulumi:"sourceInstanceParams"`
}

func LookupRegionInstanceTemplateOutput(ctx *pulumi.Context, args LookupRegionInstanceTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupRegionInstanceTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionInstanceTemplateResult, error) {
			args := v.(LookupRegionInstanceTemplateArgs)
			r, err := LookupRegionInstanceTemplate(ctx, &args, opts...)
			var s LookupRegionInstanceTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegionInstanceTemplateResultOutput)
}

type LookupRegionInstanceTemplateOutputArgs struct {
	InstanceTemplate pulumi.StringInput    `pulumi:"instanceTemplate"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
	Region           pulumi.StringInput    `pulumi:"region"`
}

func (LookupRegionInstanceTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionInstanceTemplateArgs)(nil)).Elem()
}

type LookupRegionInstanceTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupRegionInstanceTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionInstanceTemplateResult)(nil)).Elem()
}

func (o LookupRegionInstanceTemplateResultOutput) ToLookupRegionInstanceTemplateResultOutput() LookupRegionInstanceTemplateResultOutput {
	return o
}

func (o LookupRegionInstanceTemplateResultOutput) ToLookupRegionInstanceTemplateResultOutputWithContext(ctx context.Context) LookupRegionInstanceTemplateResultOutput {
	return o
}

// The creation timestamp for this instance template in RFC3339 text format.
func (o LookupRegionInstanceTemplateResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRegionInstanceTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

// The resource type, which is always compute#instanceTemplate for instance templates.
func (o LookupRegionInstanceTemplateResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionInstanceTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// The instance properties for this instance template.
func (o LookupRegionInstanceTemplateResultOutput) Properties() InstancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) InstancePropertiesResponse { return v.Properties }).(InstancePropertiesResponseOutput)
}

// URL of the region where the instance template resides. Only applicable for regional resources.
func (o LookupRegionInstanceTemplateResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) string { return v.Region }).(pulumi.StringOutput)
}

// The URL for this instance template. The server defines this URL.
func (o LookupRegionInstanceTemplateResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupRegionInstanceTemplateResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
func (o LookupRegionInstanceTemplateResultOutput) SourceInstance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) string { return v.SourceInstance }).(pulumi.StringOutput)
}

// The source instance params to use to create this instance template.
func (o LookupRegionInstanceTemplateResultOutput) SourceInstanceParams() SourceInstanceParamsResponseOutput {
	return o.ApplyT(func(v LookupRegionInstanceTemplateResult) SourceInstanceParamsResponse { return v.SourceInstanceParams }).(SourceInstanceParamsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionInstanceTemplateResultOutput{})
}
