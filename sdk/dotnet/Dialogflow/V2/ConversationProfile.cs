// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2
{
    /// <summary>
    /// Creates a conversation profile in the specified project. ConversationProfile.CreateTime and ConversationProfile.UpdateTime aren't populated in the response. You can retrieve them via GetConversationProfile API.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v2:ConversationProfile")]
    public partial class ConversationProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration for an automated agent to use with this profile.
        /// </summary>
        [Output("automatedAgentConfig")]
        public Output<Outputs.GoogleCloudDialogflowV2AutomatedAgentConfigResponse> AutomatedAgentConfig { get; private set; } = null!;

        /// <summary>
        /// Create time of the conversation profile.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Human readable name for this profile. Max length 1024 bytes.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Configuration for agent assistance to use with this profile.
        /// </summary>
        [Output("humanAgentAssistantConfig")]
        public Output<Outputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigResponse> HumanAgentAssistantConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
        /// </summary>
        [Output("humanAgentHandoffConfig")]
        public Output<Outputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigResponse> HumanAgentHandoffConfig { get; private set; } = null!;

        /// <summary>
        /// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-US languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
        /// </summary>
        [Output("languageCode")]
        public Output<string> LanguageCode { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Configuration for logging conversation lifecycle events.
        /// </summary>
        [Output("loggingConfig")]
        public Output<Outputs.GoogleCloudDialogflowV2LoggingConfigResponse> LoggingConfig { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
        /// </summary>
        [Output("newMessageEventNotificationConfig")]
        public Output<Outputs.GoogleCloudDialogflowV2NotificationConfigResponse> NewMessageEventNotificationConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration for publishing conversation lifecycle events.
        /// </summary>
        [Output("notificationConfig")]
        public Output<Outputs.GoogleCloudDialogflowV2NotificationConfigResponse> NotificationConfig { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
        /// </summary>
        [Output("securitySettings")]
        public Output<string> SecuritySettings { get; private set; } = null!;

        /// <summary>
        /// Settings for speech transcription.
        /// </summary>
        [Output("sttConfig")]
        public Output<Outputs.GoogleCloudDialogflowV2SpeechToTextConfigResponse> SttConfig { get; private set; } = null!;

        /// <summary>
        /// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
        /// </summary>
        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;

        /// <summary>
        /// Update time of the conversation profile.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ConversationProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConversationProfile(string name, ConversationProfileArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2:ConversationProfile", name, args ?? new ConversationProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConversationProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2:ConversationProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConversationProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConversationProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConversationProfile(name, id, options);
        }
    }

    public sealed class ConversationProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for an automated agent to use with this profile.
        /// </summary>
        [Input("automatedAgentConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2AutomatedAgentConfigArgs>? AutomatedAgentConfig { get; set; }

        /// <summary>
        /// Human readable name for this profile. Max length 1024 bytes.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Configuration for agent assistance to use with this profile.
        /// </summary>
        [Input("humanAgentAssistantConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigArgs>? HumanAgentAssistantConfig { get; set; }

        /// <summary>
        /// Configuration for connecting to a live agent. Currently, this feature is not general available, please contact Google to get access.
        /// </summary>
        [Input("humanAgentHandoffConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigArgs>? HumanAgentHandoffConfig { get; set; }

        /// <summary>
        /// Language code for the conversation profile. If not specified, the language is en-US. Language at ConversationProfile should be set for all non en-US languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Configuration for logging conversation lifecycle events.
        /// </summary>
        [Input("loggingConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2LoggingConfigArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// The unique identifier of this conversation profile. Format: `projects//locations//conversationProfiles/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration for publishing new message events. Event will be sent in format of ConversationEvent
        /// </summary>
        [Input("newMessageEventNotificationConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2NotificationConfigArgs>? NewMessageEventNotificationConfig { get; set; }

        /// <summary>
        /// Configuration for publishing conversation lifecycle events.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2NotificationConfigArgs>? NotificationConfig { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Name of the CX SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
        /// </summary>
        [Input("securitySettings")]
        public Input<string>? SecuritySettings { get; set; }

        /// <summary>
        /// Settings for speech transcription.
        /// </summary>
        [Input("sttConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2SpeechToTextConfigArgs>? SttConfig { get; set; }

        /// <summary>
        /// The time zone of this conversational profile from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris. Defaults to America/New_York.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public ConversationProfileArgs()
        {
        }
        public static new ConversationProfileArgs Empty => new ConversationProfileArgs();
    }
}
