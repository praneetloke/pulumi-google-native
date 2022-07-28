// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1Alpha2.Inputs
{

    /// <summary>
    /// Network describes the network configuration for a `WorkerPool`.
    /// </summary>
    public sealed class NetworkConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to WorkerPool.project_id on the default network. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number, such as `12345`, and {network} is the name of a VPC network in the project.
        /// </summary>
        [Input("peeredNetwork", required: true)]
        public Input<string> PeeredNetwork { get; set; } = null!;

        public NetworkConfigArgs()
        {
        }
        public static new NetworkConfigArgs Empty => new NetworkConfigArgs();
    }
}
