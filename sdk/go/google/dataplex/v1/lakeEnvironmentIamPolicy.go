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

// Sets the access control policy on the specified resource. Replaces any existing policy.Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED errors.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type LakeEnvironmentIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs GoogleIamV1AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
	Bindings      GoogleIamV1BindingResponseArrayOutput `pulumi:"bindings"`
	EnvironmentId pulumi.StringOutput                   `pulumi:"environmentId"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag     pulumi.StringOutput `pulumi:"etag"`
	LakeId   pulumi.StringOutput `pulumi:"lakeId"`
	Location pulumi.StringOutput `pulumi:"location"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewLakeEnvironmentIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewLakeEnvironmentIamPolicy(ctx *pulumi.Context,
	name string, args *LakeEnvironmentIamPolicyArgs, opts ...pulumi.ResourceOption) (*LakeEnvironmentIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.LakeId == nil {
		return nil, errors.New("invalid value for required argument 'LakeId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"environmentId",
		"lakeId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LakeEnvironmentIamPolicy
	err := ctx.RegisterResource("google-native:dataplex/v1:LakeEnvironmentIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLakeEnvironmentIamPolicy gets an existing LakeEnvironmentIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLakeEnvironmentIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LakeEnvironmentIamPolicyState, opts ...pulumi.ResourceOption) (*LakeEnvironmentIamPolicy, error) {
	var resource LakeEnvironmentIamPolicy
	err := ctx.ReadResource("google-native:dataplex/v1:LakeEnvironmentIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LakeEnvironmentIamPolicy resources.
type lakeEnvironmentIamPolicyState struct {
}

type LakeEnvironmentIamPolicyState struct {
}

func (LakeEnvironmentIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lakeEnvironmentIamPolicyState)(nil)).Elem()
}

type lakeEnvironmentIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []GoogleIamV1AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
	Bindings      []GoogleIamV1Binding `pulumi:"bindings"`
	EnvironmentId string               `pulumi:"environmentId"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag     *string `pulumi:"etag"`
	LakeId   string  `pulumi:"lakeId"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used:paths: "bindings, etag"
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a LakeEnvironmentIamPolicy resource.
type LakeEnvironmentIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs GoogleIamV1AuditConfigArrayInput
	// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
	Bindings      GoogleIamV1BindingArrayInput
	EnvironmentId pulumi.StringInput
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag     pulumi.StringPtrInput
	LakeId   pulumi.StringInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used:paths: "bindings, etag"
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (LakeEnvironmentIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lakeEnvironmentIamPolicyArgs)(nil)).Elem()
}

type LakeEnvironmentIamPolicyInput interface {
	pulumi.Input

	ToLakeEnvironmentIamPolicyOutput() LakeEnvironmentIamPolicyOutput
	ToLakeEnvironmentIamPolicyOutputWithContext(ctx context.Context) LakeEnvironmentIamPolicyOutput
}

func (*LakeEnvironmentIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**LakeEnvironmentIamPolicy)(nil)).Elem()
}

func (i *LakeEnvironmentIamPolicy) ToLakeEnvironmentIamPolicyOutput() LakeEnvironmentIamPolicyOutput {
	return i.ToLakeEnvironmentIamPolicyOutputWithContext(context.Background())
}

func (i *LakeEnvironmentIamPolicy) ToLakeEnvironmentIamPolicyOutputWithContext(ctx context.Context) LakeEnvironmentIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LakeEnvironmentIamPolicyOutput)
}

type LakeEnvironmentIamPolicyOutput struct{ *pulumi.OutputState }

func (LakeEnvironmentIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LakeEnvironmentIamPolicy)(nil)).Elem()
}

func (o LakeEnvironmentIamPolicyOutput) ToLakeEnvironmentIamPolicyOutput() LakeEnvironmentIamPolicyOutput {
	return o
}

func (o LakeEnvironmentIamPolicyOutput) ToLakeEnvironmentIamPolicyOutputWithContext(ctx context.Context) LakeEnvironmentIamPolicyOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o LakeEnvironmentIamPolicyOutput) AuditConfigs() GoogleIamV1AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *LakeEnvironmentIamPolicy) GoogleIamV1AuditConfigResponseArrayOutput { return v.AuditConfigs }).(GoogleIamV1AuditConfigResponseArrayOutput)
}

// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
func (o LakeEnvironmentIamPolicyOutput) Bindings() GoogleIamV1BindingResponseArrayOutput {
	return o.ApplyT(func(v *LakeEnvironmentIamPolicy) GoogleIamV1BindingResponseArrayOutput { return v.Bindings }).(GoogleIamV1BindingResponseArrayOutput)
}

func (o LakeEnvironmentIamPolicyOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeEnvironmentIamPolicy) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
func (o LakeEnvironmentIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeEnvironmentIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o LakeEnvironmentIamPolicyOutput) LakeId() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeEnvironmentIamPolicy) pulumi.StringOutput { return v.LakeId }).(pulumi.StringOutput)
}

func (o LakeEnvironmentIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeEnvironmentIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o LakeEnvironmentIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeEnvironmentIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
func (o LakeEnvironmentIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *LakeEnvironmentIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LakeEnvironmentIamPolicyInput)(nil)).Elem(), &LakeEnvironmentIamPolicy{})
	pulumi.RegisterOutputType(LakeEnvironmentIamPolicyOutput{})
}
