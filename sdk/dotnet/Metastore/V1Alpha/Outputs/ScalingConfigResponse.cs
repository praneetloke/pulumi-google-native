// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Alpha.Outputs
{

    /// <summary>
    /// Represents the scaling configuration of a metastore service.
    /// </summary>
    [OutputType]
    public sealed class ScalingConfigResponse
    {
        /// <summary>
        /// An enum of readable instance sizes, with each instance size mapping to a float value (e.g. InstanceSize.EXTRA_SMALL = scaling_factor(0.1))
        /// </summary>
        public readonly string InstanceSize;
        /// <summary>
        /// Scaling factor, increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
        /// </summary>
        public readonly double ScalingFactor;

        [OutputConstructor]
        private ScalingConfigResponse(
            string instanceSize,

            double scalingFactor)
        {
            InstanceSize = instanceSize;
            ScalingFactor = scalingFactor;
        }
    }
}
