// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Inputs
{

    /// <summary>
    /// The Autoscaling targets for a Cluster. These determine the recommended nodes.
    /// </summary>
    public sealed class AutoscalingTargetsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cpu utilization that the Autoscaler should be trying to achieve. This number is on a scale from 0 (no utilization) to 100 (total utilization), and is limited between 10 and 80, otherwise it will return INVALID_ARGUMENT error.
        /// </summary>
        [Input("cpuUtilizationPercent")]
        public Input<int>? CpuUtilizationPercent { get; set; }

        /// <summary>
        /// The storage utilization that the Autoscaler should be trying to achieve. This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD cluster and between 8192 (8TiB) and 16384 (16TiB) for an HDD cluster, otherwise it will return INVALID_ARGUMENT error. If this value is set to 0, it will be treated as if it were set to the default value: 2560 for SSD, 8192 for HDD.
        /// </summary>
        [Input("storageUtilizationGibPerNode")]
        public Input<int>? StorageUtilizationGibPerNode { get; set; }

        public AutoscalingTargetsArgs()
        {
        }
        public static new AutoscalingTargetsArgs Empty => new AutoscalingTargetsArgs();
    }
}
