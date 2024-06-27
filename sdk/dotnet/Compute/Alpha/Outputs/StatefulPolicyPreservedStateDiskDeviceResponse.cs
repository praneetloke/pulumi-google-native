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
    public sealed class StatefulPolicyPreservedStateDiskDeviceResponse
    {
        /// <summary>
        /// These stateful disks will never be deleted during autohealing, update or VM instance recreate operations. This flag is used to configure if the disk should be deleted after it is no longer used by the group, e.g. when the given instance or the whole group is deleted. Note: disks attached in READ_ONLY mode cannot be auto-deleted.
        /// </summary>
        public readonly string AutoDelete;

        [OutputConstructor]
        private StatefulPolicyPreservedStateDiskDeviceResponse(string autoDelete)
        {
            AutoDelete = autoDelete;
        }
    }
}