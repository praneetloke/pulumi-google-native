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
    /// **Metering**: Per-Membership Feature State.
    /// </summary>
    [OutputType]
    public sealed class MeteringMembershipStateResponse
    {
        /// <summary>
        /// The time stamp of the most recent measurement of the number of vCPUs in the cluster.
        /// </summary>
        public readonly string LastMeasurementTime;
        /// <summary>
        /// The vCPUs capacity in the cluster according to the most recent measurement (1/1000 precision).
        /// </summary>
        public readonly double PreciseLastMeasuredClusterVcpuCapacity;

        [OutputConstructor]
        private MeteringMembershipStateResponse(
            string lastMeasurementTime,

            double preciseLastMeasuredClusterVcpuCapacity)
        {
            LastMeasurementTime = lastMeasurementTime;
            PreciseLastMeasuredClusterVcpuCapacity = preciseLastMeasuredClusterVcpuCapacity;
        }
    }
}
