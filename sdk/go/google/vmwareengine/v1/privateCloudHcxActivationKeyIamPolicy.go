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
type PrivateCloudHcxActivationKeyIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag               pulumi.StringOutput `pulumi:"etag"`
	HcxActivationKeyId pulumi.StringOutput `pulumi:"hcxActivationKeyId"`
	Location           pulumi.StringOutput `pulumi:"location"`
	PrivateCloudId     pulumi.StringOutput `pulumi:"privateCloudId"`
	Project            pulumi.StringOutput `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewPrivateCloudHcxActivationKeyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewPrivateCloudHcxActivationKeyIamPolicy(ctx *pulumi.Context,
	name string, args *PrivateCloudHcxActivationKeyIamPolicyArgs, opts ...pulumi.ResourceOption) (*PrivateCloudHcxActivationKeyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HcxActivationKeyId == nil {
		return nil, errors.New("invalid value for required argument 'HcxActivationKeyId'")
	}
	if args.PrivateCloudId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"hcxActivationKeyId",
		"location",
		"privateCloudId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateCloudHcxActivationKeyIamPolicy
	err := ctx.RegisterResource("google-native:vmwareengine/v1:PrivateCloudHcxActivationKeyIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateCloudHcxActivationKeyIamPolicy gets an existing PrivateCloudHcxActivationKeyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateCloudHcxActivationKeyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateCloudHcxActivationKeyIamPolicyState, opts ...pulumi.ResourceOption) (*PrivateCloudHcxActivationKeyIamPolicy, error) {
	var resource PrivateCloudHcxActivationKeyIamPolicy
	err := ctx.ReadResource("google-native:vmwareengine/v1:PrivateCloudHcxActivationKeyIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateCloudHcxActivationKeyIamPolicy resources.
type privateCloudHcxActivationKeyIamPolicyState struct {
}

type PrivateCloudHcxActivationKeyIamPolicyState struct {
}

func (PrivateCloudHcxActivationKeyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateCloudHcxActivationKeyIamPolicyState)(nil)).Elem()
}

type privateCloudHcxActivationKeyIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag               *string `pulumi:"etag"`
	HcxActivationKeyId string  `pulumi:"hcxActivationKeyId"`
	Location           *string `pulumi:"location"`
	PrivateCloudId     string  `pulumi:"privateCloudId"`
	Project            *string `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a PrivateCloudHcxActivationKeyIamPolicy resource.
type PrivateCloudHcxActivationKeyIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag               pulumi.StringPtrInput
	HcxActivationKeyId pulumi.StringInput
	Location           pulumi.StringPtrInput
	PrivateCloudId     pulumi.StringInput
	Project            pulumi.StringPtrInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (PrivateCloudHcxActivationKeyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateCloudHcxActivationKeyIamPolicyArgs)(nil)).Elem()
}

type PrivateCloudHcxActivationKeyIamPolicyInput interface {
	pulumi.Input

	ToPrivateCloudHcxActivationKeyIamPolicyOutput() PrivateCloudHcxActivationKeyIamPolicyOutput
	ToPrivateCloudHcxActivationKeyIamPolicyOutputWithContext(ctx context.Context) PrivateCloudHcxActivationKeyIamPolicyOutput
}

func (*PrivateCloudHcxActivationKeyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudHcxActivationKeyIamPolicy)(nil)).Elem()
}

func (i *PrivateCloudHcxActivationKeyIamPolicy) ToPrivateCloudHcxActivationKeyIamPolicyOutput() PrivateCloudHcxActivationKeyIamPolicyOutput {
	return i.ToPrivateCloudHcxActivationKeyIamPolicyOutputWithContext(context.Background())
}

func (i *PrivateCloudHcxActivationKeyIamPolicy) ToPrivateCloudHcxActivationKeyIamPolicyOutputWithContext(ctx context.Context) PrivateCloudHcxActivationKeyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudHcxActivationKeyIamPolicyOutput)
}

type PrivateCloudHcxActivationKeyIamPolicyOutput struct{ *pulumi.OutputState }

func (PrivateCloudHcxActivationKeyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudHcxActivationKeyIamPolicy)(nil)).Elem()
}

func (o PrivateCloudHcxActivationKeyIamPolicyOutput) ToPrivateCloudHcxActivationKeyIamPolicyOutput() PrivateCloudHcxActivationKeyIamPolicyOutput {
	return o
}

func (o PrivateCloudHcxActivationKeyIamPolicyOutput) ToPrivateCloudHcxActivationKeyIamPolicyOutputWithContext(ctx context.Context) PrivateCloudHcxActivationKeyIamPolicyOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o PrivateCloudHcxActivationKeyIamPolicyOutput) AuditConfigs() AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *PrivateCloudHcxActivationKeyIamPolicy) AuditConfigResponseArrayOutput { return v.AuditConfigs }).(AuditConfigResponseArrayOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o PrivateCloudHcxActivationKeyIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *PrivateCloudHcxActivationKeyIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o PrivateCloudHcxActivationKeyIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloudHcxActivationKeyIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PrivateCloudHcxActivationKeyIamPolicyOutput) HcxActivationKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloudHcxActivationKeyIamPolicy) pulumi.StringOutput { return v.HcxActivationKeyId }).(pulumi.StringOutput)
}

func (o PrivateCloudHcxActivationKeyIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloudHcxActivationKeyIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateCloudHcxActivationKeyIamPolicyOutput) PrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloudHcxActivationKeyIamPolicy) pulumi.StringOutput { return v.PrivateCloudId }).(pulumi.StringOutput)
}

func (o PrivateCloudHcxActivationKeyIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloudHcxActivationKeyIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o PrivateCloudHcxActivationKeyIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivateCloudHcxActivationKeyIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateCloudHcxActivationKeyIamPolicyInput)(nil)).Elem(), &PrivateCloudHcxActivationKeyIamPolicy{})
	pulumi.RegisterOutputType(PrivateCloudHcxActivationKeyIamPolicyOutput{})
}