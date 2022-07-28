// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1.Inputs
{

    /// <summary>
    /// The `MembershipRole` expiry details.
    /// </summary>
    public sealed class ExpiryDetailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time at which the `MembershipRole` will expire.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        public ExpiryDetailArgs()
        {
        }
        public static new ExpiryDetailArgs Empty => new ExpiryDetailArgs();
    }
}
