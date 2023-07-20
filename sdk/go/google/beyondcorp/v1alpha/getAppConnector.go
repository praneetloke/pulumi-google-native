// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single AppConnector.
func LookupAppConnector(ctx *pulumi.Context, args *LookupAppConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAppConnectorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAppConnectorResult
	err := ctx.Invoke("google-native:beyondcorp/v1alpha:getAppConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppConnectorArgs struct {
	AppConnectorId string  `pulumi:"appConnectorId"`
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
}

type LookupAppConnectorResult struct {
	// Timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. An arbitrary user-provided name for the AppConnector. Cannot exceed 64 characters.
	DisplayName string `pulumi:"displayName"`
	// Optional. Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Unique resource name of the AppConnector. The name is ignored when creating a AppConnector.
	Name string `pulumi:"name"`
	// Principal information about the Identity of the AppConnector.
	PrincipalInfo GoogleCloudBeyondcorpAppconnectorsV1alphaAppConnectorPrincipalInfoResponse `pulumi:"principalInfo"`
	// Optional. Resource info of the connector.
	ResourceInfo GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponse `pulumi:"resourceInfo"`
	// The current state of the AppConnector.
	State string `pulumi:"state"`
	// A unique identifier for the instance generated by the system.
	Uid string `pulumi:"uid"`
	// Timestamp when the resource was last modified.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupAppConnectorOutput(ctx *pulumi.Context, args LookupAppConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAppConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppConnectorResult, error) {
			args := v.(LookupAppConnectorArgs)
			r, err := LookupAppConnector(ctx, &args, opts...)
			var s LookupAppConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppConnectorResultOutput)
}

type LookupAppConnectorOutputArgs struct {
	AppConnectorId pulumi.StringInput    `pulumi:"appConnectorId"`
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupAppConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppConnectorArgs)(nil)).Elem()
}

type LookupAppConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAppConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppConnectorResult)(nil)).Elem()
}

func (o LookupAppConnectorResultOutput) ToLookupAppConnectorResultOutput() LookupAppConnectorResultOutput {
	return o
}

func (o LookupAppConnectorResultOutput) ToLookupAppConnectorResultOutputWithContext(ctx context.Context) LookupAppConnectorResultOutput {
	return o
}

// Timestamp when the resource was created.
func (o LookupAppConnectorResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. An arbitrary user-provided name for the AppConnector. Cannot exceed 64 characters.
func (o LookupAppConnectorResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Resource labels to represent user provided metadata.
func (o LookupAppConnectorResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Unique resource name of the AppConnector. The name is ignored when creating a AppConnector.
func (o LookupAppConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Principal information about the Identity of the AppConnector.
func (o LookupAppConnectorResultOutput) PrincipalInfo() GoogleCloudBeyondcorpAppconnectorsV1alphaAppConnectorPrincipalInfoResponseOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) GoogleCloudBeyondcorpAppconnectorsV1alphaAppConnectorPrincipalInfoResponse {
		return v.PrincipalInfo
	}).(GoogleCloudBeyondcorpAppconnectorsV1alphaAppConnectorPrincipalInfoResponseOutput)
}

// Optional. Resource info of the connector.
func (o LookupAppConnectorResultOutput) ResourceInfo() GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponseOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponse {
		return v.ResourceInfo
	}).(GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponseOutput)
}

// The current state of the AppConnector.
func (o LookupAppConnectorResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) string { return v.State }).(pulumi.StringOutput)
}

// A unique identifier for the instance generated by the system.
func (o LookupAppConnectorResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) string { return v.Uid }).(pulumi.StringOutput)
}

// Timestamp when the resource was last modified.
func (o LookupAppConnectorResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppConnectorResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppConnectorResultOutput{})
}
