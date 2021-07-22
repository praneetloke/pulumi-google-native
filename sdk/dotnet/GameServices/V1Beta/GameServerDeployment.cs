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
    /// Creates a new game server deployment in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:gameservices/v1beta:GameServerDeployment")]
    public partial class GameServerDeployment : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Human readable description of the game server delpoyment.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// ETag of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this game server deployment. Each label is a key-value pair.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name of the game server deployment, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-deployment`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The last-modified time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a GameServerDeployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GameServerDeployment(string name, GameServerDeploymentArgs args, CustomResourceOptions? options = null)
            : base("google-native:gameservices/v1beta:GameServerDeployment", name, args ?? new GameServerDeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GameServerDeployment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gameservices/v1beta:GameServerDeployment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GameServerDeployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GameServerDeployment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GameServerDeployment(name, id, options);
        }
    }

    public sealed class GameServerDeploymentArgs : Pulumi.ResourceArgs
    {
        [Input("deploymentId", required: true)]
        public Input<string> DeploymentId { get; set; } = null!;

        /// <summary>
        /// Human readable description of the game server delpoyment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ETag of the resource.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this game server deployment. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the game server deployment, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-deployment`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GameServerDeploymentArgs()
        {
        }
    }
}
