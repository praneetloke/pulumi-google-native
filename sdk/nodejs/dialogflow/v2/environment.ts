// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an agent environment.
 * Auto-naming is currently not supported for this resource.
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * Optional. The agent version loaded into this environment. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
     */
    public readonly agentVersion!: pulumi.Output<string>;
    /**
     * Optional. The developer-provided description for this environment. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. The fulfillment settings to use for this environment.
     */
    public readonly fulfillment!: pulumi.Output<outputs.dialogflow.v2.GoogleCloudDialogflowV2FulfillmentResponse>;
    /**
     * The unique identifier of this agent environment. Supported formats: - `projects//agent/environments/` - `projects//locations//agent/environments/` The environment ID for the default environment is `-`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The state of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Optional. Text to speech settings for this environment.
     */
    public readonly textToSpeechSettings!: pulumi.Output<outputs.dialogflow.v2.GoogleCloudDialogflowV2TextToSpeechSettingsResponse>;
    /**
     * The last update time of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            inputs["agentVersion"] = args ? args.agentVersion : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["fulfillment"] = args ? args.fulfillment : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["textToSpeechSettings"] = args ? args.textToSpeechSettings : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["agentVersion"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["fulfillment"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["textToSpeechSettings"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Environment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * Optional. The agent version loaded into this environment. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
     */
    agentVersion?: pulumi.Input<string>;
    /**
     * Optional. The developer-provided description for this environment. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    description?: pulumi.Input<string>;
    environmentId: pulumi.Input<string>;
    /**
     * Optional. The fulfillment settings to use for this environment.
     */
    fulfillment?: pulumi.Input<inputs.dialogflow.v2.GoogleCloudDialogflowV2FulfillmentArgs>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Text to speech settings for this environment.
     */
    textToSpeechSettings?: pulumi.Input<inputs.dialogflow.v2.GoogleCloudDialogflowV2TextToSpeechSettingsArgs>;
}
