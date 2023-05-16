// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single EndpointAttachment.
func LookupEndpointAttachment(ctx *pulumi.Context, args *LookupEndpointAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupEndpointAttachmentResult, error) {
	var rv LookupEndpointAttachmentResult
	err := ctx.Invoke("google-native:connectors/v1:getEndpointAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointAttachmentArgs struct {
	EndpointAttachmentId string  `pulumi:"endpointAttachmentId"`
	Location             string  `pulumi:"location"`
	Project              *string `pulumi:"project"`
}

type LookupEndpointAttachmentResult struct {
	// Created time.
	CreateTime string `pulumi:"createTime"`
	// Optional. Description of the resource.
	Description string `pulumi:"description"`
	// The Private Service Connect connection endpoint ip
	EndpointIp string `pulumi:"endpointIp"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels map[string]string `pulumi:"labels"`
	// Resource name of the Endpoint Attachment. Format: projects/{project}/locations/{location}/endpointAttachments/{endpoint_attachment}
	Name string `pulumi:"name"`
	// The path of the service attachment
	ServiceAttachment string `pulumi:"serviceAttachment"`
	// Updated time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupEndpointAttachmentOutput(ctx *pulumi.Context, args LookupEndpointAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointAttachmentResult, error) {
			args := v.(LookupEndpointAttachmentArgs)
			r, err := LookupEndpointAttachment(ctx, &args, opts...)
			var s LookupEndpointAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointAttachmentResultOutput)
}

type LookupEndpointAttachmentOutputArgs struct {
	EndpointAttachmentId pulumi.StringInput    `pulumi:"endpointAttachmentId"`
	Location             pulumi.StringInput    `pulumi:"location"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupEndpointAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointAttachmentArgs)(nil)).Elem()
}

type LookupEndpointAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointAttachmentResult)(nil)).Elem()
}

func (o LookupEndpointAttachmentResultOutput) ToLookupEndpointAttachmentResultOutput() LookupEndpointAttachmentResultOutput {
	return o
}

func (o LookupEndpointAttachmentResultOutput) ToLookupEndpointAttachmentResultOutputWithContext(ctx context.Context) LookupEndpointAttachmentResultOutput {
	return o
}

// Created time.
func (o LookupEndpointAttachmentResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointAttachmentResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the resource.
func (o LookupEndpointAttachmentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointAttachmentResult) string { return v.Description }).(pulumi.StringOutput)
}

// The Private Service Connect connection endpoint ip
func (o LookupEndpointAttachmentResultOutput) EndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointAttachmentResult) string { return v.EndpointIp }).(pulumi.StringOutput)
}

// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
func (o LookupEndpointAttachmentResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEndpointAttachmentResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource name of the Endpoint Attachment. Format: projects/{project}/locations/{location}/endpointAttachments/{endpoint_attachment}
func (o LookupEndpointAttachmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointAttachmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The path of the service attachment
func (o LookupEndpointAttachmentResultOutput) ServiceAttachment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointAttachmentResult) string { return v.ServiceAttachment }).(pulumi.StringOutput)
}

// Updated time.
func (o LookupEndpointAttachmentResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointAttachmentResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointAttachmentResultOutput{})
}
