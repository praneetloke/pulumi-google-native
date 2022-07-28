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
    /// The collection of simple response candidates. This message in `QueryResult.fulfillment_messages` and `WebhookResponse.fulfillment_messages` should contain only one `SimpleResponse`.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageSimpleResponsesArgs : global::Pulumi.ResourceArgs
    {
        [Input("simpleResponses", required: true)]
        private InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageSimpleResponseArgs>? _simpleResponses;

        /// <summary>
        /// The list of simple responses.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageSimpleResponseArgs> SimpleResponses
        {
            get => _simpleResponses ?? (_simpleResponses = new InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageSimpleResponseArgs>());
            set => _simpleResponses = value;
        }

        public GoogleCloudDialogflowV2beta1IntentMessageSimpleResponsesArgs()
        {
        }
        public static new GoogleCloudDialogflowV2beta1IntentMessageSimpleResponsesArgs Empty => new GoogleCloudDialogflowV2beta1IntentMessageSimpleResponsesArgs();
    }
}
