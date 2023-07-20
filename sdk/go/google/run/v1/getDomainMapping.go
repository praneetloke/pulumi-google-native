// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a domain mapping.
func LookupDomainMapping(ctx *pulumi.Context, args *LookupDomainMappingArgs, opts ...pulumi.InvokeOption) (*LookupDomainMappingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainMappingResult
	err := ctx.Invoke("google-native:run/v1:getDomainMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainMappingArgs struct {
	DomainmappingId string  `pulumi:"domainmappingId"`
	Location        string  `pulumi:"location"`
	Project         *string `pulumi:"project"`
}

type LookupDomainMappingResult struct {
	// The API version for this call such as "domains.cloudrun.com/v1".
	ApiVersion string `pulumi:"apiVersion"`
	// The kind of resource, in this case "DomainMapping".
	Kind string `pulumi:"kind"`
	// Metadata associated with this BuildTemplate.
	Metadata ObjectMetaResponse `pulumi:"metadata"`
	// The spec for this DomainMapping.
	Spec DomainMappingSpecResponse `pulumi:"spec"`
	// The current status of the DomainMapping.
	Status DomainMappingStatusResponse `pulumi:"status"`
}

func LookupDomainMappingOutput(ctx *pulumi.Context, args LookupDomainMappingOutputArgs, opts ...pulumi.InvokeOption) LookupDomainMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainMappingResult, error) {
			args := v.(LookupDomainMappingArgs)
			r, err := LookupDomainMapping(ctx, &args, opts...)
			var s LookupDomainMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainMappingResultOutput)
}

type LookupDomainMappingOutputArgs struct {
	DomainmappingId pulumi.StringInput    `pulumi:"domainmappingId"`
	Location        pulumi.StringInput    `pulumi:"location"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDomainMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainMappingArgs)(nil)).Elem()
}

type LookupDomainMappingResultOutput struct{ *pulumi.OutputState }

func (LookupDomainMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainMappingResult)(nil)).Elem()
}

func (o LookupDomainMappingResultOutput) ToLookupDomainMappingResultOutput() LookupDomainMappingResultOutput {
	return o
}

func (o LookupDomainMappingResultOutput) ToLookupDomainMappingResultOutputWithContext(ctx context.Context) LookupDomainMappingResultOutput {
	return o
}

// The API version for this call such as "domains.cloudrun.com/v1".
func (o LookupDomainMappingResultOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainMappingResult) string { return v.ApiVersion }).(pulumi.StringOutput)
}

// The kind of resource, in this case "DomainMapping".
func (o LookupDomainMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Metadata associated with this BuildTemplate.
func (o LookupDomainMappingResultOutput) Metadata() ObjectMetaResponseOutput {
	return o.ApplyT(func(v LookupDomainMappingResult) ObjectMetaResponse { return v.Metadata }).(ObjectMetaResponseOutput)
}

// The spec for this DomainMapping.
func (o LookupDomainMappingResultOutput) Spec() DomainMappingSpecResponseOutput {
	return o.ApplyT(func(v LookupDomainMappingResult) DomainMappingSpecResponse { return v.Spec }).(DomainMappingSpecResponseOutput)
}

// The current status of the DomainMapping.
func (o LookupDomainMappingResultOutput) Status() DomainMappingStatusResponseOutput {
	return o.ApplyT(func(v LookupDomainMappingResult) DomainMappingStatusResponse { return v.Status }).(DomainMappingStatusResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainMappingResultOutput{})
}
