// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified UrlMap resource.
 */
export function getRegionUrlMap(args: GetRegionUrlMapArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionUrlMapResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/beta:getRegionUrlMap", {
        "project": args.project,
        "region": args.region,
        "urlMap": args.urlMap,
    }, opts);
}

export interface GetRegionUrlMapArgs {
    project?: string;
    region: string;
    urlMap: string;
}

export interface GetRegionUrlMapResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
     */
    readonly defaultCustomErrorResponsePolicy: outputs.compute.beta.CustomErrorResponsePolicyResponse;
    /**
     * defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
     */
    readonly defaultRouteAction: outputs.compute.beta.HttpRouteActionResponse;
    /**
     * The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
     */
    readonly defaultService: string;
    /**
     * When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
     */
    readonly defaultUrlRedirect: outputs.compute.beta.HttpRedirectActionResponse;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field is ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a UrlMap.
     */
    readonly fingerprint: string;
    /**
     * Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly headerAction: outputs.compute.beta.HttpHeaderActionResponse;
    /**
     * The list of host rules to use against the URL.
     */
    readonly hostRules: outputs.compute.beta.HostRuleResponse[];
    /**
     * Type of the resource. Always compute#urlMaps for url maps.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The list of named PathMatchers to use against the URL.
     */
    readonly pathMatchers: outputs.compute.beta.PathMatcherResponse[];
    /**
     * URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly tests: outputs.compute.beta.UrlMapTestResponse[];
}
/**
 * Returns the specified UrlMap resource.
 */
export function getRegionUrlMapOutput(args: GetRegionUrlMapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionUrlMapResult> {
    return pulumi.output(args).apply((a: any) => getRegionUrlMap(a, opts))
}

export interface GetRegionUrlMapOutputArgs {
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    urlMap: pulumi.Input<string>;
}
