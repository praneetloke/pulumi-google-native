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
    /// Settings related to speech recognition.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use speech adaptation for speech recognition.
        /// </summary>
        [Input("enableSpeechAdaptation")]
        public Input<bool>? EnableSpeechAdaptation { get; set; }

        public GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsArgs Empty => new GoogleCloudDialogflowCxV3beta1SpeechToTextSettingsArgs();
    }
}
