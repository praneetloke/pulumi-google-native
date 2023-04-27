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
    /// Create a content.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataplex/v1:Contentitem")]
    public partial class Contentitem : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Content creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Content data in string format.
        /// </summary>
        [Output("dataText")]
        public Output<string> DataText { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the content.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. User defined labels for the content.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("lakeId")]
        public Output<string> LakeId { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The relative resource name of the content, of the form: projects/{project_id}/locations/{location_id}/lakes/{lake_id}/content/{content_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Notebook related configurations.
        /// </summary>
        [Output("notebook")]
        public Output<Outputs.GoogleCloudDataplexV1ContentNotebookResponse> Notebook { get; private set; } = null!;

        /// <summary>
        /// The path for the Content file, represented as directory structure. Unique within a lake. Limited to alphanumerics, hyphens, underscores, dots and slashes.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Sql Script related configurations.
        /// </summary>
        [Output("sqlScript")]
        public Output<Outputs.GoogleCloudDataplexV1ContentSqlScriptResponse> SqlScript { get; private set; } = null!;

        /// <summary>
        /// System generated globally unique ID for the content. This ID will be different if the content is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time when the content was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Contentitem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Contentitem(string name, ContentitemArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Contentitem", name, args ?? new ContentitemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Contentitem(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Contentitem", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "lakeId",
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
        /// Get an existing Contentitem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Contentitem Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Contentitem(name, id, options);
        }
    }

    public sealed class ContentitemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content data in string format.
        /// </summary>
        [Input("dataText", required: true)]
        public Input<string> DataText { get; set; } = null!;

        /// <summary>
        /// Optional. Description of the content.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User defined labels for the content.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("lakeId", required: true)]
        public Input<string> LakeId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Notebook related configurations.
        /// </summary>
        [Input("notebook")]
        public Input<Inputs.GoogleCloudDataplexV1ContentNotebookArgs>? Notebook { get; set; }

        /// <summary>
        /// The path for the Content file, represented as directory structure. Unique within a lake. Limited to alphanumerics, hyphens, underscores, dots and slashes.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Sql Script related configurations.
        /// </summary>
        [Input("sqlScript")]
        public Input<Inputs.GoogleCloudDataplexV1ContentSqlScriptArgs>? SqlScript { get; set; }

        public ContentitemArgs()
        {
        }
        public static new ContentitemArgs Empty => new ContentitemArgs();
    }
}
