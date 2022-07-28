// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RemoteBuildExecution.V1Alpha.Inputs
{

    /// <summary>
    /// AcceleratorConfig defines the accelerator cards to attach to the VM.
    /// </summary>
    public sealed class GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of guest accelerator cards exposed to each VM.
        /// </summary>
        [Input("acceleratorCount")]
        public Input<string>? AcceleratorCount { get; set; }

        /// <summary>
        /// The type of accelerator to attach to each VM, e.g. "nvidia-tesla-k80" for nVidia Tesla K80.
        /// </summary>
        [Input("acceleratorType")]
        public Input<string>? AcceleratorType { get; set; }

        public GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigArgs()
        {
        }
        public static new GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigArgs Empty => new GoogleDevtoolsRemotebuildexecutionAdminV1alphaAcceleratorConfigArgs();
    }
}
