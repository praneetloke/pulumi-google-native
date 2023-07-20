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
type ApiDeploymentIamPolicy struct {
	pulumi.CustomResourceState

	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings     BindingResponseArrayOutput `pulumi:"bindings"`
	DeploymentId pulumi.StringOutput        `pulumi:"deploymentId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringOutput `pulumi:"etag"`
	Location pulumi.StringOutput `pulumi:"location"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewApiDeploymentIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiDeploymentIamPolicy(ctx *pulumi.Context,
	name string, args *ApiDeploymentIamPolicyArgs, opts ...pulumi.ResourceOption) (*ApiDeploymentIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.DeploymentId == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"apiId",
		"deploymentId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiDeploymentIamPolicy
	err := ctx.RegisterResource("google-native:apigeeregistry/v1:ApiDeploymentIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiDeploymentIamPolicy gets an existing ApiDeploymentIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiDeploymentIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDeploymentIamPolicyState, opts ...pulumi.ResourceOption) (*ApiDeploymentIamPolicy, error) {
	var resource ApiDeploymentIamPolicy
	err := ctx.ReadResource("google-native:apigeeregistry/v1:ApiDeploymentIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiDeploymentIamPolicy resources.
type apiDeploymentIamPolicyState struct {
}

type ApiDeploymentIamPolicyState struct {
}

func (ApiDeploymentIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDeploymentIamPolicyState)(nil)).Elem()
}

type apiDeploymentIamPolicyArgs struct {
	ApiId string `pulumi:"apiId"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings     []Binding `pulumi:"bindings"`
	DeploymentId string    `pulumi:"deploymentId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a ApiDeploymentIamPolicy resource.
type ApiDeploymentIamPolicyArgs struct {
	ApiId pulumi.StringInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings     BindingArrayInput
	DeploymentId pulumi.StringInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (ApiDeploymentIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDeploymentIamPolicyArgs)(nil)).Elem()
}

type ApiDeploymentIamPolicyInput interface {
	pulumi.Input

	ToApiDeploymentIamPolicyOutput() ApiDeploymentIamPolicyOutput
	ToApiDeploymentIamPolicyOutputWithContext(ctx context.Context) ApiDeploymentIamPolicyOutput
}

func (*ApiDeploymentIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDeploymentIamPolicy)(nil)).Elem()
}

func (i *ApiDeploymentIamPolicy) ToApiDeploymentIamPolicyOutput() ApiDeploymentIamPolicyOutput {
	return i.ToApiDeploymentIamPolicyOutputWithContext(context.Background())
}

func (i *ApiDeploymentIamPolicy) ToApiDeploymentIamPolicyOutputWithContext(ctx context.Context) ApiDeploymentIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDeploymentIamPolicyOutput)
}

type ApiDeploymentIamPolicyOutput struct{ *pulumi.OutputState }

func (ApiDeploymentIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDeploymentIamPolicy)(nil)).Elem()
}

func (o ApiDeploymentIamPolicyOutput) ToApiDeploymentIamPolicyOutput() ApiDeploymentIamPolicyOutput {
	return o
}

func (o ApiDeploymentIamPolicyOutput) ToApiDeploymentIamPolicyOutputWithContext(ctx context.Context) ApiDeploymentIamPolicyOutput {
	return o
}

func (o ApiDeploymentIamPolicyOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDeploymentIamPolicy) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o ApiDeploymentIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *ApiDeploymentIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

func (o ApiDeploymentIamPolicyOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDeploymentIamPolicy) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o ApiDeploymentIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDeploymentIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApiDeploymentIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDeploymentIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ApiDeploymentIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDeploymentIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o ApiDeploymentIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *ApiDeploymentIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiDeploymentIamPolicyInput)(nil)).Elem(), &ApiDeploymentIamPolicy{})
	pulumi.RegisterOutputType(ApiDeploymentIamPolicyOutput{})
}
