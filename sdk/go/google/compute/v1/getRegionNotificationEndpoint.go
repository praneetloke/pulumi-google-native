// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified NotificationEndpoint resource in the given region.
func LookupRegionNotificationEndpoint(ctx *pulumi.Context, args *LookupRegionNotificationEndpointArgs, opts ...pulumi.InvokeOption) (*LookupRegionNotificationEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegionNotificationEndpointResult
	err := ctx.Invoke("google-native:compute/v1:getRegionNotificationEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionNotificationEndpointArgs struct {
	NotificationEndpoint string  `pulumi:"notificationEndpoint"`
	Project              *string `pulumi:"project"`
	Region               string  `pulumi:"region"`
}

type LookupRegionNotificationEndpointResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
	GrpcSettings NotificationEndpointGrpcSettingsResponse `pulumi:"grpcSettings"`
	// Type of the resource. Always compute#notificationEndpoint for notification endpoints.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// URL of the region where the notification endpoint resides. This field applies only to the regional resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
}

func LookupRegionNotificationEndpointOutput(ctx *pulumi.Context, args LookupRegionNotificationEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupRegionNotificationEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionNotificationEndpointResult, error) {
			args := v.(LookupRegionNotificationEndpointArgs)
			r, err := LookupRegionNotificationEndpoint(ctx, &args, opts...)
			var s LookupRegionNotificationEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegionNotificationEndpointResultOutput)
}

type LookupRegionNotificationEndpointOutputArgs struct {
	NotificationEndpoint pulumi.StringInput    `pulumi:"notificationEndpoint"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
	Region               pulumi.StringInput    `pulumi:"region"`
}

func (LookupRegionNotificationEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionNotificationEndpointArgs)(nil)).Elem()
}

type LookupRegionNotificationEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupRegionNotificationEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionNotificationEndpointResult)(nil)).Elem()
}

func (o LookupRegionNotificationEndpointResultOutput) ToLookupRegionNotificationEndpointResultOutput() LookupRegionNotificationEndpointResultOutput {
	return o
}

func (o LookupRegionNotificationEndpointResultOutput) ToLookupRegionNotificationEndpointResultOutputWithContext(ctx context.Context) LookupRegionNotificationEndpointResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupRegionNotificationEndpointResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNotificationEndpointResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRegionNotificationEndpointResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNotificationEndpointResult) string { return v.Description }).(pulumi.StringOutput)
}

// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
func (o LookupRegionNotificationEndpointResultOutput) GrpcSettings() NotificationEndpointGrpcSettingsResponseOutput {
	return o.ApplyT(func(v LookupRegionNotificationEndpointResult) NotificationEndpointGrpcSettingsResponse {
		return v.GrpcSettings
	}).(NotificationEndpointGrpcSettingsResponseOutput)
}

// Type of the resource. Always compute#notificationEndpoint for notification endpoints.
func (o LookupRegionNotificationEndpointResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNotificationEndpointResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionNotificationEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNotificationEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// URL of the region where the notification endpoint resides. This field applies only to the regional resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupRegionNotificationEndpointResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNotificationEndpointResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupRegionNotificationEndpointResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNotificationEndpointResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionNotificationEndpointResultOutput{})
}
