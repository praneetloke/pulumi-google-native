// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Dataproc.V1
{
    /// <summary>
    /// Optional. The type of IPv6 access for a cluster.
    /// </summary>
    [EnumType]
    public readonly struct GceClusterConfigPrivateIpv6GoogleAccess : IEquatable<GceClusterConfigPrivateIpv6GoogleAccess>
    {
        private readonly string _value;

        private GceClusterConfigPrivateIpv6GoogleAccess(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// If unspecified, Compute Engine default behavior will apply, which is the same as INHERIT_FROM_SUBNETWORK.
        /// </summary>
        public static GceClusterConfigPrivateIpv6GoogleAccess PrivateIpv6GoogleAccessUnspecified { get; } = new GceClusterConfigPrivateIpv6GoogleAccess("PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED");
        /// <summary>
        /// Private access to and from Google Services configuration inherited from the subnetwork configuration. This is the default Compute Engine behavior.
        /// </summary>
        public static GceClusterConfigPrivateIpv6GoogleAccess InheritFromSubnetwork { get; } = new GceClusterConfigPrivateIpv6GoogleAccess("INHERIT_FROM_SUBNETWORK");
        /// <summary>
        /// Enables outbound private IPv6 access to Google Services from the Dataproc cluster.
        /// </summary>
        public static GceClusterConfigPrivateIpv6GoogleAccess Outbound { get; } = new GceClusterConfigPrivateIpv6GoogleAccess("OUTBOUND");
        /// <summary>
        /// Enables bidirectional private IPv6 access between Google Services and the Dataproc cluster.
        /// </summary>
        public static GceClusterConfigPrivateIpv6GoogleAccess Bidirectional { get; } = new GceClusterConfigPrivateIpv6GoogleAccess("BIDIRECTIONAL");

        public static bool operator ==(GceClusterConfigPrivateIpv6GoogleAccess left, GceClusterConfigPrivateIpv6GoogleAccess right) => left.Equals(right);
        public static bool operator !=(GceClusterConfigPrivateIpv6GoogleAccess left, GceClusterConfigPrivateIpv6GoogleAccess right) => !left.Equals(right);

        public static explicit operator string(GceClusterConfigPrivateIpv6GoogleAccess value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GceClusterConfigPrivateIpv6GoogleAccess other && Equals(other);
        public bool Equals(GceClusterConfigPrivateIpv6GoogleAccess other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct GkeNodePoolTargetRolesItem : IEquatable<GkeNodePoolTargetRolesItem>
    {
        private readonly string _value;

        private GkeNodePoolTargetRolesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Role is unspecified.
        /// </summary>
        public static GkeNodePoolTargetRolesItem RoleUnspecified { get; } = new GkeNodePoolTargetRolesItem("ROLE_UNSPECIFIED");
        /// <summary>
        /// At least one node pool must have the DEFAULT role. Work assigned to a role that is not associated with a node pool is assigned to the node pool with the DEFAULT role. For example, work assigned to the CONTROLLER role will be assigned to the node pool with the DEFAULT role if no node pool has the CONTROLLER role.
        /// </summary>
        public static GkeNodePoolTargetRolesItem Default { get; } = new GkeNodePoolTargetRolesItem("DEFAULT");
        /// <summary>
        /// Run work associated with the Dataproc control plane (for example, controllers and webhooks). Very low resource requirements.
        /// </summary>
        public static GkeNodePoolTargetRolesItem Controller { get; } = new GkeNodePoolTargetRolesItem("CONTROLLER");
        /// <summary>
        /// Run work associated with a Spark driver of a job.
        /// </summary>
        public static GkeNodePoolTargetRolesItem SparkDriver { get; } = new GkeNodePoolTargetRolesItem("SPARK_DRIVER");
        /// <summary>
        /// Run work associated with a Spark executor of a job.
        /// </summary>
        public static GkeNodePoolTargetRolesItem SparkExecutor { get; } = new GkeNodePoolTargetRolesItem("SPARK_EXECUTOR");

        public static bool operator ==(GkeNodePoolTargetRolesItem left, GkeNodePoolTargetRolesItem right) => left.Equals(right);
        public static bool operator !=(GkeNodePoolTargetRolesItem left, GkeNodePoolTargetRolesItem right) => !left.Equals(right);

        public static explicit operator string(GkeNodePoolTargetRolesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GkeNodePoolTargetRolesItem other && Equals(other);
        public bool Equals(GkeNodePoolTargetRolesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional. Specifies the preemptibility of the instance group.The default value for master and worker groups is NON_PREEMPTIBLE. This default cannot be changed.The default value for secondary instances is PREEMPTIBLE.
    /// </summary>
    [EnumType]
    public readonly struct InstanceGroupConfigPreemptibility : IEquatable<InstanceGroupConfigPreemptibility>
    {
        private readonly string _value;

        private InstanceGroupConfigPreemptibility(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Preemptibility is unspecified, the system will choose the appropriate setting for each instance group.
        /// </summary>
        public static InstanceGroupConfigPreemptibility PreemptibilityUnspecified { get; } = new InstanceGroupConfigPreemptibility("PREEMPTIBILITY_UNSPECIFIED");
        /// <summary>
        /// Instances are non-preemptible.This option is allowed for all instance groups and is the only valid value for Master and Worker instance groups.
        /// </summary>
        public static InstanceGroupConfigPreemptibility NonPreemptible { get; } = new InstanceGroupConfigPreemptibility("NON_PREEMPTIBLE");
        /// <summary>
        /// Instances are preemptible (https://cloud.google.com/compute/docs/instances/preemptible).This option is allowed only for secondary worker (https://cloud.google.com/dataproc/docs/concepts/compute/secondary-vms) groups.
        /// </summary>
        public static InstanceGroupConfigPreemptibility Preemptible { get; } = new InstanceGroupConfigPreemptibility("PREEMPTIBLE");
        /// <summary>
        /// Instances are Spot VMs (https://cloud.google.com/compute/docs/instances/spot).This option is allowed only for secondary worker (https://cloud.google.com/dataproc/docs/concepts/compute/secondary-vms) groups. Spot VMs are the latest version of preemptible VMs (https://cloud.google.com/compute/docs/instances/preemptible), and provide additional features.
        /// </summary>
        public static InstanceGroupConfigPreemptibility Spot { get; } = new InstanceGroupConfigPreemptibility("SPOT");

        public static bool operator ==(InstanceGroupConfigPreemptibility left, InstanceGroupConfigPreemptibility right) => left.Equals(right);
        public static bool operator !=(InstanceGroupConfigPreemptibility left, InstanceGroupConfigPreemptibility right) => !left.Equals(right);

        public static explicit operator string(InstanceGroupConfigPreemptibility value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InstanceGroupConfigPreemptibility other && Equals(other);
        public bool Equals(InstanceGroupConfigPreemptibility other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. A standard set of metrics is collected unless metricOverrides are specified for the metric source (see Custom metrics (https://cloud.google.com/dataproc/docs/guides/dataproc-metrics#custom_metrics) for more information).
    /// </summary>
    [EnumType]
    public readonly struct MetricMetricSource : IEquatable<MetricMetricSource>
    {
        private readonly string _value;

        private MetricMetricSource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Required unspecified metric source.
        /// </summary>
        public static MetricMetricSource MetricSourceUnspecified { get; } = new MetricMetricSource("METRIC_SOURCE_UNSPECIFIED");
        /// <summary>
        /// Monitoring agent metrics. If this source is enabled, Dataproc enables the monitoring agent in Compute Engine, and collects monitoring agent metrics, which are published with an agent.googleapis.com prefix.
        /// </summary>
        public static MetricMetricSource MonitoringAgentDefaults { get; } = new MetricMetricSource("MONITORING_AGENT_DEFAULTS");
        /// <summary>
        /// HDFS metric source.
        /// </summary>
        public static MetricMetricSource Hdfs { get; } = new MetricMetricSource("HDFS");
        /// <summary>
        /// Spark metric source.
        /// </summary>
        public static MetricMetricSource Spark { get; } = new MetricMetricSource("SPARK");
        /// <summary>
        /// YARN metric source.
        /// </summary>
        public static MetricMetricSource Yarn { get; } = new MetricMetricSource("YARN");
        /// <summary>
        /// Spark History Server metric source.
        /// </summary>
        public static MetricMetricSource SparkHistoryServer { get; } = new MetricMetricSource("SPARK_HISTORY_SERVER");
        /// <summary>
        /// Hiveserver2 metric source.
        /// </summary>
        public static MetricMetricSource Hiveserver2 { get; } = new MetricMetricSource("HIVESERVER2");
        /// <summary>
        /// hivemetastore metric source
        /// </summary>
        public static MetricMetricSource Hivemetastore { get; } = new MetricMetricSource("HIVEMETASTORE");

        public static bool operator ==(MetricMetricSource left, MetricMetricSource right) => left.Equals(right);
        public static bool operator !=(MetricMetricSource left, MetricMetricSource right) => !left.Equals(right);

        public static explicit operator string(MetricMetricSource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MetricMetricSource other && Equals(other);
        public bool Equals(MetricMetricSource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NodeGroupRolesItem : IEquatable<NodeGroupRolesItem>
    {
        private readonly string _value;

        private NodeGroupRolesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Required unspecified role.
        /// </summary>
        public static NodeGroupRolesItem RoleUnspecified { get; } = new NodeGroupRolesItem("ROLE_UNSPECIFIED");
        /// <summary>
        /// Job drivers run on the node pool.
        /// </summary>
        public static NodeGroupRolesItem Driver { get; } = new NodeGroupRolesItem("DRIVER");

        public static bool operator ==(NodeGroupRolesItem left, NodeGroupRolesItem right) => left.Equals(right);
        public static bool operator !=(NodeGroupRolesItem left, NodeGroupRolesItem right) => !left.Equals(right);

        public static explicit operator string(NodeGroupRolesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NodeGroupRolesItem other && Equals(other);
        public bool Equals(NodeGroupRolesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional. Type of reservation to consume
    /// </summary>
    [EnumType]
    public readonly struct ReservationAffinityConsumeReservationType : IEquatable<ReservationAffinityConsumeReservationType>
    {
        private readonly string _value;

        private ReservationAffinityConsumeReservationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ReservationAffinityConsumeReservationType TypeUnspecified { get; } = new ReservationAffinityConsumeReservationType("TYPE_UNSPECIFIED");
        /// <summary>
        /// Do not consume from any allocated capacity.
        /// </summary>
        public static ReservationAffinityConsumeReservationType NoReservation { get; } = new ReservationAffinityConsumeReservationType("NO_RESERVATION");
        /// <summary>
        /// Consume any reservation available.
        /// </summary>
        public static ReservationAffinityConsumeReservationType AnyReservation { get; } = new ReservationAffinityConsumeReservationType("ANY_RESERVATION");
        /// <summary>
        /// Must consume from a specific reservation. Must specify key value fields for specifying the reservations.
        /// </summary>
        public static ReservationAffinityConsumeReservationType SpecificReservation { get; } = new ReservationAffinityConsumeReservationType("SPECIFIC_RESERVATION");

        public static bool operator ==(ReservationAffinityConsumeReservationType left, ReservationAffinityConsumeReservationType right) => left.Equals(right);
        public static bool operator !=(ReservationAffinityConsumeReservationType left, ReservationAffinityConsumeReservationType right) => !left.Equals(right);

        public static explicit operator string(ReservationAffinityConsumeReservationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ReservationAffinityConsumeReservationType other && Equals(other);
        public bool Equals(ReservationAffinityConsumeReservationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct SoftwareConfigOptionalComponentsItem : IEquatable<SoftwareConfigOptionalComponentsItem>
    {
        private readonly string _value;

        private SoftwareConfigOptionalComponentsItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified component. Specifying this will cause Cluster creation to fail.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem ComponentUnspecified { get; } = new SoftwareConfigOptionalComponentsItem("COMPONENT_UNSPECIFIED");
        /// <summary>
        /// The Anaconda python distribution. The Anaconda component is not supported in the Dataproc 2.0 image. The 2.0 image is pre-installed with Miniconda.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Anaconda { get; } = new SoftwareConfigOptionalComponentsItem("ANACONDA");
        /// <summary>
        /// Docker
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Docker { get; } = new SoftwareConfigOptionalComponentsItem("DOCKER");
        /// <summary>
        /// The Druid query engine. (alpha)
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Druid { get; } = new SoftwareConfigOptionalComponentsItem("DRUID");
        /// <summary>
        /// Flink
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Flink { get; } = new SoftwareConfigOptionalComponentsItem("FLINK");
        /// <summary>
        /// HBase. (beta)
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Hbase { get; } = new SoftwareConfigOptionalComponentsItem("HBASE");
        /// <summary>
        /// The Hive Web HCatalog (the REST service for accessing HCatalog).
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem HiveWebhcat { get; } = new SoftwareConfigOptionalComponentsItem("HIVE_WEBHCAT");
        /// <summary>
        /// Hudi.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Hudi { get; } = new SoftwareConfigOptionalComponentsItem("HUDI");
        /// <summary>
        /// The Jupyter Notebook.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Jupyter { get; } = new SoftwareConfigOptionalComponentsItem("JUPYTER");
        /// <summary>
        /// The Presto query engine.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Presto { get; } = new SoftwareConfigOptionalComponentsItem("PRESTO");
        /// <summary>
        /// The Trino query engine.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Trino { get; } = new SoftwareConfigOptionalComponentsItem("TRINO");
        /// <summary>
        /// The Ranger service.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Ranger { get; } = new SoftwareConfigOptionalComponentsItem("RANGER");
        /// <summary>
        /// The Solr service.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Solr { get; } = new SoftwareConfigOptionalComponentsItem("SOLR");
        /// <summary>
        /// The Zeppelin notebook.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Zeppelin { get; } = new SoftwareConfigOptionalComponentsItem("ZEPPELIN");
        /// <summary>
        /// The Zookeeper service.
        /// </summary>
        public static SoftwareConfigOptionalComponentsItem Zookeeper { get; } = new SoftwareConfigOptionalComponentsItem("ZOOKEEPER");

        public static bool operator ==(SoftwareConfigOptionalComponentsItem left, SoftwareConfigOptionalComponentsItem right) => left.Equals(right);
        public static bool operator !=(SoftwareConfigOptionalComponentsItem left, SoftwareConfigOptionalComponentsItem right) => !left.Equals(right);

        public static explicit operator string(SoftwareConfigOptionalComponentsItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SoftwareConfigOptionalComponentsItem other && Equals(other);
        public bool Equals(SoftwareConfigOptionalComponentsItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
