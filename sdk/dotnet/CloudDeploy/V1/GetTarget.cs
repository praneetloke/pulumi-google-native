// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1
{
    public static class GetTarget
    {
        /// <summary>
        /// Gets details of a single Target.
        /// </summary>
        public static Task<GetTargetResult> InvokeAsync(GetTargetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetResult>("google-native:clouddeploy/v1:getTarget", args ?? new GetTargetArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Target.
        /// </summary>
        public static Output<GetTargetResult> Invoke(GetTargetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTargetResult>("google-native:clouddeploy/v1:getTarget", args ?? new GetTargetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("targetId", required: true)]
        public string TargetId { get; set; } = null!;

        public GetTargetArgs()
        {
        }
    }

    public sealed class GetTargetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("targetId", required: true)]
        public Input<string> TargetId { get; set; } = null!;

        public GetTargetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetResult
    {
        /// <summary>
        /// Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// Information specifying an Anthos Cluster.
        /// </summary>
        public readonly Outputs.AnthosClusterResponse AnthosCluster;
        /// <summary>
        /// Time at which the `Target` was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Description of the `Target`. Max length is 255 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Configurations for all execution that relates to this `Target`. Each `ExecutionEnvironmentUsage` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values. When no configurations are specified, execution will use the default specified in `DefaultPool`.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExecutionConfigResponse> ExecutionConfigs;
        /// <summary>
        /// Information specifying a GKE Cluster.
        /// </summary>
        public readonly Outputs.GkeClusterResponse Gke;
        /// <summary>
        /// Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Name of the `Target`. Format is projects/{project}/locations/{location}/targets/a-z{0,62}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. Whether or not the `Target` requires approval.
        /// </summary>
        public readonly bool RequireApproval;
        /// <summary>
        /// Resource id of the `Target`.
        /// </summary>
        public readonly string TargetId;
        /// <summary>
        /// Unique identifier of the `Target`.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// Most recent time at which the `Target` was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetTargetResult(
            ImmutableDictionary<string, string> annotations,

            Outputs.AnthosClusterResponse anthosCluster,

            string createTime,

            string description,

            string etag,

            ImmutableArray<Outputs.ExecutionConfigResponse> executionConfigs,

            Outputs.GkeClusterResponse gke,

            ImmutableDictionary<string, string> labels,

            string name,

            bool requireApproval,

            string targetId,

            string uid,

            string updateTime)
        {
            Annotations = annotations;
            AnthosCluster = anthosCluster;
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            ExecutionConfigs = executionConfigs;
            Gke = gke;
            Labels = labels;
            Name = name;
            RequireApproval = requireApproval;
            TargetId = targetId;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
