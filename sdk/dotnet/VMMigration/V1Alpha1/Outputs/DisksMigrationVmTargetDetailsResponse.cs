// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1.Outputs
{

    /// <summary>
    /// Details for the VM created VM as part of disks migration.
    /// </summary>
    [OutputType]
    public sealed class DisksMigrationVmTargetDetailsResponse
    {
        /// <summary>
        /// The URI of the Compute Engine VM.
        /// </summary>
        public readonly string VmUri;

        [OutputConstructor]
        private DisksMigrationVmTargetDetailsResponse(string vmUri)
        {
            VmUri = vmUri;
        }
    }
}