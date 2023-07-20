// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an address resource in the specified project by using the data included in the request.
type Address struct {
	pulumi.CustomResourceState

	// The static IP address represented by this resource.
	Address pulumi.StringOutput `pulumi:"address"`
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType pulumi.StringOutput `pulumi:"addressType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this address can be used after the external IPv6 address reservation.
	Ipv6EndpointType pulumi.StringOutput `pulumi:"ipv6EndpointType"`
	// Type of the resource. Always compute#address for addresses.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this Address, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Address.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network pulumi.StringOutput `pulumi:"network"`
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.IntOutput    `pulumi:"prefixLength"`
	Project      pulumi.StringOutput `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *HA VPN over Cloud Interconnect* configuration. These addresses are regional resources. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
	Status pulumi.StringOutput `pulumi:"status"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The URLs of the resources that are using this address.
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewAddress registers a new resource with the given unique name, arguments, and options.
func NewAddress(ctx *pulumi.Context,
	name string, args *AddressArgs, opts ...pulumi.ResourceOption) (*Address, error) {
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
	var resource Address
	err := ctx.RegisterResource("google-native:compute/v1:Address", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddress gets an existing Address resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressState, opts ...pulumi.ResourceOption) (*Address, error) {
	var resource Address
	err := ctx.ReadResource("google-native:compute/v1:Address", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Address resources.
type addressState struct {
}

type AddressState struct {
}

func (AddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressState)(nil)).Elem()
}

type addressArgs struct {
	// The static IP address represented by this resource.
	Address *string `pulumi:"address"`
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType *AddressAddressType `pulumi:"addressType"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description *string `pulumi:"description"`
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion *AddressIpVersion `pulumi:"ipVersion"`
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this address can be used after the external IPv6 address reservation.
	Ipv6EndpointType *AddressIpv6EndpointType `pulumi:"ipv6EndpointType"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name *string `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network *string `pulumi:"network"`
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier *AddressNetworkTier `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength *int    `pulumi:"prefixLength"`
	Project      *string `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *HA VPN over Cloud Interconnect* configuration. These addresses are regional resources. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	Purpose *AddressPurpose `pulumi:"purpose"`
	Region  string          `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork *string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a Address resource.
type AddressArgs struct {
	// The static IP address represented by this resource.
	Address pulumi.StringPtrInput
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType AddressAddressTypePtrInput
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringPtrInput
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion AddressIpVersionPtrInput
	// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this address can be used after the external IPv6 address reservation.
	Ipv6EndpointType AddressIpv6EndpointTypePtrInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringPtrInput
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network pulumi.StringPtrInput
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier AddressNetworkTierPtrInput
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.IntPtrInput
	Project      pulumi.StringPtrInput
	// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *HA VPN over Cloud Interconnect* configuration. These addresses are regional resources. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	Purpose AddressPurposePtrInput
	Region  pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork pulumi.StringPtrInput
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressArgs)(nil)).Elem()
}

type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(ctx context.Context) AddressOutput
}

func (*Address) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *Address) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i *Address) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

type AddressOutput struct{ *pulumi.OutputState }

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

// The static IP address represented by this resource.
func (o AddressOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
func (o AddressOutput) AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.AddressType }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o AddressOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this field when you create the resource.
func (o AddressOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
func (o AddressOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

// The endpoint type of this address, which should be VM or NETLB. This is used for deciding which type of endpoint this address can be used after the external IPv6 address reservation.
func (o AddressOutput) Ipv6EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Ipv6EndpointType }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#address for addresses.
func (o AddressOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this Address, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Address.
func (o AddressOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o AddressOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Address) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
func (o AddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
func (o AddressOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
func (o AddressOutput) NetworkTier() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.NetworkTier }).(pulumi.StringOutput)
}

// The prefix length if the resource represents an IP range.
func (o AddressOutput) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v *Address) pulumi.IntOutput { return v.PrefixLength }).(pulumi.IntOutput)
}

func (o AddressOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *HA VPN over Cloud Interconnect* configuration. These addresses are regional resources. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
func (o AddressOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

func (o AddressOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o AddressOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o AddressOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
func (o AddressOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
func (o AddressOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Subnetwork }).(pulumi.StringOutput)
}

// The URLs of the resources that are using this address.
func (o AddressOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Address) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressInput)(nil)).Elem(), &Address{})
	pulumi.RegisterOutputType(AddressOutput{})
}
