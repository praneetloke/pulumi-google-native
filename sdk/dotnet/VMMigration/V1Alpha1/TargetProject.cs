// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1
{
    /// <summary>
    /// Creates a new TargetProject in a given project. NOTE: TargetProject is a global resource; hence the only supported value for location is `global`.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:vmmigration/v1alpha1:TargetProject")]
    public partial class TargetProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time this target project resource was created (not related to when the Compute Engine project it points to was created).
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The target project's description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the target project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Required. The target_project identifier.
        /// </summary>
        [Output("targetProjectId")]
        public Output<string> TargetProjectId { get; private set; } = null!;

        /// <summary>
        /// The last time the target project resource was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a TargetProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetProject(string name, TargetProjectArgs args, CustomResourceOptions? options = null)
            : base("google-native:vmmigration/v1alpha1:TargetProject", name, args ?? new TargetProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetProject(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:vmmigration/v1alpha1:TargetProject", name, null, MakeResourceOptions(options, id))
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
                    "targetProjectId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TargetProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetProject Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TargetProject(name, id, options);
        }
    }

    public sealed class TargetProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target project's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The target project ID (number) or project name.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Required. The target_project identifier.
        /// </summary>
        [Input("targetProjectId", required: true)]
        public Input<string> TargetProjectId { get; set; } = null!;

        public TargetProjectArgs()
        {
        }
        public static new TargetProjectArgs Empty => new TargetProjectArgs();
    }
}
