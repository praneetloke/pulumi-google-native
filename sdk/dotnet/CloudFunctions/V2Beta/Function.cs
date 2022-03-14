// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2Beta
{
    /// <summary>
    /// Creates a new function. If a function with the given name already exists in the specified project, the long running operation will return `ALREADY_EXISTS` error.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudfunctions/v2beta:Function")]
    public partial class Function : Pulumi.CustomResource
    {
        /// <summary>
        /// Describes the Build step of the function that builds a container from the given source.
        /// </summary>
        [Output("buildConfig")]
        public Output<Outputs.BuildConfigResponse> BuildConfig { get; private set; } = null!;

        /// <summary>
        /// User-provided description of a function.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Describe whether the function is gen1 or gen2.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
        /// </summary>
        [Output("eventTrigger")]
        public Output<Outputs.EventTriggerResponse> EventTrigger { get; private set; } = null!;

        /// <summary>
        /// Labels associated with this Cloud Function.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
        /// </summary>
        [Output("serviceConfig")]
        public Output<Outputs.ServiceConfigResponse> ServiceConfig { get; private set; } = null!;

        /// <summary>
        /// State of the function.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// State Messages for this Cloud Function.
        /// </summary>
        [Output("stateMessages")]
        public Output<ImmutableArray<Outputs.GoogleCloudFunctionsV2betaStateMessageResponse>> StateMessages { get; private set; } = null!;

        /// <summary>
        /// The last update timestamp of a Cloud Function.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Function resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Function(string name, FunctionArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:cloudfunctions/v2beta:Function", name, args ?? new FunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Function(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudfunctions/v2beta:Function", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Function resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Function Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Function(name, id, options);
        }
    }

    public sealed class FunctionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the Build step of the function that builds a container from the given source.
        /// </summary>
        [Input("buildConfig")]
        public Input<Inputs.BuildConfigArgs>? BuildConfig { get; set; }

        /// <summary>
        /// User-provided description of a function.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Describe whether the function is gen1 or gen2.
        /// </summary>
        [Input("environment")]
        public Input<Pulumi.GoogleNative.CloudFunctions.V2Beta.FunctionEnvironment>? Environment { get; set; }

        /// <summary>
        /// An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
        /// </summary>
        [Input("eventTrigger")]
        public Input<Inputs.EventTriggerArgs>? EventTrigger { get; set; }

        /// <summary>
        /// The ID to use for the function, which will become the final component of the function's resource name. This value should be 4-63 characters, and valid characters are /a-z-/.
        /// </summary>
        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels associated with this Cloud Function.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
        /// </summary>
        [Input("serviceConfig")]
        public Input<Inputs.ServiceConfigArgs>? ServiceConfig { get; set; }

        public FunctionArgs()
        {
        }
    }
}
