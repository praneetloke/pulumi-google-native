// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// Contains Properties set for the reservation.
    /// </summary>
    [OutputType]
    public sealed class AllocationResourceStatusSpecificSKUAllocationResponse
    {
        /// <summary>
        /// ID of the instance template used to populate reservation properties.
        /// </summary>
        public readonly string SourceInstanceTemplateId;

        [OutputConstructor]
        private AllocationResourceStatusSpecificSKUAllocationResponse(string sourceInstanceTemplateId)
        {
            SourceInstanceTemplateId = sourceInstanceTemplateId;
        }
    }
}
