// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    /// <summary>
    /// Represents an example that the agent is trained on to identify the intent. Next ID: 15
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponse
    {
        /// <summary>
        /// The ordered list of training phrase parts. The parts are concatenated in order to form the training phrase. Note: The API does not automatically annotate training phrases like the Dialogflow Console does. Note: Do not forget to include whitespace at part boundaries, so the training phrase is well formatted when the parts are concatenated. If the training phrase does not need to be annotated with parameters, you just need a single part with only the Part.text field set. If you want to annotate the training phrase, you must create multiple parts, where the fields of each part are populated in one of two ways: - `Part.text` is set to a part of the phrase that has no parameters. - `Part.text` is set to a part of the phrase that you want to annotate, and the `parameter_id` field is set.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1IntentTrainingPhrasePartResponse> Parts;
        /// <summary>
        /// Indicates how many times this example was added to the intent.
        /// </summary>
        public readonly int RepeatCount;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponse(
            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1IntentTrainingPhrasePartResponse> parts,

            int repeatCount)
        {
            Parts = parts;
            RepeatCount = repeatCount;
        }
    }
}
