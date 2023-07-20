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
type ApiIamMember struct {
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

// NewApiIamMember registers a new resource with the given unique name, arguments, and options.
func NewApiIamMember(ctx *pulumi.Context,
	name string, args *ApiIamMemberArgs, opts ...pulumi.ResourceOption) (*ApiIamMember, error) {
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
	var resource ApiIamMember
	err := ctx.RegisterResource("google-native:apigateway/v1:ApiIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIamMember gets an existing ApiIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIamMemberState, opts ...pulumi.ResourceOption) (*ApiIamMember, error) {
	var resource ApiIamMember
	err := ctx.ReadResource("google-native:apigateway/v1:ApiIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIamMember resources.
type apiIamMemberState struct {
}

type ApiIamMemberState struct {
}

func (ApiIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamMemberState)(nil)).Elem()
}

type apiIamMemberArgs struct {
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

// The set of arguments for constructing a ApiIamMember resource.
type ApiIamMemberArgs struct {
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

func (ApiIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamMemberArgs)(nil)).Elem()
}

type ApiIamMemberInput interface {
	pulumi.Input

	ToApiIamMemberOutput() ApiIamMemberOutput
	ToApiIamMemberOutputWithContext(ctx context.Context) ApiIamMemberOutput
}

func (*ApiIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIamMember)(nil)).Elem()
}

func (i *ApiIamMember) ToApiIamMemberOutput() ApiIamMemberOutput {
	return i.ToApiIamMemberOutputWithContext(context.Background())
}

func (i *ApiIamMember) ToApiIamMemberOutputWithContext(ctx context.Context) ApiIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamMemberOutput)
}

type ApiIamMemberOutput struct{ *pulumi.OutputState }

func (ApiIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIamMember)(nil)).Elem()
}

func (o ApiIamMemberOutput) ToApiIamMemberOutput() ApiIamMemberOutput {
	return o
}

func (o ApiIamMemberOutput) ToApiIamMemberOutputWithContext(ctx context.Context) ApiIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o ApiIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *ApiIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o ApiIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o ApiIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o ApiIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o ApiIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o ApiIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiIamMemberInput)(nil)).Elem(), &ApiIamMember{})
	pulumi.RegisterOutputType(ApiIamMemberOutput{})
}
