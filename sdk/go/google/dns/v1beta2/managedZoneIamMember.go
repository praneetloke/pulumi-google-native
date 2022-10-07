// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type ManagedZoneIamMember struct {
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

// NewManagedZoneIamMember registers a new resource with the given unique name, arguments, and options.
func NewManagedZoneIamMember(ctx *pulumi.Context,
	name string, args *ManagedZoneIamMemberArgs, opts ...pulumi.ResourceOption) (*ManagedZoneIamMember, error) {
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
	var resource ManagedZoneIamMember
	err := ctx.RegisterResource("google-native:dns/v1beta2:ManagedZoneIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedZoneIamMember gets an existing ManagedZoneIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedZoneIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedZoneIamMemberState, opts ...pulumi.ResourceOption) (*ManagedZoneIamMember, error) {
	var resource ManagedZoneIamMember
	err := ctx.ReadResource("google-native:dns/v1beta2:ManagedZoneIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedZoneIamMember resources.
type managedZoneIamMemberState struct {
}

type ManagedZoneIamMemberState struct {
}

func (ManagedZoneIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedZoneIamMemberState)(nil)).Elem()
}

type managedZoneIamMemberArgs struct {
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

// The set of arguments for constructing a ManagedZoneIamMember resource.
type ManagedZoneIamMemberArgs struct {
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

func (ManagedZoneIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedZoneIamMemberArgs)(nil)).Elem()
}

type ManagedZoneIamMemberInput interface {
	pulumi.Input

	ToManagedZoneIamMemberOutput() ManagedZoneIamMemberOutput
	ToManagedZoneIamMemberOutputWithContext(ctx context.Context) ManagedZoneIamMemberOutput
}

func (*ManagedZoneIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedZoneIamMember)(nil)).Elem()
}

func (i *ManagedZoneIamMember) ToManagedZoneIamMemberOutput() ManagedZoneIamMemberOutput {
	return i.ToManagedZoneIamMemberOutputWithContext(context.Background())
}

func (i *ManagedZoneIamMember) ToManagedZoneIamMemberOutputWithContext(ctx context.Context) ManagedZoneIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedZoneIamMemberOutput)
}

type ManagedZoneIamMemberOutput struct{ *pulumi.OutputState }

func (ManagedZoneIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedZoneIamMember)(nil)).Elem()
}

func (o ManagedZoneIamMemberOutput) ToManagedZoneIamMemberOutput() ManagedZoneIamMemberOutput {
	return o
}

func (o ManagedZoneIamMemberOutput) ToManagedZoneIamMemberOutputWithContext(ctx context.Context) ManagedZoneIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o ManagedZoneIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *ManagedZoneIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o ManagedZoneIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZoneIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o ManagedZoneIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZoneIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o ManagedZoneIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZoneIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o ManagedZoneIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZoneIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o ManagedZoneIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedZoneIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedZoneIamMemberInput)(nil)).Elem(), &ManagedZoneIamMember{})
	pulumi.RegisterOutputType(ManagedZoneIamMemberOutput{})
}
