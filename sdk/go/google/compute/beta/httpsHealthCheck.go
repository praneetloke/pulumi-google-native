// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a HttpsHealthCheck resource in the specified project using the data included in the request.
type HttpsHealthCheck struct {
	pulumi.CustomResourceState

	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec pulumi.IntOutput `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntOutput `pulumi:"healthyThreshold"`
	// The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
	Host pulumi.StringOutput `pulumi:"host"`
	// Type of the resource.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The TCP port number for the HTTPS health check request. The default value is 443.
	Port    pulumi.IntOutput    `pulumi:"port"`
	Project pulumi.StringOutput `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The request path of the HTTPS health check request. The default value is "/". Must comply with RFC3986.
	RequestPath pulumi.StringOutput `pulumi:"requestPath"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have a greater value than checkIntervalSec.
	TimeoutSec pulumi.IntOutput `pulumi:"timeoutSec"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntOutput `pulumi:"unhealthyThreshold"`
}

// NewHttpsHealthCheck registers a new resource with the given unique name, arguments, and options.
func NewHttpsHealthCheck(ctx *pulumi.Context,
	name string, args *HttpsHealthCheckArgs, opts ...pulumi.ResourceOption) (*HttpsHealthCheck, error) {
	if args == nil {
		args = &HttpsHealthCheckArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource HttpsHealthCheck
	err := ctx.RegisterResource("google-native:compute/beta:HttpsHealthCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpsHealthCheck gets an existing HttpsHealthCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpsHealthCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpsHealthCheckState, opts ...pulumi.ResourceOption) (*HttpsHealthCheck, error) {
	var resource HttpsHealthCheck
	err := ctx.ReadResource("google-native:compute/beta:HttpsHealthCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpsHealthCheck resources.
type httpsHealthCheckState struct {
}

type HttpsHealthCheckState struct {
}

func (HttpsHealthCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpsHealthCheckState)(nil)).Elem()
}

type httpsHealthCheckArgs struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec *int `pulumi:"checkIntervalSec"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
	Host *string `pulumi:"host"`
	// Type of the resource.
	Kind *string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The TCP port number for the HTTPS health check request. The default value is 443.
	Port    *int    `pulumi:"port"`
	Project *string `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The request path of the HTTPS health check request. The default value is "/". Must comply with RFC3986.
	RequestPath *string `pulumi:"requestPath"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have a greater value than checkIntervalSec.
	TimeoutSec *int `pulumi:"timeoutSec"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

// The set of arguments for constructing a HttpsHealthCheck resource.
type HttpsHealthCheckArgs struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrInput
	// The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
	Host pulumi.StringPtrInput
	// Type of the resource.
	Kind pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The TCP port number for the HTTPS health check request. The default value is 443.
	Port    pulumi.IntPtrInput
	Project pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The request path of the HTTPS health check request. The default value is "/". Must comply with RFC3986.
	RequestPath pulumi.StringPtrInput
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have a greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrInput
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (HttpsHealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpsHealthCheckArgs)(nil)).Elem()
}

type HttpsHealthCheckInput interface {
	pulumi.Input

	ToHttpsHealthCheckOutput() HttpsHealthCheckOutput
	ToHttpsHealthCheckOutputWithContext(ctx context.Context) HttpsHealthCheckOutput
}

func (*HttpsHealthCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpsHealthCheck)(nil)).Elem()
}

func (i *HttpsHealthCheck) ToHttpsHealthCheckOutput() HttpsHealthCheckOutput {
	return i.ToHttpsHealthCheckOutputWithContext(context.Background())
}

func (i *HttpsHealthCheck) ToHttpsHealthCheckOutputWithContext(ctx context.Context) HttpsHealthCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpsHealthCheckOutput)
}

type HttpsHealthCheckOutput struct{ *pulumi.OutputState }

func (HttpsHealthCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpsHealthCheck)(nil)).Elem()
}

func (o HttpsHealthCheckOutput) ToHttpsHealthCheckOutput() HttpsHealthCheckOutput {
	return o
}

func (o HttpsHealthCheckOutput) ToHttpsHealthCheckOutputWithContext(ctx context.Context) HttpsHealthCheckOutput {
	return o
}

// How often (in seconds) to send a health check. The default value is 5 seconds.
func (o HttpsHealthCheckOutput) CheckIntervalSec() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.IntOutput { return v.CheckIntervalSec }).(pulumi.IntOutput)
}

// Creation timestamp in RFC3339 text format.
func (o HttpsHealthCheckOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o HttpsHealthCheckOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
func (o HttpsHealthCheckOutput) HealthyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.IntOutput { return v.HealthyThreshold }).(pulumi.IntOutput)
}

// The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
func (o HttpsHealthCheckOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Type of the resource.
func (o HttpsHealthCheckOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o HttpsHealthCheckOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The TCP port number for the HTTPS health check request. The default value is 443.
func (o HttpsHealthCheckOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o HttpsHealthCheckOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o HttpsHealthCheckOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The request path of the HTTPS health check request. The default value is "/". Must comply with RFC3986.
func (o HttpsHealthCheckOutput) RequestPath() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringOutput { return v.RequestPath }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o HttpsHealthCheckOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have a greater value than checkIntervalSec.
func (o HttpsHealthCheckOutput) TimeoutSec() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.IntOutput { return v.TimeoutSec }).(pulumi.IntOutput)
}

// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
func (o HttpsHealthCheckOutput) UnhealthyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpsHealthCheck) pulumi.IntOutput { return v.UnhealthyThreshold }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpsHealthCheckInput)(nil)).Elem(), &HttpsHealthCheck{})
	pulumi.RegisterOutputType(HttpsHealthCheckOutput{})
}
