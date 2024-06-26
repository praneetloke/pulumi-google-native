// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

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

// AuditLogConfigLogTypeInput is an input type that accepts values of the AuditLogConfigLogType enum
// A concrete instance of `AuditLogConfigLogTypeInput` can be one of the following:
//
//	AuditLogConfigLogTypeLogTypeUnspecified
//	AuditLogConfigLogTypeAdminRead
//	AuditLogConfigLogTypeDataWrite
//	AuditLogConfigLogTypeDataRead
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

// The format of the Cloud Pub/Sub messages.
type PubsubConfigMessageFormat string

const (
	// Unspecified.
	PubsubConfigMessageFormatMessageFormatUnspecified = PubsubConfigMessageFormat("MESSAGE_FORMAT_UNSPECIFIED")
	// The message payload is a serialized protocol buffer of SourceRepoEvent.
	PubsubConfigMessageFormatProtobuf = PubsubConfigMessageFormat("PROTOBUF")
	// The message payload is a JSON string of SourceRepoEvent.
	PubsubConfigMessageFormatJson = PubsubConfigMessageFormat("JSON")
)

func (PubsubConfigMessageFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*PubsubConfigMessageFormat)(nil)).Elem()
}

func (e PubsubConfigMessageFormat) ToPubsubConfigMessageFormatOutput() PubsubConfigMessageFormatOutput {
	return pulumi.ToOutput(e).(PubsubConfigMessageFormatOutput)
}

func (e PubsubConfigMessageFormat) ToPubsubConfigMessageFormatOutputWithContext(ctx context.Context) PubsubConfigMessageFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PubsubConfigMessageFormatOutput)
}

func (e PubsubConfigMessageFormat) ToPubsubConfigMessageFormatPtrOutput() PubsubConfigMessageFormatPtrOutput {
	return e.ToPubsubConfigMessageFormatPtrOutputWithContext(context.Background())
}

func (e PubsubConfigMessageFormat) ToPubsubConfigMessageFormatPtrOutputWithContext(ctx context.Context) PubsubConfigMessageFormatPtrOutput {
	return PubsubConfigMessageFormat(e).ToPubsubConfigMessageFormatOutputWithContext(ctx).ToPubsubConfigMessageFormatPtrOutputWithContext(ctx)
}

func (e PubsubConfigMessageFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PubsubConfigMessageFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PubsubConfigMessageFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PubsubConfigMessageFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PubsubConfigMessageFormatOutput struct{ *pulumi.OutputState }

func (PubsubConfigMessageFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PubsubConfigMessageFormat)(nil)).Elem()
}

func (o PubsubConfigMessageFormatOutput) ToPubsubConfigMessageFormatOutput() PubsubConfigMessageFormatOutput {
	return o
}

func (o PubsubConfigMessageFormatOutput) ToPubsubConfigMessageFormatOutputWithContext(ctx context.Context) PubsubConfigMessageFormatOutput {
	return o
}

func (o PubsubConfigMessageFormatOutput) ToPubsubConfigMessageFormatPtrOutput() PubsubConfigMessageFormatPtrOutput {
	return o.ToPubsubConfigMessageFormatPtrOutputWithContext(context.Background())
}

func (o PubsubConfigMessageFormatOutput) ToPubsubConfigMessageFormatPtrOutputWithContext(ctx context.Context) PubsubConfigMessageFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PubsubConfigMessageFormat) *PubsubConfigMessageFormat {
		return &v
	}).(PubsubConfigMessageFormatPtrOutput)
}

func (o PubsubConfigMessageFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PubsubConfigMessageFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PubsubConfigMessageFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PubsubConfigMessageFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PubsubConfigMessageFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PubsubConfigMessageFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PubsubConfigMessageFormatPtrOutput struct{ *pulumi.OutputState }

func (PubsubConfigMessageFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PubsubConfigMessageFormat)(nil)).Elem()
}

func (o PubsubConfigMessageFormatPtrOutput) ToPubsubConfigMessageFormatPtrOutput() PubsubConfigMessageFormatPtrOutput {
	return o
}

func (o PubsubConfigMessageFormatPtrOutput) ToPubsubConfigMessageFormatPtrOutputWithContext(ctx context.Context) PubsubConfigMessageFormatPtrOutput {
	return o
}

func (o PubsubConfigMessageFormatPtrOutput) Elem() PubsubConfigMessageFormatOutput {
	return o.ApplyT(func(v *PubsubConfigMessageFormat) PubsubConfigMessageFormat {
		if v != nil {
			return *v
		}
		var ret PubsubConfigMessageFormat
		return ret
	}).(PubsubConfigMessageFormatOutput)
}

func (o PubsubConfigMessageFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PubsubConfigMessageFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PubsubConfigMessageFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PubsubConfigMessageFormatInput is an input type that accepts values of the PubsubConfigMessageFormat enum
// A concrete instance of `PubsubConfigMessageFormatInput` can be one of the following:
//
//	PubsubConfigMessageFormatMessageFormatUnspecified
//	PubsubConfigMessageFormatProtobuf
//	PubsubConfigMessageFormatJson
type PubsubConfigMessageFormatInput interface {
	pulumi.Input

	ToPubsubConfigMessageFormatOutput() PubsubConfigMessageFormatOutput
	ToPubsubConfigMessageFormatOutputWithContext(context.Context) PubsubConfigMessageFormatOutput
}

var pubsubConfigMessageFormatPtrType = reflect.TypeOf((**PubsubConfigMessageFormat)(nil)).Elem()

type PubsubConfigMessageFormatPtrInput interface {
	pulumi.Input

	ToPubsubConfigMessageFormatPtrOutput() PubsubConfigMessageFormatPtrOutput
	ToPubsubConfigMessageFormatPtrOutputWithContext(context.Context) PubsubConfigMessageFormatPtrOutput
}

type pubsubConfigMessageFormatPtr string

func PubsubConfigMessageFormatPtr(v string) PubsubConfigMessageFormatPtrInput {
	return (*pubsubConfigMessageFormatPtr)(&v)
}

func (*pubsubConfigMessageFormatPtr) ElementType() reflect.Type {
	return pubsubConfigMessageFormatPtrType
}

func (in *pubsubConfigMessageFormatPtr) ToPubsubConfigMessageFormatPtrOutput() PubsubConfigMessageFormatPtrOutput {
	return pulumi.ToOutput(in).(PubsubConfigMessageFormatPtrOutput)
}

func (in *pubsubConfigMessageFormatPtr) ToPubsubConfigMessageFormatPtrOutputWithContext(ctx context.Context) PubsubConfigMessageFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PubsubConfigMessageFormatPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypeInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypePtrInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*PubsubConfigMessageFormatInput)(nil)).Elem(), PubsubConfigMessageFormat("MESSAGE_FORMAT_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*PubsubConfigMessageFormatPtrInput)(nil)).Elem(), PubsubConfigMessageFormat("MESSAGE_FORMAT_UNSPECIFIED"))
	pulumi.RegisterOutputType(AuditLogConfigLogTypeOutput{})
	pulumi.RegisterOutputType(AuditLogConfigLogTypePtrOutput{})
	pulumi.RegisterOutputType(PubsubConfigMessageFormatOutput{})
	pulumi.RegisterOutputType(PubsubConfigMessageFormatPtrOutput{})
}
