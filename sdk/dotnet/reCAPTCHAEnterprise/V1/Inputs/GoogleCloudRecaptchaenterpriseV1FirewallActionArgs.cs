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
    /// An individual action. Each action represents what to do if a policy matches.
    /// </summary>
    public sealed class GoogleCloudRecaptchaenterpriseV1FirewallActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user request did not match any policy and should be allowed access to the requested resource.
        /// </summary>
        [Input("allow")]
        public Input<Inputs.GoogleCloudRecaptchaenterpriseV1FirewallActionAllowActionArgs>? Allow { get; set; }

        /// <summary>
        /// This action will deny access to a given page. The user will get an HTTP error code.
        /// </summary>
        [Input("block")]
        public Input<Inputs.GoogleCloudRecaptchaenterpriseV1FirewallActionBlockActionArgs>? Block { get; set; }

        /// <summary>
        /// This action will redirect the request to a ReCaptcha interstitial to attach a token.
        /// </summary>
        [Input("redirect")]
        public Input<Inputs.GoogleCloudRecaptchaenterpriseV1FirewallActionRedirectActionArgs>? Redirect { get; set; }

        /// <summary>
        /// This action will set a custom header but allow the request to continue to the customer backend.
        /// </summary>
        [Input("setHeader")]
        public Input<Inputs.GoogleCloudRecaptchaenterpriseV1FirewallActionSetHeaderActionArgs>? SetHeader { get; set; }

        /// <summary>
        /// This action will transparently serve a different page to an offending user.
        /// </summary>
        [Input("substitute")]
        public Input<Inputs.GoogleCloudRecaptchaenterpriseV1FirewallActionSubstituteActionArgs>? Substitute { get; set; }

        public GoogleCloudRecaptchaenterpriseV1FirewallActionArgs()
        {
        }
        public static new GoogleCloudRecaptchaenterpriseV1FirewallActionArgs Empty => new GoogleCloudRecaptchaenterpriseV1FirewallActionArgs();
    }
}
