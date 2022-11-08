// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Optional. The size of the Cloud Composer environment. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
type EnvironmentConfigEnvironmentSize string

const (
	// The size of the environment is unspecified.
	EnvironmentConfigEnvironmentSizeEnvironmentSizeUnspecified = EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_UNSPECIFIED")
	// The environment size is small.
	EnvironmentConfigEnvironmentSizeEnvironmentSizeSmall = EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_SMALL")
	// The environment size is medium.
	EnvironmentConfigEnvironmentSizeEnvironmentSizeMedium = EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_MEDIUM")
	// The environment size is large.
	EnvironmentConfigEnvironmentSizeEnvironmentSizeLarge = EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_LARGE")
)

func (EnvironmentConfigEnvironmentSize) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentConfigEnvironmentSize)(nil)).Elem()
}

func (e EnvironmentConfigEnvironmentSize) ToEnvironmentConfigEnvironmentSizeOutput() EnvironmentConfigEnvironmentSizeOutput {
	return pulumi.ToOutput(e).(EnvironmentConfigEnvironmentSizeOutput)
}

func (e EnvironmentConfigEnvironmentSize) ToEnvironmentConfigEnvironmentSizeOutputWithContext(ctx context.Context) EnvironmentConfigEnvironmentSizeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnvironmentConfigEnvironmentSizeOutput)
}

func (e EnvironmentConfigEnvironmentSize) ToEnvironmentConfigEnvironmentSizePtrOutput() EnvironmentConfigEnvironmentSizePtrOutput {
	return e.ToEnvironmentConfigEnvironmentSizePtrOutputWithContext(context.Background())
}

func (e EnvironmentConfigEnvironmentSize) ToEnvironmentConfigEnvironmentSizePtrOutputWithContext(ctx context.Context) EnvironmentConfigEnvironmentSizePtrOutput {
	return EnvironmentConfigEnvironmentSize(e).ToEnvironmentConfigEnvironmentSizeOutputWithContext(ctx).ToEnvironmentConfigEnvironmentSizePtrOutputWithContext(ctx)
}

func (e EnvironmentConfigEnvironmentSize) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentConfigEnvironmentSize) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentConfigEnvironmentSize) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnvironmentConfigEnvironmentSize) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnvironmentConfigEnvironmentSizeOutput struct{ *pulumi.OutputState }

func (EnvironmentConfigEnvironmentSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentConfigEnvironmentSize)(nil)).Elem()
}

func (o EnvironmentConfigEnvironmentSizeOutput) ToEnvironmentConfigEnvironmentSizeOutput() EnvironmentConfigEnvironmentSizeOutput {
	return o
}

func (o EnvironmentConfigEnvironmentSizeOutput) ToEnvironmentConfigEnvironmentSizeOutputWithContext(ctx context.Context) EnvironmentConfigEnvironmentSizeOutput {
	return o
}

func (o EnvironmentConfigEnvironmentSizeOutput) ToEnvironmentConfigEnvironmentSizePtrOutput() EnvironmentConfigEnvironmentSizePtrOutput {
	return o.ToEnvironmentConfigEnvironmentSizePtrOutputWithContext(context.Background())
}

func (o EnvironmentConfigEnvironmentSizeOutput) ToEnvironmentConfigEnvironmentSizePtrOutputWithContext(ctx context.Context) EnvironmentConfigEnvironmentSizePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentConfigEnvironmentSize) *EnvironmentConfigEnvironmentSize {
		return &v
	}).(EnvironmentConfigEnvironmentSizePtrOutput)
}

func (o EnvironmentConfigEnvironmentSizeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnvironmentConfigEnvironmentSizeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentConfigEnvironmentSize) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnvironmentConfigEnvironmentSizeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentConfigEnvironmentSizeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentConfigEnvironmentSize) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentConfigEnvironmentSizePtrOutput struct{ *pulumi.OutputState }

func (EnvironmentConfigEnvironmentSizePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentConfigEnvironmentSize)(nil)).Elem()
}

func (o EnvironmentConfigEnvironmentSizePtrOutput) ToEnvironmentConfigEnvironmentSizePtrOutput() EnvironmentConfigEnvironmentSizePtrOutput {
	return o
}

func (o EnvironmentConfigEnvironmentSizePtrOutput) ToEnvironmentConfigEnvironmentSizePtrOutputWithContext(ctx context.Context) EnvironmentConfigEnvironmentSizePtrOutput {
	return o
}

func (o EnvironmentConfigEnvironmentSizePtrOutput) Elem() EnvironmentConfigEnvironmentSizeOutput {
	return o.ApplyT(func(v *EnvironmentConfigEnvironmentSize) EnvironmentConfigEnvironmentSize {
		if v != nil {
			return *v
		}
		var ret EnvironmentConfigEnvironmentSize
		return ret
	}).(EnvironmentConfigEnvironmentSizeOutput)
}

func (o EnvironmentConfigEnvironmentSizePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentConfigEnvironmentSizePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnvironmentConfigEnvironmentSize) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EnvironmentConfigEnvironmentSizeInput is an input type that accepts EnvironmentConfigEnvironmentSizeArgs and EnvironmentConfigEnvironmentSizeOutput values.
// You can construct a concrete instance of `EnvironmentConfigEnvironmentSizeInput` via:
//
//	EnvironmentConfigEnvironmentSizeArgs{...}
type EnvironmentConfigEnvironmentSizeInput interface {
	pulumi.Input

	ToEnvironmentConfigEnvironmentSizeOutput() EnvironmentConfigEnvironmentSizeOutput
	ToEnvironmentConfigEnvironmentSizeOutputWithContext(context.Context) EnvironmentConfigEnvironmentSizeOutput
}

var environmentConfigEnvironmentSizePtrType = reflect.TypeOf((**EnvironmentConfigEnvironmentSize)(nil)).Elem()

type EnvironmentConfigEnvironmentSizePtrInput interface {
	pulumi.Input

	ToEnvironmentConfigEnvironmentSizePtrOutput() EnvironmentConfigEnvironmentSizePtrOutput
	ToEnvironmentConfigEnvironmentSizePtrOutputWithContext(context.Context) EnvironmentConfigEnvironmentSizePtrOutput
}

type environmentConfigEnvironmentSizePtr string

func EnvironmentConfigEnvironmentSizePtr(v string) EnvironmentConfigEnvironmentSizePtrInput {
	return (*environmentConfigEnvironmentSizePtr)(&v)
}

func (*environmentConfigEnvironmentSizePtr) ElementType() reflect.Type {
	return environmentConfigEnvironmentSizePtrType
}

func (in *environmentConfigEnvironmentSizePtr) ToEnvironmentConfigEnvironmentSizePtrOutput() EnvironmentConfigEnvironmentSizePtrOutput {
	return pulumi.ToOutput(in).(EnvironmentConfigEnvironmentSizePtrOutput)
}

func (in *environmentConfigEnvironmentSizePtr) ToEnvironmentConfigEnvironmentSizePtrOutputWithContext(ctx context.Context) EnvironmentConfigEnvironmentSizePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnvironmentConfigEnvironmentSizePtrOutput)
}

// The current state of the environment.
type EnvironmentStateEnum string

const (
	// The state of the environment is unknown.
	EnvironmentStateEnumStateUnspecified = EnvironmentStateEnum("STATE_UNSPECIFIED")
	// The environment is in the process of being created.
	EnvironmentStateEnumCreating = EnvironmentStateEnum("CREATING")
	// The environment is currently running and healthy. It is ready for use.
	EnvironmentStateEnumRunning = EnvironmentStateEnum("RUNNING")
	// The environment is being updated. It remains usable but cannot receive additional update requests or be deleted at this time.
	EnvironmentStateEnumUpdating = EnvironmentStateEnum("UPDATING")
	// The environment is undergoing deletion. It cannot be used.
	EnvironmentStateEnumDeleting = EnvironmentStateEnum("DELETING")
	// The environment has encountered an error and cannot be used.
	EnvironmentStateEnumError = EnvironmentStateEnum("ERROR")
)

func (EnvironmentStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStateEnum)(nil)).Elem()
}

func (e EnvironmentStateEnum) ToEnvironmentStateEnumOutput() EnvironmentStateEnumOutput {
	return pulumi.ToOutput(e).(EnvironmentStateEnumOutput)
}

func (e EnvironmentStateEnum) ToEnvironmentStateEnumOutputWithContext(ctx context.Context) EnvironmentStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnvironmentStateEnumOutput)
}

func (e EnvironmentStateEnum) ToEnvironmentStateEnumPtrOutput() EnvironmentStateEnumPtrOutput {
	return e.ToEnvironmentStateEnumPtrOutputWithContext(context.Background())
}

func (e EnvironmentStateEnum) ToEnvironmentStateEnumPtrOutputWithContext(ctx context.Context) EnvironmentStateEnumPtrOutput {
	return EnvironmentStateEnum(e).ToEnvironmentStateEnumOutputWithContext(ctx).ToEnvironmentStateEnumPtrOutputWithContext(ctx)
}

func (e EnvironmentStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnvironmentStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnvironmentStateEnumOutput struct{ *pulumi.OutputState }

func (EnvironmentStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStateEnum)(nil)).Elem()
}

func (o EnvironmentStateEnumOutput) ToEnvironmentStateEnumOutput() EnvironmentStateEnumOutput {
	return o
}

func (o EnvironmentStateEnumOutput) ToEnvironmentStateEnumOutputWithContext(ctx context.Context) EnvironmentStateEnumOutput {
	return o
}

func (o EnvironmentStateEnumOutput) ToEnvironmentStateEnumPtrOutput() EnvironmentStateEnumPtrOutput {
	return o.ToEnvironmentStateEnumPtrOutputWithContext(context.Background())
}

func (o EnvironmentStateEnumOutput) ToEnvironmentStateEnumPtrOutputWithContext(ctx context.Context) EnvironmentStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentStateEnum) *EnvironmentStateEnum {
		return &v
	}).(EnvironmentStateEnumPtrOutput)
}

func (o EnvironmentStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnvironmentStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnvironmentStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentStateEnumPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentStateEnum)(nil)).Elem()
}

func (o EnvironmentStateEnumPtrOutput) ToEnvironmentStateEnumPtrOutput() EnvironmentStateEnumPtrOutput {
	return o
}

func (o EnvironmentStateEnumPtrOutput) ToEnvironmentStateEnumPtrOutputWithContext(ctx context.Context) EnvironmentStateEnumPtrOutput {
	return o
}

func (o EnvironmentStateEnumPtrOutput) Elem() EnvironmentStateEnumOutput {
	return o.ApplyT(func(v *EnvironmentStateEnum) EnvironmentStateEnum {
		if v != nil {
			return *v
		}
		var ret EnvironmentStateEnum
		return ret
	}).(EnvironmentStateEnumOutput)
}

func (o EnvironmentStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnvironmentStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EnvironmentStateEnumInput is an input type that accepts EnvironmentStateEnumArgs and EnvironmentStateEnumOutput values.
// You can construct a concrete instance of `EnvironmentStateEnumInput` via:
//
//	EnvironmentStateEnumArgs{...}
type EnvironmentStateEnumInput interface {
	pulumi.Input

	ToEnvironmentStateEnumOutput() EnvironmentStateEnumOutput
	ToEnvironmentStateEnumOutputWithContext(context.Context) EnvironmentStateEnumOutput
}

var environmentStateEnumPtrType = reflect.TypeOf((**EnvironmentStateEnum)(nil)).Elem()

type EnvironmentStateEnumPtrInput interface {
	pulumi.Input

	ToEnvironmentStateEnumPtrOutput() EnvironmentStateEnumPtrOutput
	ToEnvironmentStateEnumPtrOutputWithContext(context.Context) EnvironmentStateEnumPtrOutput
}

type environmentStateEnumPtr string

func EnvironmentStateEnumPtr(v string) EnvironmentStateEnumPtrInput {
	return (*environmentStateEnumPtr)(&v)
}

func (*environmentStateEnumPtr) ElementType() reflect.Type {
	return environmentStateEnumPtrType
}

func (in *environmentStateEnumPtr) ToEnvironmentStateEnumPtrOutput() EnvironmentStateEnumPtrOutput {
	return pulumi.ToOutput(in).(EnvironmentStateEnumPtrOutput)
}

func (in *environmentStateEnumPtr) ToEnvironmentStateEnumPtrOutputWithContext(ctx context.Context) EnvironmentStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnvironmentStateEnumPtrOutput)
}

// Optional. Indicates the user requested specifc connection type between Tenant and Customer projects. You cannot set networking connection type in public IP environment.
type NetworkingConfigConnectionType string

const (
	// No specific connection type was requested, so the environment uses the default value corresponding to the rest of its configuration.
	NetworkingConfigConnectionTypeConnectionTypeUnspecified = NetworkingConfigConnectionType("CONNECTION_TYPE_UNSPECIFIED")
	// Requests the use of VPC peerings for connecting the Customer and Tenant projects.
	NetworkingConfigConnectionTypeVpcPeering = NetworkingConfigConnectionType("VPC_PEERING")
	// Requests the use of Private Service Connect for connecting the Customer and Tenant projects.
	NetworkingConfigConnectionTypePrivateServiceConnect = NetworkingConfigConnectionType("PRIVATE_SERVICE_CONNECT")
)

func (NetworkingConfigConnectionType) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkingConfigConnectionType)(nil)).Elem()
}

func (e NetworkingConfigConnectionType) ToNetworkingConfigConnectionTypeOutput() NetworkingConfigConnectionTypeOutput {
	return pulumi.ToOutput(e).(NetworkingConfigConnectionTypeOutput)
}

func (e NetworkingConfigConnectionType) ToNetworkingConfigConnectionTypeOutputWithContext(ctx context.Context) NetworkingConfigConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkingConfigConnectionTypeOutput)
}

func (e NetworkingConfigConnectionType) ToNetworkingConfigConnectionTypePtrOutput() NetworkingConfigConnectionTypePtrOutput {
	return e.ToNetworkingConfigConnectionTypePtrOutputWithContext(context.Background())
}

func (e NetworkingConfigConnectionType) ToNetworkingConfigConnectionTypePtrOutputWithContext(ctx context.Context) NetworkingConfigConnectionTypePtrOutput {
	return NetworkingConfigConnectionType(e).ToNetworkingConfigConnectionTypeOutputWithContext(ctx).ToNetworkingConfigConnectionTypePtrOutputWithContext(ctx)
}

func (e NetworkingConfigConnectionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkingConfigConnectionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkingConfigConnectionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkingConfigConnectionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkingConfigConnectionTypeOutput struct{ *pulumi.OutputState }

func (NetworkingConfigConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkingConfigConnectionType)(nil)).Elem()
}

func (o NetworkingConfigConnectionTypeOutput) ToNetworkingConfigConnectionTypeOutput() NetworkingConfigConnectionTypeOutput {
	return o
}

func (o NetworkingConfigConnectionTypeOutput) ToNetworkingConfigConnectionTypeOutputWithContext(ctx context.Context) NetworkingConfigConnectionTypeOutput {
	return o
}

func (o NetworkingConfigConnectionTypeOutput) ToNetworkingConfigConnectionTypePtrOutput() NetworkingConfigConnectionTypePtrOutput {
	return o.ToNetworkingConfigConnectionTypePtrOutputWithContext(context.Background())
}

func (o NetworkingConfigConnectionTypeOutput) ToNetworkingConfigConnectionTypePtrOutputWithContext(ctx context.Context) NetworkingConfigConnectionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkingConfigConnectionType) *NetworkingConfigConnectionType {
		return &v
	}).(NetworkingConfigConnectionTypePtrOutput)
}

func (o NetworkingConfigConnectionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkingConfigConnectionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkingConfigConnectionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkingConfigConnectionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkingConfigConnectionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkingConfigConnectionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkingConfigConnectionTypePtrOutput struct{ *pulumi.OutputState }

func (NetworkingConfigConnectionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingConfigConnectionType)(nil)).Elem()
}

func (o NetworkingConfigConnectionTypePtrOutput) ToNetworkingConfigConnectionTypePtrOutput() NetworkingConfigConnectionTypePtrOutput {
	return o
}

func (o NetworkingConfigConnectionTypePtrOutput) ToNetworkingConfigConnectionTypePtrOutputWithContext(ctx context.Context) NetworkingConfigConnectionTypePtrOutput {
	return o
}

func (o NetworkingConfigConnectionTypePtrOutput) Elem() NetworkingConfigConnectionTypeOutput {
	return o.ApplyT(func(v *NetworkingConfigConnectionType) NetworkingConfigConnectionType {
		if v != nil {
			return *v
		}
		var ret NetworkingConfigConnectionType
		return ret
	}).(NetworkingConfigConnectionTypeOutput)
}

func (o NetworkingConfigConnectionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkingConfigConnectionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkingConfigConnectionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NetworkingConfigConnectionTypeInput is an input type that accepts NetworkingConfigConnectionTypeArgs and NetworkingConfigConnectionTypeOutput values.
// You can construct a concrete instance of `NetworkingConfigConnectionTypeInput` via:
//
//	NetworkingConfigConnectionTypeArgs{...}
type NetworkingConfigConnectionTypeInput interface {
	pulumi.Input

	ToNetworkingConfigConnectionTypeOutput() NetworkingConfigConnectionTypeOutput
	ToNetworkingConfigConnectionTypeOutputWithContext(context.Context) NetworkingConfigConnectionTypeOutput
}

var networkingConfigConnectionTypePtrType = reflect.TypeOf((**NetworkingConfigConnectionType)(nil)).Elem()

type NetworkingConfigConnectionTypePtrInput interface {
	pulumi.Input

	ToNetworkingConfigConnectionTypePtrOutput() NetworkingConfigConnectionTypePtrOutput
	ToNetworkingConfigConnectionTypePtrOutputWithContext(context.Context) NetworkingConfigConnectionTypePtrOutput
}

type networkingConfigConnectionTypePtr string

func NetworkingConfigConnectionTypePtr(v string) NetworkingConfigConnectionTypePtrInput {
	return (*networkingConfigConnectionTypePtr)(&v)
}

func (*networkingConfigConnectionTypePtr) ElementType() reflect.Type {
	return networkingConfigConnectionTypePtrType
}

func (in *networkingConfigConnectionTypePtr) ToNetworkingConfigConnectionTypePtrOutput() NetworkingConfigConnectionTypePtrOutput {
	return pulumi.ToOutput(in).(NetworkingConfigConnectionTypePtrOutput)
}

func (in *networkingConfigConnectionTypePtr) ToNetworkingConfigConnectionTypePtrOutputWithContext(ctx context.Context) NetworkingConfigConnectionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkingConfigConnectionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentConfigEnvironmentSizeInput)(nil)).Elem(), EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentConfigEnvironmentSizePtrInput)(nil)).Elem(), EnvironmentConfigEnvironmentSize("ENVIRONMENT_SIZE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentStateEnumInput)(nil)).Elem(), EnvironmentStateEnum("STATE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentStateEnumPtrInput)(nil)).Elem(), EnvironmentStateEnum("STATE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingConfigConnectionTypeInput)(nil)).Elem(), NetworkingConfigConnectionType("CONNECTION_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingConfigConnectionTypePtrInput)(nil)).Elem(), NetworkingConfigConnectionType("CONNECTION_TYPE_UNSPECIFIED"))
	pulumi.RegisterOutputType(EnvironmentConfigEnvironmentSizeOutput{})
	pulumi.RegisterOutputType(EnvironmentConfigEnvironmentSizePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentStateEnumOutput{})
	pulumi.RegisterOutputType(EnvironmentStateEnumPtrOutput{})
	pulumi.RegisterOutputType(NetworkingConfigConnectionTypeOutput{})
	pulumi.RegisterOutputType(NetworkingConfigConnectionTypePtrOutput{})
}
