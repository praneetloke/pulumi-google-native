// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single ClientGateway.
func LookupClientGateway(ctx *pulumi.Context, args *LookupClientGatewayArgs, opts ...pulumi.InvokeOption) (*LookupClientGatewayResult, error) {
	var rv LookupClientGatewayResult
	err := ctx.Invoke("google-native:beyondcorp/v1alpha:getClientGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClientGatewayArgs struct {
	ClientGatewayId string  `pulumi:"clientGatewayId"`
	Location        string  `pulumi:"location"`
	Project         *string `pulumi:"project"`
}

type LookupClientGatewayResult struct {
	// The client connector service name that the client gateway is associated to. Client Connector Services, named as follows: `projects/{project_id}/locations/{location_id}/client_connector_services/{client_connector_service_id}`.
	ClientConnectorService string `pulumi:"clientConnectorService"`
	// [Output only] Create time stamp.
	CreateTime string `pulumi:"createTime"`
	// name of resource. The name is ignored during creation.
	Name string `pulumi:"name"`
	// The operational state of the gateway.
	State string `pulumi:"state"`
	// [Output only] Update time stamp.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupClientGatewayOutput(ctx *pulumi.Context, args LookupClientGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupClientGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClientGatewayResult, error) {
			args := v.(LookupClientGatewayArgs)
			r, err := LookupClientGateway(ctx, &args, opts...)
			var s LookupClientGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClientGatewayResultOutput)
}

type LookupClientGatewayOutputArgs struct {
	ClientGatewayId pulumi.StringInput    `pulumi:"clientGatewayId"`
	Location        pulumi.StringInput    `pulumi:"location"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupClientGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClientGatewayArgs)(nil)).Elem()
}

type LookupClientGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupClientGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClientGatewayResult)(nil)).Elem()
}

func (o LookupClientGatewayResultOutput) ToLookupClientGatewayResultOutput() LookupClientGatewayResultOutput {
	return o
}

func (o LookupClientGatewayResultOutput) ToLookupClientGatewayResultOutputWithContext(ctx context.Context) LookupClientGatewayResultOutput {
	return o
}

// The client connector service name that the client gateway is associated to. Client Connector Services, named as follows: `projects/{project_id}/locations/{location_id}/client_connector_services/{client_connector_service_id}`.
func (o LookupClientGatewayResultOutput) ClientConnectorService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientGatewayResult) string { return v.ClientConnectorService }).(pulumi.StringOutput)
}

// [Output only] Create time stamp.
func (o LookupClientGatewayResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientGatewayResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// name of resource. The name is ignored during creation.
func (o LookupClientGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// The operational state of the gateway.
func (o LookupClientGatewayResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientGatewayResult) string { return v.State }).(pulumi.StringOutput)
}

// [Output only] Update time stamp.
func (o LookupClientGatewayResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientGatewayResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClientGatewayResultOutput{})
}