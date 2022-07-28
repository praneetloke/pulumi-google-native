// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Configuration options for L7 DDoS detection.
    /// </summary>
    public sealed class SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, enables CAAP for L7 DDoS detection.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Rule visibility can be one of the following: STANDARD - opaque rules. (default) PREMIUM - transparent rules.
        /// </summary>
        [Input("ruleVisibility")]
        public Input<Pulumi.GoogleNative.Compute.V1.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigRuleVisibility>? RuleVisibility { get; set; }

        public SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs()
        {
        }
        public static new SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs Empty => new SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs();
    }
}
