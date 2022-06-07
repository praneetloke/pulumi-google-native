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
    /// AcceleratorConfig represents a Hardware Accelerator request.
    /// </summary>
    public sealed class AcceleratorConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of the accelerator cards exposed to an instance.
        /// </summary>
        [Input("acceleratorCount")]
        public Input<string>? AcceleratorCount { get; set; }

        /// <summary>
        /// The accelerator type resource name. List of supported accelerators [here](https://cloud.google.com/compute/docs/gpus)
        /// </summary>
        [Input("acceleratorType")]
        public Input<string>? AcceleratorType { get; set; }

        /// <summary>
        /// Size of partitions to create on the GPU. Valid values are described in the NVIDIA [mig user guide](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
        /// </summary>
        [Input("gpuPartitionSize")]
        public Input<string>? GpuPartitionSize { get; set; }

        /// <summary>
        /// The configuration for GPU sharing options.
        /// </summary>
        [Input("gpuSharingConfig")]
        public Input<Inputs.GPUSharingConfigArgs>? GpuSharingConfig { get; set; }

        public AcceleratorConfigArgs()
        {
        }
    }
}
