// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// The authentication parameters to provide to the specified resource or URL that requires a username and password. Currently, only Basic HTTP authentication (https://tools.ietf.org/html/rfc7617) is supported in Uptime checks.
    /// </summary>
    public sealed class BasicAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password to use when authenticating with the HTTP server.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The username to use when authenticating with the HTTP server.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public BasicAuthenticationArgs()
        {
        }
        public static new BasicAuthenticationArgs Empty => new BasicAuthenticationArgs();
    }
}
