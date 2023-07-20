// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

// A field of an index.
type GoogleFirestoreAdminV1beta1IndexField struct {
	// The path of the field. Must match the field path specification described by google.firestore.v1beta1.Document.fields. Special field path `__name__` may be used by itself or at the end of a path. `__type__` may be used only at the end of path.
	FieldPath *string `pulumi:"fieldPath"`
	// The field's mode.
	Mode *GoogleFirestoreAdminV1beta1IndexFieldMode `pulumi:"mode"`
}

// GoogleFirestoreAdminV1beta1IndexFieldInput is an input type that accepts GoogleFirestoreAdminV1beta1IndexFieldArgs and GoogleFirestoreAdminV1beta1IndexFieldOutput values.
// You can construct a concrete instance of `GoogleFirestoreAdminV1beta1IndexFieldInput` via:
//
//	GoogleFirestoreAdminV1beta1IndexFieldArgs{...}
type GoogleFirestoreAdminV1beta1IndexFieldInput interface {
	pulumi.Input

	ToGoogleFirestoreAdminV1beta1IndexFieldOutput() GoogleFirestoreAdminV1beta1IndexFieldOutput
	ToGoogleFirestoreAdminV1beta1IndexFieldOutputWithContext(context.Context) GoogleFirestoreAdminV1beta1IndexFieldOutput
}

// A field of an index.
type GoogleFirestoreAdminV1beta1IndexFieldArgs struct {
	// The path of the field. Must match the field path specification described by google.firestore.v1beta1.Document.fields. Special field path `__name__` may be used by itself or at the end of a path. `__type__` may be used only at the end of path.
	FieldPath pulumi.StringPtrInput `pulumi:"fieldPath"`
	// The field's mode.
	Mode GoogleFirestoreAdminV1beta1IndexFieldModePtrInput `pulumi:"mode"`
}

func (GoogleFirestoreAdminV1beta1IndexFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta1IndexField)(nil)).Elem()
}

func (i GoogleFirestoreAdminV1beta1IndexFieldArgs) ToGoogleFirestoreAdminV1beta1IndexFieldOutput() GoogleFirestoreAdminV1beta1IndexFieldOutput {
	return i.ToGoogleFirestoreAdminV1beta1IndexFieldOutputWithContext(context.Background())
}

func (i GoogleFirestoreAdminV1beta1IndexFieldArgs) ToGoogleFirestoreAdminV1beta1IndexFieldOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta1IndexFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleFirestoreAdminV1beta1IndexFieldOutput)
}

// GoogleFirestoreAdminV1beta1IndexFieldArrayInput is an input type that accepts GoogleFirestoreAdminV1beta1IndexFieldArray and GoogleFirestoreAdminV1beta1IndexFieldArrayOutput values.
// You can construct a concrete instance of `GoogleFirestoreAdminV1beta1IndexFieldArrayInput` via:
//
//	GoogleFirestoreAdminV1beta1IndexFieldArray{ GoogleFirestoreAdminV1beta1IndexFieldArgs{...} }
type GoogleFirestoreAdminV1beta1IndexFieldArrayInput interface {
	pulumi.Input

	ToGoogleFirestoreAdminV1beta1IndexFieldArrayOutput() GoogleFirestoreAdminV1beta1IndexFieldArrayOutput
	ToGoogleFirestoreAdminV1beta1IndexFieldArrayOutputWithContext(context.Context) GoogleFirestoreAdminV1beta1IndexFieldArrayOutput
}

type GoogleFirestoreAdminV1beta1IndexFieldArray []GoogleFirestoreAdminV1beta1IndexFieldInput

func (GoogleFirestoreAdminV1beta1IndexFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleFirestoreAdminV1beta1IndexField)(nil)).Elem()
}

func (i GoogleFirestoreAdminV1beta1IndexFieldArray) ToGoogleFirestoreAdminV1beta1IndexFieldArrayOutput() GoogleFirestoreAdminV1beta1IndexFieldArrayOutput {
	return i.ToGoogleFirestoreAdminV1beta1IndexFieldArrayOutputWithContext(context.Background())
}

func (i GoogleFirestoreAdminV1beta1IndexFieldArray) ToGoogleFirestoreAdminV1beta1IndexFieldArrayOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta1IndexFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleFirestoreAdminV1beta1IndexFieldArrayOutput)
}

// A field of an index.
type GoogleFirestoreAdminV1beta1IndexFieldOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta1IndexFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta1IndexField)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta1IndexFieldOutput) ToGoogleFirestoreAdminV1beta1IndexFieldOutput() GoogleFirestoreAdminV1beta1IndexFieldOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta1IndexFieldOutput) ToGoogleFirestoreAdminV1beta1IndexFieldOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta1IndexFieldOutput {
	return o
}

// The path of the field. Must match the field path specification described by google.firestore.v1beta1.Document.fields. Special field path `__name__` may be used by itself or at the end of a path. `__type__` may be used only at the end of path.
func (o GoogleFirestoreAdminV1beta1IndexFieldOutput) FieldPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta1IndexField) *string { return v.FieldPath }).(pulumi.StringPtrOutput)
}

// The field's mode.
func (o GoogleFirestoreAdminV1beta1IndexFieldOutput) Mode() GoogleFirestoreAdminV1beta1IndexFieldModePtrOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta1IndexField) *GoogleFirestoreAdminV1beta1IndexFieldMode {
		return v.Mode
	}).(GoogleFirestoreAdminV1beta1IndexFieldModePtrOutput)
}

type GoogleFirestoreAdminV1beta1IndexFieldArrayOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta1IndexFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleFirestoreAdminV1beta1IndexField)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta1IndexFieldArrayOutput) ToGoogleFirestoreAdminV1beta1IndexFieldArrayOutput() GoogleFirestoreAdminV1beta1IndexFieldArrayOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta1IndexFieldArrayOutput) ToGoogleFirestoreAdminV1beta1IndexFieldArrayOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta1IndexFieldArrayOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta1IndexFieldArrayOutput) Index(i pulumi.IntInput) GoogleFirestoreAdminV1beta1IndexFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GoogleFirestoreAdminV1beta1IndexField {
		return vs[0].([]GoogleFirestoreAdminV1beta1IndexField)[vs[1].(int)]
	}).(GoogleFirestoreAdminV1beta1IndexFieldOutput)
}

// A field of an index.
type GoogleFirestoreAdminV1beta1IndexFieldResponse struct {
	// The path of the field. Must match the field path specification described by google.firestore.v1beta1.Document.fields. Special field path `__name__` may be used by itself or at the end of a path. `__type__` may be used only at the end of path.
	FieldPath string `pulumi:"fieldPath"`
	// The field's mode.
	Mode string `pulumi:"mode"`
}

// A field of an index.
type GoogleFirestoreAdminV1beta1IndexFieldResponseOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta1IndexFieldResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta1IndexFieldResponse)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta1IndexFieldResponseOutput) ToGoogleFirestoreAdminV1beta1IndexFieldResponseOutput() GoogleFirestoreAdminV1beta1IndexFieldResponseOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta1IndexFieldResponseOutput) ToGoogleFirestoreAdminV1beta1IndexFieldResponseOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta1IndexFieldResponseOutput {
	return o
}

// The path of the field. Must match the field path specification described by google.firestore.v1beta1.Document.fields. Special field path `__name__` may be used by itself or at the end of a path. `__type__` may be used only at the end of path.
func (o GoogleFirestoreAdminV1beta1IndexFieldResponseOutput) FieldPath() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta1IndexFieldResponse) string { return v.FieldPath }).(pulumi.StringOutput)
}

// The field's mode.
func (o GoogleFirestoreAdminV1beta1IndexFieldResponseOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta1IndexFieldResponse) string { return v.Mode }).(pulumi.StringOutput)
}

type GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleFirestoreAdminV1beta1IndexFieldResponse)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput) ToGoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput() GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput) ToGoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput) Index(i pulumi.IntInput) GoogleFirestoreAdminV1beta1IndexFieldResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GoogleFirestoreAdminV1beta1IndexFieldResponse {
		return vs[0].([]GoogleFirestoreAdminV1beta1IndexFieldResponse)[vs[1].(int)]
	}).(GoogleFirestoreAdminV1beta1IndexFieldResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleFirestoreAdminV1beta1IndexFieldInput)(nil)).Elem(), GoogleFirestoreAdminV1beta1IndexFieldArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleFirestoreAdminV1beta1IndexFieldArrayInput)(nil)).Elem(), GoogleFirestoreAdminV1beta1IndexFieldArray{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta1IndexFieldOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta1IndexFieldArrayOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta1IndexFieldResponseOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta1IndexFieldResponseArrayOutput{})
}
