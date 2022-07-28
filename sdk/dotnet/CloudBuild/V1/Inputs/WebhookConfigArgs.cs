// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// WebhookConfig describes the configuration of a trigger that creates a build whenever a webhook is sent to a trigger's webhook URL.
    /// </summary>
    public sealed class WebhookConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource name for the secret required as a URL parameter.
        /// </summary>
        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        /// <summary>
        /// Potential issues with the underlying Pub/Sub subscription configuration. Only populated on get requests.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.CloudBuild.V1.WebhookConfigState>? State { get; set; }

        public WebhookConfigArgs()
        {
        }
        public static new WebhookConfigArgs Empty => new WebhookConfigArgs();
    }
}
