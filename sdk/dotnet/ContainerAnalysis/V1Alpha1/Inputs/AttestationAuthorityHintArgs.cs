// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// This submessage provides human-readable hints about the purpose of the AttestationAuthority. Because the name of a Note acts as its resource reference, it is important to disambiguate the canonical name of the Note (which might be a UUID for security purposes) from "readable" names more suitable for debug output. Note that these hints should NOT be used to look up AttestationAuthorities in security sensitive contexts, such as when looking up Attestations to verify.
    /// </summary>
    public sealed class AttestationAuthorityHintArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human readable name of this Attestation Authority, for example "qa".
        /// </summary>
        [Input("humanReadableName")]
        public Input<string>? HumanReadableName { get; set; }

        public AttestationAuthorityHintArgs()
        {
        }
        public static new AttestationAuthorityHintArgs Empty => new AttestationAuthorityHintArgs();
    }
}
