// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1.Inputs
{

    /// <summary>
    /// `Value` represents a dynamically typed value which can be either be a float, a integer, a string, or a datetime value. A producer of value is expected to set one of these variants. Absence of any variant indicates an error.
    /// </summary>
    public sealed class GoogleCloudContentwarehouseV1ValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents a boolean value.
        /// </summary>
        [Input("booleanValue")]
        public Input<bool>? BooleanValue { get; set; }

        /// <summary>
        /// Represents a datetime value.
        /// </summary>
        [Input("datetimeValue")]
        public Input<Inputs.GoogleTypeDateTimeArgs>? DatetimeValue { get; set; }

        /// <summary>
        /// Represents an enum value.
        /// </summary>
        [Input("enumValue")]
        public Input<Inputs.GoogleCloudContentwarehouseV1EnumValueArgs>? EnumValue { get; set; }

        /// <summary>
        /// Represents a float value.
        /// </summary>
        [Input("floatValue")]
        public Input<double>? FloatValue { get; set; }

        /// <summary>
        /// Represents a integer value.
        /// </summary>
        [Input("intValue")]
        public Input<int>? IntValue { get; set; }

        /// <summary>
        /// Represents a string value.
        /// </summary>
        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        /// <summary>
        /// Represents a timestamp value.
        /// </summary>
        [Input("timestampValue")]
        public Input<Inputs.GoogleCloudContentwarehouseV1TimestampValueArgs>? TimestampValue { get; set; }

        public GoogleCloudContentwarehouseV1ValueArgs()
        {
        }
        public static new GoogleCloudContentwarehouseV1ValueArgs Empty => new GoogleCloudContentwarehouseV1ValueArgs();
    }
}
