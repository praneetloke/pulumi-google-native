// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified instance.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:remotebuildexecution/v1alpha:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Project    *string `pulumi:"project"`
}

type LookupInstanceResult struct {
	// The policy to define whether or not RBE features can be used or how they can be used.
	FeaturePolicy GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse `pulumi:"featurePolicy"`
	// The location is a GCP region. Currently only `us-central1` is supported.
	Location string `pulumi:"location"`
	// Whether stack driver logging is enabled for the instance.
	LoggingEnabled bool `pulumi:"loggingEnabled"`
	// Instance resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]`. Name should not be populated when creating an instance since it is provided in the `instance_id` field.
	Name string `pulumi:"name"`
	// State of the instance.
	State string `pulumi:"state"`
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

// The policy to define whether or not RBE features can be used or how they can be used.
func (o LookupInstanceResultOutput) FeaturePolicy() GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponse {
		return v.FeaturePolicy
	}).(GoogleDevtoolsRemotebuildexecutionAdminV1alphaFeaturePolicyResponseOutput)
}

// The location is a GCP region. Currently only `us-central1` is supported.
func (o LookupInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Whether stack driver logging is enabled for the instance.
func (o LookupInstanceResultOutput) LoggingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.LoggingEnabled }).(pulumi.BoolOutput)
}

// Instance resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]`. Name should not be populated when creating an instance since it is provided in the `instance_id` field.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of the instance.
func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
