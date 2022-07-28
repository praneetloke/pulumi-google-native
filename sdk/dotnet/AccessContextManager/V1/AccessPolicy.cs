// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1
{
    /// <summary>
    /// Creates an access policy. This method fails if the organization already has an access policy. The long-running operation has a successful status after the access policy propagates to long-lasting storage. Syntactic and basic semantic errors are returned in `metadata` as a BadRequest proto.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:accesscontextmanager/v1:AccessPolicy")]
    public partial class AccessPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An opaque identifier for the current version of the `AccessPolicy`. This will always be a strongly validated etag, meaning that two Access Polices will be identical if and only if their etags are identical. Clients should not expect this to be in any specific format.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name of the `AccessPolicy`. Format: `accessPolicies/{access_policy}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// The scopes of a policy define which resources an ACM policy can restrict, and where ACM resources can be referenced. For example, a policy with scopes=["folders/123"] has the following behavior: - vpcsc perimeters can only restrict projects within folders/123 - access levels can only be referenced by resources within folders/123. If empty, there are no limitations on which resources can be restricted by an ACM policy, and there are no limitations on where ACM resources can be referenced. Only one policy can include a given scope (attempting to create a second policy which includes "folders/123" will result in an error). Currently, scopes cannot be modified after a policy is created. Currently, policies can only have a single scope. Format: list of `folders/{folder_number}` or `projects/{project_number}`
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Human readable title. Does not affect behavior.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPolicy(string name, AccessPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:accesscontextmanager/v1:AccessPolicy", name, args ?? new AccessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:accesscontextmanager/v1:AccessPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessPolicy(name, id, options);
        }
    }

    public sealed class AccessPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scopes of a policy define which resources an ACM policy can restrict, and where ACM resources can be referenced. For example, a policy with scopes=["folders/123"] has the following behavior: - vpcsc perimeters can only restrict projects within folders/123 - access levels can only be referenced by resources within folders/123. If empty, there are no limitations on which resources can be restricted by an ACM policy, and there are no limitations on where ACM resources can be referenced. Only one policy can include a given scope (attempting to create a second policy which includes "folders/123" will result in an error). Currently, scopes cannot be modified after a policy is created. Currently, policies can only have a single scope. Format: list of `folders/{folder_number}` or `projects/{project_number}`
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Human readable title. Does not affect behavior.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public AccessPolicyArgs()
        {
        }
        public static new AccessPolicyArgs Empty => new AccessPolicyArgs();
    }
}
