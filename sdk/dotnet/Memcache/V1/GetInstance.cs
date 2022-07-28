// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Memcache.V1
{
    public static class GetInstance
    {
        /// <summary>
        /// Gets details of a single Instance.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:memcache/v1:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Instance.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("google-native:memcache/v1:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.
        /// </summary>
        public readonly string AuthorizedNetwork;
        /// <summary>
        /// The time the instance was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Endpoint for the Discovery API.
        /// </summary>
        public readonly string DiscoveryEndpoint;
        /// <summary>
        /// User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// List of messages that describe the current state of the Memcached instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceMessageResponse> InstanceMessages;
        /// <summary>
        /// Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The maintenance policy for the instance. If not provided, the maintenance event will be performed based on Memorystore internal rollout schedule.
        /// </summary>
        public readonly Outputs.GoogleCloudMemcacheV1MaintenancePolicyResponse MaintenancePolicy;
        /// <summary>
        /// Published maintenance schedule.
        /// </summary>
        public readonly Outputs.MaintenanceScheduleResponse MaintenanceSchedule;
        /// <summary>
        /// The full version of memcached server running on this instance. System automatically determines the full memcached version for an instance based on the input MemcacheVersion. The full version format will be "memcached-1.5.16".
        /// </summary>
        public readonly string MemcacheFullVersion;
        /// <summary>
        /// List of Memcached nodes. Refer to Node message for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodeResponse> MemcacheNodes;
        /// <summary>
        /// The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
        /// </summary>
        public readonly string MemcacheVersion;
        /// <summary>
        /// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configuration for Memcached nodes.
        /// </summary>
        public readonly Outputs.NodeConfigResponse NodeConfig;
        /// <summary>
        /// Number of nodes in the Memcached instance.
        /// </summary>
        public readonly int NodeCount;
        /// <summary>
        /// User defined parameters to apply to the memcached process on each node.
        /// </summary>
        public readonly Outputs.MemcacheParametersResponse Parameters;
        /// <summary>
        /// The state of this Memcached instance.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The time the instance was updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetInstanceResult(
            string authorizedNetwork,

            string createTime,

            string discoveryEndpoint,

            string displayName,

            ImmutableArray<Outputs.InstanceMessageResponse> instanceMessages,

            ImmutableDictionary<string, string> labels,

            Outputs.GoogleCloudMemcacheV1MaintenancePolicyResponse maintenancePolicy,

            Outputs.MaintenanceScheduleResponse maintenanceSchedule,

            string memcacheFullVersion,

            ImmutableArray<Outputs.NodeResponse> memcacheNodes,

            string memcacheVersion,

            string name,

            Outputs.NodeConfigResponse nodeConfig,

            int nodeCount,

            Outputs.MemcacheParametersResponse parameters,

            string state,

            string updateTime,

            ImmutableArray<string> zones)
        {
            AuthorizedNetwork = authorizedNetwork;
            CreateTime = createTime;
            DiscoveryEndpoint = discoveryEndpoint;
            DisplayName = displayName;
            InstanceMessages = instanceMessages;
            Labels = labels;
            MaintenancePolicy = maintenancePolicy;
            MaintenanceSchedule = maintenanceSchedule;
            MemcacheFullVersion = memcacheFullVersion;
            MemcacheNodes = memcacheNodes;
            MemcacheVersion = memcacheVersion;
            Name = name;
            NodeConfig = nodeConfig;
            NodeCount = nodeCount;
            Parameters = parameters;
            State = state;
            UpdateTime = updateTime;
            Zones = zones;
        }
    }
}
