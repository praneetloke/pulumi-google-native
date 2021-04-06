// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.
type SnapshotIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringOutput `pulumi:"etag"`
	IamOwned pulumi.BoolOutput   `pulumi:"iamOwned"`
	// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
	Rules RuleResponseArrayOutput `pulumi:"rules"`
	// Specifies the format of the policy.
	//
	// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
	//
	// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
	//
	// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	//
	// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
	//
	// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewSnapshotIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewSnapshotIamPolicy(ctx *pulumi.Context,
	name string, args *SnapshotIamPolicyArgs, opts ...pulumi.ResourceOption) (*SnapshotIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	var resource SnapshotIamPolicy
	err := ctx.RegisterResource("google-cloud:compute/alpha:SnapshotIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotIamPolicy gets an existing SnapshotIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotIamPolicyState, opts ...pulumi.ResourceOption) (*SnapshotIamPolicy, error) {
	var resource SnapshotIamPolicy
	err := ctx.ReadResource("google-cloud:compute/alpha:SnapshotIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotIamPolicy resources.
type snapshotIamPolicyState struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfigResponse `pulumi:"auditConfigs"`
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings []BindingResponse `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	IamOwned *bool   `pulumi:"iamOwned"`
	// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
	Rules []RuleResponse `pulumi:"rules"`
	// Specifies the format of the policy.
	//
	// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
	//
	// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
	//
	// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	//
	// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
	//
	// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

type SnapshotIamPolicyState struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayInput
	// Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
	Bindings BindingResponseArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	IamOwned pulumi.BoolPtrInput
	// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
	Rules RuleResponseArrayInput
	// Specifies the format of the policy.
	//
	// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
	//
	// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
	//
	// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	//
	// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
	//
	// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (SnapshotIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotIamPolicyState)(nil)).Elem()
}

type snapshotIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify bindings.
	Bindings []Binding `pulumi:"bindings"`
	// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify the etag.
	Etag     *string `pulumi:"etag"`
	IamOwned *bool   `pulumi:"iamOwned"`
	Project  string  `pulumi:"project"`
	Resource string  `pulumi:"resource"`
	// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
	Rules []Rule `pulumi:"rules"`
	// Specifies the format of the policy.
	//
	// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
	//
	// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
	//
	// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	//
	// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
	//
	// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a SnapshotIamPolicy resource.
type SnapshotIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify bindings.
	Bindings BindingArrayInput
	// Flatten Policy to create a backward compatible wire-format. Deprecated. Use 'policy' to specify the etag.
	Etag     pulumi.StringPtrInput
	IamOwned pulumi.BoolPtrInput
	Project  pulumi.StringInput
	Resource pulumi.StringInput
	// If more than one rule is specified, the rules are applied in the following manner: - All matching LOG rules are always applied. - If any DENY/DENY_WITH_LOG rule matches, permission is denied. Logging will be applied if one or more matching rule requires logging. - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is granted. Logging will be applied if one or more matching rule requires logging. - Otherwise, if no rule applies, permission is denied.
	Rules RuleArrayInput
	// Specifies the format of the policy.
	//
	// Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected.
	//
	// Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations:
	//
	// * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	//
	// If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.
	//
	// To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (SnapshotIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotIamPolicyArgs)(nil)).Elem()
}

type SnapshotIamPolicyInput interface {
	pulumi.Input

	ToSnapshotIamPolicyOutput() SnapshotIamPolicyOutput
	ToSnapshotIamPolicyOutputWithContext(ctx context.Context) SnapshotIamPolicyOutput
}

func (*SnapshotIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotIamPolicy)(nil))
}

func (i *SnapshotIamPolicy) ToSnapshotIamPolicyOutput() SnapshotIamPolicyOutput {
	return i.ToSnapshotIamPolicyOutputWithContext(context.Background())
}

func (i *SnapshotIamPolicy) ToSnapshotIamPolicyOutputWithContext(ctx context.Context) SnapshotIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotIamPolicyOutput)
}

type SnapshotIamPolicyOutput struct {
	*pulumi.OutputState
}

func (SnapshotIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotIamPolicy)(nil))
}

func (o SnapshotIamPolicyOutput) ToSnapshotIamPolicyOutput() SnapshotIamPolicyOutput {
	return o
}

func (o SnapshotIamPolicyOutput) ToSnapshotIamPolicyOutputWithContext(ctx context.Context) SnapshotIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SnapshotIamPolicyOutput{})
}
