// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1.Outputs
{

    /// <summary>
    /// Specification of the GRPC Endpoint.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudNetworksecurityV1beta1GrpcEndpointResponse
    {
        /// <summary>
        /// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
        /// </summary>
        public readonly string TargetUri;

        [OutputConstructor]
        private GoogleCloudNetworksecurityV1beta1GrpcEndpointResponse(string targetUri)
        {
            TargetUri = targetUri;
        }
    }
}
