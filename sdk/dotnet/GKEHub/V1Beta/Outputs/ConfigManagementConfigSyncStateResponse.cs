// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Beta.Outputs
{

    /// <summary>
    /// State information for ConfigSync
    /// </summary>
    [OutputType]
    public sealed class ConfigManagementConfigSyncStateResponse
    {
        /// <summary>
        /// Information about the deployment of ConfigSync, including the version of the various Pods deployed
        /// </summary>
        public readonly Outputs.ConfigManagementConfigSyncDeploymentStateResponse DeploymentState;
        /// <summary>
        /// Errors pertaining to the installation of Config Sync.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigManagementConfigSyncErrorResponse> Errors;
        /// <summary>
        /// The state of ConfigSync's process to sync configs to a cluster
        /// </summary>
        public readonly Outputs.ConfigManagementSyncStateResponse SyncState;
        /// <summary>
        /// The version of ConfigSync deployed
        /// </summary>
        public readonly Outputs.ConfigManagementConfigSyncVersionResponse Version;

        [OutputConstructor]
        private ConfigManagementConfigSyncStateResponse(
            Outputs.ConfigManagementConfigSyncDeploymentStateResponse deploymentState,

            ImmutableArray<Outputs.ConfigManagementConfigSyncErrorResponse> errors,

            Outputs.ConfigManagementSyncStateResponse syncState,

            Outputs.ConfigManagementConfigSyncVersionResponse version)
        {
            DeploymentState = deploymentState;
            Errors = errors;
            SyncState = syncState;
            Version = version;
        }
    }
}
