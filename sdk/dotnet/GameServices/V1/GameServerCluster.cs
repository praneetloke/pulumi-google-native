// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1
{
    /// <summary>
    /// Creates a new game server cluster in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:gameservices/v1:GameServerCluster")]
    public partial class GameServerCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The state of the Kubernetes cluster in preview. This will be available if view is set to FULL in the relevant list/get/preview request.
        /// </summary>
        [Output("clusterState")]
        public Output<Outputs.KubernetesClusterStateResponse> ClusterState { get; private set; } = null!;

        /// <summary>
        /// The game server cluster connection information. This information is used to manage game server clusters.
        /// </summary>
        [Output("connectionInfo")]
        public Output<Outputs.GameServerClusterConnectionInfoResponse> ConnectionInfo { get; private set; } = null!;

        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Human readable description of the cluster.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this game server cluster. Each label is a key-value pair.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name of the game server cluster, in the following form: `projects/{project}/locations/{locationId}/realms/{realmId}/gameServerClusters/{gameServerClusterId}`. For example, `projects/my-project/locations/global/realms/zanzibar/gameServerClusters/my-gke-cluster`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The last-modified time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a GameServerCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GameServerCluster(string name, GameServerClusterArgs args, CustomResourceOptions? options = null)
            : base("google-native:gameservices/v1:GameServerCluster", name, args ?? new GameServerClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GameServerCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gameservices/v1:GameServerCluster", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GameServerCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GameServerCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GameServerCluster(name, id, options);
        }
    }

    public sealed class GameServerClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The game server cluster connection information. This information is used to manage game server clusters.
        /// </summary>
        [Input("connectionInfo")]
        public Input<Inputs.GameServerClusterConnectionInfoArgs>? ConnectionInfo { get; set; }

        /// <summary>
        /// Human readable description of the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Required. The ID of the game server cluster resource to create.
        /// </summary>
        [Input("gameServerClusterId", required: true)]
        public Input<string> GameServerClusterId { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this game server cluster. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the game server cluster, in the following form: `projects/{project}/locations/{locationId}/realms/{realmId}/gameServerClusters/{gameServerClusterId}`. For example, `projects/my-project/locations/global/realms/zanzibar/gameServerClusters/my-gke-cluster`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GameServerClusterArgs()
        {
        }
    }
}
