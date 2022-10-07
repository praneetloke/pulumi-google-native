// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.
type ServiceAttachmentIamBinding struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewServiceAttachmentIamBinding registers a new resource with the given unique name, arguments, and options.
func NewServiceAttachmentIamBinding(ctx *pulumi.Context,
	name string, args *ServiceAttachmentIamBindingArgs, opts ...pulumi.ResourceOption) (*ServiceAttachmentIamBinding, error) {
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
	var resource ServiceAttachmentIamBinding
	err := ctx.RegisterResource("google-native:compute/beta:ServiceAttachmentIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAttachmentIamBinding gets an existing ServiceAttachmentIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAttachmentIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAttachmentIamBindingState, opts ...pulumi.ResourceOption) (*ServiceAttachmentIamBinding, error) {
	var resource ServiceAttachmentIamBinding
	err := ctx.ReadResource("google-native:compute/beta:ServiceAttachmentIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAttachmentIamBinding resources.
type serviceAttachmentIamBindingState struct {
}

type ServiceAttachmentIamBindingState struct {
}

func (ServiceAttachmentIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAttachmentIamBindingState)(nil)).Elem()
}

type serviceAttachmentIamBindingArgs struct {
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

// The set of arguments for constructing a ServiceAttachmentIamBinding resource.
type ServiceAttachmentIamBindingArgs struct {
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

func (ServiceAttachmentIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAttachmentIamBindingArgs)(nil)).Elem()
}

type ServiceAttachmentIamBindingInput interface {
	pulumi.Input

	ToServiceAttachmentIamBindingOutput() ServiceAttachmentIamBindingOutput
	ToServiceAttachmentIamBindingOutputWithContext(ctx context.Context) ServiceAttachmentIamBindingOutput
}

func (*ServiceAttachmentIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAttachmentIamBinding)(nil)).Elem()
}

func (i *ServiceAttachmentIamBinding) ToServiceAttachmentIamBindingOutput() ServiceAttachmentIamBindingOutput {
	return i.ToServiceAttachmentIamBindingOutputWithContext(context.Background())
}

func (i *ServiceAttachmentIamBinding) ToServiceAttachmentIamBindingOutputWithContext(ctx context.Context) ServiceAttachmentIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAttachmentIamBindingOutput)
}

type ServiceAttachmentIamBindingOutput struct{ *pulumi.OutputState }

func (ServiceAttachmentIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAttachmentIamBinding)(nil)).Elem()
}

func (o ServiceAttachmentIamBindingOutput) ToServiceAttachmentIamBindingOutput() ServiceAttachmentIamBindingOutput {
	return o
}

func (o ServiceAttachmentIamBindingOutput) ToServiceAttachmentIamBindingOutputWithContext(ctx context.Context) ServiceAttachmentIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o ServiceAttachmentIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *ServiceAttachmentIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o ServiceAttachmentIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAttachmentIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`.
func (o ServiceAttachmentIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceAttachmentIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o ServiceAttachmentIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAttachmentIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o ServiceAttachmentIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAttachmentIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
func (o ServiceAttachmentIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAttachmentIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAttachmentIamBindingInput)(nil)).Elem(), &ServiceAttachmentIamBinding{})
	pulumi.RegisterOutputType(ServiceAttachmentIamBindingOutput{})
}
