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

// Creates a new Source in a given project and location.
// Auto-naming is currently not supported for this resource.
type Source struct {
	pulumi.CustomResourceState

	// AWS type source details.
	Aws AwsSourceDetailsResponseOutput `pulumi:"aws"`
	// The create time timestamp.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// User-provided description of the source.
	Description pulumi.StringOutput `pulumi:"description"`
	// The labels of the source.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The Source name.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Required. The source identifier.
	SourceId pulumi.StringOutput `pulumi:"sourceId"`
	// The update time timestamp.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Vmware type source details.
	Vmware VmwareSourceDetailsResponseOutput `pulumi:"vmware"`
}

// NewSource registers a new resource with the given unique name, arguments, and options.
func NewSource(ctx *pulumi.Context,
	name string, args *SourceArgs, opts ...pulumi.ResourceOption) (*Source, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"sourceId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Source
	err := ctx.RegisterResource("google-native:vmmigration/v1:Source", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSource gets an existing Source resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceState, opts ...pulumi.ResourceOption) (*Source, error) {
	var resource Source
	err := ctx.ReadResource("google-native:vmmigration/v1:Source", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Source resources.
type sourceState struct {
}

type SourceState struct {
}

func (SourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceState)(nil)).Elem()
}

type sourceArgs struct {
	// AWS type source details.
	Aws *AwsSourceDetails `pulumi:"aws"`
	// User-provided description of the source.
	Description *string `pulumi:"description"`
	// The labels of the source.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Required. The source identifier.
	SourceId string `pulumi:"sourceId"`
	// Vmware type source details.
	Vmware *VmwareSourceDetails `pulumi:"vmware"`
}

// The set of arguments for constructing a Source resource.
type SourceArgs struct {
	// AWS type source details.
	Aws AwsSourceDetailsPtrInput
	// User-provided description of the source.
	Description pulumi.StringPtrInput
	// The labels of the source.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Required. The source identifier.
	SourceId pulumi.StringInput
	// Vmware type source details.
	Vmware VmwareSourceDetailsPtrInput
}

func (SourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceArgs)(nil)).Elem()
}

type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(ctx context.Context) SourceOutput
}

func (*Source) ElementType() reflect.Type {
	return reflect.TypeOf((**Source)(nil)).Elem()
}

func (i *Source) ToSourceOutput() SourceOutput {
	return i.ToSourceOutputWithContext(context.Background())
}

func (i *Source) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutput)
}

type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

// AWS type source details.
func (o SourceOutput) Aws() AwsSourceDetailsResponseOutput {
	return o.ApplyT(func(v *Source) AwsSourceDetailsResponseOutput { return v.Aws }).(AwsSourceDetailsResponseOutput)
}

// The create time timestamp.
func (o SourceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// User-provided description of the source.
func (o SourceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The labels of the source.
func (o SourceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Source) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o SourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The Source name.
func (o SourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SourceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o SourceOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Source) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Required. The source identifier.
func (o SourceOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

// The update time timestamp.
func (o SourceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Vmware type source details.
func (o SourceOutput) Vmware() VmwareSourceDetailsResponseOutput {
	return o.ApplyT(func(v *Source) VmwareSourceDetailsResponseOutput { return v.Vmware }).(VmwareSourceDetailsResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceInput)(nil)).Elem(), &Source{})
	pulumi.RegisterOutputType(SourceOutput{})
}
