// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1Beta1
{
    /// <summary>
    /// Creates a tag template. The user should enable the Data Catalog API in the project identified by the `parent` parameter (see [Data Catalog Resource Project](https://cloud.google.com/data-catalog/docs/concepts/resource-project) for more information).
    /// </summary>
    [GoogleNativeResourceType("google-native:datacatalog/v1beta1:TagTemplate")]
    public partial class TagTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// The display name for this template. Defaults to an empty string.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. This map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. Field IDs can contain letters (both uppercase and lowercase), numbers (0-9) and underscores (_). Field IDs must be at least 1 character long and at most 64 characters long. Field IDs must start with a letter or underscore.
        /// </summary>
        [Output("fields")]
        public Output<ImmutableDictionary<string, string>> Fields { get; private set; } = null!;

        /// <summary>
        /// The resource name of the tag template in URL format. Example: * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id} Note that this TagTemplate and its child resources may not actually be stored in the location in this name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a TagTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagTemplate(string name, TagTemplateArgs args, CustomResourceOptions? options = null)
            : base("google-native:datacatalog/v1beta1:TagTemplate", name, args ?? new TagTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TagTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datacatalog/v1beta1:TagTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TagTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TagTemplate(name, id, options);
        }
    }

    public sealed class TagTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name for this template. Defaults to an empty string.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("fields", required: true)]
        private InputMap<string>? _fields;

        /// <summary>
        /// Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. This map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. Field IDs can contain letters (both uppercase and lowercase), numbers (0-9) and underscores (_). Field IDs must be at least 1 character long and at most 64 characters long. Field IDs must start with a letter or underscore.
        /// </summary>
        public InputMap<string> Fields
        {
            get => _fields ?? (_fields = new InputMap<string>());
            set => _fields = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the tag template in URL format. Example: * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id} Note that this TagTemplate and its child resources may not actually be stored in the location in this name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("tagTemplateId", required: true)]
        public Input<string> TagTemplateId { get; set; } = null!;

        public TagTemplateArgs()
        {
        }
    }
}
