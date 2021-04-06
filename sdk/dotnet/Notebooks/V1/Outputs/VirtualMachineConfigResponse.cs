// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Notebooks.V1.Outputs
{

    [OutputType]
    public sealed class VirtualMachineConfigResponse
    {
        /// <summary>
        /// Optional. The Compute Engine accelerator configuration for this runtime.
        /// </summary>
        public readonly Outputs.RuntimeAcceleratorConfigResponse AcceleratorConfig;
        /// <summary>
        /// Optional. Use a list of container images to start the notebook instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerImageResponse> ContainerImages;
        /// <summary>
        /// Required. Data disk option configuration settings.
        /// </summary>
        public readonly Outputs.LocalDiskResponse DataDisk;
        /// <summary>
        /// Optional. Encryption settings for virtual machine data disk.
        /// </summary>
        public readonly Outputs.EncryptionConfigResponse EncryptionConfig;
        /// <summary>
        /// The Compute Engine guest attributes. (see [Project and instance guest attributes](https://cloud.google.com/compute/docs/storing-retrieving-metadata#guest_attributes)).
        /// </summary>
        public readonly ImmutableDictionary<string, string> GuestAttributes;
        /// <summary>
        /// Optional. If true, runtime will only have internal IP addresses. By default, runtimes are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each vm. This `internal_ip_only` restriction can only be enabled for subnetwork enabled networks, and all dependencies must be configured to be accessible without external IP addresses.
        /// </summary>
        public readonly bool InternalIpOnly;
        /// <summary>
        /// Optional. The labels to associate with this runtime. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Required. The Compute Engine machine type used for runtimes. Short name is valid. Examples: * `n1-standard-2` * `e2-standard-8`
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Optional. The Compute Engine metadata entries to add to virtual machine. (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork. If neither `network` nor `subnet` is specified, the "default" network of the project is used, if it exists. A full URL or partial URI. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/global/default` * `projects/[project_id]/regions/global/default` Runtimes are managed resources inside Google Infrastructure. Runtimes support the following network configurations: * Google Managed Network (Network &amp; subnet are empty) * Consumer Project VPC (network &amp; subnet are required). Requires configuring Private Service Access. * Shared VPC (network &amp; subnet are required). Requires configuring Private Service Access.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Optional. Shielded VM Instance configuration settings.
        /// </summary>
        public readonly Outputs.RuntimeShieldedInstanceConfigResponse ShieldedInstanceConfig;
        /// <summary>
        /// Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network. A full URL or partial URI are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0` * `projects/[project_id]/regions/us-east1/subnetworks/sub0`
        /// </summary>
        public readonly string Subnet;
        /// <summary>
        /// Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The zone where the virtual machine is located. If using regional request, the notebooks service will pick a location in the corresponding runtime region. On a get request, zone will always be present. Example: * `us-central1-b`
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private VirtualMachineConfigResponse(
            Outputs.RuntimeAcceleratorConfigResponse acceleratorConfig,

            ImmutableArray<Outputs.ContainerImageResponse> containerImages,

            Outputs.LocalDiskResponse dataDisk,

            Outputs.EncryptionConfigResponse encryptionConfig,

            ImmutableDictionary<string, string> guestAttributes,

            bool internalIpOnly,

            ImmutableDictionary<string, string> labels,

            string machineType,

            ImmutableDictionary<string, string> metadata,

            string network,

            Outputs.RuntimeShieldedInstanceConfigResponse shieldedInstanceConfig,

            string subnet,

            ImmutableArray<string> tags,

            string zone)
        {
            AcceleratorConfig = acceleratorConfig;
            ContainerImages = containerImages;
            DataDisk = dataDisk;
            EncryptionConfig = encryptionConfig;
            GuestAttributes = guestAttributes;
            InternalIpOnly = internalIpOnly;
            Labels = labels;
            MachineType = machineType;
            Metadata = metadata;
            Network = network;
            ShieldedInstanceConfig = shieldedInstanceConfig;
            Subnet = subnet;
            Tags = tags;
            Zone = zone;
        }
    }
}
