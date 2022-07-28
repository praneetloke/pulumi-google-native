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
    /// RouteMatch defines the predicate used to match requests to a given action. Multiple match types are "OR"ed for evaluation. If no routeMatch field is specified, this rule will unconditionally match traffic.
    /// </summary>
    public sealed class TcpRouteRouteMatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Must be specified in the CIDR range format. A CIDR range consists of an IP Address and a prefix length to construct the subnet mask. By default, the prefix length is 32 (i.e. matches a single IP address). Only IPV4 addresses are supported. Examples: "10.0.0.1" - matches against this exact IP address. "10.0.0.0/8" - matches against any IP address within the 10.0.0.0 subnet and 255.255.255.0 mask. "0.0.0.0/0" - matches against any IP address'.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Specifies the destination port to match against.
        /// </summary>
        [Input("port", required: true)]
        public Input<string> Port { get; set; } = null!;

        public TcpRouteRouteMatchArgs()
        {
        }
        public static new TcpRouteRouteMatchArgs Empty => new TcpRouteRouteMatchArgs();
    }
}
