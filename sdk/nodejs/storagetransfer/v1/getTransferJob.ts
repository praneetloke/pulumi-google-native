// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a transfer job.
 */
export function getTransferJob(args: GetTransferJobArgs, opts?: pulumi.InvokeOptions): Promise<GetTransferJobResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:storagetransfer/v1:getTransferJob", {
        "projectId": args.projectId,
        "transferJobId": args.transferJobId,
    }, opts);
}

export interface GetTransferJobArgs {
    projectId: string;
    transferJobId: string;
}

export interface GetTransferJobResult {
    /**
     * The time that the transfer job was created.
     */
    readonly creationTime: string;
    /**
     * The time that the transfer job was deleted.
     */
    readonly deletionTime: string;
    /**
     * A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
     */
    readonly description: string;
    /**
     * The time that the transfer job was last modified.
     */
    readonly lastModificationTime: string;
    /**
     * The name of the most recently started TransferOperation of this JobConfig. Present if a TransferOperation has been created for this JobConfig.
     */
    readonly latestOperationName: string;
    /**
     * A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service assigns a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. For transfers involving PosixFilesystem, this name must start with 'transferJobs/OPI' specifically. For all other transfer types, this name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix for PosixFilesystem transfers. Non-PosixFilesystem example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` PosixFilesystem example: `"transferJobs/OPI^[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Applications must not rely on the enforcement of naming requirements involving OPI. Invalid job names fail with an INVALID_ARGUMENT error.
     */
    readonly name: string;
    /**
     * Notification configuration. This is not supported for transfers involving PosixFilesystem.
     */
    readonly notificationConfig: outputs.storagetransfer.v1.NotificationConfigResponse;
    /**
     * The ID of the Google Cloud Platform Project that owns the job.
     */
    readonly project: string;
    /**
     * Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job never executes a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
     */
    readonly schedule: outputs.storagetransfer.v1.ScheduleResponse;
    /**
     * Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
     */
    readonly status: string;
    /**
     * Transfer specification.
     */
    readonly transferSpec: outputs.storagetransfer.v1.TransferSpecResponse;
}