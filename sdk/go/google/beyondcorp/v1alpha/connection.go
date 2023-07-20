// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Connection in a given project and location.
type Connection struct {
	pulumi.CustomResourceState

	// Address of the remote application endpoint for the BeyondCorp Connection.
	ApplicationEndpoint ApplicationEndpointResponseOutput `pulumi:"applicationEndpoint"`
	// Optional. User-settable connection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	ConnectionId pulumi.StringPtrOutput `pulumi:"connectionId"`
	// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
	Connectors pulumi.StringArrayOutput `pulumi:"connectors"`
	// Timestamp when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Gateway used by the connection.
	Gateway GatewayResponseOutput `pulumi:"gateway"`
	// Optional. Resource labels to represent user provided metadata.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Unique resource name of the connection. The name is ignored when creating a connection.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The current state of the connection.
	State pulumi.StringOutput `pulumi:"state"`
	// The type of network connectivity used by the connection.
	Type pulumi.StringOutput `pulumi:"type"`
	// A unique identifier for the instance generated by the system.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Timestamp when the resource was last modified.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("google-native:beyondcorp/v1alpha:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("google-native:beyondcorp/v1alpha:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Address of the remote application endpoint for the BeyondCorp Connection.
	ApplicationEndpoint ApplicationEndpoint `pulumi:"applicationEndpoint"`
	// Optional. User-settable connection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	ConnectionId *string `pulumi:"connectionId"`
	// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
	Connectors []string `pulumi:"connectors"`
	// Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Gateway used by the connection.
	Gateway *Gateway `pulumi:"gateway"`
	// Optional. Resource labels to represent user provided metadata.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Unique resource name of the connection. The name is ignored when creating a connection.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The type of network connectivity used by the connection.
	Type ConnectionType `pulumi:"type"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Address of the remote application endpoint for the BeyondCorp Connection.
	ApplicationEndpoint ApplicationEndpointInput
	// Optional. User-settable connection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	ConnectionId pulumi.StringPtrInput
	// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
	Connectors pulumi.StringArrayInput
	// Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
	DisplayName pulumi.StringPtrInput
	// Optional. Gateway used by the connection.
	Gateway GatewayPtrInput
	// Optional. Resource labels to represent user provided metadata.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Unique resource name of the connection. The name is ignored when creating a connection.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The type of network connectivity used by the connection.
	Type ConnectionTypeInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

// Address of the remote application endpoint for the BeyondCorp Connection.
func (o ConnectionOutput) ApplicationEndpoint() ApplicationEndpointResponseOutput {
	return o.ApplyT(func(v *Connection) ApplicationEndpointResponseOutput { return v.ApplicationEndpoint }).(ApplicationEndpointResponseOutput)
}

// Optional. User-settable connection resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
func (o ConnectionOutput) ConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.ConnectionId }).(pulumi.StringPtrOutput)
}

// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
func (o ConnectionOutput) Connectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringArrayOutput { return v.Connectors }).(pulumi.StringArrayOutput)
}

// Timestamp when the resource was created.
func (o ConnectionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
func (o ConnectionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Gateway used by the connection.
func (o ConnectionOutput) Gateway() GatewayResponseOutput {
	return o.ApplyT(func(v *Connection) GatewayResponseOutput { return v.Gateway }).(GatewayResponseOutput)
}

// Optional. Resource labels to represent user provided metadata.
func (o ConnectionOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o ConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Unique resource name of the connection. The name is ignored when creating a connection.
func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o ConnectionOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The current state of the connection.
func (o ConnectionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The type of network connectivity used by the connection.
func (o ConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A unique identifier for the instance generated by the system.
func (o ConnectionOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Timestamp when the resource was last modified.
func (o ConnectionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterOutputType(ConnectionOutput{})
}
