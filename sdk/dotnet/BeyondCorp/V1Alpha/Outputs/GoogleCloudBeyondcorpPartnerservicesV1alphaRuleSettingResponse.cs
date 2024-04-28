// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha.Outputs
{

    /// <summary>
    /// Message to capture settings for a BrowserDlpRule
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudBeyondcorpPartnerservicesV1alphaRuleSettingResponse
    {
        /// <summary>
        /// Immutable. The type of the Setting. .
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value of the Setting.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Value;

        [OutputConstructor]
        private GoogleCloudBeyondcorpPartnerservicesV1alphaRuleSettingResponse(
            string type,

            ImmutableDictionary<string, object> value)
        {
            Type = type;
            Value = value;
        }
    }
}
