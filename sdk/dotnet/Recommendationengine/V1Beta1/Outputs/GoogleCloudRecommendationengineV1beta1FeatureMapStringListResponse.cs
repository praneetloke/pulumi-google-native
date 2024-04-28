// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Recommendationengine.V1Beta1.Outputs
{

    /// <summary>
    /// A list of string features.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRecommendationengineV1beta1FeatureMapStringListResponse
    {
        /// <summary>
        /// String feature value with a length limit of 128 bytes.
        /// </summary>
        public readonly ImmutableArray<string> Value;

        [OutputConstructor]
        private GoogleCloudRecommendationengineV1beta1FeatureMapStringListResponse(ImmutableArray<string> value)
        {
            Value = value;
        }
    }
}
