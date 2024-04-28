// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha.Outputs
{

    /// <summary>
    /// ResourceInfo represents the information/status of an app connector resource. Such as: - remote_agent - container - runtime - appgateway - appconnector - appconnection - tunnel - logagent
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponse
    {
        /// <summary>
        /// Specific details for the resource. This is for internal use only.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Resource;
        /// <summary>
        /// Overall health status. Overall status is derived based on the status of each sub level resources.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// List of Info for the sub level resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponse> Sub;
        /// <summary>
        /// The timestamp to collect the info. It is suggested to be set by the topmost level resource only.
        /// </summary>
        public readonly string Time;

        [OutputConstructor]
        private GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponse(
            ImmutableDictionary<string, object> resource,

            string status,

            ImmutableArray<Outputs.GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponse> sub,

            string time)
        {
            Resource = resource;
            Status = status;
            Sub = sub;
            Time = time;
        }
    }
}
