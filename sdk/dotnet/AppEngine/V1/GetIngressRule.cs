// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1
{
    public static class GetIngressRule
    {
        /// <summary>
        /// Gets the specified firewall rule.
        /// </summary>
        public static Task<GetIngressRuleResult> InvokeAsync(GetIngressRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIngressRuleResult>("google-native:appengine/v1:getIngressRule", args ?? new GetIngressRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified firewall rule.
        /// </summary>
        public static Output<GetIngressRuleResult> Invoke(GetIngressRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIngressRuleResult>("google-native:appengine/v1:getIngressRule", args ?? new GetIngressRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIngressRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        [Input("ingressRuleId", required: true)]
        public string IngressRuleId { get; set; } = null!;

        public GetIngressRuleArgs()
        {
        }
        public static new GetIngressRuleArgs Empty => new GetIngressRuleArgs();
    }

    public sealed class GetIngressRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("ingressRuleId", required: true)]
        public Input<string> IngressRuleId { get; set; } = null!;

        public GetIngressRuleInvokeArgs()
        {
        }
        public static new GetIngressRuleInvokeArgs Empty => new GetIngressRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetIngressRuleResult
    {
        /// <summary>
        /// The action to take on matched requests.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// An optional string description of this rule. This field has a maximum length of 400 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
        /// </summary>
        public readonly string SourceRange;

        [OutputConstructor]
        private GetIngressRuleResult(
            string action,

            string description,

            int priority,

            string sourceRange)
        {
            Action = action;
            Description = description;
            Priority = priority;
            SourceRange = sourceRange;
        }
    }
}
