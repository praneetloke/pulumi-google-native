// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Inputs
{

    /// <summary>
    /// Config to process conversation.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2HumanAgentAssistantConfigConversationProcessConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of recent non-small-talk sentences to use as context for article and FAQ suggestion
        /// </summary>
        [Input("recentSentencesCount")]
        public Input<int>? RecentSentencesCount { get; set; }

        public GoogleCloudDialogflowV2HumanAgentAssistantConfigConversationProcessConfigArgs()
        {
        }
        public static new GoogleCloudDialogflowV2HumanAgentAssistantConfigConversationProcessConfigArgs Empty => new GoogleCloudDialogflowV2HumanAgentAssistantConfigConversationProcessConfigArgs();
    }
}
