// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single ClientConnectorService.
func LookupClientConnectorService(ctx *pulumi.Context, args *LookupClientConnectorServiceArgs, opts ...pulumi.InvokeOption) (*LookupClientConnectorServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClientConnectorServiceResult
	err := ctx.Invoke("google-native:beyondcorp/v1alpha:getClientConnectorService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClientConnectorServiceArgs struct {
	ClientConnectorServiceId string  `pulumi:"clientConnectorServiceId"`
	Location                 string  `pulumi:"location"`
	Project                  *string `pulumi:"project"`
}

type LookupClientConnectorServiceResult struct {
	// [Output only] Create time stamp.
	CreateTime string `pulumi:"createTime"`
	// Optional. User-provided name. The display name should follow certain format. * Must be 6 to 30 characters in length. * Can only contain lowercase letters, numbers, and hyphens. * Must start with a letter.
	DisplayName string `pulumi:"displayName"`
	// The details of the egress settings.
	Egress EgressResponse `pulumi:"egress"`
	// The details of the ingress settings.
	Ingress IngressResponse `pulumi:"ingress"`
	// Name of resource. The name is ignored during creation.
	Name string `pulumi:"name"`
	// The operational state of the ClientConnectorService.
	State string `pulumi:"state"`
	// [Output only] Update time stamp.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupClientConnectorServiceOutput(ctx *pulumi.Context, args LookupClientConnectorServiceOutputArgs, opts ...pulumi.InvokeOption) LookupClientConnectorServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClientConnectorServiceResult, error) {
			args := v.(LookupClientConnectorServiceArgs)
			r, err := LookupClientConnectorService(ctx, &args, opts...)
			var s LookupClientConnectorServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClientConnectorServiceResultOutput)
}

type LookupClientConnectorServiceOutputArgs struct {
	ClientConnectorServiceId pulumi.StringInput    `pulumi:"clientConnectorServiceId"`
	Location                 pulumi.StringInput    `pulumi:"location"`
	Project                  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupClientConnectorServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClientConnectorServiceArgs)(nil)).Elem()
}

type LookupClientConnectorServiceResultOutput struct{ *pulumi.OutputState }

func (LookupClientConnectorServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClientConnectorServiceResult)(nil)).Elem()
}

func (o LookupClientConnectorServiceResultOutput) ToLookupClientConnectorServiceResultOutput() LookupClientConnectorServiceResultOutput {
	return o
}

func (o LookupClientConnectorServiceResultOutput) ToLookupClientConnectorServiceResultOutputWithContext(ctx context.Context) LookupClientConnectorServiceResultOutput {
	return o
}

// [Output only] Create time stamp.
func (o LookupClientConnectorServiceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientConnectorServiceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. User-provided name. The display name should follow certain format. * Must be 6 to 30 characters in length. * Can only contain lowercase letters, numbers, and hyphens. * Must start with a letter.
func (o LookupClientConnectorServiceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientConnectorServiceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The details of the egress settings.
func (o LookupClientConnectorServiceResultOutput) Egress() EgressResponseOutput {
	return o.ApplyT(func(v LookupClientConnectorServiceResult) EgressResponse { return v.Egress }).(EgressResponseOutput)
}

// The details of the ingress settings.
func (o LookupClientConnectorServiceResultOutput) Ingress() IngressResponseOutput {
	return o.ApplyT(func(v LookupClientConnectorServiceResult) IngressResponse { return v.Ingress }).(IngressResponseOutput)
}

// Name of resource. The name is ignored during creation.
func (o LookupClientConnectorServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientConnectorServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The operational state of the ClientConnectorService.
func (o LookupClientConnectorServiceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientConnectorServiceResult) string { return v.State }).(pulumi.StringOutput)
}

// [Output only] Update time stamp.
func (o LookupClientConnectorServiceResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientConnectorServiceResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClientConnectorServiceResultOutput{})
}
