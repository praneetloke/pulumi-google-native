// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class InstancePropertiesResponse
    {
        /// <summary>
        /// Controls for advanced machine-related behavior features. Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly Outputs.AdvancedMachineFeaturesResponse AdvancedMachineFeatures;
        /// <summary>
        /// Enables instances created based on these properties to send packets with source IP addresses other than their own and receive packets with destination IP addresses other than their own. If these instances will be used as an IP gateway or it will be set as the next-hop in a Route resource, specify true. If unsure, leave this set to false. See the Enable IP forwarding documentation for more information.
        /// </summary>
        public readonly bool CanIpForward;
        /// <summary>
        /// Specifies the Confidential Instance options. Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly Outputs.ConfidentialInstanceConfigResponse ConfidentialInstanceConfig;
        /// <summary>
        /// An optional text description for the instances that are created from these properties.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// An array of disks that are associated with the instances that are created from these properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.AttachedDiskResponse> Disks;
        /// <summary>
        /// Display Device properties to enable support for remote display products like: Teradici, VNC and TeamViewer Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly Outputs.DisplayDeviceResponse DisplayDevice;
        /// <summary>
        /// A list of guest accelerator cards' type and count to use for instances created from these properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.AcceleratorConfigResponse> GuestAccelerators;
        /// <summary>
        /// KeyRevocationActionType of the instance. Supported options are "STOP" and "NONE". The default value is "NONE" if it is not specified.
        /// </summary>
        public readonly string KeyRevocationActionType;
        /// <summary>
        /// Labels to apply to instances that are created from these properties.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The machine type to use for instances that are created from these properties.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// The metadata key/value pairs to assign to instances that are created from these properties. These pairs can consist of custom metadata or predefined keys. See Project and instance metadata for more information.
        /// </summary>
        public readonly Outputs.MetadataResponse Metadata;
        /// <summary>
        /// Minimum cpu/platform to be used by instances. The instance may be scheduled on the specified or newer cpu/platform. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge". For more information, read Specifying a Minimum CPU Platform.
        /// </summary>
        public readonly string MinCpuPlatform;
        /// <summary>
        /// An array of network access configurations for this interface.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceResponse> NetworkInterfaces;
        /// <summary>
        /// Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly Outputs.NetworkPerformanceConfigResponse NetworkPerformanceConfig;
        /// <summary>
        /// PostKeyRevocationActionType of the instance.
        /// </summary>
        public readonly string PostKeyRevocationActionType;
        /// <summary>
        /// The private IPv6 google access type for VMs. If not specified, use INHERIT_FROM_SUBNETWORK as default. Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly string PrivateIpv6GoogleAccess;
        /// <summary>
        /// Specifies the reservations that instances can consume from. Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly Outputs.ReservationAffinityResponse ReservationAffinity;
        /// <summary>
        /// Resource manager tags to be bound to the instance. Tag keys and values have the same definition as resource manager tags. Keys must be in the format `tagKeys/{tag_key_id}`, and values are in the format `tagValues/456`. The field is ignored (both PUT &amp; PATCH) when empty.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ResourceManagerTags;
        /// <summary>
        /// Resource policies (names, not URLs) applied to instances created from these properties. Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly ImmutableArray<string> ResourcePolicies;
        /// <summary>
        /// Specifies the scheduling options for the instances that are created from these properties.
        /// </summary>
        public readonly Outputs.SchedulingResponse Scheduling;
        /// <summary>
        /// [Input Only] Secure tags to apply to this instance. Maximum number of secure tags allowed is 50. Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly ImmutableArray<string> SecureTags;
        /// <summary>
        /// A list of service accounts with specified scopes. Access tokens for these service accounts are available to the instances that are created from these properties. Use metadata queries to obtain the access tokens for these instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceAccountResponse> ServiceAccounts;
        /// <summary>
        /// Note that for MachineImage, this is not supported yet.
        /// </summary>
        public readonly Outputs.ShieldedInstanceConfigResponse ShieldedInstanceConfig;
        /// <summary>
        /// Specifies the Shielded VM options for the instances that are created from these properties.
        /// </summary>
        public readonly Outputs.ShieldedVmConfigResponse ShieldedVmConfig;
        /// <summary>
        /// A list of tags to apply to the instances that are created from these properties. The tags identify valid sources or targets for network firewalls. The setTags method can modify this list of tags. Each tag within the list must comply with RFC1035.
        /// </summary>
        public readonly Outputs.TagsResponse Tags;

        [OutputConstructor]
        private InstancePropertiesResponse(
            Outputs.AdvancedMachineFeaturesResponse advancedMachineFeatures,

            bool canIpForward,

            Outputs.ConfidentialInstanceConfigResponse confidentialInstanceConfig,

            string description,

            ImmutableArray<Outputs.AttachedDiskResponse> disks,

            Outputs.DisplayDeviceResponse displayDevice,

            ImmutableArray<Outputs.AcceleratorConfigResponse> guestAccelerators,

            string keyRevocationActionType,

            ImmutableDictionary<string, string> labels,

            string machineType,

            Outputs.MetadataResponse metadata,

            string minCpuPlatform,

            ImmutableArray<Outputs.NetworkInterfaceResponse> networkInterfaces,

            Outputs.NetworkPerformanceConfigResponse networkPerformanceConfig,

            string postKeyRevocationActionType,

            string privateIpv6GoogleAccess,

            Outputs.ReservationAffinityResponse reservationAffinity,

            ImmutableDictionary<string, string> resourceManagerTags,

            ImmutableArray<string> resourcePolicies,

            Outputs.SchedulingResponse scheduling,

            ImmutableArray<string> secureTags,

            ImmutableArray<Outputs.ServiceAccountResponse> serviceAccounts,

            Outputs.ShieldedInstanceConfigResponse shieldedInstanceConfig,

            Outputs.ShieldedVmConfigResponse shieldedVmConfig,

            Outputs.TagsResponse tags)
        {
            AdvancedMachineFeatures = advancedMachineFeatures;
            CanIpForward = canIpForward;
            ConfidentialInstanceConfig = confidentialInstanceConfig;
            Description = description;
            Disks = disks;
            DisplayDevice = displayDevice;
            GuestAccelerators = guestAccelerators;
            KeyRevocationActionType = keyRevocationActionType;
            Labels = labels;
            MachineType = machineType;
            Metadata = metadata;
            MinCpuPlatform = minCpuPlatform;
            NetworkInterfaces = networkInterfaces;
            NetworkPerformanceConfig = networkPerformanceConfig;
            PostKeyRevocationActionType = postKeyRevocationActionType;
            PrivateIpv6GoogleAccess = privateIpv6GoogleAccess;
            ReservationAffinity = reservationAffinity;
            ResourceManagerTags = resourceManagerTags;
            ResourcePolicies = resourcePolicies;
            Scheduling = scheduling;
            SecureTags = secureTags;
            ServiceAccounts = serviceAccounts;
            ShieldedInstanceConfig = shieldedInstanceConfig;
            ShieldedVmConfig = shieldedVmConfig;
            Tags = tags;
        }
    }
}
