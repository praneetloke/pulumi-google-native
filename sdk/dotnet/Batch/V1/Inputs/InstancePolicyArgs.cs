// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Inputs
{

    /// <summary>
    /// InstancePolicy describes an instance type and resources attached to each VM created by this InstancePolicy.
    /// </summary>
    public sealed class InstancePolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("accelerators")]
        private InputList<Inputs.AcceleratorArgs>? _accelerators;

        /// <summary>
        /// The accelerators attached to each VM instance.
        /// </summary>
        public InputList<Inputs.AcceleratorArgs> Accelerators
        {
            get => _accelerators ?? (_accelerators = new InputList<Inputs.AcceleratorArgs>());
            set => _accelerators = value;
        }

        [Input("disks")]
        private InputList<Inputs.AttachedDiskArgs>? _disks;

        /// <summary>
        /// Non-boot disks to be attached for each VM created by this InstancePolicy. New disks will be deleted when the VM is deleted.
        /// </summary>
        public InputList<Inputs.AttachedDiskArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.AttachedDiskArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// The Compute Engine machine type.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// The minimum CPU platform. See `https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform`. Not yet implemented.
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// The provisioning model.
        /// </summary>
        [Input("provisioningModel")]
        public Input<Pulumi.GoogleNative.Batch.V1.InstancePolicyProvisioningModel>? ProvisioningModel { get; set; }

        public InstancePolicyArgs()
        {
        }
        public static new InstancePolicyArgs Empty => new InstancePolicyArgs();
    }
}
