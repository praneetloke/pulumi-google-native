// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// A selector that chooses target cluster for jobs based on metadata.
    /// </summary>
    public sealed class ClusterSelectorArgs : global::Pulumi.ResourceArgs
    {
        [Input("clusterLabels", required: true)]
        private InputMap<string>? _clusterLabels;

        /// <summary>
        /// The cluster labels. Cluster must have all labels to match.
        /// </summary>
        public InputMap<string> ClusterLabels
        {
            get => _clusterLabels ?? (_clusterLabels = new InputMap<string>());
            set => _clusterLabels = value;
        }

        /// <summary>
        /// Optional. The zone where workflow process executes. This parameter does not affect the selection of the cluster.If unspecified, the zone of the first cluster matching the selector is used.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ClusterSelectorArgs()
        {
        }
        public static new ClusterSelectorArgs Empty => new ClusterSelectorArgs();
    }
}
