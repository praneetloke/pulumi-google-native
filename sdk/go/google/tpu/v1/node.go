// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a node.
type Node struct {
	pulumi.CustomResourceState

	// Required. The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringOutput `pulumi:"acceleratorType"`
	// The API version that created this Node.
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The time when the node was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The health status of the TPU node.
	Health pulumi.StringOutput `pulumi:"health"`
	// If this field is populated, it contains a description of why the TPU Node is unhealthy.
	HealthDescription pulumi.StringOutput `pulumi:"healthDescription"`
	// DEPRECATED! Use network_endpoints instead. The network address for the TPU Node as visible to Compute Engine instances.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Immutable. The name of the TPU
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
	Network pulumi.StringOutput `pulumi:"network"`
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node reach out to the 0th entry in this map first.
	NetworkEndpoints NetworkEndpointResponseArrayOutput `pulumi:"networkEndpoints"`
	// DEPRECATED! Use network_endpoints instead. The network port for the TPU Node as visible to Compute Engine instances.
	Port pulumi.StringOutput `pulumi:"port"`
	// The scheduling options for this node.
	SchedulingConfig SchedulingConfigResponseOutput `pulumi:"schedulingConfig"`
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The current state for the TPU Node.
	State pulumi.StringOutput `pulumi:"state"`
	// The Symptoms that have occurred to the TPU Node.
	Symptoms SymptomResponseArrayOutput `pulumi:"symptoms"`
	// Required. The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringOutput `pulumi:"tensorflowVersion"`
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
	UseServiceNetworking pulumi.BoolOutput `pulumi:"useServiceNetworking"`
}

// NewNode registers a new resource with the given unique name, arguments, and options.
func NewNode(ctx *pulumi.Context,
	name string, args *NodeArgs, opts ...pulumi.ResourceOption) (*Node, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocationsId == nil {
		return nil, errors.New("invalid value for required argument 'LocationsId'")
	}
	if args.NodesId == nil {
		return nil, errors.New("invalid value for required argument 'NodesId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource Node
	err := ctx.RegisterResource("google-cloud:tpu/v1:Node", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNode gets an existing Node resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeState, opts ...pulumi.ResourceOption) (*Node, error) {
	var resource Node
	err := ctx.ReadResource("google-cloud:tpu/v1:Node", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Node resources.
type nodeState struct {
	// Required. The type of hardware accelerators associated with this node.
	AcceleratorType *string `pulumi:"acceleratorType"`
	// The API version that created this Node.
	ApiVersion *string `pulumi:"apiVersion"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The time when the node was created.
	CreateTime *string `pulumi:"createTime"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// The health status of the TPU node.
	Health *string `pulumi:"health"`
	// If this field is populated, it contains a description of why the TPU Node is unhealthy.
	HealthDescription *string `pulumi:"healthDescription"`
	// DEPRECATED! Use network_endpoints instead. The network address for the TPU Node as visible to Compute Engine instances.
	IpAddress *string `pulumi:"ipAddress"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of the TPU
	Name *string `pulumi:"name"`
	// The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
	Network *string `pulumi:"network"`
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node reach out to the 0th entry in this map first.
	NetworkEndpoints []NetworkEndpointResponse `pulumi:"networkEndpoints"`
	// DEPRECATED! Use network_endpoints instead. The network port for the TPU Node as visible to Compute Engine instances.
	Port *string `pulumi:"port"`
	// The scheduling options for this node.
	SchedulingConfig *SchedulingConfigResponse `pulumi:"schedulingConfig"`
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The current state for the TPU Node.
	State *string `pulumi:"state"`
	// The Symptoms that have occurred to the TPU Node.
	Symptoms []SymptomResponse `pulumi:"symptoms"`
	// Required. The version of Tensorflow running in the Node.
	TensorflowVersion *string `pulumi:"tensorflowVersion"`
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
	UseServiceNetworking *bool `pulumi:"useServiceNetworking"`
}

type NodeState struct {
	// Required. The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringPtrInput
	// The API version that created this Node.
	ApiVersion pulumi.StringPtrInput
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The time when the node was created.
	CreateTime pulumi.StringPtrInput
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// The health status of the TPU node.
	Health pulumi.StringPtrInput
	// If this field is populated, it contains a description of why the TPU Node is unhealthy.
	HealthDescription pulumi.StringPtrInput
	// DEPRECATED! Use network_endpoints instead. The network address for the TPU Node as visible to Compute Engine instances.
	IpAddress pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Immutable. The name of the TPU
	Name pulumi.StringPtrInput
	// The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
	Network pulumi.StringPtrInput
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node reach out to the 0th entry in this map first.
	NetworkEndpoints NetworkEndpointResponseArrayInput
	// DEPRECATED! Use network_endpoints instead. The network port for the TPU Node as visible to Compute Engine instances.
	Port pulumi.StringPtrInput
	// The scheduling options for this node.
	SchedulingConfig SchedulingConfigResponsePtrInput
	// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
	ServiceAccount pulumi.StringPtrInput
	// The current state for the TPU Node.
	State pulumi.StringPtrInput
	// The Symptoms that have occurred to the TPU Node.
	Symptoms SymptomResponseArrayInput
	// Required. The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringPtrInput
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
	UseServiceNetworking pulumi.BoolPtrInput
}

func (NodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeState)(nil)).Elem()
}

type nodeArgs struct {
	// Required. The type of hardware accelerators associated with this node.
	AcceleratorType *string `pulumi:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// The health status of the TPU node.
	Health *string `pulumi:"health"`
	// DEPRECATED! Use network_endpoints instead. The network address for the TPU Node as visible to Compute Engine instances.
	IpAddress *string `pulumi:"ipAddress"`
	// Resource labels to represent user-provided metadata.
	Labels      map[string]string `pulumi:"labels"`
	LocationsId string            `pulumi:"locationsId"`
	// The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
	Network *string `pulumi:"network"`
	NodesId string  `pulumi:"nodesId"`
	// DEPRECATED! Use network_endpoints instead. The network port for the TPU Node as visible to Compute Engine instances.
	Port       *string `pulumi:"port"`
	ProjectsId string  `pulumi:"projectsId"`
	// The scheduling options for this node.
	SchedulingConfig *SchedulingConfig `pulumi:"schedulingConfig"`
	// Required. The version of Tensorflow running in the Node.
	TensorflowVersion *string `pulumi:"tensorflowVersion"`
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
	UseServiceNetworking *bool `pulumi:"useServiceNetworking"`
}

// The set of arguments for constructing a Node resource.
type NodeArgs struct {
	// Required. The type of hardware accelerators associated with this node.
	AcceleratorType pulumi.StringPtrInput
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The user-supplied description of the TPU. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// The health status of the TPU node.
	Health pulumi.StringPtrInput
	// DEPRECATED! Use network_endpoints instead. The network address for the TPU Node as visible to Compute Engine instances.
	IpAddress pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels      pulumi.StringMapInput
	LocationsId pulumi.StringInput
	// The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
	Network pulumi.StringPtrInput
	NodesId pulumi.StringInput
	// DEPRECATED! Use network_endpoints instead. The network port for the TPU Node as visible to Compute Engine instances.
	Port       pulumi.StringPtrInput
	ProjectsId pulumi.StringInput
	// The scheduling options for this node.
	SchedulingConfig SchedulingConfigPtrInput
	// Required. The version of Tensorflow running in the Node.
	TensorflowVersion pulumi.StringPtrInput
	// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
	UseServiceNetworking pulumi.BoolPtrInput
}

func (NodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeArgs)(nil)).Elem()
}

type NodeInput interface {
	pulumi.Input

	ToNodeOutput() NodeOutput
	ToNodeOutputWithContext(ctx context.Context) NodeOutput
}

func (*Node) ElementType() reflect.Type {
	return reflect.TypeOf((*Node)(nil))
}

func (i *Node) ToNodeOutput() NodeOutput {
	return i.ToNodeOutputWithContext(context.Background())
}

func (i *Node) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeOutput)
}

type NodeOutput struct {
	*pulumi.OutputState
}

func (NodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Node)(nil))
}

func (o NodeOutput) ToNodeOutput() NodeOutput {
	return o
}

func (o NodeOutput) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NodeOutput{})
}
