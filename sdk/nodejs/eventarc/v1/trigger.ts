// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a new trigger in a particular project and location.
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:eventarc/v1:Trigger';

    /**
     * Returns true if the given object is an instance of Trigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trigger.__pulumiType;
    }

    /**
     * The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Destination specifies where the events should be sent to.
     */
    public readonly destination!: pulumi.Output<outputs.eventarc.v1.DestinationResponse>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
     */
    public readonly eventFilters!: pulumi.Output<outputs.eventarc.v1.EventFilterResponse[]>;
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    public readonly transport!: pulumi.Output<outputs.eventarc.v1.TransportResponse>;
    /**
     * Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The last-modified time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.eventFilters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventFilters'");
            }
            if ((!args || args.triggerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggerId'");
            }
            if ((!args || args.validateOnly === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validateOnly'");
            }
            inputs["destination"] = args ? args.destination : undefined;
            inputs["eventFilters"] = args ? args.eventFilters : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["transport"] = args ? args.transport : undefined;
            inputs["triggerId"] = args ? args.triggerId : undefined;
            inputs["validateOnly"] = args ? args.validateOnly : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["uid"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["destination"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["eventFilters"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["serviceAccount"] = undefined /*out*/;
            inputs["transport"] = undefined /*out*/;
            inputs["uid"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Trigger.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * Destination specifies where the events should be sent to.
     */
    destination: pulumi.Input<inputs.eventarc.v1.DestinationArgs>;
    /**
     * null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
     */
    eventFilters: pulumi.Input<pulumi.Input<inputs.eventarc.v1.EventFilterArgs>[]>;
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    transport?: pulumi.Input<inputs.eventarc.v1.TransportArgs>;
    triggerId: pulumi.Input<string>;
    validateOnly: pulumi.Input<string>;
}
