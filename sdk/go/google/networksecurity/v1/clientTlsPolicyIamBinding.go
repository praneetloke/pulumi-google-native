// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type ClientTlsPolicyIamBinding struct {
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

// NewClientTlsPolicyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewClientTlsPolicyIamBinding(ctx *pulumi.Context,
	name string, args *ClientTlsPolicyIamBindingArgs, opts ...pulumi.ResourceOption) (*ClientTlsPolicyIamBinding, error) {
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
	var resource ClientTlsPolicyIamBinding
	err := ctx.RegisterResource("google-native:networksecurity/v1:ClientTlsPolicyIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientTlsPolicyIamBinding gets an existing ClientTlsPolicyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientTlsPolicyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientTlsPolicyIamBindingState, opts ...pulumi.ResourceOption) (*ClientTlsPolicyIamBinding, error) {
	var resource ClientTlsPolicyIamBinding
	err := ctx.ReadResource("google-native:networksecurity/v1:ClientTlsPolicyIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientTlsPolicyIamBinding resources.
type clientTlsPolicyIamBindingState struct {
}

type ClientTlsPolicyIamBindingState struct {
}

func (ClientTlsPolicyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientTlsPolicyIamBindingState)(nil)).Elem()
}

type clientTlsPolicyIamBindingArgs struct {
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

// The set of arguments for constructing a ClientTlsPolicyIamBinding resource.
type ClientTlsPolicyIamBindingArgs struct {
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

func (ClientTlsPolicyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientTlsPolicyIamBindingArgs)(nil)).Elem()
}

type ClientTlsPolicyIamBindingInput interface {
	pulumi.Input

	ToClientTlsPolicyIamBindingOutput() ClientTlsPolicyIamBindingOutput
	ToClientTlsPolicyIamBindingOutputWithContext(ctx context.Context) ClientTlsPolicyIamBindingOutput
}

func (*ClientTlsPolicyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientTlsPolicyIamBinding)(nil)).Elem()
}

func (i *ClientTlsPolicyIamBinding) ToClientTlsPolicyIamBindingOutput() ClientTlsPolicyIamBindingOutput {
	return i.ToClientTlsPolicyIamBindingOutputWithContext(context.Background())
}

func (i *ClientTlsPolicyIamBinding) ToClientTlsPolicyIamBindingOutputWithContext(ctx context.Context) ClientTlsPolicyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientTlsPolicyIamBindingOutput)
}

type ClientTlsPolicyIamBindingOutput struct{ *pulumi.OutputState }

func (ClientTlsPolicyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientTlsPolicyIamBinding)(nil)).Elem()
}

func (o ClientTlsPolicyIamBindingOutput) ToClientTlsPolicyIamBindingOutput() ClientTlsPolicyIamBindingOutput {
	return o
}

func (o ClientTlsPolicyIamBindingOutput) ToClientTlsPolicyIamBindingOutputWithContext(ctx context.Context) ClientTlsPolicyIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o ClientTlsPolicyIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *ClientTlsPolicyIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o ClientTlsPolicyIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientTlsPolicyIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o ClientTlsPolicyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClientTlsPolicyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o ClientTlsPolicyIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientTlsPolicyIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o ClientTlsPolicyIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientTlsPolicyIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one `IamBinding` can be used per role.
func (o ClientTlsPolicyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientTlsPolicyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientTlsPolicyIamBindingInput)(nil)).Elem(), &ClientTlsPolicyIamBinding{})
	pulumi.RegisterOutputType(ClientTlsPolicyIamBindingOutput{})
}
