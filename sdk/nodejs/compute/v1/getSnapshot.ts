// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified Snapshot resource.
 */
export function getSnapshot(args: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/v1:getSnapshot", {
        "project": args.project,
        "snapshot": args.snapshot,
    }, opts);
}

export interface GetSnapshotArgs {
    project?: string;
    snapshot: string;
}

export interface GetSnapshotResult {
    /**
     * The architecture of the snapshot. Valid values are ARM64 or X86_64.
     */
    readonly architecture: string;
    /**
     * Set to true if snapshots are automatically created by applying resource policy on the target disk.
     */
    readonly autoCreated: boolean;
    /**
     * Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
     */
    readonly chainName: string;
    /**
     * Size in bytes of the snapshot at creation time.
     */
    readonly creationSizeBytes: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Size of the source disk, specified in GB.
     */
    readonly diskSizeGb: string;
    /**
     * Number of bytes downloaded to restore a snapshot to a disk.
     */
    readonly downloadBytes: string;
    /**
     * Type of the resource. Always compute#snapshot for Snapshot resources.
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a snapshot.
     */
    readonly labelFingerprint: string;
    /**
     * Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
     */
    readonly labels: {[key: string]: string};
    /**
     * Integer license codes indicating which licenses are attached to this snapshot.
     */
    readonly licenseCodes: string[];
    /**
     * A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image).
     */
    readonly licenses: string[];
    /**
     * An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
     */
    readonly locationHint: string;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
     */
    readonly snapshotEncryptionKey: outputs.compute.v1.CustomerEncryptionKeyResponse;
    /**
     * Indicates the type of the snapshot.
     */
    readonly snapshotType: string;
    /**
     * The source disk used to create this snapshot.
     */
    readonly sourceDisk: string;
    /**
     * The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
     */
    readonly sourceDiskEncryptionKey: outputs.compute.v1.CustomerEncryptionKeyResponse;
    /**
     * The ID value of the disk used to create this snapshot. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given disk name.
     */
    readonly sourceDiskId: string;
    /**
     * URL of the resource policy which created this scheduled snapshot.
     */
    readonly sourceSnapshotSchedulePolicy: string;
    /**
     * ID of the resource policy which created this scheduled snapshot.
     */
    readonly sourceSnapshotSchedulePolicyId: string;
    /**
     * The status of the snapshot. This can be CREATING, DELETING, FAILED, READY, or UPLOADING.
     */
    readonly status: string;
    /**
     * A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
     */
    readonly storageBytes: string;
    /**
     * An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation. This status can either be UPDATING, meaning the size of the snapshot is being updated, or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
     */
    readonly storageBytesStatus: string;
    /**
     * Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
     */
    readonly storageLocations: string[];
}
/**
 * Returns the specified Snapshot resource.
 */
export function getSnapshotOutput(args: GetSnapshotOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSnapshotResult> {
    return pulumi.output(args).apply((a: any) => getSnapshot(a, opts))
}

export interface GetSnapshotOutputArgs {
    project?: pulumi.Input<string>;
    snapshot: pulumi.Input<string>;
}
