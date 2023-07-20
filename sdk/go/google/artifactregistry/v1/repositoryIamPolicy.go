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

// Updates the IAM policy for a given resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type RepositoryIamPolicy struct {
	pulumi.CustomResourceState

	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag         pulumi.StringOutput `pulumi:"etag"`
	Location     pulumi.StringOutput `pulumi:"location"`
	Project      pulumi.StringOutput `pulumi:"project"`
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewRepositoryIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamPolicy(ctx *pulumi.Context,
	name string, args *RepositoryIamPolicyArgs, opts ...pulumi.ResourceOption) (*RepositoryIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"repositoryId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryIamPolicy
	err := ctx.RegisterResource("google-native:artifactregistry/v1:RepositoryIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryIamPolicy gets an existing RepositoryIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryIamPolicyState, opts ...pulumi.ResourceOption) (*RepositoryIamPolicy, error) {
	var resource RepositoryIamPolicy
	err := ctx.ReadResource("google-native:artifactregistry/v1:RepositoryIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryIamPolicy resources.
type repositoryIamPolicyState struct {
}

type RepositoryIamPolicyState struct {
}

func (RepositoryIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamPolicyState)(nil)).Elem()
}

type repositoryIamPolicyArgs struct {
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag         *string `pulumi:"etag"`
	Location     *string `pulumi:"location"`
	Project      *string `pulumi:"project"`
	RepositoryId string  `pulumi:"repositoryId"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a RepositoryIamPolicy resource.
type RepositoryIamPolicyArgs struct {
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag         pulumi.StringPtrInput
	Location     pulumi.StringPtrInput
	Project      pulumi.StringPtrInput
	RepositoryId pulumi.StringInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (RepositoryIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamPolicyArgs)(nil)).Elem()
}

type RepositoryIamPolicyInput interface {
	pulumi.Input

	ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput
	ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput
}

func (*RepositoryIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamPolicy)(nil)).Elem()
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput {
	return i.ToRepositoryIamPolicyOutputWithContext(context.Background())
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyOutput)
}

type RepositoryIamPolicyOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamPolicy)(nil)).Elem()
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput {
	return o
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput {
	return o
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o RepositoryIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o RepositoryIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RepositoryIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RepositoryIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o RepositoryIamPolicyOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o RepositoryIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamPolicyInput)(nil)).Elem(), &RepositoryIamPolicy{})
	pulumi.RegisterOutputType(RepositoryIamPolicyOutput{})
}
