// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    /// <summary>
    /// A predicate which describes the SBOM being referenced.
    /// </summary>
    [OutputType]
    public sealed class SbomReferenceIntotoPredicateResponse
    {
        /// <summary>
        /// A map of algorithm to digest of the contents of the SBOM.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Digest;
        /// <summary>
        /// The location of the SBOM.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The mime type of the SBOM.
        /// </summary>
        public readonly string MimeType;
        /// <summary>
        /// The person or system referring this predicate to the consumer.
        /// </summary>
        public readonly string ReferrerId;

        [OutputConstructor]
        private SbomReferenceIntotoPredicateResponse(
            ImmutableDictionary<string, string> digest,

            string location,

            string mimeType,

            string referrerId)
        {
            Digest = digest;
            Location = location;
            MimeType = mimeType;
            ReferrerId = referrerId;
        }
    }
}
