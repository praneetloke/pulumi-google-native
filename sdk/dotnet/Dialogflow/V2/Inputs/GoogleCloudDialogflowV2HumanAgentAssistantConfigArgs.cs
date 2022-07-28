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
    /// Defines the Human Agent Assist to connect to a conversation.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2HumanAgentAssistantConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for agent assistance of end user participant. Currently, this feature is not general available, please contact Google to get access.
        /// </summary>
        [Input("endUserSuggestionConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionConfigArgs>? EndUserSuggestionConfig { get; set; }

        /// <summary>
        /// Configuration for agent assistance of human agent participant.
        /// </summary>
        [Input("humanAgentSuggestionConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionConfigArgs>? HumanAgentSuggestionConfig { get; set; }

        /// <summary>
        /// Configuration for message analysis.
        /// </summary>
        [Input("messageAnalysisConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigMessageAnalysisConfigArgs>? MessageAnalysisConfig { get; set; }

        /// <summary>
        /// Pub/Sub topic on which to publish new agent assistant events.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2NotificationConfigArgs>? NotificationConfig { get; set; }

        public GoogleCloudDialogflowV2HumanAgentAssistantConfigArgs()
        {
        }
        public static new GoogleCloudDialogflowV2HumanAgentAssistantConfigArgs Empty => new GoogleCloudDialogflowV2HumanAgentAssistantConfigArgs();
    }
}
