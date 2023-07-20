// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified TargetTcpProxy resource.
func LookupRegionTargetTcpProxy(ctx *pulumi.Context, args *LookupRegionTargetTcpProxyArgs, opts ...pulumi.InvokeOption) (*LookupRegionTargetTcpProxyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegionTargetTcpProxyResult
	err := ctx.Invoke("google-native:compute/alpha:getRegionTargetTcpProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionTargetTcpProxyArgs struct {
	Project        *string `pulumi:"project"`
	Region         string  `pulumi:"region"`
	TargetTcpProxy string  `pulumi:"targetTcpProxy"`
}

type LookupRegionTargetTcpProxyResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Type of the resource. Always compute#targetTcpProxy for target TCP proxies.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	ProxyBind bool `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader string `pulumi:"proxyHeader"`
	// URL of the region where the regional TCP proxy resides. This field is not applicable to global TCP proxy.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// URL to the BackendService resource.
	Service string `pulumi:"service"`
}

func LookupRegionTargetTcpProxyOutput(ctx *pulumi.Context, args LookupRegionTargetTcpProxyOutputArgs, opts ...pulumi.InvokeOption) LookupRegionTargetTcpProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionTargetTcpProxyResult, error) {
			args := v.(LookupRegionTargetTcpProxyArgs)
			r, err := LookupRegionTargetTcpProxy(ctx, &args, opts...)
			var s LookupRegionTargetTcpProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegionTargetTcpProxyResultOutput)
}

type LookupRegionTargetTcpProxyOutputArgs struct {
	Project        pulumi.StringPtrInput `pulumi:"project"`
	Region         pulumi.StringInput    `pulumi:"region"`
	TargetTcpProxy pulumi.StringInput    `pulumi:"targetTcpProxy"`
}

func (LookupRegionTargetTcpProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionTargetTcpProxyArgs)(nil)).Elem()
}

type LookupRegionTargetTcpProxyResultOutput struct{ *pulumi.OutputState }

func (LookupRegionTargetTcpProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionTargetTcpProxyResult)(nil)).Elem()
}

func (o LookupRegionTargetTcpProxyResultOutput) ToLookupRegionTargetTcpProxyResultOutput() LookupRegionTargetTcpProxyResultOutput {
	return o
}

func (o LookupRegionTargetTcpProxyResultOutput) ToLookupRegionTargetTcpProxyResultOutputWithContext(ctx context.Context) LookupRegionTargetTcpProxyResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupRegionTargetTcpProxyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRegionTargetTcpProxyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#targetTcpProxy for target TCP proxies.
func (o LookupRegionTargetTcpProxyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionTargetTcpProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
func (o LookupRegionTargetTcpProxyResultOutput) ProxyBind() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) bool { return v.ProxyBind }).(pulumi.BoolOutput)
}

// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
func (o LookupRegionTargetTcpProxyResultOutput) ProxyHeader() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) string { return v.ProxyHeader }).(pulumi.StringOutput)
}

// URL of the region where the regional TCP proxy resides. This field is not applicable to global TCP proxy.
func (o LookupRegionTargetTcpProxyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupRegionTargetTcpProxyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// URL to the BackendService resource.
func (o LookupRegionTargetTcpProxyResultOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionTargetTcpProxyResult) string { return v.Service }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionTargetTcpProxyResultOutput{})
}
