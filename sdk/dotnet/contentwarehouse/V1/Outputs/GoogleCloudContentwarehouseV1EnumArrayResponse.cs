// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.contentwarehouse.V1.Outputs
{

    /// <summary>
    /// Enum values.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContentwarehouseV1EnumArrayResponse
    {
        /// <summary>
        /// List of enum values.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GoogleCloudContentwarehouseV1EnumArrayResponse(ImmutableArray<string> values)
        {
            Values = values;
        }
    }
}