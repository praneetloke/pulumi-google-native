// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1
{
    public static class GetTransferJob
    {
        /// <summary>
        /// Gets a transfer job.
        /// </summary>
        public static Task<GetTransferJobResult> InvokeAsync(GetTransferJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransferJobResult>("google-native:storagetransfer/v1:getTransferJob", args ?? new GetTransferJobArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a transfer job.
        /// </summary>
        public static Output<GetTransferJobResult> Invoke(GetTransferJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTransferJobResult>("google-native:storagetransfer/v1:getTransferJob", args ?? new GetTransferJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransferJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        [Input("transferJobId", required: true)]
        public string TransferJobId { get; set; } = null!;

        public GetTransferJobArgs()
        {
        }
        public static new GetTransferJobArgs Empty => new GetTransferJobArgs();
    }

    public sealed class GetTransferJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("transferJobId", required: true)]
        public Input<string> TransferJobId { get; set; } = null!;

        public GetTransferJobInvokeArgs()
        {
        }
        public static new GetTransferJobInvokeArgs Empty => new GetTransferJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransferJobResult
    {
        /// <summary>
        /// The time that the transfer job was created.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The time that the transfer job was deleted.
        /// </summary>
        public readonly string DeletionTime;
        /// <summary>
        /// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The time that the transfer job was last modified.
        /// </summary>
        public readonly string LastModificationTime;
        /// <summary>
        /// The name of the most recently started TransferOperation of this JobConfig. Present if a TransferOperation has been created for this JobConfig.
        /// </summary>
        public readonly string LatestOperationName;
        /// <summary>
        /// Logging configuration.
        /// </summary>
        public readonly Outputs.LoggingConfigResponse LoggingConfig;
        /// <summary>
        /// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service assigns a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. For transfers involving PosixFilesystem, this name must start with `transferJobs/OPI` specifically. For all other transfer types, this name must not start with `transferJobs/OPI`. Non-PosixFilesystem example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` PosixFilesystem example: `"transferJobs/OPI^[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Applications must not rely on the enforcement of naming requirements involving OPI. Invalid job names fail with an INVALID_ARGUMENT error.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Notification configuration. This is not supported for transfers involving PosixFilesystem.
        /// </summary>
        public readonly Outputs.NotificationConfigResponse NotificationConfig;
        /// <summary>
        /// The ID of the Google Cloud project that owns the job.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job never executes a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
        /// </summary>
        public readonly Outputs.ScheduleResponse Schedule;
        /// <summary>
        /// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Transfer specification.
        /// </summary>
        public readonly Outputs.TransferSpecResponse TransferSpec;

        [OutputConstructor]
        private GetTransferJobResult(
            string creationTime,

            string deletionTime,

            string description,

            string lastModificationTime,

            string latestOperationName,

            Outputs.LoggingConfigResponse loggingConfig,

            string name,

            Outputs.NotificationConfigResponse notificationConfig,

            string project,

            Outputs.ScheduleResponse schedule,

            string status,

            Outputs.TransferSpecResponse transferSpec)
        {
            CreationTime = creationTime;
            DeletionTime = deletionTime;
            Description = description;
            LastModificationTime = lastModificationTime;
            LatestOperationName = latestOperationName;
            LoggingConfig = loggingConfig;
            Name = name;
            NotificationConfig = notificationConfig;
            Project = project;
            Schedule = schedule;
            Status = status;
            TransferSpec = transferSpec;
        }
    }
}
