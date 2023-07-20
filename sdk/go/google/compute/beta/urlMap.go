// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a UrlMap resource in the specified project using the data included in the request.
type UrlMap struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
	DefaultCustomErrorResponsePolicy CustomErrorResponsePolicyResponseOutput `pulumi:"defaultCustomErrorResponsePolicy"`
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultRouteAction HttpRouteActionResponseOutput `pulumi:"defaultRouteAction"`
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultService pulumi.StringOutput `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
	DefaultUrlRedirect HttpRedirectActionResponseOutput `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field is ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a UrlMap.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction HttpHeaderActionResponseOutput `pulumi:"headerAction"`
	// The list of host rules to use against the URL.
	HostRules HostRuleResponseArrayOutput `pulumi:"hostRules"`
	// Type of the resource. Always compute#urlMaps for url maps.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers PathMatcherResponseArrayOutput `pulumi:"pathMatchers"`
	Project      pulumi.StringOutput            `pulumi:"project"`
	// URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	Tests UrlMapTestResponseArrayOutput `pulumi:"tests"`
}

// NewUrlMap registers a new resource with the given unique name, arguments, and options.
func NewUrlMap(ctx *pulumi.Context,
	name string, args *UrlMapArgs, opts ...pulumi.ResourceOption) (*UrlMap, error) {
	if args == nil {
		args = &UrlMapArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UrlMap
	err := ctx.RegisterResource("google-native:compute/beta:UrlMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUrlMap gets an existing UrlMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUrlMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UrlMapState, opts ...pulumi.ResourceOption) (*UrlMap, error) {
	var resource UrlMap
	err := ctx.ReadResource("google-native:compute/beta:UrlMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UrlMap resources.
type urlMapState struct {
}

type UrlMapState struct {
}

func (UrlMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*urlMapState)(nil)).Elem()
}

type urlMapArgs struct {
	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
	DefaultCustomErrorResponsePolicy *CustomErrorResponsePolicy `pulumi:"defaultCustomErrorResponsePolicy"`
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultRouteAction *HttpRouteAction `pulumi:"defaultRouteAction"`
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
	DefaultUrlRedirect *HttpRedirectAction `pulumi:"defaultUrlRedirect"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction *HttpHeaderAction `pulumi:"headerAction"`
	// The list of host rules to use against the URL.
	HostRules []HostRule `pulumi:"hostRules"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The list of named PathMatchers to use against the URL.
	PathMatchers []PathMatcher `pulumi:"pathMatchers"`
	Project      *string       `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	Tests []UrlMapTest `pulumi:"tests"`
}

// The set of arguments for constructing a UrlMap resource.
type UrlMapArgs struct {
	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
	DefaultCustomErrorResponsePolicy CustomErrorResponsePolicyPtrInput
	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultRouteAction HttpRouteActionPtrInput
	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
	DefaultUrlRedirect HttpRedirectActionPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	HeaderAction HttpHeaderActionPtrInput
	// The list of host rules to use against the URL.
	HostRules HostRuleArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The list of named PathMatchers to use against the URL.
	PathMatchers PathMatcherArrayInput
	Project      pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	Tests UrlMapTestArrayInput
}

func (UrlMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*urlMapArgs)(nil)).Elem()
}

type UrlMapInput interface {
	pulumi.Input

	ToUrlMapOutput() UrlMapOutput
	ToUrlMapOutputWithContext(ctx context.Context) UrlMapOutput
}

func (*UrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((**UrlMap)(nil)).Elem()
}

func (i *UrlMap) ToUrlMapOutput() UrlMapOutput {
	return i.ToUrlMapOutputWithContext(context.Background())
}

func (i *UrlMap) ToUrlMapOutputWithContext(ctx context.Context) UrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlMapOutput)
}

type UrlMapOutput struct{ *pulumi.OutputState }

func (UrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UrlMap)(nil)).Elem()
}

func (o UrlMapOutput) ToUrlMapOutput() UrlMapOutput {
	return o
}

func (o UrlMapOutput) ToUrlMapOutputWithContext(ctx context.Context) UrlMapOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o UrlMapOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the Load Balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for Global External HTTP(S) load balancing.
func (o UrlMapOutput) DefaultCustomErrorResponsePolicy() CustomErrorResponsePolicyResponseOutput {
	return o.ApplyT(func(v *UrlMap) CustomErrorResponsePolicyResponseOutput { return v.DefaultCustomErrorResponsePolicy }).(CustomErrorResponsePolicyResponseOutput)
}

// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices. Only one of defaultRouteAction or defaultUrlRedirect must be set. URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
func (o UrlMapOutput) DefaultRouteAction() HttpRouteActionResponseOutput {
	return o.ApplyT(func(v *UrlMap) HttpRouteActionResponseOutput { return v.DefaultRouteAction }).(HttpRouteActionResponseOutput)
}

// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. However, if defaultService is specified, defaultRouteAction cannot contain any weightedBackendServices. Conversely, if routeAction specifies any weightedBackendServices, service must not be specified. Only one of defaultService, defaultUrlRedirect , or defaultRouteAction.weightedBackendService must be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
func (o UrlMapOutput) DefaultService() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.DefaultService }).(pulumi.StringOutput)
}

// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or defaultRouteAction must not be set. Not supported when the URL map is bound to a target gRPC proxy.
func (o UrlMapOutput) DefaultUrlRedirect() HttpRedirectActionResponseOutput {
	return o.ApplyT(func(v *UrlMap) HttpRedirectActionResponseOutput { return v.DefaultUrlRedirect }).(HttpRedirectActionResponseOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o UrlMapOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field is ignored when inserting a UrlMap. An up-to-date fingerprint must be provided in order to update the UrlMap, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a UrlMap.
func (o UrlMapOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
func (o UrlMapOutput) HeaderAction() HttpHeaderActionResponseOutput {
	return o.ApplyT(func(v *UrlMap) HttpHeaderActionResponseOutput { return v.HeaderAction }).(HttpHeaderActionResponseOutput)
}

// The list of host rules to use against the URL.
func (o UrlMapOutput) HostRules() HostRuleResponseArrayOutput {
	return o.ApplyT(func(v *UrlMap) HostRuleResponseArrayOutput { return v.HostRules }).(HostRuleResponseArrayOutput)
}

// Type of the resource. Always compute#urlMaps for url maps.
func (o UrlMapOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o UrlMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of named PathMatchers to use against the URL.
func (o UrlMapOutput) PathMatchers() PathMatcherResponseArrayOutput {
	return o.ApplyT(func(v *UrlMap) PathMatcherResponseArrayOutput { return v.PathMatchers }).(PathMatcherResponseArrayOutput)
}

func (o UrlMapOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// URL of the region where the regional URL map resides. This field is not applicable to global URL maps. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o UrlMapOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o UrlMapOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o UrlMapOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlMap) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
func (o UrlMapOutput) Tests() UrlMapTestResponseArrayOutput {
	return o.ApplyT(func(v *UrlMap) UrlMapTestResponseArrayOutput { return v.Tests }).(UrlMapTestResponseArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UrlMapInput)(nil)).Elem(), &UrlMap{})
	pulumi.RegisterOutputType(UrlMapOutput{})
}
