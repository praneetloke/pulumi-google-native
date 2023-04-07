// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a particular instance.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:spanner/v1:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	FieldMask  *string `pulumi:"fieldMask"`
	InstanceId string  `pulumi:"instanceId"`
	Project    *string `pulumi:"project"`
}

type LookupInstanceResult struct {
	// The name of the instance's configuration. Values are of the form `projects//instanceConfigs/`. See also InstanceConfig and ListInstanceConfigs.
	Config string `pulumi:"config"`
	// The time at which the instance was created.
	CreateTime string `pulumi:"createTime"`
	// The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30 characters in length.
	DisplayName string `pulumi:"displayName"`
	// Deprecated. This field is not populated.
	//
	// Deprecated: Deprecated. This field is not populated.
	EndpointUris []string `pulumi:"endpointUris"`
	// Free instance metadata. Only populated for free instances.
	FreeInstanceMetadata FreeInstanceMetadataResponse `pulumi:"freeInstanceMetadata"`
	// The `InstanceType` of the current instance.
	InstanceType string `pulumi:"instanceType"`
	// Cloud Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. Cloud Labels can be used to filter collections of resources. They can be used to control how resource metrics are aggregated. And they can be used as arguments to policy management rules (e.g. route, firewall, load balancing, etc.). * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `[a-z0-9_-]{0,63}`. * No more than 64 labels can be associated with a given resource. See https://goo.gl/xmQnxf for more information on and examples of labels. If you plan to use labels in your own code, please note that additional characters may be allowed in the future. And so you are advised to use an internal label representation, such as JSON, which doesn't rely upon specific characters being disallowed. For example, representing labels as the string: name + "_" + value would prove problematic if we were to allow "_" in a future release.
	Labels map[string]string `pulumi:"labels"`
	// A unique identifier for the instance, which cannot be changed after the instance is created. Values are of the form `projects//instances/a-z*[a-z0-9]`. The final segment of the name must be between 2 and 64 characters in length.
	Name string `pulumi:"name"`
	// The number of nodes allocated to this instance. At most one of either node_count or processing_units should be present in the message. Users can set the node_count field to specify the target number of nodes allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
	NodeCount int `pulumi:"nodeCount"`
	// The number of processing units allocated to this instance. At most one of processing_units or node_count should be present in the message. Users can set the processing_units field to specify the target number of processing units allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
	ProcessingUnits int `pulumi:"processingUnits"`
	// The current instance state. For CreateInstance, the state must be either omitted or set to `CREATING`. For UpdateInstance, the state must be either omitted or set to `READY`.
	State string `pulumi:"state"`
	// The time at which the instance was most recently updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	FieldMask  pulumi.StringPtrInput `pulumi:"fieldMask"`
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// The name of the instance's configuration. Values are of the form `projects//instanceConfigs/`. See also InstanceConfig and ListInstanceConfigs.
func (o LookupInstanceResultOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Config }).(pulumi.StringOutput)
}

// The time at which the instance was created.
func (o LookupInstanceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30 characters in length.
func (o LookupInstanceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Deprecated. This field is not populated.
//
// Deprecated: Deprecated. This field is not populated.
func (o LookupInstanceResultOutput) EndpointUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.EndpointUris }).(pulumi.StringArrayOutput)
}

// Free instance metadata. Only populated for free instances.
func (o LookupInstanceResultOutput) FreeInstanceMetadata() FreeInstanceMetadataResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) FreeInstanceMetadataResponse { return v.FreeInstanceMetadata }).(FreeInstanceMetadataResponseOutput)
}

// The `InstanceType` of the current instance.
func (o LookupInstanceResultOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.InstanceType }).(pulumi.StringOutput)
}

// Cloud Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. Cloud Labels can be used to filter collections of resources. They can be used to control how resource metrics are aggregated. And they can be used as arguments to policy management rules (e.g. route, firewall, load balancing, etc.). * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `[a-z0-9_-]{0,63}`. * No more than 64 labels can be associated with a given resource. See https://goo.gl/xmQnxf for more information on and examples of labels. If you plan to use labels in your own code, please note that additional characters may be allowed in the future. And so you are advised to use an internal label representation, such as JSON, which doesn't rely upon specific characters being disallowed. For example, representing labels as the string: name + "_" + value would prove problematic if we were to allow "_" in a future release.
func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// A unique identifier for the instance, which cannot be changed after the instance is created. Values are of the form `projects//instances/a-z*[a-z0-9]`. The final segment of the name must be between 2 and 64 characters in length.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of nodes allocated to this instance. At most one of either node_count or processing_units should be present in the message. Users can set the node_count field to specify the target number of nodes allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
func (o LookupInstanceResultOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.NodeCount }).(pulumi.IntOutput)
}

// The number of processing units allocated to this instance. At most one of processing_units or node_count should be present in the message. Users can set the processing_units field to specify the target number of processing units allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
func (o LookupInstanceResultOutput) ProcessingUnits() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.ProcessingUnits }).(pulumi.IntOutput)
}

// The current instance state. For CreateInstance, the state must be either omitted or set to `CREATING`. For UpdateInstance, the state must be either omitted or set to `READY`.
func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

// The time at which the instance was most recently updated.
func (o LookupInstanceResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
