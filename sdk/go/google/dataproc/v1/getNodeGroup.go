// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the resource representation for a node group in a cluster.
func LookupNodeGroup(ctx *pulumi.Context, args *LookupNodeGroupArgs, opts ...pulumi.InvokeOption) (*LookupNodeGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNodeGroupResult
	err := ctx.Invoke("google-native:dataproc/v1:getNodeGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNodeGroupArgs struct {
	ClusterId   string  `pulumi:"clusterId"`
	NodeGroupId string  `pulumi:"nodeGroupId"`
	Project     *string `pulumi:"project"`
	RegionId    string  `pulumi:"regionId"`
}

type LookupNodeGroupResult struct {
	// Optional. Node group labels. Label keys must consist of from 1 to 63 characters and conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values can be empty. If specified, they must consist of from 1 to 63 characters and conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). The node group must have no more than 32 labelsn.
	Labels map[string]string `pulumi:"labels"`
	// The Node group resource name (https://aip.dev/122).
	Name string `pulumi:"name"`
	// Optional. The node group instance group configuration.
	NodeGroupConfig InstanceGroupConfigResponse `pulumi:"nodeGroupConfig"`
	// Node group roles.
	Roles []string `pulumi:"roles"`
}

func LookupNodeGroupOutput(ctx *pulumi.Context, args LookupNodeGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNodeGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNodeGroupResult, error) {
			args := v.(LookupNodeGroupArgs)
			r, err := LookupNodeGroup(ctx, &args, opts...)
			var s LookupNodeGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNodeGroupResultOutput)
}

type LookupNodeGroupOutputArgs struct {
	ClusterId   pulumi.StringInput    `pulumi:"clusterId"`
	NodeGroupId pulumi.StringInput    `pulumi:"nodeGroupId"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
	RegionId    pulumi.StringInput    `pulumi:"regionId"`
}

func (LookupNodeGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeGroupArgs)(nil)).Elem()
}

type LookupNodeGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNodeGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeGroupResult)(nil)).Elem()
}

func (o LookupNodeGroupResultOutput) ToLookupNodeGroupResultOutput() LookupNodeGroupResultOutput {
	return o
}

func (o LookupNodeGroupResultOutput) ToLookupNodeGroupResultOutputWithContext(ctx context.Context) LookupNodeGroupResultOutput {
	return o
}

// Optional. Node group labels. Label keys must consist of from 1 to 63 characters and conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values can be empty. If specified, they must consist of from 1 to 63 characters and conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). The node group must have no more than 32 labelsn.
func (o LookupNodeGroupResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The Node group resource name (https://aip.dev/122).
func (o LookupNodeGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. The node group instance group configuration.
func (o LookupNodeGroupResultOutput) NodeGroupConfig() InstanceGroupConfigResponseOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) InstanceGroupConfigResponse { return v.NodeGroupConfig }).(InstanceGroupConfigResponseOutput)
}

// Node group roles.
func (o LookupNodeGroupResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNodeGroupResultOutput{})
}
