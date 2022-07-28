// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// This submessage provides human-readable hints about the purpose of the authority. Because the name of a note acts as its resource reference, it is important to disambiguate the canonical name of the Note (which might be a UUID for security purposes) from "readable" names more suitable for debug output. Note that these hints should not be used to look up authorities in security sensitive contexts, such as when looking up attestations to verify.
    /// </summary>
    public sealed class DSSEHintArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human readable name of this attestation authority, for example "cloudbuild-prod".
        /// </summary>
        [Input("humanReadableName", required: true)]
        public Input<string> HumanReadableName { get; set; } = null!;

        public DSSEHintArgs()
        {
        }
        public static new DSSEHintArgs Empty => new DSSEHintArgs();
    }
}
