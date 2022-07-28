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
    /// A text or ssml response that is preferentially used for TTS output audio synthesis, as described in the comment on the ResponseMessage message.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1ResponseMessageOutputAudioTextArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SSML text to be synthesized. For more information, see [SSML](/speech/text-to-speech/docs/ssml).
        /// </summary>
        [Input("ssml")]
        public Input<string>? Ssml { get; set; }

        /// <summary>
        /// The raw text to be synthesized.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        public GoogleCloudDialogflowCxV3beta1ResponseMessageOutputAudioTextArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1ResponseMessageOutputAudioTextArgs Empty => new GoogleCloudDialogflowCxV3beta1ResponseMessageOutputAudioTextArgs();
    }
}
