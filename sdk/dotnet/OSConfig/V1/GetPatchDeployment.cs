// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1
{
    public static class GetPatchDeployment
    {
        /// <summary>
        /// Get an OS Config patch deployment.
        /// </summary>
        public static Task<GetPatchDeploymentResult> InvokeAsync(GetPatchDeploymentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPatchDeploymentResult>("google-native:osconfig/v1:getPatchDeployment", args ?? new GetPatchDeploymentArgs(), options.WithVersion());
    }


    public sealed class GetPatchDeploymentArgs : Pulumi.InvokeArgs
    {
        [Input("patchDeploymentId", required: true)]
        public string PatchDeploymentId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetPatchDeploymentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPatchDeploymentResult
    {
        /// <summary>
        /// Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Duration of the patch. After the duration ends, the patch times out.
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// VM instances to patch.
        /// </summary>
        public readonly Outputs.PatchInstanceFilterResponse InstanceFilter;
        /// <summary>
        /// The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        public readonly string LastExecuteTime;
        /// <summary>
        /// Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Schedule a one-time execution.
        /// </summary>
        public readonly Outputs.OneTimeScheduleResponse OneTimeSchedule;
        /// <summary>
        /// Optional. Patch configuration that is applied.
        /// </summary>
        public readonly Outputs.PatchConfigResponse PatchConfig;
        /// <summary>
        /// Schedule recurring executions.
        /// </summary>
        public readonly Outputs.RecurringScheduleResponse RecurringSchedule;
        /// <summary>
        /// Optional. Rollout strategy of the patch job.
        /// </summary>
        public readonly Outputs.PatchRolloutResponse Rollout;
        /// <summary>
        /// Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetPatchDeploymentResult(
            string createTime,

            string description,

            string duration,

            Outputs.PatchInstanceFilterResponse instanceFilter,

            string lastExecuteTime,

            string name,

            Outputs.OneTimeScheduleResponse oneTimeSchedule,

            Outputs.PatchConfigResponse patchConfig,

            Outputs.RecurringScheduleResponse recurringSchedule,

            Outputs.PatchRolloutResponse rollout,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Duration = duration;
            InstanceFilter = instanceFilter;
            LastExecuteTime = lastExecuteTime;
            Name = name;
            OneTimeSchedule = oneTimeSchedule;
            PatchConfig = patchConfig;
            RecurringSchedule = recurringSchedule;
            Rollout = rollout;
            UpdateTime = updateTime;
        }
    }
}