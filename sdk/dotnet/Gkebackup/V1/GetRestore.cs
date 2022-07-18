// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkebackup.V1
{
    public static class GetRestore
    {
        /// <summary>
        /// Retrieves the details of a single Restore.
        /// </summary>
        public static Task<GetRestoreResult> InvokeAsync(GetRestoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRestoreResult>("google-native:gkebackup/v1:getRestore", args ?? new GetRestoreArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the details of a single Restore.
        /// </summary>
        public static Output<GetRestoreResult> Invoke(GetRestoreInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRestoreResult>("google-native:gkebackup/v1:getRestore", args ?? new GetRestoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRestoreArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("restoreId", required: true)]
        public string RestoreId { get; set; } = null!;

        [Input("restorePlanId", required: true)]
        public string RestorePlanId { get; set; } = null!;

        public GetRestoreArgs()
        {
        }
    }

    public sealed class GetRestoreInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("restoreId", required: true)]
        public Input<string> RestoreId { get; set; } = null!;

        [Input("restorePlanId", required: true)]
        public Input<string> RestorePlanId { get; set; } = null!;

        public GetRestoreInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRestoreResult
    {
        /// <summary>
        /// Immutable. A reference to the Backup used as the source from which this Restore will restore. Note that this Backup must be a sub-resource of the RestorePlan's backup_plan. Format: `projects/*/locations/*/backupPlans/*/backups/*`.
        /// </summary>
        public readonly string Backup;
        /// <summary>
        /// The target cluster into which this Restore will restore data. Valid formats: - `projects/*/locations/*/clusters/*` - `projects/*/zones/*/clusters/*` Inherited from parent RestorePlan's cluster value.
        /// </summary>
        public readonly string Cluster;
        /// <summary>
        /// Timestamp of when the restore operation completed.
        /// </summary>
        public readonly string CompleteTime;
        /// <summary>
        /// The timestamp when this Restore resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// User specified descriptive string for this Restore.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a restore from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform restore updates in order to avoid race conditions: An `etag` is returned in the response to `GetRestore`, and systems are expected to put that etag in the request to `UpdateRestore` or `DeleteRestore` to ensure that their change will be applied to the same version of the resource.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// A set of custom labels supplied by user.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The full name of the Restore resource. Format: `projects/*/locations/*/restorePlans/*/restores/*`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of resources excluded during the restore execution.
        /// </summary>
        public readonly int ResourcesExcludedCount;
        /// <summary>
        /// Number of resources that failed to be restored during the restore execution.
        /// </summary>
        public readonly int ResourcesFailedCount;
        /// <summary>
        /// Number of resources restored during the restore execution.
        /// </summary>
        public readonly int ResourcesRestoredCount;
        /// <summary>
        /// Configuration of the Restore. Inherited from parent RestorePlan's restore_config.
        /// </summary>
        public readonly Outputs.RestoreConfigResponse RestoreConfig;
        /// <summary>
        /// The current state of the Restore.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Human-readable description of why the Restore is in its current state.
        /// </summary>
        public readonly string StateReason;
        /// <summary>
        /// Server generated global unique identifier of [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The timestamp when this Restore resource was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Number of volumes restored during the restore execution.
        /// </summary>
        public readonly int VolumesRestoredCount;

        [OutputConstructor]
        private GetRestoreResult(
            string backup,

            string cluster,

            string completeTime,

            string createTime,

            string description,

            string etag,

            ImmutableDictionary<string, string> labels,

            string name,

            int resourcesExcludedCount,

            int resourcesFailedCount,

            int resourcesRestoredCount,

            Outputs.RestoreConfigResponse restoreConfig,

            string state,

            string stateReason,

            string uid,

            string updateTime,

            int volumesRestoredCount)
        {
            Backup = backup;
            Cluster = cluster;
            CompleteTime = completeTime;
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            Labels = labels;
            Name = name;
            ResourcesExcludedCount = resourcesExcludedCount;
            ResourcesFailedCount = resourcesFailedCount;
            ResourcesRestoredCount = resourcesRestoredCount;
            RestoreConfig = restoreConfig;
            State = state;
            StateReason = stateReason;
            Uid = uid;
            UpdateTime = updateTime;
            VolumesRestoredCount = volumesRestoredCount;
        }
    }
}
