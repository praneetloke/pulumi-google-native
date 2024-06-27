// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Outputs
{

    /// <summary>
    /// State indicating an ACM's progress syncing configurations to a cluster
    /// </summary>
    [OutputType]
    public sealed class ConfigManagementSyncStateResponse
    {
        /// <summary>
        /// Sync status code
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// A list of errors resulting from problematic configs. This list will be truncated after 100 errors, although it is unlikely for that many errors to simultaneously exist.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigManagementSyncErrorResponse> Errors;
        /// <summary>
        /// Token indicating the state of the importer.
        /// </summary>
        public readonly string ImportToken;
        /// <summary>
        /// Deprecated: use last_sync_time instead. Timestamp of when ACM last successfully synced the repo The time format is specified in https://golang.org/pkg/time/#Time.String
        /// </summary>
        public readonly string LastSync;
        /// <summary>
        /// Timestamp type of when ACM last successfully synced the repo
        /// </summary>
        public readonly string LastSyncTime;
        /// <summary>
        /// Token indicating the state of the repo.
        /// </summary>
        public readonly string SourceToken;
        /// <summary>
        /// Token indicating the state of the syncer.
        /// </summary>
        public readonly string SyncToken;

        [OutputConstructor]
        private ConfigManagementSyncStateResponse(
            string code,

            ImmutableArray<Outputs.ConfigManagementSyncErrorResponse> errors,

            string importToken,

            string lastSync,

            string lastSyncTime,

            string sourceToken,

            string syncToken)
        {
            Code = code;
            Errors = errors;
            ImportToken = importToken;
            LastSync = lastSync;
            LastSyncTime = lastSyncTime;
            SourceToken = sourceToken;
            SyncToken = syncToken;
        }
    }
}