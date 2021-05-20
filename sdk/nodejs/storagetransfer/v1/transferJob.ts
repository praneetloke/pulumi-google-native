// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a transfer job that runs periodically.
 */
export class TransferJob extends pulumi.CustomResource {
    /**
     * Get an existing TransferJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransferJob {
        return new TransferJob(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:storagetransfer/v1:TransferJob';

    /**
     * Returns true if the given object is an instance of TransferJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransferJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransferJob.__pulumiType;
    }

    /**
     * The time that the transfer job was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The time that the transfer job was deleted.
     */
    public /*out*/ readonly deletionTime!: pulumi.Output<string>;
    /**
     * A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The time that the transfer job was last modified.
     */
    public /*out*/ readonly lastModificationTime!: pulumi.Output<string>;
    /**
     * The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
     */
    public readonly latestOperationName!: pulumi.Output<string>;
    /**
     * A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Notification configuration.
     */
    public readonly notificationConfig!: pulumi.Output<outputs.storagetransfer.v1.NotificationConfigResponse>;
    /**
     * The ID of the Google Cloud Platform Project that owns the job.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
     */
    public readonly schedule!: pulumi.Output<outputs.storagetransfer.v1.ScheduleResponse>;
    /**
     * Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Transfer specification.
     */
    public readonly transferSpec!: pulumi.Output<outputs.storagetransfer.v1.TransferSpecResponse>;

    /**
     * Create a TransferJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransferJobArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.transferJobId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transferJobId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["latestOperationName"] = args ? args.latestOperationName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["transferJobId"] = args ? args.transferJobId : undefined;
            inputs["transferSpec"] = args ? args.transferSpec : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["deletionTime"] = undefined /*out*/;
            inputs["lastModificationTime"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["deletionTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["lastModificationTime"] = undefined /*out*/;
            inputs["latestOperationName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notificationConfig"] = undefined /*out*/;
            inputs["project"] = undefined /*out*/;
            inputs["schedule"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["transferSpec"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TransferJob.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransferJob resource.
 */
export interface TransferJobArgs {
    /**
     * A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
     */
    readonly latestOperationName?: pulumi.Input<string>;
    /**
     * A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Notification configuration.
     */
    readonly notificationConfig?: pulumi.Input<inputs.storagetransfer.v1.NotificationConfigArgs>;
    /**
     * The ID of the Google Cloud Platform Project that owns the job.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
     */
    readonly schedule?: pulumi.Input<inputs.storagetransfer.v1.ScheduleArgs>;
    /**
     * Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
     */
    readonly status?: pulumi.Input<string>;
    readonly transferJobId: pulumi.Input<string>;
    /**
     * Transfer specification.
     */
    readonly transferSpec?: pulumi.Input<inputs.storagetransfer.v1.TransferSpecArgs>;
}
