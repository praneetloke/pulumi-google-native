// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1Beta
{
    /// <summary>
    /// Creates a new game server config in a given project, location, and game server deployment. Game server configs are immutable, and are not applied until referenced in the game server deployment rollout resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:gameservices/v1beta:Config")]
    public partial class Config : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the game server config.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// FleetConfig contains a list of Agones fleet specs. Only one FleetConfig is allowed.
        /// </summary>
        [Output("fleetConfigs")]
        public Output<ImmutableArray<Outputs.FleetConfigResponse>> FleetConfigs { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this game server config. Each label is a key-value pair.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name of the game server config, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/{config}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The autoscaling settings.
        /// </summary>
        [Output("scalingConfigs")]
        public Output<ImmutableArray<Outputs.ScalingConfigResponse>> ScalingConfigs { get; private set; } = null!;

        /// <summary>
        /// The last-modified time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Config resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Config(string name, ConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:gameservices/v1beta:Config", name, args ?? new ConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Config(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gameservices/v1beta:Config", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Config resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Config Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Config(name, id, options);
        }
    }

    public sealed class ConfigArgs : Pulumi.ResourceArgs
    {
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        /// <summary>
        /// The description of the game server config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fleetConfigs")]
        private InputList<Inputs.FleetConfigArgs>? _fleetConfigs;

        /// <summary>
        /// FleetConfig contains a list of Agones fleet specs. Only one FleetConfig is allowed.
        /// </summary>
        public InputList<Inputs.FleetConfigArgs> FleetConfigs
        {
            get => _fleetConfigs ?? (_fleetConfigs = new InputList<Inputs.FleetConfigArgs>());
            set => _fleetConfigs = value;
        }

        [Input("gameServerDeploymentId", required: true)]
        public Input<string> GameServerDeploymentId { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this game server config. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the game server config, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/{config}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("scalingConfigs")]
        private InputList<Inputs.ScalingConfigArgs>? _scalingConfigs;

        /// <summary>
        /// The autoscaling settings.
        /// </summary>
        public InputList<Inputs.ScalingConfigArgs> ScalingConfigs
        {
            get => _scalingConfigs ?? (_scalingConfigs = new InputList<Inputs.ScalingConfigArgs>());
            set => _scalingConfigs = value;
        }

        public ConfigArgs()
        {
        }
    }
}
