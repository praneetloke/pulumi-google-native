// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Outputs
{

    /// <summary>
    /// Online serving scaling configuration. If min_node_count and max_node_count are set to the same value, the cluster will be configured with the fixed number of node (no auto-scaling).
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1FeaturestoreOnlineServingConfigScalingResponse
    {
        /// <summary>
        /// Optional. The cpu utilization that the Autoscaler should be trying to achieve. This number is on a scale from 0 (no utilization) to 100 (total utilization), and is limited between 10 and 80. When a cluster's CPU utilization exceeds the target that you have set, Bigtable immediately adds nodes to the cluster. When CPU utilization is substantially lower than the target, Bigtable removes nodes. If not set or set to 0, default to 50.
        /// </summary>
        public readonly int CpuUtilizationTarget;
        /// <summary>
        /// The maximum number of nodes to scale up to. Must be greater than min_node_count, and less than or equal to 10 times of 'min_node_count'.
        /// </summary>
        public readonly int MaxNodeCount;
        /// <summary>
        /// The minimum number of nodes to scale down to. Must be greater than or equal to 1.
        /// </summary>
        public readonly int MinNodeCount;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1FeaturestoreOnlineServingConfigScalingResponse(
            int cpuUtilizationTarget,

            int maxNodeCount,

            int minNodeCount)
        {
            CpuUtilizationTarget = cpuUtilizationTarget;
            MaxNodeCount = maxNodeCount;
            MinNodeCount = minNodeCount;
        }
    }
}