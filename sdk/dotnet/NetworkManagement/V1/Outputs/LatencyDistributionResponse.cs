// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1.Outputs
{

    /// <summary>
    /// Describes measured latency distribution.
    /// </summary>
    [OutputType]
    public sealed class LatencyDistributionResponse
    {
        /// <summary>
        /// Representative latency percentiles.
        /// </summary>
        public readonly ImmutableArray<Outputs.LatencyPercentileResponse> LatencyPercentiles;

        [OutputConstructor]
        private LatencyDistributionResponse(ImmutableArray<Outputs.LatencyPercentileResponse> latencyPercentiles)
        {
            LatencyPercentiles = latencyPercentiles;
        }
    }
}