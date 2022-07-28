// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Configuration options for Cloud Armor Adaptive Protection (CAAP).
    /// </summary>
    public sealed class SecurityPolicyAdaptiveProtectionConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoDeployConfig")]
        public Input<Inputs.SecurityPolicyAdaptiveProtectionConfigAutoDeployConfigArgs>? AutoDeployConfig { get; set; }

        /// <summary>
        /// If set to true, enables Cloud Armor Machine Learning.
        /// </summary>
        [Input("layer7DdosDefenseConfig")]
        public Input<Inputs.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs>? Layer7DdosDefenseConfig { get; set; }

        public SecurityPolicyAdaptiveProtectionConfigArgs()
        {
        }
        public static new SecurityPolicyAdaptiveProtectionConfigArgs Empty => new SecurityPolicyAdaptiveProtectionConfigArgs();
    }
}
