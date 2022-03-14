// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a regional BackendService resource in the specified project using the data included in the request. For more information, see Backend services overview.
 */
export class RegionBackendService extends pulumi.CustomResource {
    /**
     * Get an existing RegionBackendService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionBackendService {
        return new RegionBackendService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/v1:RegionBackendService';

    /**
     * Returns true if the given object is an instance of RegionBackendService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionBackendService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionBackendService.__pulumiType;
    }

    /**
     * Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day (86,400). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly affinityCookieTtlSec!: pulumi.Output<number>;
    /**
     * The list of backends that serve this BackendService.
     */
    public readonly backends!: pulumi.Output<outputs.compute.v1.BackendResponse[]>;
    /**
     * Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
     */
    public readonly cdnPolicy!: pulumi.Output<outputs.compute.v1.BackendServiceCdnPolicyResponse>;
    public readonly circuitBreakers!: pulumi.Output<outputs.compute.v1.CircuitBreakersResponse>;
    public readonly connectionDraining!: pulumi.Output<outputs.compute.v1.ConnectionDrainingResponse>;
    /**
     * Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
     */
    public readonly connectionTrackingPolicy!: pulumi.Output<outputs.compute.v1.BackendServiceConnectionTrackingPolicyResponse>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. 
     */
    public readonly consistentHash!: pulumi.Output<outputs.compute.v1.ConsistentHashLoadBalancerSettingsResponse>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
     */
    public readonly customRequestHeaders!: pulumi.Output<string[]>;
    /**
     * Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
     */
    public readonly customResponseHeaders!: pulumi.Output<string[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The resource URL for the edge security policy associated with this backend service.
     */
    public /*out*/ readonly edgeSecurityPolicy!: pulumi.Output<string>;
    /**
     * If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
     */
    public readonly enableCDN!: pulumi.Output<boolean>;
    /**
     * Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
     */
    public readonly failoverPolicy!: pulumi.Output<outputs.compute.v1.BackendServiceFailoverPolicyResponse>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a BackendService.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
     */
    public readonly healthChecks!: pulumi.Output<string[]>;
    /**
     * The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
     */
    public readonly iap!: pulumi.Output<outputs.compute.v1.BackendServiceIAPResponse>;
    /**
     * Type of resource. Always compute#backendService for backend services.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
     */
    public readonly loadBalancingScheme!: pulumi.Output<string>;
    /**
     * The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly localityLbPolicy!: pulumi.Output<string>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.v1.BackendServiceLogConfigResponse>;
    /**
     * Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
     */
    public readonly maxStreamDuration!: pulumi.Output<outputs.compute.v1.DurationResponse>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    public readonly outlierDetection!: pulumi.Output<outputs.compute.v1.OutlierDetectionResponse>;
    /**
     * A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
     */
    public readonly portName!: pulumi.Output<string>;
    /**
     * The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The resource URL for the security policy associated with this backend service.
     */
    public /*out*/ readonly securityPolicy!: pulumi.Output<string>;
    /**
     * This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
     */
    public readonly securitySettings!: pulumi.Output<outputs.compute.v1.SecuritySettingsResponse>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
     */
    public readonly sessionAffinity!: pulumi.Output<string>;
    public readonly subsetting!: pulumi.Output<outputs.compute.v1.SubsettingResponse>;
    /**
     * The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings The default is 30 seconds. The full range of timeout values allowed is 1 - 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
     */
    public readonly timeoutSec!: pulumi.Output<number>;

    /**
     * Create a RegionBackendService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionBackendServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["affinityCookieTtlSec"] = args ? args.affinityCookieTtlSec : undefined;
            resourceInputs["backends"] = args ? args.backends : undefined;
            resourceInputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            resourceInputs["circuitBreakers"] = args ? args.circuitBreakers : undefined;
            resourceInputs["connectionDraining"] = args ? args.connectionDraining : undefined;
            resourceInputs["connectionTrackingPolicy"] = args ? args.connectionTrackingPolicy : undefined;
            resourceInputs["consistentHash"] = args ? args.consistentHash : undefined;
            resourceInputs["customRequestHeaders"] = args ? args.customRequestHeaders : undefined;
            resourceInputs["customResponseHeaders"] = args ? args.customResponseHeaders : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableCDN"] = args ? args.enableCDN : undefined;
            resourceInputs["failoverPolicy"] = args ? args.failoverPolicy : undefined;
            resourceInputs["healthChecks"] = args ? args.healthChecks : undefined;
            resourceInputs["iap"] = args ? args.iap : undefined;
            resourceInputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            resourceInputs["localityLbPolicy"] = args ? args.localityLbPolicy : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["maxStreamDuration"] = args ? args.maxStreamDuration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["outlierDetection"] = args ? args.outlierDetection : undefined;
            resourceInputs["portName"] = args ? args.portName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["securitySettings"] = args ? args.securitySettings : undefined;
            resourceInputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            resourceInputs["subsetting"] = args ? args.subsetting : undefined;
            resourceInputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["edgeSecurityPolicy"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["securityPolicy"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        } else {
            resourceInputs["affinityCookieTtlSec"] = undefined /*out*/;
            resourceInputs["backends"] = undefined /*out*/;
            resourceInputs["cdnPolicy"] = undefined /*out*/;
            resourceInputs["circuitBreakers"] = undefined /*out*/;
            resourceInputs["connectionDraining"] = undefined /*out*/;
            resourceInputs["connectionTrackingPolicy"] = undefined /*out*/;
            resourceInputs["consistentHash"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["customRequestHeaders"] = undefined /*out*/;
            resourceInputs["customResponseHeaders"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["edgeSecurityPolicy"] = undefined /*out*/;
            resourceInputs["enableCDN"] = undefined /*out*/;
            resourceInputs["failoverPolicy"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["healthChecks"] = undefined /*out*/;
            resourceInputs["iap"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["loadBalancingScheme"] = undefined /*out*/;
            resourceInputs["localityLbPolicy"] = undefined /*out*/;
            resourceInputs["logConfig"] = undefined /*out*/;
            resourceInputs["maxStreamDuration"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["outlierDetection"] = undefined /*out*/;
            resourceInputs["portName"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["securityPolicy"] = undefined /*out*/;
            resourceInputs["securitySettings"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["sessionAffinity"] = undefined /*out*/;
            resourceInputs["subsetting"] = undefined /*out*/;
            resourceInputs["timeoutSec"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionBackendService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionBackendService resource.
 */
export interface RegionBackendServiceArgs {
    /**
     * Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is one day (86,400). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    affinityCookieTtlSec?: pulumi.Input<number>;
    /**
     * The list of backends that serve this BackendService.
     */
    backends?: pulumi.Input<pulumi.Input<inputs.compute.v1.BackendArgs>[]>;
    /**
     * Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
     */
    cdnPolicy?: pulumi.Input<inputs.compute.v1.BackendServiceCdnPolicyArgs>;
    circuitBreakers?: pulumi.Input<inputs.compute.v1.CircuitBreakersArgs>;
    connectionDraining?: pulumi.Input<inputs.compute.v1.ConnectionDrainingArgs>;
    /**
     * Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
     */
    connectionTrackingPolicy?: pulumi.Input<inputs.compute.v1.BackendServiceConnectionTrackingPolicyArgs>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. 
     */
    consistentHash?: pulumi.Input<inputs.compute.v1.ConsistentHashLoadBalancerSettingsArgs>;
    /**
     * Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
     */
    customRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
     */
    customResponseHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
     */
    enableCDN?: pulumi.Input<boolean>;
    /**
     * Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
     */
    failoverPolicy?: pulumi.Input<inputs.compute.v1.BackendServiceFailoverPolicyArgs>;
    /**
     * The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
     */
    healthChecks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
     */
    iap?: pulumi.Input<inputs.compute.v1.BackendServiceIAPArgs>;
    /**
     * Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
     */
    loadBalancingScheme?: pulumi.Input<enums.compute.v1.RegionBackendServiceLoadBalancingScheme>;
    /**
     * The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    localityLbPolicy?: pulumi.Input<enums.compute.v1.RegionBackendServiceLocalityLbPolicy>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
     */
    logConfig?: pulumi.Input<inputs.compute.v1.BackendServiceLogConfigArgs>;
    /**
     * Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
     */
    maxStreamDuration?: pulumi.Input<inputs.compute.v1.DurationArgs>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    network?: pulumi.Input<string>;
    /**
     * Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    outlierDetection?: pulumi.Input<inputs.compute.v1.OutlierDetectionArgs>;
    /**
     * A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
     */
    portName?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
     */
    protocol?: pulumi.Input<enums.compute.v1.RegionBackendServiceProtocol>;
    region: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
     */
    securitySettings?: pulumi.Input<inputs.compute.v1.SecuritySettingsArgs>;
    /**
     * Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
     */
    sessionAffinity?: pulumi.Input<enums.compute.v1.RegionBackendServiceSessionAffinity>;
    subsetting?: pulumi.Input<inputs.compute.v1.SubsettingArgs>;
    /**
     * The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings The default is 30 seconds. The full range of timeout values allowed is 1 - 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
     */
    timeoutSec?: pulumi.Input<number>;
}
