// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// The configuration for a custom policy implemented by the user and deployed with the client.
    /// </summary>
    public sealed class BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional, arbitrary JSON object with configuration data, understood by a locally installed custom policy implementation.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// Identifies the custom policy. The value should match the type the custom implementation is registered with on the gRPC clients. It should follow protocol buffer message naming conventions and include the full path (e.g. myorg.CustomLbPolicy). The maximum length is 256 characters. Note that specifying the same custom policy more than once for a backend is not a valid configuration and will be rejected.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicyArgs()
        {
        }
    }
}
