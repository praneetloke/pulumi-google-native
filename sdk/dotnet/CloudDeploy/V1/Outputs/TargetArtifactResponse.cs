// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// The artifacts produced by a target render operation.
    /// </summary>
    [OutputType]
    public sealed class TargetArtifactResponse
    {
        /// <summary>
        /// URI of a directory containing the artifacts. This contains deployment configuration used by Skaffold during a rollout, and all paths are relative to this location.
        /// </summary>
        public readonly string ArtifactUri;
        /// <summary>
        /// File path of the rendered manifest relative to the URI.
        /// </summary>
        public readonly string ManifestPath;
        /// <summary>
        /// Map from the phase ID to the phase artifacts for the `Target`.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.PhaseArtifactResponse> PhaseArtifacts;
        /// <summary>
        /// File path of the resolved Skaffold configuration relative to the URI.
        /// </summary>
        public readonly string SkaffoldConfigPath;

        [OutputConstructor]
        private TargetArtifactResponse(
            string artifactUri,

            string manifestPath,

            ImmutableDictionary<string, Outputs.PhaseArtifactResponse> phaseArtifacts,

            string skaffoldConfigPath)
        {
            ArtifactUri = artifactUri;
            ManifestPath = manifestPath;
            PhaseArtifacts = phaseArtifacts;
            SkaffoldConfigPath = skaffoldConfigPath;
        }
    }
}