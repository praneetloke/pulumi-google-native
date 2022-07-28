// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.Inputs
{

    /// <summary>
    /// Settings specific to keys that can be used for WAF (Web Application Firewall).
    /// </summary>
    public sealed class GoogleCloudRecaptchaenterpriseV1WafSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The WAF feature for which this key is enabled.
        /// </summary>
        [Input("wafFeature", required: true)]
        public Input<Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature> WafFeature { get; set; } = null!;

        /// <summary>
        /// The WAF service that uses this key.
        /// </summary>
        [Input("wafService", required: true)]
        public Input<Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.GoogleCloudRecaptchaenterpriseV1WafSettingsWafService> WafService { get; set; } = null!;

        public GoogleCloudRecaptchaenterpriseV1WafSettingsArgs()
        {
        }
        public static new GoogleCloudRecaptchaenterpriseV1WafSettingsArgs Empty => new GoogleCloudRecaptchaenterpriseV1WafSettingsArgs();
    }
}
