// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type CaPoolIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewCaPoolIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewCaPoolIamPolicy(ctx *pulumi.Context,
	name string, args *CaPoolIamPolicyArgs, opts ...pulumi.ResourceOption) (*CaPoolIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPoolId == nil {
		return nil, errors.New("invalid value for required argument 'CaPoolId'")
	}
	var resource CaPoolIamPolicy
	err := ctx.RegisterResource("google-native:privateca/v1:CaPoolIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaPoolIamPolicy gets an existing CaPoolIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaPoolIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaPoolIamPolicyState, opts ...pulumi.ResourceOption) (*CaPoolIamPolicy, error) {
	var resource CaPoolIamPolicy
	err := ctx.ReadResource("google-native:privateca/v1:CaPoolIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaPoolIamPolicy resources.
type caPoolIamPolicyState struct {
}

type CaPoolIamPolicyState struct {
}

func (CaPoolIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamPolicyState)(nil)).Elem()
}

type caPoolIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	CaPoolId string    `pulumi:"caPoolId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a CaPoolIamPolicy resource.
type CaPoolIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingArrayInput
	CaPoolId pulumi.StringInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (CaPoolIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamPolicyArgs)(nil)).Elem()
}

type CaPoolIamPolicyInput interface {
	pulumi.Input

	ToCaPoolIamPolicyOutput() CaPoolIamPolicyOutput
	ToCaPoolIamPolicyOutputWithContext(ctx context.Context) CaPoolIamPolicyOutput
}

func (*CaPoolIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*CaPoolIamPolicy)(nil))
}

func (i *CaPoolIamPolicy) ToCaPoolIamPolicyOutput() CaPoolIamPolicyOutput {
	return i.ToCaPoolIamPolicyOutputWithContext(context.Background())
}

func (i *CaPoolIamPolicy) ToCaPoolIamPolicyOutputWithContext(ctx context.Context) CaPoolIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamPolicyOutput)
}

type CaPoolIamPolicyOutput struct {
	*pulumi.OutputState
}

func (CaPoolIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaPoolIamPolicy)(nil))
}

func (o CaPoolIamPolicyOutput) ToCaPoolIamPolicyOutput() CaPoolIamPolicyOutput {
	return o
}

func (o CaPoolIamPolicyOutput) ToCaPoolIamPolicyOutputWithContext(ctx context.Context) CaPoolIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CaPoolIamPolicyOutput{})
}
