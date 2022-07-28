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
    /// Configuration that allows for slower scale in so that even if Autoscaler recommends an abrupt scale in of a MIG, it will be throttled as specified by the parameters below.
    /// </summary>
    public sealed class AutoscalingPolicyScaleInControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum allowed number (or %) of VMs that can be deducted from the peak recommendation during the window autoscaler looks at when computing recommendations. Possibly all these VMs can be deleted at once so user service needs to be prepared to lose that many VMs in one step.
        /// </summary>
        [Input("maxScaledInReplicas")]
        public Input<Inputs.FixedOrPercentArgs>? MaxScaledInReplicas { get; set; }

        /// <summary>
        /// How far back autoscaling looks when computing recommendations to include directives regarding slower scale in, as described above.
        /// </summary>
        [Input("timeWindowSec")]
        public Input<int>? TimeWindowSec { get; set; }

        public AutoscalingPolicyScaleInControlArgs()
        {
        }
        public static new AutoscalingPolicyScaleInControlArgs Empty => new AutoscalingPolicyScaleInControlArgs();
    }
}
