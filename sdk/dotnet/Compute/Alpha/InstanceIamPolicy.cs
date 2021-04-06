// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Alpha
{
    /// <summary>
    /// Sets the access control policy on the specified resource. Replaces any existing policy.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:compute/alpha:InstanceIamPolicy")]
    public partial class InstanceIamPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies cloud audit logging configuration for this policy.
        /// </summary>
        [Output("auditConfigs")]
        public Output<ImmutableArray<Outputs.AuditConfigResponse>> AuditConfigs { get; private set; } = null!;

        /// <summary>
        /// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
        /// </summary>
        [Output("bindings")]
        public Output<ImmutableArray<Outputs.BindingResponse>> Bindings { get; private set; } = null!;

        /// <summary>
        /// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
        /// 
        /// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("iamOwned")]
        public Output<bool> IamOwned { get; private set; } = null!;

        /// <summary>
        /// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.RuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// Specifies the format of the policy.
        /// 
        /// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
        /// 
        /// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
        /// 
        /// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
        /// 
        /// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
        /// 
        /// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
        /// 
        /// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceIamPolicy(string name, InstanceIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:compute/alpha:InstanceIamPolicy", name, args ?? new InstanceIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceIamPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:compute/alpha:InstanceIamPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceIamPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InstanceIamPolicy(name, id, options);
        }
    }

    public sealed class InstanceIamPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("auditConfigs")]
        private InputList<Inputs.AuditConfigArgs>? _auditConfigs;

        /// <summary>
        /// Specifies cloud audit logging configuration for this policy.
        /// </summary>
        public InputList<Inputs.AuditConfigArgs> AuditConfigs
        {
            get => _auditConfigs ?? (_auditConfigs = new InputList<Inputs.AuditConfigArgs>());
            set => _auditConfigs = value;
        }

        [Input("bindings")]
        private InputList<Inputs.BindingArgs>? _bindings;

        /// <summary>
        /// Flatten Policy to create a backwacd compatible wire-format. Deprecated. Use 'policy' to specify bindings.
        /// </summary>
        public InputList<Inputs.BindingArgs> Bindings
        {
            get => _bindings ?? (_bindings = new InputList<Inputs.BindingArgs>());
            set => _bindings = value;
        }

        /// <summary>
        /// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify the etag.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("iamOwned")]
        public Input<bool>? IamOwned { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        [Input("rules")]
        private InputList<Inputs.RuleArgs>? _rules;

        /// <summary>
        /// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
        /// </summary>
        public InputList<Inputs.RuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies the format of the policy.
        /// 
        /// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
        /// 
        /// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
        /// 
        /// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
        /// 
        /// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
        /// 
        /// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
        /// 
        /// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public InstanceIamPolicyArgs()
        {
        }
    }
}
