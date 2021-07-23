// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns a specified persistent disk. Gets a list of available persistent disks by making a list() request.
 */
export function getDisk(args: GetDiskArgs, opts?: pulumi.InvokeOptions): Promise<GetDiskResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/v1:getDisk", {
        "disk": args.disk,
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

export interface GetDiskArgs {
    disk: string;
    project?: string;
    zone: string;
}

export interface GetDiskResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later (e.g. to create a disk snapshot, to create a disk image, to create a machine image, or to attach the disk to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later.
     */
    readonly diskEncryptionKey: outputs.compute.v1.CustomerEncryptionKeyResponse;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
     */
    readonly guestOsFeatures: outputs.compute.v1.GuestOsFeatureResponse[];
    /**
     * Type of the resource. Always compute#disk for disks.
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a disk.
     */
    readonly labelFingerprint: string;
    /**
     * Labels to apply to this disk. These can be later modified by the setLabels method.
     */
    readonly labels: {[key: string]: string};
    /**
     * Last attach timestamp in RFC3339 text format.
     */
    readonly lastAttachTimestamp: string;
    /**
     * Last detach timestamp in RFC3339 text format.
     */
    readonly lastDetachTimestamp: string;
    /**
     * Integer license codes indicating which licenses are attached to this disk.
     */
    readonly licenseCodes: string[];
    /**
     * A list of publicly visible licenses. Reserved for Google's use.
     */
    readonly licenses: string[];
    /**
     * An opaque location hint used to place the disk close to other resources. This field is for use by internal tools that use the public API.
     */
    readonly locationHint: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Internal use only.
     */
    readonly options: string;
    /**
     * Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. The currently supported size is 4096, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.
     */
    readonly physicalBlockSizeBytes: string;
    /**
     * Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle. Values must be between 10,000 and 120,000. For more details, see the Extreme persistent disk documentation.
     */
    readonly provisionedIops: string;
    /**
     * URL of the region where the disk resides. Only applicable for regional resources. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * URLs of the zones where the disk should be replicated to. Only applicable for regional resources.
     */
    readonly replicaZones: string[];
    /**
     * Resource policies applied to this disk for automatic snapshot creations.
     */
    readonly resourcePolicies: string[];
    /**
     * Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    readonly selfLink: string;
    /**
     * Size, in GB, of the persistent disk. You can specify this field when creating a persistent disk using the sourceImage, sourceSnapshot, or sourceDisk parameter, or specify it alone to create an empty persistent disk. If you specify this field along with a source, the value of sizeGb must not be less than the size of the source. Acceptable values are 1 to 65536, inclusive.
     */
    readonly sizeGb: string;
    /**
     * The source disk used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - https://www.googleapis.com/compute/v1/projects/project/regions/region /disks/disk - projects/project/zones/zone/disks/disk - projects/project/regions/region/disks/disk - zones/zone/disks/disk - regions/region/disks/disk 
     */
    readonly sourceDisk: string;
    /**
     * The unique ID of the disk used to create this disk. This value identifies the exact disk that was used to create this persistent disk. For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used.
     */
    readonly sourceDiskId: string;
    /**
     * The source image used to create this disk. If the source image is deleted, this field will not be set. To create a disk with one of the public operating system images, specify the image by its family name. For example, specify family/debian-9 to use the latest Debian 9 image: projects/debian-cloud/global/images/family/debian-9 Alternatively, use a specific version of a public operating system image: projects/debian-cloud/global/images/debian-9-stretch-vYYYYMMDD To create a disk with a custom image that you created, specify the image name in the following format: global/images/my-custom-image You can also specify a custom image by its image family, which returns the latest version of the image in that family. Replace the image name with family/family-name: global/images/family/my-image-family 
     */
    readonly sourceImage: string;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    readonly sourceImageEncryptionKey: outputs.compute.v1.CustomerEncryptionKeyResponse;
    /**
     * The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.
     */
    readonly sourceImageId: string;
    /**
     * The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project /global/snapshots/snapshot - projects/project/global/snapshots/snapshot - global/snapshots/snapshot 
     */
    readonly sourceSnapshot: string;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    readonly sourceSnapshotEncryptionKey: outputs.compute.v1.CustomerEncryptionKeyResponse;
    /**
     * The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
     */
    readonly sourceSnapshotId: string;
    /**
     * The full Google Cloud Storage URI where the disk image is stored. This file must be a gzip-compressed tarball whose name ends in .tar.gz or virtual machine disk whose name ends in vmdk. Valid URIs may start with gs:// or https://storage.googleapis.com/. This flag is not optimized for creating multiple disks from a source storage object. To create many disks from a source storage object, use gcloud compute images import instead.
     */
    readonly sourceStorageObject: string;
    /**
     * The status of disk creation. - CREATING: Disk is provisioning. - RESTORING: Source data is being copied into the disk. - FAILED: Disk creation failed. - READY: Disk is ready for use. - DELETING: Disk is deleting. 
     */
    readonly status: string;
    /**
     * URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk. For example: projects/project /zones/zone/diskTypes/pd-ssd . See Persistent disk types.
     */
    readonly type: string;
    /**
     * Links to the users of the disk (attached instances) in form: projects/project/zones/zone/instances/instance
     */
    readonly users: string[];
    /**
     * URL of the zone where the disk resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly zone: string;
}