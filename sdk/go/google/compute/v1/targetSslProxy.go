// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a TargetSslProxy resource in the specified project using the data included in the request.
type TargetSslProxy struct {
	pulumi.CustomResourceState

	// URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored. Accepted format is //certificatemanager.googleapis.com/projects/{project }/locations/{location}/certificateMaps/{resourceName}.
	CertificateMap pulumi.StringOutput `pulumi:"certificateMap"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#targetSslProxy for target SSL proxies.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader pulumi.StringOutput `pulumi:"proxyHeader"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// URL to the BackendService resource.
	Service pulumi.StringOutput `pulumi:"service"`
	// URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	SslCertificates pulumi.StringArrayOutput `pulumi:"sslCertificates"`
	// URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
	SslPolicy pulumi.StringOutput `pulumi:"sslPolicy"`
}

// NewTargetSslProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetSslProxy(ctx *pulumi.Context,
	name string, args *TargetSslProxyArgs, opts ...pulumi.ResourceOption) (*TargetSslProxy, error) {
	if args == nil {
		args = &TargetSslProxyArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TargetSslProxy
	err := ctx.RegisterResource("google-native:compute/v1:TargetSslProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetSslProxy gets an existing TargetSslProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetSslProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetSslProxyState, opts ...pulumi.ResourceOption) (*TargetSslProxy, error) {
	var resource TargetSslProxy
	err := ctx.ReadResource("google-native:compute/v1:TargetSslProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetSslProxy resources.
type targetSslProxyState struct {
}

type TargetSslProxyState struct {
}

func (TargetSslProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetSslProxyState)(nil)).Elem()
}

type targetSslProxyArgs struct {
	// URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored. Accepted format is //certificatemanager.googleapis.com/projects/{project }/locations/{location}/certificateMaps/{resourceName}.
	CertificateMap *string `pulumi:"certificateMap"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader *TargetSslProxyProxyHeader `pulumi:"proxyHeader"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// URL to the BackendService resource.
	Service *string `pulumi:"service"`
	// URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	SslCertificates []string `pulumi:"sslCertificates"`
	// URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
}

// The set of arguments for constructing a TargetSslProxy resource.
type TargetSslProxyArgs struct {
	// URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored. Accepted format is //certificatemanager.googleapis.com/projects/{project }/locations/{location}/certificateMaps/{resourceName}.
	CertificateMap pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	ProxyHeader TargetSslProxyProxyHeaderPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// URL to the BackendService resource.
	Service pulumi.StringPtrInput
	// URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	SslCertificates pulumi.StringArrayInput
	// URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
}

func (TargetSslProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetSslProxyArgs)(nil)).Elem()
}

type TargetSslProxyInput interface {
	pulumi.Input

	ToTargetSslProxyOutput() TargetSslProxyOutput
	ToTargetSslProxyOutputWithContext(ctx context.Context) TargetSslProxyOutput
}

func (*TargetSslProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetSslProxy)(nil)).Elem()
}

func (i *TargetSslProxy) ToTargetSslProxyOutput() TargetSslProxyOutput {
	return i.ToTargetSslProxyOutputWithContext(context.Background())
}

func (i *TargetSslProxy) ToTargetSslProxyOutputWithContext(ctx context.Context) TargetSslProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetSslProxyOutput)
}

type TargetSslProxyOutput struct{ *pulumi.OutputState }

func (TargetSslProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetSslProxy)(nil)).Elem()
}

func (o TargetSslProxyOutput) ToTargetSslProxyOutput() TargetSslProxyOutput {
	return o
}

func (o TargetSslProxyOutput) ToTargetSslProxyOutputWithContext(ctx context.Context) TargetSslProxyOutput {
	return o
}

// URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored. Accepted format is //certificatemanager.googleapis.com/projects/{project }/locations/{location}/certificateMaps/{resourceName}.
func (o TargetSslProxyOutput) CertificateMap() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.CertificateMap }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o TargetSslProxyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o TargetSslProxyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#targetSslProxy for target SSL proxies.
func (o TargetSslProxyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o TargetSslProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TargetSslProxyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
func (o TargetSslProxyOutput) ProxyHeader() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.ProxyHeader }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o TargetSslProxyOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o TargetSslProxyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// URL to the BackendService resource.
func (o TargetSslProxyOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
func (o TargetSslProxyOutput) SslCertificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringArrayOutput { return v.SslCertificates }).(pulumi.StringArrayOutput)
}

// URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
func (o TargetSslProxyOutput) SslPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetSslProxy) pulumi.StringOutput { return v.SslPolicy }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetSslProxyInput)(nil)).Elem(), &TargetSslProxy{})
	pulumi.RegisterOutputType(TargetSslProxyOutput{})
}
