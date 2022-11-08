// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// HttpRouteRuleMatch criteria for a request's query parameter.
    /// </summary>
    [OutputType]
    public sealed class HttpQueryParameterMatchResponse
    {
        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter exactly matches the contents of exactMatch. Only one of presentMatch, exactMatch, or regexMatch must be set. 
        /// </summary>
        public readonly string ExactMatch;
        /// <summary>
        /// The name of the query parameter to match. The query parameter must exist in the request, in the absence of which the request match fails.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies that the queryParameterMatch matches if the request contains the query parameter, irrespective of whether the parameter has a value or not. Only one of presentMatch, exactMatch, or regexMatch must be set. 
        /// </summary>
        public readonly bool PresentMatch;
        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter matches the regular expression specified by regexMatch. For more information about regular expression syntax, see Syntax. Only one of presentMatch, exactMatch, or regexMatch must be set. Regular expressions can only be used when the loadBalancingScheme is set to INTERNAL_SELF_MANAGED. 
        /// </summary>
        public readonly string RegexMatch;

        [OutputConstructor]
        private HttpQueryParameterMatchResponse(
            string exactMatch,

            string name,

            bool presentMatch,

            string regexMatch)
        {
            ExactMatch = exactMatch;
            Name = name;
            PresentMatch = presentMatch;
            RegexMatch = regexMatch;
        }
    }
}
