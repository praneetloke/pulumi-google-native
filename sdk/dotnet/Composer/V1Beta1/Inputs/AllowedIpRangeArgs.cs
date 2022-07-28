// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1Beta1.Inputs
{

    /// <summary>
    /// Allowed IP range with user-provided description.
    /// </summary>
    public sealed class AllowedIpRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. User-provided description. It must contain at most 300 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IP address or range, defined using CIDR notation, of requests that this rule applies to. Examples: `192.168.1.1` or `192.168.0.0/16` or `2001:db8::/32` or `2001:0db8:0000:0042:0000:8a2e:0370:7334`. IP range prefixes should be properly truncated. For example, `1.2.3.4/24` should be truncated to `1.2.3.0/24`. Similarly, for IPv6, `2001:db8::1/32` should be truncated to `2001:db8::/32`.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public AllowedIpRangeArgs()
        {
        }
        public static new AllowedIpRangeArgs Empty => new AllowedIpRangeArgs();
    }
}
