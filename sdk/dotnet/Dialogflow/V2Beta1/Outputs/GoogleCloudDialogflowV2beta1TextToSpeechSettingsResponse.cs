// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Outputs
{

    /// <summary>
    /// Instructs the speech synthesizer on how to generate the output audio content.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1TextToSpeechSettingsResponse
    {
        /// <summary>
        /// Optional. Indicates whether text to speech is enabled. Even when this field is false, other settings in this proto are still retained.
        /// </summary>
        public readonly bool EnableTextToSpeech;
        /// <summary>
        /// Audio encoding of the synthesized audio content.
        /// </summary>
        public readonly string OutputAudioEncoding;
        /// <summary>
        /// Optional. The synthesis sample rate (in hertz) for this audio. If not provided, then the synthesizer will use the default sample rate based on the audio encoding. If this is different from the voice's natural sample rate, then the synthesizer will honor this request by converting to the desired sample rate (which might result in worse audio quality).
        /// </summary>
        public readonly int SampleRateHertz;
        /// <summary>
        /// Optional. Configuration of how speech should be synthesized, mapping from language (https://cloud.google.com/dialogflow/docs/reference/language) to SynthesizeSpeechConfig.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponse SynthesizeSpeechConfigs;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1TextToSpeechSettingsResponse(
            bool enableTextToSpeech,

            string outputAudioEncoding,

            int sampleRateHertz,

            Outputs.GoogleCloudDialogflowV2beta1SynthesizeSpeechConfigResponse synthesizeSpeechConfigs)
        {
            EnableTextToSpeech = enableTextToSpeech;
            OutputAudioEncoding = outputAudioEncoding;
            SampleRateHertz = sampleRateHertz;
            SynthesizeSpeechConfigs = synthesizeSpeechConfigs;
        }
    }
}
