// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Beta
{
    /// <summary>
    /// Creates a RBACRoleBinding.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:gkehub/v1beta:Rbacrolebinding")]
    public partial class Rbacrolebinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When the rbacrolebinding was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// When the rbacrolebinding was deleted.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// group is the group, as seen by the kubernetes cluster.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name for the rbacrolebinding `projects/{project}/locations/{location}/namespaces/{namespace}/rbacrolebindings/{rbacrolebinding}` or `projects/{project}/locations/{location}/memberships/{membership}/rbacrolebindings/{rbacrolebinding}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Required. Client chosen ID for the RBACRoleBinding. `rbacrolebinding_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
        /// </summary>
        [Output("rbacrolebindingId")]
        public Output<string> RbacrolebindingId { get; private set; } = null!;

        /// <summary>
        /// Role to bind to the principal
        /// </summary>
        [Output("role")]
        public Output<Outputs.RoleResponse> Role { get; private set; } = null!;

        /// <summary>
        /// State of the rbacrolebinding resource.
        /// </summary>
        [Output("state")]
        public Output<Outputs.RBACRoleBindingLifecycleStateResponse> State { get; private set; } = null!;

        /// <summary>
        /// Google-generated UUID for this resource. This is unique across all rbacrolebinding resources. If a rbacrolebinding resource is deleted and another resource with the same name is created, it gets a different uid.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// When the rbacrolebinding was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld"
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a Rbacrolebinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rbacrolebinding(string name, RbacrolebindingArgs args, CustomResourceOptions? options = null)
            : base("google-native:gkehub/v1beta:Rbacrolebinding", name, args ?? new RbacrolebindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rbacrolebinding(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gkehub/v1beta:Rbacrolebinding", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "namespaceId",
                    "project",
                    "rbacrolebindingId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Rbacrolebinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rbacrolebinding Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Rbacrolebinding(name, id, options);
        }
    }

    public sealed class RbacrolebindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// group is the group, as seen by the kubernetes cluster.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name for the rbacrolebinding `projects/{project}/locations/{location}/namespaces/{namespace}/rbacrolebindings/{rbacrolebinding}` or `projects/{project}/locations/{location}/memberships/{membership}/rbacrolebindings/{rbacrolebinding}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Required. Client chosen ID for the RBACRoleBinding. `rbacrolebinding_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
        /// </summary>
        [Input("rbacrolebindingId", required: true)]
        public Input<string> RbacrolebindingId { get; set; } = null!;

        /// <summary>
        /// Role to bind to the principal
        /// </summary>
        [Input("role", required: true)]
        public Input<Inputs.RoleArgs> Role { get; set; } = null!;

        /// <summary>
        /// user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld"
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public RbacrolebindingArgs()
        {
        }
        public static new RbacrolebindingArgs Empty => new RbacrolebindingArgs();
    }
}
