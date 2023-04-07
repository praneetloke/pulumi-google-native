// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified UrlMap resource.
func LookupRegionUrlMap(ctx *pulumi.Context, args *LookupRegionUrlMapArgs, opts ...pulumi.InvokeOption) (*LookupRegionUrlMapResult, error) {
	var rv LookupRegionUrlMapResult
	err := ctx.Invoke("google-native:compute/beta:getRegionUrlMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionUrlMapArgs struct {
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
	UrlMap  string  `pulumi:"urlMap"`
}

type LookupRegionUrlMapResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
	DefaultCustomErrorResponsePolicy CustomErrorResponsePolicyResponse `pulumi:"defaultCustomErrorResponsePolicy"`
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultRouteAction HttpRouteActionResponse `pulumi:"defaultRouteAction"`
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultService string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
	DefaultUrlRedirect HttpRedirectActionResponse `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field is ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a UrlMap.
	Fingerprint string `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction HttpHeaderActionResponse `pulumi:"headerAction"`
	// The list of host rules to use against the URL.
	HostRules []HostRuleResponse `pulumi:"hostRules"`
	// Type of the resource. Always compute#urlMaps for url maps.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers []PathMatcherResponse `pulumi:"pathMatchers"`
	// URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	Tests []UrlMapTestResponse `pulumi:"tests"`
}

func LookupRegionUrlMapOutput(ctx *pulumi.Context, args LookupRegionUrlMapOutputArgs, opts ...pulumi.InvokeOption) LookupRegionUrlMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionUrlMapResult, error) {
			args := v.(LookupRegionUrlMapArgs)
			r, err := LookupRegionUrlMap(ctx, &args, opts...)
			var s LookupRegionUrlMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegionUrlMapResultOutput)
}

type LookupRegionUrlMapOutputArgs struct {
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringInput    `pulumi:"region"`
	UrlMap  pulumi.StringInput    `pulumi:"urlMap"`
}

func (LookupRegionUrlMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionUrlMapArgs)(nil)).Elem()
}

type LookupRegionUrlMapResultOutput struct{ *pulumi.OutputState }

func (LookupRegionUrlMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionUrlMapResult)(nil)).Elem()
}

func (o LookupRegionUrlMapResultOutput) ToLookupRegionUrlMapResultOutput() LookupRegionUrlMapResultOutput {
	return o
}

func (o LookupRegionUrlMapResultOutput) ToLookupRegionUrlMapResultOutputWithContext(ctx context.Context) LookupRegionUrlMapResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupRegionUrlMapResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
func (o LookupRegionUrlMapResultOutput) DefaultCustomErrorResponsePolicy() CustomErrorResponsePolicyResponseOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) CustomErrorResponsePolicyResponse {
		return v.DefaultCustomErrorResponsePolicy
	}).(CustomErrorResponsePolicyResponseOutput)
}

// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
func (o LookupRegionUrlMapResultOutput) DefaultRouteAction() HttpRouteActionResponseOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) HttpRouteActionResponse { return v.DefaultRouteAction }).(HttpRouteActionResponseOutput)
}

// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
func (o LookupRegionUrlMapResultOutput) DefaultService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) string { return v.DefaultService }).(pulumi.StringOutput)
}

// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
func (o LookupRegionUrlMapResultOutput) DefaultUrlRedirect() HttpRedirectActionResponseOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) HttpRedirectActionResponse { return v.DefaultUrlRedirect }).(HttpRedirectActionResponseOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRegionUrlMapResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field is ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a UrlMap.
func (o LookupRegionUrlMapResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
func (o LookupRegionUrlMapResultOutput) HeaderAction() HttpHeaderActionResponseOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) HttpHeaderActionResponse { return v.HeaderAction }).(HttpHeaderActionResponseOutput)
}

// The list of host rules to use against the URL.
func (o LookupRegionUrlMapResultOutput) HostRules() HostRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) []HostRuleResponse { return v.HostRules }).(HostRuleResponseArrayOutput)
}

// Type of the resource. Always compute#urlMaps for url maps.
func (o LookupRegionUrlMapResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionUrlMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of named PathMatchers to use against the URL.
func (o LookupRegionUrlMapResultOutput) PathMatchers() PathMatcherResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) []PathMatcherResponse { return v.PathMatchers }).(PathMatcherResponseArrayOutput)
}

// URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupRegionUrlMapResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupRegionUrlMapResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
func (o LookupRegionUrlMapResultOutput) Tests() UrlMapTestResponseArrayOutput {
	return o.ApplyT(func(v LookupRegionUrlMapResult) []UrlMapTestResponse { return v.Tests }).(UrlMapTestResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionUrlMapResultOutput{})
}
