// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1
{
    /// <summary>
    /// Creates a new CustomConnector in a given project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:connectors/v1:CustomConnector")]
    public partial class CustomConnector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Created time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Required. Identifier to assign to the CreateCustomConnector. Must be unique within scope of the parent resource.
        /// </summary>
        [Output("customConnectorId")]
        public Output<string> CustomConnectorId { get; private set; } = null!;

        /// <summary>
        /// Type of the custom connector.
        /// </summary>
        [Output("customConnectorType")]
        public Output<string> CustomConnectorType { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Optional. Logo of the resource.
        /// </summary>
        [Output("logo")]
        public Output<string> Logo { get; private set; } = null!;

        /// <summary>
        /// Identifier. Resource name of the CustomConnector. Format: projects/{project}/locations/{location}/customConnectors/{connector}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Updated time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CustomConnector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomConnector(string name, CustomConnectorArgs args, CustomResourceOptions? options = null)
            : base("google-native:connectors/v1:CustomConnector", name, args ?? new CustomConnectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomConnector(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:connectors/v1:CustomConnector", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "customConnectorId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomConnector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomConnector Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomConnector(name, id, options);
        }
    }

    public sealed class CustomConnectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Identifier to assign to the CreateCustomConnector. Must be unique within scope of the parent resource.
        /// </summary>
        [Input("customConnectorId", required: true)]
        public Input<string> CustomConnectorId { get; set; } = null!;

        /// <summary>
        /// Type of the custom connector.
        /// </summary>
        [Input("customConnectorType", required: true)]
        public Input<Pulumi.GoogleNative.Connectors.V1.CustomConnectorCustomConnectorType> CustomConnectorType { get; set; } = null!;

        /// <summary>
        /// Optional. Description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Optional. Logo of the resource.
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// Identifier. Resource name of the CustomConnector. Format: projects/{project}/locations/{location}/customConnectors/{connector}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public CustomConnectorArgs()
        {
        }
        public static new CustomConnectorArgs Empty => new CustomConnectorArgs();
    }
}