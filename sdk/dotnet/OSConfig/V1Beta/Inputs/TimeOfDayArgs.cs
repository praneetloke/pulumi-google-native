// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Inputs
{

    /// <summary>
    /// Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
    /// </summary>
    public sealed class TimeOfDayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
        /// </summary>
        [Input("hours")]
        public Input<int>? Hours { get; set; }

        /// <summary>
        /// Minutes of hour of day. Must be from 0 to 59.
        /// </summary>
        [Input("minutes")]
        public Input<int>? Minutes { get; set; }

        /// <summary>
        /// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
        /// </summary>
        [Input("nanos")]
        public Input<int>? Nanos { get; set; }

        /// <summary>
        /// Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
        /// </summary>
        [Input("seconds")]
        public Input<int>? Seconds { get; set; }

        public TimeOfDayArgs()
        {
        }
        public static new TimeOfDayArgs Empty => new TimeOfDayArgs();
    }
}
