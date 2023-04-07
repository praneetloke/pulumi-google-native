// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a network in the specified project using the data included in the request.
type Network struct {
	pulumi.CustomResourceState

	// Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
	AutoCreateSubnetworks pulumi.BoolOutput `pulumi:"autoCreateSubnetworks"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
	EnableUlaInternalIpv6 pulumi.BoolOutput `pulumi:"enableUlaInternalIpv6"`
	// URL of the firewall policy the network is associated with.
	FirewallPolicy pulumi.StringOutput `pulumi:"firewallPolicy"`
	// The gateway address for default routing out of the network, selected by Google Cloud.
	GatewayIPv4 pulumi.StringOutput `pulumi:"gatewayIPv4"`
	// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
	InternalIpv6Range pulumi.StringOutput `pulumi:"internalIpv6Range"`
	// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
	//
	// Deprecated: Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
	Ipv4Range pulumi.StringOutput `pulumi:"ipv4Range"`
	// Type of the resource. Always compute#network for networks.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Maximum Transmission Unit in bytes. The minimum value for this field is 1300 and the maximum value is 8896. The suggested value is 1500, which is the default MTU used on the Internet, or 8896 if you want to use Jumbo frames. If unspecified, the value defaults to 1460.
	Mtu pulumi.IntOutput `pulumi:"mtu"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
	NetworkFirewallPolicyEnforcementOrder pulumi.StringOutput `pulumi:"networkFirewallPolicyEnforcementOrder"`
	// A list of network peerings for the resource.
	Peerings NetworkPeeringResponseArrayOutput `pulumi:"peerings"`
	Project  pulumi.StringOutput               `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
	RoutingConfig NetworkRoutingConfigResponseOutput `pulumi:"routingConfig"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Server-defined fully-qualified URLs for all subnetworks in this VPC network.
	Subnetworks pulumi.StringArrayOutput `pulumi:"subnetworks"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		args = &NetworkArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Network
	err := ctx.RegisterResource("google-native:compute/v1:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("google-native:compute/v1:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
}

type NetworkState struct {
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
	AutoCreateSubnetworks *bool `pulumi:"autoCreateSubnetworks"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description *string `pulumi:"description"`
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
	EnableUlaInternalIpv6 *bool `pulumi:"enableUlaInternalIpv6"`
	// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
	InternalIpv6Range *string `pulumi:"internalIpv6Range"`
	// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
	//
	// Deprecated: Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
	Ipv4Range *string `pulumi:"ipv4Range"`
	// Maximum Transmission Unit in bytes. The minimum value for this field is 1300 and the maximum value is 8896. The suggested value is 1500, which is the default MTU used on the Internet, or 8896 if you want to use Jumbo frames. If unspecified, the value defaults to 1460.
	Mtu *int `pulumi:"mtu"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name *string `pulumi:"name"`
	// The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
	NetworkFirewallPolicyEnforcementOrder *NetworkNetworkFirewallPolicyEnforcementOrder `pulumi:"networkFirewallPolicyEnforcementOrder"`
	Project                               *string                                       `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
	RoutingConfig *NetworkRoutingConfig `pulumi:"routingConfig"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
	AutoCreateSubnetworks pulumi.BoolPtrInput
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringPtrInput
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
	EnableUlaInternalIpv6 pulumi.BoolPtrInput
	// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
	InternalIpv6Range pulumi.StringPtrInput
	// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
	//
	// Deprecated: Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
	Ipv4Range pulumi.StringPtrInput
	// Maximum Transmission Unit in bytes. The minimum value for this field is 1300 and the maximum value is 8896. The suggested value is 1500, which is the default MTU used on the Internet, or 8896 if you want to use Jumbo frames. If unspecified, the value defaults to 1460.
	Mtu pulumi.IntPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringPtrInput
	// The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
	NetworkFirewallPolicyEnforcementOrder NetworkNetworkFirewallPolicyEnforcementOrderPtrInput
	Project                               pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
	RoutingConfig NetworkRoutingConfigPtrInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}

type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(ctx context.Context) NetworkOutput
}

func (*Network) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

type NetworkOutput struct{ *pulumi.OutputState }

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

// Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
func (o NetworkOutput) AutoCreateSubnetworks() pulumi.BoolOutput {
	return o.ApplyT(func(v *Network) pulumi.BoolOutput { return v.AutoCreateSubnetworks }).(pulumi.BoolOutput)
}

// Creation timestamp in RFC3339 text format.
func (o NetworkOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this field when you create the resource.
func (o NetworkOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
func (o NetworkOutput) EnableUlaInternalIpv6() pulumi.BoolOutput {
	return o.ApplyT(func(v *Network) pulumi.BoolOutput { return v.EnableUlaInternalIpv6 }).(pulumi.BoolOutput)
}

// URL of the firewall policy the network is associated with.
func (o NetworkOutput) FirewallPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.FirewallPolicy }).(pulumi.StringOutput)
}

// The gateway address for default routing out of the network, selected by Google Cloud.
func (o NetworkOutput) GatewayIPv4() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.GatewayIPv4 }).(pulumi.StringOutput)
}

// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
func (o NetworkOutput) InternalIpv6Range() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.InternalIpv6Range }).(pulumi.StringOutput)
}

// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
//
// Deprecated: Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
func (o NetworkOutput) Ipv4Range() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Ipv4Range }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#network for networks.
func (o NetworkOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Maximum Transmission Unit in bytes. The minimum value for this field is 1300 and the maximum value is 8896. The suggested value is 1500, which is the default MTU used on the Internet, or 8896 if you want to use Jumbo frames. If unspecified, the value defaults to 1460.
func (o NetworkOutput) Mtu() pulumi.IntOutput {
	return o.ApplyT(func(v *Network) pulumi.IntOutput { return v.Mtu }).(pulumi.IntOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
func (o NetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
func (o NetworkOutput) NetworkFirewallPolicyEnforcementOrder() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.NetworkFirewallPolicyEnforcementOrder }).(pulumi.StringOutput)
}

// A list of network peerings for the resource.
func (o NetworkOutput) Peerings() NetworkPeeringResponseArrayOutput {
	return o.ApplyT(func(v *Network) NetworkPeeringResponseArrayOutput { return v.Peerings }).(NetworkPeeringResponseArrayOutput)
}

func (o NetworkOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o NetworkOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
func (o NetworkOutput) RoutingConfig() NetworkRoutingConfigResponseOutput {
	return o.ApplyT(func(v *Network) NetworkRoutingConfigResponseOutput { return v.RoutingConfig }).(NetworkRoutingConfigResponseOutput)
}

// Server-defined URL for the resource.
func (o NetworkOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o NetworkOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// Server-defined fully-qualified URLs for all subnetworks in this VPC network.
func (o NetworkOutput) Subnetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Network) pulumi.StringArrayOutput { return v.Subnetworks }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInput)(nil)).Elem(), &Network{})
	pulumi.RegisterOutputType(NetworkOutput{})
}
