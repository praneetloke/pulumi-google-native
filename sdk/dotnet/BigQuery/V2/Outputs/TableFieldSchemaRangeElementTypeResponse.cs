// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    /// <summary>
    /// Optional. The subtype of the RANGE, if the type of this field is RANGE. If the type is RANGE, this field is required. Possible values for the field element type of a RANGE include: - DATE - DATETIME - TIMESTAMP
    /// </summary>
    [OutputType]
    public sealed class TableFieldSchemaRangeElementTypeResponse
    {
        /// <summary>
        /// The field element type of a RANGE
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private TableFieldSchemaRangeElementTypeResponse(string type)
        {
            Type = type;
        }
    }
}