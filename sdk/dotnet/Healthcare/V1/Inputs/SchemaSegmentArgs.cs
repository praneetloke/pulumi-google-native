// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Inputs
{

    /// <summary>
    /// An HL7v2 Segment.
    /// </summary>
    public sealed class SchemaSegmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of times this segment can be present in this group. 0 or -1 means unbounded.
        /// </summary>
        [Input("maxOccurs")]
        public Input<int>? MaxOccurs { get; set; }

        /// <summary>
        /// The minimum number of times this segment can be present in this group.
        /// </summary>
        [Input("minOccurs")]
        public Input<int>? MinOccurs { get; set; }

        /// <summary>
        /// The Segment type. For example, "PID".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SchemaSegmentArgs()
        {
        }
        public static new SchemaSegmentArgs Empty => new SchemaSegmentArgs();
    }
}
