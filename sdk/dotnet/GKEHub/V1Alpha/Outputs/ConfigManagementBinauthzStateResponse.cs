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
    /// State for Binauthz
    /// </summary>
    [OutputType]
    public sealed class ConfigManagementBinauthzStateResponse
    {
        /// <summary>
        /// The version of binauthz that is installed.
        /// </summary>
        public readonly Outputs.ConfigManagementBinauthzVersionResponse Version;
        /// <summary>
        /// The state of the binauthz webhook.
        /// </summary>
        public readonly string Webhook;

        [OutputConstructor]
        private ConfigManagementBinauthzStateResponse(
            Outputs.ConfigManagementBinauthzVersionResponse version,

            string webhook)
        {
            Version = version;
            Webhook = webhook;
        }
    }
}
