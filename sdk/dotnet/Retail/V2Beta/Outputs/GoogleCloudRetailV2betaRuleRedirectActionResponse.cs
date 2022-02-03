// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Beta.Outputs
{

    /// <summary>
    /// Redirects a shopper to a specific page. * Rule Condition: - Must specify Condition. * Action Input: Request Query * Action Result: Redirects shopper to provided uri.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2betaRuleRedirectActionResponse
    {
        /// <summary>
        /// URL must have length equal or less than 2000 characters.
        /// </summary>
        public readonly string RedirectUri;

        [OutputConstructor]
        private GoogleCloudRetailV2betaRuleRedirectActionResponse(string redirectUri)
        {
            RedirectUri = redirectUri;
        }
    }
}