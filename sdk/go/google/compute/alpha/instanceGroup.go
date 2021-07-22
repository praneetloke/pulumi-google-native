// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an instance group in the specified project using the parameters that are included in the request.
type InstanceGroup struct {
	pulumi.CustomResourceState

	// The creation timestamp for this instance group in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The fingerprint of the named ports. The system uses this fingerprint to detect conflicts when multiple users change the named ports concurrently.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The resource type, which is always compute#instanceGroup for instance groups.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name pulumi.StringOutput `pulumi:"name"`
	//  Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "http", port: 80},{name: "http", port: 8080}] Named ports apply to all instances in this instance group.
	NamedPorts NamedPortResponseArrayOutput `pulumi:"namedPorts"`
	// The URL of the network to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
	Network pulumi.StringOutput `pulumi:"network"`
	// The URL of the region where the instance group is located (for regional resources).
	Region pulumi.StringOutput `pulumi:"region"`
	// The URL for this instance group. The server generates this URL.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// The total number of instances in the instance group.
	Size pulumi.IntOutput `pulumi:"size"`
	// The URL of the subnetwork to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The URL of the zone where the instance group is located (for zonal resources).
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroup(ctx *pulumi.Context,
	name string, args *InstanceGroupArgs, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	if args == nil {
		args = &InstanceGroupArgs{}
	}

	var resource InstanceGroup
	err := ctx.RegisterResource("google-native:compute/alpha:InstanceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroup gets an existing InstanceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupState, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	var resource InstanceGroup
	err := ctx.ReadResource("google-native:compute/alpha:InstanceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroup resources.
type instanceGroupState struct {
}

type InstanceGroupState struct {
}

func (InstanceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupState)(nil)).Elem()
}

type instanceGroupArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name *string `pulumi:"name"`
	//  Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "http", port: 80},{name: "http", port: 8080}] Named ports apply to all instances in this instance group.
	NamedPorts []NamedPort `pulumi:"namedPorts"`
	Project    *string     `pulumi:"project"`
	RequestId  *string     `pulumi:"requestId"`
	Zone       *string     `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroup resource.
type InstanceGroupArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name pulumi.StringPtrInput
	//  Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "http", port: 80},{name: "http", port: 8080}] Named ports apply to all instances in this instance group.
	NamedPorts NamedPortArrayInput
	Project    pulumi.StringPtrInput
	RequestId  pulumi.StringPtrInput
	Zone       pulumi.StringPtrInput
}

func (InstanceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupArgs)(nil)).Elem()
}

type InstanceGroupInput interface {
	pulumi.Input

	ToInstanceGroupOutput() InstanceGroupOutput
	ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput
}

func (*InstanceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroup)(nil))
}

func (i *InstanceGroup) ToInstanceGroupOutput() InstanceGroupOutput {
	return i.ToInstanceGroupOutputWithContext(context.Background())
}

func (i *InstanceGroup) ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupOutput)
}

type InstanceGroupOutput struct {
	*pulumi.OutputState
}

func (InstanceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroup)(nil))
}

func (o InstanceGroupOutput) ToInstanceGroupOutput() InstanceGroupOutput {
	return o
}

func (o InstanceGroupOutput) ToInstanceGroupOutputWithContext(ctx context.Context) InstanceGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceGroupOutput{})
}
