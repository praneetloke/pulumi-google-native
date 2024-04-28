// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Beta.Outputs
{

    /// <summary>
    /// ResourceList contains container resource requirements.
    /// </summary>
    [OutputType]
    public sealed class PolicyControllerResourceListResponse
    {
        /// <summary>
        /// CPU requirement expressed in Kubernetes resource units.
        /// </summary>
        public readonly string Cpu;
        /// <summary>
        /// Memory requirement expressed in Kubernetes resource units.
        /// </summary>
        public readonly string Memory;

        [OutputConstructor]
        private PolicyControllerResourceListResponse(
            string cpu,

            string memory)
        {
            Cpu = cpu;
            Memory = memory;
        }
    }
}
