// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Inputs
{

    /// <summary>
    /// A floating point interval.
    /// </summary>
    public sealed class GoogleCloudRetailV2alphaIntervalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Exclusive upper bound.
        /// </summary>
        [Input("exclusiveMaximum")]
        public Input<double>? ExclusiveMaximum { get; set; }

        /// <summary>
        /// Exclusive lower bound.
        /// </summary>
        [Input("exclusiveMinimum")]
        public Input<double>? ExclusiveMinimum { get; set; }

        /// <summary>
        /// Inclusive upper bound.
        /// </summary>
        [Input("maximum")]
        public Input<double>? Maximum { get; set; }

        /// <summary>
        /// Inclusive lower bound.
        /// </summary>
        [Input("minimum")]
        public Input<double>? Minimum { get; set; }

        public GoogleCloudRetailV2alphaIntervalArgs()
        {
        }
        public static new GoogleCloudRetailV2alphaIntervalArgs Empty => new GoogleCloudRetailV2alphaIntervalArgs();
    }
}
