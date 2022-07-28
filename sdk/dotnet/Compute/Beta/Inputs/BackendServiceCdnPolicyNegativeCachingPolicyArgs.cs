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
    /// Specify CDN TTLs for response error codes.
    /// </summary>
    public sealed class BackendServiceCdnPolicyNegativeCachingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 302, 307, 308, 404, 405, 410, 421, 451 and 501 are can be specified as values, and you cannot specify a status code more than once.
        /// </summary>
        [Input("code")]
        public Input<int>? Code { get; set; }

        /// <summary>
        /// The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s (30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public BackendServiceCdnPolicyNegativeCachingPolicyArgs()
        {
        }
        public static new BackendServiceCdnPolicyNegativeCachingPolicyArgs Empty => new BackendServiceCdnPolicyNegativeCachingPolicyArgs();
    }
}
