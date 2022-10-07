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
type RegionJobIamBinding struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the identities requesting access for a Cloud Platform resource. members can have the following values: allUsers: A special identifier that represents anyone who is on the internet; with or without a Google account. allAuthenticatedUsers: A special identifier that represents anyone who is authenticated with a Google account or a service account. user:{emailid}: An email address that represents a specific Google account. For example, alice@example.com . serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com. group:{emailid}: An email address that represents a Google group. For example, admins@example.com. deleted:user:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a user that has been recently deleted. For example, alice@example.com?uid=123456789012345678901. If the user is recovered, this value reverts to user:{emailid} and the recovered user retains the role in the binding. deleted:serviceAccount:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901. If the service account is undeleted, this value reverts to serviceAccount:{emailid} and the undeleted service account retains the role in the binding. deleted:group:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, admins@example.com?uid=123456789012345678901. If the group is recovered, this value reverts to group:{emailid} and the recovered group retains the role in the binding. domain:{domain}: The G Suite domain (primary) that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// Role that is assigned to members. For example, roles/viewer, roles/editor, or roles/owner.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRegionJobIamBinding registers a new resource with the given unique name, arguments, and options.
func NewRegionJobIamBinding(ctx *pulumi.Context,
	name string, args *RegionJobIamBindingArgs, opts ...pulumi.ResourceOption) (*RegionJobIamBinding, error) {
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
	var resource RegionJobIamBinding
	err := ctx.RegisterResource("google-native:dataproc/v1beta2:RegionJobIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionJobIamBinding gets an existing RegionJobIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionJobIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionJobIamBindingState, opts ...pulumi.ResourceOption) (*RegionJobIamBinding, error) {
	var resource RegionJobIamBinding
	err := ctx.ReadResource("google-native:dataproc/v1beta2:RegionJobIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionJobIamBinding resources.
type regionJobIamBindingState struct {
}

type RegionJobIamBindingState struct {
}

func (RegionJobIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionJobIamBindingState)(nil)).Elem()
}

type regionJobIamBindingArgs struct {
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

// The set of arguments for constructing a RegionJobIamBinding resource.
type RegionJobIamBindingArgs struct {
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

func (RegionJobIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionJobIamBindingArgs)(nil)).Elem()
}

type RegionJobIamBindingInput interface {
	pulumi.Input

	ToRegionJobIamBindingOutput() RegionJobIamBindingOutput
	ToRegionJobIamBindingOutputWithContext(ctx context.Context) RegionJobIamBindingOutput
}

func (*RegionJobIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionJobIamBinding)(nil)).Elem()
}

func (i *RegionJobIamBinding) ToRegionJobIamBindingOutput() RegionJobIamBindingOutput {
	return i.ToRegionJobIamBindingOutputWithContext(context.Background())
}

func (i *RegionJobIamBinding) ToRegionJobIamBindingOutputWithContext(ctx context.Context) RegionJobIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionJobIamBindingOutput)
}

type RegionJobIamBindingOutput struct{ *pulumi.OutputState }

func (RegionJobIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionJobIamBinding)(nil)).Elem()
}

func (o RegionJobIamBindingOutput) ToRegionJobIamBindingOutput() RegionJobIamBindingOutput {
	return o
}

func (o RegionJobIamBindingOutput) ToRegionJobIamBindingOutputWithContext(ctx context.Context) RegionJobIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o RegionJobIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *RegionJobIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o RegionJobIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionJobIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the identities requesting access for a Cloud Platform resource. members can have the following values: allUsers: A special identifier that represents anyone who is on the internet; with or without a Google account. allAuthenticatedUsers: A special identifier that represents anyone who is authenticated with a Google account or a service account. user:{emailid}: An email address that represents a specific Google account. For example, alice@example.com . serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com. group:{emailid}: An email address that represents a Google group. For example, admins@example.com. deleted:user:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a user that has been recently deleted. For example, alice@example.com?uid=123456789012345678901. If the user is recovered, this value reverts to user:{emailid} and the recovered user retains the role in the binding. deleted:serviceAccount:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901. If the service account is undeleted, this value reverts to serviceAccount:{emailid} and the undeleted service account retains the role in the binding. deleted:group:{emailid}?uid={uniqueid}: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, admins@example.com?uid=123456789012345678901. If the group is recovered, this value reverts to group:{emailid} and the recovered group retains the role in the binding. domain:{domain}: The G Suite domain (primary) that represents all the users of that domain. For example, google.com or example.com.
func (o RegionJobIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RegionJobIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o RegionJobIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionJobIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o RegionJobIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionJobIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Role that is assigned to members. For example, roles/viewer, roles/editor, or roles/owner.
func (o RegionJobIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionJobIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionJobIamBindingInput)(nil)).Elem(), &RegionJobIamBinding{})
	pulumi.RegisterOutputType(RegionJobIamBindingOutput{})
}
