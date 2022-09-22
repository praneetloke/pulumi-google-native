// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1
{
    public static class GetInstance
    {
        /// <summary>
        /// Retrieves a resource containing information about a Cloud SQL instance.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:sqladmin/v1:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a resource containing information about a Cloud SQL instance.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("google-native:sqladmin/v1:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("instance", required: true)]
        public string Instance { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

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
        /// List all maintenance versions applicable on the instance
        /// </summary>
        public readonly ImmutableArray<string> AvailableMaintenanceVersions;
        /// <summary>
        /// The backend type. `SECOND_GEN`: Cloud SQL database instance. `EXTERNAL`: A database server that is not managed by Google. This property is read-only; use the `tier` property in the `settings` object to determine the database type.
        /// </summary>
        public readonly string BackendType;
        /// <summary>
        /// Connection name of the Cloud SQL instance used in connection strings.
        /// </summary>
        public readonly string ConnectionName;
        /// <summary>
        /// The time when the instance was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The current disk usage of the instance in bytes. This property has been deprecated. Use the "cloudsql.googleapis.com/database/disk/bytes_used" metric in Cloud Monitoring API instead. Please see [this announcement](https://groups.google.com/d/msg/google-cloud-sql-announce/I_7-F9EBhT0/BtvFtdFeAgAJ) for details.
        /// </summary>
        public readonly string CurrentDiskSize;
        /// <summary>
        /// Stores the current database version running on the instance including minor version such as `MYSQL_8_0_18`.
        /// </summary>
        public readonly string DatabaseInstalledVersion;
        /// <summary>
        /// The database engine type and version. The `databaseVersion` field cannot be changed after instance creation.
        /// </summary>
        public readonly string DatabaseVersion;
        /// <summary>
        /// Disk encryption configuration specific to an instance.
        /// </summary>
        public readonly Outputs.DiskEncryptionConfigurationResponse DiskEncryptionConfiguration;
        /// <summary>
        /// Disk encryption status specific to an instance.
        /// </summary>
        public readonly Outputs.DiskEncryptionStatusResponse DiskEncryptionStatus;
        /// <summary>
        /// This field is deprecated and will be removed from a future version of the API. Use the `settings.settingsVersion` field instead.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The name and status of the failover replica.
        /// </summary>
        public readonly Outputs.InstanceFailoverReplicaResponse FailoverReplica;
        /// <summary>
        /// The Compute Engine zone that the instance is currently serving from. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary zone. WARNING: Changing this might restart the instance.
        /// </summary>
        public readonly string GceZone;
        /// <summary>
        /// The instance type.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The assigned IP addresses for the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpMappingResponse> IpAddresses;
        /// <summary>
        /// The IPv6 address assigned to the instance. (Deprecated) This property was applicable only to First Generation instances.
        /// </summary>
        public readonly string Ipv6Address;
        /// <summary>
        /// This is always `sql#instance`.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The current software version on the instance.
        /// </summary>
        public readonly string MaintenanceVersion;
        /// <summary>
        /// The name of the instance which will act as primary in the replication setup.
        /// </summary>
        public readonly string MasterInstanceName;
        /// <summary>
        /// The maximum disk size of the instance in bytes.
        /// </summary>
        public readonly string MaxDiskSize;
        /// <summary>
        /// Name of the Cloud SQL instance. This does not include the project ID.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configuration specific to on-premises instances.
        /// </summary>
        public readonly Outputs.OnPremisesConfigurationResponse OnPremisesConfiguration;
        /// <summary>
        /// This field represents the report generated by the proactive database wellness job for OutOfDisk issues. * Writers: * the proactive database wellness job for OOD. * Readers: * the proactive database wellness job
        /// </summary>
        public readonly Outputs.SqlOutOfDiskReportResponse OutOfDiskReport;
        /// <summary>
        /// The project ID of the project containing the Cloud SQL instance. The Google apps domain is prefixed if applicable.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The geographical region. Can be: * `us-central` (`FIRST_GEN` instances only) * `us-central1` (`SECOND_GEN` instances only) * `asia-east1` or `europe-west1`. Defaults to `us-central` or `us-central1` depending on the instance type. The region cannot be changed after instance creation.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Configuration specific to failover replicas and read replicas.
        /// </summary>
        public readonly Outputs.ReplicaConfigurationResponse ReplicaConfiguration;
        /// <summary>
        /// The replicas of the instance.
        /// </summary>
        public readonly ImmutableArray<string> ReplicaNames;
        /// <summary>
        /// Initial root password. Use only on creation. You must set root passwords before you can connect to PostgreSQL instances.
        /// </summary>
        public readonly string RootPassword;
        /// <summary>
        /// The status indicating if instance satisfiesPzs. Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// The start time of any upcoming scheduled maintenance for this instance.
        /// </summary>
        public readonly Outputs.SqlScheduledMaintenanceResponse ScheduledMaintenance;
        /// <summary>
        /// The Compute Engine zone that the failover instance is currently serving from for a regional instance. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary/failover zone.
        /// </summary>
        public readonly string SecondaryGceZone;
        /// <summary>
        /// The URI of this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// SSL configuration.
        /// </summary>
        public readonly Outputs.SslCertResponse ServerCaCert;
        /// <summary>
        /// The service account email address assigned to the instance.\This property is read-only.
        /// </summary>
        public readonly string ServiceAccountEmailAddress;
        /// <summary>
        /// The user settings.
        /// </summary>
        public readonly Outputs.SettingsResponse Settings;
        /// <summary>
        /// The current serving state of the Cloud SQL instance.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// If the instance state is SUSPENDED, the reason for the suspension.
        /// </summary>
        public readonly ImmutableArray<string> SuspensionReason;

        [OutputConstructor]
        private GetInstanceResult(
            ImmutableArray<string> availableMaintenanceVersions,

            string backendType,

            string connectionName,

            string createTime,

            string currentDiskSize,

            string databaseInstalledVersion,

            string databaseVersion,

            Outputs.DiskEncryptionConfigurationResponse diskEncryptionConfiguration,

            Outputs.DiskEncryptionStatusResponse diskEncryptionStatus,

            string etag,

            Outputs.InstanceFailoverReplicaResponse failoverReplica,

            string gceZone,

            string instanceType,

            ImmutableArray<Outputs.IpMappingResponse> ipAddresses,

            string ipv6Address,

            string kind,

            string maintenanceVersion,

            string masterInstanceName,

            string maxDiskSize,

            string name,

            Outputs.OnPremisesConfigurationResponse onPremisesConfiguration,

            Outputs.SqlOutOfDiskReportResponse outOfDiskReport,

            string project,

            string region,

            Outputs.ReplicaConfigurationResponse replicaConfiguration,

            ImmutableArray<string> replicaNames,

            string rootPassword,

            bool satisfiesPzs,

            Outputs.SqlScheduledMaintenanceResponse scheduledMaintenance,

            string secondaryGceZone,

            string selfLink,

            Outputs.SslCertResponse serverCaCert,

            string serviceAccountEmailAddress,

            Outputs.SettingsResponse settings,

            string state,

            ImmutableArray<string> suspensionReason)
        {
            AvailableMaintenanceVersions = availableMaintenanceVersions;
            BackendType = backendType;
            ConnectionName = connectionName;
            CreateTime = createTime;
            CurrentDiskSize = currentDiskSize;
            DatabaseInstalledVersion = databaseInstalledVersion;
            DatabaseVersion = databaseVersion;
            DiskEncryptionConfiguration = diskEncryptionConfiguration;
            DiskEncryptionStatus = diskEncryptionStatus;
            Etag = etag;
            FailoverReplica = failoverReplica;
            GceZone = gceZone;
            InstanceType = instanceType;
            IpAddresses = ipAddresses;
            Ipv6Address = ipv6Address;
            Kind = kind;
            MaintenanceVersion = maintenanceVersion;
            MasterInstanceName = masterInstanceName;
            MaxDiskSize = maxDiskSize;
            Name = name;
            OnPremisesConfiguration = onPremisesConfiguration;
            OutOfDiskReport = outOfDiskReport;
            Project = project;
            Region = region;
            ReplicaConfiguration = replicaConfiguration;
            ReplicaNames = replicaNames;
            RootPassword = rootPassword;
            SatisfiesPzs = satisfiesPzs;
            ScheduledMaintenance = scheduledMaintenance;
            SecondaryGceZone = secondaryGceZone;
            SelfLink = selfLink;
            ServerCaCert = serverCaCert;
            ServiceAccountEmailAddress = serviceAccountEmailAddress;
            Settings = settings;
            State = state;
            SuspensionReason = suspensionReason;
        }
    }
}
