// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

// File share configuration for the instance.
type FileShareConfig struct {
	// File share capacity in gigabytes (GB). Filestore defines 1 GB as 1024^3 bytes.
	CapacityGb *string `pulumi:"capacityGb"`
	// The name of the file share (must be 16 characters or less).
	Name *string `pulumi:"name"`
	// Nfs Export Options. There is a limit of 10 export options per file share.
	NfsExportOptions []NfsExportOptions `pulumi:"nfsExportOptions"`
	// The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`, that this file share has been restored from.
	SourceBackup *string `pulumi:"sourceBackup"`
}

// FileShareConfigInput is an input type that accepts FileShareConfigArgs and FileShareConfigOutput values.
// You can construct a concrete instance of `FileShareConfigInput` via:
//
//	FileShareConfigArgs{...}
type FileShareConfigInput interface {
	pulumi.Input

	ToFileShareConfigOutput() FileShareConfigOutput
	ToFileShareConfigOutputWithContext(context.Context) FileShareConfigOutput
}

// File share configuration for the instance.
type FileShareConfigArgs struct {
	// File share capacity in gigabytes (GB). Filestore defines 1 GB as 1024^3 bytes.
	CapacityGb pulumi.StringPtrInput `pulumi:"capacityGb"`
	// The name of the file share (must be 16 characters or less).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Nfs Export Options. There is a limit of 10 export options per file share.
	NfsExportOptions NfsExportOptionsArrayInput `pulumi:"nfsExportOptions"`
	// The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`, that this file share has been restored from.
	SourceBackup pulumi.StringPtrInput `pulumi:"sourceBackup"`
}

func (FileShareConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileShareConfig)(nil)).Elem()
}

func (i FileShareConfigArgs) ToFileShareConfigOutput() FileShareConfigOutput {
	return i.ToFileShareConfigOutputWithContext(context.Background())
}

func (i FileShareConfigArgs) ToFileShareConfigOutputWithContext(ctx context.Context) FileShareConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileShareConfigOutput)
}

// FileShareConfigArrayInput is an input type that accepts FileShareConfigArray and FileShareConfigArrayOutput values.
// You can construct a concrete instance of `FileShareConfigArrayInput` via:
//
//	FileShareConfigArray{ FileShareConfigArgs{...} }
type FileShareConfigArrayInput interface {
	pulumi.Input

	ToFileShareConfigArrayOutput() FileShareConfigArrayOutput
	ToFileShareConfigArrayOutputWithContext(context.Context) FileShareConfigArrayOutput
}

type FileShareConfigArray []FileShareConfigInput

func (FileShareConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FileShareConfig)(nil)).Elem()
}

func (i FileShareConfigArray) ToFileShareConfigArrayOutput() FileShareConfigArrayOutput {
	return i.ToFileShareConfigArrayOutputWithContext(context.Background())
}

func (i FileShareConfigArray) ToFileShareConfigArrayOutputWithContext(ctx context.Context) FileShareConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileShareConfigArrayOutput)
}

// File share configuration for the instance.
type FileShareConfigOutput struct{ *pulumi.OutputState }

func (FileShareConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileShareConfig)(nil)).Elem()
}

func (o FileShareConfigOutput) ToFileShareConfigOutput() FileShareConfigOutput {
	return o
}

func (o FileShareConfigOutput) ToFileShareConfigOutputWithContext(ctx context.Context) FileShareConfigOutput {
	return o
}

// File share capacity in gigabytes (GB). Filestore defines 1 GB as 1024^3 bytes.
func (o FileShareConfigOutput) CapacityGb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileShareConfig) *string { return v.CapacityGb }).(pulumi.StringPtrOutput)
}

// The name of the file share (must be 16 characters or less).
func (o FileShareConfigOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileShareConfig) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Nfs Export Options. There is a limit of 10 export options per file share.
func (o FileShareConfigOutput) NfsExportOptions() NfsExportOptionsArrayOutput {
	return o.ApplyT(func(v FileShareConfig) []NfsExportOptions { return v.NfsExportOptions }).(NfsExportOptionsArrayOutput)
}

// The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`, that this file share has been restored from.
func (o FileShareConfigOutput) SourceBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileShareConfig) *string { return v.SourceBackup }).(pulumi.StringPtrOutput)
}

type FileShareConfigArrayOutput struct{ *pulumi.OutputState }

func (FileShareConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FileShareConfig)(nil)).Elem()
}

func (o FileShareConfigArrayOutput) ToFileShareConfigArrayOutput() FileShareConfigArrayOutput {
	return o
}

func (o FileShareConfigArrayOutput) ToFileShareConfigArrayOutputWithContext(ctx context.Context) FileShareConfigArrayOutput {
	return o
}

func (o FileShareConfigArrayOutput) Index(i pulumi.IntInput) FileShareConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FileShareConfig {
		return vs[0].([]FileShareConfig)[vs[1].(int)]
	}).(FileShareConfigOutput)
}

// File share configuration for the instance.
type FileShareConfigResponse struct {
	// File share capacity in gigabytes (GB). Filestore defines 1 GB as 1024^3 bytes.
	CapacityGb string `pulumi:"capacityGb"`
	// The name of the file share (must be 16 characters or less).
	Name string `pulumi:"name"`
	// Nfs Export Options. There is a limit of 10 export options per file share.
	NfsExportOptions []NfsExportOptionsResponse `pulumi:"nfsExportOptions"`
	// The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`, that this file share has been restored from.
	SourceBackup string `pulumi:"sourceBackup"`
}

// File share configuration for the instance.
type FileShareConfigResponseOutput struct{ *pulumi.OutputState }

func (FileShareConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileShareConfigResponse)(nil)).Elem()
}

func (o FileShareConfigResponseOutput) ToFileShareConfigResponseOutput() FileShareConfigResponseOutput {
	return o
}

func (o FileShareConfigResponseOutput) ToFileShareConfigResponseOutputWithContext(ctx context.Context) FileShareConfigResponseOutput {
	return o
}

// File share capacity in gigabytes (GB). Filestore defines 1 GB as 1024^3 bytes.
func (o FileShareConfigResponseOutput) CapacityGb() pulumi.StringOutput {
	return o.ApplyT(func(v FileShareConfigResponse) string { return v.CapacityGb }).(pulumi.StringOutput)
}

// The name of the file share (must be 16 characters or less).
func (o FileShareConfigResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v FileShareConfigResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Nfs Export Options. There is a limit of 10 export options per file share.
func (o FileShareConfigResponseOutput) NfsExportOptions() NfsExportOptionsResponseArrayOutput {
	return o.ApplyT(func(v FileShareConfigResponse) []NfsExportOptionsResponse { return v.NfsExportOptions }).(NfsExportOptionsResponseArrayOutput)
}

// The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`, that this file share has been restored from.
func (o FileShareConfigResponseOutput) SourceBackup() pulumi.StringOutput {
	return o.ApplyT(func(v FileShareConfigResponse) string { return v.SourceBackup }).(pulumi.StringOutput)
}

type FileShareConfigResponseArrayOutput struct{ *pulumi.OutputState }

func (FileShareConfigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FileShareConfigResponse)(nil)).Elem()
}

func (o FileShareConfigResponseArrayOutput) ToFileShareConfigResponseArrayOutput() FileShareConfigResponseArrayOutput {
	return o
}

func (o FileShareConfigResponseArrayOutput) ToFileShareConfigResponseArrayOutputWithContext(ctx context.Context) FileShareConfigResponseArrayOutput {
	return o
}

func (o FileShareConfigResponseArrayOutput) Index(i pulumi.IntInput) FileShareConfigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FileShareConfigResponse {
		return vs[0].([]FileShareConfigResponse)[vs[1].(int)]
	}).(FileShareConfigResponseOutput)
}

// Network configuration for the instance.
type NetworkConfig struct {
	// The network connect mode of the Filestore instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode *NetworkConfigConnectMode `pulumi:"connectMode"`
	// Internet protocol versions for which the instance has IP addresses assigned. For this version, only MODE_IPV4 is supported.
	Modes []NetworkConfigModesItem `pulumi:"modes"`
	// The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected.
	Network *string `pulumi:"network"`
	// Optional, reserved_ip_range can have one of the following two types of values. * CIDR range value when using DIRECT_PEERING connect mode. * [Allocated IP address range](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address) when using PRIVATE_SERVICE_ACCESS connect mode. When the name of an allocated IP address range is specified, it must be one of the ranges associated with the private service access connection. When specified as a direct CIDR value, it must be a /29 CIDR block for Basic tier, a /24 CIDR block for High Scale tier, or a /26 CIDR block for Enterprise tier in one of the [internal IP address ranges](https://www.arin.net/reference/research/statistics/address_filters/) that identifies the range of IP addresses reserved for this instance. For example, 10.0.0.0/29, 192.168.0.0/24 or 192.168.0.0/26, respectively. The range you specify can't overlap with either existing subnets or assigned IP address ranges for other Filestore instances in the selected VPC network.
	ReservedIpRange *string `pulumi:"reservedIpRange"`
}

// NetworkConfigInput is an input type that accepts NetworkConfigArgs and NetworkConfigOutput values.
// You can construct a concrete instance of `NetworkConfigInput` via:
//
//	NetworkConfigArgs{...}
type NetworkConfigInput interface {
	pulumi.Input

	ToNetworkConfigOutput() NetworkConfigOutput
	ToNetworkConfigOutputWithContext(context.Context) NetworkConfigOutput
}

// Network configuration for the instance.
type NetworkConfigArgs struct {
	// The network connect mode of the Filestore instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode NetworkConfigConnectModePtrInput `pulumi:"connectMode"`
	// Internet protocol versions for which the instance has IP addresses assigned. For this version, only MODE_IPV4 is supported.
	Modes NetworkConfigModesItemArrayInput `pulumi:"modes"`
	// The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected.
	Network pulumi.StringPtrInput `pulumi:"network"`
	// Optional, reserved_ip_range can have one of the following two types of values. * CIDR range value when using DIRECT_PEERING connect mode. * [Allocated IP address range](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address) when using PRIVATE_SERVICE_ACCESS connect mode. When the name of an allocated IP address range is specified, it must be one of the ranges associated with the private service access connection. When specified as a direct CIDR value, it must be a /29 CIDR block for Basic tier, a /24 CIDR block for High Scale tier, or a /26 CIDR block for Enterprise tier in one of the [internal IP address ranges](https://www.arin.net/reference/research/statistics/address_filters/) that identifies the range of IP addresses reserved for this instance. For example, 10.0.0.0/29, 192.168.0.0/24 or 192.168.0.0/26, respectively. The range you specify can't overlap with either existing subnets or assigned IP address ranges for other Filestore instances in the selected VPC network.
	ReservedIpRange pulumi.StringPtrInput `pulumi:"reservedIpRange"`
}

func (NetworkConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfig)(nil)).Elem()
}

func (i NetworkConfigArgs) ToNetworkConfigOutput() NetworkConfigOutput {
	return i.ToNetworkConfigOutputWithContext(context.Background())
}

func (i NetworkConfigArgs) ToNetworkConfigOutputWithContext(ctx context.Context) NetworkConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigOutput)
}

// NetworkConfigArrayInput is an input type that accepts NetworkConfigArray and NetworkConfigArrayOutput values.
// You can construct a concrete instance of `NetworkConfigArrayInput` via:
//
//	NetworkConfigArray{ NetworkConfigArgs{...} }
type NetworkConfigArrayInput interface {
	pulumi.Input

	ToNetworkConfigArrayOutput() NetworkConfigArrayOutput
	ToNetworkConfigArrayOutputWithContext(context.Context) NetworkConfigArrayOutput
}

type NetworkConfigArray []NetworkConfigInput

func (NetworkConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkConfig)(nil)).Elem()
}

func (i NetworkConfigArray) ToNetworkConfigArrayOutput() NetworkConfigArrayOutput {
	return i.ToNetworkConfigArrayOutputWithContext(context.Background())
}

func (i NetworkConfigArray) ToNetworkConfigArrayOutputWithContext(ctx context.Context) NetworkConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigArrayOutput)
}

// Network configuration for the instance.
type NetworkConfigOutput struct{ *pulumi.OutputState }

func (NetworkConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfig)(nil)).Elem()
}

func (o NetworkConfigOutput) ToNetworkConfigOutput() NetworkConfigOutput {
	return o
}

func (o NetworkConfigOutput) ToNetworkConfigOutputWithContext(ctx context.Context) NetworkConfigOutput {
	return o
}

// The network connect mode of the Filestore instance. If not provided, the connect mode defaults to DIRECT_PEERING.
func (o NetworkConfigOutput) ConnectMode() NetworkConfigConnectModePtrOutput {
	return o.ApplyT(func(v NetworkConfig) *NetworkConfigConnectMode { return v.ConnectMode }).(NetworkConfigConnectModePtrOutput)
}

// Internet protocol versions for which the instance has IP addresses assigned. For this version, only MODE_IPV4 is supported.
func (o NetworkConfigOutput) Modes() NetworkConfigModesItemArrayOutput {
	return o.ApplyT(func(v NetworkConfig) []NetworkConfigModesItem { return v.Modes }).(NetworkConfigModesItemArrayOutput)
}

// The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected.
func (o NetworkConfigOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkConfig) *string { return v.Network }).(pulumi.StringPtrOutput)
}

// Optional, reserved_ip_range can have one of the following two types of values. * CIDR range value when using DIRECT_PEERING connect mode. * [Allocated IP address range](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address) when using PRIVATE_SERVICE_ACCESS connect mode. When the name of an allocated IP address range is specified, it must be one of the ranges associated with the private service access connection. When specified as a direct CIDR value, it must be a /29 CIDR block for Basic tier, a /24 CIDR block for High Scale tier, or a /26 CIDR block for Enterprise tier in one of the [internal IP address ranges](https://www.arin.net/reference/research/statistics/address_filters/) that identifies the range of IP addresses reserved for this instance. For example, 10.0.0.0/29, 192.168.0.0/24 or 192.168.0.0/26, respectively. The range you specify can't overlap with either existing subnets or assigned IP address ranges for other Filestore instances in the selected VPC network.
func (o NetworkConfigOutput) ReservedIpRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkConfig) *string { return v.ReservedIpRange }).(pulumi.StringPtrOutput)
}

type NetworkConfigArrayOutput struct{ *pulumi.OutputState }

func (NetworkConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkConfig)(nil)).Elem()
}

func (o NetworkConfigArrayOutput) ToNetworkConfigArrayOutput() NetworkConfigArrayOutput {
	return o
}

func (o NetworkConfigArrayOutput) ToNetworkConfigArrayOutputWithContext(ctx context.Context) NetworkConfigArrayOutput {
	return o
}

func (o NetworkConfigArrayOutput) Index(i pulumi.IntInput) NetworkConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkConfig {
		return vs[0].([]NetworkConfig)[vs[1].(int)]
	}).(NetworkConfigOutput)
}

// Network configuration for the instance.
type NetworkConfigResponse struct {
	// The network connect mode of the Filestore instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode string `pulumi:"connectMode"`
	// IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or IPv6 addresses in the format `{block1}:{block2}:{block3}:{block4}:{block5}:{block6}:{block7}:{block8}`.
	IpAddresses []string `pulumi:"ipAddresses"`
	// Internet protocol versions for which the instance has IP addresses assigned. For this version, only MODE_IPV4 is supported.
	Modes []string `pulumi:"modes"`
	// The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected.
	Network string `pulumi:"network"`
	// Optional, reserved_ip_range can have one of the following two types of values. * CIDR range value when using DIRECT_PEERING connect mode. * [Allocated IP address range](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address) when using PRIVATE_SERVICE_ACCESS connect mode. When the name of an allocated IP address range is specified, it must be one of the ranges associated with the private service access connection. When specified as a direct CIDR value, it must be a /29 CIDR block for Basic tier, a /24 CIDR block for High Scale tier, or a /26 CIDR block for Enterprise tier in one of the [internal IP address ranges](https://www.arin.net/reference/research/statistics/address_filters/) that identifies the range of IP addresses reserved for this instance. For example, 10.0.0.0/29, 192.168.0.0/24 or 192.168.0.0/26, respectively. The range you specify can't overlap with either existing subnets or assigned IP address ranges for other Filestore instances in the selected VPC network.
	ReservedIpRange string `pulumi:"reservedIpRange"`
}

// Network configuration for the instance.
type NetworkConfigResponseOutput struct{ *pulumi.OutputState }

func (NetworkConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfigResponse)(nil)).Elem()
}

func (o NetworkConfigResponseOutput) ToNetworkConfigResponseOutput() NetworkConfigResponseOutput {
	return o
}

func (o NetworkConfigResponseOutput) ToNetworkConfigResponseOutputWithContext(ctx context.Context) NetworkConfigResponseOutput {
	return o
}

// The network connect mode of the Filestore instance. If not provided, the connect mode defaults to DIRECT_PEERING.
func (o NetworkConfigResponseOutput) ConnectMode() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkConfigResponse) string { return v.ConnectMode }).(pulumi.StringOutput)
}

// IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or IPv6 addresses in the format `{block1}:{block2}:{block3}:{block4}:{block5}:{block6}:{block7}:{block8}`.
func (o NetworkConfigResponseOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkConfigResponse) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Internet protocol versions for which the instance has IP addresses assigned. For this version, only MODE_IPV4 is supported.
func (o NetworkConfigResponseOutput) Modes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkConfigResponse) []string { return v.Modes }).(pulumi.StringArrayOutput)
}

// The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected.
func (o NetworkConfigResponseOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkConfigResponse) string { return v.Network }).(pulumi.StringOutput)
}

// Optional, reserved_ip_range can have one of the following two types of values. * CIDR range value when using DIRECT_PEERING connect mode. * [Allocated IP address range](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address) when using PRIVATE_SERVICE_ACCESS connect mode. When the name of an allocated IP address range is specified, it must be one of the ranges associated with the private service access connection. When specified as a direct CIDR value, it must be a /29 CIDR block for Basic tier, a /24 CIDR block for High Scale tier, or a /26 CIDR block for Enterprise tier in one of the [internal IP address ranges](https://www.arin.net/reference/research/statistics/address_filters/) that identifies the range of IP addresses reserved for this instance. For example, 10.0.0.0/29, 192.168.0.0/24 or 192.168.0.0/26, respectively. The range you specify can't overlap with either existing subnets or assigned IP address ranges for other Filestore instances in the selected VPC network.
func (o NetworkConfigResponseOutput) ReservedIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkConfigResponse) string { return v.ReservedIpRange }).(pulumi.StringOutput)
}

type NetworkConfigResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkConfigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkConfigResponse)(nil)).Elem()
}

func (o NetworkConfigResponseArrayOutput) ToNetworkConfigResponseArrayOutput() NetworkConfigResponseArrayOutput {
	return o
}

func (o NetworkConfigResponseArrayOutput) ToNetworkConfigResponseArrayOutputWithContext(ctx context.Context) NetworkConfigResponseArrayOutput {
	return o
}

func (o NetworkConfigResponseArrayOutput) Index(i pulumi.IntInput) NetworkConfigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkConfigResponse {
		return vs[0].([]NetworkConfigResponse)[vs[1].(int)]
	}).(NetworkConfigResponseOutput)
}

// NFS export options specifications.
type NfsExportOptions struct {
	// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
	AccessMode *NfsExportOptionsAccessMode `pulumi:"accessMode"`
	// An integer representing the anonymous group id with a default value of 65534. Anon_gid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
	AnonGid *string `pulumi:"anonGid"`
	// An integer representing the anonymous user id with a default value of 65534. Anon_uid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
	AnonUid *string `pulumi:"anonUid"`
	// List of either an IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the file share. Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned. The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
	IpRanges []string `pulumi:"ipRanges"`
	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
	SquashMode *NfsExportOptionsSquashMode `pulumi:"squashMode"`
}

// NfsExportOptionsInput is an input type that accepts NfsExportOptionsArgs and NfsExportOptionsOutput values.
// You can construct a concrete instance of `NfsExportOptionsInput` via:
//
//	NfsExportOptionsArgs{...}
type NfsExportOptionsInput interface {
	pulumi.Input

	ToNfsExportOptionsOutput() NfsExportOptionsOutput
	ToNfsExportOptionsOutputWithContext(context.Context) NfsExportOptionsOutput
}

// NFS export options specifications.
type NfsExportOptionsArgs struct {
	// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
	AccessMode NfsExportOptionsAccessModePtrInput `pulumi:"accessMode"`
	// An integer representing the anonymous group id with a default value of 65534. Anon_gid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
	AnonGid pulumi.StringPtrInput `pulumi:"anonGid"`
	// An integer representing the anonymous user id with a default value of 65534. Anon_uid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
	AnonUid pulumi.StringPtrInput `pulumi:"anonUid"`
	// List of either an IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the file share. Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned. The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
	IpRanges pulumi.StringArrayInput `pulumi:"ipRanges"`
	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
	SquashMode NfsExportOptionsSquashModePtrInput `pulumi:"squashMode"`
}

func (NfsExportOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsExportOptions)(nil)).Elem()
}

func (i NfsExportOptionsArgs) ToNfsExportOptionsOutput() NfsExportOptionsOutput {
	return i.ToNfsExportOptionsOutputWithContext(context.Background())
}

func (i NfsExportOptionsArgs) ToNfsExportOptionsOutputWithContext(ctx context.Context) NfsExportOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsExportOptionsOutput)
}

// NfsExportOptionsArrayInput is an input type that accepts NfsExportOptionsArray and NfsExportOptionsArrayOutput values.
// You can construct a concrete instance of `NfsExportOptionsArrayInput` via:
//
//	NfsExportOptionsArray{ NfsExportOptionsArgs{...} }
type NfsExportOptionsArrayInput interface {
	pulumi.Input

	ToNfsExportOptionsArrayOutput() NfsExportOptionsArrayOutput
	ToNfsExportOptionsArrayOutputWithContext(context.Context) NfsExportOptionsArrayOutput
}

type NfsExportOptionsArray []NfsExportOptionsInput

func (NfsExportOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsExportOptions)(nil)).Elem()
}

func (i NfsExportOptionsArray) ToNfsExportOptionsArrayOutput() NfsExportOptionsArrayOutput {
	return i.ToNfsExportOptionsArrayOutputWithContext(context.Background())
}

func (i NfsExportOptionsArray) ToNfsExportOptionsArrayOutputWithContext(ctx context.Context) NfsExportOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsExportOptionsArrayOutput)
}

// NFS export options specifications.
type NfsExportOptionsOutput struct{ *pulumi.OutputState }

func (NfsExportOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsExportOptions)(nil)).Elem()
}

func (o NfsExportOptionsOutput) ToNfsExportOptionsOutput() NfsExportOptionsOutput {
	return o
}

func (o NfsExportOptionsOutput) ToNfsExportOptionsOutputWithContext(ctx context.Context) NfsExportOptionsOutput {
	return o
}

// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
func (o NfsExportOptionsOutput) AccessMode() NfsExportOptionsAccessModePtrOutput {
	return o.ApplyT(func(v NfsExportOptions) *NfsExportOptionsAccessMode { return v.AccessMode }).(NfsExportOptionsAccessModePtrOutput)
}

// An integer representing the anonymous group id with a default value of 65534. Anon_gid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
func (o NfsExportOptionsOutput) AnonGid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NfsExportOptions) *string { return v.AnonGid }).(pulumi.StringPtrOutput)
}

// An integer representing the anonymous user id with a default value of 65534. Anon_uid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
func (o NfsExportOptionsOutput) AnonUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NfsExportOptions) *string { return v.AnonUid }).(pulumi.StringPtrOutput)
}

// List of either an IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the file share. Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned. The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
func (o NfsExportOptionsOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NfsExportOptions) []string { return v.IpRanges }).(pulumi.StringArrayOutput)
}

// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
func (o NfsExportOptionsOutput) SquashMode() NfsExportOptionsSquashModePtrOutput {
	return o.ApplyT(func(v NfsExportOptions) *NfsExportOptionsSquashMode { return v.SquashMode }).(NfsExportOptionsSquashModePtrOutput)
}

type NfsExportOptionsArrayOutput struct{ *pulumi.OutputState }

func (NfsExportOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsExportOptions)(nil)).Elem()
}

func (o NfsExportOptionsArrayOutput) ToNfsExportOptionsArrayOutput() NfsExportOptionsArrayOutput {
	return o
}

func (o NfsExportOptionsArrayOutput) ToNfsExportOptionsArrayOutputWithContext(ctx context.Context) NfsExportOptionsArrayOutput {
	return o
}

func (o NfsExportOptionsArrayOutput) Index(i pulumi.IntInput) NfsExportOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NfsExportOptions {
		return vs[0].([]NfsExportOptions)[vs[1].(int)]
	}).(NfsExportOptionsOutput)
}

// NFS export options specifications.
type NfsExportOptionsResponse struct {
	// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
	AccessMode string `pulumi:"accessMode"`
	// An integer representing the anonymous group id with a default value of 65534. Anon_gid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
	AnonGid string `pulumi:"anonGid"`
	// An integer representing the anonymous user id with a default value of 65534. Anon_uid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
	AnonUid string `pulumi:"anonUid"`
	// List of either an IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the file share. Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned. The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
	IpRanges []string `pulumi:"ipRanges"`
	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
	SquashMode string `pulumi:"squashMode"`
}

// NFS export options specifications.
type NfsExportOptionsResponseOutput struct{ *pulumi.OutputState }

func (NfsExportOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsExportOptionsResponse)(nil)).Elem()
}

func (o NfsExportOptionsResponseOutput) ToNfsExportOptionsResponseOutput() NfsExportOptionsResponseOutput {
	return o
}

func (o NfsExportOptionsResponseOutput) ToNfsExportOptionsResponseOutputWithContext(ctx context.Context) NfsExportOptionsResponseOutput {
	return o
}

// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
func (o NfsExportOptionsResponseOutput) AccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v NfsExportOptionsResponse) string { return v.AccessMode }).(pulumi.StringOutput)
}

// An integer representing the anonymous group id with a default value of 65534. Anon_gid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
func (o NfsExportOptionsResponseOutput) AnonGid() pulumi.StringOutput {
	return o.ApplyT(func(v NfsExportOptionsResponse) string { return v.AnonGid }).(pulumi.StringOutput)
}

// An integer representing the anonymous user id with a default value of 65534. Anon_uid may only be set with squash_mode of ROOT_SQUASH. An error will be returned if this field is specified for other squash_mode settings.
func (o NfsExportOptionsResponseOutput) AnonUid() pulumi.StringOutput {
	return o.ApplyT(func(v NfsExportOptionsResponse) string { return v.AnonUid }).(pulumi.StringOutput)
}

// List of either an IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which may mount the file share. Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned. The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
func (o NfsExportOptionsResponseOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NfsExportOptionsResponse) []string { return v.IpRanges }).(pulumi.StringArrayOutput)
}

// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
func (o NfsExportOptionsResponseOutput) SquashMode() pulumi.StringOutput {
	return o.ApplyT(func(v NfsExportOptionsResponse) string { return v.SquashMode }).(pulumi.StringOutput)
}

type NfsExportOptionsResponseArrayOutput struct{ *pulumi.OutputState }

func (NfsExportOptionsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NfsExportOptionsResponse)(nil)).Elem()
}

func (o NfsExportOptionsResponseArrayOutput) ToNfsExportOptionsResponseArrayOutput() NfsExportOptionsResponseArrayOutput {
	return o
}

func (o NfsExportOptionsResponseArrayOutput) ToNfsExportOptionsResponseArrayOutputWithContext(ctx context.Context) NfsExportOptionsResponseArrayOutput {
	return o
}

func (o NfsExportOptionsResponseArrayOutput) Index(i pulumi.IntInput) NfsExportOptionsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NfsExportOptionsResponse {
		return vs[0].([]NfsExportOptionsResponse)[vs[1].(int)]
	}).(NfsExportOptionsResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileShareConfigInput)(nil)).Elem(), FileShareConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileShareConfigArrayInput)(nil)).Elem(), FileShareConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConfigInput)(nil)).Elem(), NetworkConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConfigArrayInput)(nil)).Elem(), NetworkConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NfsExportOptionsInput)(nil)).Elem(), NfsExportOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NfsExportOptionsArrayInput)(nil)).Elem(), NfsExportOptionsArray{})
	pulumi.RegisterOutputType(FileShareConfigOutput{})
	pulumi.RegisterOutputType(FileShareConfigArrayOutput{})
	pulumi.RegisterOutputType(FileShareConfigResponseOutput{})
	pulumi.RegisterOutputType(FileShareConfigResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkConfigOutput{})
	pulumi.RegisterOutputType(NetworkConfigArrayOutput{})
	pulumi.RegisterOutputType(NetworkConfigResponseOutput{})
	pulumi.RegisterOutputType(NetworkConfigResponseArrayOutput{})
	pulumi.RegisterOutputType(NfsExportOptionsOutput{})
	pulumi.RegisterOutputType(NfsExportOptionsArrayOutput{})
	pulumi.RegisterOutputType(NfsExportOptionsResponseOutput{})
	pulumi.RegisterOutputType(NfsExportOptionsResponseArrayOutput{})
}
