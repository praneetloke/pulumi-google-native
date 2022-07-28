// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ApigeeRegistry.V1
{
    public static class GetVersion
    {
        /// <summary>
        /// GetApiVersion returns a specified version.
        /// </summary>
        public static Task<GetVersionResult> InvokeAsync(GetVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVersionResult>("google-native:apigeeregistry/v1:getVersion", args ?? new GetVersionArgs(), options.WithDefaults());

        /// <summary>
        /// GetApiVersion returns a specified version.
        /// </summary>
        public static Output<GetVersionResult> Invoke(GetVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVersionResult>("google-native:apigeeregistry/v1:getVersion", args ?? new GetVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("versionId", required: true)]
        public string VersionId { get; set; } = null!;

        public GetVersionArgs()
        {
        }
        public static new GetVersionArgs Empty => new GetVersionArgs();
    }

    public sealed class GetVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public GetVersionInvokeArgs()
        {
        }
        public static new GetVersionInvokeArgs Empty => new GetVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetVersionResult
    {
        /// <summary>
        /// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
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
        /// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A user-definable description of the lifecycle phase of this API version. Format: free-form, but we expect single words that describe API maturity, e.g. "CONCEPT", "DESIGN", "DEVELOPMENT", "STAGING", "PRODUCTION", "DEPRECATED", "RETIRED".
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Last update timestamp.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetVersionResult(
            ImmutableDictionary<string, string> annotations,

            string createTime,

            string description,

            string displayName,

            ImmutableDictionary<string, string> labels,

            string name,

            string state,

            string updateTime)
        {
            Annotations = annotations;
            CreateTime = createTime;
            Description = description;
            DisplayName = displayName;
            Labels = labels;
            Name = name;
            State = state;
            UpdateTime = updateTime;
        }
    }
}
