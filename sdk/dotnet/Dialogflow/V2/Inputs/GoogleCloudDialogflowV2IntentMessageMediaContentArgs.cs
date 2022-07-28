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
    /// The media content card for Actions on Google.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2IntentMessageMediaContentArgs : global::Pulumi.ResourceArgs
    {
        [Input("mediaObjects", required: true)]
        private InputList<Inputs.GoogleCloudDialogflowV2IntentMessageMediaContentResponseMediaObjectArgs>? _mediaObjects;

        /// <summary>
        /// List of media objects.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2IntentMessageMediaContentResponseMediaObjectArgs> MediaObjects
        {
            get => _mediaObjects ?? (_mediaObjects = new InputList<Inputs.GoogleCloudDialogflowV2IntentMessageMediaContentResponseMediaObjectArgs>());
            set => _mediaObjects = value;
        }

        /// <summary>
        /// Optional. What type of media is the content (ie "audio").
        /// </summary>
        [Input("mediaType")]
        public Input<Pulumi.GoogleNative.Dialogflow.V2.GoogleCloudDialogflowV2IntentMessageMediaContentMediaType>? MediaType { get; set; }

        public GoogleCloudDialogflowV2IntentMessageMediaContentArgs()
        {
        }
        public static new GoogleCloudDialogflowV2IntentMessageMediaContentArgs Empty => new GoogleCloudDialogflowV2IntentMessageMediaContentArgs();
    }
}
