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
    /// Represents configuration for a [Service Directory](https://cloud.google.com/service-directory) service.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1WebhookServiceDirectoryConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Generic Service configuration of this webhook.
        /// </summary>
        [Input("genericWebService")]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1WebhookGenericWebServiceArgs>? GenericWebService { get; set; }

        /// <summary>
        /// The name of [Service Directory](https://cloud.google.com/service-directory) service. Format: `projects//locations//namespaces//services/`. `Location ID` of the service directory must be the same as the location of the agent.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public GoogleCloudDialogflowCxV3beta1WebhookServiceDirectoryConfigArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1WebhookServiceDirectoryConfigArgs Empty => new GoogleCloudDialogflowCxV3beta1WebhookServiceDirectoryConfigArgs();
    }
}
