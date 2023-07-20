// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified Route resource.
func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteResult
	err := ctx.Invoke("google-native:compute/beta:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	Project *string `pulumi:"project"`
	Route   string  `pulumi:"route"`
}

type LookupRouteResult struct {
	// AS path.
	AsPaths []RouteAsPathResponse `pulumi:"asPaths"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description string `pulumi:"description"`
	// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
	DestRange string `pulumi:"destRange"`
	// Type of this resource. Always compute#routes for Route resources.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name string `pulumi:"name"`
	// Fully-qualified URL of the network that this route applies to.
	Network string `pulumi:"network"`
	// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL: projects/ project/global/gateways/default-internet-gateway
	NextHopGateway string `pulumi:"nextHopGateway"`
	// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs: - 10.128.0.56 - https://www.googleapis.com/compute/v1/projects/project/regions/region /forwardingRules/forwardingRule - regions/region/forwardingRules/forwardingRule
	NextHopIlb string `pulumi:"nextHopIlb"`
	// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
	NextHopInstance string `pulumi:"nextHopInstance"`
	// The URL to an InterconnectAttachment which is the next hop for the route. This field will only be populated for the dynamic routes generated by Cloud Router with a linked interconnectAttachment.
	NextHopInterconnectAttachment string `pulumi:"nextHopInterconnectAttachment"`
	// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
	NextHopIp string `pulumi:"nextHopIp"`
	// The URL of the local network if it should handle matching packets.
	NextHopNetwork string `pulumi:"nextHopNetwork"`
	// The network peering name that should handle matching packets, which should conform to RFC1035.
	NextHopPeering string `pulumi:"nextHopPeering"`
	// The URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel string `pulumi:"nextHopVpnTunnel"`
	// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
	Priority int `pulumi:"priority"`
	// [Output only] The status of the route.
	RouteStatus string `pulumi:"routeStatus"`
	// The type of this route, which can be one of the following values: - 'TRANSIT' for a transit route that this router learned from another Cloud Router and will readvertise to one of its BGP peers - 'SUBNET' for a route from a subnet of the VPC - 'BGP' for a route learned from a BGP peer of this router - 'STATIC' for a static route
	RouteType string `pulumi:"routeType"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink string `pulumi:"selfLink"`
	// A list of instance tags to which this route applies.
	Tags []string `pulumi:"tags"`
	// If potential misconfigurations are detected for this route, this field will be populated with warning messages.
	Warnings []RouteWarningsItemResponse `pulumi:"warnings"`
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

type LookupRouteOutputArgs struct {
	Project pulumi.StringPtrInput `pulumi:"project"`
	Route   pulumi.StringInput    `pulumi:"route"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}

type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

// AS path.
func (o LookupRouteResultOutput) AsPaths() RouteAsPathResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []RouteAsPathResponse { return v.AsPaths }).(RouteAsPathResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupRouteResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this field when you create the resource.
func (o LookupRouteResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Description }).(pulumi.StringOutput)
}

// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
func (o LookupRouteResultOutput) DestRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.DestRange }).(pulumi.StringOutput)
}

// Type of this resource. Always compute#routes for Route resources.
func (o LookupRouteResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
func (o LookupRouteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Name }).(pulumi.StringOutput)
}

// Fully-qualified URL of the network that this route applies to.
func (o LookupRouteResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Network }).(pulumi.StringOutput)
}

// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL: projects/ project/global/gateways/default-internet-gateway
func (o LookupRouteResultOutput) NextHopGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopGateway }).(pulumi.StringOutput)
}

// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs: - 10.128.0.56 - https://www.googleapis.com/compute/v1/projects/project/regions/region /forwardingRules/forwardingRule - regions/region/forwardingRules/forwardingRule
func (o LookupRouteResultOutput) NextHopIlb() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopIlb }).(pulumi.StringOutput)
}

// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
func (o LookupRouteResultOutput) NextHopInstance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopInstance }).(pulumi.StringOutput)
}

// The URL to an InterconnectAttachment which is the next hop for the route. This field will only be populated for the dynamic routes generated by Cloud Router with a linked interconnectAttachment.
func (o LookupRouteResultOutput) NextHopInterconnectAttachment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopInterconnectAttachment }).(pulumi.StringOutput)
}

// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
func (o LookupRouteResultOutput) NextHopIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopIp }).(pulumi.StringOutput)
}

// The URL of the local network if it should handle matching packets.
func (o LookupRouteResultOutput) NextHopNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopNetwork }).(pulumi.StringOutput)
}

// The network peering name that should handle matching packets, which should conform to RFC1035.
func (o LookupRouteResultOutput) NextHopPeering() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopPeering }).(pulumi.StringOutput)
}

// The URL to a VpnTunnel that should handle matching packets.
func (o LookupRouteResultOutput) NextHopVpnTunnel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopVpnTunnel }).(pulumi.StringOutput)
}

// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
func (o LookupRouteResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouteResult) int { return v.Priority }).(pulumi.IntOutput)
}

// [Output only] The status of the route.
func (o LookupRouteResultOutput) RouteStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.RouteStatus }).(pulumi.StringOutput)
}

// The type of this route, which can be one of the following values: - 'TRANSIT' for a transit route that this router learned from another Cloud Router and will readvertise to one of its BGP peers - 'SUBNET' for a route from a subnet of the VPC - 'BGP' for a route learned from a BGP peer of this router - 'STATIC' for a static route
func (o LookupRouteResultOutput) RouteType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.RouteType }).(pulumi.StringOutput)
}

// Server-defined fully-qualified URL for this resource.
func (o LookupRouteResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// A list of instance tags to which this route applies.
func (o LookupRouteResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// If potential misconfigurations are detected for this route, this field will be populated with warning messages.
func (o LookupRouteResultOutput) Warnings() RouteWarningsItemResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []RouteWarningsItemResponse { return v.Warnings }).(RouteWarningsItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}
