// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1.Outputs
{

    /// <summary>
    /// A VMWare Engine Node
    /// </summary>
    [OutputType]
    public sealed class ReportSummaryVmwareNodeResponse
    {
        /// <summary>
        /// Code to identify VMware Engine node series, e.g. "ve1-standard-72". Based on the displayName of cloud.google.com/vmware-engine/docs/reference/rest/v1/projects.locations.nodeTypes
        /// </summary>
        public readonly string Code;

        [OutputConstructor]
        private ReportSummaryVmwareNodeResponse(string code)
        {
            Code = code;
        }
    }
}