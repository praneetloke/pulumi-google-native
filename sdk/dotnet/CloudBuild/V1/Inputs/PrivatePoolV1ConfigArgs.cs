// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// Configuration for a V1 `PrivatePool`.
    /// </summary>
    public sealed class PrivatePoolV1ConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network configuration for the pool.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.NetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// Machine configuration for the workers in the pool.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.WorkerConfigArgs>? WorkerConfig { get; set; }

        public PrivatePoolV1ConfigArgs()
        {
        }
        public static new PrivatePoolV1ConfigArgs Empty => new PrivatePoolV1ConfigArgs();
    }
}
