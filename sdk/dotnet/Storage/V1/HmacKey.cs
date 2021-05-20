// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1
{
    /// <summary>
    /// Creates a new HMAC key for the specified service account.
    /// </summary>
    [GoogleNativeResourceType("google-native:storage/v1:HmacKey")]
    public partial class HmacKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the HMAC Key.
        /// </summary>
        [Output("accessId")]
        public Output<string> AccessId { get; private set; } = null!;

        /// <summary>
        /// HTTP 1.1 Entity tag for the HMAC key.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The kind of item this is. For HMAC Key metadata, this is always storage#hmacKeyMetadata.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Project ID owning the service account to which the key authenticates.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The link to this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The email address of the key's associated service account.
        /// </summary>
        [Output("serviceAccountEmail")]
        public Output<string> ServiceAccountEmail { get; private set; } = null!;

        /// <summary>
        /// The state of the key. Can be one of ACTIVE, INACTIVE, or DELETED.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The creation time of the HMAC key in RFC 3339 format.
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// The last modification time of the HMAC key metadata in RFC 3339 format.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a HmacKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HmacKey(string name, HmacKeyArgs args, CustomResourceOptions? options = null)
            : base("google-native:storage/v1:HmacKey", name, args ?? new HmacKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HmacKey(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:storage/v1:HmacKey", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing HmacKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HmacKey Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HmacKey(name, id, options);
        }
    }

    public sealed class HmacKeyArgs : Pulumi.ResourceArgs
    {
        [Input("accessId", required: true)]
        public Input<string> AccessId { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        [Input("userProject")]
        public Input<string>? UserProject { get; set; }

        public HmacKeyArgs()
        {
        }
    }
}
