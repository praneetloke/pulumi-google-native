// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1
{
    /// <summary>
    /// Creates a new conversion workspace in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:datamigration/v1:ConversionWorkspace")]
    public partial class ConversionWorkspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The ID of the conversion workspace to create.
        /// </summary>
        [Output("conversionWorkspaceId")]
        public Output<string> ConversionWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the workspace resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The destination engine details.
        /// </summary>
        [Output("destination")]
        public Output<Outputs.DatabaseEngineInfoResponse> Destination { get; private set; } = null!;

        /// <summary>
        /// The display name for the workspace.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A generic list of settings for the workspace. The settings are database pair dependant and can indicate default behavior for the mapping rules engine or turn on or off specific features. Such examples can be: convert_foreign_key_to_interleave=true, skip_triggers=false, ignore_non_table_synonyms=true
        /// </summary>
        [Output("globalSettings")]
        public Output<ImmutableDictionary<string, string>> GlobalSettings { get; private set; } = null!;

        /// <summary>
        /// Whether the workspace has uncommitted changes (changes which were made after the workspace was committed).
        /// </summary>
        [Output("hasUncommittedChanges")]
        public Output<bool> HasUncommittedChanges { get; private set; } = null!;

        /// <summary>
        /// The latest commit ID.
        /// </summary>
        [Output("latestCommitId")]
        public Output<string> LatestCommitId { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the workspace was committed.
        /// </summary>
        [Output("latestCommitTime")]
        public Output<string> LatestCommitTime { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Full name of the workspace resource, in the form of: projects/{project}/locations/{location}/conversionWorkspaces/{conversion_workspace}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A unique ID used to identify the request. If the server receives two requests with the same ID, then the second request is ignored. It is recommended to always set this value to a UUID. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The source engine details.
        /// </summary>
        [Output("source")]
        public Output<Outputs.DatabaseEngineInfoResponse> Source { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the workspace resource was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ConversionWorkspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConversionWorkspace(string name, ConversionWorkspaceArgs args, CustomResourceOptions? options = null)
            : base("google-native:datamigration/v1:ConversionWorkspace", name, args ?? new ConversionWorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConversionWorkspace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datamigration/v1:ConversionWorkspace", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "conversionWorkspaceId",
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
        /// Get an existing ConversionWorkspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConversionWorkspace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConversionWorkspace(name, id, options);
        }
    }

    public sealed class ConversionWorkspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The ID of the conversion workspace to create.
        /// </summary>
        [Input("conversionWorkspaceId", required: true)]
        public Input<string> ConversionWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The destination engine details.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.DatabaseEngineInfoArgs> Destination { get; set; } = null!;

        /// <summary>
        /// The display name for the workspace.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("globalSettings")]
        private InputMap<string>? _globalSettings;

        /// <summary>
        /// A generic list of settings for the workspace. The settings are database pair dependant and can indicate default behavior for the mapping rules engine or turn on or off specific features. Such examples can be: convert_foreign_key_to_interleave=true, skip_triggers=false, ignore_non_table_synonyms=true
        /// </summary>
        public InputMap<string> GlobalSettings
        {
            get => _globalSettings ?? (_globalSettings = new InputMap<string>());
            set => _globalSettings = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Full name of the workspace resource, in the form of: projects/{project}/locations/{location}/conversionWorkspaces/{conversion_workspace}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A unique ID used to identify the request. If the server receives two requests with the same ID, then the second request is ignored. It is recommended to always set this value to a UUID. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The source engine details.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.DatabaseEngineInfoArgs> Source { get; set; } = null!;

        public ConversionWorkspaceArgs()
        {
        }
        public static new ConversionWorkspaceArgs Empty => new ConversionWorkspaceArgs();
    }
}
