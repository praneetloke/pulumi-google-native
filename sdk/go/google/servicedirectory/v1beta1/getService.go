// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a service.
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceResult
	err := ctx.Invoke("google-native:servicedirectory/v1beta1:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	Location    string  `pulumi:"location"`
	NamespaceId string  `pulumi:"namespaceId"`
	Project     *string `pulumi:"project"`
	ServiceId   string  `pulumi:"serviceId"`
}

type LookupServiceResult struct {
	// The timestamp when the service was created.
	CreateTime string `pulumi:"createTime"`
	// Endpoints associated with this service. Returned on LookupService.ResolveService. Control plane clients should use RegistrationService.ListEndpoints.
	Endpoints []EndpointResponse `pulumi:"endpoints"`
	// Optional. Metadata for the service. This data can be consumed by service clients. Restrictions: * The entire metadata dictionary may contain up to 2000 characters, spread accoss all key-value pairs. Metadata that goes beyond this limit are rejected * Valid metadata keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/). Metadata that fails to meet these requirements are rejected Note: This field is equivalent to the `annotations` field in the v1 API. They have the same syntax and read/write to the same location in Service Directory.
	Metadata map[string]string `pulumi:"metadata"`
	// Immutable. The resource name for the service in the format `projects/*/locations/*/namespaces/*/services/*`.
	Name string `pulumi:"name"`
	// A globally unique identifier (in UUID4 format) for this service.
	Uid string `pulumi:"uid"`
	// The timestamp when the service was last updated. Note: endpoints being created/deleted/updated within the service are not considered service updates for the purpose of this timestamp.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	Location    pulumi.StringInput    `pulumi:"location"`
	NamespaceId pulumi.StringInput    `pulumi:"namespaceId"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
	ServiceId   pulumi.StringInput    `pulumi:"serviceId"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

// The timestamp when the service was created.
func (o LookupServiceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Endpoints associated with this service. Returned on LookupService.ResolveService. Control plane clients should use RegistrationService.ListEndpoints.
func (o LookupServiceResultOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []EndpointResponse { return v.Endpoints }).(EndpointResponseArrayOutput)
}

// Optional. Metadata for the service. This data can be consumed by service clients. Restrictions: * The entire metadata dictionary may contain up to 2000 characters, spread accoss all key-value pairs. Metadata that goes beyond this limit are rejected * Valid metadata keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/). Metadata that fails to meet these requirements are rejected Note: This field is equivalent to the `annotations` field in the v1 API. They have the same syntax and read/write to the same location in Service Directory.
func (o LookupServiceResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// Immutable. The resource name for the service in the format `projects/*/locations/*/namespaces/*/services/*`.
func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// A globally unique identifier (in UUID4 format) for this service.
func (o LookupServiceResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The timestamp when the service was last updated. Note: endpoints being created/deleted/updated within the service are not considered service updates for the purpose of this timestamp.
func (o LookupServiceResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
