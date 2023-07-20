// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified image.
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupImageResult
	err := ctx.Invoke("google-native:compute/v1:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImageArgs struct {
	Image   string  `pulumi:"image"`
	Project *string `pulumi:"project"`
}

type LookupImageResult struct {
	// The architecture of the image. Valid values are ARM64 or X86_64.
	Architecture string `pulumi:"architecture"`
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes string `pulumi:"archiveSizeBytes"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// The deprecation status associated with this image.
	Deprecated DeprecationStatusResponse `pulumi:"deprecated"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb string `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. The image family name can be from a publicly managed image family provided by Compute Engine, or from a custom image family you create. For example, centos-stream-9 is a publicly available image family. For more information, see Image family best practices. When creating disks, you can specify an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
	Family string `pulumi:"family"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images. To see a list of available options, see the guestOSfeatures[].type parameter.
	GuestOsFeatures []GuestOsFeatureResponse `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
	ImageEncryptionKey CustomerEncryptionKeyResponse `pulumi:"imageEncryptionKey"`
	// Type of the resource. Always compute#image for images.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an image.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels to apply to this image. These can be later modified by the setLabels method.
	Labels map[string]string `pulumi:"labels"`
	// Integer license codes indicating which licenses are attached to this image.
	LicenseCodes []string `pulumi:"licenseCodes"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The parameters of the raw disk image.
	RawDisk ImageRawDiskResponse `pulumi:"rawDisk"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Set the secure boot keys of shielded instance.
	ShieldedInstanceInitialState InitialStateConfigResponse `pulumi:"shieldedInstanceInitialState"`
	// URL of the source disk used to create this image. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL
	SourceDisk string `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey CustomerEncryptionKeyResponse `pulumi:"sourceDiskEncryptionKey"`
	// The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
	SourceDiskId string `pulumi:"sourceDiskId"`
	// URL of the source image used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ images/image_name - projects/project_id/global/images/image_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL
	SourceImage string `pulumi:"sourceImage"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey CustomerEncryptionKeyResponse `pulumi:"sourceImageEncryptionKey"`
	// The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
	SourceImageId string `pulumi:"sourceImageId"`
	// URL of the source snapshot used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ snapshots/snapshot_name - projects/project_id/global/snapshots/snapshot_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL
	SourceSnapshot string `pulumi:"sourceSnapshot"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey CustomerEncryptionKeyResponse `pulumi:"sourceSnapshotEncryptionKey"`
	// The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
	SourceSnapshotId string `pulumi:"sourceSnapshotId"`
	// The type of the image used to create this disk. The default and only valid value is RAW.
	SourceType string `pulumi:"sourceType"`
	// The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
	Status string `pulumi:"status"`
	// Cloud Storage bucket storage location of the image (regional or multi-regional).
	StorageLocations []string `pulumi:"storageLocations"`
}

func LookupImageOutput(ctx *pulumi.Context, args LookupImageOutputArgs, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageResult, error) {
			args := v.(LookupImageArgs)
			r, err := LookupImage(ctx, &args, opts...)
			var s LookupImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageResultOutput)
}

type LookupImageOutputArgs struct {
	Image   pulumi.StringInput    `pulumi:"image"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

// The architecture of the image. Valid values are ARM64 or X86_64.
func (o LookupImageResultOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Architecture }).(pulumi.StringOutput)
}

// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
func (o LookupImageResultOutput) ArchiveSizeBytes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ArchiveSizeBytes }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupImageResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// The deprecation status associated with this image.
func (o LookupImageResultOutput) Deprecated() DeprecationStatusResponseOutput {
	return o.ApplyT(func(v LookupImageResult) DeprecationStatusResponse { return v.Deprecated }).(DeprecationStatusResponseOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupImageResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Description }).(pulumi.StringOutput)
}

// Size of the image when restored onto a persistent disk (in GB).
func (o LookupImageResultOutput) DiskSizeGb() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.DiskSizeGb }).(pulumi.StringOutput)
}

// The name of the image family to which this image belongs. The image family name can be from a publicly managed image family provided by Compute Engine, or from a custom image family you create. For example, centos-stream-9 is a publicly available image family. For more information, see Image family best practices. When creating disks, you can specify an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
func (o LookupImageResultOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Family }).(pulumi.StringOutput)
}

// A list of features to enable on the guest operating system. Applicable only for bootable images. To see a list of available options, see the guestOSfeatures[].type parameter.
func (o LookupImageResultOutput) GuestOsFeatures() GuestOsFeatureResponseArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []GuestOsFeatureResponse { return v.GuestOsFeatures }).(GuestOsFeatureResponseArrayOutput)
}

// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
func (o LookupImageResultOutput) ImageEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupImageResult) CustomerEncryptionKeyResponse { return v.ImageEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// Type of the resource. Always compute#image for images.
func (o LookupImageResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an image.
func (o LookupImageResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this image. These can be later modified by the setLabels method.
func (o LookupImageResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImageResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Integer license codes indicating which licenses are attached to this image.
func (o LookupImageResultOutput) LicenseCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.LicenseCodes }).(pulumi.StringArrayOutput)
}

// Any applicable license URI.
func (o LookupImageResultOutput) Licenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.Licenses }).(pulumi.StringArrayOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// The parameters of the raw disk image.
func (o LookupImageResultOutput) RawDisk() ImageRawDiskResponseOutput {
	return o.ApplyT(func(v LookupImageResult) ImageRawDiskResponse { return v.RawDisk }).(ImageRawDiskResponseOutput)
}

// Reserved for future use.
func (o LookupImageResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Server-defined URL for the resource.
func (o LookupImageResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Set the secure boot keys of shielded instance.
func (o LookupImageResultOutput) ShieldedInstanceInitialState() InitialStateConfigResponseOutput {
	return o.ApplyT(func(v LookupImageResult) InitialStateConfigResponse { return v.ShieldedInstanceInitialState }).(InitialStateConfigResponseOutput)
}

// URL of the source disk used to create this image. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL
func (o LookupImageResultOutput) SourceDisk() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceDisk }).(pulumi.StringOutput)
}

// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
func (o LookupImageResultOutput) SourceDiskEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupImageResult) CustomerEncryptionKeyResponse { return v.SourceDiskEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
func (o LookupImageResultOutput) SourceDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceDiskId }).(pulumi.StringOutput)
}

// URL of the source image used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ images/image_name - projects/project_id/global/images/image_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL
func (o LookupImageResultOutput) SourceImage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceImage }).(pulumi.StringOutput)
}

// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
func (o LookupImageResultOutput) SourceImageEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupImageResult) CustomerEncryptionKeyResponse { return v.SourceImageEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
func (o LookupImageResultOutput) SourceImageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceImageId }).(pulumi.StringOutput)
}

// URL of the source snapshot used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ snapshots/snapshot_name - projects/project_id/global/snapshots/snapshot_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL
func (o LookupImageResultOutput) SourceSnapshot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceSnapshot }).(pulumi.StringOutput)
}

// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
func (o LookupImageResultOutput) SourceSnapshotEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupImageResult) CustomerEncryptionKeyResponse { return v.SourceSnapshotEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
func (o LookupImageResultOutput) SourceSnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceSnapshotId }).(pulumi.StringOutput)
}

// The type of the image used to create this disk. The default and only valid value is RAW.
func (o LookupImageResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.SourceType }).(pulumi.StringOutput)
}

// The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
func (o LookupImageResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Status }).(pulumi.StringOutput)
}

// Cloud Storage bucket storage location of the image (regional or multi-regional).
func (o LookupImageResultOutput) StorageLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.StorageLocations }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
