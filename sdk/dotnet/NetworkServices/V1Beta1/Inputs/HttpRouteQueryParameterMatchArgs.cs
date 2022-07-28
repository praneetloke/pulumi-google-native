// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Inputs
{

    /// <summary>
    /// Specifications to match a query parameter in the request.
    /// </summary>
    public sealed class HttpRouteQueryParameterMatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value of the query parameter must exactly match the contents of exact_match. Only one of exact_match, regex_match, or present_match must be set.
        /// </summary>
        [Input("exactMatch")]
        public Input<string>? ExactMatch { get; set; }

        /// <summary>
        /// Specifies that the QueryParameterMatcher matches if request contains query parameter, irrespective of whether the parameter has a value or not. Only one of exact_match, regex_match, or present_match must be set.
        /// </summary>
        [Input("presentMatch")]
        public Input<bool>? PresentMatch { get; set; }

        /// <summary>
        /// The name of the query parameter to match.
        /// </summary>
        [Input("queryParameter")]
        public Input<string>? QueryParameter { get; set; }

        /// <summary>
        /// The value of the query parameter must match the regular expression specified by regex_match. For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax Only one of exact_match, regex_match, or present_match must be set.
        /// </summary>
        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        public HttpRouteQueryParameterMatchArgs()
        {
        }
        public static new HttpRouteQueryParameterMatchArgs Empty => new HttpRouteQueryParameterMatchArgs();
    }
}
