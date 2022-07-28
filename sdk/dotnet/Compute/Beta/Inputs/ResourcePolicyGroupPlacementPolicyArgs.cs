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
    /// A GroupPlacementPolicy specifies resource placement configuration. It specifies the failure bucket separation as well as network locality
    /// </summary>
    public sealed class ResourcePolicyGroupPlacementPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of availability domains to spread instances across. If two instances are in different availability domain, they are not in the same low latency network.
        /// </summary>
        [Input("availabilityDomainCount")]
        public Input<int>? AvailabilityDomainCount { get; set; }

        /// <summary>
        /// Specifies network collocation
        /// </summary>
        [Input("collocation")]
        public Input<Pulumi.GoogleNative.Compute.Beta.ResourcePolicyGroupPlacementPolicyCollocation>? Collocation { get; set; }

        /// <summary>
        /// Number of VMs in this placement group. Google does not recommend that you use this field unless you use a compact policy and you want your policy to work only if it contains this exact number of VMs.
        /// </summary>
        [Input("vmCount")]
        public Input<int>? VmCount { get; set; }

        public ResourcePolicyGroupPlacementPolicyArgs()
        {
        }
        public static new ResourcePolicyGroupPlacementPolicyArgs Empty => new ResourcePolicyGroupPlacementPolicyArgs();
    }
}
