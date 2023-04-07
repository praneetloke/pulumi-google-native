// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a network in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:Network")]
    public partial class Network : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
        /// </summary>
        [Output("autoCreateSubnetworks")]
        public Output<bool> AutoCreateSubnetworks { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
        /// </summary>
        [Output("enableUlaInternalIpv6")]
        public Output<bool> EnableUlaInternalIpv6 { get; private set; } = null!;

        /// <summary>
        /// URL of the firewall policy the network is associated with.
        /// </summary>
        [Output("firewallPolicy")]
        public Output<string> FirewallPolicy { get; private set; } = null!;

        /// <summary>
        /// The gateway address for default routing out of the network, selected by Google Cloud.
        /// </summary>
        [Output("gatewayIPv4")]
        public Output<string> GatewayIPv4 { get; private set; } = null!;

        /// <summary>
        /// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
        /// </summary>
        [Output("internalIpv6Range")]
        public Output<string> InternalIpv6Range { get; private set; } = null!;

        /// <summary>
        /// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
        /// </summary>
        [Output("ipv4Range")]
        public Output<string> Ipv4Range { get; private set; } = null!;

        /// <summary>
        /// Type of the resource. Always compute#network for networks.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Maximum Transmission Unit in bytes. The minimum value for this field is 1300 and the maximum value is 8896. The suggested value is 1500, which is the default MTU used on the Internet, or 8896 if you want to use Jumbo frames. If unspecified, the value defaults to 1460.
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
        /// </summary>
        [Output("networkFirewallPolicyEnforcementOrder")]
        public Output<string> NetworkFirewallPolicyEnforcementOrder { get; private set; } = null!;

        /// <summary>
        /// A list of network peerings for the resource.
        /// </summary>
        [Output("peerings")]
        public Output<ImmutableArray<Outputs.NetworkPeeringResponse>> Peerings { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
        /// </summary>
        [Output("routingConfig")]
        public Output<Outputs.NetworkRoutingConfigResponse> RoutingConfig { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// Server-defined fully-qualified URLs for all subnetworks in this VPC network.
        /// </summary>
        [Output("subnetworks")]
        public Output<ImmutableArray<string>> Subnetworks { get; private set; } = null!;


        /// <summary>
        /// Create a Network resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Network(string name, NetworkArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:Network", name, args ?? new NetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Network(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:Network", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Network resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Network Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Network(name, id, options);
        }
    }

    public sealed class NetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
        /// </summary>
        [Input("autoCreateSubnetworks")]
        public Input<bool>? AutoCreateSubnetworks { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
        /// </summary>
        [Input("enableUlaInternalIpv6")]
        public Input<bool>? EnableUlaInternalIpv6 { get; set; }

        /// <summary>
        /// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
        /// </summary>
        [Input("internalIpv6Range")]
        public Input<string>? InternalIpv6Range { get; set; }

        /// <summary>
        /// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
        /// </summary>
        [Input("ipv4Range")]
        public Input<string>? Ipv4Range { get; set; }

        /// <summary>
        /// Maximum Transmission Unit in bytes. The minimum value for this field is 1300 and the maximum value is 8896. The suggested value is 1500, which is the default MTU used on the Internet, or 8896 if you want to use Jumbo frames. If unspecified, the value defaults to 1460.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
        /// </summary>
        [Input("networkFirewallPolicyEnforcementOrder")]
        public Input<Pulumi.GoogleNative.Compute.V1.NetworkNetworkFirewallPolicyEnforcementOrder>? NetworkFirewallPolicyEnforcementOrder { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
        /// </summary>
        [Input("routingConfig")]
        public Input<Inputs.NetworkRoutingConfigArgs>? RoutingConfig { get; set; }

        public NetworkArgs()
        {
        }
        public static new NetworkArgs Empty => new NetworkArgs();
    }
}
