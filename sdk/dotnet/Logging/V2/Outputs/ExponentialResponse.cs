// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2.Outputs
{

    /// <summary>
    /// Specifies an exponential sequence of buckets that have a width that is proportional to the value of the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.There are num_finite_buckets + 2 (= N) buckets. Bucket i has the following boundaries:Upper bound (0 &lt;= i &lt; N-1): scale * (growth_factor ^ i). Lower bound (1 &lt;= i &lt; N): scale * (growth_factor ^ (i - 1)).
    /// </summary>
    [OutputType]
    public sealed class ExponentialResponse
    {
        /// <summary>
        /// Must be greater than 1.
        /// </summary>
        public readonly double GrowthFactor;
        /// <summary>
        /// Must be greater than 0.
        /// </summary>
        public readonly int NumFiniteBuckets;
        /// <summary>
        /// Must be greater than 0.
        /// </summary>
        public readonly double Scale;

        [OutputConstructor]
        private ExponentialResponse(
            double growthFactor,

            int numFiniteBuckets,

            double scale)
        {
            GrowthFactor = growthFactor;
            NumFiniteBuckets = numFiniteBuckets;
            Scale = scale;
        }
    }
}