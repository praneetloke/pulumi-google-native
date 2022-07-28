// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2
{
    public static class GetOrganizationExclusion
    {
        /// <summary>
        /// Gets the description of an exclusion in the _Default sink.
        /// </summary>
        public static Task<GetOrganizationExclusionResult> InvokeAsync(GetOrganizationExclusionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationExclusionResult>("google-native:logging/v2:getOrganizationExclusion", args ?? new GetOrganizationExclusionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the description of an exclusion in the _Default sink.
        /// </summary>
        public static Output<GetOrganizationExclusionResult> Invoke(GetOrganizationExclusionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrganizationExclusionResult>("google-native:logging/v2:getOrganizationExclusion", args ?? new GetOrganizationExclusionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationExclusionArgs : global::Pulumi.InvokeArgs
    {
        [Input("exclusionId", required: true)]
        public string ExclusionId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetOrganizationExclusionArgs()
        {
        }
        public static new GetOrganizationExclusionArgs Empty => new GetOrganizationExclusionArgs();
    }

    public sealed class GetOrganizationExclusionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("exclusionId", required: true)]
        public Input<string> ExclusionId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetOrganizationExclusionInvokeArgs()
        {
        }
        public static new GetOrganizationExclusionInvokeArgs Empty => new GetOrganizationExclusionInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationExclusionResult
    {
        /// <summary>
        /// The creation timestamp of the exclusion.This field may not be present for older exclusions.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. A description of this exclusion.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity&lt;ERROR sample(insertId, 0.99)
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The last update timestamp of the exclusion.This field may not be present for older exclusions.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetOrganizationExclusionResult(
            string createTime,

            string description,

            bool disabled,

            string filter,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Disabled = disabled;
            Filter = filter;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
