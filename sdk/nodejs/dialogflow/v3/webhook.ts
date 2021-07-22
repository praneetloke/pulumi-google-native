// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a webhook in the specified agent.
 */
export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Webhook {
        return new Webhook(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    /**
     * Indicates whether the webhook is disabled.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * The human-readable name of the webhook, unique within the agent.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Configuration for a generic web service.
     */
    public readonly genericWebService!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3WebhookGenericWebServiceResponse>;
    /**
     * The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for a [Service Directory](https://cloud.google.com/service-directory) service.
     */
    public readonly serviceDirectory!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3WebhookServiceDirectoryConfigResponse>;
    /**
     * Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
     */
    public readonly timeout!: pulumi.Output<string>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["agentId"] = args ? args.agentId : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["genericWebService"] = args ? args.genericWebService : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["serviceDirectory"] = args ? args.serviceDirectory : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
        } else {
            inputs["disabled"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["genericWebService"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["serviceDirectory"] = undefined /*out*/;
            inputs["timeout"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Webhook.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    agentId: pulumi.Input<string>;
    /**
     * Indicates whether the webhook is disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The human-readable name of the webhook, unique within the agent.
     */
    displayName: pulumi.Input<string>;
    /**
     * Configuration for a generic web service.
     */
    genericWebService?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs>;
    location?: pulumi.Input<string>;
    /**
     * The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Configuration for a [Service Directory](https://cloud.google.com/service-directory) service.
     */
    serviceDirectory?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3WebhookServiceDirectoryConfigArgs>;
    /**
     * Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
     */
    timeout?: pulumi.Input<string>;
}
