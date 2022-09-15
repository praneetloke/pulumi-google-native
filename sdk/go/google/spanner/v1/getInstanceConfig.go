// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a particular instance configuration.
func LookupInstanceConfig(ctx *pulumi.Context, args *LookupInstanceConfigArgs, opts ...pulumi.InvokeOption) (*LookupInstanceConfigResult, error) {
	var rv LookupInstanceConfigResult
	err := ctx.Invoke("google-native:spanner/v1:getInstanceConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceConfigArgs struct {
	InstanceConfigId string  `pulumi:"instanceConfigId"`
	Project          *string `pulumi:"project"`
}

type LookupInstanceConfigResult struct {
	// Base configuration name, e.g. projects//instanceConfigs/nam3, based on which this configuration is created. Only set for user managed configurations. `base_config` must refer to a configuration of type GOOGLE_MANAGED in the same project as this configuration.
	BaseConfig string `pulumi:"baseConfig"`
	// Whether this instance config is a Google or User Managed Configuration.
	ConfigType string `pulumi:"configType"`
	// The name of this instance configuration as it appears in UIs.
	DisplayName string `pulumi:"displayName"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a instance config from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform instance config updates in order to avoid race conditions: An etag is returned in the response which contains instance configs, and systems are expected to put that etag in the request to update instance config to ensure that their change will be applied to the same version of the instance config. If no etag is provided in the call to update instance config, then the existing instance config is overwritten blindly.
	Etag string `pulumi:"etag"`
	// Describes whether free instances are available to be created in this instance config.
	FreeInstanceAvailability string `pulumi:"freeInstanceAvailability"`
	// Cloud Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. Cloud Labels can be used to filter collections of resources. They can be used to control how resource metrics are aggregated. And they can be used as arguments to policy management rules (e.g. route, firewall, load balancing, etc.). * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `[a-z0-9_-]{0,63}`. * No more than 64 labels can be associated with a given resource. See https://goo.gl/xmQnxf for more information on and examples of labels. If you plan to use labels in your own code, please note that additional characters may be allowed in the future. Therefore, you are advised to use an internal label representation, such as JSON, which doesn't rely upon specific characters being disallowed. For example, representing labels as the string: name + "_" + value would prove problematic if we were to allow "_" in a future release.
	Labels map[string]string `pulumi:"labels"`
	// Allowed values of the "default_leader" schema option for databases in instances that use this instance configuration.
	LeaderOptions []string `pulumi:"leaderOptions"`
	// A unique identifier for the instance configuration. Values are of the form `projects//instanceConfigs/a-z*`.
	Name string `pulumi:"name"`
	// The available optional replicas to choose from for user managed configurations. Populated for Google managed configurations.
	OptionalReplicas []ReplicaInfoResponse `pulumi:"optionalReplicas"`
	// If true, the instance config is being created or updated. If false, there are no ongoing operations for the instance config.
	Reconciling bool `pulumi:"reconciling"`
	// The geographic placement of nodes in this instance configuration and their replication properties.
	Replicas []ReplicaInfoResponse `pulumi:"replicas"`
	// The current instance config state.
	State string `pulumi:"state"`
}

func LookupInstanceConfigOutput(ctx *pulumi.Context, args LookupInstanceConfigOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceConfigResult, error) {
			args := v.(LookupInstanceConfigArgs)
			r, err := LookupInstanceConfig(ctx, &args, opts...)
			var s LookupInstanceConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceConfigResultOutput)
}

type LookupInstanceConfigOutputArgs struct {
	InstanceConfigId pulumi.StringInput    `pulumi:"instanceConfigId"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInstanceConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceConfigArgs)(nil)).Elem()
}

type LookupInstanceConfigResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceConfigResult)(nil)).Elem()
}

func (o LookupInstanceConfigResultOutput) ToLookupInstanceConfigResultOutput() LookupInstanceConfigResultOutput {
	return o
}

func (o LookupInstanceConfigResultOutput) ToLookupInstanceConfigResultOutputWithContext(ctx context.Context) LookupInstanceConfigResultOutput {
	return o
}

// Base configuration name, e.g. projects//instanceConfigs/nam3, based on which this configuration is created. Only set for user managed configurations. `base_config` must refer to a configuration of type GOOGLE_MANAGED in the same project as this configuration.
func (o LookupInstanceConfigResultOutput) BaseConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) string { return v.BaseConfig }).(pulumi.StringOutput)
}

// Whether this instance config is a Google or User Managed Configuration.
func (o LookupInstanceConfigResultOutput) ConfigType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) string { return v.ConfigType }).(pulumi.StringOutput)
}

// The name of this instance configuration as it appears in UIs.
func (o LookupInstanceConfigResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a instance config from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform instance config updates in order to avoid race conditions: An etag is returned in the response which contains instance configs, and systems are expected to put that etag in the request to update instance config to ensure that their change will be applied to the same version of the instance config. If no etag is provided in the call to update instance config, then the existing instance config is overwritten blindly.
func (o LookupInstanceConfigResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Describes whether free instances are available to be created in this instance config.
func (o LookupInstanceConfigResultOutput) FreeInstanceAvailability() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) string { return v.FreeInstanceAvailability }).(pulumi.StringOutput)
}

// Cloud Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. Cloud Labels can be used to filter collections of resources. They can be used to control how resource metrics are aggregated. And they can be used as arguments to policy management rules (e.g. route, firewall, load balancing, etc.). * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `[a-z0-9_-]{0,63}`. * No more than 64 labels can be associated with a given resource. See https://goo.gl/xmQnxf for more information on and examples of labels. If you plan to use labels in your own code, please note that additional characters may be allowed in the future. Therefore, you are advised to use an internal label representation, such as JSON, which doesn't rely upon specific characters being disallowed. For example, representing labels as the string: name + "_" + value would prove problematic if we were to allow "_" in a future release.
func (o LookupInstanceConfigResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Allowed values of the "default_leader" schema option for databases in instances that use this instance configuration.
func (o LookupInstanceConfigResultOutput) LeaderOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) []string { return v.LeaderOptions }).(pulumi.StringArrayOutput)
}

// A unique identifier for the instance configuration. Values are of the form `projects//instanceConfigs/a-z*`.
func (o LookupInstanceConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The available optional replicas to choose from for user managed configurations. Populated for Google managed configurations.
func (o LookupInstanceConfigResultOutput) OptionalReplicas() ReplicaInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) []ReplicaInfoResponse { return v.OptionalReplicas }).(ReplicaInfoResponseArrayOutput)
}

// If true, the instance config is being created or updated. If false, there are no ongoing operations for the instance config.
func (o LookupInstanceConfigResultOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) bool { return v.Reconciling }).(pulumi.BoolOutput)
}

// The geographic placement of nodes in this instance configuration and their replication properties.
func (o LookupInstanceConfigResultOutput) Replicas() ReplicaInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) []ReplicaInfoResponse { return v.Replicas }).(ReplicaInfoResponseArrayOutput)
}

// The current instance config state.
func (o LookupInstanceConfigResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceConfigResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceConfigResultOutput{})
}