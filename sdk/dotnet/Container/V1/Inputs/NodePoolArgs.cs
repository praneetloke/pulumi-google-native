// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// NodePool contains the name and configuration for a cluster's node pool. Node pools are a set of nodes (i.e. VM's), with a common configuration and specification, under the control of the cluster master. They may have a set of Kubernetes labels applied to them, which may be used to reference them during pod scheduling. They may also be resized up or down, to accommodate the workload.
    /// </summary>
    public sealed class NodePoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        /// </summary>
        [Input("autoscaling")]
        public Input<Inputs.NodePoolAutoscalingArgs>? Autoscaling { get; set; }

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
        /// Upgrade settings control disruption and speed of the upgrade.
        /// </summary>
        [Input("upgradeSettings")]
        public Input<Inputs.UpgradeSettingsArgs>? UpgradeSettings { get; set; }

        /// <summary>
        /// The version of the Kubernetes of this node.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NodePoolArgs()
        {
        }
        public static new NodePoolArgs Empty => new NodePoolArgs();
    }
}
