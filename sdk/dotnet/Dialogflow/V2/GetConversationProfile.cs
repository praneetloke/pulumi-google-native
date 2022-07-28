// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2
{
    public static class GetConversationProfile
    {
        /// <summary>
        /// Retrieves the specified conversation profile.
        /// </summary>
        public static Task<GetConversationProfileResult> InvokeAsync(GetConversationProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConversationProfileResult>("google-native:dialogflow/v2:getConversationProfile", args ?? new GetConversationProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified conversation profile.
        /// </summary>
        public static Output<GetConversationProfileResult> Invoke(GetConversationProfileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConversationProfileResult>("google-native:dialogflow/v2:getConversationProfile", args ?? new GetConversationProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConversationProfileArgs : global::Pulumi.InvokeArgs
    {
        [Input("conversationProfileId", required: true)]
        public string ConversationProfileId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetConversationProfileArgs()
        {
        }
        public static new GetConversationProfileArgs Empty => new GetConversationProfileArgs();
    }

    public sealed class GetConversationProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("conversationProfileId", required: true)]
        public Input<string> ConversationProfileId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetConversationProfileInvokeArgs()
        {
        }
        public static new GetConversationProfileInvokeArgs Empty => new GetConversationProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetConversationProfileResult
    {
        /// <summary>
        /// Configuration for an automated agent to use with this profile.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2AutomatedAgentConfigResponse AutomatedAgentConfig;
        /// <summary>
        /// Create time of the conversation profile.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Human readable name for this profile. Max length 1024 bytes.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Configuration for agent assistance to use with this profile.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigResponse HumanAgentAssistantConfig;
        /// <summary>
        /// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigResponse HumanAgentHandoffConfig;
        /// <summary>
        /// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-US languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
        /// </summary>
        public readonly string LanguageCode;
        /// <summary>
        /// Configuration for logging conversation lifecycle events.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2LoggingConfigResponse LoggingConfig;
        /// <summary>
        /// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2NotificationConfigResponse NewMessageEventNotificationConfig;
        /// <summary>
        /// Configuration for publishing conversation lifecycle events.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2NotificationConfigResponse NotificationConfig;
        /// <summary>
        /// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
        /// </summary>
        public readonly string SecuritySettings;
        /// <summary>
        /// Settings for speech transcription.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2SpeechToTextConfigResponse SttConfig;
        /// <summary>
        /// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
        /// </summary>
        public readonly string TimeZone;
        /// <summary>
        /// Update time of the conversation profile.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetConversationProfileResult(
            Outputs.GoogleCloudDialogflowV2AutomatedAgentConfigResponse automatedAgentConfig,

            string createTime,

            string displayName,

            Outputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigResponse humanAgentAssistantConfig,

            Outputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigResponse humanAgentHandoffConfig,

            string languageCode,

            Outputs.GoogleCloudDialogflowV2LoggingConfigResponse loggingConfig,

            string name,

            Outputs.GoogleCloudDialogflowV2NotificationConfigResponse newMessageEventNotificationConfig,

            Outputs.GoogleCloudDialogflowV2NotificationConfigResponse notificationConfig,

            string securitySettings,

            Outputs.GoogleCloudDialogflowV2SpeechToTextConfigResponse sttConfig,

            string timeZone,

            string updateTime)
        {
            AutomatedAgentConfig = automatedAgentConfig;
            CreateTime = createTime;
            DisplayName = displayName;
            HumanAgentAssistantConfig = humanAgentAssistantConfig;
            HumanAgentHandoffConfig = humanAgentHandoffConfig;
            LanguageCode = languageCode;
            LoggingConfig = loggingConfig;
            Name = name;
            NewMessageEventNotificationConfig = newMessageEventNotificationConfig;
            NotificationConfig = notificationConfig;
            SecuritySettings = securitySettings;
            SttConfig = sttConfig;
            TimeZone = timeZone;
            UpdateTime = updateTime;
        }
    }
}
