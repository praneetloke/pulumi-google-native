// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Represents a rule that describes one or more match conditions along with the action to be taken when traffic matches this condition (allow or deny).
    /// </summary>
    [OutputType]
    public sealed class FirewallPolicyRuleResponse
    {
        /// <summary>
        /// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// An optional description for this resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The direction in which this rule applies.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
        /// </summary>
        public readonly bool EnableLogging;
        /// <summary>
        /// [Output only] Type of the resource. Always compute#firewallPolicyRule for firewall policy rules
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        /// </summary>
        public readonly Outputs.FirewallPolicyRuleMatcherResponse Match;
        /// <summary>
        /// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// An optional name for the rule. This field is not a unique identifier and can be updated.
        /// </summary>
        public readonly string RuleName;
        /// <summary>
        /// Calculation of the complexity of a single firewall policy rule.
        /// </summary>
        public readonly int RuleTupleCount;
        /// <summary>
        /// A fully-qualified URL of a SecurityProfile resource instance. Example: https://networksecurity.googleapis.com/v1/projects/{project}/locations/{location}/securityProfileGroups/my-security-profile-group Must be specified if action = 'apply_profile_group' and cannot be specified for other actions.
        /// </summary>
        public readonly string SecurityProfileGroup;
        /// <summary>
        /// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
        /// </summary>
        public readonly ImmutableArray<string> TargetResources;
        /// <summary>
        /// A list of secure tags that controls which instances the firewall rule applies to. If targetSecureTag are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the target_secure_tag are in INEFFECTIVE state, then this rule will be ignored. targetSecureTag may not be set at the same time as targetServiceAccounts. If neither targetServiceAccounts nor targetSecureTag are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyRuleSecureTagResponse> TargetSecureTags;
        /// <summary>
        /// A list of service accounts indicating the sets of instances that are applied with this rule.
        /// </summary>
        public readonly ImmutableArray<string> TargetServiceAccounts;

        [OutputConstructor]
        private FirewallPolicyRuleResponse(
            string action,

            string description,

            string direction,

            bool disabled,

            bool enableLogging,

            string kind,

            Outputs.FirewallPolicyRuleMatcherResponse match,

            int priority,

            string ruleName,

            int ruleTupleCount,

            string securityProfileGroup,

            ImmutableArray<string> targetResources,

            ImmutableArray<Outputs.FirewallPolicyRuleSecureTagResponse> targetSecureTags,

            ImmutableArray<string> targetServiceAccounts)
        {
            Action = action;
            Description = description;
            Direction = direction;
            Disabled = disabled;
            EnableLogging = enableLogging;
            Kind = kind;
            Match = match;
            Priority = priority;
            RuleName = ruleName;
            RuleTupleCount = ruleTupleCount;
            SecurityProfileGroup = securityProfileGroup;
            TargetResources = targetResources;
            TargetSecureTags = targetSecureTags;
            TargetServiceAccounts = targetServiceAccounts;
        }
    }
}
