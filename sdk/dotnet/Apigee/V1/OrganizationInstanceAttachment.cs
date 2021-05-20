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
    /// Creates a new attachment of an environment to an instance. **Note:** Not supported for Apigee hybrid.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:OrganizationInstanceAttachment")]
    public partial class OrganizationInstanceAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Time the attachment was created in milliseconds since epoch.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// ID of the attached environment.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// ID of the attachment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationInstanceAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationInstanceAttachment(string name, OrganizationInstanceAttachmentArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationInstanceAttachment", name, args ?? new OrganizationInstanceAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationInstanceAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:OrganizationInstanceAttachment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationInstanceAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationInstanceAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationInstanceAttachment(name, id, options);
        }
    }

    public sealed class OrganizationInstanceAttachmentArgs : Pulumi.ResourceArgs
    {
        [Input("attachmentId", required: true)]
        public Input<string> AttachmentId { get; set; } = null!;

        /// <summary>
        /// ID of the attached environment.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public OrganizationInstanceAttachmentArgs()
        {
        }
    }
}
