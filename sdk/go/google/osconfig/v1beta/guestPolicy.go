// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create an OS Config guest policy.
type GuestPolicy struct {
	pulumi.CustomResourceState

	// Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	Assignment AssignmentResponseOutput `pulumi:"assignment"`
	// Time this guest policy was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
	PackageRepositories PackageRepositoryResponseArrayOutput `pulumi:"packageRepositories"`
	// The software packages to be managed by this policy.
	Packages PackageResponseArrayOutput `pulumi:"packages"`
	// A list of Recipes to install on the VM instance.
	Recipes SoftwareRecipeResponseArrayOutput `pulumi:"recipes"`
	// Last time this guest policy was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGuestPolicy registers a new resource with the given unique name, arguments, and options.
func NewGuestPolicy(ctx *pulumi.Context,
	name string, args *GuestPolicyArgs, opts ...pulumi.ResourceOption) (*GuestPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Assignment == nil {
		return nil, errors.New("invalid value for required argument 'Assignment'")
	}
	if args.GuestPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'GuestPolicyId'")
	}
	var resource GuestPolicy
	err := ctx.RegisterResource("google-native:osconfig/v1beta:GuestPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestPolicy gets an existing GuestPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestPolicyState, opts ...pulumi.ResourceOption) (*GuestPolicy, error) {
	var resource GuestPolicy
	err := ctx.ReadResource("google-native:osconfig/v1beta:GuestPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestPolicy resources.
type guestPolicyState struct {
}

type GuestPolicyState struct {
}

func (GuestPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestPolicyState)(nil)).Elem()
}

type guestPolicyArgs struct {
	// Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	Assignment Assignment `pulumi:"assignment"`
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag *string `pulumi:"etag"`
	// Required. The logical name of the guest policy in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
	GuestPolicyId string `pulumi:"guestPolicyId"`
	// Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
	Name *string `pulumi:"name"`
	// A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
	PackageRepositories []PackageRepository `pulumi:"packageRepositories"`
	// The software packages to be managed by this policy.
	Packages []Package `pulumi:"packages"`
	Project  *string   `pulumi:"project"`
	// A list of Recipes to install on the VM instance.
	Recipes []SoftwareRecipe `pulumi:"recipes"`
}

// The set of arguments for constructing a GuestPolicy resource.
type GuestPolicyArgs struct {
	// Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	Assignment AssignmentInput
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringPtrInput
	// Required. The logical name of the guest policy in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
	GuestPolicyId pulumi.StringInput
	// Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
	Name pulumi.StringPtrInput
	// A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
	PackageRepositories PackageRepositoryArrayInput
	// The software packages to be managed by this policy.
	Packages PackageArrayInput
	Project  pulumi.StringPtrInput
	// A list of Recipes to install on the VM instance.
	Recipes SoftwareRecipeArrayInput
}

func (GuestPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestPolicyArgs)(nil)).Elem()
}

type GuestPolicyInput interface {
	pulumi.Input

	ToGuestPolicyOutput() GuestPolicyOutput
	ToGuestPolicyOutputWithContext(ctx context.Context) GuestPolicyOutput
}

func (*GuestPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestPolicy)(nil)).Elem()
}

func (i *GuestPolicy) ToGuestPolicyOutput() GuestPolicyOutput {
	return i.ToGuestPolicyOutputWithContext(context.Background())
}

func (i *GuestPolicy) ToGuestPolicyOutputWithContext(ctx context.Context) GuestPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestPolicyOutput)
}

type GuestPolicyOutput struct{ *pulumi.OutputState }

func (GuestPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestPolicy)(nil)).Elem()
}

func (o GuestPolicyOutput) ToGuestPolicyOutput() GuestPolicyOutput {
	return o
}

func (o GuestPolicyOutput) ToGuestPolicyOutputWithContext(ctx context.Context) GuestPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GuestPolicyInput)(nil)).Elem(), &GuestPolicy{})
	pulumi.RegisterOutputType(GuestPolicyOutput{})
}
