// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1.Outputs
{

    /// <summary>
    /// Data type and value of a parameter.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudChannelV1ValueResponse
    {
        /// <summary>
        /// Represents a boolean value.
        /// </summary>
        public readonly bool BoolValue;
        /// <summary>
        /// Represents a double value.
        /// </summary>
        public readonly double DoubleValue;
        /// <summary>
        /// Represents an int64 value.
        /// </summary>
        public readonly string Int64Value;
        /// <summary>
        /// Represents an 'Any' proto value.
        /// </summary>
        public readonly ImmutableDictionary<string, object> ProtoValue;
        /// <summary>
        /// Represents a string value.
        /// </summary>
        public readonly string StringValue;

        [OutputConstructor]
        private GoogleCloudChannelV1ValueResponse(
            bool boolValue,

            double doubleValue,

            string int64Value,

            ImmutableDictionary<string, object> protoValue,

            string stringValue)
        {
            BoolValue = boolValue;
            DoubleValue = doubleValue;
            Int64Value = int64Value;
            ProtoValue = protoValue;
            StringValue = stringValue;
        }
    }
}
