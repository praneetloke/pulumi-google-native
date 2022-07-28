// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Inputs
{

    /// <summary>
    /// Options for enabling sharding.
    /// </summary>
    public sealed class ShardingOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Shards test cases into the specified groups of packages, classes, and/or methods.
        /// </summary>
        [Input("manualSharding")]
        public Input<Inputs.ManualShardingArgs>? ManualSharding { get; set; }

        /// <summary>
        /// Uniformly shards test cases given a total number of shards.
        /// </summary>
        [Input("uniformSharding")]
        public Input<Inputs.UniformShardingArgs>? UniformSharding { get; set; }

        public ShardingOptionArgs()
        {
        }
        public static new ShardingOptionArgs Empty => new ShardingOptionArgs();
    }
}
