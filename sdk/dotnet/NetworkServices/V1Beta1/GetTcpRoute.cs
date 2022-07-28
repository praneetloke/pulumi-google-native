// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1
{
    public static class GetTcpRoute
    {
        /// <summary>
        /// Gets details of a single TcpRoute.
        /// </summary>
        public static Task<GetTcpRouteResult> InvokeAsync(GetTcpRouteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTcpRouteResult>("google-native:networkservices/v1beta1:getTcpRoute", args ?? new GetTcpRouteArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single TcpRoute.
        /// </summary>
        public static Output<GetTcpRouteResult> Invoke(GetTcpRouteInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTcpRouteResult>("google-native:networkservices/v1beta1:getTcpRoute", args ?? new GetTcpRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTcpRouteArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("tcpRouteId", required: true)]
        public string TcpRouteId { get; set; } = null!;

        public GetTcpRouteArgs()
        {
        }
        public static new GetTcpRouteArgs Empty => new GetTcpRouteArgs();
    }

    public sealed class GetTcpRouteInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("tcpRouteId", required: true)]
        public Input<string> TcpRouteId { get; set; } = null!;

        public GetTcpRouteInvokeArgs()
        {
        }
        public static new GetTcpRouteInvokeArgs Empty => new GetTcpRouteInvokeArgs();
    }


    [OutputType]
    public sealed class GetTcpRouteResult
    {
        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        /// </summary>
        public readonly ImmutableArray<string> Gateways;
        /// <summary>
        /// Optional. Set of label tags associated with the TcpRoute resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
        /// </summary>
        public readonly ImmutableArray<string> Meshes;
        /// <summary>
        /// Name of the TcpRoute resource. It matches pattern `projects/*/locations/global/tcpRoutes/tcp_route_name&gt;`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
        /// </summary>
        public readonly ImmutableArray<Outputs.TcpRouteRouteRuleResponse> Rules;
        /// <summary>
        /// Server-defined URL of this resource
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetTcpRouteResult(
            string createTime,

            string description,

            ImmutableArray<string> gateways,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> meshes,

            string name,

            ImmutableArray<Outputs.TcpRouteRouteRuleResponse> rules,

            string selfLink,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Gateways = gateways;
            Labels = labels;
            Meshes = meshes;
            Name = name;
            Rules = rules;
            SelfLink = selfLink;
            UpdateTime = updateTime;
        }
    }
}
