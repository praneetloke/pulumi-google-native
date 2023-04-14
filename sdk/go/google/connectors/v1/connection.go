// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Connection in a given project and location.
// Auto-naming is currently not supported for this resource.
type Connection struct {
	pulumi.CustomResourceState

	// Optional. Configuration for establishing the connection's authentication with an external system.
	AuthConfig AuthConfigResponseOutput `pulumi:"authConfig"`
	// Optional. Configuration for configuring the connection with an external system.
	ConfigVariables ConfigVariableResponseArrayOutput `pulumi:"configVariables"`
	// Required. Identifier to assign to the Connection. Must be unique within scope of the parent resource.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// Connector version on which the connection is created. The format is: projects/*/locations/*/providers/*/connectors/*/versions/* Only global location is supported for ConnectorVersion resource.
	ConnectorVersion pulumi.StringOutput `pulumi:"connectorVersion"`
	// Created time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Description of the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
	DestinationConfigs DestinationConfigResponseArrayOutput `pulumi:"destinationConfigs"`
	// GCR location where the envoy image is stored. formatted like: gcr.io/{bucketName}/{imageName}
	EnvoyImageLocation pulumi.StringOutput `pulumi:"envoyImageLocation"`
	// GCR location where the runtime image is stored. formatted like: gcr.io/{bucketName}/{imageName}
	ImageLocation pulumi.StringOutput `pulumi:"imageLocation"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Optional. Configuration that indicates whether or not the Connection can be edited.
	LockConfig LockConfigResponseOutput `pulumi:"lockConfig"`
	// Optional. Log configuration for the connection.
	LogConfig ConnectorsLogConfigResponseOutput `pulumi:"logConfig"`
	// Resource name of the Connection. Format: projects/{project}/locations/{location}/connections/{connection}
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Node configuration for the connection.
	NodeConfig NodeConfigResponseOutput `pulumi:"nodeConfig"`
	Project    pulumi.StringOutput      `pulumi:"project"`
	// Optional. Service account needed for runtime plane to access GCP resources.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g. "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
	ServiceDirectory pulumi.StringOutput `pulumi:"serviceDirectory"`
	// Optional. Ssl config of a connection
	SslConfig SslConfigResponseOutput `pulumi:"sslConfig"`
	// Current status of the connection.
	Status ConnectionStatusResponseOutput `pulumi:"status"`
	// Optional. Suspended indicates if a user has suspended a connection or not.
	Suspended pulumi.BoolOutput `pulumi:"suspended"`
	// Updated time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.ConnectorVersion == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorVersion'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"connectionId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Connection
	err := ctx.RegisterResource("google-native:connectors/v1:Connection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:connectors/v1:Connection", name, id, state, &resource, opts...)
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
	// Optional. Configuration for establishing the connection's authentication with an external system.
	AuthConfig *AuthConfig `pulumi:"authConfig"`
	// Optional. Configuration for configuring the connection with an external system.
	ConfigVariables []ConfigVariable `pulumi:"configVariables"`
	// Required. Identifier to assign to the Connection. Must be unique within scope of the parent resource.
	ConnectionId string `pulumi:"connectionId"`
	// Connector version on which the connection is created. The format is: projects/*/locations/*/providers/*/connectors/*/versions/* Only global location is supported for ConnectorVersion resource.
	ConnectorVersion string `pulumi:"connectorVersion"`
	// Optional. Description of the resource.
	Description *string `pulumi:"description"`
	// Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
	DestinationConfigs []DestinationConfig `pulumi:"destinationConfigs"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Configuration that indicates whether or not the Connection can be edited.
	LockConfig *LockConfig `pulumi:"lockConfig"`
	// Optional. Log configuration for the connection.
	LogConfig *ConnectorsLogConfig `pulumi:"logConfig"`
	// Optional. Node configuration for the connection.
	NodeConfig *NodeConfig `pulumi:"nodeConfig"`
	Project    *string     `pulumi:"project"`
	// Optional. Service account needed for runtime plane to access GCP resources.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Optional. Ssl config of a connection
	SslConfig *SslConfig `pulumi:"sslConfig"`
	// Optional. Suspended indicates if a user has suspended a connection or not.
	Suspended *bool `pulumi:"suspended"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Optional. Configuration for establishing the connection's authentication with an external system.
	AuthConfig AuthConfigPtrInput
	// Optional. Configuration for configuring the connection with an external system.
	ConfigVariables ConfigVariableArrayInput
	// Required. Identifier to assign to the Connection. Must be unique within scope of the parent resource.
	ConnectionId pulumi.StringInput
	// Connector version on which the connection is created. The format is: projects/*/locations/*/providers/*/connectors/*/versions/* Only global location is supported for ConnectorVersion resource.
	ConnectorVersion pulumi.StringInput
	// Optional. Description of the resource.
	Description pulumi.StringPtrInput
	// Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
	DestinationConfigs DestinationConfigArrayInput
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Configuration that indicates whether or not the Connection can be edited.
	LockConfig LockConfigPtrInput
	// Optional. Log configuration for the connection.
	LogConfig ConnectorsLogConfigPtrInput
	// Optional. Node configuration for the connection.
	NodeConfig NodeConfigPtrInput
	Project    pulumi.StringPtrInput
	// Optional. Service account needed for runtime plane to access GCP resources.
	ServiceAccount pulumi.StringPtrInput
	// Optional. Ssl config of a connection
	SslConfig SslConfigPtrInput
	// Optional. Suspended indicates if a user has suspended a connection or not.
	Suspended pulumi.BoolPtrInput
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

// Optional. Configuration for establishing the connection's authentication with an external system.
func (o ConnectionOutput) AuthConfig() AuthConfigResponseOutput {
	return o.ApplyT(func(v *Connection) AuthConfigResponseOutput { return v.AuthConfig }).(AuthConfigResponseOutput)
}

// Optional. Configuration for configuring the connection with an external system.
func (o ConnectionOutput) ConfigVariables() ConfigVariableResponseArrayOutput {
	return o.ApplyT(func(v *Connection) ConfigVariableResponseArrayOutput { return v.ConfigVariables }).(ConfigVariableResponseArrayOutput)
}

// Required. Identifier to assign to the Connection. Must be unique within scope of the parent resource.
func (o ConnectionOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ConnectionId }).(pulumi.StringOutput)
}

// Connector version on which the connection is created. The format is: projects/*/locations/*/providers/*/connectors/*/versions/* Only global location is supported for ConnectorVersion resource.
func (o ConnectionOutput) ConnectorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ConnectorVersion }).(pulumi.StringOutput)
}

// Created time.
func (o ConnectionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the resource.
func (o ConnectionOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
func (o ConnectionOutput) DestinationConfigs() DestinationConfigResponseArrayOutput {
	return o.ApplyT(func(v *Connection) DestinationConfigResponseArrayOutput { return v.DestinationConfigs }).(DestinationConfigResponseArrayOutput)
}

// GCR location where the envoy image is stored. formatted like: gcr.io/{bucketName}/{imageName}
func (o ConnectionOutput) EnvoyImageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.EnvoyImageLocation }).(pulumi.StringOutput)
}

// GCR location where the runtime image is stored. formatted like: gcr.io/{bucketName}/{imageName}
func (o ConnectionOutput) ImageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ImageLocation }).(pulumi.StringOutput)
}

// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
func (o ConnectionOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o ConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. Configuration that indicates whether or not the Connection can be edited.
func (o ConnectionOutput) LockConfig() LockConfigResponseOutput {
	return o.ApplyT(func(v *Connection) LockConfigResponseOutput { return v.LockConfig }).(LockConfigResponseOutput)
}

// Optional. Log configuration for the connection.
func (o ConnectionOutput) LogConfig() ConnectorsLogConfigResponseOutput {
	return o.ApplyT(func(v *Connection) ConnectorsLogConfigResponseOutput { return v.LogConfig }).(ConnectorsLogConfigResponseOutput)
}

// Resource name of the Connection. Format: projects/{project}/locations/{location}/connections/{connection}
func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. Node configuration for the connection.
func (o ConnectionOutput) NodeConfig() NodeConfigResponseOutput {
	return o.ApplyT(func(v *Connection) NodeConfigResponseOutput { return v.NodeConfig }).(NodeConfigResponseOutput)
}

func (o ConnectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. Service account needed for runtime plane to access GCP resources.
func (o ConnectionOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g. "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
func (o ConnectionOutput) ServiceDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ServiceDirectory }).(pulumi.StringOutput)
}

// Optional. Ssl config of a connection
func (o ConnectionOutput) SslConfig() SslConfigResponseOutput {
	return o.ApplyT(func(v *Connection) SslConfigResponseOutput { return v.SslConfig }).(SslConfigResponseOutput)
}

// Current status of the connection.
func (o ConnectionOutput) Status() ConnectionStatusResponseOutput {
	return o.ApplyT(func(v *Connection) ConnectionStatusResponseOutput { return v.Status }).(ConnectionStatusResponseOutput)
}

// Optional. Suspended indicates if a user has suspended a connection or not.
func (o ConnectionOutput) Suspended() pulumi.BoolOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolOutput { return v.Suspended }).(pulumi.BoolOutput)
}

// Updated time.
func (o ConnectionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterOutputType(ConnectionOutput{})
}
