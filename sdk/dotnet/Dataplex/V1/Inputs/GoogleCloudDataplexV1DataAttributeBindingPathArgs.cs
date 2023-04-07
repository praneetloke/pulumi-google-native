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
    /// Represents a subresource of the given resource, and associated bindings with it. Currently supported subresources are column and partition schema fields within a table.
    /// </summary>
    public sealed class GoogleCloudDataplexV1DataAttributeBindingPathArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<string>? _attributes;

        /// <summary>
        /// Optional. List of attributes to be associated with the path of the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
        /// </summary>
        public InputList<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<string>());
            set => _attributes = value;
        }

        /// <summary>
        /// The name identifier of the path. Nested columns should be of the form: 'address.city'.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GoogleCloudDataplexV1DataAttributeBindingPathArgs()
        {
        }
        public static new GoogleCloudDataplexV1DataAttributeBindingPathArgs Empty => new GoogleCloudDataplexV1DataAttributeBindingPathArgs();
    }
}
