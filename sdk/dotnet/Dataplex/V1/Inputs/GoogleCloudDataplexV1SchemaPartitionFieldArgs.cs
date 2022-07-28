// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// Represents a key field within the entity's partition structure. You could have up to 20 partition fields, but only the first 10 partitions have the filtering ability due to performance consideration. Note: Partition fields are immutable.
    /// </summary>
    public sealed class GoogleCloudDataplexV1SchemaPartitionFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Partition field name must consist of letters, numbers, and underscores only, with a maximum of length of 256 characters, and must begin with a letter or underscore..
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Immutable. The type of field.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Dataplex.V1.GoogleCloudDataplexV1SchemaPartitionFieldType> Type { get; set; } = null!;

        public GoogleCloudDataplexV1SchemaPartitionFieldArgs()
        {
        }
        public static new GoogleCloudDataplexV1SchemaPartitionFieldArgs Empty => new GoogleCloudDataplexV1SchemaPartitionFieldArgs();
    }
}
