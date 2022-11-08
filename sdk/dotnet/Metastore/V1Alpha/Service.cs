// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Alpha
{
    /// <summary>
    /// Creates a metastore service in a project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:metastore/v1alpha:Service")]
    public partial class Service : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
        /// </summary>
        [Output("artifactGcsUri")]
        public Output<string> ArtifactGcsUri { get; private set; } = null!;

        /// <summary>
        /// The time when the metastore service was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Immutable. The database type that the Metastore service stores its data.
        /// </summary>
        [Output("databaseType")]
        public Output<string> DatabaseType { get; private set; } = null!;

        /// <summary>
        /// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
        /// </summary>
        [Output("encryptionConfig")]
        public Output<Outputs.EncryptionConfigResponse> EncryptionConfig { get; private set; } = null!;

        /// <summary>
        /// The URI of the endpoint used to access the metastore service.
        /// </summary>
        [Output("endpointUri")]
        public Output<string> EndpointUri { get; private set; } = null!;

        /// <summary>
        /// Configuration information specific to running Hive metastore software as the metastore service.
        /// </summary>
        [Output("hiveMetastoreConfig")]
        public Output<Outputs.HiveMetastoreConfigResponse> HiveMetastoreConfig { get; private set; } = null!;

        /// <summary>
        /// User-defined labels for the metastore service.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.MaintenanceWindowResponse> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The setting that defines how metastore metadata should be integrated with external services and systems.
        /// </summary>
        [Output("metadataIntegration")]
        public Output<Outputs.MetadataIntegrationResponse> MetadataIntegration { get; private set; } = null!;

        /// <summary>
        /// The metadata management activities of the metastore service.
        /// </summary>
        [Output("metadataManagementActivity")]
        public Output<Outputs.MetadataManagementActivityResponse> MetadataManagementActivity { get; private set; } = null!;

        /// <summary>
        /// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The configuration specifying the network settings for the Dataproc Metastore service.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.NetworkConfigResponse> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// The TCP port at which the metastore service is reached. Default: 9083.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
        /// </summary>
        [Output("releaseChannel")]
        public Output<string> ReleaseChannel { get; private set; } = null!;

        /// <summary>
        /// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The current state of the metastore service.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Additional information about the current state of the metastore service, if available.
        /// </summary>
        [Output("stateMessage")]
        public Output<string> StateMessage { get; private set; } = null!;

        /// <summary>
        /// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
        /// </summary>
        [Output("telemetryConfig")]
        public Output<Outputs.TelemetryConfigResponse> TelemetryConfig { get; private set; } = null!;

        /// <summary>
        /// The tier of the service.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;

        /// <summary>
        /// The globally unique resource identifier of the metastore service.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time when the metastore service was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("google-native:metastore/v1alpha:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:metastore/v1alpha:Service", name, null, MakeResourceOptions(options, id))
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
                    "serviceId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. The database type that the Metastore service stores its data.
        /// </summary>
        [Input("databaseType")]
        public Input<Pulumi.GoogleNative.Metastore.V1Alpha.ServiceDatabaseType>? DatabaseType { get; set; }

        /// <summary>
        /// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.EncryptionConfigArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// Configuration information specific to running Hive metastore software as the metastore service.
        /// </summary>
        [Input("hiveMetastoreConfig")]
        public Input<Inputs.HiveMetastoreConfigArgs>? HiveMetastoreConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the metastore service.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The setting that defines how metastore metadata should be integrated with external services and systems.
        /// </summary>
        [Input("metadataIntegration")]
        public Input<Inputs.MetadataIntegrationArgs>? MetadataIntegration { get; set; }

        /// <summary>
        /// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The configuration specifying the network settings for the Dataproc Metastore service.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.NetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// The TCP port at which the metastore service is reached. Default: 9083.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
        /// </summary>
        [Input("releaseChannel")]
        public Input<Pulumi.GoogleNative.Metastore.V1Alpha.ServiceReleaseChannel>? ReleaseChannel { get; set; }

        /// <summary>
        /// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
        /// </summary>
        [Input("telemetryConfig")]
        public Input<Inputs.TelemetryConfigArgs>? TelemetryConfig { get; set; }

        /// <summary>
        /// The tier of the service.
        /// </summary>
        [Input("tier")]
        public Input<Pulumi.GoogleNative.Metastore.V1Alpha.ServiceTier>? Tier { get; set; }

        public ServiceArgs()
        {
        }
        public static new ServiceArgs Empty => new ServiceArgs();
    }
}
