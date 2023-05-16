// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns information about the reservation.
 */
export function getReservation(args: GetReservationArgs, opts?: pulumi.InvokeOptions): Promise<GetReservationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:bigqueryreservation/v1:getReservation", {
        "location": args.location,
        "project": args.project,
        "reservationId": args.reservationId,
    }, opts);
}

export interface GetReservationArgs {
    location: string;
    project?: string;
    reservationId: string;
}

export interface GetReservationResult {
    /**
     * The configuration parameters for the auto scaling feature.
     */
    readonly autoscale: outputs.bigqueryreservation.v1.AutoscaleResponse;
    /**
     * Job concurrency target which sets a soft upper bound on the number of jobs that can run concurrently in this reservation. This is a soft target due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency target will be automatically computed by the system. NOTE: this field is exposed as `target_job_concurrency` in the Information Schema, DDL and BQ CLI.
     */
    readonly concurrency: string;
    /**
     * Creation time of the reservation.
     */
    readonly creationTime: string;
    /**
     * Edition of the reservation.
     */
    readonly edition: string;
    /**
     * If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
     */
    readonly ignoreIdleSlots: boolean;
    /**
     * Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
     */
    readonly multiRegionAuxiliary: boolean;
    /**
     * The resource name of the reservation, e.g., `projects/*&#47;locations/*&#47;reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
     */
    readonly name: string;
    /**
     * Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If edition is EDITION_UNSPECIFIED and total slot_capacity of the reservation and its siblings exceeds the total slot_count of all capacity commitments, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. If edition is any value but EDITION_UNSPECIFIED, then the above requirement is not needed. The total slot_capacity of the reservation and its siblings may exceed the total slot_count of capacity commitments. In that case, the exceeding slots will be charged with the autoscale SKU. You can increase the number of baseline slots in a reservation every few minutes. If you want to decrease your baseline slots, you are limited to once an hour if you have recently changed your baseline slot capacity and your baseline slots exceed your committed slots. Otherwise, you can decrease your baseline slots every few minutes. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
     */
    readonly slotCapacity: string;
    /**
     * Last update time of the reservation.
     */
    readonly updateTime: string;
}
/**
 * Returns information about the reservation.
 */
export function getReservationOutput(args: GetReservationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReservationResult> {
    return pulumi.output(args).apply((a: any) => getReservation(a, opts))
}

export interface GetReservationOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    reservationId: pulumi.Input<string>;
}
