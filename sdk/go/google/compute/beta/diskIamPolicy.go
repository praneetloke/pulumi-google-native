// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.
type DiskIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// This is deprecated and has no effect. Do not use.
	IamOwned pulumi.BoolOutput `pulumi:"iamOwned"`
	// This is deprecated and has no effect. Do not use.
	Rules RuleResponseArrayOutput `pulumi:"rules"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewDiskIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDiskIamPolicy(ctx *pulumi.Context,
	name string, args *DiskIamPolicyArgs, opts ...pulumi.ResourceOption) (*DiskIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	var resource DiskIamPolicy
	err := ctx.RegisterResource("google-native:compute/beta:DiskIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskIamPolicy gets an existing DiskIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskIamPolicyState, opts ...pulumi.ResourceOption) (*DiskIamPolicy, error) {
	var resource DiskIamPolicy
	err := ctx.ReadResource("google-native:compute/beta:DiskIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskIamPolicy resources.
type diskIamPolicyState struct {
}

type DiskIamPolicyState struct {
}

func (DiskIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskIamPolicyState)(nil)).Elem()
}

type diskIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag *string `pulumi:"etag"`
	// This is deprecated and has no effect. Do not use.
	IamOwned *bool   `pulumi:"iamOwned"`
	Project  *string `pulumi:"project"`
	Resource string  `pulumi:"resource"`
	// This is deprecated and has no effect. Do not use.
	Rules []Rule `pulumi:"rules"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int    `pulumi:"version"`
	Zone    *string `pulumi:"zone"`
}

// The set of arguments for constructing a DiskIamPolicy resource.
type DiskIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag pulumi.StringPtrInput
	// This is deprecated and has no effect. Do not use.
	IamOwned pulumi.BoolPtrInput
	Project  pulumi.StringPtrInput
	Resource pulumi.StringInput
	// This is deprecated and has no effect. Do not use.
	Rules RuleArrayInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
	Zone    pulumi.StringPtrInput
}

func (DiskIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskIamPolicyArgs)(nil)).Elem()
}

type DiskIamPolicyInput interface {
	pulumi.Input

	ToDiskIamPolicyOutput() DiskIamPolicyOutput
	ToDiskIamPolicyOutputWithContext(ctx context.Context) DiskIamPolicyOutput
}

func (*DiskIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskIamPolicy)(nil))
}

func (i *DiskIamPolicy) ToDiskIamPolicyOutput() DiskIamPolicyOutput {
	return i.ToDiskIamPolicyOutputWithContext(context.Background())
}

func (i *DiskIamPolicy) ToDiskIamPolicyOutputWithContext(ctx context.Context) DiskIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskIamPolicyOutput)
}

type DiskIamPolicyOutput struct {
	*pulumi.OutputState
}

func (DiskIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskIamPolicy)(nil))
}

func (o DiskIamPolicyOutput) ToDiskIamPolicyOutput() DiskIamPolicyOutput {
	return o
}

func (o DiskIamPolicyOutput) ToDiskIamPolicyOutputWithContext(ctx context.Context) DiskIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DiskIamPolicyOutput{})
}
