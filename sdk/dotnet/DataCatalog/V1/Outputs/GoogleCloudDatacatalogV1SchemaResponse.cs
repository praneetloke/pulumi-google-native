// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDatacatalogV1SchemaResponse
    {
        /// <summary>
        /// The unified GoogleSQL-like schema of columns. The overall maximum number of columns and nested columns is 10,000. The maximum nested depth is 15 levels.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDatacatalogV1ColumnSchemaResponse> Columns;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1SchemaResponse(ImmutableArray<Outputs.GoogleCloudDatacatalogV1ColumnSchemaResponse> columns)
        {
            Columns = columns;
        }
    }
}