// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type GatewayIamMember struct {
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

// NewGatewayIamMember registers a new resource with the given unique name, arguments, and options.
func NewGatewayIamMember(ctx *pulumi.Context,
	name string, args *GatewayIamMemberArgs, opts ...pulumi.ResourceOption) (*GatewayIamMember, error) {
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
	var resource GatewayIamMember
	err := ctx.RegisterResource("google-native:apigateway/v1beta:GatewayIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayIamMember gets an existing GatewayIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayIamMemberState, opts ...pulumi.ResourceOption) (*GatewayIamMember, error) {
	var resource GatewayIamMember
	err := ctx.ReadResource("google-native:apigateway/v1beta:GatewayIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayIamMember resources.
type gatewayIamMemberState struct {
}

type GatewayIamMemberState struct {
}

func (GatewayIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamMemberState)(nil)).Elem()
}

type gatewayIamMemberArgs struct {
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

// The set of arguments for constructing a GatewayIamMember resource.
type GatewayIamMemberArgs struct {
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

func (GatewayIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamMemberArgs)(nil)).Elem()
}

type GatewayIamMemberInput interface {
	pulumi.Input

	ToGatewayIamMemberOutput() GatewayIamMemberOutput
	ToGatewayIamMemberOutputWithContext(ctx context.Context) GatewayIamMemberOutput
}

func (*GatewayIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayIamMember)(nil)).Elem()
}

func (i *GatewayIamMember) ToGatewayIamMemberOutput() GatewayIamMemberOutput {
	return i.ToGatewayIamMemberOutputWithContext(context.Background())
}

func (i *GatewayIamMember) ToGatewayIamMemberOutputWithContext(ctx context.Context) GatewayIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamMemberOutput)
}

type GatewayIamMemberOutput struct{ *pulumi.OutputState }

func (GatewayIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayIamMember)(nil)).Elem()
}

func (o GatewayIamMemberOutput) ToGatewayIamMemberOutput() GatewayIamMemberOutput {
	return o
}

func (o GatewayIamMemberOutput) ToGatewayIamMemberOutputWithContext(ctx context.Context) GatewayIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o GatewayIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *GatewayIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o GatewayIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o GatewayIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o GatewayIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o GatewayIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o GatewayIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayIamMemberInput)(nil)).Elem(), &GatewayIamMember{})
	pulumi.RegisterOutputType(GatewayIamMemberOutput{})
}
