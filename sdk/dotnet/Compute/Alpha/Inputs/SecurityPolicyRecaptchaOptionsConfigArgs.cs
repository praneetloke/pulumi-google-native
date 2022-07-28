// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class SecurityPolicyRecaptchaOptionsConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional field to supply a reCAPTCHA site key to be used for all the rules using the redirect action with the type of GOOGLE_RECAPTCHA under the security policy. The specified site key needs to be created from the reCAPTCHA API. The user is responsible for the validity of the specified site key. If not specified, a Google-managed site key is used.
        /// </summary>
        [Input("redirectSiteKey")]
        public Input<string>? RedirectSiteKey { get; set; }

        public SecurityPolicyRecaptchaOptionsConfigArgs()
        {
        }
        public static new SecurityPolicyRecaptchaOptionsConfigArgs Empty => new SecurityPolicyRecaptchaOptionsConfigArgs();
    }
}
