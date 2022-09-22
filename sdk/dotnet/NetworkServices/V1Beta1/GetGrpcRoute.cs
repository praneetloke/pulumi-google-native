// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1
{
    public static class GetGrpcRoute
    {
        /// <summary>
        /// Gets details of a single GrpcRoute.
        /// </summary>
        public static Task<GetGrpcRouteResult> InvokeAsync(GetGrpcRouteArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGrpcRouteResult>("google-native:networkservices/v1beta1:getGrpcRoute", args ?? new GetGrpcRouteArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single GrpcRoute.
        /// </summary>
        public static Output<GetGrpcRouteResult> Invoke(GetGrpcRouteInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGrpcRouteResult>("google-native:networkservices/v1beta1:getGrpcRoute", args ?? new GetGrpcRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGrpcRouteArgs : global::Pulumi.InvokeArgs
    {
        [Input("grpcRouteId", required: true)]
        public string GrpcRouteId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGrpcRouteArgs()
        {
        }
        public static new GetGrpcRouteArgs Empty => new GetGrpcRouteArgs();
    }

    public sealed class GetGrpcRouteInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("grpcRouteId", required: true)]
        public Input<string> GrpcRouteId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGrpcRouteInvokeArgs()
        {
        }
        public static new GetGrpcRouteInvokeArgs Empty => new GetGrpcRouteInvokeArgs();
    }


    [OutputType]
    public sealed class GetGrpcRouteResult
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
        /// Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        /// </summary>
        public readonly ImmutableArray<string> Gateways;
        /// <summary>
        /// Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.example.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateway must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
        /// </summary>
        public readonly ImmutableArray<string> Hostnames;
        /// <summary>
        /// Optional. Set of label tags associated with the GrpcRoute resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`
        /// </summary>
        public readonly ImmutableArray<string> Meshes;
        /// <summary>
        /// Name of the GrpcRoute resource. It matches pattern `projects/*/locations/global/grpcRoutes/`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
        /// </summary>
        public readonly ImmutableArray<Outputs.GrpcRouteRouteRuleResponse> Rules;
        /// <summary>
        /// Server-defined URL of this resource
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetGrpcRouteResult(
            string createTime,

            string description,

            ImmutableArray<string> gateways,

            ImmutableArray<string> hostnames,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> meshes,

            string name,

            ImmutableArray<Outputs.GrpcRouteRouteRuleResponse> rules,

            string selfLink,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Gateways = gateways;
            Hostnames = hostnames;
            Labels = labels;
            Meshes = meshes;
            Name = name;
            Rules = rules;
            SelfLink = selfLink;
            UpdateTime = updateTime;
        }
    }
}
