// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a managed instance group using the information that you specify in the request. After the group is created, instances in the group are created using the specified instance template. This operation is marked as DONE when the group is created even if the instances in the group have not yet been created. You must separately verify the status of the individual instances with the listmanagedinstances method. A managed instance group can have up to 1000 VM instances per group. Please contact Cloud Support if you need an increase in this limit.
type InstanceGroupManager struct {
	pulumi.CustomResourceState

	// The autohealing policy for this managed instance group. You can specify only one value.
	AutoHealingPolicies InstanceGroupManagerAutoHealingPolicyResponseArrayOutput `pulumi:"autoHealingPolicies"`
	// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
	BaseInstanceName pulumi.StringOutput `pulumi:"baseInstanceName"`
	// The creation timestamp for this managed instance group in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
	CurrentActions InstanceGroupManagerActionsSummaryResponseOutput `pulumi:"currentActions"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	DistributionPolicy DistributionPolicyResponseOutput `pulumi:"distributionPolicy"`
	// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
	FailoverAction pulumi.StringOutput `pulumi:"failoverAction"`
	// Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The URL of the Instance Group resource.
	InstanceGroup pulumi.StringOutput `pulumi:"instanceGroup"`
	// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
	InstanceTemplate pulumi.StringOutput `pulumi:"instanceTemplate"`
	// The resource type, which is always compute#instanceGroupManager for managed instance groups.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name pulumi.StringOutput `pulumi:"name"`
	// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
	NamedPorts NamedPortResponseArrayOutput `pulumi:"namedPorts"`
	// The URL of the region where the managed instance group resides (for regional resources).
	Region pulumi.StringOutput `pulumi:"region"`
	// The URL for this managed instance group. The server defines this URL.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Stateful configuration for this Instanced Group Manager
	StatefulPolicy StatefulPolicyResponseOutput `pulumi:"statefulPolicy"`
	// The status of this managed instance group.
	Status InstanceGroupManagerStatusResponseOutput `pulumi:"status"`
	// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
	TargetPools pulumi.StringArrayOutput `pulumi:"targetPools"`
	// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
	TargetSize pulumi.IntOutput `pulumi:"targetSize"`
	// The update policy for this managed instance group.
	UpdatePolicy InstanceGroupManagerUpdatePolicyResponseOutput `pulumi:"updatePolicy"`
	// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
	Versions InstanceGroupManagerVersionResponseArrayOutput `pulumi:"versions"`
	// The URL of a zone where the managed instance group is located (for zonal resources).
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroupManager registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroupManager(ctx *pulumi.Context,
	name string, args *InstanceGroupManagerArgs, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	if args == nil {
		args = &InstanceGroupManagerArgs{}
	}

	var resource InstanceGroupManager
	err := ctx.RegisterResource("google-native:compute/beta:InstanceGroupManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroupManager gets an existing InstanceGroupManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroupManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupManagerState, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	var resource InstanceGroupManager
	err := ctx.ReadResource("google-native:compute/beta:InstanceGroupManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroupManager resources.
type instanceGroupManagerState struct {
}

type InstanceGroupManagerState struct {
}

func (InstanceGroupManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupManagerState)(nil)).Elem()
}

type instanceGroupManagerArgs struct {
	// The autohealing policy for this managed instance group. You can specify only one value.
	AutoHealingPolicies []InstanceGroupManagerAutoHealingPolicy `pulumi:"autoHealingPolicies"`
	// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
	BaseInstanceName *string `pulumi:"baseInstanceName"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	DistributionPolicy *DistributionPolicy `pulumi:"distributionPolicy"`
	// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
	FailoverAction *InstanceGroupManagerFailoverAction `pulumi:"failoverAction"`
	// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
	InstanceTemplate *string `pulumi:"instanceTemplate"`
	// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name *string `pulumi:"name"`
	// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
	NamedPorts []NamedPort `pulumi:"namedPorts"`
	Project    *string     `pulumi:"project"`
	RequestId  *string     `pulumi:"requestId"`
	// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Stateful configuration for this Instanced Group Manager
	StatefulPolicy *StatefulPolicy `pulumi:"statefulPolicy"`
	// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
	TargetPools []string `pulumi:"targetPools"`
	// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group.
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
	Versions []InstanceGroupManagerVersion `pulumi:"versions"`
	Zone     *string                       `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroupManager resource.
type InstanceGroupManagerArgs struct {
	// The autohealing policy for this managed instance group. You can specify only one value.
	AutoHealingPolicies InstanceGroupManagerAutoHealingPolicyArrayInput
	// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
	BaseInstanceName pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	DistributionPolicy DistributionPolicyPtrInput
	// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
	FailoverAction InstanceGroupManagerFailoverActionPtrInput
	// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
	InstanceTemplate pulumi.StringPtrInput
	// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name pulumi.StringPtrInput
	// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
	NamedPorts NamedPortArrayInput
	Project    pulumi.StringPtrInput
	RequestId  pulumi.StringPtrInput
	// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
	ServiceAccount pulumi.StringPtrInput
	// Stateful configuration for this Instanced Group Manager
	StatefulPolicy StatefulPolicyPtrInput
	// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
	TargetPools pulumi.StringArrayInput
	// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group.
	UpdatePolicy InstanceGroupManagerUpdatePolicyPtrInput
	// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
	Versions InstanceGroupManagerVersionArrayInput
	Zone     pulumi.StringPtrInput
}

func (InstanceGroupManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupManagerArgs)(nil)).Elem()
}

type InstanceGroupManagerInput interface {
	pulumi.Input

	ToInstanceGroupManagerOutput() InstanceGroupManagerOutput
	ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput
}

func (*InstanceGroupManager) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroupManager)(nil))
}

func (i *InstanceGroupManager) ToInstanceGroupManagerOutput() InstanceGroupManagerOutput {
	return i.ToInstanceGroupManagerOutputWithContext(context.Background())
}

func (i *InstanceGroupManager) ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerOutput)
}

type InstanceGroupManagerOutput struct {
	*pulumi.OutputState
}

func (InstanceGroupManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGroupManager)(nil))
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerOutput() InstanceGroupManagerOutput {
	return o
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceGroupManagerOutput{})
}
