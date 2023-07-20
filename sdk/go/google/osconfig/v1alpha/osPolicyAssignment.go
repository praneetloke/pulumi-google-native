// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create an OS policy assignment. This method also creates the first revision of the OS policy assignment. This method returns a long running operation (LRO) that contains the rollout details. The rollout can be cancelled by cancelling the LRO. For more information, see [Method: projects.locations.osPolicyAssignments.operations.cancel](https://cloud.google.com/compute/docs/osconfig/rest/v1alpha/projects.locations.osPolicyAssignments.operations/cancel).
type OsPolicyAssignment struct {
	pulumi.CustomResourceState

	// Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS policies from this revision. For a given OS policy assignment, there is only one revision with a value of `true` for this field.
	Baseline pulumi.BoolOutput `pulumi:"baseline"`
	// Indicates that this revision deletes the OS policy assignment.
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Filter to select VMs.
	InstanceFilter OSPolicyAssignmentInstanceFilterResponseOutput `pulumi:"instanceFilter"`
	Location       pulumi.StringOutput                            `pulumi:"location"`
	// Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of OS policies to be applied to the VMs.
	OsPolicies OSPolicyResponseArrayOutput `pulumi:"osPolicies"`
	// Required. The logical name of the OS policy assignment in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
	OsPolicyAssignmentId pulumi.StringOutput `pulumi:"osPolicyAssignmentId"`
	Project              pulumi.StringOutput `pulumi:"project"`
	// Indicates that reconciliation is in progress for the revision. This value is `true` when the `rollout_state` is one of: * IN_PROGRESS * CANCELLING
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// The timestamp that the revision was created.
	RevisionCreateTime pulumi.StringOutput `pulumi:"revisionCreateTime"`
	// The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy assignment
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout OSPolicyAssignmentRolloutResponseOutput `pulumi:"rollout"`
	// OS policy assignment rollout state
	RolloutState pulumi.StringOutput `pulumi:"rolloutState"`
	// Server generated unique id for the OS policy assignment resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewOsPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewOsPolicyAssignment(ctx *pulumi.Context,
	name string, args *OsPolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*OsPolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceFilter == nil {
		return nil, errors.New("invalid value for required argument 'InstanceFilter'")
	}
	if args.OsPolicies == nil {
		return nil, errors.New("invalid value for required argument 'OsPolicies'")
	}
	if args.OsPolicyAssignmentId == nil {
		return nil, errors.New("invalid value for required argument 'OsPolicyAssignmentId'")
	}
	if args.Rollout == nil {
		return nil, errors.New("invalid value for required argument 'Rollout'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"osPolicyAssignmentId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OsPolicyAssignment
	err := ctx.RegisterResource("google-native:osconfig/v1alpha:OsPolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOsPolicyAssignment gets an existing OsPolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOsPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OsPolicyAssignmentState, opts ...pulumi.ResourceOption) (*OsPolicyAssignment, error) {
	var resource OsPolicyAssignment
	err := ctx.ReadResource("google-native:osconfig/v1alpha:OsPolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OsPolicyAssignment resources.
type osPolicyAssignmentState struct {
}

type OsPolicyAssignmentState struct {
}

func (OsPolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*osPolicyAssignmentState)(nil)).Elem()
}

type osPolicyAssignmentArgs struct {
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	Etag *string `pulumi:"etag"`
	// Filter to select VMs.
	InstanceFilter OSPolicyAssignmentInstanceFilter `pulumi:"instanceFilter"`
	Location       *string                          `pulumi:"location"`
	// Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
	Name *string `pulumi:"name"`
	// List of OS policies to be applied to the VMs.
	OsPolicies []OSPolicy `pulumi:"osPolicies"`
	// Required. The logical name of the OS policy assignment in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
	OsPolicyAssignmentId string  `pulumi:"osPolicyAssignmentId"`
	Project              *string `pulumi:"project"`
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout OSPolicyAssignmentRollout `pulumi:"rollout"`
}

// The set of arguments for constructing a OsPolicyAssignment resource.
type OsPolicyAssignmentArgs struct {
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	Etag pulumi.StringPtrInput
	// Filter to select VMs.
	InstanceFilter OSPolicyAssignmentInstanceFilterInput
	Location       pulumi.StringPtrInput
	// Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
	Name pulumi.StringPtrInput
	// List of OS policies to be applied to the VMs.
	OsPolicies OSPolicyArrayInput
	// Required. The logical name of the OS policy assignment in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
	OsPolicyAssignmentId pulumi.StringInput
	Project              pulumi.StringPtrInput
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout OSPolicyAssignmentRolloutInput
}

func (OsPolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*osPolicyAssignmentArgs)(nil)).Elem()
}

type OsPolicyAssignmentInput interface {
	pulumi.Input

	ToOsPolicyAssignmentOutput() OsPolicyAssignmentOutput
	ToOsPolicyAssignmentOutputWithContext(ctx context.Context) OsPolicyAssignmentOutput
}

func (*OsPolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**OsPolicyAssignment)(nil)).Elem()
}

func (i *OsPolicyAssignment) ToOsPolicyAssignmentOutput() OsPolicyAssignmentOutput {
	return i.ToOsPolicyAssignmentOutputWithContext(context.Background())
}

func (i *OsPolicyAssignment) ToOsPolicyAssignmentOutputWithContext(ctx context.Context) OsPolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsPolicyAssignmentOutput)
}

type OsPolicyAssignmentOutput struct{ *pulumi.OutputState }

func (OsPolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsPolicyAssignment)(nil)).Elem()
}

func (o OsPolicyAssignmentOutput) ToOsPolicyAssignmentOutput() OsPolicyAssignmentOutput {
	return o
}

func (o OsPolicyAssignmentOutput) ToOsPolicyAssignmentOutputWithContext(ctx context.Context) OsPolicyAssignmentOutput {
	return o
}

// Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS policies from this revision. For a given OS policy assignment, there is only one revision with a value of `true` for this field.
func (o OsPolicyAssignmentOutput) Baseline() pulumi.BoolOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.BoolOutput { return v.Baseline }).(pulumi.BoolOutput)
}

// Indicates that this revision deletes the OS policy assignment.
func (o OsPolicyAssignmentOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.BoolOutput { return v.Deleted }).(pulumi.BoolOutput)
}

// OS policy assignment description. Length of the description is limited to 1024 characters.
func (o OsPolicyAssignmentOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
func (o OsPolicyAssignmentOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Filter to select VMs.
func (o OsPolicyAssignmentOutput) InstanceFilter() OSPolicyAssignmentInstanceFilterResponseOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) OSPolicyAssignmentInstanceFilterResponseOutput { return v.InstanceFilter }).(OSPolicyAssignmentInstanceFilterResponseOutput)
}

func (o OsPolicyAssignmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
func (o OsPolicyAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of OS policies to be applied to the VMs.
func (o OsPolicyAssignmentOutput) OsPolicies() OSPolicyResponseArrayOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) OSPolicyResponseArrayOutput { return v.OsPolicies }).(OSPolicyResponseArrayOutput)
}

// Required. The logical name of the OS policy assignment in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
func (o OsPolicyAssignmentOutput) OsPolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.OsPolicyAssignmentId }).(pulumi.StringOutput)
}

func (o OsPolicyAssignmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Indicates that reconciliation is in progress for the revision. This value is `true` when the `rollout_state` is one of: * IN_PROGRESS * CANCELLING
func (o OsPolicyAssignmentOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// The timestamp that the revision was created.
func (o OsPolicyAssignmentOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy assignment
func (o OsPolicyAssignmentOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.RevisionId }).(pulumi.StringOutput)
}

// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
func (o OsPolicyAssignmentOutput) Rollout() OSPolicyAssignmentRolloutResponseOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) OSPolicyAssignmentRolloutResponseOutput { return v.Rollout }).(OSPolicyAssignmentRolloutResponseOutput)
}

// OS policy assignment rollout state
func (o OsPolicyAssignmentOutput) RolloutState() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.RolloutState }).(pulumi.StringOutput)
}

// Server generated unique id for the OS policy assignment resource.
func (o OsPolicyAssignmentOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *OsPolicyAssignment) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OsPolicyAssignmentInput)(nil)).Elem(), &OsPolicyAssignment{})
	pulumi.RegisterOutputType(OsPolicyAssignmentOutput{})
}
