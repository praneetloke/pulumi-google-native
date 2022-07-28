// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RuntimeConfig.V1Beta1.Inputs
{

    /// <summary>
    /// A Cardinality condition for the Waiter resource. A cardinality condition is met when the number of variables under a specified path prefix reaches a predefined number. For example, if you set a Cardinality condition where the `path` is set to `/foo` and the number of paths is set to `2`, the following variables would meet the condition in a RuntimeConfig resource: + `/foo/variable1 = "value1"` + `/foo/variable2 = "value2"` + `/bar/variable3 = "value3"` It would not satisfy the same condition with the `number` set to `3`, however, because there is only 2 paths that start with `/foo`. Cardinality conditions are recursive; all subtrees under the specific path prefix are counted.
    /// </summary>
    public sealed class CardinalityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number variables under the `path` that must exist to meet this condition. Defaults to 1 if not specified.
        /// </summary>
        [Input("number")]
        public Input<int>? Number { get; set; }

        /// <summary>
        /// The root of the variable subtree to monitor. For example, `/foo`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public CardinalityArgs()
        {
        }
        public static new CardinalityArgs Empty => new CardinalityArgs();
    }
}
