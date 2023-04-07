// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.Outputs
{

    /// <summary>
    /// An individual action. Each action represents what to do if a policy matches.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRecaptchaenterpriseV1FirewallActionResponse
    {
        /// <summary>
        /// The user request did not match any policy and should be allowed access to the requested resource.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionAllowActionResponse Allow;
        /// <summary>
        /// This action will deny access to a given page. The user will get an HTTP error code.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionBlockActionResponse Block;
        /// <summary>
        /// This action will redirect the request to a ReCaptcha interstitial to attach a token.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionRedirectActionResponse Redirect;
        /// <summary>
        /// This action will set a custom header but allow the request to continue to the customer backend.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionSetHeaderActionResponse SetHeader;
        /// <summary>
        /// This action will transparently serve a different page to an offending user.
        /// </summary>
        public readonly Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionSubstituteActionResponse Substitute;

        [OutputConstructor]
        private GoogleCloudRecaptchaenterpriseV1FirewallActionResponse(
            Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionAllowActionResponse allow,

            Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionBlockActionResponse block,

            Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionRedirectActionResponse redirect,

            Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionSetHeaderActionResponse setHeader,

            Outputs.GoogleCloudRecaptchaenterpriseV1FirewallActionSubstituteActionResponse substitute)
        {
            Allow = allow;
            Block = block;
            Redirect = redirect;
            SetHeader = setHeader;
            Substitute = substitute;
        }
    }
}
