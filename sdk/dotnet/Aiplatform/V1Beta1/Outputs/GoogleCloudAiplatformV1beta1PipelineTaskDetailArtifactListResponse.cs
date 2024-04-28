// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Outputs
{

    /// <summary>
    /// A list of artifact metadata.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1PipelineTaskDetailArtifactListResponse
    {
        /// <summary>
        /// A list of artifact metadata.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1ArtifactResponse> Artifacts;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1PipelineTaskDetailArtifactListResponse(ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1ArtifactResponse> artifacts)
        {
            Artifacts = artifacts;
        }
    }
}
