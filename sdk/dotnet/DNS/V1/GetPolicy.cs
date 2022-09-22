// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1
{
    public static class GetPolicy
    {
        /// <summary>
        /// Fetches the representation of an existing Policy.
        /// </summary>
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("google-native:dns/v1:getPolicy", args ?? new GetPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Fetches the representation of an existing Policy.
        /// </summary>
        public static Output<GetPolicyResult> Invoke(GetPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPolicyResult>("google-native:dns/v1:getPolicy", args ?? new GetPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("clientOperationId")]
        public string? ClientOperationId { get; set; }

        [Input("policy", required: true)]
        public string Policy { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetPolicyArgs()
        {
        }
        public static new GetPolicyArgs Empty => new GetPolicyArgs();
    }

    public sealed class GetPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetPolicyInvokeArgs()
        {
        }
        public static new GetPolicyInvokeArgs Empty => new GetPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        /// </summary>
        public readonly Outputs.PolicyAlternativeNameServerConfigResponse AlternativeNameServerConfig;
        /// <summary>
        /// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
        /// </summary>
        public readonly bool EnableInboundForwarding;
        /// <summary>
        /// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
        /// </summary>
        public readonly bool EnableLogging;
        public readonly string Kind;
        /// <summary>
        /// User-assigned name for this policy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of network names specifying networks to which this policy is applied.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyNetworkResponse> Networks;

        [OutputConstructor]
        private GetPolicyResult(
            Outputs.PolicyAlternativeNameServerConfigResponse alternativeNameServerConfig,

            string description,

            bool enableInboundForwarding,

            bool enableLogging,

            string kind,

            string name,

            ImmutableArray<Outputs.PolicyNetworkResponse> networks)
        {
            AlternativeNameServerConfig = alternativeNameServerConfig;
            Description = description;
            EnableInboundForwarding = enableInboundForwarding;
            EnableLogging = enableLogging;
            Kind = kind;
            Name = name;
            Networks = networks;
        }
    }
}
