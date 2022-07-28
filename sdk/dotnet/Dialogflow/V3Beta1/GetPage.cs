// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1
{
    public static class GetPage
    {
        /// <summary>
        /// Retrieves the specified page.
        /// </summary>
        public static Task<GetPageResult> InvokeAsync(GetPageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPageResult>("google-native:dialogflow/v3beta1:getPage", args ?? new GetPageArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified page.
        /// </summary>
        public static Output<GetPageResult> Invoke(GetPageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPageResult>("google-native:dialogflow/v3beta1:getPage", args ?? new GetPageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPageArgs : global::Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        [Input("flowId", required: true)]
        public string FlowId { get; set; } = null!;

        [Input("languageCode")]
        public string? LanguageCode { get; set; }

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("pageId", required: true)]
        public string PageId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetPageArgs()
        {
        }
        public static new GetPageArgs Empty => new GetPageArgs();
    }

    public sealed class GetPageInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("flowId", required: true)]
        public Input<string> FlowId { get; set; } = null!;

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("pageId", required: true)]
        public Input<string> PageId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetPageInvokeArgs()
        {
        }
        public static new GetPageInvokeArgs Empty => new GetPageInvokeArgs();
    }


    [OutputType]
    public sealed class GetPageResult
    {
        /// <summary>
        /// The human-readable name of the page, unique within the flow.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The fulfillment to call when the session is entering the page.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3beta1FulfillmentResponse EntryFulfillment;
        /// <summary>
        /// Handlers associated with the page to handle events such as webhook errors, no match or no input.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EventHandlerResponse> EventHandlers;
        /// <summary>
        /// The form associated with the page, used for collecting parameters relevant to the page.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3beta1FormResponse Form;
        /// <summary>
        /// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -&gt; page's transition route group -&gt; flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        public readonly ImmutableArray<string> TransitionRouteGroups;
        /// <summary>
        /// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse> TransitionRoutes;

        [OutputConstructor]
        private GetPageResult(
            string displayName,

            Outputs.GoogleCloudDialogflowCxV3beta1FulfillmentResponse entryFulfillment,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EventHandlerResponse> eventHandlers,

            Outputs.GoogleCloudDialogflowCxV3beta1FormResponse form,

            string name,

            ImmutableArray<string> transitionRouteGroups,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse> transitionRoutes)
        {
            DisplayName = displayName;
            EntryFulfillment = entryFulfillment;
            EventHandlers = eventHandlers;
            Form = form;
            Name = name;
            TransitionRouteGroups = transitionRouteGroups;
            TransitionRoutes = transitionRoutes;
        }
    }
}
