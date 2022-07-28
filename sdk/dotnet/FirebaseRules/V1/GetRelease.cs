// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseRules.V1
{
    public static class GetRelease
    {
        /// <summary>
        /// Get a `Release` by name.
        /// </summary>
        public static Task<GetReleaseResult> InvokeAsync(GetReleaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReleaseResult>("google-native:firebaserules/v1:getRelease", args ?? new GetReleaseArgs(), options.WithDefaults());

        /// <summary>
        /// Get a `Release` by name.
        /// </summary>
        public static Output<GetReleaseResult> Invoke(GetReleaseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReleaseResult>("google-native:firebaserules/v1:getRelease", args ?? new GetReleaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReleaseArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("releaseId", required: true)]
        public string ReleaseId { get; set; } = null!;

        public GetReleaseArgs()
        {
        }
        public static new GetReleaseArgs Empty => new GetReleaseArgs();
    }

    public sealed class GetReleaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("releaseId", required: true)]
        public Input<string> ReleaseId { get; set; } = null!;

        public GetReleaseInvokeArgs()
        {
        }
        public static new GetReleaseInvokeArgs Empty => new GetReleaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetReleaseResult
    {
        /// <summary>
        /// Time the release was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Format: `projects/{project_id}/releases/{release_id}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        /// </summary>
        public readonly string RulesetName;
        /// <summary>
        /// Time the release was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetReleaseResult(
            string createTime,

            string name,

            string rulesetName,

            string updateTime)
        {
            CreateTime = createTime;
            Name = name;
            RulesetName = rulesetName;
            UpdateTime = updateTime;
        }
    }
}
