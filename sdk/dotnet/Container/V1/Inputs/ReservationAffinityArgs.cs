// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// [ReservationAffinity](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources) is the configuration of desired reservation which instances could take capacity from.
    /// </summary>
    public sealed class ReservationAffinityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Corresponds to the type of reservation consumption.
        /// </summary>
        [Input("consumeReservationType")]
        public Input<Pulumi.GoogleNative.Container.V1.ReservationAffinityConsumeReservationType>? ConsumeReservationType { get; set; }

        /// <summary>
        /// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify "googleapis.com/reservation-name" as the key and specify the name of your reservation as its value.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Corresponds to the label value(s) of reservation resource(s).
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ReservationAffinityArgs()
        {
        }
        public static new ReservationAffinityArgs Empty => new ReservationAffinityArgs();
    }
}
