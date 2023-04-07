// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1
{
    public static class GetRule
    {
        /// <summary>
        /// Gets details of a single GatewaySecurityPolicyRule.
        /// </summary>
        public static Task<GetRuleResult> InvokeAsync(GetRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRuleResult>("google-native:networksecurity/v1beta1:getRule", args ?? new GetRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single GatewaySecurityPolicyRule.
        /// </summary>
        public static Output<GetRuleResult> Invoke(GetRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRuleResult>("google-native:networksecurity/v1beta1:getRule", args ?? new GetRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("gatewaySecurityPolicyId", required: true)]
        public string GatewaySecurityPolicyId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("ruleId", required: true)]
        public string RuleId { get; set; } = null!;

        public GetRuleArgs()
        {
        }
        public static new GetRuleArgs Empty => new GetRuleArgs();
    }

    public sealed class GetRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("gatewaySecurityPolicyId", required: true)]
        public Input<string> GatewaySecurityPolicyId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        public GetRuleInvokeArgs()
        {
        }
        public static new GetRuleInvokeArgs Empty => new GetRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetRuleResult
    {
        /// <summary>
        /// Optional. CEL expression for matching on L7/application level criteria.
        /// </summary>
        public readonly string ApplicationMatcher;
        /// <summary>
        /// Profile which tells what the primitive action should be.
        /// </summary>
        public readonly string BasicProfile;
        /// <summary>
        /// Time when the rule was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Free-text description of the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether the rule is enforced.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Immutable. Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule} rule should match the pattern: (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Priority of the rule. Lower number corresponds to higher precedence.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// CEL expression for matching on session criteria.
        /// </summary>
        public readonly string SessionMatcher;
        /// <summary>
        /// Optional. Flag to enable TLS inspection of traffic matching on , can only be true if the parent GatewaySecurityPolicy references a TLSInspectionConfig.
        /// </summary>
        public readonly bool TlsInspectionEnabled;
        /// <summary>
        /// Time when the rule was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetRuleResult(
            string applicationMatcher,

            string basicProfile,

            string createTime,

            string description,

            bool enabled,

            string name,

            int priority,

            string sessionMatcher,

            bool tlsInspectionEnabled,

            string updateTime)
        {
            ApplicationMatcher = applicationMatcher;
            BasicProfile = basicProfile;
            CreateTime = createTime;
            Description = description;
            Enabled = enabled;
            Name = name;
            Priority = priority;
            SessionMatcher = sessionMatcher;
            TlsInspectionEnabled = tlsInspectionEnabled;
            UpdateTime = updateTime;
        }
    }
}
