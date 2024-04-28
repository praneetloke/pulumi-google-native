// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Inputs
{

    /// <summary>
    /// The input from the human user.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1ConversationTurnUserInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether sentiment analysis is enabled.
        /// </summary>
        [Input("enableSentimentAnalysis")]
        public Input<bool>? EnableSentimentAnalysis { get; set; }

        [Input("injectedParameters")]
        private InputMap<object>? _injectedParameters;

        /// <summary>
        /// Parameters that need to be injected into the conversation during intent detection.
        /// </summary>
        public InputMap<object> InjectedParameters
        {
            get => _injectedParameters ?? (_injectedParameters = new InputMap<object>());
            set => _injectedParameters = value;
        }

        /// <summary>
        /// Supports text input, event input, dtmf input in the test case.
        /// </summary>
        [Input("input")]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1QueryInputArgs>? Input { get; set; }

        /// <summary>
        /// If webhooks should be allowed to trigger in response to the user utterance. Often if parameters are injected, webhooks should not be enabled.
        /// </summary>
        [Input("isWebhookEnabled")]
        public Input<bool>? IsWebhookEnabled { get; set; }

        public GoogleCloudDialogflowCxV3beta1ConversationTurnUserInputArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1ConversationTurnUserInputArgs Empty => new GoogleCloudDialogflowCxV3beta1ConversationTurnUserInputArgs();
    }
}
