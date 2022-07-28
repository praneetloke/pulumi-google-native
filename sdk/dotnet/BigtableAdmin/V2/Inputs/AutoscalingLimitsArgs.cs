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
    /// Limits for the number of nodes a Cluster can autoscale up/down to.
    /// </summary>
    public sealed class AutoscalingLimitsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of nodes to scale up to.
        /// </summary>
        [Input("maxServeNodes", required: true)]
        public Input<int> MaxServeNodes { get; set; } = null!;

        /// <summary>
        /// Minimum number of nodes to scale down to.
        /// </summary>
        [Input("minServeNodes", required: true)]
        public Input<int> MinServeNodes { get; set; } = null!;

        public AutoscalingLimitsArgs()
        {
        }
        public static new AutoscalingLimitsArgs Empty => new AutoscalingLimitsArgs();
    }
}
