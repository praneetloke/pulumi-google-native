// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get a single trigger.
 */
export function getTrigger(args: GetTriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetTriggerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:eventarc/v1:getTrigger", {
        "location": args.location,
        "project": args.project,
        "triggerId": args.triggerId,
    }, opts);
}

export interface GetTriggerArgs {
    location: string;
    project?: string;
    triggerId: string;
}

export interface GetTriggerResult {
    /**
     * Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
     */
    readonly channel: string;
    /**
     * The reason(s) why a trigger is in FAILED state.
     */
    readonly conditions: {[key: string]: string};
    /**
     * The creation time.
     */
    readonly createTime: string;
    /**
     * Destination specifies where the events should be sent to.
     */
    readonly destination: outputs.eventarc.v1.DestinationResponse;
    /**
     * This checksum is computed by the server based on the value of other fields, and might be sent only on create requests to ensure that the client has an up-to-date value before proceeding.
     */
    readonly etag: string;
    /**
     * Unordered list. The list of filters that applies to event attributes. Only events that match all the provided filters are sent to the destination.
     */
    readonly eventFilters: outputs.eventarc.v1.EventFilterResponse[];
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource name of the trigger. Must be unique within the location of the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
     */
    readonly name: string;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have the `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. To create Audit Log triggers, the service account should also have the `roles/eventarc.eventReceiver` IAM role.
     */
    readonly serviceAccount: string;
    /**
     * Optional. To deliver messages, Eventarc might use other GCP products as a transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    readonly transport: outputs.eventarc.v1.TransportResponse;
    /**
     * Server-assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     */
    readonly uid: string;
    /**
     * The last-modified time.
     */
    readonly updateTime: string;
}

export function getTriggerOutput(args: GetTriggerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTriggerResult> {
    return pulumi.output(args).apply(a => getTrigger(a, opts))
}

export interface GetTriggerOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    triggerId: pulumi.Input<string>;
}
