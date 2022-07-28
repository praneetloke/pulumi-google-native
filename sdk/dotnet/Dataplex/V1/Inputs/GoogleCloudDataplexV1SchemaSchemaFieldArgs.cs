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
    /// Represents a column field within a table schema.
    /// </summary>
    public sealed class GoogleCloudDataplexV1SchemaSchemaFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. User friendly field description. Must be less than or equal to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fields")]
        private InputList<Inputs.GoogleCloudDataplexV1SchemaSchemaFieldArgs>? _fields;

        /// <summary>
        /// Optional. Any nested field for complex types.
        /// </summary>
        public InputList<Inputs.GoogleCloudDataplexV1SchemaSchemaFieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.GoogleCloudDataplexV1SchemaSchemaFieldArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// Additional field semantics.
        /// </summary>
        [Input("mode", required: true)]
        public Input<Pulumi.GoogleNative.Dataplex.V1.GoogleCloudDataplexV1SchemaSchemaFieldMode> Mode { get; set; } = null!;

        /// <summary>
        /// The name of the field. Must contain only letters, numbers and underscores, with a maximum length of 767 characters, and must begin with a letter or underscore.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type of field.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.Dataplex.V1.GoogleCloudDataplexV1SchemaSchemaFieldType> Type { get; set; } = null!;

        public GoogleCloudDataplexV1SchemaSchemaFieldArgs()
        {
        }
        public static new GoogleCloudDataplexV1SchemaSchemaFieldArgs Empty => new GoogleCloudDataplexV1SchemaSchemaFieldArgs();
    }
}
