// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BinaryAuthorization.V1
{
    /// <summary>
    /// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
    /// </summary>
    [GoogleNativeResourceType("google-native:binaryauthorization/v1:AttestorIamMember")]
    public partial class AttestorIamMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
        /// </summary>
        [Output("condition")]
        public Output<Pulumi.GoogleNative.IAM.V1.Outputs.Condition?> Condition { get; private set; } = null!;

        /// <summary>
        /// The etag of the resource's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Identity that will be granted the privilege in role. The entry can have one of the following values:
        /// 
        ///  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        ///  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        ///  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
        ///  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The name of the resource to manage IAM policies for.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it is not provided, a default will be supplied.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a AttestorIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttestorIamMember(string name, AttestorIamMemberArgs args, CustomResourceOptions? options = null)
            : base("google-native:binaryauthorization/v1:AttestorIamMember", name, args ?? new AttestorIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AttestorIamMember(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:binaryauthorization/v1:AttestorIamMember", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AttestorIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttestorIamMember Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AttestorIamMember(name, id, options);
        }
    }

    public sealed class AttestorIamMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An IAM Condition for a given binding.
        /// </summary>
        [Input("condition")]
        public Input<Pulumi.GoogleNative.IAM.V1.Inputs.ConditionArgs>? Condition { get; set; }

        /// <summary>
        /// Identity that will be granted the privilege in role. The entry can have one of the following values:
        /// 
        ///  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        ///  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        ///  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
        ///  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The name of the resource to manage IAM policies for.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The role that should be applied.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public AttestorIamMemberArgs()
        {
        }
        public static new AttestorIamMemberArgs Empty => new AttestorIamMemberArgs();
    }
}
