// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class SecurityPolicyRuleHttpHeaderActionHttpHeaderOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the header to set.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// The value to set the named header to.
        /// </summary>
        [Input("headerValue")]
        public Input<string>? HeaderValue { get; set; }

        public SecurityPolicyRuleHttpHeaderActionHttpHeaderOptionArgs()
        {
        }
        public static new SecurityPolicyRuleHttpHeaderActionHttpHeaderOptionArgs Empty => new SecurityPolicyRuleHttpHeaderActionHttpHeaderOptionArgs();
    }
}
