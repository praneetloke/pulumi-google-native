// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1Alpha1
{
    /// <summary>
    /// Creates a new source in a given project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:migrationcenter/v1alpha1:Source")]
    public partial class Source : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp when the source was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Free-text description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// User-friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The number of frames that were reported by the source and contained errors.
        /// </summary>
        [Output("errorFrameCount")]
        public Output<int> ErrorFrameCount { get; private set; } = null!;

        /// <summary>
        /// If `true`, the source is managed by other service(s).
        /// </summary>
        [Output("isManaged")]
        public Output<bool> IsManaged { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The full name of the source.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of frames that are still being processed.
        /// </summary>
        [Output("pendingFrameCount")]
        public Output<int> PendingFrameCount { get; private set; } = null!;

        /// <summary>
        /// The information confidence of the source. The higher the value, the higher the confidence.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Required. User specified ID for the source. It will become the last component of the source name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression: `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
        /// </summary>
        [Output("sourceId")]
        public Output<string> SourceId { get; private set; } = null!;

        /// <summary>
        /// The state of the source.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Data source type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the source was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Source resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Source(string name, SourceArgs args, CustomResourceOptions? options = null)
            : base("google-native:migrationcenter/v1alpha1:Source", name, args ?? new SourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Source(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:migrationcenter/v1alpha1:Source", name, null, MakeResourceOptions(options, id))
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
                    "project",
                    "sourceId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Source resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Source Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Source(name, id, options);
        }
    }

    public sealed class SourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Free-text description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// If `true`, the source is managed by other service(s).
        /// </summary>
        [Input("isManaged")]
        public Input<bool>? IsManaged { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The information confidence of the source. The higher the value, the higher the confidence.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Required. User specified ID for the source. It will become the last component of the source name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression: `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
        /// </summary>
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        /// <summary>
        /// Data source type.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.MigrationCenter.V1Alpha1.SourceType>? Type { get; set; }

        public SourceArgs()
        {
        }
        public static new SourceArgs Empty => new SourceArgs();
    }
}
