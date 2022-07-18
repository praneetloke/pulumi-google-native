// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2.Inputs
{

    /// <summary>
    /// Filters events based on exact matches on the CloudEvents attributes.
    /// </summary>
    public sealed class EventFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of a CloudEvents attribute.
        /// </summary>
        [Input("attribute", required: true)]
        public Input<string> Attribute { get; set; } = null!;

        /// <summary>
        /// Optional. The operator used for matching the events with the value of the filter. If not specified, only events that have an exact key-value pair specified in the filter are matched. The only allowed value is `match-path-pattern`.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// The value for the attribute.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public EventFilterArgs()
        {
        }
    }
}
