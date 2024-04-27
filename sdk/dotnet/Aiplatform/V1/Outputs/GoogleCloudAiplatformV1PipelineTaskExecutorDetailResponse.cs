// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Outputs
{

    /// <summary>
    /// The runtime detail of a pipeline executor.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1PipelineTaskExecutorDetailResponse
    {
        /// <summary>
        /// The detailed info for a container executor.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1PipelineTaskExecutorDetailContainerDetailResponse ContainerDetail;
        /// <summary>
        /// The detailed info for a custom job executor.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1PipelineTaskExecutorDetailCustomJobDetailResponse CustomJobDetail;

        [OutputConstructor]
        private GoogleCloudAiplatformV1PipelineTaskExecutorDetailResponse(
            Outputs.GoogleCloudAiplatformV1PipelineTaskExecutorDetailContainerDetailResponse containerDetail,

            Outputs.GoogleCloudAiplatformV1PipelineTaskExecutorDetailCustomJobDetailResponse customJobDetail)
        {
            ContainerDetail = containerDetail;
            CustomJobDetail = customJobDetail;
        }
    }
}