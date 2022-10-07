// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type AppConnectionIamMember struct {
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

// NewAppConnectionIamMember registers a new resource with the given unique name, arguments, and options.
func NewAppConnectionIamMember(ctx *pulumi.Context,
	name string, args *AppConnectionIamMemberArgs, opts ...pulumi.ResourceOption) (*AppConnectionIamMember, error) {
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
	var resource AppConnectionIamMember
	err := ctx.RegisterResource("google-native:beyondcorp/v1alpha:AppConnectionIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppConnectionIamMember gets an existing AppConnectionIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppConnectionIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppConnectionIamMemberState, opts ...pulumi.ResourceOption) (*AppConnectionIamMember, error) {
	var resource AppConnectionIamMember
	err := ctx.ReadResource("google-native:beyondcorp/v1alpha:AppConnectionIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppConnectionIamMember resources.
type appConnectionIamMemberState struct {
}

type AppConnectionIamMemberState struct {
}

func (AppConnectionIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectionIamMemberState)(nil)).Elem()
}

type appConnectionIamMemberArgs struct {
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

// The set of arguments for constructing a AppConnectionIamMember resource.
type AppConnectionIamMemberArgs struct {
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

func (AppConnectionIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectionIamMemberArgs)(nil)).Elem()
}

type AppConnectionIamMemberInput interface {
	pulumi.Input

	ToAppConnectionIamMemberOutput() AppConnectionIamMemberOutput
	ToAppConnectionIamMemberOutputWithContext(ctx context.Context) AppConnectionIamMemberOutput
}

func (*AppConnectionIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnectionIamMember)(nil)).Elem()
}

func (i *AppConnectionIamMember) ToAppConnectionIamMemberOutput() AppConnectionIamMemberOutput {
	return i.ToAppConnectionIamMemberOutputWithContext(context.Background())
}

func (i *AppConnectionIamMember) ToAppConnectionIamMemberOutputWithContext(ctx context.Context) AppConnectionIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConnectionIamMemberOutput)
}

type AppConnectionIamMemberOutput struct{ *pulumi.OutputState }

func (AppConnectionIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnectionIamMember)(nil)).Elem()
}

func (o AppConnectionIamMemberOutput) ToAppConnectionIamMemberOutput() AppConnectionIamMemberOutput {
	return o
}

func (o AppConnectionIamMemberOutput) ToAppConnectionIamMemberOutputWithContext(ctx context.Context) AppConnectionIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o AppConnectionIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *AppConnectionIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o AppConnectionIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o AppConnectionIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o AppConnectionIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o AppConnectionIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o AppConnectionIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectionIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppConnectionIamMemberInput)(nil)).Elem(), &AppConnectionIamMember{})
	pulumi.RegisterOutputType(AppConnectionIamMemberOutput{})
}
