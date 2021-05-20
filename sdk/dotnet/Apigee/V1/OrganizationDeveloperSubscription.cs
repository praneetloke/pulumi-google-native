// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates a subscription to an API product.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:OrganizationDeveloperSubscription")]
    public partial class OrganizationDeveloperSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the API product for which the developer is purchasing a subscription.
        /// </summary>
        [Output("apiproduct")]
        public Output<string> Apiproduct { get; private set; } = null!;

        /// <summary>
        /// Time when the API product subscription was created in milliseconds since epoch.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Time when the API product subscription ends in milliseconds since epoch.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Time when the API product subscription was last modified in milliseconds since epoch.
        /// </summary>
        [Output("lastModifiedAt")]
        public Output<string> LastModifiedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the API product subscription.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Time when the API product subscription starts in milliseconds since epoch.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationDeveloperSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationDeveloperSubscription(string name, OrganizationDeveloperSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationDeveloperSubscription", name, args ?? new OrganizationDeveloperSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationDeveloperSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationDeveloperSubscription", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationDeveloperSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationDeveloperSubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationDeveloperSubscription(name, id, options);
        }
    }

    public sealed class OrganizationDeveloperSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the API product for which the developer is purchasing a subscription.
        /// </summary>
        [Input("apiproduct")]
        public Input<string>? Apiproduct { get; set; }

        [Input("developerId", required: true)]
        public Input<string> DeveloperId { get; set; } = null!;

        /// <summary>
        /// Time when the API product subscription ends in milliseconds since epoch.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Time when the API product subscription starts in milliseconds since epoch.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        public OrganizationDeveloperSubscriptionArgs()
        {
        }
    }
}
