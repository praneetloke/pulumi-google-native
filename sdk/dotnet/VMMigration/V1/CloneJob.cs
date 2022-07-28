// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1
{
    /// <summary>
    /// Initiates a Clone of a specific migrating VM.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:vmmigration/v1:CloneJob")]
    public partial class CloneJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The clone job identifier.
        /// </summary>
        [Output("cloneJobId")]
        public Output<string> CloneJobId { get; private set; } = null!;

        /// <summary>
        /// Details of the target VM in Compute Engine.
        /// </summary>
        [Output("computeEngineTargetDetails")]
        public Output<Outputs.ComputeEngineTargetDetailsResponse> ComputeEngineTargetDetails { get; private set; } = null!;

        /// <summary>
        /// The time the clone job was created (as an API call, not when it was actually created in the target).
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time the clone job was ended.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Provides details for the errors that led to the Clone Job's state.
        /// </summary>
        [Output("error")]
        public Output<Outputs.StatusResponse> Error { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("migratingVmId")]
        public Output<string> MigratingVmId { get; private set; } = null!;

        /// <summary>
        /// The name of the clone.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        [Output("sourceId")]
        public Output<string> SourceId { get; private set; } = null!;

        /// <summary>
        /// State of the clone job.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The time the state was last updated.
        /// </summary>
        [Output("stateTime")]
        public Output<string> StateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CloneJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloneJob(string name, CloneJobArgs args, CustomResourceOptions? options = null)
            : base("google-native:vmmigration/v1:CloneJob", name, args ?? new CloneJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloneJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:vmmigration/v1:CloneJob", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "cloneJobId",
                    "location",
                    "migratingVmId",
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
        /// Get an existing CloneJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloneJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CloneJob(name, id, options);
        }
    }

    public sealed class CloneJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The clone job identifier.
        /// </summary>
        [Input("cloneJobId", required: true)]
        public Input<string> CloneJobId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("migratingVmId", required: true)]
        public Input<string> MigratingVmId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public CloneJobArgs()
        {
        }
        public static new CloneJobArgs Empty => new CloneJobArgs();
    }
}
