// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Outputs
{

    /// <summary>
    /// Specifies anti affinity group config for the VMware user cluster.
    /// </summary>
    [OutputType]
    public sealed class VmwareAAGConfigResponse
    {
        /// <summary>
        /// Spread nodes across at least three physical hosts (requires at least three hosts). Enabled by default.
        /// </summary>
        public readonly bool AagConfigDisabled;

        [OutputConstructor]
        private VmwareAAGConfigResponse(bool aagConfigDisabled)
        {
            AagConfigDisabled = aagConfigDisabled;
        }
    }
}
