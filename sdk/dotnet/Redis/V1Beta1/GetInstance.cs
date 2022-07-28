// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Redis.V1Beta1
{
    public static class GetInstance
    {
        /// <summary>
        /// Gets the details of a specific Redis instance.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:redis/v1beta1:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of a specific Redis instance.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("google-native:redis/v1beta1:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
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
        /// Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
        /// </summary>
        public readonly string AlternativeLocationId;
        /// <summary>
        /// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
        /// </summary>
        public readonly bool AuthEnabled;
        /// <summary>
        /// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
        /// </summary>
        public readonly string AuthorizedNetwork;
        /// <summary>
        /// Optional. The available maintenance versions that an instance could update to.
        /// </summary>
        public readonly ImmutableArray<string> AvailableMaintenanceVersions;
        /// <summary>
        /// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
        /// </summary>
        public readonly string ConnectMode;
        /// <summary>
        /// The time the instance was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.
        /// </summary>
        public readonly string CurrentLocationId;
        /// <summary>
        /// Optional. The KMS key reference that the customer provides when trying to create the instance.
        /// </summary>
        public readonly string CustomerManagedKey;
        /// <summary>
        /// An arbitrary and optional user-provided name for the instance.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Resource labels to represent user provided metadata
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
        /// </summary>
        public readonly Outputs.MaintenancePolicyResponse MaintenancePolicy;
        /// <summary>
        /// Date and time of upcoming maintenance events which have been scheduled.
        /// </summary>
        public readonly Outputs.MaintenanceScheduleResponse MaintenanceSchedule;
        /// <summary>
        /// Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
        /// </summary>
        public readonly string MaintenanceVersion;
        /// <summary>
        /// Redis memory size in GiB.
        /// </summary>
        public readonly int MemorySizeGb;
        /// <summary>
        /// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Info per node.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodeInfoResponse> Nodes;
        /// <summary>
        /// Optional. Persistence configuration parameters
        /// </summary>
        public readonly Outputs.PersistenceConfigResponse PersistenceConfig;
        /// <summary>
        /// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
        /// </summary>
        public readonly string PersistenceIamIdentity;
        /// <summary>
        /// The port number of the exposed Redis endpoint.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.
        /// </summary>
        public readonly string ReadEndpoint;
        /// <summary>
        /// The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.
        /// </summary>
        public readonly int ReadEndpointPort;
        /// <summary>
        /// Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
        /// </summary>
        public readonly string ReadReplicasMode;
        /// <summary>
        /// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
        /// </summary>
        public readonly ImmutableDictionary<string, string> RedisConfigs;
        /// <summary>
        /// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
        /// </summary>
        public readonly string RedisVersion;
        /// <summary>
        /// Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
        /// </summary>
        public readonly int ReplicaCount;
        /// <summary>
        /// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
        /// </summary>
        public readonly string ReservedIpRange;
        /// <summary>
        /// Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
        /// </summary>
        public readonly string SecondaryIpRange;
        /// <summary>
        /// List of server CA certificates for the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.TlsCertificateResponse> ServerCaCerts;
        /// <summary>
        /// The current state of this instance.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Additional information about the current status of this instance, if available.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// Optional. reasons that causes instance in "SUSPENDED" state.
        /// </summary>
        public readonly ImmutableArray<string> SuspensionReasons;
        /// <summary>
        /// The service tier of the instance.
        /// </summary>
        public readonly string Tier;
        /// <summary>
        /// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
        /// </summary>
        public readonly string TransitEncryptionMode;

        [OutputConstructor]
        private GetInstanceResult(
            string alternativeLocationId,

            bool authEnabled,

            string authorizedNetwork,

            ImmutableArray<string> availableMaintenanceVersions,

            string connectMode,

            string createTime,

            string currentLocationId,

            string customerManagedKey,

            string displayName,

            string host,

            ImmutableDictionary<string, string> labels,

            string location,

            Outputs.MaintenancePolicyResponse maintenancePolicy,

            Outputs.MaintenanceScheduleResponse maintenanceSchedule,

            string maintenanceVersion,

            int memorySizeGb,

            string name,

            ImmutableArray<Outputs.NodeInfoResponse> nodes,

            Outputs.PersistenceConfigResponse persistenceConfig,

            string persistenceIamIdentity,

            int port,

            string readEndpoint,

            int readEndpointPort,

            string readReplicasMode,

            ImmutableDictionary<string, string> redisConfigs,

            string redisVersion,

            int replicaCount,

            string reservedIpRange,

            string secondaryIpRange,

            ImmutableArray<Outputs.TlsCertificateResponse> serverCaCerts,

            string state,

            string statusMessage,

            ImmutableArray<string> suspensionReasons,

            string tier,

            string transitEncryptionMode)
        {
            AlternativeLocationId = alternativeLocationId;
            AuthEnabled = authEnabled;
            AuthorizedNetwork = authorizedNetwork;
            AvailableMaintenanceVersions = availableMaintenanceVersions;
            ConnectMode = connectMode;
            CreateTime = createTime;
            CurrentLocationId = currentLocationId;
            CustomerManagedKey = customerManagedKey;
            DisplayName = displayName;
            Host = host;
            Labels = labels;
            Location = location;
            MaintenancePolicy = maintenancePolicy;
            MaintenanceSchedule = maintenanceSchedule;
            MaintenanceVersion = maintenanceVersion;
            MemorySizeGb = memorySizeGb;
            Name = name;
            Nodes = nodes;
            PersistenceConfig = persistenceConfig;
            PersistenceIamIdentity = persistenceIamIdentity;
            Port = port;
            ReadEndpoint = readEndpoint;
            ReadEndpointPort = readEndpointPort;
            ReadReplicasMode = readReplicasMode;
            RedisConfigs = redisConfigs;
            RedisVersion = redisVersion;
            ReplicaCount = replicaCount;
            ReservedIpRange = reservedIpRange;
            SecondaryIpRange = secondaryIpRange;
            ServerCaCerts = serverCaCerts;
            State = state;
            StatusMessage = statusMessage;
            SuspensionReasons = suspensionReasons;
            Tier = tier;
            TransitEncryptionMode = transitEncryptionMode;
        }
    }
}
