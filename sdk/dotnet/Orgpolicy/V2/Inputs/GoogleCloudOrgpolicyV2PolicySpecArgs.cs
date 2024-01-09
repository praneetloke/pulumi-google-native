// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Orgpolicy.V2.Inputs
{

    /// <summary>
    /// Defines a Google Cloud policy specification which is used to specify constraints for configurations of Google Cloud resources.
    /// </summary>
    public sealed class GoogleCloudOrgpolicyV2PolicySpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An opaque tag indicating the current version of the policySpec, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the policy is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current policySpec to use when executing a read-modify-write loop. When the policy is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Determines the inheritance behavior for this policy. If `inherit_from_parent` is true, policy rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the new root for evaluation. This field can be set only for policies which configure list constraints.
        /// </summary>
        [Input("inheritFromParent")]
        public Input<bool>? InheritFromParent { get; set; }

        /// <summary>
        /// Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
        /// </summary>
        [Input("reset")]
        public Input<bool>? Reset { get; set; }

        [Input("rules")]
        private InputList<Inputs.GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs>? _rules;

        /// <summary>
        /// In policies for boolean constraints, the following requirements apply: - There must be one and only one policy rule where condition is unset. - Boolean policy rules with conditions must set `enforced` to the opposite of the policy rule without a condition. - During policy evaluation, policy rules with conditions that are true for a target resource take precedence.
        /// </summary>
        public InputList<Inputs.GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.GoogleCloudOrgpolicyV2PolicySpecPolicyRuleArgs>());
            set => _rules = value;
        }

        public GoogleCloudOrgpolicyV2PolicySpecArgs()
        {
        }
        public static new GoogleCloudOrgpolicyV2PolicySpecArgs Empty => new GoogleCloudOrgpolicyV2PolicySpecArgs();
    }
}
