// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1
{
    /// <summary>
    /// Creates an TransitionRouteGroup in the specified flow. Note: You should always train a flow prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3beta1:TransitionRouteGroup")]
    public partial class TransitionRouteGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Transition routes associated with the TransitionRouteGroup.
        /// </summary>
        [Output("transitionRoutes")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse>> TransitionRoutes { get; private set; } = null!;


        /// <summary>
        /// Create a TransitionRouteGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitionRouteGroup(string name, TransitionRouteGroupArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:TransitionRouteGroup", name, args ?? new TransitionRouteGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitionRouteGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3beta1:TransitionRouteGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TransitionRouteGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitionRouteGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TransitionRouteGroup(name, id, options);
        }
    }

    public sealed class TransitionRouteGroupArgs : Pulumi.ResourceArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        /// <summary>
        /// The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("flowId", required: true)]
        public Input<string> FlowId { get; set; } = null!;

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("transitionRoutes")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3beta1TransitionRouteArgs>? _transitionRoutes;

        /// <summary>
        /// Transition routes associated with the TransitionRouteGroup.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3beta1TransitionRouteArgs> TransitionRoutes
        {
            get => _transitionRoutes ?? (_transitionRoutes = new InputList<Inputs.GoogleCloudDialogflowCxV3beta1TransitionRouteArgs>());
            set => _transitionRoutes = value;
        }

        public TransitionRouteGroupArgs()
        {
        }
    }
}
