// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified TargetGrpcProxy resource in the given scope.
func LookupTargetGrpcProxy(ctx *pulumi.Context, args *LookupTargetGrpcProxyArgs, opts ...pulumi.InvokeOption) (*LookupTargetGrpcProxyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTargetGrpcProxyResult
	err := ctx.Invoke("google-native:compute/alpha:getTargetGrpcProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetGrpcProxyArgs struct {
	Project         *string `pulumi:"project"`
	TargetGrpcProxy string  `pulumi:"targetGrpcProxy"`
}

type LookupTargetGrpcProxyResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetGrpcProxy.
	Fingerprint string `pulumi:"fingerprint"`
	// Type of the resource. Always compute#targetGrpcProxy for target grpc proxies.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL with id for the resource.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService. The protocol field in the BackendService must be set to GRPC.
	UrlMap string `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by the urlMap may be accessed by gRPC applications without using a sidecar proxy. This will enable configuration checks on urlMap and its referenced BackendServices to not allow unsupported features. A gRPC application must use "xds:///" scheme in the target URI of the service it is connecting to. If false, indicates that the BackendServices referenced by the urlMap will be accessed by gRPC applications via a sidecar proxy. In this case, a gRPC application must not use "xds:///" scheme in the target URI of the service it is connecting to
	ValidateForProxyless bool `pulumi:"validateForProxyless"`
}

func LookupTargetGrpcProxyOutput(ctx *pulumi.Context, args LookupTargetGrpcProxyOutputArgs, opts ...pulumi.InvokeOption) LookupTargetGrpcProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetGrpcProxyResult, error) {
			args := v.(LookupTargetGrpcProxyArgs)
			r, err := LookupTargetGrpcProxy(ctx, &args, opts...)
			var s LookupTargetGrpcProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTargetGrpcProxyResultOutput)
}

type LookupTargetGrpcProxyOutputArgs struct {
	Project         pulumi.StringPtrInput `pulumi:"project"`
	TargetGrpcProxy pulumi.StringInput    `pulumi:"targetGrpcProxy"`
}

func (LookupTargetGrpcProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGrpcProxyArgs)(nil)).Elem()
}

type LookupTargetGrpcProxyResultOutput struct{ *pulumi.OutputState }

func (LookupTargetGrpcProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGrpcProxyResult)(nil)).Elem()
}

func (o LookupTargetGrpcProxyResultOutput) ToLookupTargetGrpcProxyResultOutput() LookupTargetGrpcProxyResultOutput {
	return o
}

func (o LookupTargetGrpcProxyResultOutput) ToLookupTargetGrpcProxyResultOutputWithContext(ctx context.Context) LookupTargetGrpcProxyResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupTargetGrpcProxyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupTargetGrpcProxyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetGrpcProxy.
func (o LookupTargetGrpcProxyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#targetGrpcProxy for target grpc proxies.
func (o LookupTargetGrpcProxyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupTargetGrpcProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupTargetGrpcProxyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL with id for the resource.
func (o LookupTargetGrpcProxyResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// URL to the UrlMap resource that defines the mapping from URL to the BackendService. The protocol field in the BackendService must be set to GRPC.
func (o LookupTargetGrpcProxyResultOutput) UrlMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) string { return v.UrlMap }).(pulumi.StringOutput)
}

// If true, indicates that the BackendServices referenced by the urlMap may be accessed by gRPC applications without using a sidecar proxy. This will enable configuration checks on urlMap and its referenced BackendServices to not allow unsupported features. A gRPC application must use "xds:///" scheme in the target URI of the service it is connecting to. If false, indicates that the BackendServices referenced by the urlMap will be accessed by gRPC applications via a sidecar proxy. In this case, a gRPC application must not use "xds:///" scheme in the target URI of the service it is connecting to
func (o LookupTargetGrpcProxyResultOutput) ValidateForProxyless() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTargetGrpcProxyResult) bool { return v.ValidateForProxyless }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetGrpcProxyResultOutput{})
}
