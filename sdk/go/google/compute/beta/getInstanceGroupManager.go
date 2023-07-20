// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns all of the details about the specified managed instance group.
func LookupInstanceGroupManager(ctx *pulumi.Context, args *LookupInstanceGroupManagerArgs, opts ...pulumi.InvokeOption) (*LookupInstanceGroupManagerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceGroupManagerResult
	err := ctx.Invoke("google-native:compute/beta:getInstanceGroupManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceGroupManagerArgs struct {
	InstanceGroupManager string  `pulumi:"instanceGroupManager"`
	Project              *string `pulumi:"project"`
	Zone                 string  `pulumi:"zone"`
}

type LookupInstanceGroupManagerResult struct {
	// Specifies configuration that overrides the instance template configuration for the group.
	AllInstancesConfig InstanceGroupManagerAllInstancesConfigResponse `pulumi:"allInstancesConfig"`
	// The autohealing policy for this managed instance group. You can specify only one value.
	AutoHealingPolicies []InstanceGroupManagerAutoHealingPolicyResponse `pulumi:"autoHealingPolicies"`
	// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
	BaseInstanceName string `pulumi:"baseInstanceName"`
	// The creation timestamp for this managed instance group in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
	CurrentActions InstanceGroupManagerActionsSummaryResponse `pulumi:"currentActions"`
	// An optional description of this resource.
	Description string `pulumi:"description"`
	// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
	DistributionPolicy DistributionPolicyResponse `pulumi:"distributionPolicy"`
	// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
	FailoverAction string `pulumi:"failoverAction"`
	// Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
	Fingerprint string `pulumi:"fingerprint"`
	// The URL of the Instance Group resource.
	InstanceGroup string `pulumi:"instanceGroup"`
	// The repair policy for this managed instance group.
	InstanceLifecyclePolicy InstanceGroupManagerInstanceLifecyclePolicyResponse `pulumi:"instanceLifecyclePolicy"`
	// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
	InstanceTemplate string `pulumi:"instanceTemplate"`
	// The resource type, which is always compute#instanceGroupManager for managed instance groups.
	Kind string `pulumi:"kind"`
	// Pagination behavior of the listManagedInstances API method for this managed instance group.
	ListManagedInstancesResults string `pulumi:"listManagedInstancesResults"`
	// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name string `pulumi:"name"`
	// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
	NamedPorts []NamedPortResponse `pulumi:"namedPorts"`
	// The URL of the region where the managed instance group resides (for regional resources).
	Region string `pulumi:"region"`
	// The URL for this managed instance group. The server defines this URL.
	SelfLink string `pulumi:"selfLink"`
	// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Stateful configuration for this Instanced Group Manager
	StatefulPolicy StatefulPolicyResponse `pulumi:"statefulPolicy"`
	// The status of this managed instance group.
	Status InstanceGroupManagerStatusResponse `pulumi:"status"`
	// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
	TargetPools []string `pulumi:"targetPools"`
	// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
	TargetSize int `pulumi:"targetSize"`
	// The update policy for this managed instance group.
	UpdatePolicy InstanceGroupManagerUpdatePolicyResponse `pulumi:"updatePolicy"`
	// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
	Versions []InstanceGroupManagerVersionResponse `pulumi:"versions"`
	// The URL of a zone where the managed instance group is located (for zonal resources).
	Zone string `pulumi:"zone"`
}

func LookupInstanceGroupManagerOutput(ctx *pulumi.Context, args LookupInstanceGroupManagerOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceGroupManagerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceGroupManagerResult, error) {
			args := v.(LookupInstanceGroupManagerArgs)
			r, err := LookupInstanceGroupManager(ctx, &args, opts...)
			var s LookupInstanceGroupManagerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceGroupManagerResultOutput)
}

type LookupInstanceGroupManagerOutputArgs struct {
	InstanceGroupManager pulumi.StringInput    `pulumi:"instanceGroupManager"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
	Zone                 pulumi.StringInput    `pulumi:"zone"`
}

func (LookupInstanceGroupManagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceGroupManagerArgs)(nil)).Elem()
}

type LookupInstanceGroupManagerResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceGroupManagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceGroupManagerResult)(nil)).Elem()
}

func (o LookupInstanceGroupManagerResultOutput) ToLookupInstanceGroupManagerResultOutput() LookupInstanceGroupManagerResultOutput {
	return o
}

func (o LookupInstanceGroupManagerResultOutput) ToLookupInstanceGroupManagerResultOutputWithContext(ctx context.Context) LookupInstanceGroupManagerResultOutput {
	return o
}

// Specifies configuration that overrides the instance template configuration for the group.
func (o LookupInstanceGroupManagerResultOutput) AllInstancesConfig() InstanceGroupManagerAllInstancesConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) InstanceGroupManagerAllInstancesConfigResponse {
		return v.AllInstancesConfig
	}).(InstanceGroupManagerAllInstancesConfigResponseOutput)
}

// The autohealing policy for this managed instance group. You can specify only one value.
func (o LookupInstanceGroupManagerResultOutput) AutoHealingPolicies() InstanceGroupManagerAutoHealingPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) []InstanceGroupManagerAutoHealingPolicyResponse {
		return v.AutoHealingPolicies
	}).(InstanceGroupManagerAutoHealingPolicyResponseArrayOutput)
}

// The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
func (o LookupInstanceGroupManagerResultOutput) BaseInstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.BaseInstanceName }).(pulumi.StringOutput)
}

// The creation timestamp for this managed instance group in RFC3339 text format.
func (o LookupInstanceGroupManagerResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
func (o LookupInstanceGroupManagerResultOutput) CurrentActions() InstanceGroupManagerActionsSummaryResponseOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) InstanceGroupManagerActionsSummaryResponse {
		return v.CurrentActions
	}).(InstanceGroupManagerActionsSummaryResponseOutput)
}

// An optional description of this resource.
func (o LookupInstanceGroupManagerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.Description }).(pulumi.StringOutput)
}

// Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
func (o LookupInstanceGroupManagerResultOutput) DistributionPolicy() DistributionPolicyResponseOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) DistributionPolicyResponse { return v.DistributionPolicy }).(DistributionPolicyResponseOutput)
}

// The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
func (o LookupInstanceGroupManagerResultOutput) FailoverAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.FailoverAction }).(pulumi.StringOutput)
}

// Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
func (o LookupInstanceGroupManagerResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The URL of the Instance Group resource.
func (o LookupInstanceGroupManagerResultOutput) InstanceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.InstanceGroup }).(pulumi.StringOutput)
}

// The repair policy for this managed instance group.
func (o LookupInstanceGroupManagerResultOutput) InstanceLifecyclePolicy() InstanceGroupManagerInstanceLifecyclePolicyResponseOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) InstanceGroupManagerInstanceLifecyclePolicyResponse {
		return v.InstanceLifecyclePolicy
	}).(InstanceGroupManagerInstanceLifecyclePolicyResponseOutput)
}

// The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
func (o LookupInstanceGroupManagerResultOutput) InstanceTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.InstanceTemplate }).(pulumi.StringOutput)
}

// The resource type, which is always compute#instanceGroupManager for managed instance groups.
func (o LookupInstanceGroupManagerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Pagination behavior of the listManagedInstances API method for this managed instance group.
func (o LookupInstanceGroupManagerResultOutput) ListManagedInstancesResults() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.ListManagedInstancesResults }).(pulumi.StringOutput)
}

// The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
func (o LookupInstanceGroupManagerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Named ports configured for the Instance Groups complementary to this Instance Group Manager.
func (o LookupInstanceGroupManagerResultOutput) NamedPorts() NamedPortResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) []NamedPortResponse { return v.NamedPorts }).(NamedPortResponseArrayOutput)
}

// The URL of the region where the managed instance group resides (for regional resources).
func (o LookupInstanceGroupManagerResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.Region }).(pulumi.StringOutput)
}

// The URL for this managed instance group. The server defines this URL.
func (o LookupInstanceGroupManagerResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
func (o LookupInstanceGroupManagerResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Stateful configuration for this Instanced Group Manager
func (o LookupInstanceGroupManagerResultOutput) StatefulPolicy() StatefulPolicyResponseOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) StatefulPolicyResponse { return v.StatefulPolicy }).(StatefulPolicyResponseOutput)
}

// The status of this managed instance group.
func (o LookupInstanceGroupManagerResultOutput) Status() InstanceGroupManagerStatusResponseOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) InstanceGroupManagerStatusResponse { return v.Status }).(InstanceGroupManagerStatusResponseOutput)
}

// The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
func (o LookupInstanceGroupManagerResultOutput) TargetPools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) []string { return v.TargetPools }).(pulumi.StringArrayOutput)
}

// The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
func (o LookupInstanceGroupManagerResultOutput) TargetSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) int { return v.TargetSize }).(pulumi.IntOutput)
}

// The update policy for this managed instance group.
func (o LookupInstanceGroupManagerResultOutput) UpdatePolicy() InstanceGroupManagerUpdatePolicyResponseOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) InstanceGroupManagerUpdatePolicyResponse {
		return v.UpdatePolicy
	}).(InstanceGroupManagerUpdatePolicyResponseOutput)
}

// Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
func (o LookupInstanceGroupManagerResultOutput) Versions() InstanceGroupManagerVersionResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) []InstanceGroupManagerVersionResponse { return v.Versions }).(InstanceGroupManagerVersionResponseArrayOutput)
}

// The URL of a zone where the managed instance group is located (for zonal resources).
func (o LookupInstanceGroupManagerResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupManagerResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceGroupManagerResultOutput{})
}
