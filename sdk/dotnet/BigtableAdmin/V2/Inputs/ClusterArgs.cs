// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Inputs
{

    /// <summary>
    /// A resizable group of nodes in a particular cloud location, capable of serving all Tables in the parent Instance.
    /// </summary>
    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for this cluster.
        /// </summary>
        [Input("clusterConfig")]
        public Input<Inputs.ClusterConfigArgs>? ClusterConfig { get; set; }

        /// <summary>
        /// Immutable. The type of storage used by this cluster to serve its parent instance's tables, unless explicitly overridden.
        /// </summary>
        [Input("defaultStorageType")]
        public Input<Pulumi.GoogleNative.BigtableAdmin.V2.ClusterDefaultStorageType>? DefaultStorageType { get; set; }

        /// <summary>
        /// Immutable. The encryption configuration for CMEK-protected clusters.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.EncryptionConfigArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// Immutable. The location where this cluster's nodes and storage reside. For best performance, clients should be located as close as possible to this cluster. Currently only zones are supported, so values should be of the form `projects/{project}/locations/{zone}`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The unique name of the cluster. Values are of the form `projects/{project}/instances/{instance}/clusters/a-z*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of nodes in the cluster. If no value is set, Cloud Bigtable automatically allocates nodes based on your data footprint and optimized for 50% storage utilization.
        /// </summary>
        [Input("serveNodes")]
        public Input<int>? ServeNodes { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }
}
