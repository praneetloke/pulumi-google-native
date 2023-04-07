// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified conversation profile.
 */
export function getConversationProfile(args: GetConversationProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetConversationProfileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dialogflow/v2beta1:getConversationProfile", {
        "conversationProfileId": args.conversationProfileId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetConversationProfileArgs {
    conversationProfileId: string;
    location: string;
    project?: string;
}

export interface GetConversationProfileResult {
    /**
     * Configuration for an automated agent to use with this profile.
     */
    readonly automatedAgentConfig: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1AutomatedAgentConfigResponse;
    /**
     * Create time of the conversation profile.
     */
    readonly createTime: string;
    /**
     * Human readable name for this profile. Max length 1024 bytes.
     */
    readonly displayName: string;
    /**
     * Configuration for agent assistance to use with this profile.
     */
    readonly humanAgentAssistantConfig: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponse;
    /**
     * Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
     */
    readonly humanAgentHandoffConfig: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1HumanAgentHandoffConfigResponse;
    /**
     * Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
     */
    readonly languageCode: string;
    /**
     * Configuration for logging conversation lifecycle events.
     */
    readonly loggingConfig: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1LoggingConfigResponse;
    /**
     * The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
     */
    readonly name: string;
    /**
     * Configuration for publishing new message events. Event will be sent in format of ConversationEvent
     */
    readonly newMessageEventNotificationConfig: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1NotificationConfigResponse;
    /**
     * Configuration for publishing conversation lifecycle events.
     */
    readonly notificationConfig: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1NotificationConfigResponse;
    /**
     * Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
     */
    readonly securitySettings: string;
    /**
     * Settings for speech transcription.
     */
    readonly sttConfig: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1SpeechToTextConfigResponse;
    /**
     * The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
     */
    readonly timeZone: string;
    /**
     * Configuration for Text-to-Speech synthesization. Used by Phone Gateway to specify synthesization options. If agent defines synthesization options as well, agent settings overrides the option here.
     */
    readonly ttsConfig: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponse;
    /**
     * Update time of the conversation profile.
     */
    readonly updateTime: string;
}
/**
 * Retrieves the specified conversation profile.
 */
export function getConversationProfileOutput(args: GetConversationProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConversationProfileResult> {
    return pulumi.output(args).apply((a: any) => getConversationProfile(a, opts))
}

export interface GetConversationProfileOutputArgs {
    conversationProfileId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
