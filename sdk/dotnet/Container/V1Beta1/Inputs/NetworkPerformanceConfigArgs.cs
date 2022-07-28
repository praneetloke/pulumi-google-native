// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration of all network bandwidth tiers
    /// </summary>
    public sealed class NetworkPerformanceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the network bandwidth tier for the NodePool for traffic to external/public IP addresses.
        /// </summary>
        [Input("externalIpEgressBandwidthTier")]
        public Input<Pulumi.GoogleNative.Container.V1Beta1.NetworkPerformanceConfigExternalIpEgressBandwidthTier>? ExternalIpEgressBandwidthTier { get; set; }

        /// <summary>
        /// Specifies the total network bandwidth tier for the NodePool.
        /// </summary>
        [Input("totalEgressBandwidthTier")]
        public Input<Pulumi.GoogleNative.Container.V1Beta1.NetworkPerformanceConfigTotalEgressBandwidthTier>? TotalEgressBandwidthTier { get; set; }

        public NetworkPerformanceConfigArgs()
        {
        }
        public static new NetworkPerformanceConfigArgs Empty => new NetworkPerformanceConfigArgs();
    }
}
