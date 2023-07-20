// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Peering for Managed AD instance.
// Auto-naming is currently not supported for this resource.
type Peering struct {
	pulumi.CustomResourceState

	// The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
	AuthorizedNetwork pulumi.StringOutput `pulumi:"authorizedNetwork"`
	// The time the instance was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
	DomainResource pulumi.StringOutput `pulumi:"domainResource"`
	// Optional. Resource labels to represent user provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Unique name of the peering in this scope including projects and location using the form: `projects/{project_id}/locations/global/peerings/{peering_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. Peering Id, unique name to identify peering.
	PeeringId pulumi.StringOutput `pulumi:"peeringId"`
	Project   pulumi.StringOutput `pulumi:"project"`
	// The current state of this Peering.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current status of this peering, if available.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Last update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewPeering registers a new resource with the given unique name, arguments, and options.
func NewPeering(ctx *pulumi.Context,
	name string, args *PeeringArgs, opts ...pulumi.ResourceOption) (*Peering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizedNetwork == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizedNetwork'")
	}
	if args.DomainResource == nil {
		return nil, errors.New("invalid value for required argument 'DomainResource'")
	}
	if args.PeeringId == nil {
		return nil, errors.New("invalid value for required argument 'PeeringId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"peeringId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Peering
	err := ctx.RegisterResource("google-native:managedidentities/v1alpha1:Peering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeering gets an existing Peering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringState, opts ...pulumi.ResourceOption) (*Peering, error) {
	var resource Peering
	err := ctx.ReadResource("google-native:managedidentities/v1alpha1:Peering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Peering resources.
type peeringState struct {
}

type PeeringState struct {
}

func (PeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringState)(nil)).Elem()
}

type peeringArgs struct {
	// The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
	AuthorizedNetwork string `pulumi:"authorizedNetwork"`
	// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
	DomainResource string `pulumi:"domainResource"`
	// Optional. Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Required. Peering Id, unique name to identify peering.
	PeeringId string  `pulumi:"peeringId"`
	Project   *string `pulumi:"project"`
}

// The set of arguments for constructing a Peering resource.
type PeeringArgs struct {
	// The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
	AuthorizedNetwork pulumi.StringInput
	// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
	DomainResource pulumi.StringInput
	// Optional. Resource labels to represent user provided metadata.
	Labels pulumi.StringMapInput
	// Required. Peering Id, unique name to identify peering.
	PeeringId pulumi.StringInput
	Project   pulumi.StringPtrInput
}

func (PeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringArgs)(nil)).Elem()
}

type PeeringInput interface {
	pulumi.Input

	ToPeeringOutput() PeeringOutput
	ToPeeringOutputWithContext(ctx context.Context) PeeringOutput
}

func (*Peering) ElementType() reflect.Type {
	return reflect.TypeOf((**Peering)(nil)).Elem()
}

func (i *Peering) ToPeeringOutput() PeeringOutput {
	return i.ToPeeringOutputWithContext(context.Background())
}

func (i *Peering) ToPeeringOutputWithContext(ctx context.Context) PeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringOutput)
}

type PeeringOutput struct{ *pulumi.OutputState }

func (PeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Peering)(nil)).Elem()
}

func (o PeeringOutput) ToPeeringOutput() PeeringOutput {
	return o
}

func (o PeeringOutput) ToPeeringOutputWithContext(ctx context.Context) PeeringOutput {
	return o
}

// The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
func (o PeeringOutput) AuthorizedNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.AuthorizedNetwork }).(pulumi.StringOutput)
}

// The time the instance was created.
func (o PeeringOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
func (o PeeringOutput) DomainResource() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.DomainResource }).(pulumi.StringOutput)
}

// Optional. Resource labels to represent user provided metadata.
func (o PeeringOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Unique name of the peering in this scope including projects and location using the form: `projects/{project_id}/locations/global/peerings/{peering_id}`.
func (o PeeringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Required. Peering Id, unique name to identify peering.
func (o PeeringOutput) PeeringId() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.PeeringId }).(pulumi.StringOutput)
}

func (o PeeringOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The current state of this Peering.
func (o PeeringOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current status of this peering, if available.
func (o PeeringOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

// Last update time.
func (o PeeringOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringInput)(nil)).Elem(), &Peering{})
	pulumi.RegisterOutputType(PeeringOutput{})
}
