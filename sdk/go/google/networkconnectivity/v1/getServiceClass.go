// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single ServiceClass.
func LookupServiceClass(ctx *pulumi.Context, args *LookupServiceClassArgs, opts ...pulumi.InvokeOption) (*LookupServiceClassResult, error) {
	var rv LookupServiceClassResult
	err := ctx.Invoke("google-native:networkconnectivity/v1:getServiceClass", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceClassArgs struct {
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
	ServiceClassId string  `pulumi:"serviceClassId"`
}

type LookupServiceClassResult struct {
	// Time when the ServiceClass was created.
	CreateTime string `pulumi:"createTime"`
	// A description of this resource.
	Description string `pulumi:"description"`
	// User-defined labels.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of a ServiceClass resource. Format: projects/{project}/locations/{location}/serviceClasses/{service_class} See: https://google.aip.dev/122#fields-representing-resource-names
	Name string `pulumi:"name"`
	// The generated service class name. Use this name to refer to the Service class in Service Connection Maps and Service Connection Policies.
	ServiceClass string `pulumi:"serviceClass"`
	// URIs of all Service Connection Maps using this service class.
	ServiceConnectionMaps []string `pulumi:"serviceConnectionMaps"`
	// Time when the ServiceClass was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupServiceClassOutput(ctx *pulumi.Context, args LookupServiceClassOutputArgs, opts ...pulumi.InvokeOption) LookupServiceClassResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceClassResult, error) {
			args := v.(LookupServiceClassArgs)
			r, err := LookupServiceClass(ctx, &args, opts...)
			var s LookupServiceClassResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceClassResultOutput)
}

type LookupServiceClassOutputArgs struct {
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
	ServiceClassId pulumi.StringInput    `pulumi:"serviceClassId"`
}

func (LookupServiceClassOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceClassArgs)(nil)).Elem()
}

type LookupServiceClassResultOutput struct{ *pulumi.OutputState }

func (LookupServiceClassResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceClassResult)(nil)).Elem()
}

func (o LookupServiceClassResultOutput) ToLookupServiceClassResultOutput() LookupServiceClassResultOutput {
	return o
}

func (o LookupServiceClassResultOutput) ToLookupServiceClassResultOutputWithContext(ctx context.Context) LookupServiceClassResultOutput {
	return o
}

// Time when the ServiceClass was created.
func (o LookupServiceClassResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceClassResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A description of this resource.
func (o LookupServiceClassResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceClassResult) string { return v.Description }).(pulumi.StringOutput)
}

// User-defined labels.
func (o LookupServiceClassResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceClassResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Immutable. The name of a ServiceClass resource. Format: projects/{project}/locations/{location}/serviceClasses/{service_class} See: https://google.aip.dev/122#fields-representing-resource-names
func (o LookupServiceClassResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceClassResult) string { return v.Name }).(pulumi.StringOutput)
}

// The generated service class name. Use this name to refer to the Service class in Service Connection Maps and Service Connection Policies.
func (o LookupServiceClassResultOutput) ServiceClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceClassResult) string { return v.ServiceClass }).(pulumi.StringOutput)
}

// URIs of all Service Connection Maps using this service class.
func (o LookupServiceClassResultOutput) ServiceConnectionMaps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServiceClassResult) []string { return v.ServiceConnectionMaps }).(pulumi.StringArrayOutput)
}

// Time when the ServiceClass was updated.
func (o LookupServiceClassResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceClassResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceClassResultOutput{})
}
