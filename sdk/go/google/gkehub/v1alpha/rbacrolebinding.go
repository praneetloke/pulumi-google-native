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

// Creates a RBACRoleBinding.
// Auto-naming is currently not supported for this resource.
type Rbacrolebinding struct {
	pulumi.CustomResourceState

	// When the rbacrolebinding was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// When the rbacrolebinding was deleted.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// group is the group, as seen by the kubernetes cluster.
	Group    pulumi.StringOutput `pulumi:"group"`
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for the rbacrolebinding `projects/{project}/locations/{location}/namespaces/{namespace}/rbacrolebindings/{rbacrolebinding}` or `projects/{project}/locations/{location}/memberships/{membership}/rbacrolebindings/{rbacrolebinding}`
	Name        pulumi.StringOutput `pulumi:"name"`
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	Project     pulumi.StringOutput `pulumi:"project"`
	// Required. Client chosen ID for the RBACRoleBinding. `rbacrolebinding_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	RbacrolebindingId pulumi.StringOutput `pulumi:"rbacrolebindingId"`
	// Role to bind to the principal
	Role RoleResponseOutput `pulumi:"role"`
	// State of the rbacrolebinding resource.
	State RBACRoleBindingLifecycleStateResponseOutput `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all rbacrolebinding resources. If a rbacrolebinding resource is deleted and another resource with the same name is created, it gets a different uid.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// When the rbacrolebinding was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld"
	User pulumi.StringOutput `pulumi:"user"`
}

// NewRbacrolebinding registers a new resource with the given unique name, arguments, and options.
func NewRbacrolebinding(ctx *pulumi.Context,
	name string, args *RbacrolebindingArgs, opts ...pulumi.ResourceOption) (*Rbacrolebinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.RbacrolebindingId == nil {
		return nil, errors.New("invalid value for required argument 'RbacrolebindingId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"namespaceId",
		"project",
		"rbacrolebindingId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rbacrolebinding
	err := ctx.RegisterResource("google-native:gkehub/v1alpha:Rbacrolebinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRbacrolebinding gets an existing Rbacrolebinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRbacrolebinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RbacrolebindingState, opts ...pulumi.ResourceOption) (*Rbacrolebinding, error) {
	var resource Rbacrolebinding
	err := ctx.ReadResource("google-native:gkehub/v1alpha:Rbacrolebinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rbacrolebinding resources.
type rbacrolebindingState struct {
}

type RbacrolebindingState struct {
}

func (RbacrolebindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacrolebindingState)(nil)).Elem()
}

type rbacrolebindingArgs struct {
	// group is the group, as seen by the kubernetes cluster.
	Group    *string `pulumi:"group"`
	Location *string `pulumi:"location"`
	// The resource name for the rbacrolebinding `projects/{project}/locations/{location}/namespaces/{namespace}/rbacrolebindings/{rbacrolebinding}` or `projects/{project}/locations/{location}/memberships/{membership}/rbacrolebindings/{rbacrolebinding}`
	Name        *string `pulumi:"name"`
	NamespaceId string  `pulumi:"namespaceId"`
	Project     *string `pulumi:"project"`
	// Required. Client chosen ID for the RBACRoleBinding. `rbacrolebinding_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	RbacrolebindingId string `pulumi:"rbacrolebindingId"`
	// Role to bind to the principal
	Role Role `pulumi:"role"`
	// user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld"
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a Rbacrolebinding resource.
type RbacrolebindingArgs struct {
	// group is the group, as seen by the kubernetes cluster.
	Group    pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// The resource name for the rbacrolebinding `projects/{project}/locations/{location}/namespaces/{namespace}/rbacrolebindings/{rbacrolebinding}` or `projects/{project}/locations/{location}/memberships/{membership}/rbacrolebindings/{rbacrolebinding}`
	Name        pulumi.StringPtrInput
	NamespaceId pulumi.StringInput
	Project     pulumi.StringPtrInput
	// Required. Client chosen ID for the RBACRoleBinding. `rbacrolebinding_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	RbacrolebindingId pulumi.StringInput
	// Role to bind to the principal
	Role RoleInput
	// user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld"
	User pulumi.StringPtrInput
}

func (RbacrolebindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacrolebindingArgs)(nil)).Elem()
}

type RbacrolebindingInput interface {
	pulumi.Input

	ToRbacrolebindingOutput() RbacrolebindingOutput
	ToRbacrolebindingOutputWithContext(ctx context.Context) RbacrolebindingOutput
}

func (*Rbacrolebinding) ElementType() reflect.Type {
	return reflect.TypeOf((**Rbacrolebinding)(nil)).Elem()
}

func (i *Rbacrolebinding) ToRbacrolebindingOutput() RbacrolebindingOutput {
	return i.ToRbacrolebindingOutputWithContext(context.Background())
}

func (i *Rbacrolebinding) ToRbacrolebindingOutputWithContext(ctx context.Context) RbacrolebindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacrolebindingOutput)
}

type RbacrolebindingOutput struct{ *pulumi.OutputState }

func (RbacrolebindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rbacrolebinding)(nil)).Elem()
}

func (o RbacrolebindingOutput) ToRbacrolebindingOutput() RbacrolebindingOutput {
	return o
}

func (o RbacrolebindingOutput) ToRbacrolebindingOutputWithContext(ctx context.Context) RbacrolebindingOutput {
	return o
}

// When the rbacrolebinding was created.
func (o RbacrolebindingOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// When the rbacrolebinding was deleted.
func (o RbacrolebindingOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// group is the group, as seen by the kubernetes cluster.
func (o RbacrolebindingOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

func (o RbacrolebindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name for the rbacrolebinding `projects/{project}/locations/{location}/namespaces/{namespace}/rbacrolebindings/{rbacrolebinding}` or `projects/{project}/locations/{location}/memberships/{membership}/rbacrolebindings/{rbacrolebinding}`
func (o RbacrolebindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RbacrolebindingOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

func (o RbacrolebindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. Client chosen ID for the RBACRoleBinding. `rbacrolebinding_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
func (o RbacrolebindingOutput) RbacrolebindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.RbacrolebindingId }).(pulumi.StringOutput)
}

// Role to bind to the principal
func (o RbacrolebindingOutput) Role() RoleResponseOutput {
	return o.ApplyT(func(v *Rbacrolebinding) RoleResponseOutput { return v.Role }).(RoleResponseOutput)
}

// State of the rbacrolebinding resource.
func (o RbacrolebindingOutput) State() RBACRoleBindingLifecycleStateResponseOutput {
	return o.ApplyT(func(v *Rbacrolebinding) RBACRoleBindingLifecycleStateResponseOutput { return v.State }).(RBACRoleBindingLifecycleStateResponseOutput)
}

// Google-generated UUID for this resource. This is unique across all rbacrolebinding resources. If a rbacrolebinding resource is deleted and another resource with the same name is created, it gets a different uid.
func (o RbacrolebindingOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// When the rbacrolebinding was last updated.
func (o RbacrolebindingOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld"
func (o RbacrolebindingOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *Rbacrolebinding) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RbacrolebindingInput)(nil)).Elem(), &Rbacrolebinding{})
	pulumi.RegisterOutputType(RbacrolebindingOutput{})
}
