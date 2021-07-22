// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1
{
    /// <summary>
    /// Creates a cluster, consisting of the specified number and type of Google Compute Engine instances. By default, the cluster is created in the project's [default network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks). One firewall is added for the cluster. After cluster creation, the Kubelet creates routes for each node to allow the containers on that node to communicate with all other instances in the cluster. Finally, an entry is added to the project's global metadata indicating which CIDR range the cluster is using.
    /// </summary>
    [GoogleNativeResourceType("google-native:container/v1beta1:Cluster")]
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Configurations for the various addons available to run in the cluster.
        /// </summary>
        [Output("addonsConfig")]
        public Output<Outputs.AddonsConfigResponse> AddonsConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration controlling RBAC group membership information.
        /// </summary>
        [Output("authenticatorGroupsConfig")]
        public Output<Outputs.AuthenticatorGroupsConfigResponse> AuthenticatorGroupsConfig { get; private set; } = null!;

        /// <summary>
        /// Autopilot configuration for the cluster.
        /// </summary>
        [Output("autopilot")]
        public Output<Outputs.AutopilotResponse> Autopilot { get; private set; } = null!;

        /// <summary>
        /// Cluster-level autoscaling configuration.
        /// </summary>
        [Output("autoscaling")]
        public Output<Outputs.ClusterAutoscalingResponse> Autoscaling { get; private set; } = null!;

        /// <summary>
        /// Configuration for Binary Authorization.
        /// </summary>
        [Output("binaryAuthorization")]
        public Output<Outputs.BinaryAuthorizationResponse> BinaryAuthorization { get; private set; } = null!;

        /// <summary>
        /// The IP address range of the container pods in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`). Leave blank to have one automatically chosen or specify a `/14` block in `10.0.0.0/8`.
        /// </summary>
        [Output("clusterIpv4Cidr")]
        public Output<string> ClusterIpv4Cidr { get; private set; } = null!;

        /// <summary>
        /// Telemetry integration for the cluster.
        /// </summary>
        [Output("clusterTelemetry")]
        public Output<Outputs.ClusterTelemetryResponse> ClusterTelemetry { get; private set; } = null!;

        /// <summary>
        /// Which conditions caused the current cluster state.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.StatusConditionResponse>> Conditions { get; private set; } = null!;

        /// <summary>
        /// Configuration of Confidential Nodes
        /// </summary>
        [Output("confidentialNodes")]
        public Output<Outputs.ConfidentialNodesResponse> ConfidentialNodes { get; private set; } = null!;

        /// <summary>
        /// [Output only] The time the cluster was created, in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// [Output only] The current software version of the master endpoint.
        /// </summary>
        [Output("currentMasterVersion")]
        public Output<string> CurrentMasterVersion { get; private set; } = null!;

        /// <summary>
        /// [Output only] Deprecated, use [NodePool.version](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters.nodePools) instead. The current version of the node software components. If they are currently at multiple versions because they're in the process of being upgraded, this reflects the minimum version of all nodes.
        /// </summary>
        [Output("currentNodeVersion")]
        public Output<string> CurrentNodeVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration of etcd encryption.
        /// </summary>
        [Output("databaseEncryption")]
        public Output<Outputs.DatabaseEncryptionResponse> DatabaseEncryption { get; private set; } = null!;

        /// <summary>
        /// The default constraint on the maximum number of pods that can be run simultaneously on a node in the node pool of this cluster. Only honored if cluster created with IP Alias support.
        /// </summary>
        [Output("defaultMaxPodsConstraint")]
        public Output<Outputs.MaxPodsConstraintResponse> DefaultMaxPodsConstraint { get; private set; } = null!;

        /// <summary>
        /// An optional description of this cluster.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Kubernetes alpha features are enabled on this cluster. This includes alpha API groups (e.g. v1beta1) and features that may not be production ready in the kubernetes version of the master and nodes. The cluster has no SLA for uptime and master/node upgrades are disabled. Alpha enabled clusters are automatically deleted thirty days after creation.
        /// </summary>
        [Output("enableKubernetesAlpha")]
        public Output<bool> EnableKubernetesAlpha { get; private set; } = null!;

        /// <summary>
        /// [Output only] The IP address of this cluster's master endpoint. The endpoint can be accessed from the internet at `https://username:password@endpoint/`. See the `masterAuth` property of this resource for username and password information.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// [Output only] The time the cluster will be automatically deleted in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// The initial Kubernetes version for this cluster. Valid versions are those found in validMasterVersions returned by getServerConfig. The version can be upgraded over time; such upgrades are reflected in currentMasterVersion and currentNodeVersion. Users may specify either explicit versions offered by Kubernetes Engine or version aliases, which have the following behavior: - "latest": picks the highest valid Kubernetes version - "1.X": picks the highest valid patch+gke.N patch in the 1.X version - "1.X.Y": picks the highest valid gke.N patch in the 1.X.Y version - "1.X.Y-gke.N": picks an explicit Kubernetes version - "","-": picks the default Kubernetes version
        /// </summary>
        [Output("initialClusterVersion")]
        public Output<string> InitialClusterVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration for cluster IP allocation.
        /// </summary>
        [Output("ipAllocationPolicy")]
        public Output<Outputs.IPAllocationPolicyResponse> IpAllocationPolicy { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the set of labels for this cluster.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Configuration for the legacy ABAC authorization mode.
        /// </summary>
        [Output("legacyAbac")]
        public Output<Outputs.LegacyAbacResponse> LegacyAbac { get; private set; } = null!;

        /// <summary>
        /// [Output only] The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/regions-zones/regions-zones#available) or [region](https://cloud.google.com/compute/docs/regions-zones/regions-zones#available) in which the cluster resides.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the cluster's nodes should be located. This field provides a default value if [NodePool.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations) are not specified during node pool creation. Warning: changing cluster locations will update the [NodePool.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations) of all node pools and will result in nodes being added and/or removed.
        /// </summary>
        [Output("locations")]
        public Output<ImmutableArray<string>> Locations { get; private set; } = null!;

        /// <summary>
        /// The logging service the cluster should use to write logs. Currently available options: * `logging.googleapis.com/kubernetes` - The Cloud Logging service with a Kubernetes-native resource model * `logging.googleapis.com` - The legacy Cloud Logging service (no longer available as of GKE 1.15). * `none` - no logs will be exported from the cluster. If left as an empty string,`logging.googleapis.com/kubernetes` will be used for GKE 1.14+ or `logging.googleapis.com` for earlier versions.
        /// </summary>
        [Output("loggingService")]
        public Output<string> LoggingService { get; private set; } = null!;

        /// <summary>
        /// Configure the maintenance policy for this cluster.
        /// </summary>
        [Output("maintenancePolicy")]
        public Output<Outputs.MaintenancePolicyResponse> MaintenancePolicy { get; private set; } = null!;

        /// <summary>
        /// Configuration for master components.
        /// </summary>
        [Output("master")]
        public Output<Outputs.MasterResponse> Master { get; private set; } = null!;

        /// <summary>
        /// The authentication information for accessing the master endpoint. If unspecified, the defaults are used: For clusters before v1.12, if master_auth is unspecified, `username` will be set to "admin", a random password will be generated, and a client certificate will be issued.
        /// </summary>
        [Output("masterAuth")]
        public Output<Outputs.MasterAuthResponse> MasterAuth { get; private set; } = null!;

        /// <summary>
        /// The configuration options for master authorized networks feature.
        /// </summary>
        [Output("masterAuthorizedNetworksConfig")]
        public Output<Outputs.MasterAuthorizedNetworksConfigResponse> MasterAuthorizedNetworksConfig { get; private set; } = null!;

        /// <summary>
        /// The monitoring service the cluster should use to write metrics. Currently available options: * "monitoring.googleapis.com/kubernetes" - The Cloud Monitoring service with a Kubernetes-native resource model * `monitoring.googleapis.com` - The legacy Cloud Monitoring service (no longer available as of GKE 1.15). * `none` - No metrics will be exported from the cluster. If left as an empty string,`monitoring.googleapis.com/kubernetes` will be used for GKE 1.14+ or `monitoring.googleapis.com` for earlier versions.
        /// </summary>
        [Output("monitoringService")]
        public Output<string> MonitoringService { get; private set; } = null!;

        /// <summary>
        /// The name of this cluster. The name must be unique within this project and location (e.g. zone or region), and can be up to 40 characters with the following restrictions: * Lowercase letters, numbers, and hyphens only. * Must start with a letter. * Must end with a number or a letter.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Google Compute Engine [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks) to which the cluster is connected. If left unspecified, the `default` network will be used. On output this shows the network ID instead of the name.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Configuration for cluster networking.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.NetworkConfigResponse> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration options for the NetworkPolicy feature.
        /// </summary>
        [Output("networkPolicy")]
        public Output<Outputs.NetworkPolicyResponse> NetworkPolicy { get; private set; } = null!;

        /// <summary>
        /// [Output only] The size of the address space on each node for hosting containers. This is provisioned from within the `container_ipv4_cidr` range. This field will only be set when cluster is in route-based network mode.
        /// </summary>
        [Output("nodeIpv4CidrSize")]
        public Output<int> NodeIpv4CidrSize { get; private set; } = null!;

        /// <summary>
        /// Default NodePool settings for the entire cluster. These settings are overridden if specified on the specific NodePool object.
        /// </summary>
        [Output("nodePoolDefaults")]
        public Output<Outputs.NodePoolDefaultsResponse> NodePoolDefaults { get; private set; } = null!;

        /// <summary>
        /// The node pools associated with this cluster. This field should not be set if "node_config" or "initial_node_count" are specified.
        /// </summary>
        [Output("nodePools")]
        public Output<ImmutableArray<Outputs.NodePoolResponse>> NodePools { get; private set; } = null!;

        /// <summary>
        /// Notification configuration of the cluster.
        /// </summary>
        [Output("notificationConfig")]
        public Output<Outputs.NotificationConfigResponse> NotificationConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration for the PodSecurityPolicy feature.
        /// </summary>
        [Output("podSecurityPolicyConfig")]
        public Output<Outputs.PodSecurityPolicyConfigResponse> PodSecurityPolicyConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration for private cluster.
        /// </summary>
        [Output("privateClusterConfig")]
        public Output<Outputs.PrivateClusterConfigResponse> PrivateClusterConfig { get; private set; } = null!;

        /// <summary>
        /// Release channel configuration.
        /// </summary>
        [Output("releaseChannel")]
        public Output<Outputs.ReleaseChannelResponse> ReleaseChannel { get; private set; } = null!;

        /// <summary>
        /// The resource labels for the cluster to use to annotate any related Google Compute Engine resources.
        /// </summary>
        [Output("resourceLabels")]
        public Output<ImmutableDictionary<string, string>> ResourceLabels { get; private set; } = null!;

        /// <summary>
        /// Configuration for exporting resource usages. Resource usage export is disabled when this config unspecified.
        /// </summary>
        [Output("resourceUsageExportConfig")]
        public Output<Outputs.ResourceUsageExportConfigResponse> ResourceUsageExportConfig { get; private set; } = null!;

        /// <summary>
        /// [Output only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output only] The IP address range of the Kubernetes services in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `1.2.3.4/29`). Service addresses are typically put in the last `/16` from the container CIDR.
        /// </summary>
        [Output("servicesIpv4Cidr")]
        public Output<string> ServicesIpv4Cidr { get; private set; } = null!;

        /// <summary>
        /// Shielded Nodes configuration.
        /// </summary>
        [Output("shieldedNodes")]
        public Output<Outputs.ShieldedNodesResponse> ShieldedNodes { get; private set; } = null!;

        /// <summary>
        /// [Output only] The current status of this cluster.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The name of the Google Compute Engine [subnetwork](https://cloud.google.com/compute/docs/subnetworks) to which the cluster is connected. On output this shows the subnetwork ID instead of the name.
        /// </summary>
        [Output("subnetwork")]
        public Output<string> Subnetwork { get; private set; } = null!;

        /// <summary>
        /// Configuration for Cloud TPU support;
        /// </summary>
        [Output("tpuConfig")]
        public Output<Outputs.TpuConfigResponse> TpuConfig { get; private set; } = null!;

        /// <summary>
        /// [Output only] The IP address range of the Cloud TPUs in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `1.2.3.4/29`).
        /// </summary>
        [Output("tpuIpv4CidrBlock")]
        public Output<string> TpuIpv4CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Cluster-level Vertical Pod Autoscaling configuration.
        /// </summary>
        [Output("verticalPodAutoscaling")]
        public Output<Outputs.VerticalPodAutoscalingResponse> VerticalPodAutoscaling { get; private set; } = null!;

        /// <summary>
        /// Configuration for issuance of mTLS keys and certificates to Kubernetes pods.
        /// </summary>
        [Output("workloadCertificates")]
        public Output<Outputs.WorkloadCertificatesResponse> WorkloadCertificates { get; private set; } = null!;

        /// <summary>
        /// Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.
        /// </summary>
        [Output("workloadIdentityConfig")]
        public Output<Outputs.WorkloadIdentityConfigResponse> WorkloadIdentityConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:container/v1beta1:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:container/v1beta1:Cluster", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configurations for the various addons available to run in the cluster.
        /// </summary>
        [Input("addonsConfig")]
        public Input<Inputs.AddonsConfigArgs>? AddonsConfig { get; set; }

        /// <summary>
        /// Configuration controlling RBAC group membership information.
        /// </summary>
        [Input("authenticatorGroupsConfig")]
        public Input<Inputs.AuthenticatorGroupsConfigArgs>? AuthenticatorGroupsConfig { get; set; }

        /// <summary>
        /// Autopilot configuration for the cluster.
        /// </summary>
        [Input("autopilot")]
        public Input<Inputs.AutopilotArgs>? Autopilot { get; set; }

        /// <summary>
        /// Cluster-level autoscaling configuration.
        /// </summary>
        [Input("autoscaling")]
        public Input<Inputs.ClusterAutoscalingArgs>? Autoscaling { get; set; }

        /// <summary>
        /// Configuration for Binary Authorization.
        /// </summary>
        [Input("binaryAuthorization")]
        public Input<Inputs.BinaryAuthorizationArgs>? BinaryAuthorization { get; set; }

        /// <summary>
        /// The IP address range of the container pods in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`). Leave blank to have one automatically chosen or specify a `/14` block in `10.0.0.0/8`.
        /// </summary>
        [Input("clusterIpv4Cidr")]
        public Input<string>? ClusterIpv4Cidr { get; set; }

        /// <summary>
        /// Telemetry integration for the cluster.
        /// </summary>
        [Input("clusterTelemetry")]
        public Input<Inputs.ClusterTelemetryArgs>? ClusterTelemetry { get; set; }

        [Input("conditions")]
        private InputList<Inputs.StatusConditionArgs>? _conditions;

        /// <summary>
        /// Which conditions caused the current cluster state.
        /// </summary>
        public InputList<Inputs.StatusConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.StatusConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Configuration of Confidential Nodes
        /// </summary>
        [Input("confidentialNodes")]
        public Input<Inputs.ConfidentialNodesArgs>? ConfidentialNodes { get; set; }

        /// <summary>
        /// Configuration of etcd encryption.
        /// </summary>
        [Input("databaseEncryption")]
        public Input<Inputs.DatabaseEncryptionArgs>? DatabaseEncryption { get; set; }

        /// <summary>
        /// The default constraint on the maximum number of pods that can be run simultaneously on a node in the node pool of this cluster. Only honored if cluster created with IP Alias support.
        /// </summary>
        [Input("defaultMaxPodsConstraint")]
        public Input<Inputs.MaxPodsConstraintArgs>? DefaultMaxPodsConstraint { get; set; }

        /// <summary>
        /// An optional description of this cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Kubernetes alpha features are enabled on this cluster. This includes alpha API groups (e.g. v1beta1) and features that may not be production ready in the kubernetes version of the master and nodes. The cluster has no SLA for uptime and master/node upgrades are disabled. Alpha enabled clusters are automatically deleted thirty days after creation.
        /// </summary>
        [Input("enableKubernetesAlpha")]
        public Input<bool>? EnableKubernetesAlpha { get; set; }

        /// <summary>
        /// The initial Kubernetes version for this cluster. Valid versions are those found in validMasterVersions returned by getServerConfig. The version can be upgraded over time; such upgrades are reflected in currentMasterVersion and currentNodeVersion. Users may specify either explicit versions offered by Kubernetes Engine or version aliases, which have the following behavior: - "latest": picks the highest valid Kubernetes version - "1.X": picks the highest valid patch+gke.N patch in the 1.X version - "1.X.Y": picks the highest valid gke.N patch in the 1.X.Y version - "1.X.Y-gke.N": picks an explicit Kubernetes version - "","-": picks the default Kubernetes version
        /// </summary>
        [Input("initialClusterVersion")]
        public Input<string>? InitialClusterVersion { get; set; }

        /// <summary>
        /// Configuration for cluster IP allocation.
        /// </summary>
        [Input("ipAllocationPolicy")]
        public Input<Inputs.IPAllocationPolicyArgs>? IpAllocationPolicy { get; set; }

        /// <summary>
        /// Configuration for the legacy ABAC authorization mode.
        /// </summary>
        [Input("legacyAbac")]
        public Input<Inputs.LegacyAbacArgs>? LegacyAbac { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the cluster's nodes should be located. This field provides a default value if [NodePool.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations) are not specified during node pool creation. Warning: changing cluster locations will update the [NodePool.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations) of all node pools and will result in nodes being added and/or removed.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// The logging service the cluster should use to write logs. Currently available options: * `logging.googleapis.com/kubernetes` - The Cloud Logging service with a Kubernetes-native resource model * `logging.googleapis.com` - The legacy Cloud Logging service (no longer available as of GKE 1.15). * `none` - no logs will be exported from the cluster. If left as an empty string,`logging.googleapis.com/kubernetes` will be used for GKE 1.14+ or `logging.googleapis.com` for earlier versions.
        /// </summary>
        [Input("loggingService")]
        public Input<string>? LoggingService { get; set; }

        /// <summary>
        /// Configure the maintenance policy for this cluster.
        /// </summary>
        [Input("maintenancePolicy")]
        public Input<Inputs.MaintenancePolicyArgs>? MaintenancePolicy { get; set; }

        /// <summary>
        /// Configuration for master components.
        /// </summary>
        [Input("master")]
        public Input<Inputs.MasterArgs>? Master { get; set; }

        /// <summary>
        /// The authentication information for accessing the master endpoint. If unspecified, the defaults are used: For clusters before v1.12, if master_auth is unspecified, `username` will be set to "admin", a random password will be generated, and a client certificate will be issued.
        /// </summary>
        [Input("masterAuth")]
        public Input<Inputs.MasterAuthArgs>? MasterAuth { get; set; }

        /// <summary>
        /// The configuration options for master authorized networks feature.
        /// </summary>
        [Input("masterAuthorizedNetworksConfig")]
        public Input<Inputs.MasterAuthorizedNetworksConfigArgs>? MasterAuthorizedNetworksConfig { get; set; }

        /// <summary>
        /// The monitoring service the cluster should use to write metrics. Currently available options: * "monitoring.googleapis.com/kubernetes" - The Cloud Monitoring service with a Kubernetes-native resource model * `monitoring.googleapis.com` - The legacy Cloud Monitoring service (no longer available as of GKE 1.15). * `none` - No metrics will be exported from the cluster. If left as an empty string,`monitoring.googleapis.com/kubernetes` will be used for GKE 1.14+ or `monitoring.googleapis.com` for earlier versions.
        /// </summary>
        [Input("monitoringService")]
        public Input<string>? MonitoringService { get; set; }

        /// <summary>
        /// The name of this cluster. The name must be unique within this project and location (e.g. zone or region), and can be up to 40 characters with the following restrictions: * Lowercase letters, numbers, and hyphens only. * Must start with a letter. * Must end with a number or a letter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Google Compute Engine [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks) to which the cluster is connected. If left unspecified, the `default` network will be used. On output this shows the network ID instead of the name.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Configuration for cluster networking.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.NetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// Configuration options for the NetworkPolicy feature.
        /// </summary>
        [Input("networkPolicy")]
        public Input<Inputs.NetworkPolicyArgs>? NetworkPolicy { get; set; }

        /// <summary>
        /// Default NodePool settings for the entire cluster. These settings are overridden if specified on the specific NodePool object.
        /// </summary>
        [Input("nodePoolDefaults")]
        public Input<Inputs.NodePoolDefaultsArgs>? NodePoolDefaults { get; set; }

        [Input("nodePools")]
        private InputList<Inputs.NodePoolArgs>? _nodePools;

        /// <summary>
        /// The node pools associated with this cluster. This field should not be set if "node_config" or "initial_node_count" are specified.
        /// </summary>
        public InputList<Inputs.NodePoolArgs> NodePools
        {
            get => _nodePools ?? (_nodePools = new InputList<Inputs.NodePoolArgs>());
            set => _nodePools = value;
        }

        /// <summary>
        /// Notification configuration of the cluster.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.NotificationConfigArgs>? NotificationConfig { get; set; }

        /// <summary>
        /// The parent (project and location) where the cluster will be created. Specified in the format `projects/*/locations/*`.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Configuration for the PodSecurityPolicy feature.
        /// </summary>
        [Input("podSecurityPolicyConfig")]
        public Input<Inputs.PodSecurityPolicyConfigArgs>? PodSecurityPolicyConfig { get; set; }

        /// <summary>
        /// Configuration for private cluster.
        /// </summary>
        [Input("privateClusterConfig")]
        public Input<Inputs.PrivateClusterConfigArgs>? PrivateClusterConfig { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Release channel configuration.
        /// </summary>
        [Input("releaseChannel")]
        public Input<Inputs.ReleaseChannelArgs>? ReleaseChannel { get; set; }

        [Input("resourceLabels")]
        private InputMap<string>? _resourceLabels;

        /// <summary>
        /// The resource labels for the cluster to use to annotate any related Google Compute Engine resources.
        /// </summary>
        public InputMap<string> ResourceLabels
        {
            get => _resourceLabels ?? (_resourceLabels = new InputMap<string>());
            set => _resourceLabels = value;
        }

        /// <summary>
        /// Configuration for exporting resource usages. Resource usage export is disabled when this config unspecified.
        /// </summary>
        [Input("resourceUsageExportConfig")]
        public Input<Inputs.ResourceUsageExportConfigArgs>? ResourceUsageExportConfig { get; set; }

        /// <summary>
        /// Shielded Nodes configuration.
        /// </summary>
        [Input("shieldedNodes")]
        public Input<Inputs.ShieldedNodesArgs>? ShieldedNodes { get; set; }

        /// <summary>
        /// The name of the Google Compute Engine [subnetwork](https://cloud.google.com/compute/docs/subnetworks) to which the cluster is connected. On output this shows the subnetwork ID instead of the name.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        /// <summary>
        /// Configuration for Cloud TPU support;
        /// </summary>
        [Input("tpuConfig")]
        public Input<Inputs.TpuConfigArgs>? TpuConfig { get; set; }

        /// <summary>
        /// Cluster-level Vertical Pod Autoscaling configuration.
        /// </summary>
        [Input("verticalPodAutoscaling")]
        public Input<Inputs.VerticalPodAutoscalingArgs>? VerticalPodAutoscaling { get; set; }

        /// <summary>
        /// Configuration for issuance of mTLS keys and certificates to Kubernetes pods.
        /// </summary>
        [Input("workloadCertificates")]
        public Input<Inputs.WorkloadCertificatesArgs>? WorkloadCertificates { get; set; }

        /// <summary>
        /// Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.
        /// </summary>
        [Input("workloadIdentityConfig")]
        public Input<Inputs.WorkloadIdentityConfigArgs>? WorkloadIdentityConfig { get; set; }

        public ClusterArgs()
        {
        }
    }
}
