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
    /// Initiates a Cutover of a specific migrating VM. The returned LRO is completed when the cutover job resource is created and the job is initiated.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:vmmigration/v1alpha1:CutoverJob")]
    public partial class CutoverJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Details of the target VM in Compute Engine.
        /// </summary>
        [Output("computeEngineTargetDetails")]
        public Output<Outputs.ComputeEngineTargetDetailsResponse> ComputeEngineTargetDetails { get; private set; } = null!;

        /// <summary>
        /// Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_details instead.
        /// </summary>
        [Output("computeEngineVmDetails")]
        public Output<Outputs.TargetVMDetailsResponse> ComputeEngineVmDetails { get; private set; } = null!;

        /// <summary>
        /// The time the cutover job was created (as an API call, not when it was actually created in the target).
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Required. The cutover job identifier.
        /// </summary>
        [Output("cutoverJobId")]
        public Output<string> CutoverJobId { get; private set; } = null!;

        /// <summary>
        /// The time the cutover job had finished.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Provides details for the errors that led to the Cutover Job's state.
        /// </summary>
        [Output("error")]
        public Output<Outputs.StatusResponse> Error { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("migratingVmId")]
        public Output<string> MigratingVmId { get; private set; } = null!;

        /// <summary>
        /// The name of the cutover job.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current progress in percentage of the cutover job.
        /// </summary>
        [Output("progress")]
        public Output<int> Progress { get; private set; } = null!;

        /// <summary>
        /// The current progress in percentage of the cutover job.
        /// </summary>
        [Output("progressPercent")]
        public Output<int> ProgressPercent { get; private set; } = null!;

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
        /// State of the cutover job.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// A message providing possible extra details about the current state.
        /// </summary>
        [Output("stateMessage")]
        public Output<string> StateMessage { get; private set; } = null!;

        /// <summary>
        /// The time the state was last updated.
        /// </summary>
        [Output("stateTime")]
        public Output<string> StateTime { get; private set; } = null!;

        /// <summary>
        /// The cutover steps list representing its progress.
        /// </summary>
        [Output("steps")]
        public Output<ImmutableArray<Outputs.CutoverStepResponse>> Steps { get; private set; } = null!;

        /// <summary>
        /// Details of the VM to create as the target of this cutover job. Deprecated: Use compute_engine_target_details instead.
        /// </summary>
        [Output("targetDetails")]
        public Output<Outputs.TargetVMDetailsResponse> TargetDetails { get; private set; } = null!;


        /// <summary>
        /// Create a CutoverJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CutoverJob(string name, CutoverJobArgs args, CustomResourceOptions? options = null)
            : base("google-native:vmmigration/v1alpha1:CutoverJob", name, args ?? new CutoverJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CutoverJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:vmmigration/v1alpha1:CutoverJob", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "cutoverJobId",
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
        /// Get an existing CutoverJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CutoverJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CutoverJob(name, id, options);
        }
    }

    public sealed class CutoverJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The cutover job identifier.
        /// </summary>
        [Input("cutoverJobId", required: true)]
        public Input<string> CutoverJobId { get; set; } = null!;

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

        public CutoverJobArgs()
        {
        }
        public static new CutoverJobArgs Empty => new CutoverJobArgs();
    }
}
