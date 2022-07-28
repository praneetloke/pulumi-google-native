// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Inputs
{

    /// <summary>
    /// An intent represents a user's intent to interact with a conversational agent. You can provide information for the Dialogflow API to use to match user input to an intent by adding training phrases (i.e., examples of user input) to your intent.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3IntentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The human-readable name of the intent, unique within the agent.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
        /// </summary>
        [Input("isFallback")]
        public Input<bool>? IsFallback { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3IntentParameterArgs>? _parameters;

        /// <summary>
        /// The collection of parameters associated with the intent.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3IntentParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.GoogleCloudDialogflowCxV3IntentParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("trainingPhrases")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs>? _trainingPhrases;

        /// <summary>
        /// The collection of training phrases the agent is trained on to identify the intent.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs> TrainingPhrases
        {
            get => _trainingPhrases ?? (_trainingPhrases = new InputList<Inputs.GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs>());
            set => _trainingPhrases = value;
        }

        public GoogleCloudDialogflowCxV3IntentArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3IntentArgs Empty => new GoogleCloudDialogflowCxV3IntentArgs();
    }
}
