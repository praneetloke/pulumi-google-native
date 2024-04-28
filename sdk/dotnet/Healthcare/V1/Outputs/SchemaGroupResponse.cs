// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Outputs
{

    /// <summary>
    /// An HL7v2 logical group construct.
    /// </summary>
    [OutputType]
    public sealed class SchemaGroupResponse
    {
        /// <summary>
        /// True indicates that this is a choice group, meaning that only one of its segments can exist in a given message.
        /// </summary>
        public readonly bool Choice;
        /// <summary>
        /// The maximum number of times this group can be repeated. 0 or -1 means unbounded.
        /// </summary>
        public readonly int MaxOccurs;
        /// <summary>
        /// Nested groups and/or segments.
        /// </summary>
        public readonly ImmutableArray<Outputs.GroupOrSegmentResponse> Members;
        /// <summary>
        /// The minimum number of times this group must be present/repeated.
        /// </summary>
        public readonly int MinOccurs;
        /// <summary>
        /// The name of this group. For example, "ORDER_DETAIL".
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private SchemaGroupResponse(
            bool choice,

            int maxOccurs,

            ImmutableArray<Outputs.GroupOrSegmentResponse> members,

            int minOccurs,

            string name)
        {
            Choice = choice;
            MaxOccurs = maxOccurs;
            Members = members;
            MinOccurs = minOccurs;
            Name = name;
        }
    }
}
