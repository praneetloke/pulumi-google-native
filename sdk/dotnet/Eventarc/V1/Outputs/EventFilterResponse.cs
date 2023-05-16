// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1.Outputs
{

    /// <summary>
    /// Filters events based on exact matches on the CloudEvents attributes.
    /// </summary>
    [OutputType]
    public sealed class EventFilterResponse
    {
        /// <summary>
        /// The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. You can [retrieve a specific provider's supported event types](/eventarc/docs/list-providers#describe-provider). All triggers MUST provide a filter for the 'type' attribute.
        /// </summary>
        public readonly string Attribute;
        /// <summary>
        /// Optional. The operator used for matching the events with the value of the filter. If not specified, only events that have an exact key-value pair specified in the filter are matched. The only allowed value is `match-path-pattern`.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// The value for the attribute.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private EventFilterResponse(
            string attribute,

            string @operator,

            string value)
        {
            Attribute = attribute;
            Operator = @operator;
            Value = value;
        }
    }
}
