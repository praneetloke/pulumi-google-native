// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified image.
 */
export function getImage(args: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/beta:getImage", {
        "image": args.image,
        "project": args.project,
    }, opts);
}

export interface GetImageArgs {
    image: string;
    project?: string;
}

export interface GetImageResult {
    /**
     * The architecture of the image. Valid values are ARM64 or X86_64.
     */
    readonly architecture: string;
    /**
     * Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
     */
    readonly archiveSizeBytes: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * The deprecation status associated with this image.
     */
    readonly deprecated: outputs.compute.beta.DeprecationStatusResponse;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Size of the image when restored onto a persistent disk (in GB).
     */
    readonly diskSizeGb: string;
    /**
     * The name of the image family to which this image belongs. The image family name can be from a publicly managed image family provided by Compute Engine, or from a custom image family you create. For example, centos-stream-9 is a publicly available image family. For more information, see Image family best practices. When creating disks, you can specify an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.
     */
    readonly family: string;
    /**
     * A list of features to enable on the guest operating system. Applicable only for bootable images. To see a list of available options, see the guestOSfeatures[].type parameter.
     */
    readonly guestOsFeatures: outputs.compute.beta.GuestOsFeatureResponse[];
    /**
     * Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key, you must provide the same key if you use the image later (e.g. to create a disk from the image). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the image, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the image later.
     */
    readonly imageEncryptionKey: outputs.compute.beta.CustomerEncryptionKeyResponse;
    /**
     * Type of the resource. Always compute#image for images.
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this image, which is essentially a hash of the labels used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an image.
     */
    readonly labelFingerprint: string;
    /**
     * Labels to apply to this image. These can be later modified by the setLabels method.
     */
    readonly labels: {[key: string]: string};
    /**
     * Integer license codes indicating which licenses are attached to this image.
     */
    readonly licenseCodes: string[];
    /**
     * Any applicable license URI.
     */
    readonly licenses: string[];
    /**
     * A flag for marketplace VM disk created from the image, which is designed for marketplace VM disk to prevent the proprietary data on the disk from being accessed unwantedly. The flag will be inherited by the disk created from the image. The disk with locked flag set to true will be prohibited from performing the operations below: - R/W or R/O disk attach - Disk detach, if disk is created via create-on-create - Create images - Create snapshots - Create disk clone (create disk from the current disk) The image with the locked field set to true will be prohibited from performing the operations below: - Create images from the current image - Update the locked field for the current image The instance with at least one disk with locked flag set to true will be prohibited from performing the operations below: - Secondary disk attach - Create instant snapshot - Create machine images - Create instance template - Delete the instance with --keep-disk parameter set to true 
     */
    readonly locked: boolean;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The parameters of the raw disk image.
     */
    readonly rawDisk: outputs.compute.beta.ImageRawDiskResponse;
    /**
     * A rollout policy to apply to this image. When specified, the rollout policy overrides per-zone references to the image via the associated image family. The rollout policy restricts the zones where this image is accessible when using a zonal image family reference. When the rollout policy does not include the user specified zone, or if the zone is rolled out, this image is accessible. The rollout policy for this image is read-only, except for allowlisted users. This field might not be configured. To view the latest non-deprecated image in a specific zone, use the imageFamilyViews.get method.
     */
    readonly rolloutOverride: outputs.compute.beta.RolloutPolicyResponse;
    /**
     * Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Set the secure boot keys of shielded instance.
     */
    readonly shieldedInstanceInitialState: outputs.compute.beta.InitialStateConfigResponse;
    /**
     * URL of the source disk used to create this image. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    readonly sourceDisk: string;
    /**
     * The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
     */
    readonly sourceDiskEncryptionKey: outputs.compute.beta.CustomerEncryptionKeyResponse;
    /**
     * The ID value of the disk used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name.
     */
    readonly sourceDiskId: string;
    /**
     * URL of the source image used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ images/image_name - projects/project_id/global/images/image_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    readonly sourceImage: string;
    /**
     * The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key.
     */
    readonly sourceImageEncryptionKey: outputs.compute.beta.CustomerEncryptionKeyResponse;
    /**
     * The ID value of the image used to create this image. This value may be used to determine whether the image was taken from the current or a previous instance of a given image name.
     */
    readonly sourceImageId: string;
    /**
     * URL of the source snapshot used to create this image. The following are valid formats for the URL: - https://www.googleapis.com/compute/v1/projects/project_id/global/ snapshots/snapshot_name - projects/project_id/global/snapshots/snapshot_name In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
     */
    readonly sourceSnapshot: string;
    /**
     * The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key.
     */
    readonly sourceSnapshotEncryptionKey: outputs.compute.beta.CustomerEncryptionKeyResponse;
    /**
     * The ID value of the snapshot used to create this image. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name.
     */
    readonly sourceSnapshotId: string;
    /**
     * The type of the image used to create this disk. The default and only valid value is RAW.
     */
    readonly sourceType: string;
    /**
     * The status of the image. An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY. Possible values are FAILED, PENDING, or READY.
     */
    readonly status: string;
    /**
     * Cloud Storage bucket storage location of the image (regional or multi-regional).
     */
    readonly storageLocations: string[];
    /**
     * A list of publicly visible user-licenses. Unlike regular licenses, user provided licenses can be modified after the disk is created. This includes a list of URLs to the license resource. For example, to provide a debian license: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-9-stretch 
     */
    readonly userLicenses: string[];
}
/**
 * Returns the specified image.
 */
export function getImageOutput(args: GetImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageResult> {
    return pulumi.output(args).apply((a: any) => getImage(a, opts))
}

export interface GetImageOutputArgs {
    image: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
