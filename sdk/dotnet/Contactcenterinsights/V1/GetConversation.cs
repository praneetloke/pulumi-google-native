// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1
{
    public static class GetConversation
    {
        /// <summary>
        /// Gets a conversation.
        /// </summary>
        public static Task<GetConversationResult> InvokeAsync(GetConversationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConversationResult>("google-native:contactcenterinsights/v1:getConversation", args ?? new GetConversationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a conversation.
        /// </summary>
        public static Output<GetConversationResult> Invoke(GetConversationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConversationResult>("google-native:contactcenterinsights/v1:getConversation", args ?? new GetConversationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConversationArgs : global::Pulumi.InvokeArgs
    {
        [Input("conversationId", required: true)]
        public string ConversationId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        public GetConversationArgs()
        {
        }
        public static new GetConversationArgs Empty => new GetConversationArgs();
    }

    public sealed class GetConversationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("conversationId", required: true)]
        public Input<string> ConversationId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetConversationInvokeArgs()
        {
        }
        public static new GetConversationInvokeArgs Empty => new GetConversationInvokeArgs();
    }


    [OutputType]
    public sealed class GetConversationResult
    {
        /// <summary>
        /// An opaque, user-specified string representing the human agent who handled the conversation.
        /// </summary>
        public readonly string AgentId;
        /// <summary>
        /// Call-specific metadata.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1ConversationCallMetadataResponse CallMetadata;
        /// <summary>
        /// The time at which the conversation was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The source of the audio and transcription for the conversation.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1ConversationDataSourceResponse DataSource;
        /// <summary>
        /// All the matched Dialogflow intents in the call. The key corresponds to a Dialogflow intent, format: projects/{project}/agent/{agent}/intents/{intent}
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1DialogflowIntentResponse DialogflowIntents;
        /// <summary>
        /// The duration of the conversation.
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// The time at which this conversation should expire. After this time, the conversation data and any associated analyses will be deleted.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// A map for the user to specify any custom fields. A maximum of 20 labels per conversation is allowed, with a maximum of 256 characters per entry.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// A user-specified language code for the conversation.
        /// </summary>
        public readonly string LanguageCode;
        /// <summary>
        /// The conversation's latest analysis, if one exists.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1AnalysisResponse LatestAnalysis;
        /// <summary>
        /// Latest summary of the conversation.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1ConversationSummarizationSuggestionDataResponse LatestSummary;
        /// <summary>
        /// Immutable. The conversation medium, if unspecified will default to PHONE_CALL.
        /// </summary>
        public readonly string Medium;
        /// <summary>
        /// Immutable. The resource name of the conversation. Format: projects/{project}/locations/{location}/conversations/{conversation}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Obfuscated user ID which the customer sent to us.
        /// </summary>
        public readonly string ObfuscatedUserId;
        /// <summary>
        /// The annotations that were generated during the customer and agent interaction.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudContactcenterinsightsV1RuntimeAnnotationResponse> RuntimeAnnotations;
        /// <summary>
        /// The time at which the conversation started.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The conversation transcript.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1ConversationTranscriptResponse Transcript;
        /// <summary>
        /// Input only. The TTL for this resource. If specified, then this TTL will be used to calculate the expire time.
        /// </summary>
        public readonly string Ttl;
        /// <summary>
        /// The number of turns in the conversation.
        /// </summary>
        public readonly int TurnCount;
        /// <summary>
        /// The most recent time at which the conversation was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetConversationResult(
            string agentId,

            Outputs.GoogleCloudContactcenterinsightsV1ConversationCallMetadataResponse callMetadata,

            string createTime,

            Outputs.GoogleCloudContactcenterinsightsV1ConversationDataSourceResponse dataSource,

            Outputs.GoogleCloudContactcenterinsightsV1DialogflowIntentResponse dialogflowIntents,

            string duration,

            string expireTime,

            ImmutableDictionary<string, string> labels,

            string languageCode,

            Outputs.GoogleCloudContactcenterinsightsV1AnalysisResponse latestAnalysis,

            Outputs.GoogleCloudContactcenterinsightsV1ConversationSummarizationSuggestionDataResponse latestSummary,

            string medium,

            string name,

            string obfuscatedUserId,

            ImmutableArray<Outputs.GoogleCloudContactcenterinsightsV1RuntimeAnnotationResponse> runtimeAnnotations,

            string startTime,

            Outputs.GoogleCloudContactcenterinsightsV1ConversationTranscriptResponse transcript,

            string ttl,

            int turnCount,

            string updateTime)
        {
            AgentId = agentId;
            CallMetadata = callMetadata;
            CreateTime = createTime;
            DataSource = dataSource;
            DialogflowIntents = dialogflowIntents;
            Duration = duration;
            ExpireTime = expireTime;
            Labels = labels;
            LanguageCode = languageCode;
            LatestAnalysis = latestAnalysis;
            LatestSummary = latestSummary;
            Medium = medium;
            Name = name;
            ObfuscatedUserId = obfuscatedUserId;
            RuntimeAnnotations = runtimeAnnotations;
            StartTime = startTime;
            Transcript = transcript;
            Ttl = ttl;
            TurnCount = turnCount;
            UpdateTime = updateTime;
        }
    }
}
