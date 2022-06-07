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
    /// Identifies the event that kicked off the build.
    /// </summary>
    [OutputType]
    public sealed class GrafeasV1SlsaProvenanceZeroTwoSlsaInvocationResponse
    {
        public readonly Outputs.GrafeasV1SlsaProvenanceZeroTwoSlsaConfigSourceResponse ConfigSource;
        public readonly ImmutableDictionary<string, string> Environment;
        public readonly ImmutableDictionary<string, string> Parameters;

        [OutputConstructor]
        private GrafeasV1SlsaProvenanceZeroTwoSlsaInvocationResponse(
            Outputs.GrafeasV1SlsaProvenanceZeroTwoSlsaConfigSourceResponse configSource,

            ImmutableDictionary<string, string> environment,

            ImmutableDictionary<string, string> parameters)
        {
            ConfigSource = configSource;
            Environment = environment;
            Parameters = parameters;
        }
    }
}
