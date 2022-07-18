// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Backup for the given BackupPlan.
 * Auto-naming is currently not supported for this resource.
 */
export class Backup extends pulumi.CustomResource {
    /**
     * Get an existing Backup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Backup {
        return new Backup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gkebackup/v1:Backup';

    /**
     * Returns true if the given object is an instance of Backup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backup.__pulumiType;
    }

    /**
     * If True, all namespaces were included in the Backup.
     */
    public /*out*/ readonly allNamespaces!: pulumi.Output<boolean>;
    /**
     * The client-provided short name for the Backup resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of Backups in this BackupPlan
     */
    public readonly backupId!: pulumi.Output<string | undefined>;
    public readonly backupPlanId!: pulumi.Output<string>;
    /**
     * Information about the GKE cluster from which this Backup was created.
     */
    public /*out*/ readonly clusterMetadata!: pulumi.Output<outputs.gkebackup.v1.ClusterMetadataResponse>;
    /**
     * Completion time of the Backup
     */
    public /*out*/ readonly completeTime!: pulumi.Output<string>;
    /**
     * The size of the config backup in bytes.
     */
    public /*out*/ readonly configBackupSizeBytes!: pulumi.Output<string>;
    /**
     * Whether or not the Backup contains Kubernetes Secrets. Controlled by the parent BackupPlan's include_secrets value.
     */
    public /*out*/ readonly containsSecrets!: pulumi.Output<boolean>;
    /**
     * Whether or not the Backup contains volume data. Controlled by the parent BackupPlan's include_volume_data value.
     */
    public /*out*/ readonly containsVolumeData!: pulumi.Output<boolean>;
    /**
     * The timestamp when this Backup resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Minimum age for this Backup (in days). If this field is set to a non-zero value, the Backup will be "locked" against deletion (either manual or automatic deletion) for the number of days provided (measured from the creation time of the Backup). MUST be an integer value between 0-90 (inclusive). Defaults to parent BackupPlan's backup_delete_lock_days setting and may only be increased (either at creation time or in a subsequent update).
     */
    public readonly deleteLockDays!: pulumi.Output<number>;
    /**
     * The time at which an existing delete lock will expire for this backup (calculated from create_time + delete_lock_days).
     */
    public /*out*/ readonly deleteLockExpireTime!: pulumi.Output<string>;
    /**
     * User specified descriptive string for this Backup.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The customer managed encryption key that was used to encrypt the Backup's artifacts. Inherited from the parent BackupPlan's encryption_key value.
     */
    public /*out*/ readonly encryptionKey!: pulumi.Output<outputs.gkebackup.v1.EncryptionKeyResponse>;
    /**
     * `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a backup from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform backup updates in order to avoid race conditions: An `etag` is returned in the response to `GetBackup`, and systems are expected to put that etag in the request to `UpdateBackup` or `DeleteBackup` to ensure that their change will be applied to the same version of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A set of custom labels supplied by user.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * This flag indicates whether this Backup resource was created manually by a user or via a schedule in the BackupPlan. A value of True means that the Backup was created manually.
     */
    public /*out*/ readonly manual!: pulumi.Output<boolean>;
    /**
     * The fully qualified name of the Backup. `projects/*&#47;locations/*&#47;backupPlans/*&#47;backups/*`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The total number of Kubernetes Pods contained in the Backup.
     */
    public /*out*/ readonly podCount!: pulumi.Output<number>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The total number of Kubernetes resources included in the Backup.
     */
    public /*out*/ readonly resourceCount!: pulumi.Output<number>;
    /**
     * The age (in days) after which this Backup will be automatically deleted. Must be an integer value >= 0: - If 0, no automatic deletion will occur for this Backup. - If not 0, this must be >= delete_lock_days and <= 365. Once a Backup is created, this value may only be increased. Defaults to the parent BackupPlan's backup_retain_days value.
     */
    public readonly retainDays!: pulumi.Output<number>;
    /**
     * The time at which this Backup will be automatically deleted (calculated from create_time + retain_days).
     */
    public /*out*/ readonly retainExpireTime!: pulumi.Output<string>;
    /**
     * If set, the list of ProtectedApplications whose resources were included in the Backup.
     */
    public /*out*/ readonly selectedApplications!: pulumi.Output<outputs.gkebackup.v1.NamespacedNamesResponse>;
    /**
     * If set, the list of namespaces that were included in the Backup.
     */
    public /*out*/ readonly selectedNamespaces!: pulumi.Output<outputs.gkebackup.v1.NamespacesResponse>;
    /**
     * The total size of the Backup in bytes = config backup size + sum(volume backup sizes)
     */
    public /*out*/ readonly sizeBytes!: pulumi.Output<string>;
    /**
     * Current state of the Backup
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Human-readable description of why the backup is in the current `state`.
     */
    public /*out*/ readonly stateReason!: pulumi.Output<string>;
    /**
     * Server generated global unique identifier of [UUID4](https://en.wikipedia.org/wiki/Universally_unique_identifier)
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The timestamp when this Backup resource was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The total number of volume backups contained in the Backup.
     */
    public /*out*/ readonly volumeCount!: pulumi.Output<number>;

    /**
     * Create a Backup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.backupPlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupPlanId'");
            }
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["backupPlanId"] = args ? args.backupPlanId : undefined;
            resourceInputs["deleteLockDays"] = args ? args.deleteLockDays : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["retainDays"] = args ? args.retainDays : undefined;
            resourceInputs["allNamespaces"] = undefined /*out*/;
            resourceInputs["clusterMetadata"] = undefined /*out*/;
            resourceInputs["completeTime"] = undefined /*out*/;
            resourceInputs["configBackupSizeBytes"] = undefined /*out*/;
            resourceInputs["containsSecrets"] = undefined /*out*/;
            resourceInputs["containsVolumeData"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteLockExpireTime"] = undefined /*out*/;
            resourceInputs["encryptionKey"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["manual"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["podCount"] = undefined /*out*/;
            resourceInputs["resourceCount"] = undefined /*out*/;
            resourceInputs["retainExpireTime"] = undefined /*out*/;
            resourceInputs["selectedApplications"] = undefined /*out*/;
            resourceInputs["selectedNamespaces"] = undefined /*out*/;
            resourceInputs["sizeBytes"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateReason"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["volumeCount"] = undefined /*out*/;
        } else {
            resourceInputs["allNamespaces"] = undefined /*out*/;
            resourceInputs["backupId"] = undefined /*out*/;
            resourceInputs["backupPlanId"] = undefined /*out*/;
            resourceInputs["clusterMetadata"] = undefined /*out*/;
            resourceInputs["completeTime"] = undefined /*out*/;
            resourceInputs["configBackupSizeBytes"] = undefined /*out*/;
            resourceInputs["containsSecrets"] = undefined /*out*/;
            resourceInputs["containsVolumeData"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteLockDays"] = undefined /*out*/;
            resourceInputs["deleteLockExpireTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["encryptionKey"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["manual"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["podCount"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["resourceCount"] = undefined /*out*/;
            resourceInputs["retainDays"] = undefined /*out*/;
            resourceInputs["retainExpireTime"] = undefined /*out*/;
            resourceInputs["selectedApplications"] = undefined /*out*/;
            resourceInputs["selectedNamespaces"] = undefined /*out*/;
            resourceInputs["sizeBytes"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateReason"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["volumeCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["backupPlanId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Backup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Backup resource.
 */
export interface BackupArgs {
    /**
     * The client-provided short name for the Backup resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of Backups in this BackupPlan
     */
    backupId?: pulumi.Input<string>;
    backupPlanId: pulumi.Input<string>;
    /**
     * Minimum age for this Backup (in days). If this field is set to a non-zero value, the Backup will be "locked" against deletion (either manual or automatic deletion) for the number of days provided (measured from the creation time of the Backup). MUST be an integer value between 0-90 (inclusive). Defaults to parent BackupPlan's backup_delete_lock_days setting and may only be increased (either at creation time or in a subsequent update).
     */
    deleteLockDays?: pulumi.Input<number>;
    /**
     * User specified descriptive string for this Backup.
     */
    description?: pulumi.Input<string>;
    /**
     * A set of custom labels supplied by user.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The age (in days) after which this Backup will be automatically deleted. Must be an integer value >= 0: - If 0, no automatic deletion will occur for this Backup. - If not 0, this must be >= delete_lock_days and <= 365. Once a Backup is created, this value may only be increased. Defaults to the parent BackupPlan's backup_retain_days value.
     */
    retainDays?: pulumi.Input<number>;
}
