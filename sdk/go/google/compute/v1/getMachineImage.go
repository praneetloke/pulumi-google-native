// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified machine image.
func LookupMachineImage(ctx *pulumi.Context, args *LookupMachineImageArgs, opts ...pulumi.InvokeOption) (*LookupMachineImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMachineImageResult
	err := ctx.Invoke("google-native:compute/v1:getMachineImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineImageArgs struct {
	MachineImage string  `pulumi:"machineImage"`
	Project      *string `pulumi:"project"`
}

type LookupMachineImageResult struct {
	// The creation timestamp for this machine image in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process.
	GuestFlush bool `pulumi:"guestFlush"`
	// Properties of source instance
	InstanceProperties InstancePropertiesResponse `pulumi:"instanceProperties"`
	// The resource type, which is always compute#machineImage for machine image.
	Kind string `pulumi:"kind"`
	// Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
	MachineImageEncryptionKey CustomerEncryptionKeyResponse `pulumi:"machineImageEncryptionKey"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// An array of Machine Image specific properties for disks attached to the source instance
	SavedDisks []SavedDiskResponse `pulumi:"savedDisks"`
	// The URL for this machine image. The server defines this URL.
	SelfLink string `pulumi:"selfLink"`
	// [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKeys []SourceDiskEncryptionKeyResponse `pulumi:"sourceDiskEncryptionKeys"`
	// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
	SourceInstance string `pulumi:"sourceInstance"`
	// DEPRECATED: Please use instance_properties instead for source instance related properties. New properties will not be added to this field.
	SourceInstanceProperties SourceInstancePropertiesResponse `pulumi:"sourceInstanceProperties"`
	// The status of the machine image. One of the following values: INVALID, CREATING, READY, DELETING, and UPLOADING.
	Status string `pulumi:"status"`
	// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
	StorageLocations []string `pulumi:"storageLocations"`
	// Total size of the storage used by the machine image.
	TotalStorageBytes string `pulumi:"totalStorageBytes"`
}

func LookupMachineImageOutput(ctx *pulumi.Context, args LookupMachineImageOutputArgs, opts ...pulumi.InvokeOption) LookupMachineImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineImageResult, error) {
			args := v.(LookupMachineImageArgs)
			r, err := LookupMachineImage(ctx, &args, opts...)
			var s LookupMachineImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineImageResultOutput)
}

type LookupMachineImageOutputArgs struct {
	MachineImage pulumi.StringInput    `pulumi:"machineImage"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupMachineImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineImageArgs)(nil)).Elem()
}

type LookupMachineImageResultOutput struct{ *pulumi.OutputState }

func (LookupMachineImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineImageResult)(nil)).Elem()
}

func (o LookupMachineImageResultOutput) ToLookupMachineImageResultOutput() LookupMachineImageResultOutput {
	return o
}

func (o LookupMachineImageResultOutput) ToLookupMachineImageResultOutputWithContext(ctx context.Context) LookupMachineImageResultOutput {
	return o
}

// The creation timestamp for this machine image in RFC3339 text format.
func (o LookupMachineImageResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineImageResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupMachineImageResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineImageResult) string { return v.Description }).(pulumi.StringOutput)
}

// [Input Only] Whether to attempt an application consistent machine image by informing the OS to prepare for the snapshot process.
func (o LookupMachineImageResultOutput) GuestFlush() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMachineImageResult) bool { return v.GuestFlush }).(pulumi.BoolOutput)
}

// Properties of source instance
func (o LookupMachineImageResultOutput) InstanceProperties() InstancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupMachineImageResult) InstancePropertiesResponse { return v.InstanceProperties }).(InstancePropertiesResponseOutput)
}

// The resource type, which is always compute#machineImage for machine image.
func (o LookupMachineImageResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineImageResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Encrypts the machine image using a customer-supplied encryption key. After you encrypt a machine image using a customer-supplied key, you must provide the same key if you use the machine image later. For example, you must provide the encryption key when you create an instance from the encrypted machine image in a future request. Customer-supplied encryption keys do not protect access to metadata of the machine image. If you do not provide an encryption key when creating the machine image, then the machine image will be encrypted using an automatically generated key and you do not need to provide a key to use the machine image later.
func (o LookupMachineImageResultOutput) MachineImageEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupMachineImageResult) CustomerEncryptionKeyResponse { return v.MachineImageEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupMachineImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Reserved for future use.
func (o LookupMachineImageResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMachineImageResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// An array of Machine Image specific properties for disks attached to the source instance
func (o LookupMachineImageResultOutput) SavedDisks() SavedDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupMachineImageResult) []SavedDiskResponse { return v.SavedDisks }).(SavedDiskResponseArrayOutput)
}

// The URL for this machine image. The server defines this URL.
func (o LookupMachineImageResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineImageResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// [Input Only] The customer-supplied encryption key of the disks attached to the source instance. Required if the source disk is protected by a customer-supplied encryption key.
func (o LookupMachineImageResultOutput) SourceDiskEncryptionKeys() SourceDiskEncryptionKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupMachineImageResult) []SourceDiskEncryptionKeyResponse { return v.SourceDiskEncryptionKeys }).(SourceDiskEncryptionKeyResponseArrayOutput)
}

// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance
func (o LookupMachineImageResultOutput) SourceInstance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineImageResult) string { return v.SourceInstance }).(pulumi.StringOutput)
}

// DEPRECATED: Please use instance_properties instead for source instance related properties. New properties will not be added to this field.
func (o LookupMachineImageResultOutput) SourceInstanceProperties() SourceInstancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupMachineImageResult) SourceInstancePropertiesResponse { return v.SourceInstanceProperties }).(SourceInstancePropertiesResponseOutput)
}

// The status of the machine image. One of the following values: INVALID, CREATING, READY, DELETING, and UPLOADING.
func (o LookupMachineImageResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineImageResult) string { return v.Status }).(pulumi.StringOutput)
}

// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
func (o LookupMachineImageResultOutput) StorageLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMachineImageResult) []string { return v.StorageLocations }).(pulumi.StringArrayOutput)
}

// Total size of the storage used by the machine image.
func (o LookupMachineImageResultOutput) TotalStorageBytes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineImageResult) string { return v.TotalStorageBytes }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineImageResultOutput{})
}
