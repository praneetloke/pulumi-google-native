// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the size of the disk in base-2 GB.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<string>? DiskSizeGb { get; set; }

        /// <summary>
        /// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. For performance characteristics of SCSI over NVMe, see Local SSD performance.
        /// </summary>
        [Input("interface")]
        public Input<Pulumi.GoogleNative.Compute.Beta.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskInterface>? Interface { get; set; }

        public AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskArgs()
        {
        }
        public static new AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskArgs Empty => new AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskArgs();
    }
}
