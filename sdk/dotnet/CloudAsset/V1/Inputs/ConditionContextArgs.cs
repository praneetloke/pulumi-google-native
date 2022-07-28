// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudAsset.V1.Inputs
{

    /// <summary>
    /// The IAM conditions context.
    /// </summary>
    public sealed class ConditionContextArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hypothetical access timestamp to evaluate IAM conditions. Note that this value must not be earlier than the current time; otherwise, an INVALID_ARGUMENT error will be returned.
        /// </summary>
        [Input("accessTime")]
        public Input<string>? AccessTime { get; set; }

        public ConditionContextArgs()
        {
        }
        public static new ConditionContextArgs Empty => new ConditionContextArgs();
    }
}
