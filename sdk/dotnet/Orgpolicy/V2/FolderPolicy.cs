// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Orgpolicy.V2
{
    /// <summary>
    /// Creates a Policy. Returns a `google.rpc.Status` with `google.rpc.Code.NOT_FOUND` if the constraint does not exist. Returns a `google.rpc.Status` with `google.rpc.Code.ALREADY_EXISTS` if the policy already exists on the given Cloud resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:orgpolicy/v2:FolderPolicy")]
    public partial class FolderPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Basic information about the Organization Policy.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.GoogleCloudOrgpolicyV2PolicySpecResponse> Spec { get; private set; } = null!;


        /// <summary>
        /// Create a FolderPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FolderPolicy(string name, FolderPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:orgpolicy/v2:FolderPolicy", name, args ?? new FolderPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FolderPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:orgpolicy/v2:FolderPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing FolderPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FolderPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FolderPolicy(name, id, options);
        }
    }

    public sealed class FolderPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("folderId", required: true)]
        public Input<string> FolderId { get; set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// Basic information about the Organization Policy.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.GoogleCloudOrgpolicyV2PolicySpecArgs>? Spec { get; set; }

        public FolderPolicyArgs()
        {
        }
    }
}
