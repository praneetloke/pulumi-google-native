// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Configuration options for Adaptive Protection auto-deploy feature.
    /// </summary>
    [OutputType]
    public sealed class SecurityPolicyAdaptiveProtectionConfigAutoDeployConfigResponse
    {
        public readonly double ConfidenceThreshold;
        public readonly int ExpirationSec;
        public readonly double ImpactedBaselineThreshold;
        public readonly double LoadThreshold;

        [OutputConstructor]
        private SecurityPolicyAdaptiveProtectionConfigAutoDeployConfigResponse(
            double confidenceThreshold,

            int expirationSec,

            double impactedBaselineThreshold,

            double loadThreshold)
        {
            ConfidenceThreshold = confidenceThreshold;
            ExpirationSec = expirationSec;
            ImpactedBaselineThreshold = impactedBaselineThreshold;
            LoadThreshold = loadThreshold;
        }
    }
}
