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
    /// The specification for allowing client-side cross-origin requests. For more information about the W3C recommendation for cross-origin resource sharing (CORS), see Fetch API Living Standard.
    /// </summary>
    public sealed class CorsPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// In response to a preflight request, setting this to true indicates that the actual request can include user credentials. This field translates to the Access-Control-Allow-Credentials header. Default is false.
        /// </summary>
        [Input("allowCredentials")]
        public Input<bool>? AllowCredentials { get; set; }

        [Input("allowHeaders")]
        private InputList<string>? _allowHeaders;

        /// <summary>
        /// Specifies the content for the Access-Control-Allow-Headers header.
        /// </summary>
        public InputList<string> AllowHeaders
        {
            get => _allowHeaders ?? (_allowHeaders = new InputList<string>());
            set => _allowHeaders = value;
        }

        [Input("allowMethods")]
        private InputList<string>? _allowMethods;

        /// <summary>
        /// Specifies the content for the Access-Control-Allow-Methods header.
        /// </summary>
        public InputList<string> AllowMethods
        {
            get => _allowMethods ?? (_allowMethods = new InputList<string>());
            set => _allowMethods = value;
        }

        [Input("allowOriginRegexes")]
        private InputList<string>? _allowOriginRegexes;

        /// <summary>
        /// Specifies a regular expression that matches allowed origins. For more information about the regular expression syntax, see Syntax. An origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes. Regular expressions can only be used when the loadBalancingScheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        public InputList<string> AllowOriginRegexes
        {
            get => _allowOriginRegexes ?? (_allowOriginRegexes = new InputList<string>());
            set => _allowOriginRegexes = value;
        }

        [Input("allowOrigins")]
        private InputList<string>? _allowOrigins;

        /// <summary>
        /// Specifies the list of origins that is allowed to do CORS requests. An origin is allowed if it matches either an item in allowOrigins or an item in allowOriginRegexes.
        /// </summary>
        public InputList<string> AllowOrigins
        {
            get => _allowOrigins ?? (_allowOrigins = new InputList<string>());
            set => _allowOrigins = value;
        }

        /// <summary>
        /// If true, the setting specifies the CORS policy is disabled. The default value of false, which indicates that the CORS policy is in effect.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("exposeHeaders")]
        private InputList<string>? _exposeHeaders;

        /// <summary>
        /// Specifies the content for the Access-Control-Expose-Headers header.
        /// </summary>
        public InputList<string> ExposeHeaders
        {
            get => _exposeHeaders ?? (_exposeHeaders = new InputList<string>());
            set => _exposeHeaders = value;
        }

        /// <summary>
        /// Specifies how long results of a preflight request can be cached in seconds. This field translates to the Access-Control-Max-Age header.
        /// </summary>
        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        public CorsPolicyArgs()
        {
        }
        public static new CorsPolicyArgs Empty => new CorsPolicyArgs();
    }
}
