// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.CloudDeploy.V1
{
    public static class GetRollout
    {
        /// <summary>
        /// Gets details of a single Rollout.
        /// </summary>
        public static Task<GetRolloutResult> InvokeAsync(GetRolloutArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRolloutResult>("google-native:clouddeploy/v1:getRollout", args ?? new GetRolloutArgs(), options.WithVersion());

        /// <summary>
        /// Gets details of a single Rollout.
        /// </summary>
        public static Output<GetRolloutResult> Invoke(GetRolloutInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRolloutResult>("google-native:clouddeploy/v1:getRollout", args ?? new GetRolloutInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRolloutArgs : Pulumi.InvokeArgs
    {
        [Input("deliveryPipelineId", required: true)]
        public string DeliveryPipelineId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("releaseId", required: true)]
        public string ReleaseId { get; set; } = null!;

        [Input("rolloutId", required: true)]
        public string RolloutId { get; set; } = null!;

        public GetRolloutArgs()
        {
        }
    }

    public sealed class GetRolloutInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("deliveryPipelineId", required: true)]
        public Input<string> DeliveryPipelineId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("releaseId", required: true)]
        public Input<string> ReleaseId { get; set; } = null!;

        [Input("rolloutId", required: true)]
        public Input<string> RolloutId { get; set; } = null!;

        public GetRolloutInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRolloutResult
    {
        /// <summary>
        /// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// Approval state of the `Rollout`.
        /// </summary>
        public readonly string ApprovalState;
        /// <summary>
        /// Time at which the `Rollout` was approved.
        /// </summary>
        public readonly string ApproveTime;
        /// <summary>
        /// Time at which the `Rollout` was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Time at which the `Rollout` finished deploying.
        /// </summary>
        public readonly string DeployEndTime;
        /// <summary>
        /// Time at which the `Rollout` started deploying.
        /// </summary>
        public readonly string DeployStartTime;
        /// <summary>
        /// The resource name of the Cloud Build `Build` object that is used to deploy the Rollout. Format is `projects/{project}/locations/{location}/builds/{build}`.
        /// </summary>
        public readonly string DeployingBuild;
        /// <summary>
        /// Description of the `Rollout` for user purposes. Max length is 255 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Time at which the `Rollout` was enqueued.
        /// </summary>
        public readonly string EnqueueTime;
        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Reason the build failed. Empty if the build succeeded.
        /// </summary>
        public readonly string FailureReason;
        /// <summary>
        /// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Name of the `Rollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Current state of the `Rollout`.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The ID of Target to which this `Rollout` is deploying.
        /// </summary>
        public readonly string TargetId;
        /// <summary>
        /// Unique identifier of the `Rollout`.
        /// </summary>
        public readonly string Uid;

        [OutputConstructor]
        private GetRolloutResult(
            ImmutableDictionary<string, string> annotations,

            string approvalState,

            string approveTime,

            string createTime,

            string deployEndTime,

            string deployStartTime,

            string deployingBuild,

            string description,

            string enqueueTime,

            string etag,

            string failureReason,

            ImmutableDictionary<string, string> labels,

            string name,

            string state,

            string targetId,

            string uid)
        {
            Annotations = annotations;
            ApprovalState = approvalState;
            ApproveTime = approveTime;
            CreateTime = createTime;
            DeployEndTime = deployEndTime;
            DeployStartTime = deployStartTime;
            DeployingBuild = deployingBuild;
            Description = description;
            EnqueueTime = enqueueTime;
            Etag = etag;
            FailureReason = failureReason;
            Labels = labels;
            Name = name;
            State = state;
            TargetId = targetId;
            Uid = uid;
        }
    }
}