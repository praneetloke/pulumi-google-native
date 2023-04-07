// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Represents a reservation resource. A reservation ensures that capacity is held in a specific zone even if the reserved VMs are not running. For more information, read Reserving zonal resources.
    /// </summary>
    public sealed class ReservationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resourcePolicies")]
        private InputMap<string>? _resourcePolicies;

        /// <summary>
        /// Resource policies to be added to this reservation. The key is defined by user, and the value is resource policy url. This is to define placement policy with reservation.
        /// </summary>
        public InputMap<string> ResourcePolicies
        {
            get => _resourcePolicies ?? (_resourcePolicies = new InputMap<string>());
            set => _resourcePolicies = value;
        }

        /// <summary>
        /// Specify share-settings to create a shared reservation. This property is optional. For more information about the syntax and options for this field and its subfields, see the guide for creating a shared reservation.
        /// </summary>
        [Input("shareSettings")]
        public Input<Inputs.ShareSettingsArgs>? ShareSettings { get; set; }

        /// <summary>
        /// Reservation for instances with specific machine shapes.
        /// </summary>
        [Input("specificReservation")]
        public Input<Inputs.AllocationSpecificSKUReservationArgs>? SpecificReservation { get; set; }

        /// <summary>
        /// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
        /// </summary>
        [Input("specificReservationRequired")]
        public Input<bool>? SpecificReservationRequired { get; set; }

        /// <summary>
        /// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ReservationArgs()
        {
        }
        public static new ReservationArgs Empty => new ReservationArgs();
    }
}
