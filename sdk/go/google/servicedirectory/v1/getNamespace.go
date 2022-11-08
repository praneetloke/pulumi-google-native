// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a namespace.
func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("google-native:servicedirectory/v1:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceArgs struct {
	Location    string  `pulumi:"location"`
	NamespaceId string  `pulumi:"namespaceId"`
	Project     *string `pulumi:"project"`
}

type LookupNamespaceResult struct {
	// Optional. Resource labels associated with this namespace. No more than 64 user labels can be associated with a given resource. Label keys and values can be no longer than 63 characters.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The resource name for the namespace in the format `projects/*/locations/*/namespaces/*`.
	Name string `pulumi:"name"`
	// The globally unique identifier of the namespace in the UUID4 format.
	Uid string `pulumi:"uid"`
}

func LookupNamespaceOutput(ctx *pulumi.Context, args LookupNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceResult, error) {
			args := v.(LookupNamespaceArgs)
			r, err := LookupNamespace(ctx, &args, opts...)
			var s LookupNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceResultOutput)
}

type LookupNamespaceOutputArgs struct {
	Location    pulumi.StringInput    `pulumi:"location"`
	NamespaceId pulumi.StringInput    `pulumi:"namespaceId"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceArgs)(nil)).Elem()
}

type LookupNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceResult)(nil)).Elem()
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutput() LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutputWithContext(ctx context.Context) LookupNamespaceResultOutput {
	return o
}

// Optional. Resource labels associated with this namespace. No more than 64 user labels can be associated with a given resource. Label keys and values can be no longer than 63 characters.
func (o LookupNamespaceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Immutable. The resource name for the namespace in the format `projects/*/locations/*/namespaces/*`.
func (o LookupNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The globally unique identifier of the namespace in the UUID4 format.
func (o LookupNamespaceResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceResultOutput{})
}
