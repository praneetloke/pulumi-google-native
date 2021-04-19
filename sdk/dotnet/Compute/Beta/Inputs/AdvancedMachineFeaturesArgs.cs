// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Specifies options for controlling advanced machine features. Options that would traditionally be configured in a BIOS belong here. Features that require operating system support may have corresponding entries in the GuestOsFeatures of an Image (e.g., whether or not the OS in the Image supports nested virtualization being enabled or disabled).
    /// </summary>
    public sealed class AdvancedMachineFeaturesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable nested virtualization or not (default is false).
        /// </summary>
        [Input("enableNestedVirtualization")]
        public Input<bool>? EnableNestedVirtualization { get; set; }

        /// <summary>
        /// The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
        /// </summary>
        [Input("threadsPerCore")]
        public Input<int>? ThreadsPerCore { get; set; }

        public AdvancedMachineFeaturesArgs()
        {
        }
    }
}
