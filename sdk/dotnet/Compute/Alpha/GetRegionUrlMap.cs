// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetRegionUrlMap
    {
        /// <summary>
        /// Returns the specified UrlMap resource. Gets a list of available URL maps by making a list() request.
        /// </summary>
        public static Task<GetRegionUrlMapResult> InvokeAsync(GetRegionUrlMapArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegionUrlMapResult>("google-native:compute/alpha:getRegionUrlMap", args ?? new GetRegionUrlMapArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified UrlMap resource. Gets a list of available URL maps by making a list() request.
        /// </summary>
        public static Output<GetRegionUrlMapResult> Invoke(GetRegionUrlMapInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionUrlMapResult>("google-native:compute/alpha:getRegionUrlMap", args ?? new GetRegionUrlMapInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionUrlMapArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("urlMap", required: true)]
        public string UrlMap { get; set; } = null!;

        public GetRegionUrlMapArgs()
        {
        }
        public static new GetRegionUrlMapArgs Empty => new GetRegionUrlMapArgs();
    }

    public sealed class GetRegionUrlMapInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("urlMap", required: true)]
        public Input<string> UrlMap { get; set; } = null!;

        public GetRegionUrlMapInvokeArgs()
        {
        }
        public static new GetRegionUrlMapInvokeArgs Empty => new GetRegionUrlMapInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionUrlMapResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
        /// </summary>
        public readonly Outputs.CustomErrorResponsePolicyResponse DefaultCustomErrorResponsePolicy;
        /// <summary>
        /// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
        /// </summary>
        public readonly Outputs.HttpRouteActionResponse DefaultRouteAction;
        /// <summary>
        /// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
        /// </summary>
        public readonly string DefaultService;
        /// <summary>
        /// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
        /// </summary>
        public readonly Outputs.HttpRedirectActionResponse DefaultUrlRedirect;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field is ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a UrlMap.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        public readonly Outputs.HttpHeaderActionResponse HeaderAction;
        /// <summary>
        /// The list of host rules to use against the URL.
        /// </summary>
        public readonly ImmutableArray<Outputs.HostRuleResponse> HostRules;
        /// <summary>
        /// Type of the resource. Always compute#urlMaps for url maps.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of named PathMatchers to use against the URL.
        /// </summary>
        public readonly ImmutableArray<Outputs.PathMatcherResponse> PathMatchers;
        /// <summary>
        /// URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        public readonly ImmutableArray<Outputs.UrlMapTestResponse> Tests;

        [OutputConstructor]
        private GetRegionUrlMapResult(
            string creationTimestamp,

            Outputs.CustomErrorResponsePolicyResponse defaultCustomErrorResponsePolicy,

            Outputs.HttpRouteActionResponse defaultRouteAction,

            string defaultService,

            Outputs.HttpRedirectActionResponse defaultUrlRedirect,

            string description,

            string fingerprint,

            Outputs.HttpHeaderActionResponse headerAction,

            ImmutableArray<Outputs.HostRuleResponse> hostRules,

            string kind,

            string name,

            ImmutableArray<Outputs.PathMatcherResponse> pathMatchers,

            string region,

            string selfLink,

            ImmutableArray<Outputs.UrlMapTestResponse> tests)
        {
            CreationTimestamp = creationTimestamp;
            DefaultCustomErrorResponsePolicy = defaultCustomErrorResponsePolicy;
            DefaultRouteAction = defaultRouteAction;
            DefaultService = defaultService;
            DefaultUrlRedirect = defaultUrlRedirect;
            Description = description;
            Fingerprint = fingerprint;
            HeaderAction = headerAction;
            HostRules = hostRules;
            Kind = kind;
            Name = name;
            PathMatchers = pathMatchers;
            Region = region;
            SelfLink = selfLink;
            Tests = tests;
        }
    }
}
