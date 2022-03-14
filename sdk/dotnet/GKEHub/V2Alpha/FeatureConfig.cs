// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V2Alpha
{
    /// <summary>
    /// Creates FeatureConfig under a given parent.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:gkehub/v2alpha:FeatureConfig")]
    public partial class FeatureConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// When the FeatureConfig resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// When the FeatureConfig resource was deleted.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// GCP labels for this FeatureConfig.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Resource name of this FeatureConfig, in the format: `projects/{project}/locations/global/FeatureConfigs/{feature_type}/{feature_config}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Input only. Immutable. User input of feature spec. Note that this field is immutable. Must create a new FeatureConfig if a new feature spec is needed.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.FeatureSpecResponse> Spec { get; private set; } = null!;

        /// <summary>
        /// Lifecycle information of the FeatureConfig.
        /// </summary>
        [Output("state")]
        public Output<Outputs.FeatureConfigStateResponse> State { get; private set; } = null!;

        /// <summary>
        /// Google-generated UUID for this resource. This is unique across all FeatureConfig resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.
        /// </summary>
        [Output("uniqueId")]
        public Output<string> UniqueId { get; private set; } = null!;

        /// <summary>
        /// When the FeatureConfig resource was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a FeatureConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FeatureConfig(string name, FeatureConfigArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:gkehub/v2alpha:FeatureConfig", name, args ?? new FeatureConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FeatureConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gkehub/v2alpha:FeatureConfig", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing FeatureConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FeatureConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FeatureConfig(name, id, options);
        }
    }

    public sealed class FeatureConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the feature config to create.
        /// </summary>
        [Input("featureConfigId")]
        public Input<string>? FeatureConfigId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// GCP labels for this FeatureConfig.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Idempotent request UUID.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Input only. Immutable. User input of feature spec. Note that this field is immutable. Must create a new FeatureConfig if a new feature spec is needed.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.FeatureSpecArgs>? Spec { get; set; }

        /// <summary>
        /// Lifecycle information of the FeatureConfig.
        /// </summary>
        [Input("state")]
        public Input<Inputs.FeatureConfigStateArgs>? State { get; set; }

        public FeatureConfigArgs()
        {
        }
    }
}
