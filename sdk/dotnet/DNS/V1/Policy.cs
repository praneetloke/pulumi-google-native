// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1
{
    /// <summary>
    /// Creates a new Policy.
    /// </summary>
    [GoogleNativeResourceType("google-native:dns/v1:Policy")]
    public partial class Policy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        /// </summary>
        [Output("alternativeNameServerConfig")]
        public Output<Outputs.PolicyAlternativeNameServerConfigResponse> AlternativeNameServerConfig { get; private set; } = null!;

        /// <summary>
        /// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        /// </summary>
        [Output("clientOperationId")]
        public Output<string?> ClientOperationId { get; private set; } = null!;

        /// <summary>
        /// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
        /// </summary>
        [Output("enableInboundForwarding")]
        public Output<bool> EnableInboundForwarding { get; private set; } = null!;

        /// <summary>
        /// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
        /// </summary>
        [Output("enableLogging")]
        public Output<bool> EnableLogging { get; private set; } = null!;

        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// User-assigned name for this policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of network names specifying networks to which this policy is applied.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<Outputs.PolicyNetworkResponse>> Networks { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:dns/v1:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dns/v1:Policy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, options);
        }
    }

    public sealed class PolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        /// </summary>
        [Input("alternativeNameServerConfig")]
        public Input<Inputs.PolicyAlternativeNameServerConfigArgs>? AlternativeNameServerConfig { get; set; }

        /// <summary>
        /// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        /// </summary>
        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        /// <summary>
        /// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
        /// </summary>
        [Input("enableInboundForwarding")]
        public Input<bool>? EnableInboundForwarding { get; set; }

        /// <summary>
        /// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// User-assigned name for this policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networks")]
        private InputList<Inputs.PolicyNetworkArgs>? _networks;

        /// <summary>
        /// List of network names specifying networks to which this policy is applied.
        /// </summary>
        public InputList<Inputs.PolicyNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.PolicyNetworkArgs>());
            set => _networks = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public PolicyArgs()
        {
        }
        public static new PolicyArgs Empty => new PolicyArgs();
    }
}
