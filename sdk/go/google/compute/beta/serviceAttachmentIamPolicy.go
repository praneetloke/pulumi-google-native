// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.
type ServiceAttachmentIamPolicy struct {
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

// NewServiceAttachmentIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceAttachmentIamPolicy(ctx *pulumi.Context,
	name string, args *ServiceAttachmentIamPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceAttachmentIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	var resource ServiceAttachmentIamPolicy
	err := ctx.RegisterResource("google-cloud:compute/beta:ServiceAttachmentIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAttachmentIamPolicy gets an existing ServiceAttachmentIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAttachmentIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAttachmentIamPolicyState, opts ...pulumi.ResourceOption) (*ServiceAttachmentIamPolicy, error) {
	var resource ServiceAttachmentIamPolicy
	err := ctx.ReadResource("google-cloud:compute/beta:ServiceAttachmentIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAttachmentIamPolicy resources.
type serviceAttachmentIamPolicyState struct {
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

type ServiceAttachmentIamPolicyState struct {
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

func (ServiceAttachmentIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAttachmentIamPolicyState)(nil)).Elem()
}

type serviceAttachmentIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Flatten Policy to create a backwacd compatible wire-format. Deprecated. Use 'policy' to specify bindings.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	IamOwned *bool   `pulumi:"iamOwned"`
	Project  string  `pulumi:"project"`
	Region   string  `pulumi:"region"`
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

// The set of arguments for constructing a ServiceAttachmentIamPolicy resource.
type ServiceAttachmentIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Flatten Policy to create a backwacd compatible wire-format. Deprecated. Use 'policy' to specify bindings.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy.
	//
	// **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	IamOwned pulumi.BoolPtrInput
	Project  pulumi.StringInput
	Region   pulumi.StringInput
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

func (ServiceAttachmentIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAttachmentIamPolicyArgs)(nil)).Elem()
}

type ServiceAttachmentIamPolicyInput interface {
	pulumi.Input

	ToServiceAttachmentIamPolicyOutput() ServiceAttachmentIamPolicyOutput
	ToServiceAttachmentIamPolicyOutputWithContext(ctx context.Context) ServiceAttachmentIamPolicyOutput
}

func (*ServiceAttachmentIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAttachmentIamPolicy)(nil))
}

func (i *ServiceAttachmentIamPolicy) ToServiceAttachmentIamPolicyOutput() ServiceAttachmentIamPolicyOutput {
	return i.ToServiceAttachmentIamPolicyOutputWithContext(context.Background())
}

func (i *ServiceAttachmentIamPolicy) ToServiceAttachmentIamPolicyOutputWithContext(ctx context.Context) ServiceAttachmentIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAttachmentIamPolicyOutput)
}

type ServiceAttachmentIamPolicyOutput struct {
	*pulumi.OutputState
}

func (ServiceAttachmentIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAttachmentIamPolicy)(nil))
}

func (o ServiceAttachmentIamPolicyOutput) ToServiceAttachmentIamPolicyOutput() ServiceAttachmentIamPolicyOutput {
	return o
}

func (o ServiceAttachmentIamPolicyOutput) ToServiceAttachmentIamPolicyOutputWithContext(ctx context.Context) ServiceAttachmentIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceAttachmentIamPolicyOutput{})
}
