// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Inputs
{

    /// <summary>
    /// The type of an input artifact.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1beta1PipelineJobRuntimeConfigInputArtifactArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Artifact resource id from MLMD. Which is the last portion of an artifact resource name: `projects/{project}/locations/{location}/metadataStores/default/artifacts/{artifact_id}`. The artifact must stay within the same project, location and default metadatastore as the pipeline.
        /// </summary>
        [Input("artifactId")]
        public Input<string>? ArtifactId { get; set; }

        public GoogleCloudAiplatformV1beta1PipelineJobRuntimeConfigInputArtifactArgs()
        {
        }
        public static new GoogleCloudAiplatformV1beta1PipelineJobRuntimeConfigInputArtifactArgs Empty => new GoogleCloudAiplatformV1beta1PipelineJobRuntimeConfigInputArtifactArgs();
    }
}
