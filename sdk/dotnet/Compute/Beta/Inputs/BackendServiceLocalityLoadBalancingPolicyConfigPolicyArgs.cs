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
    /// The configuration for a built-in load balancing policy.
    /// </summary>
    public sealed class BackendServiceLocalityLoadBalancingPolicyConfigPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of a locality load balancer policy to be used. The value should be one of the predefined ones as supported by localityLbPolicy, although at the moment only ROUND_ROBIN is supported. This field should only be populated when the customPolicy field is not used. Note that specifying the same policy more than once for a backend is not a valid configuration and will be rejected.
        /// </summary>
        [Input("name")]
        public Input<Pulumi.GoogleNative.Compute.Beta.BackendServiceLocalityLoadBalancingPolicyConfigPolicyName>? Name { get; set; }

        public BackendServiceLocalityLoadBalancingPolicyConfigPolicyArgs()
        {
        }
        public static new BackendServiceLocalityLoadBalancingPolicyConfigPolicyArgs Empty => new BackendServiceLocalityLoadBalancingPolicyConfigPolicyArgs();
    }
}
