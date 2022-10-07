// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type AttestorIamBinding struct {
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

// NewAttestorIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAttestorIamBinding(ctx *pulumi.Context,
	name string, args *AttestorIamBindingArgs, opts ...pulumi.ResourceOption) (*AttestorIamBinding, error) {
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
	var resource AttestorIamBinding
	err := ctx.RegisterResource("google-native:binaryauthorization/v1beta1:AttestorIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestorIamBinding gets an existing AttestorIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestorIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestorIamBindingState, opts ...pulumi.ResourceOption) (*AttestorIamBinding, error) {
	var resource AttestorIamBinding
	err := ctx.ReadResource("google-native:binaryauthorization/v1beta1:AttestorIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttestorIamBinding resources.
type attestorIamBindingState struct {
}

type AttestorIamBindingState struct {
}

func (AttestorIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamBindingState)(nil)).Elem()
}

type attestorIamBindingArgs struct {
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

// The set of arguments for constructing a AttestorIamBinding resource.
type AttestorIamBindingArgs struct {
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

func (AttestorIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamBindingArgs)(nil)).Elem()
}

type AttestorIamBindingInput interface {
	pulumi.Input

	ToAttestorIamBindingOutput() AttestorIamBindingOutput
	ToAttestorIamBindingOutputWithContext(ctx context.Context) AttestorIamBindingOutput
}

func (*AttestorIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestorIamBinding)(nil)).Elem()
}

func (i *AttestorIamBinding) ToAttestorIamBindingOutput() AttestorIamBindingOutput {
	return i.ToAttestorIamBindingOutputWithContext(context.Background())
}

func (i *AttestorIamBinding) ToAttestorIamBindingOutputWithContext(ctx context.Context) AttestorIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorIamBindingOutput)
}

type AttestorIamBindingOutput struct{ *pulumi.OutputState }

func (AttestorIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestorIamBinding)(nil)).Elem()
}

func (o AttestorIamBindingOutput) ToAttestorIamBindingOutput() AttestorIamBindingOutput {
	return o
}

func (o AttestorIamBindingOutput) ToAttestorIamBindingOutputWithContext(ctx context.Context) AttestorIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o AttestorIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *AttestorIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o AttestorIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestorIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o AttestorIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttestorIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o AttestorIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestorIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o AttestorIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestorIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one `IamBinding` can be used per role.
func (o AttestorIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestorIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttestorIamBindingInput)(nil)).Elem(), &AttestorIamBinding{})
	pulumi.RegisterOutputType(AttestorIamBindingOutput{})
}
