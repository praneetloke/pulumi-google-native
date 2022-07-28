// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1
{
    public static class GetCloneJob
    {
        /// <summary>
        /// Gets details of a single CloneJob.
        /// </summary>
        public static Task<GetCloneJobResult> InvokeAsync(GetCloneJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloneJobResult>("google-native:vmmigration/v1:getCloneJob", args ?? new GetCloneJobArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single CloneJob.
        /// </summary>
        public static Output<GetCloneJobResult> Invoke(GetCloneJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloneJobResult>("google-native:vmmigration/v1:getCloneJob", args ?? new GetCloneJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloneJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloneJobId", required: true)]
        public string CloneJobId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("migratingVmId", required: true)]
        public string MigratingVmId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetCloneJobArgs()
        {
        }
        public static new GetCloneJobArgs Empty => new GetCloneJobArgs();
    }

    public sealed class GetCloneJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloneJobId", required: true)]
        public Input<string> CloneJobId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("migratingVmId", required: true)]
        public Input<string> MigratingVmId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetCloneJobInvokeArgs()
        {
        }
        public static new GetCloneJobInvokeArgs Empty => new GetCloneJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloneJobResult
    {
        /// <summary>
        /// Details of the target VM in Compute Engine.
        /// </summary>
        public readonly Outputs.ComputeEngineTargetDetailsResponse ComputeEngineTargetDetails;
        /// <summary>
        /// The time the clone job was created (as an API call, not when it was actually created in the target).
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The time the clone job was ended.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Provides details for the errors that led to the Clone Job's state.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// The name of the clone.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the clone job.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The time the state was last updated.
        /// </summary>
        public readonly string StateTime;

        [OutputConstructor]
        private GetCloneJobResult(
            Outputs.ComputeEngineTargetDetailsResponse computeEngineTargetDetails,

            string createTime,

            string endTime,

            Outputs.StatusResponse error,

            string name,

            string state,

            string stateTime)
        {
            ComputeEngineTargetDetails = computeEngineTargetDetails;
            CreateTime = createTime;
            EndTime = endTime;
            Error = error;
            Name = name;
            State = state;
            StateTime = stateTime;
        }
    }
}
