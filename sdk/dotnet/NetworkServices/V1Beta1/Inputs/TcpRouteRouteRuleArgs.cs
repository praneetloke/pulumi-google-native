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
    /// Specifies how to match traffic and how to route traffic when traffic is matched.
    /// </summary>
    public sealed class TcpRouteRouteRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The detailed rule defining how to route matched traffic.
        /// </summary>
        [Input("action", required: true)]
        public Input<Inputs.TcpRouteRouteActionArgs> Action { get; set; } = null!;

        [Input("matches")]
        private InputList<Inputs.TcpRouteRouteMatchArgs>? _matches;

        /// <summary>
        /// Optional. RouteMatch defines the predicate used to match requests to a given action. Multiple match types are "OR"ed for evaluation. If no routeMatch field is specified, this rule will unconditionally match traffic.
        /// </summary>
        public InputList<Inputs.TcpRouteRouteMatchArgs> Matches
        {
            get => _matches ?? (_matches = new InputList<Inputs.TcpRouteRouteMatchArgs>());
            set => _matches = value;
        }

        public TcpRouteRouteRuleArgs()
        {
        }
        public static new TcpRouteRouteRuleArgs Empty => new TcpRouteRouteRuleArgs();
    }
}
