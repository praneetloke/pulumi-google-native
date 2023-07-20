// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single EndpointPolicy.
func LookupEndpointPolicy(ctx *pulumi.Context, args *LookupEndpointPolicyArgs, opts ...pulumi.InvokeOption) (*LookupEndpointPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointPolicyResult
	err := ctx.Invoke("google-native:networkservices/v1beta1:getEndpointPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointPolicyArgs struct {
	EndpointPolicyId string  `pulumi:"endpointPolicyId"`
	Location         string  `pulumi:"location"`
	Project          *string `pulumi:"project"`
}

type LookupEndpointPolicyResult struct {
	// Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
	AuthorizationPolicy string `pulumi:"authorizationPolicy"`
	// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
	ClientTlsPolicy string `pulumi:"clientTlsPolicy"`
	// The timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `pulumi:"description"`
	// A matcher that selects endpoints to which the policies should be applied.
	EndpointMatcher EndpointMatcherResponse `pulumi:"endpointMatcher"`
	// Optional. Set of label tags associated with the EndpointPolicy resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
	Name string `pulumi:"name"`
	// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
	ServerTlsPolicy string `pulumi:"serverTlsPolicy"`
	// Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
	TrafficPortSelector TrafficPortSelectorResponse `pulumi:"trafficPortSelector"`
	// The type of endpoint policy. This is primarily used to validate the configuration.
	Type string `pulumi:"type"`
	// The timestamp when the resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupEndpointPolicyOutput(ctx *pulumi.Context, args LookupEndpointPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointPolicyResult, error) {
			args := v.(LookupEndpointPolicyArgs)
			r, err := LookupEndpointPolicy(ctx, &args, opts...)
			var s LookupEndpointPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointPolicyResultOutput)
}

type LookupEndpointPolicyOutputArgs struct {
	EndpointPolicyId pulumi.StringInput    `pulumi:"endpointPolicyId"`
	Location         pulumi.StringInput    `pulumi:"location"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupEndpointPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointPolicyArgs)(nil)).Elem()
}

type LookupEndpointPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointPolicyResult)(nil)).Elem()
}

func (o LookupEndpointPolicyResultOutput) ToLookupEndpointPolicyResultOutput() LookupEndpointPolicyResultOutput {
	return o
}

func (o LookupEndpointPolicyResultOutput) ToLookupEndpointPolicyResultOutputWithContext(ctx context.Context) LookupEndpointPolicyResultOutput {
	return o
}

// Optional. This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints. Refer to Authorization. If this field is not specified, authorization is disabled(no authz checks) for this endpoint.
func (o LookupEndpointPolicyResultOutput) AuthorizationPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) string { return v.AuthorizationPolicy }).(pulumi.StringOutput)
}

// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints. More specifically, it is applied to the outgoing traffic from the proxy to the endpoint. This is typically used for sidecar model where the proxy identifies itself as endpoint to the control plane, with the connection between sidecar and endpoint requiring authentication. If this field is not set, authentication is disabled(open). Applicable only when EndpointPolicyType is SIDECAR_PROXY.
func (o LookupEndpointPolicyResultOutput) ClientTlsPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) string { return v.ClientTlsPolicy }).(pulumi.StringOutput)
}

// The timestamp when the resource was created.
func (o LookupEndpointPolicyResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A free-text description of the resource. Max length 1024 characters.
func (o LookupEndpointPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// A matcher that selects endpoints to which the policies should be applied.
func (o LookupEndpointPolicyResultOutput) EndpointMatcher() EndpointMatcherResponseOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) EndpointMatcherResponse { return v.EndpointMatcher }).(EndpointMatcherResponseOutput)
}

// Optional. Set of label tags associated with the EndpointPolicy resource.
func (o LookupEndpointPolicyResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the EndpointPolicy resource. It matches pattern `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
func (o LookupEndpointPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends. If this field is not set, authentication is disabled(open) for this endpoint.
func (o LookupEndpointPolicyResultOutput) ServerTlsPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) string { return v.ServerTlsPolicy }).(pulumi.StringOutput)
}

// Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
func (o LookupEndpointPolicyResultOutput) TrafficPortSelector() TrafficPortSelectorResponseOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) TrafficPortSelectorResponse { return v.TrafficPortSelector }).(TrafficPortSelectorResponseOutput)
}

// The type of endpoint policy. This is primarily used to validate the configuration.
func (o LookupEndpointPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o LookupEndpointPolicyResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointPolicyResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointPolicyResultOutput{})
}
