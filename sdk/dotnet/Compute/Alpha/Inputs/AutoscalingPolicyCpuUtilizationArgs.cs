// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// CPU utilization policy.
    /// </summary>
    public sealed class AutoscalingPolicyCpuUtilizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether predictive autoscaling based on CPU metric is enabled. Valid values are: * NONE (default). No predictive method is used. The autoscaler scales the group to meet current demand based on real-time metrics. * OPTIMIZE_AVAILABILITY. Predictive autoscaling improves availability by monitoring daily and weekly load patterns and scaling out ahead of anticipated demand.
        /// </summary>
        [Input("predictiveMethod")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.AutoscalingPolicyCpuUtilizationPredictiveMethod>? PredictiveMethod { get; set; }

        /// <summary>
        /// The target CPU utilization that the autoscaler maintains. Must be a float value in the range (0, 1]. If not specified, the default is 0.6. If the CPU level is below the target utilization, the autoscaler scales in the number of instances until it reaches the minimum number of instances you specified or until the average CPU of your instances reaches the target utilization. If the average CPU is above the target utilization, the autoscaler scales out until it reaches the maximum number of instances you specified or until the average utilization reaches the target utilization.
        /// </summary>
        [Input("utilizationTarget")]
        public Input<double>? UtilizationTarget { get; set; }

        public AutoscalingPolicyCpuUtilizationArgs()
        {
        }
        public static new AutoscalingPolicyCpuUtilizationArgs Empty => new AutoscalingPolicyCpuUtilizationArgs();
    }
}
