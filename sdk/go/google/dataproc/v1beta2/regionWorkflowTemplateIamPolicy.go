// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED errors.
type RegionWorkflowTemplateIamPolicy struct {
	pulumi.CustomResourceState

	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewRegionWorkflowTemplateIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRegionWorkflowTemplateIamPolicy(ctx *pulumi.Context,
	name string, args *RegionWorkflowTemplateIamPolicyArgs, opts ...pulumi.ResourceOption) (*RegionWorkflowTemplateIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.RegionId == nil {
		return nil, errors.New("invalid value for required argument 'RegionId'")
	}
	if args.WorkflowTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'WorkflowTemplateId'")
	}
	var resource RegionWorkflowTemplateIamPolicy
	err := ctx.RegisterResource("google-native:dataproc/v1beta2:RegionWorkflowTemplateIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionWorkflowTemplateIamPolicy gets an existing RegionWorkflowTemplateIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionWorkflowTemplateIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionWorkflowTemplateIamPolicyState, opts ...pulumi.ResourceOption) (*RegionWorkflowTemplateIamPolicy, error) {
	var resource RegionWorkflowTemplateIamPolicy
	err := ctx.ReadResource("google-native:dataproc/v1beta2:RegionWorkflowTemplateIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionWorkflowTemplateIamPolicy resources.
type regionWorkflowTemplateIamPolicyState struct {
	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings []BindingResponse `pulumi:"bindings"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag *string `pulumi:"etag"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

type RegionWorkflowTemplateIamPolicyState struct {
	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings BindingResponseArrayInput
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag pulumi.StringPtrInput
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (RegionWorkflowTemplateIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionWorkflowTemplateIamPolicyState)(nil)).Elem()
}

type regionWorkflowTemplateIamPolicyArgs struct {
	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings []Binding `pulumi:"bindings"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag     *string `pulumi:"etag"`
	Project  string  `pulumi:"project"`
	RegionId string  `pulumi:"regionId"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version            *int   `pulumi:"version"`
	WorkflowTemplateId string `pulumi:"workflowTemplateId"`
}

// The set of arguments for constructing a RegionWorkflowTemplateIamPolicy resource.
type RegionWorkflowTemplateIamPolicyArgs struct {
	// Associates a list of members to a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one member.
	Bindings BindingArrayInput
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag     pulumi.StringPtrInput
	Project  pulumi.StringInput
	RegionId pulumi.StringInput
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version            pulumi.IntPtrInput
	WorkflowTemplateId pulumi.StringInput
}

func (RegionWorkflowTemplateIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionWorkflowTemplateIamPolicyArgs)(nil)).Elem()
}

type RegionWorkflowTemplateIamPolicyInput interface {
	pulumi.Input

	ToRegionWorkflowTemplateIamPolicyOutput() RegionWorkflowTemplateIamPolicyOutput
	ToRegionWorkflowTemplateIamPolicyOutputWithContext(ctx context.Context) RegionWorkflowTemplateIamPolicyOutput
}

func (*RegionWorkflowTemplateIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionWorkflowTemplateIamPolicy)(nil))
}

func (i *RegionWorkflowTemplateIamPolicy) ToRegionWorkflowTemplateIamPolicyOutput() RegionWorkflowTemplateIamPolicyOutput {
	return i.ToRegionWorkflowTemplateIamPolicyOutputWithContext(context.Background())
}

func (i *RegionWorkflowTemplateIamPolicy) ToRegionWorkflowTemplateIamPolicyOutputWithContext(ctx context.Context) RegionWorkflowTemplateIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionWorkflowTemplateIamPolicyOutput)
}

type RegionWorkflowTemplateIamPolicyOutput struct {
	*pulumi.OutputState
}

func (RegionWorkflowTemplateIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionWorkflowTemplateIamPolicy)(nil))
}

func (o RegionWorkflowTemplateIamPolicyOutput) ToRegionWorkflowTemplateIamPolicyOutput() RegionWorkflowTemplateIamPolicyOutput {
	return o
}

func (o RegionWorkflowTemplateIamPolicyOutput) ToRegionWorkflowTemplateIamPolicyOutputWithContext(ctx context.Context) RegionWorkflowTemplateIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionWorkflowTemplateIamPolicyOutput{})
}
