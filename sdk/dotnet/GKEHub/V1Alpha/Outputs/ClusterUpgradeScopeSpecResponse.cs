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
    /// **ClusterUpgrade**: The configuration for the scope-level ClusterUpgrade feature.
    /// </summary>
    [OutputType]
    public sealed class ClusterUpgradeScopeSpecResponse
    {
        /// <summary>
        /// Allow users to override some properties of each GKE upgrade.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterUpgradeGKEUpgradeOverrideResponse> GkeUpgradeOverrides;
        /// <summary>
        /// Post conditions to evaluate to mark an upgrade COMPLETE. Required.
        /// </summary>
        public readonly Outputs.ClusterUpgradePostConditionsResponse PostConditions;
        /// <summary>
        /// This scope consumes upgrades that have COMPLETE status code in the upstream scopes. See UpgradeStatus.Code for code definitions. The scope name should be in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. This is defined as repeated for future proof reasons. Initial implementation will enforce at most one upstream scope.
        /// </summary>
        public readonly ImmutableArray<string> UpstreamScopes;

        [OutputConstructor]
        private ClusterUpgradeScopeSpecResponse(
            ImmutableArray<Outputs.ClusterUpgradeGKEUpgradeOverrideResponse> gkeUpgradeOverrides,

            Outputs.ClusterUpgradePostConditionsResponse postConditions,

            ImmutableArray<string> upstreamScopes)
        {
            GkeUpgradeOverrides = gkeUpgradeOverrides;
            PostConditions = postConditions;
            UpstreamScopes = upstreamScopes;
        }
    }
}
