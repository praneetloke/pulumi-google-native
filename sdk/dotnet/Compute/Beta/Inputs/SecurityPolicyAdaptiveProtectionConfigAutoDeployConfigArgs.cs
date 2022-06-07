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
    /// Configuration options for Adaptive Protection auto-deploy feature.
    /// </summary>
    public sealed class SecurityPolicyAdaptiveProtectionConfigAutoDeployConfigArgs : Pulumi.ResourceArgs
    {
        [Input("confidenceThreshold")]
        public Input<double>? ConfidenceThreshold { get; set; }

        [Input("expirationSec")]
        public Input<int>? ExpirationSec { get; set; }

        [Input("impactedBaselineThreshold")]
        public Input<double>? ImpactedBaselineThreshold { get; set; }

        [Input("loadThreshold")]
        public Input<double>? LoadThreshold { get; set; }

        public SecurityPolicyAdaptiveProtectionConfigAutoDeployConfigArgs()
        {
        }
    }
}
