// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Settings controlling the volume of requests, connections and retries to this backend service.
    /// </summary>
    public sealed class CircuitBreakersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timeout for new network connections to hosts.
        /// </summary>
        [Input("connectTimeout")]
        public Input<Inputs.DurationArgs>? ConnectTimeout { get; set; }

        /// <summary>
        /// The maximum number of connections to the backend service. If not specified, there is no limit. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// The maximum number of pending requests allowed to the backend service. If not specified, there is no limit. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("maxPendingRequests")]
        public Input<int>? MaxPendingRequests { get; set; }

        /// <summary>
        /// The maximum number of parallel requests that allowed to the backend service. If not specified, there is no limit.
        /// </summary>
        [Input("maxRequests")]
        public Input<int>? MaxRequests { get; set; }

        /// <summary>
        /// Maximum requests for a single connection to the backend service. This parameter is respected by both the HTTP/1.1 and HTTP/2 implementations. If not specified, there is no limit. Setting this parameter to 1 will effectively disable keep alive. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("maxRequestsPerConnection")]
        public Input<int>? MaxRequestsPerConnection { get; set; }

        /// <summary>
        /// The maximum number of parallel retries allowed to the backend cluster. If not specified, the default is 1. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        public CircuitBreakersArgs()
        {
        }
        public static new CircuitBreakersArgs Empty => new CircuitBreakersArgs();
    }
}
