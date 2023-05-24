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
    /// Creates a new EndpointAttachment in a given project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:connectors/v1:EndpointAttachment")]
    public partial class EndpointAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Created time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Required. Identifier to assign to the EndpointAttachment. Must be unique within scope of the parent resource.
        /// </summary>
        [Output("endpointAttachmentId")]
        public Output<string> EndpointAttachmentId { get; private set; } = null!;

        /// <summary>
        /// The Private Service Connect connection endpoint ip
        /// </summary>
        [Output("endpointIp")]
        public Output<string> EndpointIp { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name of the Endpoint Attachment. Format: projects/{project}/locations/{location}/endpointAttachments/{endpoint_attachment}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The path of the service attachment
        /// </summary>
        [Output("serviceAttachment")]
        public Output<string> ServiceAttachment { get; private set; } = null!;

        /// <summary>
        /// Updated time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointAttachment(string name, EndpointAttachmentArgs args, CustomResourceOptions? options = null)
            : base("google-native:connectors/v1:EndpointAttachment", name, args ?? new EndpointAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:connectors/v1:EndpointAttachment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "endpointAttachmentId",
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
        /// Get an existing EndpointAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EndpointAttachment(name, id, options);
        }
    }

    public sealed class EndpointAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. Identifier to assign to the EndpointAttachment. Must be unique within scope of the parent resource.
        /// </summary>
        [Input("endpointAttachmentId", required: true)]
        public Input<string> EndpointAttachmentId { get; set; } = null!;

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

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The path of the service attachment
        /// </summary>
        [Input("serviceAttachment", required: true)]
        public Input<string> ServiceAttachment { get; set; } = null!;

        public EndpointAttachmentArgs()
        {
        }
        public static new EndpointAttachmentArgs Empty => new EndpointAttachmentArgs();
    }
}