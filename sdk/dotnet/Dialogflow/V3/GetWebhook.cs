// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3
{
    public static class GetWebhook
    {
        /// <summary>
        /// Retrieves the specified webhook.
        /// </summary>
        public static Task<GetWebhookResult> InvokeAsync(GetWebhookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebhookResult>("google-native:dialogflow/v3:getWebhook", args ?? new GetWebhookArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified webhook.
        /// </summary>
        public static Output<GetWebhookResult> Invoke(GetWebhookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebhookResult>("google-native:dialogflow/v3:getWebhook", args ?? new GetWebhookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebhookArgs : global::Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("webhookId", required: true)]
        public string WebhookId { get; set; } = null!;

        public GetWebhookArgs()
        {
        }
        public static new GetWebhookArgs Empty => new GetWebhookArgs();
    }

    public sealed class GetWebhookInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("webhookId", required: true)]
        public Input<string> WebhookId { get; set; } = null!;

        public GetWebhookInvokeArgs()
        {
        }
        public static new GetWebhookInvokeArgs Empty => new GetWebhookInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebhookResult
    {
        /// <summary>
        /// Indicates whether the webhook is disabled.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// The human-readable name of the webhook, unique within the agent.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Configuration for a generic web service.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3WebhookGenericWebServiceResponse GenericWebService;
        /// <summary>
        /// The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configuration for a [Service Directory](https://cloud.google.com/service-directory) service.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3WebhookServiceDirectoryConfigResponse ServiceDirectory;
        /// <summary>
        /// Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
        /// </summary>
        public readonly string Timeout;

        [OutputConstructor]
        private GetWebhookResult(
            bool disabled,

            string displayName,

            Outputs.GoogleCloudDialogflowCxV3WebhookGenericWebServiceResponse genericWebService,

            string name,

            Outputs.GoogleCloudDialogflowCxV3WebhookServiceDirectoryConfigResponse serviceDirectory,

            string timeout)
        {
            Disabled = disabled;
            DisplayName = displayName;
            GenericWebService = genericWebService;
            Name = name;
            ServiceDirectory = serviceDirectory;
            Timeout = timeout;
        }
    }
}
