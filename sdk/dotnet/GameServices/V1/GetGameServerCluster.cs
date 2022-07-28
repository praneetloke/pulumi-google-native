// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1
{
    public static class GetGameServerCluster
    {
        /// <summary>
        /// Gets details of a single game server cluster.
        /// </summary>
        public static Task<GetGameServerClusterResult> InvokeAsync(GetGameServerClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGameServerClusterResult>("google-native:gameservices/v1:getGameServerCluster", args ?? new GetGameServerClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single game server cluster.
        /// </summary>
        public static Output<GetGameServerClusterResult> Invoke(GetGameServerClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGameServerClusterResult>("google-native:gameservices/v1:getGameServerCluster", args ?? new GetGameServerClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGameServerClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("gameServerClusterId", required: true)]
        public string GameServerClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        [Input("view")]
        public string? View { get; set; }

        public GetGameServerClusterArgs()
        {
        }
        public static new GetGameServerClusterArgs Empty => new GetGameServerClusterArgs();
    }

    public sealed class GetGameServerClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("gameServerClusterId", required: true)]
        public Input<string> GameServerClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetGameServerClusterInvokeArgs()
        {
        }
        public static new GetGameServerClusterInvokeArgs Empty => new GetGameServerClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetGameServerClusterResult
    {
        /// <summary>
        /// The state of the Kubernetes cluster in preview. This will be available if view is set to FULL in the relevant list/get/preview request.
        /// </summary>
        public readonly Outputs.KubernetesClusterStateResponse ClusterState;
        /// <summary>
        /// The game server cluster connection information. This information is used to manage game server clusters.
        /// </summary>
        public readonly Outputs.GameServerClusterConnectionInfoResponse ConnectionInfo;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Human readable description of the cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The labels associated with this game server cluster. Each label is a key-value pair.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name of the game server cluster, in the following form: `projects/{project}/locations/{locationId}/realms/{realmId}/gameServerClusters/{gameServerClusterId}`. For example, `projects/my-project/locations/global/realms/zanzibar/gameServerClusters/my-gke-cluster`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The last-modified time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetGameServerClusterResult(
            Outputs.KubernetesClusterStateResponse clusterState,

            Outputs.GameServerClusterConnectionInfoResponse connectionInfo,

            string createTime,

            string description,

            string etag,

            ImmutableDictionary<string, string> labels,

            string name,

            string updateTime)
        {
            ClusterState = clusterState;
            ConnectionInfo = connectionInfo;
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            Labels = labels;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
