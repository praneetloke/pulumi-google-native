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

// Sets the access control policy on the specified resource. Replaces any existing policy.Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED errors.
type RegionWorkflowTemplateIamMember struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the identities requesting access for a Cloud Platform resource. members can have the following values: allUsers: A special identifier that represents anyone who is on the internet; with or without a Google account. allAuthenticatedUsers: A special identifier that represents anyone who is authenticated with a Google account or a service account. user:{emailid}: An email address that represents a specific Google account. For example, alice@example.com . serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com. group:{emailid}: An email address that represents a Google group. For example, admins@example.com. deleted:user:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a user that has been recently deleted. For example, alice@example.com?uid=123456789012345678901. If the user is recovered, this value reverts to user:{emailid} and the recovered user retains the role in the binding. deleted:serviceAccount:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901. If the service account is undeleted, this value reverts to serviceAccount:{emailid} and the undeleted service account retains the role in the binding. deleted:group:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, admins@example.com?uid=123456789012345678901. If the group is recovered, this value reverts to group:{emailid} and the recovered group retains the role in the binding. domain:{domain}: The G Suite domain (primary) that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringOutput `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// Role that is assigned to members. For example, roles/viewer, roles/editor, or roles/owner.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRegionWorkflowTemplateIamMember registers a new resource with the given unique name, arguments, and options.
func NewRegionWorkflowTemplateIamMember(ctx *pulumi.Context,
	name string, args *RegionWorkflowTemplateIamMemberArgs, opts ...pulumi.ResourceOption) (*RegionWorkflowTemplateIamMember, error) {
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
	var resource RegionWorkflowTemplateIamMember
	err := ctx.RegisterResource("google-native:dataproc/v1beta2:RegionWorkflowTemplateIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionWorkflowTemplateIamMember gets an existing RegionWorkflowTemplateIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionWorkflowTemplateIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionWorkflowTemplateIamMemberState, opts ...pulumi.ResourceOption) (*RegionWorkflowTemplateIamMember, error) {
	var resource RegionWorkflowTemplateIamMember
	err := ctx.ReadResource("google-native:dataproc/v1beta2:RegionWorkflowTemplateIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionWorkflowTemplateIamMember resources.
type regionWorkflowTemplateIamMemberState struct {
}

type RegionWorkflowTemplateIamMemberState struct {
}

func (RegionWorkflowTemplateIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionWorkflowTemplateIamMemberState)(nil)).Elem()
}

type regionWorkflowTemplateIamMemberArgs struct {
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

// The set of arguments for constructing a RegionWorkflowTemplateIamMember resource.
type RegionWorkflowTemplateIamMemberArgs struct {
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

func (RegionWorkflowTemplateIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionWorkflowTemplateIamMemberArgs)(nil)).Elem()
}

type RegionWorkflowTemplateIamMemberInput interface {
	pulumi.Input

	ToRegionWorkflowTemplateIamMemberOutput() RegionWorkflowTemplateIamMemberOutput
	ToRegionWorkflowTemplateIamMemberOutputWithContext(ctx context.Context) RegionWorkflowTemplateIamMemberOutput
}

func (*RegionWorkflowTemplateIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionWorkflowTemplateIamMember)(nil)).Elem()
}

func (i *RegionWorkflowTemplateIamMember) ToRegionWorkflowTemplateIamMemberOutput() RegionWorkflowTemplateIamMemberOutput {
	return i.ToRegionWorkflowTemplateIamMemberOutputWithContext(context.Background())
}

func (i *RegionWorkflowTemplateIamMember) ToRegionWorkflowTemplateIamMemberOutputWithContext(ctx context.Context) RegionWorkflowTemplateIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionWorkflowTemplateIamMemberOutput)
}

type RegionWorkflowTemplateIamMemberOutput struct{ *pulumi.OutputState }

func (RegionWorkflowTemplateIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionWorkflowTemplateIamMember)(nil)).Elem()
}

func (o RegionWorkflowTemplateIamMemberOutput) ToRegionWorkflowTemplateIamMemberOutput() RegionWorkflowTemplateIamMemberOutput {
	return o
}

func (o RegionWorkflowTemplateIamMemberOutput) ToRegionWorkflowTemplateIamMemberOutputWithContext(ctx context.Context) RegionWorkflowTemplateIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o RegionWorkflowTemplateIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *RegionWorkflowTemplateIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o RegionWorkflowTemplateIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionWorkflowTemplateIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the identities requesting access for a Cloud Platform resource. members can have the following values: allUsers: A special identifier that represents anyone who is on the internet; with or without a Google account. allAuthenticatedUsers: A special identifier that represents anyone who is authenticated with a Google account or a service account. user:{emailid}: An email address that represents a specific Google account. For example, alice@example.com . serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com. group:{emailid}: An email address that represents a Google group. For example, admins@example.com. deleted:user:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a user that has been recently deleted. For example, alice@example.com?uid=123456789012345678901. If the user is recovered, this value reverts to user:{emailid} and the recovered user retains the role in the binding. deleted:serviceAccount:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901. If the service account is undeleted, this value reverts to serviceAccount:{emailid} and the undeleted service account retains the role in the binding. deleted:group:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, admins@example.com?uid=123456789012345678901. If the group is recovered, this value reverts to group:{emailid} and the recovered group retains the role in the binding. domain:{domain}: The G Suite domain (primary) that represents all the users of that domain. For example, google.com or example.com.
func (o RegionWorkflowTemplateIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionWorkflowTemplateIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o RegionWorkflowTemplateIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionWorkflowTemplateIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o RegionWorkflowTemplateIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionWorkflowTemplateIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Role that is assigned to members. For example, roles/viewer, roles/editor, or roles/owner.
func (o RegionWorkflowTemplateIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionWorkflowTemplateIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionWorkflowTemplateIamMemberInput)(nil)).Elem(), &RegionWorkflowTemplateIamMember{})
	pulumi.RegisterOutputType(RegionWorkflowTemplateIamMemberOutput{})
}
