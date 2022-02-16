// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Outputs
{

    /// <summary>
    /// Represents a column field within a table schema.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1SchemaSchemaFieldResponse
    {
        /// <summary>
        /// Optional. User friendly field description. Must be less than or equal to 1024 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Any nested field for complex types.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDataplexV1SchemaSchemaFieldResponse> Fields;
        /// <summary>
        /// Additional field semantics.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The name of the field. The maximum length is 767 characters. The name must begins with a letter and not contains : and ..
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of field.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GoogleCloudDataplexV1SchemaSchemaFieldResponse(
            string description,

            ImmutableArray<Outputs.GoogleCloudDataplexV1SchemaSchemaFieldResponse> fields,

            string mode,

            string name,

            string type)
        {
            Description = description;
            Fields = fields;
            Mode = mode;
            Name = name;
            Type = type;
        }
    }
}