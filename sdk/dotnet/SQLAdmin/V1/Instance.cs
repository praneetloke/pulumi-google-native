// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1
{
    /// <summary>
    /// Creates a new Cloud SQL instance.
    /// </summary>
    [GoogleNativeResourceType("google-native:sqladmin/v1:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List all maintenance versions applicable on the instance
        /// </summary>
        [Output("availableMaintenanceVersions")]
        public Output<ImmutableArray<string>> AvailableMaintenanceVersions { get; private set; } = null!;

        /// <summary>
        /// The backend type. `SECOND_GEN`: Cloud SQL database instance. `EXTERNAL`: A database server that is not managed by Google. This property is read-only; use the `tier` property in the `settings` object to determine the database type.
        /// </summary>
        [Output("backendType")]
        public Output<string> BackendType { get; private set; } = null!;

        /// <summary>
        /// Connection name of the Cloud SQL instance used in connection strings.
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// The time when the instance was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The current disk usage of the instance in bytes. This property has been deprecated. Use the "cloudsql.googleapis.com/database/disk/bytes_used" metric in Cloud Monitoring API instead. Please see [this announcement](https://groups.google.com/d/msg/google-cloud-sql-announce/I_7-F9EBhT0/BtvFtdFeAgAJ) for details.
        /// </summary>
        [Output("currentDiskSize")]
        public Output<string> CurrentDiskSize { get; private set; } = null!;

        /// <summary>
        /// Stores the current database version running on the instance including minor version such as `MYSQL_8_0_18`.
        /// </summary>
        [Output("databaseInstalledVersion")]
        public Output<string> DatabaseInstalledVersion { get; private set; } = null!;

        /// <summary>
        /// The database engine type and version. The `databaseVersion` field cannot be changed after instance creation.
        /// </summary>
        [Output("databaseVersion")]
        public Output<string> DatabaseVersion { get; private set; } = null!;

        /// <summary>
        /// Disk encryption configuration specific to an instance.
        /// </summary>
        [Output("diskEncryptionConfiguration")]
        public Output<Outputs.DiskEncryptionConfigurationResponse> DiskEncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// Disk encryption status specific to an instance.
        /// </summary>
        [Output("diskEncryptionStatus")]
        public Output<Outputs.DiskEncryptionStatusResponse> DiskEncryptionStatus { get; private set; } = null!;

        /// <summary>
        /// This field is deprecated and will be removed from a future version of the API. Use the `settings.settingsVersion` field instead.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name and status of the failover replica.
        /// </summary>
        [Output("failoverReplica")]
        public Output<Outputs.InstanceFailoverReplicaResponse> FailoverReplica { get; private set; } = null!;

        /// <summary>
        /// The Compute Engine zone that the instance is currently serving from. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary zone. WARNING: Changing this might restart the instance.
        /// </summary>
        [Output("gceZone")]
        public Output<string> GceZone { get; private set; } = null!;

        /// <summary>
        /// The instance type.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The assigned IP addresses for the instance.
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<Outputs.IpMappingResponse>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// The IPv6 address assigned to the instance. (Deprecated) This property was applicable only to First Generation instances.
        /// </summary>
        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        /// <summary>
        /// This is always `sql#instance`.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The current software version on the instance.
        /// </summary>
        [Output("maintenanceVersion")]
        public Output<string> MaintenanceVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the instance which will act as primary in the replication setup.
        /// </summary>
        [Output("masterInstanceName")]
        public Output<string> MasterInstanceName { get; private set; } = null!;

        /// <summary>
        /// The maximum disk size of the instance in bytes.
        /// </summary>
        [Output("maxDiskSize")]
        public Output<string> MaxDiskSize { get; private set; } = null!;

        /// <summary>
        /// Name of the Cloud SQL instance. This does not include the project ID.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configuration specific to on-premises instances.
        /// </summary>
        [Output("onPremisesConfiguration")]
        public Output<Outputs.OnPremisesConfigurationResponse> OnPremisesConfiguration { get; private set; } = null!;

        /// <summary>
        /// This field represents the report generated by the proactive database wellness job for OutOfDisk issues. * Writers: * the proactive database wellness job for OOD. * Readers: * the proactive database wellness job
        /// </summary>
        [Output("outOfDiskReport")]
        public Output<Outputs.SqlOutOfDiskReportResponse> OutOfDiskReport { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The geographical region. Can be: * `us-central` (`FIRST_GEN` instances only) * `us-central1` (`SECOND_GEN` instances only) * `asia-east1` or `europe-west1`. Defaults to `us-central` or `us-central1` depending on the instance type. The region cannot be changed after instance creation.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Configuration specific to failover replicas and read replicas.
        /// </summary>
        [Output("replicaConfiguration")]
        public Output<Outputs.ReplicaConfigurationResponse> ReplicaConfiguration { get; private set; } = null!;

        /// <summary>
        /// The replicas of the instance.
        /// </summary>
        [Output("replicaNames")]
        public Output<ImmutableArray<string>> ReplicaNames { get; private set; } = null!;

        /// <summary>
        /// Initial root password. Use only on creation. You must set root passwords before you can connect to PostgreSQL instances.
        /// </summary>
        [Output("rootPassword")]
        public Output<string> RootPassword { get; private set; } = null!;

        /// <summary>
        /// The status indicating if instance satisfiesPzs. Reserved for future use.
        /// </summary>
        [Output("satisfiesPzs")]
        public Output<bool> SatisfiesPzs { get; private set; } = null!;

        /// <summary>
        /// The start time of any upcoming scheduled maintenance for this instance.
        /// </summary>
        [Output("scheduledMaintenance")]
        public Output<Outputs.SqlScheduledMaintenanceResponse> ScheduledMaintenance { get; private set; } = null!;

        /// <summary>
        /// The Compute Engine zone that the failover instance is currently serving from for a regional instance. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary/failover zone.
        /// </summary>
        [Output("secondaryGceZone")]
        public Output<string> SecondaryGceZone { get; private set; } = null!;

        /// <summary>
        /// The URI of this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// SSL configuration.
        /// </summary>
        [Output("serverCaCert")]
        public Output<Outputs.SslCertResponse> ServerCaCert { get; private set; } = null!;

        /// <summary>
        /// The service account email address assigned to the instance.\This property is read-only.
        /// </summary>
        [Output("serviceAccountEmailAddress")]
        public Output<string> ServiceAccountEmailAddress { get; private set; } = null!;

        /// <summary>
        /// The user settings.
        /// </summary>
        [Output("settings")]
        public Output<Outputs.SettingsResponse> Settings { get; private set; } = null!;

        /// <summary>
        /// The current serving state of the Cloud SQL instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// If the instance state is SUSPENDED, the reason for the suspension.
        /// </summary>
        [Output("suspensionReason")]
        public Output<ImmutableArray<string>> SuspensionReason { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:sqladmin/v1:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:sqladmin/v1:Instance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
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

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backend type. `SECOND_GEN`: Cloud SQL database instance. `EXTERNAL`: A database server that is not managed by Google. This property is read-only; use the `tier` property in the `settings` object to determine the database type.
        /// </summary>
        [Input("backendType")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.InstanceBackendType>? BackendType { get; set; }

        /// <summary>
        /// Connection name of the Cloud SQL instance used in connection strings.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// The current disk usage of the instance in bytes. This property has been deprecated. Use the "cloudsql.googleapis.com/database/disk/bytes_used" metric in Cloud Monitoring API instead. Please see [this announcement](https://groups.google.com/d/msg/google-cloud-sql-announce/I_7-F9EBhT0/BtvFtdFeAgAJ) for details.
        /// </summary>
        [Input("currentDiskSize")]
        public Input<string>? CurrentDiskSize { get; set; }

        /// <summary>
        /// The database engine type and version. The `databaseVersion` field cannot be changed after instance creation.
        /// </summary>
        [Input("databaseVersion")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.InstanceDatabaseVersion>? DatabaseVersion { get; set; }

        /// <summary>
        /// Disk encryption configuration specific to an instance.
        /// </summary>
        [Input("diskEncryptionConfiguration")]
        public Input<Inputs.DiskEncryptionConfigurationArgs>? DiskEncryptionConfiguration { get; set; }

        /// <summary>
        /// Disk encryption status specific to an instance.
        /// </summary>
        [Input("diskEncryptionStatus")]
        public Input<Inputs.DiskEncryptionStatusArgs>? DiskEncryptionStatus { get; set; }

        /// <summary>
        /// This field is deprecated and will be removed from a future version of the API. Use the `settings.settingsVersion` field instead.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name and status of the failover replica.
        /// </summary>
        [Input("failoverReplica")]
        public Input<Inputs.InstanceFailoverReplicaArgs>? FailoverReplica { get; set; }

        /// <summary>
        /// The Compute Engine zone that the instance is currently serving from. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary zone. WARNING: Changing this might restart the instance.
        /// </summary>
        [Input("gceZone")]
        public Input<string>? GceZone { get; set; }

        /// <summary>
        /// The instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.InstanceInstanceType>? InstanceType { get; set; }

        [Input("ipAddresses")]
        private InputList<Inputs.IpMappingArgs>? _ipAddresses;

        /// <summary>
        /// The assigned IP addresses for the instance.
        /// </summary>
        public InputList<Inputs.IpMappingArgs> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<Inputs.IpMappingArgs>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// The IPv6 address assigned to the instance. (Deprecated) This property was applicable only to First Generation instances.
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        /// <summary>
        /// This is always `sql#instance`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The current software version on the instance.
        /// </summary>
        [Input("maintenanceVersion")]
        public Input<string>? MaintenanceVersion { get; set; }

        /// <summary>
        /// The name of the instance which will act as primary in the replication setup.
        /// </summary>
        [Input("masterInstanceName")]
        public Input<string>? MasterInstanceName { get; set; }

        /// <summary>
        /// The maximum disk size of the instance in bytes.
        /// </summary>
        [Input("maxDiskSize")]
        public Input<string>? MaxDiskSize { get; set; }

        /// <summary>
        /// Name of the Cloud SQL instance. This does not include the project ID.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration specific to on-premises instances.
        /// </summary>
        [Input("onPremisesConfiguration")]
        public Input<Inputs.OnPremisesConfigurationArgs>? OnPremisesConfiguration { get; set; }

        /// <summary>
        /// This field represents the report generated by the proactive database wellness job for OutOfDisk issues. * Writers: * the proactive database wellness job for OOD. * Readers: * the proactive database wellness job
        /// </summary>
        [Input("outOfDiskReport")]
        public Input<Inputs.SqlOutOfDiskReportArgs>? OutOfDiskReport { get; set; }

        /// <summary>
        /// The project ID of the project containing the Cloud SQL instance. The Google apps domain is prefixed if applicable.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The geographical region. Can be: * `us-central` (`FIRST_GEN` instances only) * `us-central1` (`SECOND_GEN` instances only) * `asia-east1` or `europe-west1`. Defaults to `us-central` or `us-central1` depending on the instance type. The region cannot be changed after instance creation.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Configuration specific to failover replicas and read replicas.
        /// </summary>
        [Input("replicaConfiguration")]
        public Input<Inputs.ReplicaConfigurationArgs>? ReplicaConfiguration { get; set; }

        [Input("replicaNames")]
        private InputList<string>? _replicaNames;

        /// <summary>
        /// The replicas of the instance.
        /// </summary>
        public InputList<string> ReplicaNames
        {
            get => _replicaNames ?? (_replicaNames = new InputList<string>());
            set => _replicaNames = value;
        }

        /// <summary>
        /// Initial root password. Use only on creation. You must set root passwords before you can connect to PostgreSQL instances.
        /// </summary>
        [Input("rootPassword")]
        public Input<string>? RootPassword { get; set; }

        /// <summary>
        /// The status indicating if instance satisfiesPzs. Reserved for future use.
        /// </summary>
        [Input("satisfiesPzs")]
        public Input<bool>? SatisfiesPzs { get; set; }

        /// <summary>
        /// The start time of any upcoming scheduled maintenance for this instance.
        /// </summary>
        [Input("scheduledMaintenance")]
        public Input<Inputs.SqlScheduledMaintenanceArgs>? ScheduledMaintenance { get; set; }

        /// <summary>
        /// The Compute Engine zone that the failover instance is currently serving from for a regional instance. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary/failover zone.
        /// </summary>
        [Input("secondaryGceZone")]
        public Input<string>? SecondaryGceZone { get; set; }

        /// <summary>
        /// The URI of this resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// SSL configuration.
        /// </summary>
        [Input("serverCaCert")]
        public Input<Inputs.SslCertArgs>? ServerCaCert { get; set; }

        /// <summary>
        /// The service account email address assigned to the instance.\This property is read-only.
        /// </summary>
        [Input("serviceAccountEmailAddress")]
        public Input<string>? ServiceAccountEmailAddress { get; set; }

        /// <summary>
        /// The user settings.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.SettingsArgs>? Settings { get; set; }

        /// <summary>
        /// The current serving state of the Cloud SQL instance.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.InstanceState>? State { get; set; }

        [Input("suspensionReason")]
        private InputList<Pulumi.GoogleNative.SQLAdmin.V1.InstanceSuspensionReasonItem>? _suspensionReason;

        /// <summary>
        /// If the instance state is SUSPENDED, the reason for the suspension.
        /// </summary>
        public InputList<Pulumi.GoogleNative.SQLAdmin.V1.InstanceSuspensionReasonItem> SuspensionReason
        {
            get => _suspensionReason ?? (_suspensionReason = new InputList<Pulumi.GoogleNative.SQLAdmin.V1.InstanceSuspensionReasonItem>());
            set => _suspensionReason = value;
        }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }
}
