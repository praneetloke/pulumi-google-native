// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified PacketMirroring resource.
func LookupPacketMirroring(ctx *pulumi.Context, args *LookupPacketMirroringArgs, opts ...pulumi.InvokeOption) (*LookupPacketMirroringResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPacketMirroringResult
	err := ctx.Invoke("google-native:compute/alpha:getPacketMirroring", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPacketMirroringArgs struct {
	PacketMirroring string  `pulumi:"packetMirroring"`
	Project         *string `pulumi:"project"`
	Region          string  `pulumi:"region"`
}

type LookupPacketMirroringResult struct {
	// The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
	CollectorIlb PacketMirroringForwardingRuleInfoResponse `pulumi:"collectorIlb"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
	Enable string `pulumi:"enable"`
	// Filter for mirrored traffic. If unspecified, all traffic is mirrored.
	Filter PacketMirroringFilterResponse `pulumi:"filter"`
	// Type of the resource. Always compute#packetMirroring for packet mirrorings.
	Kind string `pulumi:"kind"`
	// PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
	MirroredResources PacketMirroringMirroredResourceInfoResponse `pulumi:"mirroredResources"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
	Network PacketMirroringNetworkInfoResponse `pulumi:"network"`
	// The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
	Priority int `pulumi:"priority"`
	// URI of the region where the packetMirroring resides.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
}

func LookupPacketMirroringOutput(ctx *pulumi.Context, args LookupPacketMirroringOutputArgs, opts ...pulumi.InvokeOption) LookupPacketMirroringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPacketMirroringResult, error) {
			args := v.(LookupPacketMirroringArgs)
			r, err := LookupPacketMirroring(ctx, &args, opts...)
			var s LookupPacketMirroringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPacketMirroringResultOutput)
}

type LookupPacketMirroringOutputArgs struct {
	PacketMirroring pulumi.StringInput    `pulumi:"packetMirroring"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
	Region          pulumi.StringInput    `pulumi:"region"`
}

func (LookupPacketMirroringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketMirroringArgs)(nil)).Elem()
}

type LookupPacketMirroringResultOutput struct{ *pulumi.OutputState }

func (LookupPacketMirroringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketMirroringResult)(nil)).Elem()
}

func (o LookupPacketMirroringResultOutput) ToLookupPacketMirroringResultOutput() LookupPacketMirroringResultOutput {
	return o
}

func (o LookupPacketMirroringResultOutput) ToLookupPacketMirroringResultOutputWithContext(ctx context.Context) LookupPacketMirroringResultOutput {
	return o
}

// The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
func (o LookupPacketMirroringResultOutput) CollectorIlb() PacketMirroringForwardingRuleInfoResponseOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) PacketMirroringForwardingRuleInfoResponse { return v.CollectorIlb }).(PacketMirroringForwardingRuleInfoResponseOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupPacketMirroringResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupPacketMirroringResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) string { return v.Description }).(pulumi.StringOutput)
}

// Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
func (o LookupPacketMirroringResultOutput) Enable() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) string { return v.Enable }).(pulumi.StringOutput)
}

// Filter for mirrored traffic. If unspecified, all traffic is mirrored.
func (o LookupPacketMirroringResultOutput) Filter() PacketMirroringFilterResponseOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) PacketMirroringFilterResponse { return v.Filter }).(PacketMirroringFilterResponseOutput)
}

// Type of the resource. Always compute#packetMirroring for packet mirrorings.
func (o LookupPacketMirroringResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) string { return v.Kind }).(pulumi.StringOutput)
}

// PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
func (o LookupPacketMirroringResultOutput) MirroredResources() PacketMirroringMirroredResourceInfoResponseOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) PacketMirroringMirroredResourceInfoResponse {
		return v.MirroredResources
	}).(PacketMirroringMirroredResourceInfoResponseOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupPacketMirroringResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
func (o LookupPacketMirroringResultOutput) Network() PacketMirroringNetworkInfoResponseOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) PacketMirroringNetworkInfoResponse { return v.Network }).(PacketMirroringNetworkInfoResponseOutput)
}

// The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
func (o LookupPacketMirroringResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) int { return v.Priority }).(pulumi.IntOutput)
}

// URI of the region where the packetMirroring resides.
func (o LookupPacketMirroringResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupPacketMirroringResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupPacketMirroringResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketMirroringResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPacketMirroringResultOutput{})
}
