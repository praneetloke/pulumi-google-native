// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type KeyRingCryptoKeyIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings    BindingResponseArrayOutput `pulumi:"bindings"`
	CryptoKeyId pulumi.StringOutput        `pulumi:"cryptoKeyId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag      pulumi.StringOutput `pulumi:"etag"`
	KeyRingId pulumi.StringOutput `pulumi:"keyRingId"`
	Location  pulumi.StringOutput `pulumi:"location"`
	Project   pulumi.StringOutput `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewKeyRingCryptoKeyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewKeyRingCryptoKeyIamPolicy(ctx *pulumi.Context,
	name string, args *KeyRingCryptoKeyIamPolicyArgs, opts ...pulumi.ResourceOption) (*KeyRingCryptoKeyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CryptoKeyId == nil {
		return nil, errors.New("invalid value for required argument 'CryptoKeyId'")
	}
	if args.KeyRingId == nil {
		return nil, errors.New("invalid value for required argument 'KeyRingId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"cryptoKeyId",
		"keyRingId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KeyRingCryptoKeyIamPolicy
	err := ctx.RegisterResource("google-native:cloudkms/v1:KeyRingCryptoKeyIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyRingCryptoKeyIamPolicy gets an existing KeyRingCryptoKeyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRingCryptoKeyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyRingCryptoKeyIamPolicyState, opts ...pulumi.ResourceOption) (*KeyRingCryptoKeyIamPolicy, error) {
	var resource KeyRingCryptoKeyIamPolicy
	err := ctx.ReadResource("google-native:cloudkms/v1:KeyRingCryptoKeyIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyRingCryptoKeyIamPolicy resources.
type keyRingCryptoKeyIamPolicyState struct {
}

type KeyRingCryptoKeyIamPolicyState struct {
}

func (KeyRingCryptoKeyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingCryptoKeyIamPolicyState)(nil)).Elem()
}

type keyRingCryptoKeyIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings    []Binding `pulumi:"bindings"`
	CryptoKeyId string    `pulumi:"cryptoKeyId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag      *string `pulumi:"etag"`
	KeyRingId string  `pulumi:"keyRingId"`
	Location  *string `pulumi:"location"`
	Project   *string `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a KeyRingCryptoKeyIamPolicy resource.
type KeyRingCryptoKeyIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings    BindingArrayInput
	CryptoKeyId pulumi.StringInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag      pulumi.StringPtrInput
	KeyRingId pulumi.StringInput
	Location  pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (KeyRingCryptoKeyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingCryptoKeyIamPolicyArgs)(nil)).Elem()
}

type KeyRingCryptoKeyIamPolicyInput interface {
	pulumi.Input

	ToKeyRingCryptoKeyIamPolicyOutput() KeyRingCryptoKeyIamPolicyOutput
	ToKeyRingCryptoKeyIamPolicyOutputWithContext(ctx context.Context) KeyRingCryptoKeyIamPolicyOutput
}

func (*KeyRingCryptoKeyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRingCryptoKeyIamPolicy)(nil)).Elem()
}

func (i *KeyRingCryptoKeyIamPolicy) ToKeyRingCryptoKeyIamPolicyOutput() KeyRingCryptoKeyIamPolicyOutput {
	return i.ToKeyRingCryptoKeyIamPolicyOutputWithContext(context.Background())
}

func (i *KeyRingCryptoKeyIamPolicy) ToKeyRingCryptoKeyIamPolicyOutputWithContext(ctx context.Context) KeyRingCryptoKeyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingCryptoKeyIamPolicyOutput)
}

type KeyRingCryptoKeyIamPolicyOutput struct{ *pulumi.OutputState }

func (KeyRingCryptoKeyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRingCryptoKeyIamPolicy)(nil)).Elem()
}

func (o KeyRingCryptoKeyIamPolicyOutput) ToKeyRingCryptoKeyIamPolicyOutput() KeyRingCryptoKeyIamPolicyOutput {
	return o
}

func (o KeyRingCryptoKeyIamPolicyOutput) ToKeyRingCryptoKeyIamPolicyOutputWithContext(ctx context.Context) KeyRingCryptoKeyIamPolicyOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o KeyRingCryptoKeyIamPolicyOutput) AuditConfigs() AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *KeyRingCryptoKeyIamPolicy) AuditConfigResponseArrayOutput { return v.AuditConfigs }).(AuditConfigResponseArrayOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o KeyRingCryptoKeyIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *KeyRingCryptoKeyIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

func (o KeyRingCryptoKeyIamPolicyOutput) CryptoKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRingCryptoKeyIamPolicy) pulumi.StringOutput { return v.CryptoKeyId }).(pulumi.StringOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o KeyRingCryptoKeyIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRingCryptoKeyIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o KeyRingCryptoKeyIamPolicyOutput) KeyRingId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRingCryptoKeyIamPolicy) pulumi.StringOutput { return v.KeyRingId }).(pulumi.StringOutput)
}

func (o KeyRingCryptoKeyIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRingCryptoKeyIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o KeyRingCryptoKeyIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRingCryptoKeyIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o KeyRingCryptoKeyIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *KeyRingCryptoKeyIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRingCryptoKeyIamPolicyInput)(nil)).Elem(), &KeyRingCryptoKeyIamPolicy{})
	pulumi.RegisterOutputType(KeyRingCryptoKeyIamPolicyOutput{})
}
