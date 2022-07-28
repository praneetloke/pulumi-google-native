// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsublite.V1.Inputs
{

    /// <summary>
    /// The settings for this topic's Reservation usage.
    /// </summary>
    public sealed class ReservationConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Reservation to use for this topic's throughput capacity. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
        /// </summary>
        [Input("throughputReservation")]
        public Input<string>? ThroughputReservation { get; set; }

        public ReservationConfigArgs()
        {
        }
        public static new ReservationConfigArgs Empty => new ReservationConfigArgs();
    }
}
