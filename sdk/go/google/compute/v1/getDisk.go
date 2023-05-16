// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified persistent disk.
func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	var rv LookupDiskResult
	err := ctx.Invoke("google-native:compute/v1:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskArgs struct {
	Disk    string  `pulumi:"disk"`
	Project *string `pulumi:"project"`
	Zone    string  `pulumi:"zone"`
}

type LookupDiskResult struct {
	// The architecture of the disk. Valid values are ARM64 or X86_64.
	Architecture string `pulumi:"architecture"`
	// Disk asynchronously replicated into this disk.
	AsyncPrimaryDisk DiskAsyncReplicationResponse `pulumi:"asyncPrimaryDisk"`
	// A list of disks this disk is asynchronously replicated to.
	AsyncSecondaryDisks map[string]string `pulumi:"asyncSecondaryDisks"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
	DiskEncryptionKey CustomerEncryptionKeyResponse `pulumi:"diskEncryptionKey"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
	GuestOsFeatures []GuestOsFeatureResponse `pulumi:"guestOsFeatures"`
	// Type of the resource. Always compute#disk for disks.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a disk.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels to apply to this disk. These can be later modified by the setLabels method.
	Labels map[string]string `pulumi:"labels"`
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp string `pulumi:"lastAttachTimestamp"`
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp string `pulumi:"lastDetachTimestamp"`
	// Integer license codes indicating which licenses are attached to this disk.
	LicenseCodes []string `pulumi:"licenseCodes"`
	// A list of publicly visible licenses. Reserved for Google's use.
	Licenses []string `pulumi:"licenses"`
	// An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
	LocationHint string `pulumi:"locationHint"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Internal use only.
	Options string `pulumi:"options"`
	// Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
	Params DiskParamsResponse `pulumi:"params"`
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes string `pulumi:"physicalBlockSizeBytes"`
	// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
	ProvisionedIops string `pulumi:"provisionedIops"`
	// Indicates how much throughput to provision for the disk. This sets the number of throughput mb per second that the disk can handle. Values must be between 1 and 7,124.
	ProvisionedThroughput string `pulumi:"provisionedThroughput"`
	// URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
	ReplicaZones []string `pulumi:"replicaZones"`
	// Resource policies applied to this disk for automatic snapshot creations.
	ResourcePolicies []string `pulumi:"resourcePolicies"`
	// Status information for the disk resource.
	ResourceStatus DiskResourceStatusResponse `pulumi:"resourceStatus"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink string `pulumi:"selfLink"`
	// Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
	SizeGb string `pulumi:"sizeGb"`
	// URL of the DiskConsistencyGroupPolicy for a secondary disk that was created using a consistency group.
	SourceConsistencyGroupPolicy string `pulumi:"sourceConsistencyGroupPolicy"`
	// ID of the DiskConsistencyGroupPolicy for a secondary disk that was created using a consistency group.
	SourceConsistencyGroupPolicyId string `pulumi:"sourceConsistencyGroupPolicyId"`
	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk
	SourceDisk string `pulumi:"sourceDisk"`
	// The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
	SourceDiskId string `pulumi:"sourceDiskId"`
	// The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family
	SourceImage string `pulumi:"sourceImage"`
	// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
	SourceImageEncryptionKey CustomerEncryptionKeyResponse `pulumi:"sourceImageEncryptionKey"`
	// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
	SourceImageId string `pulumi:"sourceImageId"`
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot
	SourceSnapshot string `pulumi:"sourceSnapshot"`
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
	SourceSnapshotEncryptionKey CustomerEncryptionKeyResponse `pulumi:"sourceSnapshotEncryptionKey"`
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId string `pulumi:"sourceSnapshotId"`
	// The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
	SourceStorageObject string `pulumi:"sourceStorageObject"`
	// The status of disk creation. - CREATING: Disk is provisioning. - RESTORING: Source data is being copied into the disk. - FAILED: Disk creation failed. - READY: Disk is ready for use. - DELETING: Disk is deleting.
	Status string `pulumi:"status"`
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
	Type string `pulumi:"type"`
	// Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
	Users []string `pulumi:"users"`
	// URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone string `pulumi:"zone"`
}

func LookupDiskOutput(ctx *pulumi.Context, args LookupDiskOutputArgs, opts ...pulumi.InvokeOption) LookupDiskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskResult, error) {
			args := v.(LookupDiskArgs)
			r, err := LookupDisk(ctx, &args, opts...)
			var s LookupDiskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskResultOutput)
}

type LookupDiskOutputArgs struct {
	Disk    pulumi.StringInput    `pulumi:"disk"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Zone    pulumi.StringInput    `pulumi:"zone"`
}

func (LookupDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskArgs)(nil)).Elem()
}

type LookupDiskResultOutput struct{ *pulumi.OutputState }

func (LookupDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskResult)(nil)).Elem()
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutput() LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutputWithContext(ctx context.Context) LookupDiskResultOutput {
	return o
}

// The architecture of the disk. Valid values are ARM64 or X86_64.
func (o LookupDiskResultOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Architecture }).(pulumi.StringOutput)
}

// Disk asynchronously replicated into this disk.
func (o LookupDiskResultOutput) AsyncPrimaryDisk() DiskAsyncReplicationResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) DiskAsyncReplicationResponse { return v.AsyncPrimaryDisk }).(DiskAsyncReplicationResponseOutput)
}

// A list of disks this disk is asynchronously replicated to.
func (o LookupDiskResultOutput) AsyncSecondaryDisks() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskResult) map[string]string { return v.AsyncSecondaryDisks }).(pulumi.StringMapOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupDiskResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupDiskResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Description }).(pulumi.StringOutput)
}

// Encrypts the disk using a customer-supplied encryption key or a customer-managed encryption key. Encryption keys do not protect access to metadata of the disk. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later. For example, to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine. After you encrypt a disk with a customer-managed key, the diskEncryptionKey.kmsKeyName is set to a key *version* name once the disk is created. The disk is encrypted with this version of the key. In the response, diskEncryptionKey.kmsKeyName appears in the following format: "diskEncryptionKey.kmsKeyName": "projects/kms_project_id/locations/region/keyRings/ key_region/cryptoKeys/key /cryptoKeysVersions/version If you do not provide an encryption key when creating the disk, then the disk is encrypted using an automatically generated key and you don't need to provide a key to use the disk later.
func (o LookupDiskResultOutput) DiskEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) CustomerEncryptionKeyResponse { return v.DiskEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
func (o LookupDiskResultOutput) GuestOsFeatures() GuestOsFeatureResponseArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []GuestOsFeatureResponse { return v.GuestOsFeatures }).(GuestOsFeatureResponseArrayOutput)
}

// Type of the resource. Always compute#disk for disks.
func (o LookupDiskResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a disk.
func (o LookupDiskResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this disk. These can be later modified by the setLabels method.
func (o LookupDiskResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Last attach timestamp in RFC3339 text format.
func (o LookupDiskResultOutput) LastAttachTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.LastAttachTimestamp }).(pulumi.StringOutput)
}

// Last detach timestamp in RFC3339 text format.
func (o LookupDiskResultOutput) LastDetachTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.LastDetachTimestamp }).(pulumi.StringOutput)
}

// Integer license codes indicating which licenses are attached to this disk.
func (o LookupDiskResultOutput) LicenseCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.LicenseCodes }).(pulumi.StringArrayOutput)
}

// A list of publicly visible licenses. Reserved for Google's use.
func (o LookupDiskResultOutput) Licenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.Licenses }).(pulumi.StringArrayOutput)
}

// An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
func (o LookupDiskResultOutput) LocationHint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.LocationHint }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

// Internal use only.
func (o LookupDiskResultOutput) Options() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Options }).(pulumi.StringOutput)
}

// Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
func (o LookupDiskResultOutput) Params() DiskParamsResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) DiskParamsResponse { return v.Params }).(DiskParamsResponseOutput)
}

// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
func (o LookupDiskResultOutput) PhysicalBlockSizeBytes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.PhysicalBlockSizeBytes }).(pulumi.StringOutput)
}

// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
func (o LookupDiskResultOutput) ProvisionedIops() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.ProvisionedIops }).(pulumi.StringOutput)
}

// Indicates how much throughput to provision for the disk. This sets the number of throughput mb per second that the disk can handle. Values must be between 1 and 7,124.
func (o LookupDiskResultOutput) ProvisionedThroughput() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.ProvisionedThroughput }).(pulumi.StringOutput)
}

// URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupDiskResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Region }).(pulumi.StringOutput)
}

// URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
func (o LookupDiskResultOutput) ReplicaZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.ReplicaZones }).(pulumi.StringArrayOutput)
}

// Resource policies applied to this disk for automatic snapshot creations.
func (o LookupDiskResultOutput) ResourcePolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.ResourcePolicies }).(pulumi.StringArrayOutput)
}

// Status information for the disk resource.
func (o LookupDiskResultOutput) ResourceStatus() DiskResourceStatusResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) DiskResourceStatusResponse { return v.ResourceStatus }).(DiskResourceStatusResponseOutput)
}

// Reserved for future use.
func (o LookupDiskResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDiskResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Server-defined fully-qualified URL for this resource.
func (o LookupDiskResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
func (o LookupDiskResultOutput) SizeGb() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SizeGb }).(pulumi.StringOutput)
}

// URL of the DiskConsistencyGroupPolicy for a secondary disk that was created using a consistency group.
func (o LookupDiskResultOutput) SourceConsistencyGroupPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceConsistencyGroupPolicy }).(pulumi.StringOutput)
}

// ID of the DiskConsistencyGroupPolicy for a secondary disk that was created using a consistency group.
func (o LookupDiskResultOutput) SourceConsistencyGroupPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceConsistencyGroupPolicyId }).(pulumi.StringOutput)
}

// The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk
func (o LookupDiskResultOutput) SourceDisk() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceDisk }).(pulumi.StringOutput)
}

// The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
func (o LookupDiskResultOutput) SourceDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceDiskId }).(pulumi.StringOutput)
}

// The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family
func (o LookupDiskResultOutput) SourceImage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceImage }).(pulumi.StringOutput)
}

// The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
func (o LookupDiskResultOutput) SourceImageEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) CustomerEncryptionKeyResponse { return v.SourceImageEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
func (o LookupDiskResultOutput) SourceImageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceImageId }).(pulumi.StringOutput)
}

// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot
func (o LookupDiskResultOutput) SourceSnapshot() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceSnapshot }).(pulumi.StringOutput)
}

// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
func (o LookupDiskResultOutput) SourceSnapshotEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) CustomerEncryptionKeyResponse { return v.SourceSnapshotEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
func (o LookupDiskResultOutput) SourceSnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceSnapshotId }).(pulumi.StringOutput)
}

// The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
func (o LookupDiskResultOutput) SourceStorageObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.SourceStorageObject }).(pulumi.StringOutput)
}

// The status of disk creation. - CREATING: Disk is provisioning. - RESTORING: Source data is being copied into the disk. - FAILED: Disk creation failed. - READY: Disk is ready for use. - DELETING: Disk is deleting.
func (o LookupDiskResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Status }).(pulumi.StringOutput)
}

// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
func (o LookupDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

// Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
func (o LookupDiskResultOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.Users }).(pulumi.StringArrayOutput)
}

// URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupDiskResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskResultOutput{})
}
