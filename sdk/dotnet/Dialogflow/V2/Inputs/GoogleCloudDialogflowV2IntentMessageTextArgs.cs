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
    /// The text response message.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2IntentMessageTextArgs : global::Pulumi.ResourceArgs
    {
        [Input("text")]
        private InputList<string>? _text;

        /// <summary>
        /// Optional. The collection of the agent's responses.
        /// </summary>
        public InputList<string> Text
        {
            get => _text ?? (_text = new InputList<string>());
            set => _text = value;
        }

        public GoogleCloudDialogflowV2IntentMessageTextArgs()
        {
        }
        public static new GoogleCloudDialogflowV2IntentMessageTextArgs Empty => new GoogleCloudDialogflowV2IntentMessageTextArgs();
    }
}
