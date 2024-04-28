// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Outputs
{

    /// <summary>
    /// VM instance status.
    /// </summary>
    [OutputType]
    public sealed class InstanceStatusResponse
    {
        /// <summary>
        /// The VM boot disk.
        /// </summary>
        public readonly Outputs.DiskResponse BootDisk;
        /// <summary>
        /// The Compute Engine machine type.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// The VM instance provisioning model.
        /// </summary>
        public readonly string ProvisioningModel;
        /// <summary>
        /// The max number of tasks can be assigned to this instance type.
        /// </summary>
        public readonly string TaskPack;

        [OutputConstructor]
        private InstanceStatusResponse(
            Outputs.DiskResponse bootDisk,

            string machineType,

            string provisioningModel,

            string taskPack)
        {
            BootDisk = bootDisk;
            MachineType = machineType;
            ProvisioningModel = provisioningModel;
            TaskPack = taskPack;
        }
    }
}
