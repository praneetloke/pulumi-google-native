// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an instance template in the specified project and region using the global instance template whose URL is included in the request.
type RegionInstanceTemplate struct {
	pulumi.CustomResourceState

	// The creation timestamp for this instance template in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource type, which is always compute#instanceTemplate for instance templates.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The instance properties for this instance template.
	Properties InstancePropertiesResponseOutput `pulumi:"properties"`
	Region     pulumi.StringOutput              `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The URL for this instance template. The server defines this URL.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
	SourceInstance pulumi.StringOutput `pulumi:"sourceInstance"`
	// The source instance params to use to create this instance template.
	SourceInstanceParams SourceInstanceParamsResponseOutput `pulumi:"sourceInstanceParams"`
}

// NewRegionInstanceTemplate registers a new resource with the given unique name, arguments, and options.
func NewRegionInstanceTemplate(ctx *pulumi.Context,
	name string, args *RegionInstanceTemplateArgs, opts ...pulumi.ResourceOption) (*RegionInstanceTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"region",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegionInstanceTemplate
	err := ctx.RegisterResource("google-native:compute/v1:RegionInstanceTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionInstanceTemplate gets an existing RegionInstanceTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionInstanceTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionInstanceTemplateState, opts ...pulumi.ResourceOption) (*RegionInstanceTemplate, error) {
	var resource RegionInstanceTemplate
	err := ctx.ReadResource("google-native:compute/v1:RegionInstanceTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionInstanceTemplate resources.
type regionInstanceTemplateState struct {
}

type RegionInstanceTemplateState struct {
}

func (RegionInstanceTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionInstanceTemplateState)(nil)).Elem()
}

type regionInstanceTemplateArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The instance properties for this instance template.
	Properties *InstanceProperties `pulumi:"properties"`
	Region     string              `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
	SourceInstance *string `pulumi:"sourceInstance"`
	// The source instance params to use to create this instance template.
	SourceInstanceParams *SourceInstanceParams `pulumi:"sourceInstanceParams"`
}

// The set of arguments for constructing a RegionInstanceTemplate resource.
type RegionInstanceTemplateArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The instance properties for this instance template.
	Properties InstancePropertiesPtrInput
	Region     pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
	SourceInstance pulumi.StringPtrInput
	// The source instance params to use to create this instance template.
	SourceInstanceParams SourceInstanceParamsPtrInput
}

func (RegionInstanceTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionInstanceTemplateArgs)(nil)).Elem()
}

type RegionInstanceTemplateInput interface {
	pulumi.Input

	ToRegionInstanceTemplateOutput() RegionInstanceTemplateOutput
	ToRegionInstanceTemplateOutputWithContext(ctx context.Context) RegionInstanceTemplateOutput
}

func (*RegionInstanceTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionInstanceTemplate)(nil)).Elem()
}

func (i *RegionInstanceTemplate) ToRegionInstanceTemplateOutput() RegionInstanceTemplateOutput {
	return i.ToRegionInstanceTemplateOutputWithContext(context.Background())
}

func (i *RegionInstanceTemplate) ToRegionInstanceTemplateOutputWithContext(ctx context.Context) RegionInstanceTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceTemplateOutput)
}

type RegionInstanceTemplateOutput struct{ *pulumi.OutputState }

func (RegionInstanceTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionInstanceTemplate)(nil)).Elem()
}

func (o RegionInstanceTemplateOutput) ToRegionInstanceTemplateOutput() RegionInstanceTemplateOutput {
	return o
}

func (o RegionInstanceTemplateOutput) ToRegionInstanceTemplateOutputWithContext(ctx context.Context) RegionInstanceTemplateOutput {
	return o
}

// The creation timestamp for this instance template in RFC3339 text format.
func (o RegionInstanceTemplateOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o RegionInstanceTemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The resource type, which is always compute#instanceTemplate for instance templates.
func (o RegionInstanceTemplateOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o RegionInstanceTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegionInstanceTemplateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The instance properties for this instance template.
func (o RegionInstanceTemplateOutput) Properties() InstancePropertiesResponseOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) InstancePropertiesResponseOutput { return v.Properties }).(InstancePropertiesResponseOutput)
}

func (o RegionInstanceTemplateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o RegionInstanceTemplateOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The URL for this instance template. The server defines this URL.
func (o RegionInstanceTemplateOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
func (o RegionInstanceTemplateOutput) SourceInstance() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) pulumi.StringOutput { return v.SourceInstance }).(pulumi.StringOutput)
}

// The source instance params to use to create this instance template.
func (o RegionInstanceTemplateOutput) SourceInstanceParams() SourceInstanceParamsResponseOutput {
	return o.ApplyT(func(v *RegionInstanceTemplate) SourceInstanceParamsResponseOutput { return v.SourceInstanceParams }).(SourceInstanceParamsResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionInstanceTemplateInput)(nil)).Elem(), &RegionInstanceTemplate{})
	pulumi.RegisterOutputType(RegionInstanceTemplateOutput{})
}
