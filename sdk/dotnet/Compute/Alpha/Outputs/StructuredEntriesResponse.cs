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
    public sealed class StructuredEntriesResponse
    {
        /// <summary>
        /// Map of a partner metadata that belong to the same subdomain. It accepts any value including google.protobuf.Struct.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Entries;

        [OutputConstructor]
        private StructuredEntriesResponse(ImmutableDictionary<string, object> entries)
        {
            Entries = entries;
        }
    }
}
