// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a Lien which applies to the resource denoted by the `parent` field. Callers of this method will require permission on the `parent` resource. For example, applying to `projects/1234` requires permission `resourcemanager.projects.updateLiens`. NOTE: Some resources may limit the number of Liens which may be applied.
type Lien struct {
	pulumi.CustomResourceState

	// The creation time of this Lien.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
	Name pulumi.StringOutput `pulumi:"name"`
	// A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
	Origin pulumi.StringOutput `pulumi:"origin"`
	// A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
	Reason pulumi.StringOutput `pulumi:"reason"`
	// The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
	Restrictions pulumi.StringArrayOutput `pulumi:"restrictions"`
}

// NewLien registers a new resource with the given unique name, arguments, and options.
func NewLien(ctx *pulumi.Context,
	name string, args *LienArgs, opts ...pulumi.ResourceOption) (*Lien, error) {
	if args == nil {
		args = &LienArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Lien
	err := ctx.RegisterResource("google-native:cloudresourcemanager/v1:Lien", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLien gets an existing Lien resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLien(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LienState, opts ...pulumi.ResourceOption) (*Lien, error) {
	var resource Lien
	err := ctx.ReadResource("google-native:cloudresourcemanager/v1:Lien", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lien resources.
type lienState struct {
}

type LienState struct {
}

func (LienState) ElementType() reflect.Type {
	return reflect.TypeOf((*lienState)(nil)).Elem()
}

type lienArgs struct {
	// The creation time of this Lien.
	CreateTime *string `pulumi:"createTime"`
	// A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
	Name *string `pulumi:"name"`
	// A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
	Origin *string `pulumi:"origin"`
	// A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
	Parent *string `pulumi:"parent"`
	// Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
	Reason *string `pulumi:"reason"`
	// The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
	Restrictions []string `pulumi:"restrictions"`
}

// The set of arguments for constructing a Lien resource.
type LienArgs struct {
	// The creation time of this Lien.
	CreateTime pulumi.StringPtrInput
	// A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
	Name pulumi.StringPtrInput
	// A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
	Origin pulumi.StringPtrInput
	// A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
	Parent pulumi.StringPtrInput
	// Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
	Reason pulumi.StringPtrInput
	// The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
	Restrictions pulumi.StringArrayInput
}

func (LienArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lienArgs)(nil)).Elem()
}

type LienInput interface {
	pulumi.Input

	ToLienOutput() LienOutput
	ToLienOutputWithContext(ctx context.Context) LienOutput
}

func (*Lien) ElementType() reflect.Type {
	return reflect.TypeOf((**Lien)(nil)).Elem()
}

func (i *Lien) ToLienOutput() LienOutput {
	return i.ToLienOutputWithContext(context.Background())
}

func (i *Lien) ToLienOutputWithContext(ctx context.Context) LienOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LienOutput)
}

type LienOutput struct{ *pulumi.OutputState }

func (LienOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lien)(nil)).Elem()
}

func (o LienOutput) ToLienOutput() LienOutput {
	return o
}

func (o LienOutput) ToLienOutputWithContext(ctx context.Context) LienOutput {
	return o
}

// The creation time of this Lien.
func (o LienOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A system-generated unique identifier for this Lien. Example: `liens/1234abcd`
func (o LienOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters. Example: 'compute.googleapis.com'
func (o LienOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Example: `projects/1234`
func (o LienOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters. Example: 'Holds production API key'
func (o LienOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringOutput { return v.Reason }).(pulumi.StringOutput)
}

// The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. Example: ['resourcemanager.projects.delete']
func (o LienOutput) Restrictions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Lien) pulumi.StringArrayOutput { return v.Restrictions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LienInput)(nil)).Elem(), &Lien{})
	pulumi.RegisterOutputType(LienOutput{})
}
