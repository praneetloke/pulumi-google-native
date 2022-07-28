// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Encapsulates numeric value that can be either absolute or relative.
    /// </summary>
    public sealed class FixedOrPercentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a fixed number of VM instances. This must be a positive integer.
        /// </summary>
        [Input("fixed")]
        public Input<int>? Fixed { get; set; }

        /// <summary>
        /// Specifies a percentage of instances between 0 to 100%, inclusive. For example, specify 80 for 80%.
        /// </summary>
        [Input("percent")]
        public Input<int>? Percent { get; set; }

        public FixedOrPercentArgs()
        {
        }
        public static new FixedOrPercentArgs Empty => new FixedOrPercentArgs();
    }
}
