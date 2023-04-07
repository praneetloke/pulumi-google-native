// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetGlobalForwardingRule
    {
        /// <summary>
        /// Returns the specified GlobalForwardingRule resource. Gets a list of available forwarding rules by making a list() request.
        /// </summary>
        public static Task<GetGlobalForwardingRuleResult> InvokeAsync(GetGlobalForwardingRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGlobalForwardingRuleResult>("google-native:compute/v1:getGlobalForwardingRule", args ?? new GetGlobalForwardingRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified GlobalForwardingRule resource. Gets a list of available forwarding rules by making a list() request.
        /// </summary>
        public static Output<GetGlobalForwardingRuleResult> Invoke(GetGlobalForwardingRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGlobalForwardingRuleResult>("google-native:compute/v1:getGlobalForwardingRule", args ?? new GetGlobalForwardingRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGlobalForwardingRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("forwardingRule", required: true)]
        public string ForwardingRule { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGlobalForwardingRuleArgs()
        {
        }
        public static new GetGlobalForwardingRuleArgs Empty => new GetGlobalForwardingRuleArgs();
    }

    public sealed class GetGlobalForwardingRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("forwardingRule", required: true)]
        public Input<string> ForwardingRule { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGlobalForwardingRuleInvokeArgs()
        {
        }
        public static new GetGlobalForwardingRuleInvokeArgs Empty => new GetGlobalForwardingRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetGlobalForwardingRuleResult
    {
        /// <summary>
        /// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By internal TCP/UDP load balancers, backend service-based network load balancers, and internal and external protocol forwarding. Set this field to true to allow packets addressed to any port or packets lacking destination port information (for example, UDP fragments after the first fragment) to be forwarded to the backends configured with this forwarding rule. The ports, port_range, and allPorts fields are mutually exclusive.
        /// </summary>
        public readonly bool AllPorts;
        /// <summary>
        /// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
        /// </summary>
        public readonly bool AllowGlobalAccess;
        /// <summary>
        /// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
        /// </summary>
        public readonly string BackendService;
        /// <summary>
        /// The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.
        /// </summary>
        public readonly string BaseForwardingRule;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// IP address for which this forwarding rule accepts traffic. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the referenced target or backendService. While creating a forwarding rule, specifying an IPAddress is required under the following circumstances: - When the target is set to targetGrpcProxy and validateForProxyless is set to true, the IPAddress should be set to 0.0.0.0. - When the target is a Private Service Connect Google APIs bundle, you must specify an IPAddress. Otherwise, you can optionally specify an IP address that references an existing static (reserved) IP address resource. When omitted, Google Cloud assigns an ephemeral IP address. Use one of the following formats to specify an IP address while creating a forwarding rule: * IP address number, as in `100.1.2.3` * IPv6 address range, as in `2600:1234::/96` * Full resource URL, as in https://www.googleapis.com/compute/v1/projects/ project_id/regions/region/addresses/address-name * Partial URL or by name, as in: - projects/project_id/regions/region/addresses/address-name - regions/region/addresses/address-name - global/addresses/address-name - address-name The forwarding rule's target or backendService, and in most cases, also the loadBalancingScheme, determine the type of IP address that you can use. For detailed information, see [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). When reading an IPAddress, the API always returns the IP address number.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The IP protocol to which this rule applies. For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP, ICMP and L3_DEFAULT. The valid IP protocols are different for different load balancing products as described in [Load balancing features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
        /// </summary>
        public readonly string IpProtocol;
        /// <summary>
        /// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6.
        /// </summary>
        public readonly string IpVersion;
        /// <summary>
        /// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
        /// </summary>
        public readonly bool IsMirroringCollector;
        /// <summary>
        /// Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Specifies the forwarding rule type. For more information about forwarding rules, refer to Forwarding rule concepts.
        /// </summary>
        public readonly string LoadBalancingScheme;
        /// <summary>
        /// Opaque filter criteria used by load balancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to load balancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule are not visible to those proxies. For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match. metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references. metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetadataFilterResponse> MetadataFilters;
        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. For Private Service Connect forwarding rules that forward traffic to Google APIs, the forwarding rule name must be a 1-20 characters string with lowercase letters and numbers and must start with a letter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// This field is not used for external load balancing. For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used. For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD. For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM. If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
        /// </summary>
        public readonly string NetworkTier;
        /// <summary>
        /// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
        /// </summary>
        public readonly bool NoAutomateDnsZone;
        /// <summary>
        /// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By backend service-based network load balancers, target pool-based network load balancers, internal proxy load balancers, external proxy load balancers, Traffic Director, external protocol forwarding, and Classic VPN. Some products have restrictions on what ports can be used. See port specifications for details. Only packets addressed to ports in the specified range will be forwarded to the backends configured with this forwarding rule. The ports, port_range, and allPorts fields are mutually exclusive. For external forwarding rules, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot have overlapping portRanges. For internal forwarding rules within the same VPC network, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot have overlapping portRanges. @pattern: \\d+(?:-\\d+)?
        /// </summary>
        public readonly string PortRange;
        /// <summary>
        /// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By internal TCP/UDP load balancers, backend service-based network load balancers, and internal protocol forwarding. You can specify a list of up to five ports by number, separated by commas. The ports can be contiguous or discontiguous. Only packets addressed to these ports will be forwarded to the backends configured with this forwarding rule. For external forwarding rules, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot share any values defined in ports. For internal forwarding rules within the same VPC network, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot share any values defined in ports. The ports, port_range, and allPorts fields are mutually exclusive. @pattern: \\d+(?:-\\d+)?
        /// </summary>
        public readonly ImmutableArray<string> Ports;
        /// <summary>
        /// The PSC connection id of the PSC Forwarding Rule.
        /// </summary>
        public readonly string PscConnectionId;
        public readonly string PscConnectionStatus;
        /// <summary>
        /// URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ForwardingRuleServiceDirectoryRegistrationResponse> ServiceDirectoryRegistrations;
        /// <summary>
        /// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing.
        /// </summary>
        public readonly string ServiceLabel;
        /// <summary>
        /// The internal fully qualified service name for this Forwarding Rule. This field is only used for internal load balancing.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each source_ip_range entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
        /// </summary>
        public readonly ImmutableArray<string> SourceIpRanges;
        /// <summary>
        /// This field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule, used in internal load balancing and network load balancing with IPv6. If the network specified is in auto subnet mode, this field is optional. However, a subnetwork must be specified if the network is in custom subnet mode or when creating external forwarding rule with IPv6.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must be in the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing resource. The forwarded traffic must be of a type appropriate to the target object. - For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). - For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle: - vpc-sc - APIs that support VPC Service Controls. - all-apis - All supported Google APIs. - For Private Service Connect forwarding rules that forward traffic to managed services, the target must be a service attachment. 
        /// </summary>
        public readonly string Target;

        [OutputConstructor]
        private GetGlobalForwardingRuleResult(
            bool allPorts,

            bool allowGlobalAccess,

            string backendService,

            string baseForwardingRule,

            string creationTimestamp,

            string description,

            string fingerprint,

            string ipAddress,

            string ipProtocol,

            string ipVersion,

            bool isMirroringCollector,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string loadBalancingScheme,

            ImmutableArray<Outputs.MetadataFilterResponse> metadataFilters,

            string name,

            string network,

            string networkTier,

            bool noAutomateDnsZone,

            string portRange,

            ImmutableArray<string> ports,

            string pscConnectionId,

            string pscConnectionStatus,

            string region,

            string selfLink,

            ImmutableArray<Outputs.ForwardingRuleServiceDirectoryRegistrationResponse> serviceDirectoryRegistrations,

            string serviceLabel,

            string serviceName,

            ImmutableArray<string> sourceIpRanges,

            string subnetwork,

            string target)
        {
            AllPorts = allPorts;
            AllowGlobalAccess = allowGlobalAccess;
            BackendService = backendService;
            BaseForwardingRule = baseForwardingRule;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            IpAddress = ipAddress;
            IpProtocol = ipProtocol;
            IpVersion = ipVersion;
            IsMirroringCollector = isMirroringCollector;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            LoadBalancingScheme = loadBalancingScheme;
            MetadataFilters = metadataFilters;
            Name = name;
            Network = network;
            NetworkTier = networkTier;
            NoAutomateDnsZone = noAutomateDnsZone;
            PortRange = portRange;
            Ports = ports;
            PscConnectionId = pscConnectionId;
            PscConnectionStatus = pscConnectionStatus;
            Region = region;
            SelfLink = selfLink;
            ServiceDirectoryRegistrations = serviceDirectoryRegistrations;
            ServiceLabel = serviceLabel;
            ServiceName = serviceName;
            SourceIpRanges = sourceIpRanges;
            Subnetwork = subnetwork;
            Target = target;
        }
    }
}
