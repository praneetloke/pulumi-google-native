// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// A path-matching rule for a URL. If matched, will use the specified BackendService to handle the traffic arriving at this URL.
    /// </summary>
    public sealed class PathRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("paths")]
        private InputList<string>? _paths;

        /// <summary>
        /// The list of path patterns to match. Each must start with / and the only place a * is allowed is at the end following a /. The string fed to the path matcher does not include any text after the first ? or #, and those chars are not allowed here.
        /// </summary>
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        /// <summary>
        /// In response to a matching path, the load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If routeAction specifies any weightedBackendServices, service must not be set. Conversely if service is set, routeAction cannot contain any weightedBackendServices. Only one of routeAction or urlRedirect must be set. URL maps for external HTTP(S) load balancers support only the urlRewrite action within a path rule's routeAction.
        /// </summary>
        [Input("routeAction")]
        public Input<Inputs.HttpRouteActionArgs>? RouteAction { get; set; }

        /// <summary>
        /// The full or partial URL of the backend service resource to which traffic is directed if this rule is matched. If routeAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if service is specified, routeAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of urlRedirect, service or routeAction.weightedBackendService must be set.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// When a path pattern is matched, the request is redirected to a URL specified by urlRedirect. If urlRedirect is specified, service or routeAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
        /// </summary>
        [Input("urlRedirect")]
        public Input<Inputs.HttpRedirectActionArgs>? UrlRedirect { get; set; }

        public PathRuleArgs()
        {
        }
        public static new PathRuleArgs Empty => new PathRuleArgs();
    }
}
