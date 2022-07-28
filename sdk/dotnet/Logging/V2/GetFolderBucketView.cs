// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2
{
    public static class GetFolderBucketView
    {
        /// <summary>
        /// Gets a view on a log bucket..
        /// </summary>
        public static Task<GetFolderBucketViewResult> InvokeAsync(GetFolderBucketViewArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFolderBucketViewResult>("google-native:logging/v2:getFolderBucketView", args ?? new GetFolderBucketViewArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a view on a log bucket..
        /// </summary>
        public static Output<GetFolderBucketViewResult> Invoke(GetFolderBucketViewInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFolderBucketViewResult>("google-native:logging/v2:getFolderBucketView", args ?? new GetFolderBucketViewInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFolderBucketViewArgs : global::Pulumi.InvokeArgs
    {
        [Input("bucketId", required: true)]
        public string BucketId { get; set; } = null!;

        [Input("folderId", required: true)]
        public string FolderId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("viewId", required: true)]
        public string ViewId { get; set; } = null!;

        public GetFolderBucketViewArgs()
        {
        }
        public static new GetFolderBucketViewArgs Empty => new GetFolderBucketViewArgs();
    }

    public sealed class GetFolderBucketViewInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("bucketId", required: true)]
        public Input<string> BucketId { get; set; } = null!;

        [Input("folderId", required: true)]
        public Input<string> FolderId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("viewId", required: true)]
        public Input<string> ViewId { get; set; } = null!;

        public GetFolderBucketViewInvokeArgs()
        {
        }
        public static new GetFolderBucketViewInvokeArgs Empty => new GetFolderBucketViewInvokeArgs();
    }


    [OutputType]
    public sealed class GetFolderBucketViewResult
    {
        /// <summary>
        /// The creation timestamp of the view.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Describes this view.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The last update timestamp of the view.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetFolderBucketViewResult(
            string createTime,

            string description,

            string filter,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Filter = filter;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
