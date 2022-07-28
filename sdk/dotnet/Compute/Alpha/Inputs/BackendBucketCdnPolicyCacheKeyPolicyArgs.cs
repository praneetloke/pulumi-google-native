// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Message containing what to include in the cache key for a request for Cloud CDN.
    /// </summary>
    public sealed class BackendBucketCdnPolicyCacheKeyPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("includeHttpHeaders")]
        private InputList<string>? _includeHttpHeaders;

        /// <summary>
        /// Allows HTTP request headers (by name) to be used in the cache key.
        /// </summary>
        public InputList<string> IncludeHttpHeaders
        {
            get => _includeHttpHeaders ?? (_includeHttpHeaders = new InputList<string>());
            set => _includeHttpHeaders = value;
        }

        [Input("queryStringWhitelist")]
        private InputList<string>? _queryStringWhitelist;

        /// <summary>
        /// Names of query string parameters to include in cache keys. Default parameters are always included. '&amp;' and '=' will be percent encoded and not treated as delimiters.
        /// </summary>
        public InputList<string> QueryStringWhitelist
        {
            get => _queryStringWhitelist ?? (_queryStringWhitelist = new InputList<string>());
            set => _queryStringWhitelist = value;
        }

        public BackendBucketCdnPolicyCacheKeyPolicyArgs()
        {
        }
        public static new BackendBucketCdnPolicyCacheKeyPolicyArgs Empty => new BackendBucketCdnPolicyCacheKeyPolicyArgs();
    }
}
