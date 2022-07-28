// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1
{
    public static class GetConversation
    {
        /// <summary>
        /// Retrieves the specific conversation.
        /// </summary>
        public static Task<GetConversationResult> InvokeAsync(GetConversationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConversationResult>("google-native:dialogflow/v2beta1:getConversation", args ?? new GetConversationArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specific conversation.
        /// </summary>
        public static Output<GetConversationResult> Invoke(GetConversationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConversationResult>("google-native:dialogflow/v2beta1:getConversation", args ?? new GetConversationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConversationArgs : global::Pulumi.InvokeArgs
    {
        [Input("conversationId", required: true)]
        public string ConversationId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

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

        public GetConversationInvokeArgs()
        {
        }
        public static new GetConversationInvokeArgs Empty => new GetConversationInvokeArgs();
    }


    [OutputType]
    public sealed class GetConversationResult
    {
        /// <summary>
        /// The Conversation Profile to be used to configure this Conversation. This field cannot be updated. Format: `projects//locations//conversationProfiles/`.
        /// </summary>
        public readonly string ConversationProfile;
        /// <summary>
        /// The stage of a conversation. It indicates whether the virtual agent or a human agent is handling the conversation. If the conversation is created with the conversation profile that has Dialogflow config set, defaults to ConversationStage.VIRTUAL_AGENT_STAGE; Otherwise, defaults to ConversationStage.HUMAN_ASSIST_STAGE. If the conversation is created with the conversation profile that has Dialogflow config set but explicitly sets conversation_stage to ConversationStage.HUMAN_ASSIST_STAGE, it skips ConversationStage.VIRTUAL_AGENT_STAGE stage and directly goes to ConversationStage.HUMAN_ASSIST_STAGE.
        /// </summary>
        public readonly string ConversationStage;
        /// <summary>
        /// The time the conversation was finished.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The current state of the Conversation.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// The unique identifier of this conversation. Format: `projects//locations//conversations/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Required if the conversation is to be connected over telephony.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2beta1ConversationPhoneNumberResponse PhoneNumber;
        /// <summary>
        /// The time the conversation was started.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private GetConversationResult(
            string conversationProfile,

            string conversationStage,

            string endTime,

            string lifecycleState,

            string name,

            Outputs.GoogleCloudDialogflowV2beta1ConversationPhoneNumberResponse phoneNumber,

            string startTime)
        {
            ConversationProfile = conversationProfile;
            ConversationStage = conversationStage;
            EndTime = endTime;
            LifecycleState = lifecycleState;
            Name = name;
            PhoneNumber = phoneNumber;
            StartTime = startTime;
        }
    }
}
