// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ApigeeRegistry.V1
{
    public static class GetApi
    {
        /// <summary>
        /// Returns a specified API.
        /// </summary>
        public static Task<GetApiResult> InvokeAsync(GetApiArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApiResult>("google-native:apigeeregistry/v1:getApi", args ?? new GetApiArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a specified API.
        /// </summary>
        public static Output<GetApiResult> Invoke(GetApiInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiResult>("google-native:apigeeregistry/v1:getApi", args ?? new GetApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetApiArgs()
        {
        }
        public static new GetApiArgs Empty => new GetApiArgs();
    }

    public sealed class GetApiInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetApiInvokeArgs()
        {
        }
        public static new GetApiInvokeArgs Empty => new GetApiInvokeArgs();
    }


    [OutputType]
    public sealed class GetApiResult
    {
        /// <summary>
        /// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// A user-definable description of the availability of this service. Format: free-form, but we expect single words that describe availability, e.g., "NONE", "TESTING", "PREVIEW", "GENERAL", "DEPRECATED", "SHUTDOWN".
        /// </summary>
        public readonly string Availability;
        /// <summary>
        /// Creation timestamp.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// A detailed description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Human-meaningful name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores, and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The recommended deployment of the API. Format: `apis/{api}/deployments/{deployment}`
        /// </summary>
        public readonly string RecommendedDeployment;
        /// <summary>
        /// The recommended version of the API. Format: `apis/{api}/versions/{version}`
        /// </summary>
        public readonly string RecommendedVersion;
        /// <summary>
        /// Last update timestamp.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetApiResult(
            ImmutableDictionary<string, string> annotations,

            string availability,

            string createTime,

            string description,

            string displayName,

            ImmutableDictionary<string, string> labels,

            string name,

            string recommendedDeployment,

            string recommendedVersion,

            string updateTime)
        {
            Annotations = annotations;
            Availability = availability;
            CreateTime = createTime;
            Description = description;
            DisplayName = displayName;
            Labels = labels;
            Name = name;
            RecommendedDeployment = recommendedDeployment;
            RecommendedVersion = recommendedVersion;
            UpdateTime = updateTime;
        }
    }
}
