// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Beta.Inputs
{

    /// <summary>
    /// Represents the scaling configuration of a metastore service.
    /// </summary>
    public sealed class ScalingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An enum of readable instance sizes, with each instance size mapping to a float value (e.g. InstanceSize.EXTRA_SMALL = scaling_factor(0.1))
        /// </summary>
        [Input("instanceSize")]
        public Input<Pulumi.GoogleNative.Metastore.V1Beta.ScalingConfigInstanceSize>? InstanceSize { get; set; }

        /// <summary>
        /// Scaling factor, increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
        /// </summary>
        [Input("scalingFactor")]
        public Input<double>? ScalingFactor { get; set; }

        public ScalingConfigArgs()
        {
        }
        public static new ScalingConfigArgs Empty => new ScalingConfigArgs();
    }
}
