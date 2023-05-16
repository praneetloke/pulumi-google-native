// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// The spec for modifying the path before sending the request to the matched backend service.
    /// </summary>
    [OutputType]
    public sealed class UrlRewriteResponse
    {
        /// <summary>
        /// Before forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite. The value must be from 1 to 255 characters.
        /// </summary>
        public readonly string HostRewrite;
        /// <summary>
        /// Before forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite. The value must be from 1 to 1024 characters.
        /// </summary>
        public readonly string PathPrefixRewrite;
        /// <summary>
        ///  If specified, the pattern rewrites the URL path (based on the :path header) using the HTTP template syntax. A corresponding path_template_match must be specified. Any template variables must exist in the path_template_match field. - -At least one variable must be specified in the path_template_match field - You can omit variables from the rewritten URL - The * and ** operators cannot be matched unless they have a corresponding variable name - e.g. {format=*} or {var=**}. For example, a path_template_match of /static/{format=**} could be rewritten as /static/content/{format} to prefix /content to the URL. Variables can also be re-ordered in a rewrite, so that /{country}/{format}/{suffix=**} can be rewritten as /content/{format}/{country}/{suffix}. At least one non-empty routeRules[].matchRules[].path_template_match is required. Only one of path_prefix_rewrite or path_template_rewrite may be specified.
        /// </summary>
        public readonly string PathTemplateRewrite;

        [OutputConstructor]
        private UrlRewriteResponse(
            string hostRewrite,

            string pathPrefixRewrite,

            string pathTemplateRewrite)
        {
            HostRewrite = hostRewrite;
            PathPrefixRewrite = pathPrefixRewrite;
            PathTemplateRewrite = pathTemplateRewrite;
        }
    }
}
