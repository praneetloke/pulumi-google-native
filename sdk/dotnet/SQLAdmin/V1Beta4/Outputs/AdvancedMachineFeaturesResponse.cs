// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Outputs
{

    /// <summary>
    /// Specifies options for controlling advanced machine features.
    /// </summary>
    [OutputType]
    public sealed class AdvancedMachineFeaturesResponse
    {
        /// <summary>
        /// The number of threads per physical core.
        /// </summary>
        public readonly int ThreadsPerCore;

        [OutputConstructor]
        private AdvancedMachineFeaturesResponse(int threadsPerCore)
        {
            ThreadsPerCore = threadsPerCore;
        }
    }
}
