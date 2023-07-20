// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified regional BackendService resource.
func LookupRegionBackendService(ctx *pulumi.Context, args *LookupRegionBackendServiceArgs, opts ...pulumi.InvokeOption) (*LookupRegionBackendServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegionBackendServiceResult
	err := ctx.Invoke("google-native:compute/v1:getRegionBackendService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionBackendServiceArgs struct {
	BackendService string  `pulumi:"backendService"`
	Project        *string `pulumi:"project"`
	Region         string  `pulumi:"region"`
}

type LookupRegionBackendServiceResult struct {
	// Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is two weeks (1,209,600). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	AffinityCookieTtlSec int `pulumi:"affinityCookieTtlSec"`
	// The list of backends that serve this BackendService.
	Backends []BackendResponse `pulumi:"backends"`
	// Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
	CdnPolicy       BackendServiceCdnPolicyResponse `pulumi:"cdnPolicy"`
	CircuitBreakers CircuitBreakersResponse         `pulumi:"circuitBreakers"`
	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
	CompressionMode    string                     `pulumi:"compressionMode"`
	ConnectionDraining ConnectionDrainingResponse `pulumi:"connectionDraining"`
	// Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
	ConnectionTrackingPolicy BackendServiceConnectionTrackingPolicyResponse `pulumi:"connectionTrackingPolicy"`
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	ConsistentHash ConsistentHashLoadBalancerSettingsResponse `pulumi:"consistentHash"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	CustomRequestHeaders []string `pulumi:"customRequestHeaders"`
	// Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	CustomResponseHeaders []string `pulumi:"customResponseHeaders"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// The resource URL for the edge security policy associated with this backend service.
	EdgeSecurityPolicy string `pulumi:"edgeSecurityPolicy"`
	// If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
	EnableCDN bool `pulumi:"enableCDN"`
	// Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
	FailoverPolicy BackendServiceFailoverPolicyResponse `pulumi:"failoverPolicy"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a BackendService.
	Fingerprint string `pulumi:"fingerprint"`
	// The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
	HealthChecks []string `pulumi:"healthChecks"`
	// The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
	Iap BackendServiceIAPResponse `pulumi:"iap"`
	// Type of resource. Always compute#backendService for backend services.
	Kind string `pulumi:"kind"`
	// Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
	LoadBalancingScheme string `pulumi:"loadBalancingScheme"`
	// A list of locality load-balancing policies to be used in order of preference. When you use localityLbPolicies, you must set at least one value for either the localityLbPolicies[].policy or the localityLbPolicies[].customPolicy field. localityLbPolicies overrides any value set in the localityLbPolicy field. For an example of how to use this field, see Define a list of preferred policies. Caution: This field and its children are intended for use in a service mesh that includes gRPC clients only. Envoy proxies can't use backend services that have this configuration.
	LocalityLbPolicies []BackendServiceLocalityLoadBalancingPolicyConfigResponse `pulumi:"localityLbPolicies"`
	// The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	LocalityLbPolicy string `pulumi:"localityLbPolicy"`
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
	LogConfig BackendServiceLogConfigResponse `pulumi:"logConfig"`
	// Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
	MaxStreamDuration DurationResponse `pulumi:"maxStreamDuration"`
	// Deployment metadata associated with the resource to be set by a GKE hub controller and read by the backend RCTH
	Metadatas map[string]string `pulumi:"metadatas"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network string `pulumi:"network"`
	// Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2, or GRPC, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	OutlierDetection OutlierDetectionResponse `pulumi:"outlierDetection"`
	// Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
	//
	// Deprecated: Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
	Port int `pulumi:"port"`
	// A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
	PortName string `pulumi:"portName"`
	// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
	Protocol string `pulumi:"protocol"`
	// URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// The resource URL for the security policy associated with this backend service.
	SecurityPolicy string `pulumi:"securityPolicy"`
	// This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	SecuritySettings SecuritySettingsResponse `pulumi:"securitySettings"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// URLs of networkservices.ServiceBinding resources. Can only be set if load balancing scheme is INTERNAL_SELF_MANAGED. If set, lists of backends and health checks must be both empty.
	ServiceBindings []string `pulumi:"serviceBindings"`
	// Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
	SessionAffinity string             `pulumi:"sessionAffinity"`
	Subsetting      SubsettingResponse `pulumi:"subsetting"`
	// The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings. The default is 30 seconds. The full range of timeout values allowed goes from 1 through 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
	TimeoutSec int `pulumi:"timeoutSec"`
}

func LookupRegionBackendServiceOutput(ctx *pulumi.Context, args LookupRegionBackendServiceOutputArgs, opts ...pulumi.InvokeOption) LookupRegionBackendServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionBackendServiceResult, error) {
			args := v.(LookupRegionBackendServiceArgs)
			r, err := LookupRegionBackendService(ctx, &args, opts...)
			var s LookupRegionBackendServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegionBackendServiceResultOutput)
}

type LookupRegionBackendServiceOutputArgs struct {
	BackendService pulumi.StringInput    `pulumi:"backendService"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
	Region         pulumi.StringInput    `pulumi:"region"`
}

func (LookupRegionBackendServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionBackendServiceArgs)(nil)).Elem()
}

type LookupRegionBackendServiceResultOutput struct{ *pulumi.OutputState }

func (LookupRegionBackendServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionBackendServiceResult)(nil)).Elem()
}

func (o LookupRegionBackendServiceResultOutput) ToLookupRegionBackendServiceResultOutput() LookupRegionBackendServiceResultOutput {
	return o
}

func (o LookupRegionBackendServiceResultOutput) ToLookupRegionBackendServiceResultOutputWithContext(ctx context.Context) LookupRegionBackendServiceResultOutput {
	return o
}

// Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is two weeks (1,209,600). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
func (o LookupRegionBackendServiceResultOutput) AffinityCookieTtlSec() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) int { return v.AffinityCookieTtlSec }).(pulumi.IntOutput)
}

// The list of backends that serve this BackendService.
func (o LookupRegionBackendServiceResultOutput) Backends() BackendResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) []BackendResponse { return v.Backends }).(BackendResponseArrayOutput)
}

// Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
func (o LookupRegionBackendServiceResultOutput) CdnPolicy() BackendServiceCdnPolicyResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) BackendServiceCdnPolicyResponse { return v.CdnPolicy }).(BackendServiceCdnPolicyResponseOutput)
}

func (o LookupRegionBackendServiceResultOutput) CircuitBreakers() CircuitBreakersResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) CircuitBreakersResponse { return v.CircuitBreakers }).(CircuitBreakersResponseOutput)
}

// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
func (o LookupRegionBackendServiceResultOutput) CompressionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.CompressionMode }).(pulumi.StringOutput)
}

func (o LookupRegionBackendServiceResultOutput) ConnectionDraining() ConnectionDrainingResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) ConnectionDrainingResponse { return v.ConnectionDraining }).(ConnectionDrainingResponseOutput)
}

// Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
func (o LookupRegionBackendServiceResultOutput) ConnectionTrackingPolicy() BackendServiceConnectionTrackingPolicyResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) BackendServiceConnectionTrackingPolicyResponse {
		return v.ConnectionTrackingPolicy
	}).(BackendServiceConnectionTrackingPolicyResponseOutput)
}

// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
func (o LookupRegionBackendServiceResultOutput) ConsistentHash() ConsistentHashLoadBalancerSettingsResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) ConsistentHashLoadBalancerSettingsResponse {
		return v.ConsistentHash
	}).(ConsistentHashLoadBalancerSettingsResponseOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupRegionBackendServiceResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
func (o LookupRegionBackendServiceResultOutput) CustomRequestHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) []string { return v.CustomRequestHeaders }).(pulumi.StringArrayOutput)
}

// Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
func (o LookupRegionBackendServiceResultOutput) CustomResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) []string { return v.CustomResponseHeaders }).(pulumi.StringArrayOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRegionBackendServiceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The resource URL for the edge security policy associated with this backend service.
func (o LookupRegionBackendServiceResultOutput) EdgeSecurityPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.EdgeSecurityPolicy }).(pulumi.StringOutput)
}

// If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
func (o LookupRegionBackendServiceResultOutput) EnableCDN() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) bool { return v.EnableCDN }).(pulumi.BoolOutput)
}

// Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
func (o LookupRegionBackendServiceResultOutput) FailoverPolicy() BackendServiceFailoverPolicyResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) BackendServiceFailoverPolicyResponse { return v.FailoverPolicy }).(BackendServiceFailoverPolicyResponseOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a BackendService.
func (o LookupRegionBackendServiceResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
func (o LookupRegionBackendServiceResultOutput) HealthChecks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) []string { return v.HealthChecks }).(pulumi.StringArrayOutput)
}

// The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
func (o LookupRegionBackendServiceResultOutput) Iap() BackendServiceIAPResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) BackendServiceIAPResponse { return v.Iap }).(BackendServiceIAPResponseOutput)
}

// Type of resource. Always compute#backendService for backend services.
func (o LookupRegionBackendServiceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
func (o LookupRegionBackendServiceResultOutput) LoadBalancingScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.LoadBalancingScheme }).(pulumi.StringOutput)
}

// A list of locality load-balancing policies to be used in order of preference. When you use localityLbPolicies, you must set at least one value for either the localityLbPolicies[].policy or the localityLbPolicies[].customPolicy field. localityLbPolicies overrides any value set in the localityLbPolicy field. For an example of how to use this field, see Define a list of preferred policies. Caution: This field and its children are intended for use in a service mesh that includes gRPC clients only. Envoy proxies can't use backend services that have this configuration.
func (o LookupRegionBackendServiceResultOutput) LocalityLbPolicies() BackendServiceLocalityLoadBalancingPolicyConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) []BackendServiceLocalityLoadBalancingPolicyConfigResponse {
		return v.LocalityLbPolicies
	}).(BackendServiceLocalityLoadBalancingPolicyConfigResponseArrayOutput)
}

// The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
func (o LookupRegionBackendServiceResultOutput) LocalityLbPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.LocalityLbPolicy }).(pulumi.StringOutput)
}

// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
func (o LookupRegionBackendServiceResultOutput) LogConfig() BackendServiceLogConfigResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) BackendServiceLogConfigResponse { return v.LogConfig }).(BackendServiceLogConfigResponseOutput)
}

// Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
func (o LookupRegionBackendServiceResultOutput) MaxStreamDuration() DurationResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) DurationResponse { return v.MaxStreamDuration }).(DurationResponseOutput)
}

// Deployment metadata associated with the resource to be set by a GKE hub controller and read by the backend RCTH
func (o LookupRegionBackendServiceResultOutput) Metadatas() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) map[string]string { return v.Metadatas }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionBackendServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
func (o LookupRegionBackendServiceResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.Network }).(pulumi.StringOutput)
}

// Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2, or GRPC, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
func (o LookupRegionBackendServiceResultOutput) OutlierDetection() OutlierDetectionResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) OutlierDetectionResponse { return v.OutlierDetection }).(OutlierDetectionResponseOutput)
}

// Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
//
// Deprecated: Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
func (o LookupRegionBackendServiceResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) int { return v.Port }).(pulumi.IntOutput)
}

// A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
func (o LookupRegionBackendServiceResultOutput) PortName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.PortName }).(pulumi.StringOutput)
}

// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
func (o LookupRegionBackendServiceResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupRegionBackendServiceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.Region }).(pulumi.StringOutput)
}

// The resource URL for the security policy associated with this backend service.
func (o LookupRegionBackendServiceResultOutput) SecurityPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.SecurityPolicy }).(pulumi.StringOutput)
}

// This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
func (o LookupRegionBackendServiceResultOutput) SecuritySettings() SecuritySettingsResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) SecuritySettingsResponse { return v.SecuritySettings }).(SecuritySettingsResponseOutput)
}

// Server-defined URL for the resource.
func (o LookupRegionBackendServiceResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// URLs of networkservices.ServiceBinding resources. Can only be set if load balancing scheme is INTERNAL_SELF_MANAGED. If set, lists of backends and health checks must be both empty.
func (o LookupRegionBackendServiceResultOutput) ServiceBindings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) []string { return v.ServiceBindings }).(pulumi.StringArrayOutput)
}

// Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
func (o LookupRegionBackendServiceResultOutput) SessionAffinity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) string { return v.SessionAffinity }).(pulumi.StringOutput)
}

func (o LookupRegionBackendServiceResultOutput) Subsetting() SubsettingResponseOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) SubsettingResponse { return v.Subsetting }).(SubsettingResponseOutput)
}

// The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings. The default is 30 seconds. The full range of timeout values allowed goes from 1 through 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
func (o LookupRegionBackendServiceResultOutput) TimeoutSec() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRegionBackendServiceResult) int { return v.TimeoutSec }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionBackendServiceResultOutput{})
}
