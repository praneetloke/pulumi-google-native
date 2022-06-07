// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// SlsaProvenanceZeroTwo is the slsa provenance as defined by the slsa spec. See full explanation of fields at slsa.dev/provenance/v0.2.
    /// </summary>
    [OutputType]
    public sealed class SlsaProvenanceZeroTwoResponse
    {
        /// <summary>
        /// Lists the steps in the build.
        /// </summary>
        public readonly ImmutableDictionary<string, string> BuildConfig;
        /// <summary>
        /// URI indicating what type of build was performed.
        /// </summary>
        public readonly string BuildType;
        /// <summary>
        /// Identifies the entity that executed the recipe, which is trusted to have correctly performed the operation and populated this provenance.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaBuilderResponse Builder;
        /// <summary>
        /// Identifies the event that kicked off the build.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaInvocationResponse Invocation;
        /// <summary>
        /// The collection of artifacts that influenced the build including sources, dependencies, build tools, base images, and so on.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaMaterialResponse> Materials;
        /// <summary>
        /// Other properties of the build.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaMetadataResponse Metadata;

        [OutputConstructor]
        private SlsaProvenanceZeroTwoResponse(
            ImmutableDictionary<string, string> buildConfig,

            string buildType,

            Outputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaBuilderResponse builder,

            Outputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaInvocationResponse invocation,

            ImmutableArray<Outputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaMaterialResponse> materials,

            Outputs.GoogleDevtoolsContaineranalysisV1alpha1SlsaProvenanceZeroTwoSlsaMetadataResponse metadata)
        {
            BuildConfig = buildConfig;
            BuildType = buildType;
            Builder = builder;
            Invocation = invocation;
            Materials = materials;
            Metadata = metadata;
        }
    }
}
