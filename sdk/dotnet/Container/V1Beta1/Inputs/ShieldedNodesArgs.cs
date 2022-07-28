// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration of Shielded Nodes feature.
    /// </summary>
    public sealed class ShieldedNodesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether Shielded Nodes features are enabled on all nodes in this cluster.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ShieldedNodesArgs()
        {
        }
        public static new ShieldedNodesArgs Empty => new ShieldedNodesArgs();
    }
}
