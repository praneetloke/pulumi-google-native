// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

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
type ServiceDatabaseTableIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
	Bindings   BindingResponseArrayOutput `pulumi:"bindings"`
	DatabaseId pulumi.StringOutput        `pulumi:"databaseId"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag      pulumi.StringOutput `pulumi:"etag"`
	Location  pulumi.StringOutput `pulumi:"location"`
	Project   pulumi.StringOutput `pulumi:"project"`
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	TableId   pulumi.StringOutput `pulumi:"tableId"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewServiceDatabaseTableIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceDatabaseTableIamPolicy(ctx *pulumi.Context,
	name string, args *ServiceDatabaseTableIamPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceDatabaseTableIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.TableId == nil {
		return nil, errors.New("invalid value for required argument 'TableId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"databaseId",
		"location",
		"project",
		"serviceId",
		"tableId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceDatabaseTableIamPolicy
	err := ctx.RegisterResource("google-native:metastore/v1alpha:ServiceDatabaseTableIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceDatabaseTableIamPolicy gets an existing ServiceDatabaseTableIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceDatabaseTableIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceDatabaseTableIamPolicyState, opts ...pulumi.ResourceOption) (*ServiceDatabaseTableIamPolicy, error) {
	var resource ServiceDatabaseTableIamPolicy
	err := ctx.ReadResource("google-native:metastore/v1alpha:ServiceDatabaseTableIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceDatabaseTableIamPolicy resources.
type serviceDatabaseTableIamPolicyState struct {
}

type ServiceDatabaseTableIamPolicyState struct {
}

func (ServiceDatabaseTableIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceDatabaseTableIamPolicyState)(nil)).Elem()
}

type serviceDatabaseTableIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
	Bindings   []Binding `pulumi:"bindings"`
	DatabaseId string    `pulumi:"databaseId"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag      *string `pulumi:"etag"`
	Location  *string `pulumi:"location"`
	Project   *string `pulumi:"project"`
	ServiceId string  `pulumi:"serviceId"`
	TableId   string  `pulumi:"tableId"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used:paths: "bindings, etag"
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a ServiceDatabaseTableIamPolicy resource.
type ServiceDatabaseTableIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
	Bindings   BindingArrayInput
	DatabaseId pulumi.StringInput
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
	Etag      pulumi.StringPtrInput
	Location  pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	ServiceId pulumi.StringInput
	TableId   pulumi.StringInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used:paths: "bindings, etag"
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (ServiceDatabaseTableIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceDatabaseTableIamPolicyArgs)(nil)).Elem()
}

type ServiceDatabaseTableIamPolicyInput interface {
	pulumi.Input

	ToServiceDatabaseTableIamPolicyOutput() ServiceDatabaseTableIamPolicyOutput
	ToServiceDatabaseTableIamPolicyOutputWithContext(ctx context.Context) ServiceDatabaseTableIamPolicyOutput
}

func (*ServiceDatabaseTableIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDatabaseTableIamPolicy)(nil)).Elem()
}

func (i *ServiceDatabaseTableIamPolicy) ToServiceDatabaseTableIamPolicyOutput() ServiceDatabaseTableIamPolicyOutput {
	return i.ToServiceDatabaseTableIamPolicyOutputWithContext(context.Background())
}

func (i *ServiceDatabaseTableIamPolicy) ToServiceDatabaseTableIamPolicyOutputWithContext(ctx context.Context) ServiceDatabaseTableIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDatabaseTableIamPolicyOutput)
}

type ServiceDatabaseTableIamPolicyOutput struct{ *pulumi.OutputState }

func (ServiceDatabaseTableIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDatabaseTableIamPolicy)(nil)).Elem()
}

func (o ServiceDatabaseTableIamPolicyOutput) ToServiceDatabaseTableIamPolicyOutput() ServiceDatabaseTableIamPolicyOutput {
	return o
}

func (o ServiceDatabaseTableIamPolicyOutput) ToServiceDatabaseTableIamPolicyOutputWithContext(ctx context.Context) ServiceDatabaseTableIamPolicyOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o ServiceDatabaseTableIamPolicyOutput) AuditConfigs() AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) AuditConfigResponseArrayOutput { return v.AuditConfigs }).(AuditConfigResponseArrayOutput)
}

// Associates a list of members, or principals, with a role. Optionally, may specify a condition that determines how and when the bindings are applied. Each of the bindings must contain at least one principal.The bindings in a Policy can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the bindings grant 50 different roles to user:alice@example.com, and not to any other principal, then you can add another 1,450 principals to the bindings in the Policy.
func (o ServiceDatabaseTableIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

func (o ServiceDatabaseTableIamPolicyOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the etag in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An etag is returned in the response to getIamPolicy, and systems are expected to put that etag in the request to setIamPolicy to ensure that their change will be applied to the same version of the policy.Important: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.
func (o ServiceDatabaseTableIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ServiceDatabaseTableIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServiceDatabaseTableIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ServiceDatabaseTableIamPolicyOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

func (o ServiceDatabaseTableIamPolicyOutput) TableId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) pulumi.StringOutput { return v.TableId }).(pulumi.StringOutput)
}

// Specifies the format of the policy.Valid values are 0, 1, and 3. Requests that specify an invalid value are rejected.Any operation that affects conditional role bindings must specify version 3. This requirement applies to the following operations: Getting a policy that includes a conditional role binding Adding a conditional role binding to a policy Changing a conditional role binding in a policy Removing any role binding, with or without a condition, from a policy that includes conditionsImportant: If you use IAM Conditions, you must include the etag field whenever you call setIamPolicy. If you omit this field, then IAM allows you to overwrite a version 3 policy with a version 1 policy, and all of the conditions in the version 3 policy are lost.If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset.To learn which resources support conditions in their IAM policies, see the IAM documentation (https://cloud.google.com/iam/help/conditions/resource-policies).
func (o ServiceDatabaseTableIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *ServiceDatabaseTableIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceDatabaseTableIamPolicyInput)(nil)).Elem(), &ServiceDatabaseTableIamPolicy{})
	pulumi.RegisterOutputType(ServiceDatabaseTableIamPolicyOutput{})
}
