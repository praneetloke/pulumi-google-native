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
    /// Config for suggestion query.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Confidence threshold of query result. Agent Assist gives each suggestion a score in the range [0.0, 1.0], based on the relevance between the suggestion and the current conversation context. A score of 0.0 has no relevance, while a score of 1.0 has high relevance. Only suggestions with a score greater than or equal to the value of this field are included in the results. For a baseline model (the default), the recommended value is in the range [0.05, 0.1]. For a custom model, there is no recommended value. Tune this value by starting from a very low value and slowly increasing until you have desired results. If this field is not set, it defaults to 0.0, which means that all suggestions are returned. Supported features: ARTICLE_SUGGESTION, FAQ, SMART_REPLY, SMART_COMPOSE.
        /// </summary>
        [Input("confidenceThreshold")]
        public Input<double>? ConfidenceThreshold { get; set; }

        /// <summary>
        /// Determines how recent conversation context is filtered when generating suggestions. If unspecified, no messages will be dropped.
        /// </summary>
        [Input("contextFilterSettings")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigContextFilterSettingsArgs>? ContextFilterSettings { get; set; }

        /// <summary>
        /// Query from Dialogflow agent. It is used by DIALOGFLOW_ASSIST.
        /// </summary>
        [Input("dialogflowQuerySource")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigDialogflowQuerySourceArgs>? DialogflowQuerySource { get; set; }

        /// <summary>
        /// Query from knowledge base document. It is used by: SMART_REPLY, SMART_COMPOSE.
        /// </summary>
        [Input("documentQuerySource")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigDocumentQuerySourceArgs>? DocumentQuerySource { get; set; }

        /// <summary>
        /// Query from knowledgebase. It is used by: ARTICLE_SUGGESTION, FAQ.
        /// </summary>
        [Input("knowledgeBaseQuerySource")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigKnowledgeBaseQuerySourceArgs>? KnowledgeBaseQuerySource { get; set; }

        /// <summary>
        /// Maximum number of results to return. Currently, if unset, defaults to 10. And the max number is 20.
        /// </summary>
        [Input("maxResults")]
        public Input<int>? MaxResults { get; set; }

        public GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigArgs()
        {
        }
        public static new GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigArgs Empty => new GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigArgs();
    }
}
