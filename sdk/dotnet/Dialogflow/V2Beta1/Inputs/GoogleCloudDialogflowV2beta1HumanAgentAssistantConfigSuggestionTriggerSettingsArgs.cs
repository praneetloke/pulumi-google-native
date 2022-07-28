// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Inputs
{

    /// <summary>
    /// Settings of suggestion trigger.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionTriggerSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Do not trigger if last utterance is small talk.
        /// </summary>
        [Input("noSmallTalk")]
        public Input<bool>? NoSmallTalk { get; set; }

        /// <summary>
        /// Only trigger suggestion if participant role of last utterance is END_USER.
        /// </summary>
        [Input("onlyEndUser")]
        public Input<bool>? OnlyEndUser { get; set; }

        public GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionTriggerSettingsArgs()
        {
        }
        public static new GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionTriggerSettingsArgs Empty => new GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionTriggerSettingsArgs();
    }
}
