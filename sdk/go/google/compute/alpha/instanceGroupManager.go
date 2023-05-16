// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a managed instance group using the information that you specify in the request. After the group is created, instances in the group are created using the specified instance template. This operation is marked as DONE when the group is created even if the instances in the group have not yet been created. You must separately verify the status of the individual instances with the listmanagedinstances method. A managed instance group can have up to 1000 VM instances per group. Please contact Cloud Support if you need an increase in this limit.
type InstanceGroupManager struct {
	pulumi.CustomResourceState

	// Specifies configuration that overrides the instance template configuration for the group.
	AllInstancesConfig InstanceGroupManagerAllInstancesConfigResponseOutput `pulumi:"allInstancesConfig"`
	// The autohealing policy for this managed instance group. You can specify only one value.
	AutoHealingPolicies InstanceGroupManagerAutoHealingPolicyResponseArrayOutput `pulumi:"autoHealingPolicies"`
	// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
	BaseInstanceName pulumi.StringOutput `pulumi:"baseInstanceName"`
	// The creation timestamp for this managed instance group in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
	CurrentActions InstanceGroupManagerActionsSummaryResponseOutput `pulumi:"currentActions"`
	// An optional description of this resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	DistributionPolicy DistributionPolicyResponseOutput `pulumi:"distributionPolicy"`
	// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
	FailoverAction pulumi.StringOutput `pulumi:"failoverAction"`
	// Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Instance flexibility allowing MIG to create VMs from multiple types of machines. Instance flexibility configuration on MIG overrides instance template configuration.
	InstanceFlexibilityPolicy InstanceGroupManagerInstanceFlexibilityPolicyResponseOutput `pulumi:"instanceFlexibilityPolicy"`
	// The URL of the Instance Group resource.
	InstanceGroup pulumi.StringOutput `pulumi:"instanceGroup"`
	// The repair policy for this managed instance group.
	InstanceLifecyclePolicy InstanceGroupManagerInstanceLifecyclePolicyResponseOutput `pulumi:"instanceLifecyclePolicy"`
	// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
	InstanceTemplate pulumi.StringOutput `pulumi:"instanceTemplate"`
	// The resource type, which is always compute#instanceGroupManager for managed instance groups.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Pagination behavior of the listManagedInstances API method for this managed instance group.
	ListManagedInstancesResults pulumi.StringOutput `pulumi:"listManagedInstancesResults"`
	// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name pulumi.StringOutput `pulumi:"name"`
	// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
	NamedPorts NamedPortResponseArrayOutput `pulumi:"namedPorts"`
	Project    pulumi.StringOutput          `pulumi:"project"`
	// The URL of the region where the managed instance group resides (for regional resources).
	Region pulumi.StringOutput `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The URL for this managed instance group. The server defines this URL.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Standby policy for stopped and suspended instances.
	StandbyPolicy InstanceGroupManagerStandbyPolicyResponseOutput `pulumi:"standbyPolicy"`
	// Stateful configuration for this Instanced Group Manager
	StatefulPolicy StatefulPolicyResponseOutput `pulumi:"statefulPolicy"`
	// The status of this managed instance group.
	Status InstanceGroupManagerStatusResponseOutput `pulumi:"status"`
	// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
	TargetPools pulumi.StringArrayOutput `pulumi:"targetPools"`
	// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
	TargetSize pulumi.IntOutput `pulumi:"targetSize"`
	// The unit of measure for the target size.
	TargetSizeUnit pulumi.StringOutput `pulumi:"targetSizeUnit"`
	// The target number of stopped instances for this managed instance group. This number changes when you: - Stop instance using the stopInstances method or start instances using the startInstances method. - Manually change the targetStoppedSize using the update method.
	TargetStoppedSize pulumi.IntOutput `pulumi:"targetStoppedSize"`
	// The target number of suspended instances for this managed instance group. This number changes when you: - Suspend instance using the suspendInstances method or resume instances using the resumeInstances method. - Manually change the targetSuspendedSize using the update method.
	TargetSuspendedSize pulumi.IntOutput `pulumi:"targetSuspendedSize"`
	// The update policy for this managed instance group.
	UpdatePolicy InstanceGroupManagerUpdatePolicyResponseOutput `pulumi:"updatePolicy"`
	// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
	Versions InstanceGroupManagerVersionResponseArrayOutput `pulumi:"versions"`
	Zone     pulumi.StringOutput                            `pulumi:"zone"`
}

// NewInstanceGroupManager registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroupManager(ctx *pulumi.Context,
	name string, args *InstanceGroupManagerArgs, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	if args == nil {
		args = &InstanceGroupManagerArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"zone",
	})
	opts = append(opts, replaceOnChanges)
	var resource InstanceGroupManager
	err := ctx.RegisterResource("google-native:compute/alpha:InstanceGroupManager", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:compute/alpha:InstanceGroupManager", name, id, state, &resource, opts...)
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
	// Specifies configuration that overrides the instance template configuration for the group.
	AllInstancesConfig *InstanceGroupManagerAllInstancesConfig `pulumi:"allInstancesConfig"`
	// The autohealing policy for this managed instance group. You can specify only one value.
	AutoHealingPolicies []InstanceGroupManagerAutoHealingPolicy `pulumi:"autoHealingPolicies"`
	// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
	BaseInstanceName *string `pulumi:"baseInstanceName"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	DistributionPolicy *DistributionPolicy `pulumi:"distributionPolicy"`
	// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
	FailoverAction *InstanceGroupManagerFailoverAction `pulumi:"failoverAction"`
	// Instance flexibility allowing MIG to create VMs from multiple types of machines. Instance flexibility configuration on MIG overrides instance template configuration.
	InstanceFlexibilityPolicy *InstanceGroupManagerInstanceFlexibilityPolicy `pulumi:"instanceFlexibilityPolicy"`
	// The repair policy for this managed instance group.
	InstanceLifecyclePolicy *InstanceGroupManagerInstanceLifecyclePolicy `pulumi:"instanceLifecyclePolicy"`
	// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
	InstanceTemplate *string `pulumi:"instanceTemplate"`
	// Pagination behavior of the listManagedInstances API method for this managed instance group.
	ListManagedInstancesResults *InstanceGroupManagerListManagedInstancesResults `pulumi:"listManagedInstancesResults"`
	// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name *string `pulumi:"name"`
	// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
	NamedPorts []NamedPort `pulumi:"namedPorts"`
	Project    *string     `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Standby policy for stopped and suspended instances.
	StandbyPolicy *InstanceGroupManagerStandbyPolicy `pulumi:"standbyPolicy"`
	// Stateful configuration for this Instanced Group Manager
	StatefulPolicy *StatefulPolicy `pulumi:"statefulPolicy"`
	// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
	TargetPools []string `pulumi:"targetPools"`
	// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
	TargetSize *int `pulumi:"targetSize"`
	// The unit of measure for the target size.
	TargetSizeUnit *InstanceGroupManagerTargetSizeUnit `pulumi:"targetSizeUnit"`
	// The target number of stopped instances for this managed instance group. This number changes when you: - Stop instance using the stopInstances method or start instances using the startInstances method. - Manually change the targetStoppedSize using the update method.
	TargetStoppedSize *int `pulumi:"targetStoppedSize"`
	// The target number of suspended instances for this managed instance group. This number changes when you: - Suspend instance using the suspendInstances method or resume instances using the resumeInstances method. - Manually change the targetSuspendedSize using the update method.
	TargetSuspendedSize *int `pulumi:"targetSuspendedSize"`
	// The update policy for this managed instance group.
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
	Versions []InstanceGroupManagerVersion `pulumi:"versions"`
	Zone     *string                       `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroupManager resource.
type InstanceGroupManagerArgs struct {
	// Specifies configuration that overrides the instance template configuration for the group.
	AllInstancesConfig InstanceGroupManagerAllInstancesConfigPtrInput
	// The autohealing policy for this managed instance group. You can specify only one value.
	AutoHealingPolicies InstanceGroupManagerAutoHealingPolicyArrayInput
	// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
	BaseInstanceName pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	DistributionPolicy DistributionPolicyPtrInput
	// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
	FailoverAction InstanceGroupManagerFailoverActionPtrInput
	// Instance flexibility allowing MIG to create VMs from multiple types of machines. Instance flexibility configuration on MIG overrides instance template configuration.
	InstanceFlexibilityPolicy InstanceGroupManagerInstanceFlexibilityPolicyPtrInput
	// The repair policy for this managed instance group.
	InstanceLifecyclePolicy InstanceGroupManagerInstanceLifecyclePolicyPtrInput
	// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
	InstanceTemplate pulumi.StringPtrInput
	// Pagination behavior of the listManagedInstances API method for this managed instance group.
	ListManagedInstancesResults InstanceGroupManagerListManagedInstancesResultsPtrInput
	// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name pulumi.StringPtrInput
	// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
	NamedPorts NamedPortArrayInput
	Project    pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
	ServiceAccount pulumi.StringPtrInput
	// Standby policy for stopped and suspended instances.
	StandbyPolicy InstanceGroupManagerStandbyPolicyPtrInput
	// Stateful configuration for this Instanced Group Manager
	StatefulPolicy StatefulPolicyPtrInput
	// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
	TargetPools pulumi.StringArrayInput
	// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
	TargetSize pulumi.IntPtrInput
	// The unit of measure for the target size.
	TargetSizeUnit InstanceGroupManagerTargetSizeUnitPtrInput
	// The target number of stopped instances for this managed instance group. This number changes when you: - Stop instance using the stopInstances method or start instances using the startInstances method. - Manually change the targetStoppedSize using the update method.
	TargetStoppedSize pulumi.IntPtrInput
	// The target number of suspended instances for this managed instance group. This number changes when you: - Suspend instance using the suspendInstances method or resume instances using the resumeInstances method. - Manually change the targetSuspendedSize using the update method.
	TargetSuspendedSize pulumi.IntPtrInput
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
	return reflect.TypeOf((**InstanceGroupManager)(nil)).Elem()
}

func (i *InstanceGroupManager) ToInstanceGroupManagerOutput() InstanceGroupManagerOutput {
	return i.ToInstanceGroupManagerOutputWithContext(context.Background())
}

func (i *InstanceGroupManager) ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerOutput)
}

type InstanceGroupManagerOutput struct{ *pulumi.OutputState }

func (InstanceGroupManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupManager)(nil)).Elem()
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerOutput() InstanceGroupManagerOutput {
	return o
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput {
	return o
}

// Specifies configuration that overrides the instance template configuration for the group.
func (o InstanceGroupManagerOutput) AllInstancesConfig() InstanceGroupManagerAllInstancesConfigResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerAllInstancesConfigResponseOutput {
		return v.AllInstancesConfig
	}).(InstanceGroupManagerAllInstancesConfigResponseOutput)
}

// The autohealing policy for this managed instance group. You can specify only one value.
func (o InstanceGroupManagerOutput) AutoHealingPolicies() InstanceGroupManagerAutoHealingPolicyResponseArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerAutoHealingPolicyResponseArrayOutput {
		return v.AutoHealingPolicies
	}).(InstanceGroupManagerAutoHealingPolicyResponseArrayOutput)
}

// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
func (o InstanceGroupManagerOutput) BaseInstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.BaseInstanceName }).(pulumi.StringOutput)
}

// The creation timestamp for this managed instance group in RFC3339 text format.
func (o InstanceGroupManagerOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
func (o InstanceGroupManagerOutput) CurrentActions() InstanceGroupManagerActionsSummaryResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerActionsSummaryResponseOutput {
		return v.CurrentActions
	}).(InstanceGroupManagerActionsSummaryResponseOutput)
}

// An optional description of this resource.
func (o InstanceGroupManagerOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
func (o InstanceGroupManagerOutput) DistributionPolicy() DistributionPolicyResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) DistributionPolicyResponseOutput { return v.DistributionPolicy }).(DistributionPolicyResponseOutput)
}

// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
func (o InstanceGroupManagerOutput) FailoverAction() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.FailoverAction }).(pulumi.StringOutput)
}

// Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
func (o InstanceGroupManagerOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Instance flexibility allowing MIG to create VMs from multiple types of machines. Instance flexibility configuration on MIG overrides instance template configuration.
func (o InstanceGroupManagerOutput) InstanceFlexibilityPolicy() InstanceGroupManagerInstanceFlexibilityPolicyResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerInstanceFlexibilityPolicyResponseOutput {
		return v.InstanceFlexibilityPolicy
	}).(InstanceGroupManagerInstanceFlexibilityPolicyResponseOutput)
}

// The URL of the Instance Group resource.
func (o InstanceGroupManagerOutput) InstanceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.InstanceGroup }).(pulumi.StringOutput)
}

// The repair policy for this managed instance group.
func (o InstanceGroupManagerOutput) InstanceLifecyclePolicy() InstanceGroupManagerInstanceLifecyclePolicyResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerInstanceLifecyclePolicyResponseOutput {
		return v.InstanceLifecyclePolicy
	}).(InstanceGroupManagerInstanceLifecyclePolicyResponseOutput)
}

// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
func (o InstanceGroupManagerOutput) InstanceTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.InstanceTemplate }).(pulumi.StringOutput)
}

// The resource type, which is always compute#instanceGroupManager for managed instance groups.
func (o InstanceGroupManagerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Pagination behavior of the listManagedInstances API method for this managed instance group.
func (o InstanceGroupManagerOutput) ListManagedInstancesResults() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.ListManagedInstancesResults }).(pulumi.StringOutput)
}

// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
func (o InstanceGroupManagerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
func (o InstanceGroupManagerOutput) NamedPorts() NamedPortResponseArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) NamedPortResponseArrayOutput { return v.NamedPorts }).(NamedPortResponseArrayOutput)
}

func (o InstanceGroupManagerOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URL of the region where the managed instance group resides (for regional resources).
func (o InstanceGroupManagerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o InstanceGroupManagerOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The URL for this managed instance group. The server defines this URL.
func (o InstanceGroupManagerOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o InstanceGroupManagerOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
func (o InstanceGroupManagerOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Standby policy for stopped and suspended instances.
func (o InstanceGroupManagerOutput) StandbyPolicy() InstanceGroupManagerStandbyPolicyResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerStandbyPolicyResponseOutput { return v.StandbyPolicy }).(InstanceGroupManagerStandbyPolicyResponseOutput)
}

// Stateful configuration for this Instanced Group Manager
func (o InstanceGroupManagerOutput) StatefulPolicy() StatefulPolicyResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) StatefulPolicyResponseOutput { return v.StatefulPolicy }).(StatefulPolicyResponseOutput)
}

// The status of this managed instance group.
func (o InstanceGroupManagerOutput) Status() InstanceGroupManagerStatusResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerStatusResponseOutput { return v.Status }).(InstanceGroupManagerStatusResponseOutput)
}

// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
func (o InstanceGroupManagerOutput) TargetPools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringArrayOutput { return v.TargetPools }).(pulumi.StringArrayOutput)
}

// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
func (o InstanceGroupManagerOutput) TargetSize() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.IntOutput { return v.TargetSize }).(pulumi.IntOutput)
}

// The unit of measure for the target size.
func (o InstanceGroupManagerOutput) TargetSizeUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.TargetSizeUnit }).(pulumi.StringOutput)
}

// The target number of stopped instances for this managed instance group. This number changes when you: - Stop instance using the stopInstances method or start instances using the startInstances method. - Manually change the targetStoppedSize using the update method.
func (o InstanceGroupManagerOutput) TargetStoppedSize() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.IntOutput { return v.TargetStoppedSize }).(pulumi.IntOutput)
}

// The target number of suspended instances for this managed instance group. This number changes when you: - Suspend instance using the suspendInstances method or resume instances using the resumeInstances method. - Manually change the targetSuspendedSize using the update method.
func (o InstanceGroupManagerOutput) TargetSuspendedSize() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.IntOutput { return v.TargetSuspendedSize }).(pulumi.IntOutput)
}

// The update policy for this managed instance group.
func (o InstanceGroupManagerOutput) UpdatePolicy() InstanceGroupManagerUpdatePolicyResponseOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerUpdatePolicyResponseOutput { return v.UpdatePolicy }).(InstanceGroupManagerUpdatePolicyResponseOutput)
}

// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
func (o InstanceGroupManagerOutput) Versions() InstanceGroupManagerVersionResponseArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerVersionResponseArrayOutput { return v.Versions }).(InstanceGroupManagerVersionResponseArrayOutput)
}

func (o InstanceGroupManagerOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupManagerInput)(nil)).Elem(), &InstanceGroupManager{})
	pulumi.RegisterOutputType(InstanceGroupManagerOutput{})
}
