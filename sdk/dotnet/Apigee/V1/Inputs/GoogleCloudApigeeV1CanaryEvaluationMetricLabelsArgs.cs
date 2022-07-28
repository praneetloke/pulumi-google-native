// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// Labels that can be used to filter Apigee metrics.
    /// </summary>
    public sealed class GoogleCloudApigeeV1CanaryEvaluationMetricLabelsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The environment ID associated with the metrics.
        /// </summary>
        [Input("env")]
        public Input<string>? Env { get; set; }

        /// <summary>
        /// The instance ID associated with the metrics. In Apigee Hybrid, the value is configured during installation.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The location associated with the metrics.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        public GoogleCloudApigeeV1CanaryEvaluationMetricLabelsArgs()
        {
        }
        public static new GoogleCloudApigeeV1CanaryEvaluationMetricLabelsArgs Empty => new GoogleCloudApigeeV1CanaryEvaluationMetricLabelsArgs();
    }
}
