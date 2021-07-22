// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new reservation resource.
 */
export class Reservation extends pulumi.CustomResource {
    /**
     * Get an existing Reservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Reservation {
        return new Reservation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:bigqueryreservation/v1beta1:Reservation';

    /**
     * Returns true if the given object is an instance of Reservation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Reservation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Reservation.__pulumiType;
    }

    /**
     * Creation time of the reservation.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * If false, any query using this reservation will use idle slots from other reservations within the same admin project. If true, a query using this reservation will execute with the slot capacity specified above at most.
     */
    public readonly ignoreIdleSlots!: pulumi.Output<boolean>;
    /**
     * The resource name of the reservation, e.g., `projects/*&#47;locations/*&#47;reservations/team1-prod`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceed the parent's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the parent's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`.
     */
    public readonly slotCapacity!: pulumi.Output<string>;
    /**
     * Last update time of the reservation.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Reservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ReservationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["ignoreIdleSlots"] = args ? args.ignoreIdleSlots : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["reservationId"] = args ? args.reservationId : undefined;
            inputs["slotCapacity"] = args ? args.slotCapacity : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["ignoreIdleSlots"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["slotCapacity"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Reservation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Reservation resource.
 */
export interface ReservationArgs {
    /**
     * If false, any query using this reservation will use idle slots from other reservations within the same admin project. If true, a query using this reservation will execute with the slot capacity specified above at most.
     */
    ignoreIdleSlots?: pulumi.Input<boolean>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the reservation, e.g., `projects/*&#47;locations/*&#47;reservations/team1-prod`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    reservationId?: pulumi.Input<string>;
    /**
     * Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If the new reservation's slot capacity exceed the parent's slot capacity or if total slot capacity of the new reservation and its siblings exceeds the parent's slot capacity, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`.
     */
    slotCapacity?: pulumi.Input<string>;
}
