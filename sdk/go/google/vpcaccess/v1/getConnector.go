// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Serverless VPC Access connector. Returns NOT_FOUND if the resource does not exist.
func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectorResult
	err := ctx.Invoke("google-native:vpcaccess/v1:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	ConnectorId string  `pulumi:"connectorId"`
	Location    string  `pulumi:"location"`
	Project     *string `pulumi:"project"`
}

type LookupConnectorResult struct {
	// List of projects using the connector.
	ConnectedProjects []string `pulumi:"connectedProjects"`
	// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
	IpCidrRange string `pulumi:"ipCidrRange"`
	// Machine type of VM Instance underlying connector. Default is e2-micro
	MachineType string `pulumi:"machineType"`
	// Maximum value of instances in autoscaling group underlying the connector.
	MaxInstances int `pulumi:"maxInstances"`
	// Maximum throughput of the connector in Mbps. Default is 300, max is 1000.
	MaxThroughput int `pulumi:"maxThroughput"`
	// Minimum value of instances in autoscaling group underlying the connector.
	MinInstances int `pulumi:"minInstances"`
	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput int `pulumi:"minThroughput"`
	// The resource name in the format `projects/*/locations/*/connectors/*`.
	Name string `pulumi:"name"`
	// Name of a VPC network.
	Network string `pulumi:"network"`
	// State of the VPC access connector.
	State string `pulumi:"state"`
	// The subnet in which to house the VPC Access Connector.
	Subnet SubnetResponse `pulumi:"subnet"`
}

func LookupConnectorOutput(ctx *pulumi.Context, args LookupConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorResult, error) {
			args := v.(LookupConnectorArgs)
			r, err := LookupConnector(ctx, &args, opts...)
			var s LookupConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorResultOutput)
}

type LookupConnectorOutputArgs struct {
	ConnectorId pulumi.StringInput    `pulumi:"connectorId"`
	Location    pulumi.StringInput    `pulumi:"location"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorArgs)(nil)).Elem()
}

type LookupConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorResult)(nil)).Elem()
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutput() LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutputWithContext(ctx context.Context) LookupConnectorResultOutput {
	return o
}

// List of projects using the connector.
func (o LookupConnectorResultOutput) ConnectedProjects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectorResult) []string { return v.ConnectedProjects }).(pulumi.StringArrayOutput)
}

// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
func (o LookupConnectorResultOutput) IpCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.IpCidrRange }).(pulumi.StringOutput)
}

// Machine type of VM Instance underlying connector. Default is e2-micro
func (o LookupConnectorResultOutput) MachineType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.MachineType }).(pulumi.StringOutput)
}

// Maximum value of instances in autoscaling group underlying the connector.
func (o LookupConnectorResultOutput) MaxInstances() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectorResult) int { return v.MaxInstances }).(pulumi.IntOutput)
}

// Maximum throughput of the connector in Mbps. Default is 300, max is 1000.
func (o LookupConnectorResultOutput) MaxThroughput() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectorResult) int { return v.MaxThroughput }).(pulumi.IntOutput)
}

// Minimum value of instances in autoscaling group underlying the connector.
func (o LookupConnectorResultOutput) MinInstances() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectorResult) int { return v.MinInstances }).(pulumi.IntOutput)
}

// Minimum throughput of the connector in Mbps. Default and min is 200.
func (o LookupConnectorResultOutput) MinThroughput() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectorResult) int { return v.MinThroughput }).(pulumi.IntOutput)
}

// The resource name in the format `projects/*/locations/*/connectors/*`.
func (o LookupConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Name of a VPC network.
func (o LookupConnectorResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Network }).(pulumi.StringOutput)
}

// State of the VPC access connector.
func (o LookupConnectorResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.State }).(pulumi.StringOutput)
}

// The subnet in which to house the VPC Access Connector.
func (o LookupConnectorResultOutput) Subnet() SubnetResponseOutput {
	return o.ApplyT(func(v LookupConnectorResult) SubnetResponse { return v.Subnet }).(SubnetResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}
