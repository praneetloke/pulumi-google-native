// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new AppConnection in a given project and location.
type AppConnection struct {
	pulumi.CustomResourceState

	// Optional. User-settable AppConnection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	AppConnectionId pulumi.StringPtrOutput `pulumi:"appConnectionId"`
	// Address of the remote application endpoint for the BeyondCorp AppConnection.
	ApplicationEndpoint GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionApplicationEndpointResponseOutput `pulumi:"applicationEndpoint"`
	// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
	Connectors pulumi.StringArrayOutput `pulumi:"connectors"`
	// Timestamp when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Gateway used by the AppConnection.
	Gateway GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayResponseOutput `pulumi:"gateway"`
	// Optional. Resource labels to represent user provided metadata.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The current state of the AppConnection.
	State pulumi.StringOutput `pulumi:"state"`
	// The type of network connectivity used by the AppConnection.
	Type pulumi.StringOutput `pulumi:"type"`
	// A unique identifier for the instance generated by the system.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Timestamp when the resource was last modified.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAppConnection registers a new resource with the given unique name, arguments, and options.
func NewAppConnection(ctx *pulumi.Context,
	name string, args *AppConnectionArgs, opts ...pulumi.ResourceOption) (*AppConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationEndpoint'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource AppConnection
	err := ctx.RegisterResource("google-native:beyondcorp/v1alpha:AppConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppConnection gets an existing AppConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppConnectionState, opts ...pulumi.ResourceOption) (*AppConnection, error) {
	var resource AppConnection
	err := ctx.ReadResource("google-native:beyondcorp/v1alpha:AppConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppConnection resources.
type appConnectionState struct {
}

type AppConnectionState struct {
}

func (AppConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectionState)(nil)).Elem()
}

type appConnectionArgs struct {
	// Optional. User-settable AppConnection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	AppConnectionId *string `pulumi:"appConnectionId"`
	// Address of the remote application endpoint for the BeyondCorp AppConnection.
	ApplicationEndpoint GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionApplicationEndpoint `pulumi:"applicationEndpoint"`
	// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
	Connectors []string `pulumi:"connectors"`
	// Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Gateway used by the AppConnection.
	Gateway *GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGateway `pulumi:"gateway"`
	// Optional. Resource labels to represent user provided metadata.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The type of network connectivity used by the AppConnection.
	Type AppConnectionType `pulumi:"type"`
}

// The set of arguments for constructing a AppConnection resource.
type AppConnectionArgs struct {
	// Optional. User-settable AppConnection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	AppConnectionId pulumi.StringPtrInput
	// Address of the remote application endpoint for the BeyondCorp AppConnection.
	ApplicationEndpoint GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionApplicationEndpointInput
	// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
	Connectors pulumi.StringArrayInput
	// Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
	DisplayName pulumi.StringPtrInput
	// Optional. Gateway used by the AppConnection.
	Gateway GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayPtrInput
	// Optional. Resource labels to represent user provided metadata.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The type of network connectivity used by the AppConnection.
	Type AppConnectionTypeInput
}

func (AppConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectionArgs)(nil)).Elem()
}

type AppConnectionInput interface {
	pulumi.Input

	ToAppConnectionOutput() AppConnectionOutput
	ToAppConnectionOutputWithContext(ctx context.Context) AppConnectionOutput
}

func (*AppConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnection)(nil)).Elem()
}

func (i *AppConnection) ToAppConnectionOutput() AppConnectionOutput {
	return i.ToAppConnectionOutputWithContext(context.Background())
}

func (i *AppConnection) ToAppConnectionOutputWithContext(ctx context.Context) AppConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConnectionOutput)
}

type AppConnectionOutput struct{ *pulumi.OutputState }

func (AppConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnection)(nil)).Elem()
}

func (o AppConnectionOutput) ToAppConnectionOutput() AppConnectionOutput {
	return o
}

func (o AppConnectionOutput) ToAppConnectionOutputWithContext(ctx context.Context) AppConnectionOutput {
	return o
}

// Optional. User-settable AppConnection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
func (o AppConnectionOutput) AppConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringPtrOutput { return v.AppConnectionId }).(pulumi.StringPtrOutput)
}

// Address of the remote application endpoint for the BeyondCorp AppConnection.
func (o AppConnectionOutput) ApplicationEndpoint() GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionApplicationEndpointResponseOutput {
	return o.ApplyT(func(v *AppConnection) GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionApplicationEndpointResponseOutput {
		return v.ApplicationEndpoint
	}).(GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionApplicationEndpointResponseOutput)
}

// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this AppConnection.
func (o AppConnectionOutput) Connectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringArrayOutput { return v.Connectors }).(pulumi.StringArrayOutput)
}

// Timestamp when the resource was created.
func (o AppConnectionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. An arbitrary user-provided name for the AppConnection. Cannot exceed 64 characters.
func (o AppConnectionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Gateway used by the AppConnection.
func (o AppConnectionOutput) Gateway() GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayResponseOutput {
	return o.ApplyT(func(v *AppConnection) GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayResponseOutput {
		return v.Gateway
	}).(GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayResponseOutput)
}

// Optional. Resource labels to represent user provided metadata.
func (o AppConnectionOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o AppConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Unique resource name of the AppConnection. The name is ignored when creating a AppConnection.
func (o AppConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppConnectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o AppConnectionOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The current state of the AppConnection.
func (o AppConnectionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The type of network connectivity used by the AppConnection.
func (o AppConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A unique identifier for the instance generated by the system.
func (o AppConnectionOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Timestamp when the resource was last modified.
func (o AppConnectionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnection) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppConnectionInput)(nil)).Elem(), &AppConnection{})
	pulumi.RegisterOutputType(AppConnectionOutput{})
}
