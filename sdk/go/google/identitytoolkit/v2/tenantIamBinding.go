// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy for a resource. If the policy exists, it is replaced. Caller must have the right Google IAM permission on the resource.
type TenantIamBinding struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identities that will be granted the privilege in role. Each entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one `IamBinding` can be used per role.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewTenantIamBinding registers a new resource with the given unique name, arguments, and options.
func NewTenantIamBinding(ctx *pulumi.Context,
	name string, args *TenantIamBindingArgs, opts ...pulumi.ResourceOption) (*TenantIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TenantIamBinding
	err := ctx.RegisterResource("google-native:identitytoolkit/v2:TenantIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenantIamBinding gets an existing TenantIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenantIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantIamBindingState, opts ...pulumi.ResourceOption) (*TenantIamBinding, error) {
	var resource TenantIamBinding
	err := ctx.ReadResource("google-native:identitytoolkit/v2:TenantIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TenantIamBinding resources.
type tenantIamBindingState struct {
}

type TenantIamBindingState struct {
}

func (TenantIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantIamBindingState)(nil)).Elem()
}

type tenantIamBindingArgs struct {
	// An IAM Condition for a given binding.
	Condition *iam.Condition `pulumi:"condition"`
	// Identities that will be granted the privilege in role. Each entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members []string `pulumi:"members"`
	// The name of the resource to manage IAM policies for.
	Name string `pulumi:"name"`
	// The role that should be applied. Only one `IamBinding` can be used per role.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a TenantIamBinding resource.
type TenantIamBindingArgs struct {
	// An IAM Condition for a given binding.
	Condition iam.ConditionPtrInput
	// Identities that will be granted the privilege in role. Each entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayInput
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringInput
	// The role that should be applied. Only one `IamBinding` can be used per role.
	Role pulumi.StringInput
}

func (TenantIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantIamBindingArgs)(nil)).Elem()
}

type TenantIamBindingInput interface {
	pulumi.Input

	ToTenantIamBindingOutput() TenantIamBindingOutput
	ToTenantIamBindingOutputWithContext(ctx context.Context) TenantIamBindingOutput
}

func (*TenantIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantIamBinding)(nil)).Elem()
}

func (i *TenantIamBinding) ToTenantIamBindingOutput() TenantIamBindingOutput {
	return i.ToTenantIamBindingOutputWithContext(context.Background())
}

func (i *TenantIamBinding) ToTenantIamBindingOutputWithContext(ctx context.Context) TenantIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantIamBindingOutput)
}

type TenantIamBindingOutput struct{ *pulumi.OutputState }

func (TenantIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantIamBinding)(nil)).Elem()
}

func (o TenantIamBindingOutput) ToTenantIamBindingOutput() TenantIamBindingOutput {
	return o
}

func (o TenantIamBindingOutput) ToTenantIamBindingOutputWithContext(ctx context.Context) TenantIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o TenantIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *TenantIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o TenantIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o TenantIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TenantIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o TenantIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o TenantIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one `IamBinding` can be used per role.
func (o TenantIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TenantIamBindingInput)(nil)).Elem(), &TenantIamBinding{})
	pulumi.RegisterOutputType(TenantIamBindingOutput{})
}
