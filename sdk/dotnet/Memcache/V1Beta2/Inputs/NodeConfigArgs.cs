// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Memcache.V1Beta2.Inputs
{

    /// <summary>
    /// Configuration for a Memcached Node.
    /// </summary>
    public sealed class NodeConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of cpus per Memcached node.
        /// </summary>
        [Input("cpuCount", required: true)]
        public Input<int> CpuCount { get; set; } = null!;

        /// <summary>
        /// Memory size in MiB for each Memcached node.
        /// </summary>
        [Input("memorySizeMb", required: true)]
        public Input<int> MemorySizeMb { get; set; } = null!;

        public NodeConfigArgs()
        {
        }
        public static new NodeConfigArgs Empty => new NodeConfigArgs();
    }
}
