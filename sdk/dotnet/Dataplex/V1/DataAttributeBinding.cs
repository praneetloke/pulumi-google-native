// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1
{
    /// <summary>
    /// Create a DataAttributeBinding resource.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataplex/v1:DataAttributeBinding")]
    public partial class DataAttributeBinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. List of attributes to be associated with the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableArray<string>> Attributes { get; private set; } = null!;

        /// <summary>
        /// The time when the DataAttributeBinding was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Required. DataAttributeBinding identifier. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the Location.
        /// </summary>
        [Output("dataAttributeBindingId")]
        public Output<string> DataAttributeBindingId { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the DataAttributeBinding.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Etags must be used when calling the DeleteDataAttributeBinding and the UpdateDataAttributeBinding method.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Optional. User-defined labels for the DataAttributeBinding.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The relative resource name of the Data Attribute Binding, of the form: projects/{project_number}/locations/{location}/dataAttributeBindings/{data_attribute_binding_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. The list of paths for items within the associated resource (eg. columns and partitions within a table) along with attribute bindings.
        /// </summary>
        [Output("paths")]
        public Output<ImmutableArray<Outputs.GoogleCloudDataplexV1DataAttributeBindingPathResponse>> Paths { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. Immutable. The resource name of the resource that is associated to attributes. Presently, only entity resource is supported in the form: projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id} Must belong in the same project and region as the attribute binding, and there can only exist one active binding for a resource.
        /// </summary>
        [Output("resource")]
        public Output<string> Resource { get; private set; } = null!;

        /// <summary>
        /// System generated globally unique ID for the DataAttributeBinding. This ID will be different if the DataAttributeBinding is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time when the DataAttributeBinding was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Only validate the request, but do not perform mutations. The default is false.
        /// </summary>
        [Output("validateOnly")]
        public Output<bool?> ValidateOnly { get; private set; } = null!;


        /// <summary>
        /// Create a DataAttributeBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataAttributeBinding(string name, DataAttributeBindingArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:DataAttributeBinding", name, args ?? new DataAttributeBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataAttributeBinding(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:DataAttributeBinding", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "dataAttributeBindingId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataAttributeBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataAttributeBinding Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataAttributeBinding(name, id, options);
        }
    }

    public sealed class DataAttributeBindingArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<string>? _attributes;

        /// <summary>
        /// Optional. List of attributes to be associated with the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
        /// </summary>
        public InputList<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<string>());
            set => _attributes = value;
        }

        /// <summary>
        /// Required. DataAttributeBinding identifier. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the Location.
        /// </summary>
        [Input("dataAttributeBindingId", required: true)]
        public Input<string> DataAttributeBindingId { get; set; } = null!;

        /// <summary>
        /// Optional. Description of the DataAttributeBinding.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Etags must be used when calling the DeleteDataAttributeBinding and the UpdateDataAttributeBinding method.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User-defined labels for the DataAttributeBinding.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("paths")]
        private InputList<Inputs.GoogleCloudDataplexV1DataAttributeBindingPathArgs>? _paths;

        /// <summary>
        /// Optional. The list of paths for items within the associated resource (eg. columns and partitions within a table) along with attribute bindings.
        /// </summary>
        public InputList<Inputs.GoogleCloudDataplexV1DataAttributeBindingPathArgs> Paths
        {
            get => _paths ?? (_paths = new InputList<Inputs.GoogleCloudDataplexV1DataAttributeBindingPathArgs>());
            set => _paths = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. Immutable. The resource name of the resource that is associated to attributes. Presently, only entity resource is supported in the form: projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id} Must belong in the same project and region as the attribute binding, and there can only exist one active binding for a resource.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        /// <summary>
        /// Optional. Only validate the request, but do not perform mutations. The default is false.
        /// </summary>
        [Input("validateOnly")]
        public Input<bool>? ValidateOnly { get; set; }

        public DataAttributeBindingArgs()
        {
        }
        public static new DataAttributeBindingArgs Empty => new DataAttributeBindingArgs();
    }
}
