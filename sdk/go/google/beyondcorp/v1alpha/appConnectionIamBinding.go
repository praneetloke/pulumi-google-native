// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type AppConnectionIamBinding struct {
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

// NewAppConnectionIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAppConnectionIamBinding(ctx *pulumi.Context,
	name string, args *AppConnectionIamBindingArgs, opts ...pulumi.ResourceOption) (*AppConnectionIamBinding, error) {
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
	var resource AppConnectionIamBinding
	err := ctx.RegisterResource("google-native:beyondcorp/v1alpha:AppConnectionIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppConnectionIamBinding gets an existing AppConnectionIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppConnectionIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppConnectionIamBindingState, opts ...pulumi.ResourceOption) (*AppConnectionIamBinding, error) {
	var resource AppConnectionIamBinding
	err := ctx.ReadResource("google-native:beyondcorp/v1alpha:AppConnectionIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppConnectionIamBinding resources.
type appConnectionIamBindingState struct {
}

type AppConnectionIamBindingState struct {
}

func (AppConnectionIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectionIamBindingState)(nil)).Elem()
}

type appConnectionIamBindingArgs struct {
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

// The set of arguments for constructing a AppConnectionIamBinding resource.
type AppConnectionIamBindingArgs struct {
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

func (AppConnectionIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectionIamBindingArgs)(nil)).Elem()
}

type AppConnectionIamBindingInput interface {
	pulumi.Input

	ToAppConnectionIamBindingOutput() AppConnectionIamBindingOutput
	ToAppConnectionIamBindingOutputWithContext(ctx context.Context) AppConnectionIamBindingOutput
}

func (*AppConnectionIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnectionIamBinding)(nil)).Elem()
}

func (i *AppConnectionIamBinding) ToAppConnectionIamBindingOutput() AppConnectionIamBindingOutput {
	return i.ToAppConnectionIamBindingOutputWithContext(context.Background())
}

func (i *AppConnectionIamBinding) ToAppConnectionIamBindingOutputWithContext(ctx context.Context) AppConnectionIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConnectionIamBindingOutput)
}

type AppConnectionIamBindingOutput struct{ *pulumi.OutputState }

func (AppConnectionIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnectionIamBinding)(nil)).Elem()
}

func (o AppConnectionIamBindingOutput) ToAppConnectionIamBindingOutput() AppConnectionIamBindingOutput {
	return o
}

func (o AppConnectionIamBindingOutput) ToAppConnectionIamBindingOutputWithContext(ctx context.Context) AppConnectionIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o AppConnectionIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *AppConnectionIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o AppConnectionIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o AppConnectionIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppConnectionIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o AppConnectionIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o AppConnectionIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one `IamBinding` can be used per role.
func (o AppConnectionIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppConnectionIamBindingInput)(nil)).Elem(), &AppConnectionIamBinding{})
	pulumi.RegisterOutputType(AppConnectionIamBindingOutput{})
}
