// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1
{
    /// <summary>
    /// Adds a new Feature.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:gkehub/v1:Feature")]
    public partial class Feature : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When the Feature resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// When the Feature resource was deleted.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the feature to create.
        /// </summary>
        [Output("featureId")]
        public Output<string?> FeatureId { get; private set; } = null!;

        /// <summary>
        /// GCP labels for this Feature.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        /// </summary>
        [Output("membershipSpecs")]
        public Output<ImmutableDictionary<string, string>> MembershipSpecs { get; private set; } = null!;

        /// <summary>
        /// Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        /// </summary>
        [Output("membershipStates")]
        public Output<ImmutableDictionary<string, string>> MembershipStates { get; private set; } = null!;

        /// <summary>
        /// The full, unique name of this Feature resource in the format `projects/*/locations/*/features/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// State of the Feature resource itself.
        /// </summary>
        [Output("resourceState")]
        public Output<Outputs.FeatureResourceStateResponse> ResourceState { get; private set; } = null!;

        /// <summary>
        /// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.CommonFeatureSpecResponse> Spec { get; private set; } = null!;

        /// <summary>
        /// The Hub-wide Feature state.
        /// </summary>
        [Output("state")]
        public Output<Outputs.CommonFeatureStateResponse> State { get; private set; } = null!;

        /// <summary>
        /// When the Feature resource was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Feature resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Feature(string name, FeatureArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:gkehub/v1:Feature", name, args ?? new FeatureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Feature(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gkehub/v1:Feature", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Feature resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Feature Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Feature(name, id, options);
        }
    }

    public sealed class FeatureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the feature to create.
        /// </summary>
        [Input("featureId")]
        public Input<string>? FeatureId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// GCP labels for this Feature.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("membershipSpecs")]
        private InputMap<string>? _membershipSpecs;

        /// <summary>
        /// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        /// </summary>
        public InputMap<string> MembershipSpecs
        {
            get => _membershipSpecs ?? (_membershipSpecs = new InputMap<string>());
            set => _membershipSpecs = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.CommonFeatureSpecArgs>? Spec { get; set; }

        public FeatureArgs()
        {
        }
        public static new FeatureArgs Empty => new FeatureArgs();
    }
}
