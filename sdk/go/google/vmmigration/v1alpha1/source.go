// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Source in a given project and location.
// Auto-naming is currently not supported for this resource.
type Source struct {
	pulumi.CustomResourceState

	// The create time timestamp.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// User-provided description of the source.
	Description pulumi.StringOutput `pulumi:"description"`
	// Provides details on the state of the Source in case of an error.
	Error StatusResponseOutput `pulumi:"error"`
	// The labels of the source.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The Source name.
	Name pulumi.StringOutput `pulumi:"name"`
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
	var resource Source
	err := ctx.RegisterResource("google-native:vmmigration/v1alpha1:Source", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:vmmigration/v1alpha1:Source", name, id, state, &resource, opts...)
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
	// User-provided description of the source.
	Description *string `pulumi:"description"`
	// The labels of the source.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Required. The source identifier.
	SourceId string `pulumi:"sourceId"`
	// Vmware type source details.
	Vmware *VmwareSourceDetails `pulumi:"vmware"`
}

// The set of arguments for constructing a Source resource.
type SourceArgs struct {
	// User-provided description of the source.
	Description pulumi.StringPtrInput
	// The labels of the source.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceInput)(nil)).Elem(), &Source{})
	pulumi.RegisterOutputType(SourceOutput{})
}
