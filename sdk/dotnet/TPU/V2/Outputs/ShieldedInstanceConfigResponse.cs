// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2.Outputs
{

    /// <summary>
    /// A set of Shielded Instance options.
    /// </summary>
    [OutputType]
    public sealed class ShieldedInstanceConfigResponse
    {
        /// <summary>
        /// Defines whether the instance has Secure Boot enabled.
        /// </summary>
        public readonly bool EnableSecureBoot;

        [OutputConstructor]
        private ShieldedInstanceConfigResponse(bool enableSecureBoot)
        {
            EnableSecureBoot = enableSecureBoot;
        }
    }
}
