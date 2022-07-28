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
    /// Configuration for webhooks.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1EnvironmentWebhookConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("webhookOverrides")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3beta1WebhookArgs>? _webhookOverrides;

        /// <summary>
        /// The list of webhooks to override for the agent environment. The webhook must exist in the agent. You can override fields in `generic_web_service` and `service_directory`.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3beta1WebhookArgs> WebhookOverrides
        {
            get => _webhookOverrides ?? (_webhookOverrides = new InputList<Inputs.GoogleCloudDialogflowCxV3beta1WebhookArgs>());
            set => _webhookOverrides = value;
        }

        public GoogleCloudDialogflowCxV3beta1EnvironmentWebhookConfigArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1EnvironmentWebhookConfigArgs Empty => new GoogleCloudDialogflowCxV3beta1EnvironmentWebhookConfigArgs();
    }
}
