// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the IAM Policy for a resource (namespace or service only).
type NamespaceServiceIamPolicy struct {
	pulumi.CustomResourceState

	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewNamespaceServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewNamespaceServiceIamPolicy(ctx *pulumi.Context,
	name string, args *NamespaceServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*NamespaceServiceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	var resource NamespaceServiceIamPolicy
	err := ctx.RegisterResource("google-native:servicedirectory/v1beta1:NamespaceServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceServiceIamPolicy gets an existing NamespaceServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceServiceIamPolicyState, opts ...pulumi.ResourceOption) (*NamespaceServiceIamPolicy, error) {
	var resource NamespaceServiceIamPolicy
	err := ctx.ReadResource("google-native:servicedirectory/v1beta1:NamespaceServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceServiceIamPolicy resources.
type namespaceServiceIamPolicyState struct {
}

type NamespaceServiceIamPolicyState struct {
}

func (NamespaceServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceServiceIamPolicyState)(nil)).Elem()
}

type namespaceServiceIamPolicyArgs struct {
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag        *string `pulumi:"etag"`
	Location    *string `pulumi:"location"`
	NamespaceId string  `pulumi:"namespaceId"`
	Project     *string `pulumi:"project"`
	ServiceId   string  `pulumi:"serviceId"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a NamespaceServiceIamPolicy resource.
type NamespaceServiceIamPolicyArgs struct {
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag        pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	NamespaceId pulumi.StringInput
	Project     pulumi.StringPtrInput
	ServiceId   pulumi.StringInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (NamespaceServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceServiceIamPolicyArgs)(nil)).Elem()
}

type NamespaceServiceIamPolicyInput interface {
	pulumi.Input

	ToNamespaceServiceIamPolicyOutput() NamespaceServiceIamPolicyOutput
	ToNamespaceServiceIamPolicyOutputWithContext(ctx context.Context) NamespaceServiceIamPolicyOutput
}

func (*NamespaceServiceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceServiceIamPolicy)(nil))
}

func (i *NamespaceServiceIamPolicy) ToNamespaceServiceIamPolicyOutput() NamespaceServiceIamPolicyOutput {
	return i.ToNamespaceServiceIamPolicyOutputWithContext(context.Background())
}

func (i *NamespaceServiceIamPolicy) ToNamespaceServiceIamPolicyOutputWithContext(ctx context.Context) NamespaceServiceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceServiceIamPolicyOutput)
}

type NamespaceServiceIamPolicyOutput struct {
	*pulumi.OutputState
}

func (NamespaceServiceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceServiceIamPolicy)(nil))
}

func (o NamespaceServiceIamPolicyOutput) ToNamespaceServiceIamPolicyOutput() NamespaceServiceIamPolicyOutput {
	return o
}

func (o NamespaceServiceIamPolicyOutput) ToNamespaceServiceIamPolicyOutputWithContext(ctx context.Context) NamespaceServiceIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceServiceIamPolicyOutput{})
}
