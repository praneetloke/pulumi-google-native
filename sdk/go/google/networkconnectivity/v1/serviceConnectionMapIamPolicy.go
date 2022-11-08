// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type ServiceConnectionMapIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag                   pulumi.StringOutput `pulumi:"etag"`
	Location               pulumi.StringOutput `pulumi:"location"`
	Project                pulumi.StringOutput `pulumi:"project"`
	ServiceConnectionMapId pulumi.StringOutput `pulumi:"serviceConnectionMapId"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewServiceConnectionMapIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceConnectionMapIamPolicy(ctx *pulumi.Context,
	name string, args *ServiceConnectionMapIamPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceConnectionMapIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceConnectionMapId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceConnectionMapId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"serviceConnectionMapId",
	})
	opts = append(opts, replaceOnChanges)
	var resource ServiceConnectionMapIamPolicy
	err := ctx.RegisterResource("google-native:networkconnectivity/v1:ServiceConnectionMapIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceConnectionMapIamPolicy gets an existing ServiceConnectionMapIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceConnectionMapIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceConnectionMapIamPolicyState, opts ...pulumi.ResourceOption) (*ServiceConnectionMapIamPolicy, error) {
	var resource ServiceConnectionMapIamPolicy
	err := ctx.ReadResource("google-native:networkconnectivity/v1:ServiceConnectionMapIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceConnectionMapIamPolicy resources.
type serviceConnectionMapIamPolicyState struct {
}

type ServiceConnectionMapIamPolicyState struct {
}

func (ServiceConnectionMapIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceConnectionMapIamPolicyState)(nil)).Elem()
}

type serviceConnectionMapIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag                   *string `pulumi:"etag"`
	Location               *string `pulumi:"location"`
	Project                *string `pulumi:"project"`
	ServiceConnectionMapId string  `pulumi:"serviceConnectionMapId"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a ServiceConnectionMapIamPolicy resource.
type ServiceConnectionMapIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag                   pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	Project                pulumi.StringPtrInput
	ServiceConnectionMapId pulumi.StringInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (ServiceConnectionMapIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceConnectionMapIamPolicyArgs)(nil)).Elem()
}

type ServiceConnectionMapIamPolicyInput interface {
	pulumi.Input

	ToServiceConnectionMapIamPolicyOutput() ServiceConnectionMapIamPolicyOutput
	ToServiceConnectionMapIamPolicyOutputWithContext(ctx context.Context) ServiceConnectionMapIamPolicyOutput
}

func (*ServiceConnectionMapIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceConnectionMapIamPolicy)(nil)).Elem()
}

func (i *ServiceConnectionMapIamPolicy) ToServiceConnectionMapIamPolicyOutput() ServiceConnectionMapIamPolicyOutput {
	return i.ToServiceConnectionMapIamPolicyOutputWithContext(context.Background())
}

func (i *ServiceConnectionMapIamPolicy) ToServiceConnectionMapIamPolicyOutputWithContext(ctx context.Context) ServiceConnectionMapIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceConnectionMapIamPolicyOutput)
}

type ServiceConnectionMapIamPolicyOutput struct{ *pulumi.OutputState }

func (ServiceConnectionMapIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceConnectionMapIamPolicy)(nil)).Elem()
}

func (o ServiceConnectionMapIamPolicyOutput) ToServiceConnectionMapIamPolicyOutput() ServiceConnectionMapIamPolicyOutput {
	return o
}

func (o ServiceConnectionMapIamPolicyOutput) ToServiceConnectionMapIamPolicyOutputWithContext(ctx context.Context) ServiceConnectionMapIamPolicyOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o ServiceConnectionMapIamPolicyOutput) AuditConfigs() AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *ServiceConnectionMapIamPolicy) AuditConfigResponseArrayOutput { return v.AuditConfigs }).(AuditConfigResponseArrayOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o ServiceConnectionMapIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *ServiceConnectionMapIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o ServiceConnectionMapIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceConnectionMapIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ServiceConnectionMapIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceConnectionMapIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServiceConnectionMapIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceConnectionMapIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ServiceConnectionMapIamPolicyOutput) ServiceConnectionMapId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceConnectionMapIamPolicy) pulumi.StringOutput { return v.ServiceConnectionMapId }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o ServiceConnectionMapIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *ServiceConnectionMapIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceConnectionMapIamPolicyInput)(nil)).Elem(), &ServiceConnectionMapIamPolicy{})
	pulumi.RegisterOutputType(ServiceConnectionMapIamPolicyOutput{})
}
