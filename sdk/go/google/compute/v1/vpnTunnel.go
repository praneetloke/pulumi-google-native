// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VpnTunnel resource in the specified project and region using the data included in the request.
type VpnTunnel struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Detailed status message for the VPN tunnel.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
	IkeVersion pulumi.IntOutput `pulumi:"ikeVersion"`
	// Type of resource. Always compute#vpnTunnel for VPN tunnels.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
	LocalTrafficSelector pulumi.StringArrayOutput `pulumi:"localTrafficSelector"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
	PeerExternalGateway pulumi.StringOutput `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. Possible values are: `0`, `1`, `2`, `3`. The number of IDs in use depends on the external VPN gateway redundancy type.
	PeerExternalGatewayInterface pulumi.IntOutput `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
	PeerGcpGateway pulumi.StringOutput `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp  pulumi.StringOutput `pulumi:"peerIp"`
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelector pulumi.StringArrayOutput `pulumi:"remoteTrafficSelector"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// URL of the router resource to be used for dynamic routing.
	Router pulumi.StringOutput `pulumi:"router"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret pulumi.StringOutput `pulumi:"sharedSecret"`
	// Hash of the shared secret.
	SharedSecretHash pulumi.StringOutput `pulumi:"sharedSecretHash"`
	// The status of the VPN tunnel, which can be one of the following: - PROVISIONING: Resource is being allocated for the VPN tunnel. - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. - ESTABLISHED: Secure session is successfully established with the peer VPN. - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). - NEGOTIATION_FAILURE: Handshake failed. - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. - NO_INCOMING_PACKETS: No incoming packets from peer. - REJECTED: Tunnel configuration was rejected, can be result of being denied access. - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel.
	Status pulumi.StringOutput `pulumi:"status"`
	// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
	TargetVpnGateway pulumi.StringOutput `pulumi:"targetVpnGateway"`
	// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
	VpnGateway pulumi.StringOutput `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated. Possible values are: `0`, `1`.
	VpnGatewayInterface pulumi.IntOutput `pulumi:"vpnGatewayInterface"`
}

// NewVpnTunnel registers a new resource with the given unique name, arguments, and options.
func NewVpnTunnel(ctx *pulumi.Context,
	name string, args *VpnTunnelArgs, opts ...pulumi.ResourceOption) (*VpnTunnel, error) {
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
	var resource VpnTunnel
	err := ctx.RegisterResource("google-native:compute/v1:VpnTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnTunnel gets an existing VpnTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnTunnelState, opts ...pulumi.ResourceOption) (*VpnTunnel, error) {
	var resource VpnTunnel
	err := ctx.ReadResource("google-native:compute/v1:VpnTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnTunnel resources.
type vpnTunnelState struct {
}

type VpnTunnelState struct {
}

func (VpnTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnTunnelState)(nil)).Elem()
}

type vpnTunnelArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
	IkeVersion *int `pulumi:"ikeVersion"`
	// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
	LocalTrafficSelector []string `pulumi:"localTrafficSelector"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
	PeerExternalGateway *string `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. Possible values are: `0`, `1`, `2`, `3`. The number of IDs in use depends on the external VPN gateway redundancy type.
	PeerExternalGatewayInterface *int `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
	PeerGcpGateway *string `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp  *string `pulumi:"peerIp"`
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
	// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelector []string `pulumi:"remoteTrafficSelector"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// URL of the router resource to be used for dynamic routing.
	Router *string `pulumi:"router"`
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret *string `pulumi:"sharedSecret"`
	// Hash of the shared secret.
	SharedSecretHash *string `pulumi:"sharedSecretHash"`
	// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
	TargetVpnGateway *string `pulumi:"targetVpnGateway"`
	// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
	VpnGateway *string `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated. Possible values are: `0`, `1`.
	VpnGatewayInterface *int `pulumi:"vpnGatewayInterface"`
}

// The set of arguments for constructing a VpnTunnel resource.
type VpnTunnelArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
	IkeVersion pulumi.IntPtrInput
	// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
	LocalTrafficSelector pulumi.StringArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
	PeerExternalGateway pulumi.StringPtrInput
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. Possible values are: `0`, `1`, `2`, `3`. The number of IDs in use depends on the external VPN gateway redundancy type.
	PeerExternalGatewayInterface pulumi.IntPtrInput
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
	PeerGcpGateway pulumi.StringPtrInput
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp  pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	Region  pulumi.StringInput
	// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelector pulumi.StringArrayInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// URL of the router resource to be used for dynamic routing.
	Router pulumi.StringPtrInput
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret pulumi.StringPtrInput
	// Hash of the shared secret.
	SharedSecretHash pulumi.StringPtrInput
	// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
	TargetVpnGateway pulumi.StringPtrInput
	// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
	VpnGateway pulumi.StringPtrInput
	// The interface ID of the VPN gateway with which this VPN tunnel is associated. Possible values are: `0`, `1`.
	VpnGatewayInterface pulumi.IntPtrInput
}

func (VpnTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnTunnelArgs)(nil)).Elem()
}

type VpnTunnelInput interface {
	pulumi.Input

	ToVpnTunnelOutput() VpnTunnelOutput
	ToVpnTunnelOutputWithContext(ctx context.Context) VpnTunnelOutput
}

func (*VpnTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnTunnel)(nil)).Elem()
}

func (i *VpnTunnel) ToVpnTunnelOutput() VpnTunnelOutput {
	return i.ToVpnTunnelOutputWithContext(context.Background())
}

func (i *VpnTunnel) ToVpnTunnelOutputWithContext(ctx context.Context) VpnTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnTunnelOutput)
}

type VpnTunnelOutput struct{ *pulumi.OutputState }

func (VpnTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnTunnel)(nil)).Elem()
}

func (o VpnTunnelOutput) ToVpnTunnelOutput() VpnTunnelOutput {
	return o
}

func (o VpnTunnelOutput) ToVpnTunnelOutputWithContext(ctx context.Context) VpnTunnelOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o VpnTunnelOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o VpnTunnelOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Detailed status message for the VPN tunnel.
func (o VpnTunnelOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
func (o VpnTunnelOutput) IkeVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.IntOutput { return v.IkeVersion }).(pulumi.IntOutput)
}

// Type of resource. Always compute#vpnTunnel for VPN tunnels.
func (o VpnTunnelOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
func (o VpnTunnelOutput) LocalTrafficSelector() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringArrayOutput { return v.LocalTrafficSelector }).(pulumi.StringArrayOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o VpnTunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
func (o VpnTunnelOutput) PeerExternalGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.PeerExternalGateway }).(pulumi.StringOutput)
}

// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. Possible values are: `0`, `1`, `2`, `3`. The number of IDs in use depends on the external VPN gateway redundancy type.
func (o VpnTunnelOutput) PeerExternalGatewayInterface() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.IntOutput { return v.PeerExternalGatewayInterface }).(pulumi.IntOutput)
}

// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
func (o VpnTunnelOutput) PeerGcpGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.PeerGcpGateway }).(pulumi.StringOutput)
}

// IP address of the peer VPN gateway. Only IPv4 is supported.
func (o VpnTunnelOutput) PeerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.PeerIp }).(pulumi.StringOutput)
}

func (o VpnTunnelOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o VpnTunnelOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
func (o VpnTunnelOutput) RemoteTrafficSelector() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringArrayOutput { return v.RemoteTrafficSelector }).(pulumi.StringArrayOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o VpnTunnelOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// URL of the router resource to be used for dynamic routing.
func (o VpnTunnelOutput) Router() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.Router }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o VpnTunnelOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
func (o VpnTunnelOutput) SharedSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.SharedSecret }).(pulumi.StringOutput)
}

// Hash of the shared secret.
func (o VpnTunnelOutput) SharedSecretHash() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.SharedSecretHash }).(pulumi.StringOutput)
}

// The status of the VPN tunnel, which can be one of the following: - PROVISIONING: Resource is being allocated for the VPN tunnel. - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. - ESTABLISHED: Secure session is successfully established with the peer VPN. - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). - NEGOTIATION_FAILURE: Handshake failed. - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. - NO_INCOMING_PACKETS: No incoming packets from peer. - REJECTED: Tunnel configuration was rejected, can be result of being denied access. - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel.
func (o VpnTunnelOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
func (o VpnTunnelOutput) TargetVpnGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.TargetVpnGateway }).(pulumi.StringOutput)
}

// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
func (o VpnTunnelOutput) VpnGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.StringOutput { return v.VpnGateway }).(pulumi.StringOutput)
}

// The interface ID of the VPN gateway with which this VPN tunnel is associated. Possible values are: `0`, `1`.
func (o VpnTunnelOutput) VpnGatewayInterface() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnTunnel) pulumi.IntOutput { return v.VpnGatewayInterface }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnTunnelInput)(nil)).Elem(), &VpnTunnel{})
	pulumi.RegisterOutputType(VpnTunnelOutput{})
}
