// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified device.
func LookupDevice(ctx *pulumi.Context, args *LookupDeviceArgs, opts ...pulumi.InvokeOption) (*LookupDeviceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeviceResult
	err := ctx.Invoke("google-native:cloudidentity/v1:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceArgs struct {
	Customer *string `pulumi:"customer"`
	DeviceId string  `pulumi:"deviceId"`
}

type LookupDeviceResult struct {
	// Attributes specific to Android devices.
	AndroidSpecificAttributes GoogleAppsCloudidentityDevicesV1AndroidAttributesResponse `pulumi:"androidSpecificAttributes"`
	// Asset tag of the device.
	AssetTag string `pulumi:"assetTag"`
	// Baseband version of the device.
	BasebandVersion string `pulumi:"basebandVersion"`
	// Device bootloader version. Example: 0.6.7.
	BootloaderVersion string `pulumi:"bootloaderVersion"`
	// Device brand. Example: Samsung.
	Brand string `pulumi:"brand"`
	// Build number of the device.
	BuildNumber string `pulumi:"buildNumber"`
	// Represents whether the Device is compromised.
	CompromisedState string `pulumi:"compromisedState"`
	// When the Company-Owned device was imported. This field is empty for BYOD devices.
	CreateTime string `pulumi:"createTime"`
	// Unique identifier for the device.
	DeviceId string `pulumi:"deviceId"`
	// Type of device.
	DeviceType string `pulumi:"deviceType"`
	// Whether developer options is enabled on device.
	EnabledDeveloperOptions bool `pulumi:"enabledDeveloperOptions"`
	// Whether USB debugging is enabled on device.
	EnabledUsbDebugging bool `pulumi:"enabledUsbDebugging"`
	// Device encryption state.
	EncryptionState string `pulumi:"encryptionState"`
	// IMEI number of device if GSM device; empty otherwise.
	Imei string `pulumi:"imei"`
	// Kernel version of the device.
	KernelVersion string `pulumi:"kernelVersion"`
	// Most recent time when device synced with this service.
	LastSyncTime string `pulumi:"lastSyncTime"`
	// Management state of the device
	ManagementState string `pulumi:"managementState"`
	// Device manufacturer. Example: Motorola.
	Manufacturer string `pulumi:"manufacturer"`
	// MEID number of device if CDMA device; empty otherwise.
	Meid string `pulumi:"meid"`
	// Model name of device. Example: Pixel 3.
	Model string `pulumi:"model"`
	// [Resource name](https://cloud.google.com/apis/design/resource_names) of the Device in format: `devices/{device}`, where device is the unique id assigned to the Device.
	Name string `pulumi:"name"`
	// Mobile or network operator of device, if available.
	NetworkOperator string `pulumi:"networkOperator"`
	// OS version of the device. Example: Android 8.1.0.
	OsVersion string `pulumi:"osVersion"`
	// Domain name for Google accounts on device. Type for other accounts on device. On Android, will only be populated if |ownership_privilege| is |PROFILE_OWNER| or |DEVICE_OWNER|. Does not include the account signed in to the device policy app if that account's domain has only one account. Examples: "com.example", "xyz.com".
	OtherAccounts []string `pulumi:"otherAccounts"`
	// Whether the device is owned by the company or an individual
	OwnerType string `pulumi:"ownerType"`
	// OS release version. Example: 6.0.
	ReleaseVersion string `pulumi:"releaseVersion"`
	// OS security patch update time on device.
	SecurityPatchTime string `pulumi:"securityPatchTime"`
	// Serial Number of device. Example: HT82V1A01076.
	SerialNumber string `pulumi:"serialNumber"`
	// WiFi MAC addresses of device.
	WifiMacAddresses []string `pulumi:"wifiMacAddresses"`
}

func LookupDeviceOutput(ctx *pulumi.Context, args LookupDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceResult, error) {
			args := v.(LookupDeviceArgs)
			r, err := LookupDevice(ctx, &args, opts...)
			var s LookupDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceResultOutput)
}

type LookupDeviceOutputArgs struct {
	Customer pulumi.StringPtrInput `pulumi:"customer"`
	DeviceId pulumi.StringInput    `pulumi:"deviceId"`
}

func (LookupDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceArgs)(nil)).Elem()
}

type LookupDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceResult)(nil)).Elem()
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutput() LookupDeviceResultOutput {
	return o
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutputWithContext(ctx context.Context) LookupDeviceResultOutput {
	return o
}

// Attributes specific to Android devices.
func (o LookupDeviceResultOutput) AndroidSpecificAttributes() GoogleAppsCloudidentityDevicesV1AndroidAttributesResponseOutput {
	return o.ApplyT(func(v LookupDeviceResult) GoogleAppsCloudidentityDevicesV1AndroidAttributesResponse {
		return v.AndroidSpecificAttributes
	}).(GoogleAppsCloudidentityDevicesV1AndroidAttributesResponseOutput)
}

// Asset tag of the device.
func (o LookupDeviceResultOutput) AssetTag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.AssetTag }).(pulumi.StringOutput)
}

// Baseband version of the device.
func (o LookupDeviceResultOutput) BasebandVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.BasebandVersion }).(pulumi.StringOutput)
}

// Device bootloader version. Example: 0.6.7.
func (o LookupDeviceResultOutput) BootloaderVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.BootloaderVersion }).(pulumi.StringOutput)
}

// Device brand. Example: Samsung.
func (o LookupDeviceResultOutput) Brand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Brand }).(pulumi.StringOutput)
}

// Build number of the device.
func (o LookupDeviceResultOutput) BuildNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.BuildNumber }).(pulumi.StringOutput)
}

// Represents whether the Device is compromised.
func (o LookupDeviceResultOutput) CompromisedState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.CompromisedState }).(pulumi.StringOutput)
}

// When the Company-Owned device was imported. This field is empty for BYOD devices.
func (o LookupDeviceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Unique identifier for the device.
func (o LookupDeviceResultOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.DeviceId }).(pulumi.StringOutput)
}

// Type of device.
func (o LookupDeviceResultOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.DeviceType }).(pulumi.StringOutput)
}

// Whether developer options is enabled on device.
func (o LookupDeviceResultOutput) EnabledDeveloperOptions() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDeviceResult) bool { return v.EnabledDeveloperOptions }).(pulumi.BoolOutput)
}

// Whether USB debugging is enabled on device.
func (o LookupDeviceResultOutput) EnabledUsbDebugging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDeviceResult) bool { return v.EnabledUsbDebugging }).(pulumi.BoolOutput)
}

// Device encryption state.
func (o LookupDeviceResultOutput) EncryptionState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.EncryptionState }).(pulumi.StringOutput)
}

// IMEI number of device if GSM device; empty otherwise.
func (o LookupDeviceResultOutput) Imei() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Imei }).(pulumi.StringOutput)
}

// Kernel version of the device.
func (o LookupDeviceResultOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.KernelVersion }).(pulumi.StringOutput)
}

// Most recent time when device synced with this service.
func (o LookupDeviceResultOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

// Management state of the device
func (o LookupDeviceResultOutput) ManagementState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.ManagementState }).(pulumi.StringOutput)
}

// Device manufacturer. Example: Motorola.
func (o LookupDeviceResultOutput) Manufacturer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Manufacturer }).(pulumi.StringOutput)
}

// MEID number of device if CDMA device; empty otherwise.
func (o LookupDeviceResultOutput) Meid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Meid }).(pulumi.StringOutput)
}

// Model name of device. Example: Pixel 3.
func (o LookupDeviceResultOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Model }).(pulumi.StringOutput)
}

// [Resource name](https://cloud.google.com/apis/design/resource_names) of the Device in format: `devices/{device}`, where device is the unique id assigned to the Device.
func (o LookupDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Mobile or network operator of device, if available.
func (o LookupDeviceResultOutput) NetworkOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.NetworkOperator }).(pulumi.StringOutput)
}

// OS version of the device. Example: Android 8.1.0.
func (o LookupDeviceResultOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.OsVersion }).(pulumi.StringOutput)
}

// Domain name for Google accounts on device. Type for other accounts on device. On Android, will only be populated if |ownership_privilege| is |PROFILE_OWNER| or |DEVICE_OWNER|. Does not include the account signed in to the device policy app if that account's domain has only one account. Examples: "com.example", "xyz.com".
func (o LookupDeviceResultOutput) OtherAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []string { return v.OtherAccounts }).(pulumi.StringArrayOutput)
}

// Whether the device is owned by the company or an individual
func (o LookupDeviceResultOutput) OwnerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.OwnerType }).(pulumi.StringOutput)
}

// OS release version. Example: 6.0.
func (o LookupDeviceResultOutput) ReleaseVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.ReleaseVersion }).(pulumi.StringOutput)
}

// OS security patch update time on device.
func (o LookupDeviceResultOutput) SecurityPatchTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.SecurityPatchTime }).(pulumi.StringOutput)
}

// Serial Number of device. Example: HT82V1A01076.
func (o LookupDeviceResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

// WiFi MAC addresses of device.
func (o LookupDeviceResultOutput) WifiMacAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []string { return v.WifiMacAddresses }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceResultOutput{})
}
