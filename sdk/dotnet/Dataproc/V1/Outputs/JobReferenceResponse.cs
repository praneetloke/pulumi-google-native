// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    [OutputType]
    public sealed class JobReferenceResponse
    {
        /// <summary>
        /// Optional. The job ID, which must be unique within the project.The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or hyphens (-). The maximum length is 100 characters.If not specified by the caller, the job ID will be provided by the server.
        /// </summary>
        public readonly string JobId;
        /// <summary>
        /// Optional. The ID of the Google Cloud Platform project that the job belongs to. If specified, must match the request project ID.
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private JobReferenceResponse(
            string jobId,

            string project)
        {
            JobId = jobId;
            Project = project;
        }
    }
}
