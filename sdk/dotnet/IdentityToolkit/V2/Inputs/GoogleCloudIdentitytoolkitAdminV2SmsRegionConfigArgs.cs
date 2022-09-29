// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IdentityToolkit.V2.Inputs
{

    /// <summary>
    /// Configures the regions where users are allowed to send verification SMS for the project or tenant. This is based on the calling code of the destination phone number.
    /// </summary>
    public sealed class GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A policy of allowing SMS to every region by default and adding disallowed regions to a disallow list.
        /// </summary>
        [Input("allowByDefault")]
        public Input<Inputs.GoogleCloudIdentitytoolkitAdminV2AllowByDefaultArgs>? AllowByDefault { get; set; }

        /// <summary>
        /// A policy of only allowing regions by explicitly adding them to an allowlist.
        /// </summary>
        [Input("allowlistOnly")]
        public Input<Inputs.GoogleCloudIdentitytoolkitAdminV2AllowlistOnlyArgs>? AllowlistOnly { get; set; }

        public GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs()
        {
        }
        public static new GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs Empty => new GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs();
    }
}
