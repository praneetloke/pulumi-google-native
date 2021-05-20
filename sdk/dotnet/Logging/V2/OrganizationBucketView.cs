// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2
{
    /// <summary>
    /// Creates a view over logs in a bucket. A bucket may contain a maximum of 50 views.
    /// </summary>
    [GoogleNativeResourceType("google-native:logging/v2:OrganizationBucketView")]
    public partial class OrganizationBucketView : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation timestamp of the view.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Describes this view.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log id Example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        /// <summary>
        /// The resource name of the view. For example "projects/my-project-id/locations/my-location/buckets/my-bucket-id/views/my-view
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The last update timestamp of the view.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationBucketView resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationBucketView(string name, OrganizationBucketViewArgs args, CustomResourceOptions? options = null)
            : base("google-native:logging/v2:OrganizationBucketView", name, args ?? new OrganizationBucketViewArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationBucketView(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:logging/v2:OrganizationBucketView", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrganizationBucketView resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationBucketView Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationBucketView(name, id, options);
        }
    }

    public sealed class OrganizationBucketViewArgs : Pulumi.ResourceArgs
    {
        [Input("bucketId", required: true)]
        public Input<string> BucketId { get; set; } = null!;

        /// <summary>
        /// Describes this view.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log id Example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The resource name of the view. For example "projects/my-project-id/locations/my-location/buckets/my-bucket-id/views/my-view
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("viewId", required: true)]
        public Input<string> ViewId { get; set; } = null!;

        public OrganizationBucketViewArgs()
        {
        }
    }
}
