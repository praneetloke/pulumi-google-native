// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Inputs
{

    /// <summary>
    /// Defines a threshold for categorizing time series values.
    /// </summary>
    public sealed class ThresholdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The state color for this threshold. Color is not allowed in a XyChart.
        /// </summary>
        [Input("color")]
        public Input<Pulumi.GoogleNative.Monitoring.V1.ThresholdColor>? Color { get; set; }

        /// <summary>
        /// The direction for the current threshold. Direction is not allowed in a XyChart.
        /// </summary>
        [Input("direction")]
        public Input<Pulumi.GoogleNative.Monitoring.V1.ThresholdDirection>? Direction { get; set; }

        /// <summary>
        /// A label for the threshold.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The target axis to use for plotting the threshold. Target axis is not allowed in a Scorecard.
        /// </summary>
        [Input("targetAxis")]
        public Input<Pulumi.GoogleNative.Monitoring.V1.ThresholdTargetAxis>? TargetAxis { get; set; }

        /// <summary>
        /// The value of the threshold. The value should be defined in the native scale of the metric.
        /// </summary>
        [Input("value")]
        public Input<double>? Value { get; set; }

        public ThresholdArgs()
        {
        }
        public static new ThresholdArgs Empty => new ThresholdArgs();
    }
}
