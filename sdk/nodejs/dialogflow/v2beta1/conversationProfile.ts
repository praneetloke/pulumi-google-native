// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a conversation profile in the specified project. ConversationProfile.CreateTime and ConversationProfile.UpdateTime aren't populated in the response. You can retrieve them via GetConversationProfile API.
 */
export class ConversationProfile extends pulumi.CustomResource {
    /**
     * Get an existing ConversationProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConversationProfile {
        return new ConversationProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2beta1:ConversationProfile';

    /**
     * Returns true if the given object is an instance of ConversationProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConversationProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConversationProfile.__pulumiType;
    }

    /**
     * Configuration for an automated agent to use with this profile.
     */
    public readonly automatedAgentConfig!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponse>;
    /**
     * Create time of the conversation profile.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Human readable name for this profile. Max length 1024 bytes.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Configuration for agent assistance to use with this profile.
     */
    public readonly humanAgentAssistantConfig!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponse>;
    /**
     * Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
     */
    public readonly humanAgentHandoffConfig!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponse>;
    /**
     * Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages.
     */
    public readonly languageCode!: pulumi.Output<string>;
    /**
     * Configuration for logging conversation lifecycle events.
     */
    public readonly loggingConfig!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1LoggingConfigResponse>;
    /**
     * The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for publishing new message events. Event will be sent in format of ConversationEvent
     */
    public readonly newMessageEventNotificationConfig!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1NotificationConfigResponse>;
    /**
     * Configuration for publishing conversation lifecycle events.
     */
    public readonly notificationConfig!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1NotificationConfigResponse>;
    /**
     * Settings for speech transcription.
     */
    public readonly sttConfig!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1SpeechToTextConfigResponse>;
    /**
     * Update time of the conversation profile.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ConversationProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConversationProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["automatedAgentConfig"] = args ? args.automatedAgentConfig : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["humanAgentAssistantConfig"] = args ? args.humanAgentAssistantConfig : undefined;
            inputs["humanAgentHandoffConfig"] = args ? args.humanAgentHandoffConfig : undefined;
            inputs["languageCode"] = args ? args.languageCode : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["loggingConfig"] = args ? args.loggingConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["newMessageEventNotificationConfig"] = args ? args.newMessageEventNotificationConfig : undefined;
            inputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sttConfig"] = args ? args.sttConfig : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["automatedAgentConfig"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["humanAgentAssistantConfig"] = undefined /*out*/;
            inputs["humanAgentHandoffConfig"] = undefined /*out*/;
            inputs["languageCode"] = undefined /*out*/;
            inputs["loggingConfig"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["newMessageEventNotificationConfig"] = undefined /*out*/;
            inputs["notificationConfig"] = undefined /*out*/;
            inputs["sttConfig"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConversationProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConversationProfile resource.
 */
export interface ConversationProfileArgs {
    /**
     * Configuration for an automated agent to use with this profile.
     */
    automatedAgentConfig?: pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1AutomatedAgentConfigArgs>;
    /**
     * Human readable name for this profile. Max length 1024 bytes.
     */
    displayName: pulumi.Input<string>;
    /**
     * Configuration for agent assistance to use with this profile.
     */
    humanAgentAssistantConfig?: pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigArgs>;
    /**
     * Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
     */
    humanAgentHandoffConfig?: pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigArgs>;
    /**
     * Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages.
     */
    languageCode?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Configuration for logging conversation lifecycle events.
     */
    loggingConfig?: pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1LoggingConfigArgs>;
    /**
     * The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration for publishing new message events. Event will be sent in format of ConversationEvent
     */
    newMessageEventNotificationConfig?: pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1NotificationConfigArgs>;
    /**
     * Configuration for publishing conversation lifecycle events.
     */
    notificationConfig?: pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1NotificationConfigArgs>;
    project?: pulumi.Input<string>;
    /**
     * Settings for speech transcription.
     */
    sttConfig?: pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1SpeechToTextConfigArgs>;
}
