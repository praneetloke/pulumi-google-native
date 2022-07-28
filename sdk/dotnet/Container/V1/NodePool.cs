// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1
{
    /// <summary>
    /// Creates a node pool for a cluster.
    /// </summary>
    [GoogleNativeResourceType("google-native:container/v1:NodePool")]
    public partial class NodePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        /// </summary>
        [Output("autoscaling")]
        public Output<Outputs.NodePoolAutoscalingResponse> Autoscaling { get; private set; } = null!;

        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Which conditions caused the current node pool state.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.StatusConditionResponse>> Conditions { get; private set; } = null!;

        /// <summary>
        /// The node configuration of the pool.
        /// </summary>
        [Output("config")]
        public Output<Outputs.NodeConfigResponse> Config { get; private set; } = null!;

        /// <summary>
        /// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        /// </summary>
        [Output("initialNodeCount")]
        public Output<int> InitialNodeCount { get; private set; } = null!;

        /// <summary>
        /// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool. During the node pool blue-green upgrade operation, the URLs contain both blue and green resources.
        /// </summary>
        [Output("instanceGroupUrls")]
        public Output<ImmutableArray<string>> InstanceGroupUrls { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        /// </summary>
        [Output("locations")]
        public Output<ImmutableArray<string>> Locations { get; private set; } = null!;

        /// <summary>
        /// NodeManagement configuration for this NodePool.
        /// </summary>
        [Output("management")]
        public Output<Outputs.NodeManagementResponse> Management { get; private set; } = null!;

        /// <summary>
        /// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        /// </summary>
        [Output("maxPodsConstraint")]
        public Output<Outputs.MaxPodsConstraintResponse> MaxPodsConstraint { get; private set; } = null!;

        /// <summary>
        /// The name of the node pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.NodeNetworkConfigResponse> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// [Output only] The pod CIDR block size per node in this node pool.
        /// </summary>
        [Output("podIpv4CidrSize")]
        public Output<int> PodIpv4CidrSize { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// [Output only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output only] The status of the nodes in this pool instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// [Output only] Deprecated. Use conditions instead. Additional information about the current status of this node pool instance, if available.
        /// </summary>
        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;

        /// <summary>
        /// [Output only] Update info contains relevant information during a node pool update.
        /// </summary>
        [Output("updateInfo")]
        public Output<Outputs.UpdateInfoResponse> UpdateInfo { get; private set; } = null!;

        /// <summary>
        /// Upgrade settings control disruption and speed of the upgrade.
        /// </summary>
        [Output("upgradeSettings")]
        public Output<Outputs.UpgradeSettingsResponse> UpgradeSettings { get; private set; } = null!;

        /// <summary>
        /// The version of the Kubernetes of this node.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a NodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodePool(string name, NodePoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:container/v1:NodePool", name, args ?? new NodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodePool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:container/v1:NodePool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "clusterId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodePool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NodePool(name, id, options);
        }
    }

    public sealed class NodePoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        /// </summary>
        [Input("autoscaling")]
        public Input<Inputs.NodePoolAutoscalingArgs>? Autoscaling { get; set; }

        /// <summary>
        /// Deprecated. The name of the cluster. This field has been deprecated and replaced by the parent field.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("conditions")]
        private InputList<Inputs.StatusConditionArgs>? _conditions;

        /// <summary>
        /// Which conditions caused the current node pool state.
        /// </summary>
        public InputList<Inputs.StatusConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.StatusConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The node configuration of the pool.
        /// </summary>
        [Input("config")]
        public Input<Inputs.NodeConfigArgs>? Config { get; set; }

        /// <summary>
        /// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        /// </summary>
        [Input("initialNodeCount")]
        public Input<int>? InitialNodeCount { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// NodeManagement configuration for this NodePool.
        /// </summary>
        [Input("management")]
        public Input<Inputs.NodeManagementArgs>? Management { get; set; }

        /// <summary>
        /// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        /// </summary>
        [Input("maxPodsConstraint")]
        public Input<Inputs.MaxPodsConstraintArgs>? MaxPodsConstraint { get; set; }

        /// <summary>
        /// The name of the node pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.NodeNetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// The parent (project, location, cluster name) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Deprecated. The Google Developers Console [project ID or project number](https://cloud.google.com/resource-manager/docs/creating-managing-projects). This field has been deprecated and replaced by the parent field.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Upgrade settings control disruption and speed of the upgrade.
        /// </summary>
        [Input("upgradeSettings")]
        public Input<Inputs.UpgradeSettingsArgs>? UpgradeSettings { get; set; }

        /// <summary>
        /// The version of the Kubernetes of this node.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Deprecated. The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field has been deprecated and replaced by the parent field.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public NodePoolArgs()
        {
        }
        public static new NodePoolArgs Empty => new NodePoolArgs();
    }
}
