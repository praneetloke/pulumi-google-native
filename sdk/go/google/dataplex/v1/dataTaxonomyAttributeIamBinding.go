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

// Sets the access control policy on the specified resource. Replaces any existing policy.Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED errors.
type DataTaxonomyAttributeIamBinding struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identities that will be granted the privilege in role. Each entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one `IamBinding` can be used per role.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDataTaxonomyAttributeIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDataTaxonomyAttributeIamBinding(ctx *pulumi.Context,
	name string, args *DataTaxonomyAttributeIamBindingArgs, opts ...pulumi.ResourceOption) (*DataTaxonomyAttributeIamBinding, error) {
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
	var resource DataTaxonomyAttributeIamBinding
	err := ctx.RegisterResource("google-native:dataplex/v1:DataTaxonomyAttributeIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataTaxonomyAttributeIamBinding gets an existing DataTaxonomyAttributeIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataTaxonomyAttributeIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataTaxonomyAttributeIamBindingState, opts ...pulumi.ResourceOption) (*DataTaxonomyAttributeIamBinding, error) {
	var resource DataTaxonomyAttributeIamBinding
	err := ctx.ReadResource("google-native:dataplex/v1:DataTaxonomyAttributeIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataTaxonomyAttributeIamBinding resources.
type dataTaxonomyAttributeIamBindingState struct {
}

type DataTaxonomyAttributeIamBindingState struct {
}

func (DataTaxonomyAttributeIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataTaxonomyAttributeIamBindingState)(nil)).Elem()
}

type dataTaxonomyAttributeIamBindingArgs struct {
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

// The set of arguments for constructing a DataTaxonomyAttributeIamBinding resource.
type DataTaxonomyAttributeIamBindingArgs struct {
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

func (DataTaxonomyAttributeIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataTaxonomyAttributeIamBindingArgs)(nil)).Elem()
}

type DataTaxonomyAttributeIamBindingInput interface {
	pulumi.Input

	ToDataTaxonomyAttributeIamBindingOutput() DataTaxonomyAttributeIamBindingOutput
	ToDataTaxonomyAttributeIamBindingOutputWithContext(ctx context.Context) DataTaxonomyAttributeIamBindingOutput
}

func (*DataTaxonomyAttributeIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataTaxonomyAttributeIamBinding)(nil)).Elem()
}

func (i *DataTaxonomyAttributeIamBinding) ToDataTaxonomyAttributeIamBindingOutput() DataTaxonomyAttributeIamBindingOutput {
	return i.ToDataTaxonomyAttributeIamBindingOutputWithContext(context.Background())
}

func (i *DataTaxonomyAttributeIamBinding) ToDataTaxonomyAttributeIamBindingOutputWithContext(ctx context.Context) DataTaxonomyAttributeIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataTaxonomyAttributeIamBindingOutput)
}

type DataTaxonomyAttributeIamBindingOutput struct{ *pulumi.OutputState }

func (DataTaxonomyAttributeIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataTaxonomyAttributeIamBinding)(nil)).Elem()
}

func (o DataTaxonomyAttributeIamBindingOutput) ToDataTaxonomyAttributeIamBindingOutput() DataTaxonomyAttributeIamBindingOutput {
	return o
}

func (o DataTaxonomyAttributeIamBindingOutput) ToDataTaxonomyAttributeIamBindingOutputWithContext(ctx context.Context) DataTaxonomyAttributeIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o DataTaxonomyAttributeIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *DataTaxonomyAttributeIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o DataTaxonomyAttributeIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataTaxonomyAttributeIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o DataTaxonomyAttributeIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataTaxonomyAttributeIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o DataTaxonomyAttributeIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataTaxonomyAttributeIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o DataTaxonomyAttributeIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataTaxonomyAttributeIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one `IamBinding` can be used per role.
func (o DataTaxonomyAttributeIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataTaxonomyAttributeIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataTaxonomyAttributeIamBindingInput)(nil)).Elem(), &DataTaxonomyAttributeIamBinding{})
	pulumi.RegisterOutputType(DataTaxonomyAttributeIamBindingOutput{})
}
