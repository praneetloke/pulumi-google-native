// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1
{
    /// <summary>
    /// Creates a transfer job that runs periodically.
    /// </summary>
    [GoogleNativeResourceType("google-native:storagetransfer/v1:TransferJob")]
    public partial class TransferJob : Pulumi.CustomResource
    {
        /// <summary>
        /// The time that the transfer job was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The time that the transfer job was deleted.
        /// </summary>
        [Output("deletionTime")]
        public Output<string> DeletionTime { get; private set; } = null!;

        /// <summary>
        /// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The time that the transfer job was last modified.
        /// </summary>
        [Output("lastModificationTime")]
        public Output<string> LastModificationTime { get; private set; } = null!;

        /// <summary>
        /// The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
        /// </summary>
        [Output("latestOperationName")]
        public Output<string> LatestOperationName { get; private set; } = null!;

        /// <summary>
        /// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Notification configuration.
        /// </summary>
        [Output("notificationConfig")]
        public Output<Outputs.NotificationConfigResponse> NotificationConfig { get; private set; } = null!;

        /// <summary>
        /// The ID of the Google Cloud Platform Project that owns the job.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.ScheduleResponse> Schedule { get; private set; } = null!;

        /// <summary>
        /// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Transfer specification.
        /// </summary>
        [Output("transferSpec")]
        public Output<Outputs.TransferSpecResponse> TransferSpec { get; private set; } = null!;


        /// <summary>
        /// Create a TransferJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransferJob(string name, TransferJobArgs args, CustomResourceOptions? options = null)
            : base("google-native:storagetransfer/v1:TransferJob", name, args ?? new TransferJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransferJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:storagetransfer/v1:TransferJob", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TransferJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransferJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TransferJob(name, id, options);
        }
    }

    public sealed class TransferJobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description provided by the user for the job. Its max length is 1024 bytes when Unicode-encoded.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the most recently started TransferOperation of this JobConfig. Present if and only if at least one TransferOperation has been created for this JobConfig.
        /// </summary>
        [Input("latestOperationName")]
        public Input<string>? LatestOperationName { get; set; }

        /// <summary>
        /// A unique name (within the transfer project) assigned when the job is created. If this field is empty in a CreateTransferJobRequest, Storage Transfer Service will assign a unique name. Otherwise, the specified name is used as the unique name for this job. If the specified name is in use by a job, the creation request fails with an ALREADY_EXISTS error. This name must start with `"transferJobs/"` prefix and end with a letter or a number, and should be no more than 128 characters. This name must not start with 'transferJobs/OPI'. 'transferJobs/OPI' is a reserved prefix. Example: `"transferJobs/^(?!OPI)[A-Za-z0-9-._~]*[A-Za-z0-9]$"` Invalid job names will fail with an INVALID_ARGUMENT error.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Notification configuration.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.NotificationConfigArgs>? NotificationConfig { get; set; }

        /// <summary>
        /// The ID of the Google Cloud Platform Project that owns the job.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Specifies schedule for the transfer job. This is an optional field. When the field is not set, the job will never execute a transfer, unless you invoke RunTransferJob or update the job to have a non-empty schedule.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// Status of the job. This value MUST be specified for `CreateTransferJobRequests`. **Note:** The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("transferJobsId", required: true)]
        public Input<string> TransferJobsId { get; set; } = null!;

        /// <summary>
        /// Transfer specification.
        /// </summary>
        [Input("transferSpec")]
        public Input<Inputs.TransferSpecArgs>? TransferSpec { get; set; }

        public TransferJobArgs()
        {
        }
    }
}
