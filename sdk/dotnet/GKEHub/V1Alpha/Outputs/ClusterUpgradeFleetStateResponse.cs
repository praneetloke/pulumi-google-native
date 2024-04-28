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
    /// **ClusterUpgrade**: The state for the fleet-level ClusterUpgrade feature.
    /// </summary>
    [OutputType]
    public sealed class ClusterUpgradeFleetStateResponse
    {
        /// <summary>
        /// This fleets whose upstream_fleets contain the current fleet. The fleet name should be either fleet project number or id.
        /// </summary>
        public readonly ImmutableArray<string> DownstreamFleets;
        /// <summary>
        /// Feature state for GKE clusters.
        /// </summary>
        public readonly Outputs.ClusterUpgradeGKEUpgradeFeatureStateResponse GkeState;
        /// <summary>
        /// A list of memberships ignored by the feature. For example, manually upgraded clusters can be ignored if they are newer than the default versions of its release channel. The membership resource is in the format: `projects/{p}/locations/{l}/membership/{m}`.
        /// </summary>
        public readonly Outputs.ClusterUpgradeIgnoredMembershipResponse Ignored;

        [OutputConstructor]
        private ClusterUpgradeFleetStateResponse(
            ImmutableArray<string> downstreamFleets,

            Outputs.ClusterUpgradeGKEUpgradeFeatureStateResponse gkeState,

            Outputs.ClusterUpgradeIgnoredMembershipResponse ignored)
        {
            DownstreamFleets = downstreamFleets;
            GkeState = gkeState;
            Ignored = ignored;
        }
    }
}
