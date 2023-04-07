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
    /// A set header action sets a header and forwards the request to the backend. This can be used to trigger custom protection implemented on the backend.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRecaptchaenterpriseV1FirewallActionSetHeaderActionResponse
    {
        /// <summary>
        /// The header key to set in the request to the backend server.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The header value to set in the request to the backend server.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GoogleCloudRecaptchaenterpriseV1FirewallActionSetHeaderActionResponse(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
