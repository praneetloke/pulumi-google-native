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
    /// Creates a new conversation. Conversations are auto-completed after 24 hours. Conversation Lifecycle: There are two stages during a conversation: Automated Agent Stage and Assist Stage. For Automated Agent Stage, there will be a dialogflow agent responding to user queries. For Assist Stage, there's no dialogflow agent responding to user queries. But we will provide suggestions which are generated from conversation. If Conversation.conversation_profile is configured for a dialogflow agent, conversation will start from `Automated Agent Stage`, otherwise, it will start from `Assist Stage`. And during `Automated Agent Stage`, once an Intent with Intent.live_agent_handoff is triggered, conversation will transfer to Assist Stage.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v2:Conversation")]
    public partial class Conversation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Identifier of the conversation. Generally it's auto generated by Google. Only set it if you cannot wait for the response to return a auto-generated one to you. The conversation ID must be compliant with the regression fomula "a-zA-Z*" with the characters length in range of [3,64]. If the field is provided, the caller is resposible for 1. the uniqueness of the ID, otherwise the request will be rejected. 2. the consistency for whether to use custom ID or not under a project to better ensure uniqueness.
        /// </summary>
        [Output("conversationId")]
        public Output<string?> ConversationId { get; private set; } = null!;

        /// <summary>
        /// The Conversation Profile to be used to configure this Conversation. This field cannot be updated. Format: `projects//locations//conversationProfiles/`.
        /// </summary>
        [Output("conversationProfile")]
        public Output<string> ConversationProfile { get; private set; } = null!;

        /// <summary>
        /// The stage of a conversation. It indicates whether the virtual agent or a human agent is handling the conversation. If the conversation is created with the conversation profile that has Dialogflow config set, defaults to ConversationStage.VIRTUAL_AGENT_STAGE; Otherwise, defaults to ConversationStage.HUMAN_ASSIST_STAGE. If the conversation is created with the conversation profile that has Dialogflow config set but explicitly sets conversation_stage to ConversationStage.HUMAN_ASSIST_STAGE, it skips ConversationStage.VIRTUAL_AGENT_STAGE stage and directly goes to ConversationStage.HUMAN_ASSIST_STAGE.
        /// </summary>
        [Output("conversationStage")]
        public Output<string> ConversationStage { get; private set; } = null!;

        /// <summary>
        /// The time the conversation was finished.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// The current state of the Conversation.
        /// </summary>
        [Output("lifecycleState")]
        public Output<string> LifecycleState { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of this conversation. Format: `projects//locations//conversations/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// It will not be empty if the conversation is to be connected over telephony.
        /// </summary>
        [Output("phoneNumber")]
        public Output<Outputs.GoogleCloudDialogflowV2ConversationPhoneNumberResponse> PhoneNumber { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The time the conversation was started.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a Conversation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Conversation(string name, ConversationArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2:Conversation", name, args ?? new ConversationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Conversation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2:Conversation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Conversation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Conversation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Conversation(name, id, options);
        }
    }

    public sealed class ConversationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Identifier of the conversation. Generally it's auto generated by Google. Only set it if you cannot wait for the response to return a auto-generated one to you. The conversation ID must be compliant with the regression fomula "a-zA-Z*" with the characters length in range of [3,64]. If the field is provided, the caller is resposible for 1. the uniqueness of the ID, otherwise the request will be rejected. 2. the consistency for whether to use custom ID or not under a project to better ensure uniqueness.
        /// </summary>
        [Input("conversationId")]
        public Input<string>? ConversationId { get; set; }

        /// <summary>
        /// The Conversation Profile to be used to configure this Conversation. This field cannot be updated. Format: `projects//locations//conversationProfiles/`.
        /// </summary>
        [Input("conversationProfile", required: true)]
        public Input<string> ConversationProfile { get; set; } = null!;

        /// <summary>
        /// The stage of a conversation. It indicates whether the virtual agent or a human agent is handling the conversation. If the conversation is created with the conversation profile that has Dialogflow config set, defaults to ConversationStage.VIRTUAL_AGENT_STAGE; Otherwise, defaults to ConversationStage.HUMAN_ASSIST_STAGE. If the conversation is created with the conversation profile that has Dialogflow config set but explicitly sets conversation_stage to ConversationStage.HUMAN_ASSIST_STAGE, it skips ConversationStage.VIRTUAL_AGENT_STAGE stage and directly goes to ConversationStage.HUMAN_ASSIST_STAGE.
        /// </summary>
        [Input("conversationStage")]
        public Input<Pulumi.GoogleNative.Dialogflow.V2.ConversationConversationStage>? ConversationStage { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public ConversationArgs()
        {
        }
        public static new ConversationArgs Empty => new ConversationArgs();
    }
}
