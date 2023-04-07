// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// A network interface resource attached to an instance.
    /// </summary>
    [OutputType]
    public sealed class NetworkInterfaceResponse
    {
        /// <summary>
        /// An array of configurations for this interface. Currently, only one access config, ONE_TO_ONE_NAT, is supported. If there are no accessConfigs specified, then this instance will have no external internet access.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessConfigResponse> AccessConfigs;
        /// <summary>
        /// An array of alias IP ranges for this network interface. You can only specify this field for network interfaces in VPC networks.
        /// </summary>
        public readonly ImmutableArray<Outputs.AliasIpRangeResponse> AliasIpRanges;
        /// <summary>
        /// Fingerprint hash of contents stored in this network interface. This field will be ignored when inserting an Instance or adding a NetworkInterface. An up-to-date fingerprint must be provided in order to update the NetworkInterface. The request will fail with error 400 Bad Request if the fingerprint is not provided, or 412 Precondition Failed if the fingerprint is out of date.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The prefix length of the primary internal IPv6 range.
        /// </summary>
        public readonly int InternalIpv6PrefixLength;
        /// <summary>
        /// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessConfigResponse> Ipv6AccessConfigs;
        /// <summary>
        /// One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork. Valid only if stackType is IPV4_IPV6.
        /// </summary>
        public readonly string Ipv6AccessType;
        /// <summary>
        /// An IPv6 internal network address for this network interface. To use a static internal IP address, it must be unused and in the same region as the instance's zone. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
        /// </summary>
        public readonly string Ipv6Address;
        /// <summary>
        /// Type of the resource. Always compute#networkInterface for network interfaces.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the network interface, which is generated by the server. For a VM, the network interface uses the nicN naming format. Where N is a value between 0 and 7. The default interface value is nic0.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL of the VPC network resource for this instance. When creating an instance, if neither the network nor the subnetwork is specified, the default network global/networks/default is used. If the selected project doesn't have the default network, you must specify a network or subnet. If the network is not specified but the subnetwork is specified, the network is inferred. If you specify this property, you can specify the network as a full or partial URL. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/global/networks/ network - projects/project/global/networks/network - global/networks/default 
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// The URL of the network attachment that this interface should connect to in the following format: projects/{project_number}/regions/{region_name}/networkAttachments/{network_attachment_name}.
        /// </summary>
        public readonly string NetworkAttachment;
        /// <summary>
        /// An IPv4 internal IP address to assign to the instance for this network interface. If not specified by the user, an unused internal IP is assigned by the system.
        /// </summary>
        public readonly string NetworkIP;
        /// <summary>
        /// The type of vNIC to be used on this interface. This may be gVNIC or VirtioNet.
        /// </summary>
        public readonly string NicType;
        /// <summary>
        /// The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It'll be empty if not specified by the users.
        /// </summary>
        public readonly int QueueCount;
        /// <summary>
        /// The stack type for this network interface. To assign only IPv4 addresses, use IPV4_ONLY. To assign both IPv4 and IPv6 addresses, use IPV4_IPV6. If not specified, IPV4_ONLY is used. This field can be both set at instance creation and update network interface operations.
        /// </summary>
        public readonly string StackType;
        /// <summary>
        /// The URL of the Subnetwork resource for this instance. If the network resource is in legacy mode, do not specify this field. If the network is in auto subnet mode, specifying the subnetwork is optional. If the network is in custom subnet mode, specifying the subnetwork is required. If you specify this field, you can specify the subnetwork as a full or partial URL. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/regions/region /subnetworks/subnetwork - regions/region/subnetworks/subnetwork 
        /// </summary>
        public readonly string Subnetwork;

        [OutputConstructor]
        private NetworkInterfaceResponse(
            ImmutableArray<Outputs.AccessConfigResponse> accessConfigs,

            ImmutableArray<Outputs.AliasIpRangeResponse> aliasIpRanges,

            string fingerprint,

            int internalIpv6PrefixLength,

            ImmutableArray<Outputs.AccessConfigResponse> ipv6AccessConfigs,

            string ipv6AccessType,

            string ipv6Address,

            string kind,

            string name,

            string network,

            string networkAttachment,

            string networkIP,

            string nicType,

            int queueCount,

            string stackType,

            string subnetwork)
        {
            AccessConfigs = accessConfigs;
            AliasIpRanges = aliasIpRanges;
            Fingerprint = fingerprint;
            InternalIpv6PrefixLength = internalIpv6PrefixLength;
            Ipv6AccessConfigs = ipv6AccessConfigs;
            Ipv6AccessType = ipv6AccessType;
            Ipv6Address = ipv6Address;
            Kind = kind;
            Name = name;
            Network = network;
            NetworkAttachment = networkAttachment;
            NetworkIP = networkIP;
            NicType = nicType;
            QueueCount = queueCount;
            StackType = stackType;
            Subnetwork = subnetwork;
        }
    }
}
