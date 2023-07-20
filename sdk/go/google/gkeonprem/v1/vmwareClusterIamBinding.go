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
type VmwareClusterIamBinding struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewVmwareClusterIamBinding registers a new resource with the given unique name, arguments, and options.
func NewVmwareClusterIamBinding(ctx *pulumi.Context,
	name string, args *VmwareClusterIamBindingArgs, opts ...pulumi.ResourceOption) (*VmwareClusterIamBinding, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VmwareClusterIamBinding
	err := ctx.RegisterResource("google-native:gkeonprem/v1:VmwareClusterIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmwareClusterIamBinding gets an existing VmwareClusterIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmwareClusterIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmwareClusterIamBindingState, opts ...pulumi.ResourceOption) (*VmwareClusterIamBinding, error) {
	var resource VmwareClusterIamBinding
	err := ctx.ReadResource("google-native:gkeonprem/v1:VmwareClusterIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmwareClusterIamBinding resources.
type vmwareClusterIamBindingState struct {
}

type VmwareClusterIamBindingState struct {
}

func (VmwareClusterIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareClusterIamBindingState)(nil)).Elem()
}

type vmwareClusterIamBindingArgs struct {
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

// The set of arguments for constructing a VmwareClusterIamBinding resource.
type VmwareClusterIamBindingArgs struct {
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

func (VmwareClusterIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareClusterIamBindingArgs)(nil)).Elem()
}

type VmwareClusterIamBindingInput interface {
	pulumi.Input

	ToVmwareClusterIamBindingOutput() VmwareClusterIamBindingOutput
	ToVmwareClusterIamBindingOutputWithContext(ctx context.Context) VmwareClusterIamBindingOutput
}

func (*VmwareClusterIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**VmwareClusterIamBinding)(nil)).Elem()
}

func (i *VmwareClusterIamBinding) ToVmwareClusterIamBindingOutput() VmwareClusterIamBindingOutput {
	return i.ToVmwareClusterIamBindingOutputWithContext(context.Background())
}

func (i *VmwareClusterIamBinding) ToVmwareClusterIamBindingOutputWithContext(ctx context.Context) VmwareClusterIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmwareClusterIamBindingOutput)
}

type VmwareClusterIamBindingOutput struct{ *pulumi.OutputState }

func (VmwareClusterIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmwareClusterIamBinding)(nil)).Elem()
}

func (o VmwareClusterIamBindingOutput) ToVmwareClusterIamBindingOutput() VmwareClusterIamBindingOutput {
	return o
}

func (o VmwareClusterIamBindingOutput) ToVmwareClusterIamBindingOutputWithContext(ctx context.Context) VmwareClusterIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o VmwareClusterIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *VmwareClusterIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o VmwareClusterIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareClusterIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
func (o VmwareClusterIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VmwareClusterIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o VmwareClusterIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareClusterIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o VmwareClusterIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareClusterIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
func (o VmwareClusterIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareClusterIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmwareClusterIamBindingInput)(nil)).Elem(), &VmwareClusterIamBinding{})
	pulumi.RegisterOutputType(VmwareClusterIamBindingOutput{})
}
