// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// An access configuration attached to an instance's network interface. Only one access config per instance is supported.
    /// </summary>
    [OutputType]
    public sealed class AccessConfigResponse
    {
        /// <summary>
        /// Applies to ipv6AccessConfigs only. The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. To use a static external IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an external IPv6 address from the instance's subnetwork.
        /// </summary>
        public readonly string ExternalIpv6;
        /// <summary>
        /// Applies to ipv6AccessConfigs only. The prefix length of the external IPv6 range.
        /// </summary>
        public readonly int ExternalIpv6PrefixLength;
        /// <summary>
        /// Type of the resource. Always compute#accessConfig for access configs.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of this access configuration. In accessConfigs (IPv4), the default and recommended name is External NAT, but you can use any arbitrary string, such as My external IP or Network Access. In ipv6AccessConfigs, the recommend name is External IPv6.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Applies to accessConfigs (IPv4) only. An external IP address associated with this instance. Specify an unused static external IP address available to the project or leave this field undefined to use an IP from a shared ephemeral IP address pool. If you specify a static external IP address, it must live in the same region as the zone of the instance.
        /// </summary>
        public readonly string NatIP;
        /// <summary>
        /// This signifies the networking tier used for configuring this access configuration and can only take the following values: PREMIUM, STANDARD. If an AccessConfig is specified without a valid external IP address, an ephemeral IP will be created with this networkTier. If an AccessConfig with a valid external IP address is specified, it must match that of the networkTier associated with the Address resource owning that IP.
        /// </summary>
        public readonly string NetworkTier;
        /// <summary>
        /// The DNS domain name for the public PTR record. You can set this field only if the `setPublicPtr` field is enabled in accessConfig. If this field is unspecified in ipv6AccessConfig, a default PTR record will be createc for first IP in associated external IPv6 range.
        /// </summary>
        public readonly string PublicPtrDomainName;
        /// <summary>
        /// Specifies whether a public DNS 'PTR' record should be created to map the external IP address of the instance to a DNS domain name. This field is not used in ipv6AccessConfig. A default PTR record will be created if the VM has external IPv6 range associated.
        /// </summary>
        public readonly bool SetPublicPtr;
        /// <summary>
        /// The type of configuration. In accessConfigs (IPv4), the default and only option is ONE_TO_ONE_NAT. In ipv6AccessConfigs, the default and only option is DIRECT_IPV6.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AccessConfigResponse(
            string externalIpv6,

            int externalIpv6PrefixLength,

            string kind,

            string name,

            string natIP,

            string networkTier,

            string publicPtrDomainName,

            bool setPublicPtr,

            string type)
        {
            ExternalIpv6 = externalIpv6;
            ExternalIpv6PrefixLength = externalIpv6PrefixLength;
            Kind = kind;
            Name = name;
            NatIP = natIP;
            NetworkTier = networkTier;
            PublicPtrDomainName = publicPtrDomainName;
            SetPublicPtr = setPublicPtr;
            Type = type;
        }
    }
}
