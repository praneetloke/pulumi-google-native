// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2
{
    public static class GetBillingAccountBucket
    {
        /// <summary>
        /// Gets a log bucket.
        /// </summary>
        public static Task<GetBillingAccountBucketResult> InvokeAsync(GetBillingAccountBucketArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBillingAccountBucketResult>("google-native:logging/v2:getBillingAccountBucket", args ?? new GetBillingAccountBucketArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a log bucket.
        /// </summary>
        public static Output<GetBillingAccountBucketResult> Invoke(GetBillingAccountBucketInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBillingAccountBucketResult>("google-native:logging/v2:getBillingAccountBucket", args ?? new GetBillingAccountBucketInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBillingAccountBucketArgs : global::Pulumi.InvokeArgs
    {
        [Input("billingAccountId", required: true)]
        public string BillingAccountId { get; set; } = null!;

        [Input("bucketId", required: true)]
        public string BucketId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        public GetBillingAccountBucketArgs()
        {
        }
        public static new GetBillingAccountBucketArgs Empty => new GetBillingAccountBucketArgs();
    }

    public sealed class GetBillingAccountBucketInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("billingAccountId", required: true)]
        public Input<string> BillingAccountId { get; set; } = null!;

        [Input("bucketId", required: true)]
        public Input<string> BucketId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        public GetBillingAccountBucketInvokeArgs()
        {
        }
        public static new GetBillingAccountBucketInvokeArgs Empty => new GetBillingAccountBucketInvokeArgs();
    }


    [OutputType]
    public sealed class GetBillingAccountBucketResult
    {
        /// <summary>
        /// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
        /// </summary>
        public readonly Outputs.CmekSettingsResponse CmekSettings;
        /// <summary>
        /// The creation timestamp of the bucket. This is not set for any of the default buckets.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Describes this bucket.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A list of indexed fields and related configuration data.
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexConfigResponse> IndexConfigs;
        /// <summary>
        /// The bucket lifecycle state.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
        /// </summary>
        public readonly bool Locked;
        /// <summary>
        /// The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
        /// </summary>
        public readonly ImmutableArray<string> RestrictedFields;
        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        /// </summary>
        public readonly int RetentionDays;
        /// <summary>
        /// The last update timestamp of the bucket.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetBillingAccountBucketResult(
            Outputs.CmekSettingsResponse cmekSettings,

            string createTime,

            string description,

            ImmutableArray<Outputs.IndexConfigResponse> indexConfigs,

            string lifecycleState,

            bool locked,

            string name,

            ImmutableArray<string> restrictedFields,

            int retentionDays,

            string updateTime)
        {
            CmekSettings = cmekSettings;
            CreateTime = createTime;
            Description = description;
            IndexConfigs = indexConfigs;
            LifecycleState = lifecycleState;
            Locked = locked;
            Name = name;
            RestrictedFields = restrictedFields;
            RetentionDays = retentionDays;
            UpdateTime = updateTime;
        }
    }
}
