// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// Configuration options for the master authorized networks feature. Enabled master authorized networks will disallow all external traffic to access Kubernetes master through HTTPS except traffic from the given CIDR blocks, Google Compute Engine Public IPs and Google Prod IPs.
    /// </summary>
    [OutputType]
    public sealed class MasterAuthorizedNetworksConfigResponse
    {
        /// <summary>
        /// cidr_blocks define up to 50 external networks that could access Kubernetes master through HTTPS.
        /// </summary>
        public readonly ImmutableArray<Outputs.CidrBlockResponse> CidrBlocks;
        /// <summary>
        /// Whether or not master authorized networks is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Whether master is accessbile via Google Compute Engine Public IP addresses.
        /// </summary>
        public readonly bool GcpPublicCidrsAccessEnabled;

        [OutputConstructor]
        private MasterAuthorizedNetworksConfigResponse(
            ImmutableArray<Outputs.CidrBlockResponse> cidrBlocks,

            bool enabled,

            bool gcpPublicCidrsAccessEnabled)
        {
            CidrBlocks = cidrBlocks;
            Enabled = enabled;
            GcpPublicCidrsAccessEnabled = gcpPublicCidrsAccessEnabled;
        }
    }
}
