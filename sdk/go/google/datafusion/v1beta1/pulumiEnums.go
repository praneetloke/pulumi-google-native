// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The log type that this config enables.
type AuditLogConfigLogType string

const (
	// Default case. Should never be this.
	AuditLogConfigLogTypeLogTypeUnspecified = AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED")
	// Admin reads. Example: CloudIAM getIamPolicy
	AuditLogConfigLogTypeAdminRead = AuditLogConfigLogType("ADMIN_READ")
	// Data writes. Example: CloudSQL Users create
	AuditLogConfigLogTypeDataWrite = AuditLogConfigLogType("DATA_WRITE")
	// Data reads. Example: CloudSQL Users list
	AuditLogConfigLogTypeDataRead = AuditLogConfigLogType("DATA_READ")
)

func (AuditLogConfigLogType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return pulumi.ToOutput(e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return e.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return AuditLogConfigLogType(e).ToAuditLogConfigLogTypeOutputWithContext(ctx).ToAuditLogConfigLogTypePtrOutputWithContext(ctx)
}

func (e AuditLogConfigLogType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuditLogConfigLogTypeOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuditLogConfigLogType) *AuditLogConfigLogType {
		return &v
	}).(AuditLogConfigLogTypePtrOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuditLogConfigLogTypePtrOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) Elem() AuditLogConfigLogTypeOutput {
	return o.ApplyT(func(v *AuditLogConfigLogType) AuditLogConfigLogType {
		if v != nil {
			return *v
		}
		var ret AuditLogConfigLogType
		return ret
	}).(AuditLogConfigLogTypeOutput)
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuditLogConfigLogType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AuditLogConfigLogTypeInput is an input type that accepts AuditLogConfigLogTypeArgs and AuditLogConfigLogTypeOutput values.
// You can construct a concrete instance of `AuditLogConfigLogTypeInput` via:
//
//	AuditLogConfigLogTypeArgs{...}
type AuditLogConfigLogTypeInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput
	ToAuditLogConfigLogTypeOutputWithContext(context.Context) AuditLogConfigLogTypeOutput
}

var auditLogConfigLogTypePtrType = reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()

type AuditLogConfigLogTypePtrInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput
	ToAuditLogConfigLogTypePtrOutputWithContext(context.Context) AuditLogConfigLogTypePtrOutput
}

type auditLogConfigLogTypePtr string

func AuditLogConfigLogTypePtr(v string) AuditLogConfigLogTypePtrInput {
	return (*auditLogConfigLogTypePtr)(&v)
}

func (*auditLogConfigLogTypePtr) ElementType() reflect.Type {
	return auditLogConfigLogTypePtrType
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutput(in).(AuditLogConfigLogTypePtrOutput)
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuditLogConfigLogTypePtrOutput)
}

// Required. Instance type.
type InstanceType string

const (
	// No type specified. The instance creation will fail.
	InstanceTypeTypeUnspecified = InstanceType("TYPE_UNSPECIFIED")
	// Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines using point and click UI. However, there are certain limitations, such as fewer number of concurrent pipelines, no support for streaming pipelines, etc.
	InstanceTypeBasic = InstanceType("BASIC")
	// Enterprise Data Fusion instance. In Enterprise type, the user will have all features available, such as support for streaming pipelines, unlimited number of concurrent pipelines, etc.
	InstanceTypeEnterprise = InstanceType("ENTERPRISE")
	// Developer Data Fusion instance. In Developer type, the user will have all features available but with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration pipelines at low cost.
	InstanceTypeDeveloper = InstanceType("DEVELOPER")
)

func (InstanceType) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceType)(nil)).Elem()
}

func (e InstanceType) ToInstanceTypeOutput() InstanceTypeOutput {
	return pulumi.ToOutput(e).(InstanceTypeOutput)
}

func (e InstanceType) ToInstanceTypeOutputWithContext(ctx context.Context) InstanceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InstanceTypeOutput)
}

func (e InstanceType) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return e.ToInstanceTypePtrOutputWithContext(context.Background())
}

func (e InstanceType) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return InstanceType(e).ToInstanceTypeOutputWithContext(ctx).ToInstanceTypePtrOutputWithContext(ctx)
}

func (e InstanceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstanceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstanceTypeOutput struct{ *pulumi.OutputState }

func (InstanceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceType)(nil)).Elem()
}

func (o InstanceTypeOutput) ToInstanceTypeOutput() InstanceTypeOutput {
	return o
}

func (o InstanceTypeOutput) ToInstanceTypeOutputWithContext(ctx context.Context) InstanceTypeOutput {
	return o
}

func (o InstanceTypeOutput) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return o.ToInstanceTypePtrOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceType) *InstanceType {
		return &v
	}).(InstanceTypePtrOutput)
}

func (o InstanceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InstanceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InstanceTypePtrOutput struct{ *pulumi.OutputState }

func (InstanceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceType)(nil)).Elem()
}

func (o InstanceTypePtrOutput) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return o
}

func (o InstanceTypePtrOutput) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return o
}

func (o InstanceTypePtrOutput) Elem() InstanceTypeOutput {
	return o.ApplyT(func(v *InstanceType) InstanceType {
		if v != nil {
			return *v
		}
		var ret InstanceType
		return ret
	}).(InstanceTypeOutput)
}

func (o InstanceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InstanceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InstanceTypeInput is an input type that accepts InstanceTypeArgs and InstanceTypeOutput values.
// You can construct a concrete instance of `InstanceTypeInput` via:
//
//	InstanceTypeArgs{...}
type InstanceTypeInput interface {
	pulumi.Input

	ToInstanceTypeOutput() InstanceTypeOutput
	ToInstanceTypeOutputWithContext(context.Context) InstanceTypeOutput
}

var instanceTypePtrType = reflect.TypeOf((**InstanceType)(nil)).Elem()

type InstanceTypePtrInput interface {
	pulumi.Input

	ToInstanceTypePtrOutput() InstanceTypePtrOutput
	ToInstanceTypePtrOutputWithContext(context.Context) InstanceTypePtrOutput
}

type instanceTypePtr string

func InstanceTypePtr(v string) InstanceTypePtrInput {
	return (*instanceTypePtr)(&v)
}

func (*instanceTypePtr) ElementType() reflect.Type {
	return instanceTypePtrType
}

func (in *instanceTypePtr) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return pulumi.ToOutput(in).(InstanceTypePtrOutput)
}

func (in *instanceTypePtr) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InstanceTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypeInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypePtrInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTypeInput)(nil)).Elem(), InstanceType("TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTypePtrInput)(nil)).Elem(), InstanceType("TYPE_UNSPECIFIED"))
	pulumi.RegisterOutputType(AuditLogConfigLogTypeOutput{})
	pulumi.RegisterOutputType(AuditLogConfigLogTypePtrOutput{})
	pulumi.RegisterOutputType(InstanceTypeOutput{})
	pulumi.RegisterOutputType(InstanceTypePtrOutput{})
}
