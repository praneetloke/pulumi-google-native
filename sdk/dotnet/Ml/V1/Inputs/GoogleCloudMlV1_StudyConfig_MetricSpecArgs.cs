// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Inputs
{

    /// <summary>
    /// Represents a metric to optimize.
    /// </summary>
    public sealed class GoogleCloudMlV1_StudyConfig_MetricSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The optimization goal of the metric.
        /// </summary>
        [Input("goal", required: true)]
        public Input<Pulumi.GoogleNative.Ml.V1.GoogleCloudMlV1_StudyConfig_MetricSpecGoal> Goal { get; set; } = null!;

        /// <summary>
        /// The name of the metric.
        /// </summary>
        [Input("metric", required: true)]
        public Input<string> Metric { get; set; } = null!;

        public GoogleCloudMlV1_StudyConfig_MetricSpecArgs()
        {
        }
        public static new GoogleCloudMlV1_StudyConfig_MetricSpecArgs Empty => new GoogleCloudMlV1_StudyConfig_MetricSpecArgs();
    }
}
