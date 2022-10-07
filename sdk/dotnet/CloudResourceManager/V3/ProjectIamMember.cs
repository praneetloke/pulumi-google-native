// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudResourceManager.V3
{
    /// <summary>
    /// Sets the IAM access control policy for the specified project, in the format `projects/{ProjectIdOrNumber}` e.g. projects/123. CAUTION: This method will replace the existing policy, and cannot be used to append additional IAM settings. Note: Removing service accounts from policies or changing their roles can render services completely inoperable. It is important to understand how the service account is being used before removing or updating its roles. The following constraints apply when using `setIamPolicy()`: + Project does not support `allUsers` and `allAuthenticatedUsers` as `members` in a `Binding` of a `Policy`. + The owner role can be granted to a `user`, `serviceAccount`, or a group that is part of an organization. For example, group@myownpersonaldomain.com could be added as an owner to a project in the myownpersonaldomain.com organization, but not the examplepetstore.com organization. + Service accounts can be made owners of a project directly without any restrictions. However, to be added as an owner, a user must be invited using the Cloud Platform console and must accept the invitation. + A user cannot be granted the owner role using `setIamPolicy()`. The user must be granted the owner role using the Cloud Platform Console and must explicitly accept the invitation. + Invitations to grant the owner role cannot be sent using `setIamPolicy()`; they must be sent only using the Cloud Platform Console. + If the project is not part of an organization, there must be at least one owner who has accepted the Terms of Service (ToS) agreement in the policy. Calling `setIamPolicy()` to remove the last ToS-accepted owner from the policy will fail. This restriction also applies to legacy projects that no longer have owners who have accepted the ToS. Edits to IAM policies will be rejected until the lack of a ToS-accepting owner is rectified. If the project is part of an organization, you can remove all owners, potentially making the organization inaccessible.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudresourcemanager/v3:ProjectIamMember")]
    public partial class ProjectIamMember : global::Pulumi.CustomResource
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
        /// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
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
        /// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectIamMember(string name, ProjectIamMemberArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudresourcemanager/v3:ProjectIamMember", name, args ?? new ProjectIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectIamMember(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudresourcemanager/v3:ProjectIamMember", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectIamMember Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProjectIamMember(name, id, options);
        }
    }

    public sealed class ProjectIamMemberArgs : global::Pulumi.ResourceArgs
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

        public ProjectIamMemberArgs()
        {
        }
        public static new ProjectIamMemberArgs Empty => new ProjectIamMemberArgs();
    }
}
