// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3
{
    /// <summary>
    /// Creates an intent in the specified agent. Note: You should always train a flow prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3:Intent")]
    public partial class Intent : Pulumi.CustomResource
    {
        /// <summary>
        /// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The human-readable name of the intent, unique within the agent.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
        /// </summary>
        [Output("isFallback")]
        public Output<bool> IsFallback { get; private set; } = null!;

        /// <summary>
        /// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The collection of parameters associated with the intent.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3IntentParameterResponse>> Parameters { get; private set; } = null!;

        /// <summary>
        /// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The collection of training phrases the agent is trained on to identify the intent.
        /// </summary>
        [Output("trainingPhrases")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3IntentTrainingPhraseResponse>> TrainingPhrases { get; private set; } = null!;


        /// <summary>
        /// Create a Intent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Intent(string name, IntentArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3:Intent", name, args ?? new IntentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Intent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3:Intent", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Intent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Intent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Intent(name, id, options);
        }
    }

    public sealed class IntentArgs : Pulumi.ResourceArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

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
        /// The language of the following fields in `intent`: * `Intent.training_phrases.parts.text` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

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

        [Input("project")]
        public Input<string>? Project { get; set; }

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

        public IntentArgs()
        {
        }
    }
}
