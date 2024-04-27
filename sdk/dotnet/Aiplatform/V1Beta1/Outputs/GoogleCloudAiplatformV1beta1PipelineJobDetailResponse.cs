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
    /// The runtime detail of PipelineJob.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1PipelineJobDetailResponse
    {
        /// <summary>
        /// The context of the pipeline.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1ContextResponse PipelineContext;
        /// <summary>
        /// The context of the current pipeline run.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1beta1ContextResponse PipelineRunContext;
        /// <summary>
        /// The runtime details of the tasks under the pipeline.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1PipelineTaskDetailResponse> TaskDetails;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1PipelineJobDetailResponse(
            Outputs.GoogleCloudAiplatformV1beta1ContextResponse pipelineContext,

            Outputs.GoogleCloudAiplatformV1beta1ContextResponse pipelineRunContext,

            ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1PipelineTaskDetailResponse> taskDetails)
        {
            PipelineContext = pipelineContext;
            PipelineRunContext = pipelineRunContext;
            TaskDetails = taskDetails;
        }
    }
}