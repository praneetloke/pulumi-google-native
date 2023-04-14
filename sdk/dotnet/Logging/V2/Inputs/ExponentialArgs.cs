// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2.Inputs
{

    /// <summary>
    /// Specifies an exponential sequence of buckets that have a width that is proportional to the value of the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.There are num_finite_buckets + 2 (= N) buckets. Bucket i has the following boundaries:Upper bound (0 &lt;= i &lt; N-1): scale * (growth_factor ^ i).Lower bound (1 &lt;= i &lt; N): scale * (growth_factor ^ (i - 1)).
    /// </summary>
    public sealed class ExponentialArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Must be greater than 1.
        /// </summary>
        [Input("growthFactor")]
        public Input<double>? GrowthFactor { get; set; }

        /// <summary>
        /// Must be greater than 0.
        /// </summary>
        [Input("numFiniteBuckets")]
        public Input<int>? NumFiniteBuckets { get; set; }

        /// <summary>
        /// Must be greater than 0.
        /// </summary>
        [Input("scale")]
        public Input<double>? Scale { get; set; }

        public ExponentialArgs()
        {
        }
        public static new ExponentialArgs Empty => new ExponentialArgs();
    }
}
