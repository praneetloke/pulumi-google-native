// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1.Inputs
{

    /// <summary>
    /// The specifications for retries.
    /// </summary>
    public sealed class HttpRouteRetryPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the allowed number of retries. This number must be &gt; 0. If not specified, default to 1.
        /// </summary>
        [Input("numRetries")]
        public Input<int>? NumRetries { get; set; }

        /// <summary>
        /// Specifies a non-zero timeout per retry attempt.
        /// </summary>
        [Input("perTryTimeout")]
        public Input<string>? PerTryTimeout { get; set; }

        [Input("retryConditions")]
        private InputList<string>? _retryConditions;

        /// <summary>
        /// Specifies one or more conditions when this retry policy applies. Valid values are: 5xx: Proxy will attempt a retry if the destination service responds with any 5xx response code, of if the destination service does not respond at all, example: disconnect, reset, read timeout, connection failure and refused streams. gateway-error: Similar to 5xx, but only applies to response codes 502, 503, 504. reset: Proxy will attempt a retry if the destination service does not respond at all (disconnect/reset/read timeout) connect-failure: Proxy will retry on failures connecting to destination for example due to connection timeouts. retriable-4xx: Proxy will retry fro retriable 4xx response codes. Currently the only retriable error supported is 409. refused-stream: Proxy will retry if the destination resets the stream with a REFUSED_STREAM error code. This reset type indicates that it is safe to retry.
        /// </summary>
        public InputList<string> RetryConditions
        {
            get => _retryConditions ?? (_retryConditions = new InputList<string>());
            set => _retryConditions = value;
        }

        public HttpRouteRetryPolicyArgs()
        {
        }
        public static new HttpRouteRetryPolicyArgs Empty => new HttpRouteRetryPolicyArgs();
    }
}
