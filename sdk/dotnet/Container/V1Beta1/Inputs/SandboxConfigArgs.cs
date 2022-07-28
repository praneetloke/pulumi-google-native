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
    /// SandboxConfig contains configurations of the sandbox to use for the node.
    /// </summary>
    public sealed class SandboxConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the sandbox to use for the node (e.g. 'gvisor')
        /// </summary>
        [Input("sandboxType")]
        public Input<string>? SandboxType { get; set; }

        /// <summary>
        /// Type of the sandbox to use for the node.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Container.V1Beta1.SandboxConfigType>? Type { get; set; }

        public SandboxConfigArgs()
        {
        }
        public static new SandboxConfigArgs Empty => new SandboxConfigArgs();
    }
}
