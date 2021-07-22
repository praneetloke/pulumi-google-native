// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1
{
    /// <summary>
    /// Creates an agent environment.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v2beta1:Environment")]
    public partial class Environment : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The agent version loaded into this environment. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
        /// </summary>
        [Output("agentVersion")]
        public Output<string> AgentVersion { get; private set; } = null!;

        /// <summary>
        /// Optional. The developer-provided description for this environment. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. The fulfillment settings to use for this environment.
        /// </summary>
        [Output("fulfillment")]
        public Output<Outputs.GoogleCloudDialogflowV2beta1FulfillmentResponse> Fulfillment { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of this agent environment. Supported formats: - `projects//agent/environments/` - `projects//locations//agent/environments/`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The state of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Optional. Text to speech settings for this environment.
        /// </summary>
        [Output("textToSpeechSettings")]
        public Output<Outputs.GoogleCloudDialogflowV2beta1TextToSpeechSettingsResponse> TextToSpeechSettings { get; private set; } = null!;

        /// <summary>
        /// The last update time of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2beta1:Environment", name, args ?? new EnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2beta1:Environment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, options);
        }
    }

    public sealed class EnvironmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The agent version loaded into this environment. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
        /// </summary>
        [Input("agentVersion")]
        public Input<string>? AgentVersion { get; set; }

        /// <summary>
        /// Optional. The developer-provided description for this environment. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Optional. The fulfillment settings to use for this environment.
        /// </summary>
        [Input("fulfillment")]
        public Input<Inputs.GoogleCloudDialogflowV2beta1FulfillmentArgs>? Fulfillment { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. Text to speech settings for this environment.
        /// </summary>
        [Input("textToSpeechSettings")]
        public Input<Inputs.GoogleCloudDialogflowV2beta1TextToSpeechSettingsArgs>? TextToSpeechSettings { get; set; }

        public EnvironmentArgs()
        {
        }
    }
}
