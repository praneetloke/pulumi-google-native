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
type AuthorizationPolicyIamMember struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identity that will be granted the privilege in role. The entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringOutput `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAuthorizationPolicyIamMember registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationPolicyIamMember(ctx *pulumi.Context,
	name string, args *AuthorizationPolicyIamMemberArgs, opts ...pulumi.ResourceOption) (*AuthorizationPolicyIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthorizationPolicyIamMember
	err := ctx.RegisterResource("google-native:networksecurity/v1:AuthorizationPolicyIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationPolicyIamMember gets an existing AuthorizationPolicyIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationPolicyIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationPolicyIamMemberState, opts ...pulumi.ResourceOption) (*AuthorizationPolicyIamMember, error) {
	var resource AuthorizationPolicyIamMember
	err := ctx.ReadResource("google-native:networksecurity/v1:AuthorizationPolicyIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationPolicyIamMember resources.
type authorizationPolicyIamMemberState struct {
}

type AuthorizationPolicyIamMemberState struct {
}

func (AuthorizationPolicyIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationPolicyIamMemberState)(nil)).Elem()
}

type authorizationPolicyIamMemberArgs struct {
	// An IAM Condition for a given binding.
	Condition *iam.Condition `pulumi:"condition"`
	// Identity that will be granted the privilege in role. The entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member string `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name string `pulumi:"name"`
	// The role that should be applied.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AuthorizationPolicyIamMember resource.
type AuthorizationPolicyIamMemberArgs struct {
	// An IAM Condition for a given binding.
	Condition iam.ConditionPtrInput
	// Identity that will be granted the privilege in role. The entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringInput
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringInput
	// The role that should be applied.
	Role pulumi.StringInput
}

func (AuthorizationPolicyIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationPolicyIamMemberArgs)(nil)).Elem()
}

type AuthorizationPolicyIamMemberInput interface {
	pulumi.Input

	ToAuthorizationPolicyIamMemberOutput() AuthorizationPolicyIamMemberOutput
	ToAuthorizationPolicyIamMemberOutputWithContext(ctx context.Context) AuthorizationPolicyIamMemberOutput
}

func (*AuthorizationPolicyIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationPolicyIamMember)(nil)).Elem()
}

func (i *AuthorizationPolicyIamMember) ToAuthorizationPolicyIamMemberOutput() AuthorizationPolicyIamMemberOutput {
	return i.ToAuthorizationPolicyIamMemberOutputWithContext(context.Background())
}

func (i *AuthorizationPolicyIamMember) ToAuthorizationPolicyIamMemberOutputWithContext(ctx context.Context) AuthorizationPolicyIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationPolicyIamMemberOutput)
}

type AuthorizationPolicyIamMemberOutput struct{ *pulumi.OutputState }

func (AuthorizationPolicyIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationPolicyIamMember)(nil)).Elem()
}

func (o AuthorizationPolicyIamMemberOutput) ToAuthorizationPolicyIamMemberOutput() AuthorizationPolicyIamMemberOutput {
	return o
}

func (o AuthorizationPolicyIamMemberOutput) ToAuthorizationPolicyIamMemberOutputWithContext(ctx context.Context) AuthorizationPolicyIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o AuthorizationPolicyIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *AuthorizationPolicyIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o AuthorizationPolicyIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationPolicyIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o AuthorizationPolicyIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationPolicyIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o AuthorizationPolicyIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationPolicyIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o AuthorizationPolicyIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationPolicyIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o AuthorizationPolicyIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationPolicyIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationPolicyIamMemberInput)(nil)).Elem(), &AuthorizationPolicyIamMember{})
	pulumi.RegisterOutputType(AuthorizationPolicyIamMemberOutput{})
}
