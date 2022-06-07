// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    /// <summary>
    /// The collection of artifacts that influenced the build including sources, dependencies, build tools, base images, and so on.
    /// </summary>
    [OutputType]
    public sealed class GrafeasV1SlsaProvenanceZeroTwoSlsaMaterialResponse
    {
        public readonly ImmutableDictionary<string, string> Digest;
        public readonly string Uri;

        [OutputConstructor]
        private GrafeasV1SlsaProvenanceZeroTwoSlsaMaterialResponse(
            ImmutableDictionary<string, string> digest,

            string uri)
        {
            Digest = digest;
            Uri = uri;
        }
    }
}
