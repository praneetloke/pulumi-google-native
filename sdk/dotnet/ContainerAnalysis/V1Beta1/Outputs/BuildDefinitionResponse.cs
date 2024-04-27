// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    [OutputType]
    public sealed class BuildDefinitionResponse
    {
        public readonly string BuildType;
        public readonly ImmutableDictionary<string, string> ExternalParameters;
        public readonly ImmutableDictionary<string, string> InternalParameters;
        public readonly ImmutableArray<Outputs.ResourceDescriptorResponse> ResolvedDependencies;

        [OutputConstructor]
        private BuildDefinitionResponse(
            string buildType,

            ImmutableDictionary<string, string> externalParameters,

            ImmutableDictionary<string, string> internalParameters,

            ImmutableArray<Outputs.ResourceDescriptorResponse> resolvedDependencies)
        {
            BuildType = buildType;
            ExternalParameters = externalParameters;
            InternalParameters = internalParameters;
            ResolvedDependencies = resolvedDependencies;
        }
    }
}