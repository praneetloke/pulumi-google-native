// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the IAM Policy for a resource
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type NamespaceServiceWorkloadIamPolicy struct {
	pulumi.CustomResourceState

	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag              pulumi.StringOutput `pulumi:"etag"`
	Location          pulumi.StringOutput `pulumi:"location"`
	NamespaceId       pulumi.StringOutput `pulumi:"namespaceId"`
	Project           pulumi.StringOutput `pulumi:"project"`
	ServiceWorkloadId pulumi.StringOutput `pulumi:"serviceWorkloadId"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewNamespaceServiceWorkloadIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewNamespaceServiceWorkloadIamPolicy(ctx *pulumi.Context,
	name string, args *NamespaceServiceWorkloadIamPolicyArgs, opts ...pulumi.ResourceOption) (*NamespaceServiceWorkloadIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.ServiceWorkloadId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceWorkloadId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"namespaceId",
		"project",
		"serviceWorkloadId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NamespaceServiceWorkloadIamPolicy
	err := ctx.RegisterResource("google-native:servicedirectory/v1beta1:NamespaceServiceWorkloadIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceServiceWorkloadIamPolicy gets an existing NamespaceServiceWorkloadIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceServiceWorkloadIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceServiceWorkloadIamPolicyState, opts ...pulumi.ResourceOption) (*NamespaceServiceWorkloadIamPolicy, error) {
	var resource NamespaceServiceWorkloadIamPolicy
	err := ctx.ReadResource("google-native:servicedirectory/v1beta1:NamespaceServiceWorkloadIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceServiceWorkloadIamPolicy resources.
type namespaceServiceWorkloadIamPolicyState struct {
}

type NamespaceServiceWorkloadIamPolicyState struct {
}

func (NamespaceServiceWorkloadIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceServiceWorkloadIamPolicyState)(nil)).Elem()
}

type namespaceServiceWorkloadIamPolicyArgs struct {
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag              *string `pulumi:"etag"`
	Location          *string `pulumi:"location"`
	NamespaceId       string  `pulumi:"namespaceId"`
	Project           *string `pulumi:"project"`
	ServiceWorkloadId string  `pulumi:"serviceWorkloadId"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a NamespaceServiceWorkloadIamPolicy resource.
type NamespaceServiceWorkloadIamPolicyArgs struct {
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	NamespaceId       pulumi.StringInput
	Project           pulumi.StringPtrInput
	ServiceWorkloadId pulumi.StringInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (NamespaceServiceWorkloadIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceServiceWorkloadIamPolicyArgs)(nil)).Elem()
}

type NamespaceServiceWorkloadIamPolicyInput interface {
	pulumi.Input

	ToNamespaceServiceWorkloadIamPolicyOutput() NamespaceServiceWorkloadIamPolicyOutput
	ToNamespaceServiceWorkloadIamPolicyOutputWithContext(ctx context.Context) NamespaceServiceWorkloadIamPolicyOutput
}

func (*NamespaceServiceWorkloadIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceServiceWorkloadIamPolicy)(nil)).Elem()
}

func (i *NamespaceServiceWorkloadIamPolicy) ToNamespaceServiceWorkloadIamPolicyOutput() NamespaceServiceWorkloadIamPolicyOutput {
	return i.ToNamespaceServiceWorkloadIamPolicyOutputWithContext(context.Background())
}

func (i *NamespaceServiceWorkloadIamPolicy) ToNamespaceServiceWorkloadIamPolicyOutputWithContext(ctx context.Context) NamespaceServiceWorkloadIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceServiceWorkloadIamPolicyOutput)
}

type NamespaceServiceWorkloadIamPolicyOutput struct{ *pulumi.OutputState }

func (NamespaceServiceWorkloadIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceServiceWorkloadIamPolicy)(nil)).Elem()
}

func (o NamespaceServiceWorkloadIamPolicyOutput) ToNamespaceServiceWorkloadIamPolicyOutput() NamespaceServiceWorkloadIamPolicyOutput {
	return o
}

func (o NamespaceServiceWorkloadIamPolicyOutput) ToNamespaceServiceWorkloadIamPolicyOutputWithContext(ctx context.Context) NamespaceServiceWorkloadIamPolicyOutput {
	return o
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o NamespaceServiceWorkloadIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *NamespaceServiceWorkloadIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o NamespaceServiceWorkloadIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceServiceWorkloadIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NamespaceServiceWorkloadIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceServiceWorkloadIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NamespaceServiceWorkloadIamPolicyOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceServiceWorkloadIamPolicy) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

func (o NamespaceServiceWorkloadIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceServiceWorkloadIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o NamespaceServiceWorkloadIamPolicyOutput) ServiceWorkloadId() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceServiceWorkloadIamPolicy) pulumi.StringOutput { return v.ServiceWorkloadId }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o NamespaceServiceWorkloadIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *NamespaceServiceWorkloadIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceServiceWorkloadIamPolicyInput)(nil)).Elem(), &NamespaceServiceWorkloadIamPolicy{})
	pulumi.RegisterOutputType(NamespaceServiceWorkloadIamPolicyOutput{})
}
