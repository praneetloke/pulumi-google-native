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
type AppConnectorIamMember struct {
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

// NewAppConnectorIamMember registers a new resource with the given unique name, arguments, and options.
func NewAppConnectorIamMember(ctx *pulumi.Context,
	name string, args *AppConnectorIamMemberArgs, opts ...pulumi.ResourceOption) (*AppConnectorIamMember, error) {
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
	var resource AppConnectorIamMember
	err := ctx.RegisterResource("google-native:beyondcorp/v1alpha:AppConnectorIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppConnectorIamMember gets an existing AppConnectorIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppConnectorIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppConnectorIamMemberState, opts ...pulumi.ResourceOption) (*AppConnectorIamMember, error) {
	var resource AppConnectorIamMember
	err := ctx.ReadResource("google-native:beyondcorp/v1alpha:AppConnectorIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppConnectorIamMember resources.
type appConnectorIamMemberState struct {
}

type AppConnectorIamMemberState struct {
}

func (AppConnectorIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectorIamMemberState)(nil)).Elem()
}

type appConnectorIamMemberArgs struct {
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

// The set of arguments for constructing a AppConnectorIamMember resource.
type AppConnectorIamMemberArgs struct {
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

func (AppConnectorIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appConnectorIamMemberArgs)(nil)).Elem()
}

type AppConnectorIamMemberInput interface {
	pulumi.Input

	ToAppConnectorIamMemberOutput() AppConnectorIamMemberOutput
	ToAppConnectorIamMemberOutputWithContext(ctx context.Context) AppConnectorIamMemberOutput
}

func (*AppConnectorIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnectorIamMember)(nil)).Elem()
}

func (i *AppConnectorIamMember) ToAppConnectorIamMemberOutput() AppConnectorIamMemberOutput {
	return i.ToAppConnectorIamMemberOutputWithContext(context.Background())
}

func (i *AppConnectorIamMember) ToAppConnectorIamMemberOutputWithContext(ctx context.Context) AppConnectorIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppConnectorIamMemberOutput)
}

type AppConnectorIamMemberOutput struct{ *pulumi.OutputState }

func (AppConnectorIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppConnectorIamMember)(nil)).Elem()
}

func (o AppConnectorIamMemberOutput) ToAppConnectorIamMemberOutput() AppConnectorIamMemberOutput {
	return o
}

func (o AppConnectorIamMemberOutput) ToAppConnectorIamMemberOutputWithContext(ctx context.Context) AppConnectorIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o AppConnectorIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *AppConnectorIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o AppConnectorIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectorIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o AppConnectorIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectorIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o AppConnectorIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectorIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o AppConnectorIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectorIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o AppConnectorIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AppConnectorIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppConnectorIamMemberInput)(nil)).Elem(), &AppConnectorIamMember{})
	pulumi.RegisterOutputType(AppConnectorIamMemberOutput{})
}
