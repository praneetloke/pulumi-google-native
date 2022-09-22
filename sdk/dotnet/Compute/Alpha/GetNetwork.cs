// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetNetwork
    {
        /// <summary>
        /// Returns the specified network. Gets a list of available networks by making a list() request.
        /// </summary>
        public static Task<GetNetworkResult> InvokeAsync(GetNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("google-native:compute/alpha:getNetwork", args ?? new GetNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified network. Gets a list of available networks by making a list() request.
        /// </summary>
        public static Output<GetNetworkResult> Invoke(GetNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkResult>("google-native:compute/alpha:getNetwork", args ?? new GetNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkArgs : global::Pulumi.InvokeArgs
    {
        [Input("network", required: true)]
        public string Network { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetNetworkArgs()
        {
        }
        public static new GetNetworkArgs Empty => new GetNetworkArgs();
    }

    public sealed class GetNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetNetworkInvokeArgs()
        {
        }
        public static new GetNetworkInvokeArgs Empty => new GetNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkResult
    {
        /// <summary>
        /// Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
        /// </summary>
        public readonly bool AutoCreateSubnetworks;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
        /// </summary>
        public readonly bool EnableUlaInternalIpv6;
        /// <summary>
        /// URL of the firewall policy the network is associated with.
        /// </summary>
        public readonly string FirewallPolicy;
        /// <summary>
        /// The gateway address for default routing out of the network, selected by GCP.
        /// </summary>
        public readonly string GatewayIPv4;
        /// <summary>
        /// When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
        /// </summary>
        public readonly string InternalIpv6Range;
        /// <summary>
        /// Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
        /// </summary>
        public readonly string Ipv4Range;
        /// <summary>
        /// Type of the resource. Always compute#network for networks.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Maximum Transmission Unit in bytes. The minimum value for this field is 1300 and the maximum value is 8896. The suggested value is 1500, which is the default MTU used on the Internet, or 8896 if you want to use Jumbo frames. If unspecified, the value defaults to 1460.
        /// </summary>
        public readonly int Mtu;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
        /// </summary>
        public readonly string NetworkFirewallPolicyEnforcementOrder;
        /// <summary>
        /// A list of network peerings for the resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkPeeringResponse> Peerings;
        /// <summary>
        /// URL of the region where the regional network resides. This field is not applicable to global network. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
        /// </summary>
        public readonly Outputs.NetworkRoutingConfigResponse RoutingConfig;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// Server-defined fully-qualified URLs for all subnetworks in this VPC network.
        /// </summary>
        public readonly ImmutableArray<string> Subnetworks;

        [OutputConstructor]
        private GetNetworkResult(
            bool autoCreateSubnetworks,

            string creationTimestamp,

            string description,

            bool enableUlaInternalIpv6,

            string firewallPolicy,

            string gatewayIPv4,

            string internalIpv6Range,

            string ipv4Range,

            string kind,

            int mtu,

            string name,

            string networkFirewallPolicyEnforcementOrder,

            ImmutableArray<Outputs.NetworkPeeringResponse> peerings,

            string region,

            Outputs.NetworkRoutingConfigResponse routingConfig,

            string selfLink,

            string selfLinkWithId,

            ImmutableArray<string> subnetworks)
        {
            AutoCreateSubnetworks = autoCreateSubnetworks;
            CreationTimestamp = creationTimestamp;
            Description = description;
            EnableUlaInternalIpv6 = enableUlaInternalIpv6;
            FirewallPolicy = firewallPolicy;
            GatewayIPv4 = gatewayIPv4;
            InternalIpv6Range = internalIpv6Range;
            Ipv4Range = ipv4Range;
            Kind = kind;
            Mtu = mtu;
            Name = name;
            NetworkFirewallPolicyEnforcementOrder = networkFirewallPolicyEnforcementOrder;
            Peerings = peerings;
            Region = region;
            RoutingConfig = routingConfig;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            Subnetworks = subnetworks;
        }
    }
}
