// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Inputs
{

    /// <summary>
    /// Metadata for a Pub/Sub connector used by the job.
    /// </summary>
    public sealed class PubSubIODetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Subscription used in the connection.
        /// </summary>
        [Input("subscription")]
        public Input<string>? Subscription { get; set; }

        /// <summary>
        /// Topic accessed in the connection.
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public PubSubIODetailsArgs()
        {
        }
        public static new PubSubIODetailsArgs Empty => new PubSubIODetailsArgs();
    }
}
