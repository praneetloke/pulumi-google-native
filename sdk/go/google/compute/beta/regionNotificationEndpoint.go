// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a NotificationEndpoint in the specified project in the given region using the parameters that are included in the request.
type RegionNotificationEndpoint struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
	GrpcSettings NotificationEndpointGrpcSettingsResponseOutput `pulumi:"grpcSettings"`
	// Type of the resource. Always compute#notificationEndpoint for notification endpoints.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewRegionNotificationEndpoint registers a new resource with the given unique name, arguments, and options.
func NewRegionNotificationEndpoint(ctx *pulumi.Context,
	name string, args *RegionNotificationEndpointArgs, opts ...pulumi.ResourceOption) (*RegionNotificationEndpoint, error) {
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
	var resource RegionNotificationEndpoint
	err := ctx.RegisterResource("google-native:compute/beta:RegionNotificationEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionNotificationEndpoint gets an existing RegionNotificationEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionNotificationEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionNotificationEndpointState, opts ...pulumi.ResourceOption) (*RegionNotificationEndpoint, error) {
	var resource RegionNotificationEndpoint
	err := ctx.ReadResource("google-native:compute/beta:RegionNotificationEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionNotificationEndpoint resources.
type regionNotificationEndpointState struct {
}

type RegionNotificationEndpointState struct {
}

func (RegionNotificationEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNotificationEndpointState)(nil)).Elem()
}

type regionNotificationEndpointArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
	GrpcSettings *NotificationEndpointGrpcSettings `pulumi:"grpcSettings"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a RegionNotificationEndpoint resource.
type RegionNotificationEndpointArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
	GrpcSettings NotificationEndpointGrpcSettingsPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	Region  pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
}

func (RegionNotificationEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNotificationEndpointArgs)(nil)).Elem()
}

type RegionNotificationEndpointInput interface {
	pulumi.Input

	ToRegionNotificationEndpointOutput() RegionNotificationEndpointOutput
	ToRegionNotificationEndpointOutputWithContext(ctx context.Context) RegionNotificationEndpointOutput
}

func (*RegionNotificationEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionNotificationEndpoint)(nil)).Elem()
}

func (i *RegionNotificationEndpoint) ToRegionNotificationEndpointOutput() RegionNotificationEndpointOutput {
	return i.ToRegionNotificationEndpointOutputWithContext(context.Background())
}

func (i *RegionNotificationEndpoint) ToRegionNotificationEndpointOutputWithContext(ctx context.Context) RegionNotificationEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNotificationEndpointOutput)
}

type RegionNotificationEndpointOutput struct{ *pulumi.OutputState }

func (RegionNotificationEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionNotificationEndpoint)(nil)).Elem()
}

func (o RegionNotificationEndpointOutput) ToRegionNotificationEndpointOutput() RegionNotificationEndpointOutput {
	return o
}

func (o RegionNotificationEndpointOutput) ToRegionNotificationEndpointOutputWithContext(ctx context.Context) RegionNotificationEndpointOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o RegionNotificationEndpointOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o RegionNotificationEndpointOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
func (o RegionNotificationEndpointOutput) GrpcSettings() NotificationEndpointGrpcSettingsResponseOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) NotificationEndpointGrpcSettingsResponseOutput {
		return v.GrpcSettings
	}).(NotificationEndpointGrpcSettingsResponseOutput)
}

// Type of the resource. Always compute#notificationEndpoint for notification endpoints.
func (o RegionNotificationEndpointOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o RegionNotificationEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegionNotificationEndpointOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o RegionNotificationEndpointOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o RegionNotificationEndpointOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o RegionNotificationEndpointOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNotificationEndpoint) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionNotificationEndpointInput)(nil)).Elem(), &RegionNotificationEndpoint{})
	pulumi.RegisterOutputType(RegionNotificationEndpointOutput{})
}
