// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Iap.V1
{
    /// <summary>
    /// Constructs a new OAuth brand for the project if one does not exist. The created brand is "internal only", meaning that OAuth clients created under it only accept requests from users who belong to the same G Suite organization as the project. The brand is created in an un-reviewed status. NOTE: The "internal only" status can be manually changed in the Google Cloud console. Requires that a brand does not already exist for the project, and that the specified support email is owned by the caller.
    /// </summary>
    [GoogleNativeResourceType("google-native:iap/v1:Brand")]
    public partial class Brand : Pulumi.CustomResource
    {
        /// <summary>
        /// Application name displayed on OAuth consent screen.
        /// </summary>
        [Output("applicationTitle")]
        public Output<string> ApplicationTitle { get; private set; } = null!;

        /// <summary>
        /// Identifier of the brand. NOTE: GCP project number achieves the same brand identification purpose as only one brand per project can be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether the brand is only intended for usage inside the G Suite organization only.
        /// </summary>
        [Output("orgInternalOnly")]
        public Output<bool> OrgInternalOnly { get; private set; } = null!;

        /// <summary>
        /// Support email displayed on the OAuth consent screen.
        /// </summary>
        [Output("supportEmail")]
        public Output<string> SupportEmail { get; private set; } = null!;


        /// <summary>
        /// Create a Brand resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Brand(string name, BrandArgs args, CustomResourceOptions? options = null)
            : base("google-native:iap/v1:Brand", name, args ?? new BrandArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Brand(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:iap/v1:Brand", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Brand resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Brand Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Brand(name, id, options);
        }
    }

    public sealed class BrandArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application name displayed on OAuth consent screen.
        /// </summary>
        [Input("applicationTitle")]
        public Input<string>? ApplicationTitle { get; set; }

        [Input("brandId", required: true)]
        public Input<string> BrandId { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Support email displayed on the OAuth consent screen.
        /// </summary>
        [Input("supportEmail")]
        public Input<string>? SupportEmail { get; set; }

        public BrandArgs()
        {
        }
    }
}
