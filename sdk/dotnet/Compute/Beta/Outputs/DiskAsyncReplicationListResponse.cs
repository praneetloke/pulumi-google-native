// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class DiskAsyncReplicationListResponse
    {
        public readonly Outputs.DiskAsyncReplicationResponse AsyncReplicationDisk;

        [OutputConstructor]
        private DiskAsyncReplicationListResponse(Outputs.DiskAsyncReplicationResponse asyncReplicationDisk)
        {
            AsyncReplicationDisk = asyncReplicationDisk;
        }
    }
}
