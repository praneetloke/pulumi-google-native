// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// State information for an ACM's Operator
    /// </summary>
    [OutputType]
    public sealed class ConfigManagementOperatorStateResponse
    {
        /// <summary>
        /// The state of the Operator's deployment
        /// </summary>
        public readonly string DeploymentState;
        /// <summary>
        /// Install errors.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigManagementInstallErrorResponse> Errors;
        /// <summary>
        /// The semenatic version number of the operator
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private ConfigManagementOperatorStateResponse(
            string deploymentState,

            ImmutableArray<Outputs.ConfigManagementInstallErrorResponse> errors,

            string version)
        {
            DeploymentState = deploymentState;
            Errors = errors;
            Version = version;
        }
    }
}