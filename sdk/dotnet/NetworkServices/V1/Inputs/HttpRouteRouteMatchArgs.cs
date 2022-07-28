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
    /// RouteMatch defines specifications used to match requests. If multiple match types are set, this RouteMatch will match if ALL type of matches are matched.
    /// </summary>
    public sealed class HttpRouteRouteMatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HTTP request path value should exactly match this value. Only one of full_path_match, prefix_match, or regex_match should be used.
        /// </summary>
        [Input("fullPathMatch")]
        public Input<string>? FullPathMatch { get; set; }

        [Input("headers")]
        private InputList<Inputs.HttpRouteHeaderMatchArgs>? _headers;

        /// <summary>
        /// Specifies a list of HTTP request headers to match against. ALL of the supplied headers must be matched.
        /// </summary>
        public InputList<Inputs.HttpRouteHeaderMatchArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.HttpRouteHeaderMatchArgs>());
            set => _headers = value;
        }

        /// <summary>
        /// Specifies if prefix_match and full_path_match matches are case sensitive. The default value is false.
        /// </summary>
        [Input("ignoreCase")]
        public Input<bool>? IgnoreCase { get; set; }

        /// <summary>
        /// The HTTP request path value must begin with specified prefix_match. prefix_match must begin with a /. Only one of full_path_match, prefix_match, or regex_match should be used.
        /// </summary>
        [Input("prefixMatch")]
        public Input<string>? PrefixMatch { get; set; }

        [Input("queryParameters")]
        private InputList<Inputs.HttpRouteQueryParameterMatchArgs>? _queryParameters;

        /// <summary>
        /// Specifies a list of query parameters to match against. ALL of the query parameters must be matched.
        /// </summary>
        public InputList<Inputs.HttpRouteQueryParameterMatchArgs> QueryParameters
        {
            get => _queryParameters ?? (_queryParameters = new InputList<Inputs.HttpRouteQueryParameterMatchArgs>());
            set => _queryParameters = value;
        }

        /// <summary>
        /// The HTTP request path value must satisfy the regular expression specified by regex_match after removing any query parameters and anchor supplied with the original URL. For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax Only one of full_path_match, prefix_match, or regex_match should be used.
        /// </summary>
        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        public HttpRouteRouteMatchArgs()
        {
        }
        public static new HttpRouteRouteMatchArgs Empty => new HttpRouteRouteMatchArgs();
    }
}
