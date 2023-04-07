// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type CertificateAuthorityCertificateRevocationListIamBinding struct {
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

// NewCertificateAuthorityCertificateRevocationListIamBinding registers a new resource with the given unique name, arguments, and options.
func NewCertificateAuthorityCertificateRevocationListIamBinding(ctx *pulumi.Context,
	name string, args *CertificateAuthorityCertificateRevocationListIamBindingArgs, opts ...pulumi.ResourceOption) (*CertificateAuthorityCertificateRevocationListIamBinding, error) {
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
	var resource CertificateAuthorityCertificateRevocationListIamBinding
	err := ctx.RegisterResource("google-native:privateca/v1beta1:CertificateAuthorityCertificateRevocationListIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateAuthorityCertificateRevocationListIamBinding gets an existing CertificateAuthorityCertificateRevocationListIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateAuthorityCertificateRevocationListIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateAuthorityCertificateRevocationListIamBindingState, opts ...pulumi.ResourceOption) (*CertificateAuthorityCertificateRevocationListIamBinding, error) {
	var resource CertificateAuthorityCertificateRevocationListIamBinding
	err := ctx.ReadResource("google-native:privateca/v1beta1:CertificateAuthorityCertificateRevocationListIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateAuthorityCertificateRevocationListIamBinding resources.
type certificateAuthorityCertificateRevocationListIamBindingState struct {
}

type CertificateAuthorityCertificateRevocationListIamBindingState struct {
}

func (CertificateAuthorityCertificateRevocationListIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityCertificateRevocationListIamBindingState)(nil)).Elem()
}

type certificateAuthorityCertificateRevocationListIamBindingArgs struct {
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

// The set of arguments for constructing a CertificateAuthorityCertificateRevocationListIamBinding resource.
type CertificateAuthorityCertificateRevocationListIamBindingArgs struct {
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

func (CertificateAuthorityCertificateRevocationListIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityCertificateRevocationListIamBindingArgs)(nil)).Elem()
}

type CertificateAuthorityCertificateRevocationListIamBindingInput interface {
	pulumi.Input

	ToCertificateAuthorityCertificateRevocationListIamBindingOutput() CertificateAuthorityCertificateRevocationListIamBindingOutput
	ToCertificateAuthorityCertificateRevocationListIamBindingOutputWithContext(ctx context.Context) CertificateAuthorityCertificateRevocationListIamBindingOutput
}

func (*CertificateAuthorityCertificateRevocationListIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthorityCertificateRevocationListIamBinding)(nil)).Elem()
}

func (i *CertificateAuthorityCertificateRevocationListIamBinding) ToCertificateAuthorityCertificateRevocationListIamBindingOutput() CertificateAuthorityCertificateRevocationListIamBindingOutput {
	return i.ToCertificateAuthorityCertificateRevocationListIamBindingOutputWithContext(context.Background())
}

func (i *CertificateAuthorityCertificateRevocationListIamBinding) ToCertificateAuthorityCertificateRevocationListIamBindingOutputWithContext(ctx context.Context) CertificateAuthorityCertificateRevocationListIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificateRevocationListIamBindingOutput)
}

type CertificateAuthorityCertificateRevocationListIamBindingOutput struct{ *pulumi.OutputState }

func (CertificateAuthorityCertificateRevocationListIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthorityCertificateRevocationListIamBinding)(nil)).Elem()
}

func (o CertificateAuthorityCertificateRevocationListIamBindingOutput) ToCertificateAuthorityCertificateRevocationListIamBindingOutput() CertificateAuthorityCertificateRevocationListIamBindingOutput {
	return o
}

func (o CertificateAuthorityCertificateRevocationListIamBindingOutput) ToCertificateAuthorityCertificateRevocationListIamBindingOutputWithContext(ctx context.Context) CertificateAuthorityCertificateRevocationListIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o CertificateAuthorityCertificateRevocationListIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *CertificateAuthorityCertificateRevocationListIamBinding) iam.ConditionPtrOutput {
		return v.Condition
	}).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o CertificateAuthorityCertificateRevocationListIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthorityCertificateRevocationListIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
func (o CertificateAuthorityCertificateRevocationListIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertificateAuthorityCertificateRevocationListIamBinding) pulumi.StringArrayOutput {
		return v.Members
	}).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o CertificateAuthorityCertificateRevocationListIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthorityCertificateRevocationListIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o CertificateAuthorityCertificateRevocationListIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthorityCertificateRevocationListIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
func (o CertificateAuthorityCertificateRevocationListIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateAuthorityCertificateRevocationListIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateAuthorityCertificateRevocationListIamBindingInput)(nil)).Elem(), &CertificateAuthorityCertificateRevocationListIamBinding{})
	pulumi.RegisterOutputType(CertificateAuthorityCertificateRevocationListIamBindingOutput{})
}
