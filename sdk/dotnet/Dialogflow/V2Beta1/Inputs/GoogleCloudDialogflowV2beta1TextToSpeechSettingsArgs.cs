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
    /// Instructs the speech synthesizer on how to generate the output audio content.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1TextToSpeechSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Indicates whether text to speech is enabled. Even when this field is false, other settings in this proto are still retained.
        /// </summary>
        [Input("enableTextToSpeech")]
        public Input<bool>? EnableTextToSpeech { get; set; }

        /// <summary>
        /// Required. Audio encoding of the synthesized audio content.
        /// </summary>
        [Input("outputAudioEncoding")]
        public Input<string>? OutputAudioEncoding { get; set; }

        /// <summary>
        /// Optional. The synthesis sample rate (in hertz) for this audio. If not provided, then the synthesizer will use the default sample rate based on the audio encoding. If this is different from the voice's natural sample rate, then the synthesizer will honor this request by converting to the desired sample rate (which might result in worse audio quality).
        /// </summary>
        [Input("sampleRateHertz")]
        public Input<int>? SampleRateHertz { get; set; }

        [Input("synthesizeSpeechConfigs")]
        private InputMap<string>? _synthesizeSpeechConfigs;

        /// <summary>
        /// Optional. Configuration of how speech should be synthesized, mapping from language (https://cloud.google.com/dialogflow/docs/reference/language) to SynthesizeSpeechConfig.
        /// </summary>
        public InputMap<string> SynthesizeSpeechConfigs
        {
            get => _synthesizeSpeechConfigs ?? (_synthesizeSpeechConfigs = new InputMap<string>());
            set => _synthesizeSpeechConfigs = value;
        }

        public GoogleCloudDialogflowV2beta1TextToSpeechSettingsArgs()
        {
        }
    }
}
