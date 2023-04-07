// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Beta.Outputs
{

    /// <summary>
    /// A managed metastore service that serves metadata queries.
    /// </summary>
    [OutputType]
    public sealed class ServiceResponse
    {
        /// <summary>
        /// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
        /// </summary>
        public readonly string ArtifactGcsUri;
        /// <summary>
        /// The time when the metastore service was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Immutable. The database type that the Metastore service stores its data.
        /// </summary>
        public readonly string DatabaseType;
        /// <summary>
        /// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
        /// </summary>
        public readonly Outputs.EncryptionConfigResponse EncryptionConfig;
        /// <summary>
        /// The URI of the endpoint used to access the metastore service.
        /// </summary>
        public readonly string EndpointUri;
        /// <summary>
        /// Configuration information specific to running Hive metastore software as the metastore service.
        /// </summary>
        public readonly Outputs.HiveMetastoreConfigResponse HiveMetastoreConfig;
        /// <summary>
        /// User-defined labels for the metastore service.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
        /// </summary>
        public readonly Outputs.MaintenanceWindowResponse MaintenanceWindow;
        /// <summary>
        /// The setting that defines how metastore metadata should be integrated with external services and systems.
        /// </summary>
        public readonly Outputs.MetadataIntegrationResponse MetadataIntegration;
        /// <summary>
        /// The metadata management activities of the metastore service.
        /// </summary>
        public readonly Outputs.MetadataManagementActivityResponse MetadataManagementActivity;
        /// <summary>
        /// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// The configuration specifying the network settings for the Dataproc Metastore service.
        /// </summary>
        public readonly Outputs.NetworkConfigResponse NetworkConfig;
        /// <summary>
        /// The TCP port at which the metastore service is reached. Default: 9083.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
        /// </summary>
        public readonly string ReleaseChannel;
        /// <summary>
        /// Scaling configuration of the metastore service.
        /// </summary>
        public readonly Outputs.ScalingConfigResponse ScalingConfig;
        /// <summary>
        /// The current state of the metastore service.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Additional information about the current state of the metastore service, if available.
        /// </summary>
        public readonly string StateMessage;
        /// <summary>
        /// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
        /// </summary>
        public readonly Outputs.TelemetryConfigResponse TelemetryConfig;
        /// <summary>
        /// The tier of the service.
        /// </summary>
        public readonly string Tier;
        /// <summary>
        /// The globally unique resource identifier of the metastore service.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The time when the metastore service was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private ServiceResponse(
            string artifactGcsUri,

            string createTime,

            string databaseType,

            Outputs.EncryptionConfigResponse encryptionConfig,

            string endpointUri,

            Outputs.HiveMetastoreConfigResponse hiveMetastoreConfig,

            ImmutableDictionary<string, string> labels,

            Outputs.MaintenanceWindowResponse maintenanceWindow,

            Outputs.MetadataIntegrationResponse metadataIntegration,

            Outputs.MetadataManagementActivityResponse metadataManagementActivity,

            string name,

            string network,

            Outputs.NetworkConfigResponse networkConfig,

            int port,

            string releaseChannel,

            Outputs.ScalingConfigResponse scalingConfig,

            string state,

            string stateMessage,

            Outputs.TelemetryConfigResponse telemetryConfig,

            string tier,

            string uid,

            string updateTime)
        {
            ArtifactGcsUri = artifactGcsUri;
            CreateTime = createTime;
            DatabaseType = databaseType;
            EncryptionConfig = encryptionConfig;
            EndpointUri = endpointUri;
            HiveMetastoreConfig = hiveMetastoreConfig;
            Labels = labels;
            MaintenanceWindow = maintenanceWindow;
            MetadataIntegration = metadataIntegration;
            MetadataManagementActivity = metadataManagementActivity;
            Name = name;
            Network = network;
            NetworkConfig = networkConfig;
            Port = port;
            ReleaseChannel = releaseChannel;
            ScalingConfig = scalingConfig;
            State = state;
            StateMessage = stateMessage;
            TelemetryConfig = telemetryConfig;
            Tier = tier;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
