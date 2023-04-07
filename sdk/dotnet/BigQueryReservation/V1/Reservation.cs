// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQueryReservation.V1
{
    /// <summary>
    /// Creates a new reservation resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:bigqueryreservation/v1:Reservation")]
    public partial class Reservation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration parameters for the auto scaling feature. Note this is an alpha feature.
        /// </summary>
        [Output("autoscale")]
        public Output<Outputs.AutoscaleResponse> Autoscale { get; private set; } = null!;

        /// <summary>
        /// Job concurrency target which sets a soft upper bound on the number of jobs that can run concurrently in this reservation. This is a soft target due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency target will be automatically computed by the system. NOTE: this field is exposed as `target_job_concurrency` in the Information Schema, DDL and BQ CLI.
        /// </summary>
        [Output("concurrency")]
        public Output<string> Concurrency { get; private set; } = null!;

        /// <summary>
        /// Creation time of the reservation.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Edition of the reservation.
        /// </summary>
        [Output("edition")]
        public Output<string> Edition { get; private set; } = null!;

        /// <summary>
        /// If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
        /// </summary>
        [Output("ignoreIdleSlots")]
        public Output<bool> IgnoreIdleSlots { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
        /// </summary>
        [Output("multiRegionAuxiliary")]
        public Output<bool> MultiRegionAuxiliary { get; private set; } = null!;

        /// <summary>
        /// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The reservation ID. It must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        /// </summary>
        [Output("reservationId")]
        public Output<string?> ReservationId { get; private set; } = null!;

        /// <summary>
        /// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If total slot_capacity of the reservation and its siblings exceeds the total slot_count of all capacity commitments, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
        /// </summary>
        [Output("slotCapacity")]
        public Output<string> SlotCapacity { get; private set; } = null!;

        /// <summary>
        /// Last update time of the reservation.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Reservation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Reservation(string name, ReservationArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:bigqueryreservation/v1:Reservation", name, args ?? new ReservationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Reservation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:bigqueryreservation/v1:Reservation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Reservation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Reservation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Reservation(name, id, options);
        }
    }

    public sealed class ReservationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration parameters for the auto scaling feature. Note this is an alpha feature.
        /// </summary>
        [Input("autoscale")]
        public Input<Inputs.AutoscaleArgs>? Autoscale { get; set; }

        /// <summary>
        /// Job concurrency target which sets a soft upper bound on the number of jobs that can run concurrently in this reservation. This is a soft target due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency target will be automatically computed by the system. NOTE: this field is exposed as `target_job_concurrency` in the Information Schema, DDL and BQ CLI.
        /// </summary>
        [Input("concurrency")]
        public Input<string>? Concurrency { get; set; }

        /// <summary>
        /// Edition of the reservation.
        /// </summary>
        [Input("edition")]
        public Input<Pulumi.GoogleNative.BigQueryReservation.V1.ReservationEdition>? Edition { get; set; }

        /// <summary>
        /// If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
        /// </summary>
        [Input("ignoreIdleSlots")]
        public Input<bool>? IgnoreIdleSlots { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
        /// </summary>
        [Input("multiRegionAuxiliary")]
        public Input<bool>? MultiRegionAuxiliary { get; set; }

        /// <summary>
        /// The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The reservation ID. It must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        /// </summary>
        [Input("reservationId")]
        public Input<string>? ReservationId { get; set; }

        /// <summary>
        /// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If total slot_capacity of the reservation and its siblings exceeds the total slot_count of all capacity commitments, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
        /// </summary>
        [Input("slotCapacity")]
        public Input<string>? SlotCapacity { get; set; }

        public ReservationArgs()
        {
        }
        public static new ReservationArgs Empty => new ReservationArgs();
    }
}
