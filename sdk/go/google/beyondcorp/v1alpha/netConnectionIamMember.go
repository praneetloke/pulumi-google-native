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
type NetConnectionIamMember struct {
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

// NewNetConnectionIamMember registers a new resource with the given unique name, arguments, and options.
func NewNetConnectionIamMember(ctx *pulumi.Context,
	name string, args *NetConnectionIamMemberArgs, opts ...pulumi.ResourceOption) (*NetConnectionIamMember, error) {
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
	var resource NetConnectionIamMember
	err := ctx.RegisterResource("google-native:beyondcorp/v1alpha:NetConnectionIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetConnectionIamMember gets an existing NetConnectionIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetConnectionIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetConnectionIamMemberState, opts ...pulumi.ResourceOption) (*NetConnectionIamMember, error) {
	var resource NetConnectionIamMember
	err := ctx.ReadResource("google-native:beyondcorp/v1alpha:NetConnectionIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetConnectionIamMember resources.
type netConnectionIamMemberState struct {
}

type NetConnectionIamMemberState struct {
}

func (NetConnectionIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*netConnectionIamMemberState)(nil)).Elem()
}

type netConnectionIamMemberArgs struct {
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

// The set of arguments for constructing a NetConnectionIamMember resource.
type NetConnectionIamMemberArgs struct {
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

func (NetConnectionIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*netConnectionIamMemberArgs)(nil)).Elem()
}

type NetConnectionIamMemberInput interface {
	pulumi.Input

	ToNetConnectionIamMemberOutput() NetConnectionIamMemberOutput
	ToNetConnectionIamMemberOutputWithContext(ctx context.Context) NetConnectionIamMemberOutput
}

func (*NetConnectionIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**NetConnectionIamMember)(nil)).Elem()
}

func (i *NetConnectionIamMember) ToNetConnectionIamMemberOutput() NetConnectionIamMemberOutput {
	return i.ToNetConnectionIamMemberOutputWithContext(context.Background())
}

func (i *NetConnectionIamMember) ToNetConnectionIamMemberOutputWithContext(ctx context.Context) NetConnectionIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetConnectionIamMemberOutput)
}

type NetConnectionIamMemberOutput struct{ *pulumi.OutputState }

func (NetConnectionIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetConnectionIamMember)(nil)).Elem()
}

func (o NetConnectionIamMemberOutput) ToNetConnectionIamMemberOutput() NetConnectionIamMemberOutput {
	return o
}

func (o NetConnectionIamMemberOutput) ToNetConnectionIamMemberOutputWithContext(ctx context.Context) NetConnectionIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o NetConnectionIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *NetConnectionIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o NetConnectionIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetConnectionIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o NetConnectionIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *NetConnectionIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o NetConnectionIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetConnectionIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o NetConnectionIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetConnectionIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o NetConnectionIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *NetConnectionIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetConnectionIamMemberInput)(nil)).Elem(), &NetConnectionIamMember{})
	pulumi.RegisterOutputType(NetConnectionIamMemberOutput{})
}
