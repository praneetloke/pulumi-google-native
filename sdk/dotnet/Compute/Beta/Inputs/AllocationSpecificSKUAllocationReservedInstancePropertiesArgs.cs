// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Properties of the SKU instances being reserved. Next ID: 9
    /// </summary>
    public sealed class AllocationSpecificSKUAllocationReservedInstancePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("guestAccelerators")]
        private InputList<Inputs.AcceleratorConfigArgs>? _guestAccelerators;

        /// <summary>
        /// Specifies accelerator type and count.
        /// </summary>
        public InputList<Inputs.AcceleratorConfigArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<Inputs.AcceleratorConfigArgs>());
            set => _guestAccelerators = value;
        }

        [Input("localSsds")]
        private InputList<Inputs.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskArgs>? _localSsds;

        /// <summary>
        /// Specifies amount of local ssd to reserve with each instance. The type of disk is local-ssd.
        /// </summary>
        public InputList<Inputs.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskArgs> LocalSsds
        {
            get => _localSsds ?? (_localSsds = new InputList<Inputs.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDiskArgs>());
            set => _localSsds = value;
        }

        /// <summary>
        /// An opaque location hint used to place the allocation close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        [Input("locationHint")]
        public Input<string>? LocationHint { get; set; }

        /// <summary>
        /// Specifies type of machine (name only) which has fixed number of vCPUs and fixed amount of memory. This also includes specifying custom machine type following custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY pattern.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// Specifies the number of hours after reservation creation where instances using the reservation won't be scheduled for maintenance.
        /// </summary>
        [Input("maintenanceFreezeDurationHours")]
        public Input<int>? MaintenanceFreezeDurationHours { get; set; }

        /// <summary>
        /// For more information about maintenance intervals, see Setting maintenance intervals.
        /// </summary>
        [Input("maintenanceInterval")]
        public Input<Pulumi.GoogleNative.Compute.Beta.AllocationSpecificSKUAllocationReservedInstancePropertiesMaintenanceInterval>? MaintenanceInterval { get; set; }

        /// <summary>
        /// Minimum cpu platform the reservation.
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        public AllocationSpecificSKUAllocationReservedInstancePropertiesArgs()
        {
        }
        public static new AllocationSpecificSKUAllocationReservedInstancePropertiesArgs Empty => new AllocationSpecificSKUAllocationReservedInstancePropertiesArgs();
    }
}
