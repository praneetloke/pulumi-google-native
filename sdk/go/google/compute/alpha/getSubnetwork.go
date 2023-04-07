// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified subnetwork.
func LookupSubnetwork(ctx *pulumi.Context, args *LookupSubnetworkArgs, opts ...pulumi.InvokeOption) (*LookupSubnetworkResult, error) {
	var rv LookupSubnetworkResult
	err := ctx.Invoke("google-native:compute/alpha:getSubnetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubnetworkArgs struct {
	Project    *string `pulumi:"project"`
	Region     string  `pulumi:"region"`
	Subnetwork string  `pulumi:"subnetwork"`
}

type LookupSubnetworkResult struct {
	// Can only be specified if VPC flow logging for this subnetwork is enabled. Sets the aggregation interval for collecting flow logs. Increasing the interval time reduces the amount of generated flow logs for long-lasting connections. Default is an interval of 5 seconds per connection. Valid values: INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN.
	AggregationInterval string `pulumi:"aggregationInterval"`
	// Whether this subnetwork's ranges can conflict with existing static routes. Setting this to true allows this subnetwork's primary and secondary ranges to overlap with (and contain) static routes that have already been configured on the corresponding network. For example if a static route has range 10.1.0.0/16, a subnet range 10.0.0.0/8 could only be created if allow_conflicting_routes=true. Overlapping is only allowed on subnetwork operations; routes whose ranges conflict with this subnetwork's ranges won't be allowed unless route.allow_conflicting_subnetworks is set to true. Typically packets destined to IPs within the subnetwork (which may contain private/sensitive data) are prevented from leaving the virtual network. Setting this field to true will disable this feature. The default value is false and applies to all existing subnetworks and automatically created subnetworks. This field cannot be set to true at resource creation time.
	AllowSubnetCidrRoutesOverlap bool `pulumi:"allowSubnetCidrRoutesOverlap"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
	Description string `pulumi:"description"`
	// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is determined by the org policy, if there is no org policy specified, then it will default to disabled. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	EnableFlowLogs bool `pulumi:"enableFlowLogs"`
	// Enables Layer2 communication on the subnetwork.
	EnableL2 bool `pulumi:"enableL2"`
	// Deprecated in favor of enable in PrivateIpv6GoogleAccess. Whether the VMs in this subnet can directly access Google services via internal IPv6 addresses. This field can be both set at resource creation time and updated using patch.
	//
	// Deprecated: Deprecated in favor of enable in PrivateIpv6GoogleAccess. Whether the VMs in this subnet can directly access Google services via internal IPv6 addresses. This field can be both set at resource creation time and updated using patch.
	EnablePrivateV6Access bool `pulumi:"enablePrivateV6Access"`
	// The external IPv6 address range that is owned by this subnetwork.
	ExternalIpv6Prefix string `pulumi:"externalIpv6Prefix"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
	Fingerprint string `pulumi:"fingerprint"`
	// Can only be specified if VPC flow logging for this subnetwork is enabled. The value of the field must be in [0, 1]. Set the sampling rate of VPC flow logs within the subnetwork where 1.0 means all collected logs are reported and 0.0 means no logs are reported. Default is 0.5 unless otherwise specified by the org policy, which means half of all collected logs are reported.
	FlowSampling float64 `pulumi:"flowSampling"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress string `pulumi:"gatewayAddress"`
	// The internal IPv6 address range that is assigned to this subnetwork.
	InternalIpv6Prefix string `pulumi:"internalIpv6Prefix"`
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
	IpCidrRange string `pulumi:"ipCidrRange"`
	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first time the subnet is updated into IPV4_IPV6 dual stack.
	Ipv6AccessType string `pulumi:"ipv6AccessType"`
	// This field is for internal use.
	Ipv6CidrRange string `pulumi:"ipv6CidrRange"`
	// Type of the resource. Always compute#subnetwork for Subnetwork resources.
	Kind string `pulumi:"kind"`
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
	LogConfig SubnetworkLogConfigResponse `pulumi:"logConfig"`
	// Can only be specified if VPC flow logging for this subnetwork is enabled. Configures whether metadata fields should be added to the reported VPC flow logs. Options are INCLUDE_ALL_METADATA, EXCLUDE_ALL_METADATA, and CUSTOM_METADATA. Default is EXCLUDE_ALL_METADATA.
	Metadata string `pulumi:"metadata"`
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. This field can be set only at resource creation time.
	Network string `pulumi:"network"`
	// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
	PrivateIpGoogleAccess bool `pulumi:"privateIpGoogleAccess"`
	// This field is for internal use. This field can be both set at resource creation time and updated using patch.
	PrivateIpv6GoogleAccess string `pulumi:"privateIpv6GoogleAccess"`
	// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	Purpose string `pulumi:"purpose"`
	// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
	Region string `pulumi:"region"`
	// The URL of the reserved internal range.
	ReservedInternalRange string `pulumi:"reservedInternalRange"`
	// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
	Role string `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
	SecondaryIpRanges []SubnetworkSecondaryRangeResponse `pulumi:"secondaryIpRanges"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// The stack type for the subnet. If set to IPV4_ONLY, new VMs in the subnet are assigned IPv4 addresses only. If set to IPV4_IPV6, new VMs in the subnet can be assigned both IPv4 and IPv6 addresses. If not specified, IPV4_ONLY is used. This field can be both set at resource creation time and updated using patch.
	StackType string `pulumi:"stackType"`
	// The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY
	State string `pulumi:"state"`
	// A repeated field indicating the VLAN IDs supported on this subnetwork. During Subnet creation, specifying vlan is valid only if enable_l2 is true. During Subnet Update, specifying vlan is allowed only for l2 enabled subnets. Restricted to only one VLAN.
	Vlans []int `pulumi:"vlans"`
}

func LookupSubnetworkOutput(ctx *pulumi.Context, args LookupSubnetworkOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetworkResult, error) {
			args := v.(LookupSubnetworkArgs)
			r, err := LookupSubnetwork(ctx, &args, opts...)
			var s LookupSubnetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetworkResultOutput)
}

type LookupSubnetworkOutputArgs struct {
	Project    pulumi.StringPtrInput `pulumi:"project"`
	Region     pulumi.StringInput    `pulumi:"region"`
	Subnetwork pulumi.StringInput    `pulumi:"subnetwork"`
}

func (LookupSubnetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetworkArgs)(nil)).Elem()
}

type LookupSubnetworkResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetworkResult)(nil)).Elem()
}

func (o LookupSubnetworkResultOutput) ToLookupSubnetworkResultOutput() LookupSubnetworkResultOutput {
	return o
}

func (o LookupSubnetworkResultOutput) ToLookupSubnetworkResultOutputWithContext(ctx context.Context) LookupSubnetworkResultOutput {
	return o
}

// Can only be specified if VPC flow logging for this subnetwork is enabled. Sets the aggregation interval for collecting flow logs. Increasing the interval time reduces the amount of generated flow logs for long-lasting connections. Default is an interval of 5 seconds per connection. Valid values: INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN.
func (o LookupSubnetworkResultOutput) AggregationInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.AggregationInterval }).(pulumi.StringOutput)
}

// Whether this subnetwork's ranges can conflict with existing static routes. Setting this to true allows this subnetwork's primary and secondary ranges to overlap with (and contain) static routes that have already been configured on the corresponding network. For example if a static route has range 10.1.0.0/16, a subnet range 10.0.0.0/8 could only be created if allow_conflicting_routes=true. Overlapping is only allowed on subnetwork operations; routes whose ranges conflict with this subnetwork's ranges won't be allowed unless route.allow_conflicting_subnetworks is set to true. Typically packets destined to IPs within the subnetwork (which may contain private/sensitive data) are prevented from leaving the virtual network. Setting this field to true will disable this feature. The default value is false and applies to all existing subnetworks and automatically created subnetworks. This field cannot be set to true at resource creation time.
func (o LookupSubnetworkResultOutput) AllowSubnetCidrRoutesOverlap() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) bool { return v.AllowSubnetCidrRoutesOverlap }).(pulumi.BoolOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupSubnetworkResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
func (o LookupSubnetworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Description }).(pulumi.StringOutput)
}

// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is determined by the org policy, if there is no org policy specified, then it will default to disabled. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
func (o LookupSubnetworkResultOutput) EnableFlowLogs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) bool { return v.EnableFlowLogs }).(pulumi.BoolOutput)
}

// Enables Layer2 communication on the subnetwork.
func (o LookupSubnetworkResultOutput) EnableL2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) bool { return v.EnableL2 }).(pulumi.BoolOutput)
}

// Deprecated in favor of enable in PrivateIpv6GoogleAccess. Whether the VMs in this subnet can directly access Google services via internal IPv6 addresses. This field can be both set at resource creation time and updated using patch.
//
// Deprecated: Deprecated in favor of enable in PrivateIpv6GoogleAccess. Whether the VMs in this subnet can directly access Google services via internal IPv6 addresses. This field can be both set at resource creation time and updated using patch.
func (o LookupSubnetworkResultOutput) EnablePrivateV6Access() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) bool { return v.EnablePrivateV6Access }).(pulumi.BoolOutput)
}

// The external IPv6 address range that is owned by this subnetwork.
func (o LookupSubnetworkResultOutput) ExternalIpv6Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.ExternalIpv6Prefix }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
func (o LookupSubnetworkResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// Can only be specified if VPC flow logging for this subnetwork is enabled. The value of the field must be in [0, 1]. Set the sampling rate of VPC flow logs within the subnetwork where 1.0 means all collected logs are reported and 0.0 means no logs are reported. Default is 0.5 unless otherwise specified by the org policy, which means half of all collected logs are reported.
func (o LookupSubnetworkResultOutput) FlowSampling() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSubnetworkResult) float64 { return v.FlowSampling }).(pulumi.Float64Output)
}

// The gateway address for default routes to reach destination addresses outside this subnetwork.
func (o LookupSubnetworkResultOutput) GatewayAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.GatewayAddress }).(pulumi.StringOutput)
}

// The internal IPv6 address range that is assigned to this subnetwork.
func (o LookupSubnetworkResultOutput) InternalIpv6Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.InternalIpv6Prefix }).(pulumi.StringOutput)
}

// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
func (o LookupSubnetworkResultOutput) IpCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.IpCidrRange }).(pulumi.StringOutput)
}

// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first time the subnet is updated into IPV4_IPV6 dual stack.
func (o LookupSubnetworkResultOutput) Ipv6AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Ipv6AccessType }).(pulumi.StringOutput)
}

// This field is for internal use.
func (o LookupSubnetworkResultOutput) Ipv6CidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Ipv6CidrRange }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#subnetwork for Subnetwork resources.
func (o LookupSubnetworkResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Kind }).(pulumi.StringOutput)
}

// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
func (o LookupSubnetworkResultOutput) LogConfig() SubnetworkLogConfigResponseOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) SubnetworkLogConfigResponse { return v.LogConfig }).(SubnetworkLogConfigResponseOutput)
}

// Can only be specified if VPC flow logging for this subnetwork is enabled. Configures whether metadata fields should be added to the reported VPC flow logs. Options are INCLUDE_ALL_METADATA, EXCLUDE_ALL_METADATA, and CUSTOM_METADATA. Default is EXCLUDE_ALL_METADATA.
func (o LookupSubnetworkResultOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Metadata }).(pulumi.StringOutput)
}

// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupSubnetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. This field can be set only at resource creation time.
func (o LookupSubnetworkResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Network }).(pulumi.StringOutput)
}

// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
func (o LookupSubnetworkResultOutput) PrivateIpGoogleAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) bool { return v.PrivateIpGoogleAccess }).(pulumi.BoolOutput)
}

// This field is for internal use. This field can be both set at resource creation time and updated using patch.
func (o LookupSubnetworkResultOutput) PrivateIpv6GoogleAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.PrivateIpv6GoogleAccess }).(pulumi.StringOutput)
}

// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
func (o LookupSubnetworkResultOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Purpose }).(pulumi.StringOutput)
}

// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
func (o LookupSubnetworkResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Region }).(pulumi.StringOutput)
}

// The URL of the reserved internal range.
func (o LookupSubnetworkResultOutput) ReservedInternalRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.ReservedInternalRange }).(pulumi.StringOutput)
}

// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
func (o LookupSubnetworkResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.Role }).(pulumi.StringOutput)
}

// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
func (o LookupSubnetworkResultOutput) SecondaryIpRanges() SubnetworkSecondaryRangeResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) []SubnetworkSecondaryRangeResponse { return v.SecondaryIpRanges }).(SubnetworkSecondaryRangeResponseArrayOutput)
}

// Server-defined URL for the resource.
func (o LookupSubnetworkResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupSubnetworkResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The stack type for the subnet. If set to IPV4_ONLY, new VMs in the subnet are assigned IPv4 addresses only. If set to IPV4_IPV6, new VMs in the subnet can be assigned both IPv4 and IPv6 addresses. If not specified, IPV4_ONLY is used. This field can be both set at resource creation time and updated using patch.
func (o LookupSubnetworkResultOutput) StackType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.StackType }).(pulumi.StringOutput)
}

// The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY
func (o LookupSubnetworkResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) string { return v.State }).(pulumi.StringOutput)
}

// A repeated field indicating the VLAN IDs supported on this subnetwork. During Subnet creation, specifying vlan is valid only if enable_l2 is true. During Subnet Update, specifying vlan is allowed only for l2 enabled subnets. Restricted to only one VLAN.
func (o LookupSubnetworkResultOutput) Vlans() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupSubnetworkResult) []int { return v.Vlans }).(pulumi.IntArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetworkResultOutput{})
}
