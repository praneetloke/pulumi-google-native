// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Redis.V1
{
    /// <summary>
    /// Creates a Redis instance based on the specified tier and memory size. By default, the instance is accessible from the project's [default network](https://cloud.google.com/vpc/docs/vpc). The creation is executed asynchronously and callers may check the returned operation to track its progress. Once the operation is completed the Redis instance will be fully functional. Completed longrunning.Operation will contain the new instance object in the response field. The returned operation is automatically deleted after a few hours, so there is no need to call DeleteOperation.
    /// </summary>
    [GoogleNativeResourceType("google-native:redis/v1:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
        /// </summary>
        [Output("alternativeLocationId")]
        public Output<string> AlternativeLocationId { get; private set; } = null!;

        /// <summary>
        /// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
        /// </summary>
        [Output("authEnabled")]
        public Output<bool> AuthEnabled { get; private set; } = null!;

        /// <summary>
        /// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
        /// </summary>
        [Output("authorizedNetwork")]
        public Output<string> AuthorizedNetwork { get; private set; } = null!;

        /// <summary>
        /// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
        /// </summary>
        [Output("connectMode")]
        public Output<string> ConnectMode { get; private set; } = null!;

        /// <summary>
        /// The time the instance was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the location_id provided by the user at creation time. For Standard Tier instances, this can be either location_id or alternative_location_id and can change after a failover event.
        /// </summary>
        [Output("currentLocationId")]
        public Output<string> CurrentLocationId { get; private set; } = null!;

        /// <summary>
        /// An arbitrary and optional user-provided name for the instance.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Resource labels to represent user provided metadata
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Required. Redis memory size in GiB.
        /// </summary>
        [Output("memorySizeGb")]
        public Output<int> MemorySizeGb { get; private set; } = null!;

        /// <summary>
        /// Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
        /// </summary>
        [Output("persistenceIamIdentity")]
        public Output<string> PersistenceIamIdentity { get; private set; } = null!;

        /// <summary>
        /// The port number of the exposed Redis endpoint.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
        /// </summary>
        [Output("redisConfigs")]
        public Output<ImmutableDictionary<string, string>> RedisConfigs { get; private set; } = null!;

        /// <summary>
        /// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
        /// </summary>
        [Output("redisVersion")]
        public Output<string> RedisVersion { get; private set; } = null!;

        /// <summary>
        /// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
        /// </summary>
        [Output("reservedIpRange")]
        public Output<string> ReservedIpRange { get; private set; } = null!;

        /// <summary>
        /// List of server CA certificates for the instance.
        /// </summary>
        [Output("serverCaCerts")]
        public Output<ImmutableArray<Outputs.TlsCertificateResponse>> ServerCaCerts { get; private set; } = null!;

        /// <summary>
        /// The current state of this instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Additional information about the current status of this instance, if available.
        /// </summary>
        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;

        /// <summary>
        /// Required. The service tier of the instance.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;

        /// <summary>
        /// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
        /// </summary>
        [Output("transitEncryptionMode")]
        public Output<string> TransitEncryptionMode { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("google-native:redis/v1:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:redis/v1:Instance", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
        /// </summary>
        [Input("alternativeLocationId")]
        public Input<string>? AlternativeLocationId { get; set; }

        /// <summary>
        /// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
        /// </summary>
        [Input("authEnabled")]
        public Input<bool>? AuthEnabled { get; set; }

        /// <summary>
        /// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
        /// </summary>
        [Input("connectMode")]
        public Input<string>? ConnectMode { get; set; }

        /// <summary>
        /// An arbitrary and optional user-provided name for the instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels to represent user provided metadata
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Required. Redis memory size in GiB.
        /// </summary>
        [Input("memorySizeGb")]
        public Input<int>? MemorySizeGb { get; set; }

        /// <summary>
        /// Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("redisConfigs")]
        private InputMap<string>? _redisConfigs;

        /// <summary>
        /// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
        /// </summary>
        public InputMap<string> RedisConfigs
        {
            get => _redisConfigs ?? (_redisConfigs = new InputMap<string>());
            set => _redisConfigs = value;
        }

        /// <summary>
        /// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
        /// </summary>
        [Input("redisVersion")]
        public Input<string>? RedisVersion { get; set; }

        /// <summary>
        /// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
        /// </summary>
        [Input("reservedIpRange")]
        public Input<string>? ReservedIpRange { get; set; }

        /// <summary>
        /// Required. The service tier of the instance.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        /// <summary>
        /// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
        /// </summary>
        [Input("transitEncryptionMode")]
        public Input<string>? TransitEncryptionMode { get; set; }

        public InstanceArgs()
        {
        }
    }
}
