// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Policy. Returns a `google.rpc.Status` with `google.rpc.Code.NOT_FOUND` if the constraint does not exist. Returns a `google.rpc.Status` with `google.rpc.Code.ALREADY_EXISTS` if the policy already exists on the given Cloud resource.
type FolderPolicy struct {
	pulumi.CustomResourceState

	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
	Name pulumi.StringOutput `pulumi:"name"`
	// Basic information about the Organization Policy.
	Spec GoogleCloudOrgpolicyV2PolicySpecResponseOutput `pulumi:"spec"`
}

// NewFolderPolicy registers a new resource with the given unique name, arguments, and options.
func NewFolderPolicy(ctx *pulumi.Context,
	name string, args *FolderPolicyArgs, opts ...pulumi.ResourceOption) (*FolderPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	var resource FolderPolicy
	err := ctx.RegisterResource("google-native:orgpolicy/v2:FolderPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderPolicy gets an existing FolderPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderPolicyState, opts ...pulumi.ResourceOption) (*FolderPolicy, error) {
	var resource FolderPolicy
	err := ctx.ReadResource("google-native:orgpolicy/v2:FolderPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderPolicy resources.
type folderPolicyState struct {
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
	Name *string `pulumi:"name"`
	// Basic information about the Organization Policy.
	Spec *GoogleCloudOrgpolicyV2PolicySpecResponse `pulumi:"spec"`
}

type FolderPolicyState struct {
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
	Name pulumi.StringPtrInput
	// Basic information about the Organization Policy.
	Spec GoogleCloudOrgpolicyV2PolicySpecResponsePtrInput
}

func (FolderPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderPolicyState)(nil)).Elem()
}

type folderPolicyArgs struct {
	FolderId string `pulumi:"folderId"`
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
	Name     *string `pulumi:"name"`
	PolicyId string  `pulumi:"policyId"`
	// Basic information about the Organization Policy.
	Spec *GoogleCloudOrgpolicyV2PolicySpec `pulumi:"spec"`
}

// The set of arguments for constructing a FolderPolicy resource.
type FolderPolicyArgs struct {
	FolderId pulumi.StringInput
	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
	Name     pulumi.StringPtrInput
	PolicyId pulumi.StringInput
	// Basic information about the Organization Policy.
	Spec GoogleCloudOrgpolicyV2PolicySpecPtrInput
}

func (FolderPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderPolicyArgs)(nil)).Elem()
}

type FolderPolicyInput interface {
	pulumi.Input

	ToFolderPolicyOutput() FolderPolicyOutput
	ToFolderPolicyOutputWithContext(ctx context.Context) FolderPolicyOutput
}

func (*FolderPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderPolicy)(nil))
}

func (i *FolderPolicy) ToFolderPolicyOutput() FolderPolicyOutput {
	return i.ToFolderPolicyOutputWithContext(context.Background())
}

func (i *FolderPolicy) ToFolderPolicyOutputWithContext(ctx context.Context) FolderPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderPolicyOutput)
}

type FolderPolicyOutput struct {
	*pulumi.OutputState
}

func (FolderPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderPolicy)(nil))
}

func (o FolderPolicyOutput) ToFolderPolicyOutput() FolderPolicyOutput {
	return o
}

func (o FolderPolicyOutput) ToFolderPolicyOutputWithContext(ctx context.Context) FolderPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FolderPolicyOutput{})
}
