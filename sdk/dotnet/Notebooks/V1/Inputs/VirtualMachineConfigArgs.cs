// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Notebooks.V1.Inputs
{

    /// <summary>
    /// The config settings for virtual machine.
    /// </summary>
    public sealed class VirtualMachineConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The Compute Engine accelerator configuration for this runtime.
        /// </summary>
        [Input("acceleratorConfig")]
        public Input<Inputs.RuntimeAcceleratorConfigArgs>? AcceleratorConfig { get; set; }

        [Input("containerImages")]
        private InputList<Inputs.ContainerImageArgs>? _containerImages;

        /// <summary>
        /// Optional. Use a list of container images to start the notebook instance.
        /// </summary>
        public InputList<Inputs.ContainerImageArgs> ContainerImages
        {
            get => _containerImages ?? (_containerImages = new InputList<Inputs.ContainerImageArgs>());
            set => _containerImages = value;
        }

        /// <summary>
        /// Required. Data disk option configuration settings.
        /// </summary>
        [Input("dataDisk")]
        public Input<Inputs.LocalDiskArgs>? DataDisk { get; set; }

        /// <summary>
        /// Optional. Encryption settings for virtual machine data disk.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.EncryptionConfigArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// Optional. If true, runtime will only have internal IP addresses. By default, runtimes are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each vm. This `internal_ip_only` restriction can only be enabled for subnetwork enabled networks, and all dependencies must be configured to be accessible without external IP addresses.
        /// </summary>
        [Input("internalIpOnly")]
        public Input<bool>? InternalIpOnly { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The labels to associate with this runtime. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Required. The Compute Engine machine type used for runtimes. Short name is valid. Examples: * `n1-standard-2` * `e2-standard-8`
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Optional. The Compute Engine metadata entries to add to virtual machine. (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork. If neither `network` nor `subnet` is specified, the "default" network of the project is used, if it exists. A full URL or partial URI. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/global/default` * `projects/[project_id]/regions/global/default` Runtimes are managed resources inside Google Infrastructure. Runtimes support the following network configurations: * Google Managed Network (Network &amp; subnet are empty) * Consumer Project VPC (network &amp; subnet are required). Requires configuring Private Service Access. * Shared VPC (network &amp; subnet are required). Requires configuring Private Service Access.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Optional. Shielded VM Instance configuration settings.
        /// </summary>
        [Input("shieldedInstanceConfig")]
        public Input<Inputs.RuntimeShieldedInstanceConfigArgs>? ShieldedInstanceConfig { get; set; }

        /// <summary>
        /// Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network. A full URL or partial URI are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0` * `projects/[project_id]/regions/us-east1/subnetworks/sub0`
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Optional. The Compute Engine tags to add to runtime (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public VirtualMachineConfigArgs()
        {
        }
    }
}
