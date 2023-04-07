// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified address resource.
func LookupGlobalAddress(ctx *pulumi.Context, args *LookupGlobalAddressArgs, opts ...pulumi.InvokeOption) (*LookupGlobalAddressResult, error) {
	var rv LookupGlobalAddressResult
	err := ctx.Invoke("google-native:compute/alpha:getGlobalAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalAddressArgs struct {
	Address string  `pulumi:"address"`
	Project *string `pulumi:"project"`
}

type LookupGlobalAddressResult struct {
	// The static IP address represented by this resource.
	Address string `pulumi:"address"`
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType string `pulumi:"addressType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description string `pulumi:"description"`
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion string `pulumi:"ipVersion"`
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this address can be used after the external IPv6 address reservation.
	Ipv6EndpointType string `pulumi:"ipv6EndpointType"`
	// Type of the resource. Always compute#address for addresses.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this Address, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Address.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name string `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network string `pulumi:"network"`
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier string `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength int `pulumi:"prefixLength"`
	// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *HA VPN over Cloud Interconnect* configuration. These addresses are regional resources. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	Purpose string `pulumi:"purpose"`
	// The URL of the region where a regional address resides. For regional addresses, you must specify the region as a path parameter in the HTTP request URL. *This field is not applicable to global addresses.*
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
	Status string `pulumi:"status"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork string `pulumi:"subnetwork"`
	// The URLs of the resources that are using this address.
	Users []string `pulumi:"users"`
}

func LookupGlobalAddressOutput(ctx *pulumi.Context, args LookupGlobalAddressOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalAddressResult, error) {
			args := v.(LookupGlobalAddressArgs)
			r, err := LookupGlobalAddress(ctx, &args, opts...)
			var s LookupGlobalAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalAddressResultOutput)
}

type LookupGlobalAddressOutputArgs struct {
	Address pulumi.StringInput    `pulumi:"address"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGlobalAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalAddressArgs)(nil)).Elem()
}

type LookupGlobalAddressResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalAddressResult)(nil)).Elem()
}

func (o LookupGlobalAddressResultOutput) ToLookupGlobalAddressResultOutput() LookupGlobalAddressResultOutput {
	return o
}

func (o LookupGlobalAddressResultOutput) ToLookupGlobalAddressResultOutputWithContext(ctx context.Context) LookupGlobalAddressResultOutput {
	return o
}

// The static IP address represented by this resource.
func (o LookupGlobalAddressResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Address }).(pulumi.StringOutput)
}

// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
func (o LookupGlobalAddressResultOutput) AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.AddressType }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupGlobalAddressResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this field when you create the resource.
func (o LookupGlobalAddressResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Description }).(pulumi.StringOutput)
}

// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
func (o LookupGlobalAddressResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this address can be used after the external IPv6 address reservation.
func (o LookupGlobalAddressResultOutput) Ipv6EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Ipv6EndpointType }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#address for addresses.
func (o LookupGlobalAddressResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this Address, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Address.
func (o LookupGlobalAddressResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupGlobalAddressResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
func (o LookupGlobalAddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
func (o LookupGlobalAddressResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Network }).(pulumi.StringOutput)
}

// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
func (o LookupGlobalAddressResultOutput) NetworkTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.NetworkTier }).(pulumi.StringOutput)
}

// The prefix length if the resource represents an IP range.
func (o LookupGlobalAddressResultOutput) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) int { return v.PrefixLength }).(pulumi.IntOutput)
}

// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *HA VPN over Cloud Interconnect* configuration. These addresses are regional resources. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
func (o LookupGlobalAddressResultOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Purpose }).(pulumi.StringOutput)
}

// The URL of the region where a regional address resides. For regional addresses, you must specify the region as a path parameter in the HTTP request URL. *This field is not applicable to global addresses.*
func (o LookupGlobalAddressResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupGlobalAddressResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupGlobalAddressResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
func (o LookupGlobalAddressResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Status }).(pulumi.StringOutput)
}

// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
func (o LookupGlobalAddressResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

// The URLs of the resources that are using this address.
func (o LookupGlobalAddressResultOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGlobalAddressResult) []string { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalAddressResultOutput{})
}
