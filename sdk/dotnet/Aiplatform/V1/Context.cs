// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1
{
    /// <summary>
    /// Creates a Context associated with a MetadataStore.
    /// </summary>
    [GoogleNativeResourceType("google-native:aiplatform/v1:Context")]
    public partial class Context : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The {context} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}/contexts/{context}`. If not provided, the Context's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all Contexts in the parent MetadataStore. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting Context.)
        /// </summary>
        [Output("contextId")]
        public Output<string?> ContextId { get; private set; } = null!;

        /// <summary>
        /// Timestamp when this Context was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the Context
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// User provided display name of the Context. May be up to 128 Unicode characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// An eTag used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The labels with user-defined metadata to organize your Contexts. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Context (System labels are excluded).
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Properties of the Context. Top level metadata keys' heading and trailing spaces will be trimmed. The size of this field should not exceed 200KB.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        [Output("metadataStoreId")]
        public Output<string> MetadataStoreId { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the Context.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of resource names of Contexts that are parents of this Context. A Context may have at most 10 parent_contexts.
        /// </summary>
        [Output("parentContexts")]
        public Output<ImmutableArray<string>> ParentContexts { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The title of the schema describing the metadata. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
        /// </summary>
        [Output("schemaTitle")]
        public Output<string> SchemaTitle { get; private set; } = null!;

        /// <summary>
        /// The version of the schema in schema_name to use. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
        /// </summary>
        [Output("schemaVersion")]
        public Output<string> SchemaVersion { get; private set; } = null!;

        /// <summary>
        /// Timestamp when this Context was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Context resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Context(string name, ContextArgs args, CustomResourceOptions? options = null)
            : base("google-native:aiplatform/v1:Context", name, args ?? new ContextArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Context(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:aiplatform/v1:Context", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "metadataStoreId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Context resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Context Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Context(name, id, options);
        }
    }

    public sealed class ContextArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The {context} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}/contexts/{context}`. If not provided, the Context's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all Contexts in the parent MetadataStore. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting Context.)
        /// </summary>
        [Input("contextId")]
        public Input<string>? ContextId { get; set; }

        /// <summary>
        /// Description of the Context
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User provided display name of the Context. May be up to 128 Unicode characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// An eTag used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your Contexts. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Context (System labels are excluded).
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Properties of the Context. Top level metadata keys' heading and trailing spaces will be trimmed. The size of this field should not exceed 200KB.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataStoreId", required: true)]
        public Input<string> MetadataStoreId { get; set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the Context.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The title of the schema describing the metadata. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
        /// </summary>
        [Input("schemaTitle")]
        public Input<string>? SchemaTitle { get; set; }

        /// <summary>
        /// The version of the schema in schema_name to use. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
        /// </summary>
        [Input("schemaVersion")]
        public Input<string>? SchemaVersion { get; set; }

        public ContextArgs()
        {
        }
        public static new ContextArgs Empty => new ContextArgs();
    }
}
