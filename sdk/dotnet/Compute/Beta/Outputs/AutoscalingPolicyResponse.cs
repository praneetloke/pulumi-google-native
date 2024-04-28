// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Cloud Autoscaler policy.
    /// </summary>
    [OutputType]
    public sealed class AutoscalingPolicyResponse
    {
        /// <summary>
        /// The number of seconds that your application takes to initialize on a VM instance. This is referred to as the [initialization period](/compute/docs/autoscaler#cool_down_period). Specifying an accurate initialization period improves autoscaler decisions. For example, when scaling out, the autoscaler ignores data from VMs that are still initializing because those VMs might not yet represent normal usage of your application. The default initialization period is 60 seconds. Initialization periods might vary because of numerous factors. We recommend that you test how long your application takes to initialize. To do this, create a VM and time your application's startup process.
        /// </summary>
        public readonly int CoolDownPeriodSec;
        /// <summary>
        /// Defines the CPU utilization policy that allows the autoscaler to scale based on the average CPU utilization of a managed instance group.
        /// </summary>
        public readonly Outputs.AutoscalingPolicyCpuUtilizationResponse CpuUtilization;
        /// <summary>
        /// Configuration parameters of autoscaling based on a custom metric.
        /// </summary>
        public readonly ImmutableArray<Outputs.AutoscalingPolicyCustomMetricUtilizationResponse> CustomMetricUtilizations;
        /// <summary>
        /// Configuration parameters of autoscaling based on load balancer.
        /// </summary>
        public readonly Outputs.AutoscalingPolicyLoadBalancingUtilizationResponse LoadBalancingUtilization;
        /// <summary>
        /// The maximum number of instances that the autoscaler can scale out to. This is required when creating or updating an autoscaler. The maximum number of replicas must not be lower than minimal number of replicas.
        /// </summary>
        public readonly int MaxNumReplicas;
        /// <summary>
        /// The minimum number of replicas that the autoscaler can scale in to. This cannot be less than 0. If not provided, autoscaler chooses a default value depending on maximum number of instances allowed.
        /// </summary>
        public readonly int MinNumReplicas;
        /// <summary>
        /// Defines the operating mode for this policy. The following modes are available: - OFF: Disables the autoscaler but maintains its configuration. - ONLY_SCALE_OUT: Restricts the autoscaler to add VM instances only. - ON: Enables all autoscaler activities according to its policy. For more information, see "Turning off or restricting an autoscaler"
        /// </summary>
        public readonly string Mode;
        public readonly Outputs.AutoscalingPolicyScaleDownControlResponse ScaleDownControl;
        public readonly Outputs.AutoscalingPolicyScaleInControlResponse ScaleInControl;
        /// <summary>
        /// Scaling schedules defined for an autoscaler. Multiple schedules can be set on an autoscaler, and they can overlap. During overlapping periods the greatest min_required_replicas of all scaling schedules is applied. Up to 128 scaling schedules are allowed.
        /// </summary>
        public readonly Outputs.AutoscalingPolicyScalingScheduleResponse ScalingSchedules;

        [OutputConstructor]
        private AutoscalingPolicyResponse(
            int coolDownPeriodSec,

            Outputs.AutoscalingPolicyCpuUtilizationResponse cpuUtilization,

            ImmutableArray<Outputs.AutoscalingPolicyCustomMetricUtilizationResponse> customMetricUtilizations,

            Outputs.AutoscalingPolicyLoadBalancingUtilizationResponse loadBalancingUtilization,

            int maxNumReplicas,

            int minNumReplicas,

            string mode,

            Outputs.AutoscalingPolicyScaleDownControlResponse scaleDownControl,

            Outputs.AutoscalingPolicyScaleInControlResponse scaleInControl,

            Outputs.AutoscalingPolicyScalingScheduleResponse scalingSchedules)
        {
            CoolDownPeriodSec = coolDownPeriodSec;
            CpuUtilization = cpuUtilization;
            CustomMetricUtilizations = customMetricUtilizations;
            LoadBalancingUtilization = loadBalancingUtilization;
            MaxNumReplicas = maxNumReplicas;
            MinNumReplicas = minNumReplicas;
            Mode = mode;
            ScaleDownControl = scaleDownControl;
            ScaleInControl = scaleInControl;
            ScalingSchedules = scalingSchedules;
        }
    }
}
