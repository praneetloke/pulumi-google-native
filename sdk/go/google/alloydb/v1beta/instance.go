// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Instance in a given project and location.
// Auto-naming is currently not supported for this resource.
type Instance struct {
	pulumi.CustomResourceState

	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Availability type of an Instance. If empty, defaults to REGIONAL for primary instances. For read pools, availability_type is always UNSPECIFIED. Instances in the read pools are evenly distributed across available zones within the region (i.e. read pools with more than one node will have a node in at least two zones).
	AvailabilityType pulumi.StringOutput `pulumi:"availabilityType"`
	// Optional. Client connection specific configurations
	ClientConnectionConfig ClientConnectionConfigResponseOutput `pulumi:"clientConnectionConfig"`
	ClusterId              pulumi.StringOutput                  `pulumi:"clusterId"`
	// Create time stamp
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary. This is a list of "key": "value" pairs. "key": The name of the flag. These flags are passed at instance setup time, so include both server options and system variables for Postgres. Flags are specified with underscores, not hyphens. "value": The value of the flag. Booleans are set to **on** for true and **off** for false. This field must be omitted if the flag doesn't take a value.
	DatabaseFlags pulumi.StringMapOutput `pulumi:"databaseFlags"`
	// Delete time stamp
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// User-settable and human-readable display name for the Instance.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// For Resource freshness validation (https://google.aip.dev/154)
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	GceZone pulumi.StringOutput `pulumi:"gceZone"`
	// Required. ID of the requesting object.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The type of the instance. Specified at creation time.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The IP address for the Instance. This is the connection endpoint for an end-user application.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Labels as key value pairs
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Configurations for the machines that host the underlying database engine.
	MachineConfig MachineConfigResponseOutput `pulumi:"machineConfig"`
	// The name of the instance resource with the format: * projects/{project}/locations/{region}/clusters/{cluster_id}/instances/{instance_id} where the cluster and instance ID segments should satisfy the regex expression `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`, e.g. 1-63 characters of lowercase letters, numbers, and dashes, starting with a letter, and ending with a letter or number. For more details see https://google.aip.dev/122. The prefix of the instance resource name is the name of the parent resource: * projects/{project}/locations/{region}/clusters/{cluster_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// List of available read-only VMs in this instance, including the standby for a PRIMARY instance.
	Nodes   NodeResponseArrayOutput `pulumi:"nodes"`
	Project pulumi.StringOutput     `pulumi:"project"`
	// Configuration for query insights.
	QueryInsightsConfig QueryInsightsInstanceConfigResponseOutput `pulumi:"queryInsightsConfig"`
	// Read pool instance configuration. This is required if the value of instanceType is READ_POOL.
	ReadPoolConfig ReadPoolConfigResponseOutput `pulumi:"readPoolConfig"`
	// Reconciling (https://google.aip.dev/128#reconciliation). Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The current serving state of the instance.
	State pulumi.StringOutput `pulumi:"state"`
	// The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Update policy that will be applied during instance update. This field is not persisted when you update the instance. To use a non-default update policy, you must specify explicitly specify the value in each update request.
	UpdatePolicy UpdatePolicyResponseOutput `pulumi:"updatePolicy"`
	// Update time stamp
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// This is set for the read-write VM of the PRIMARY instance only.
	WritableNode NodeResponseOutput `pulumi:"writableNode"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clusterId",
		"instanceId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("google-native:alloydb/v1beta:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("google-native:alloydb/v1beta:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
	Annotations map[string]string `pulumi:"annotations"`
	// Availability type of an Instance. If empty, defaults to REGIONAL for primary instances. For read pools, availability_type is always UNSPECIFIED. Instances in the read pools are evenly distributed across available zones within the region (i.e. read pools with more than one node will have a node in at least two zones).
	AvailabilityType *InstanceAvailabilityType `pulumi:"availabilityType"`
	// Optional. Client connection specific configurations
	ClientConnectionConfig *ClientConnectionConfig `pulumi:"clientConnectionConfig"`
	ClusterId              string                  `pulumi:"clusterId"`
	// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary. This is a list of "key": "value" pairs. "key": The name of the flag. These flags are passed at instance setup time, so include both server options and system variables for Postgres. Flags are specified with underscores, not hyphens. "value": The value of the flag. Booleans are set to **on** for true and **off** for false. This field must be omitted if the flag doesn't take a value.
	DatabaseFlags map[string]string `pulumi:"databaseFlags"`
	// User-settable and human-readable display name for the Instance.
	DisplayName *string `pulumi:"displayName"`
	// For Resource freshness validation (https://google.aip.dev/154)
	Etag *string `pulumi:"etag"`
	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	GceZone *string `pulumi:"gceZone"`
	// Required. ID of the requesting object.
	InstanceId string `pulumi:"instanceId"`
	// The type of the instance. Specified at creation time.
	InstanceType InstanceInstanceType `pulumi:"instanceType"`
	// Labels as key value pairs
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Configurations for the machines that host the underlying database engine.
	MachineConfig *MachineConfig `pulumi:"machineConfig"`
	Project       *string        `pulumi:"project"`
	// Configuration for query insights.
	QueryInsightsConfig *QueryInsightsInstanceConfig `pulumi:"queryInsightsConfig"`
	// Read pool instance configuration. This is required if the value of instanceType is READ_POOL.
	ReadPoolConfig *ReadPoolConfig `pulumi:"readPoolConfig"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Update policy that will be applied during instance update. This field is not persisted when you update the instance. To use a non-default update policy, you must specify explicitly specify the value in each update request.
	UpdatePolicy *UpdatePolicy `pulumi:"updatePolicy"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
	Annotations pulumi.StringMapInput
	// Availability type of an Instance. If empty, defaults to REGIONAL for primary instances. For read pools, availability_type is always UNSPECIFIED. Instances in the read pools are evenly distributed across available zones within the region (i.e. read pools with more than one node will have a node in at least two zones).
	AvailabilityType InstanceAvailabilityTypePtrInput
	// Optional. Client connection specific configurations
	ClientConnectionConfig ClientConnectionConfigPtrInput
	ClusterId              pulumi.StringInput
	// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary. This is a list of "key": "value" pairs. "key": The name of the flag. These flags are passed at instance setup time, so include both server options and system variables for Postgres. Flags are specified with underscores, not hyphens. "value": The value of the flag. Booleans are set to **on** for true and **off** for false. This field must be omitted if the flag doesn't take a value.
	DatabaseFlags pulumi.StringMapInput
	// User-settable and human-readable display name for the Instance.
	DisplayName pulumi.StringPtrInput
	// For Resource freshness validation (https://google.aip.dev/154)
	Etag pulumi.StringPtrInput
	// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
	GceZone pulumi.StringPtrInput
	// Required. ID of the requesting object.
	InstanceId pulumi.StringInput
	// The type of the instance. Specified at creation time.
	InstanceType InstanceInstanceTypeInput
	// Labels as key value pairs
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Configurations for the machines that host the underlying database engine.
	MachineConfig MachineConfigPtrInput
	Project       pulumi.StringPtrInput
	// Configuration for query insights.
	QueryInsightsConfig QueryInsightsInstanceConfigPtrInput
	// Read pool instance configuration. This is required if the value of instanceType is READ_POOL.
	ReadPoolConfig ReadPoolConfigPtrInput
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Update policy that will be applied during instance update. This field is not persisted when you update the instance. To use a non-default update policy, you must specify explicitly specify the value in each update request.
	UpdatePolicy UpdatePolicyPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
func (o InstanceOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Availability type of an Instance. If empty, defaults to REGIONAL for primary instances. For read pools, availability_type is always UNSPECIFIED. Instances in the read pools are evenly distributed across available zones within the region (i.e. read pools with more than one node will have a node in at least two zones).
func (o InstanceOutput) AvailabilityType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AvailabilityType }).(pulumi.StringOutput)
}

// Optional. Client connection specific configurations
func (o InstanceOutput) ClientConnectionConfig() ClientConnectionConfigResponseOutput {
	return o.ApplyT(func(v *Instance) ClientConnectionConfigResponseOutput { return v.ClientConnectionConfig }).(ClientConnectionConfigResponseOutput)
}

func (o InstanceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Create time stamp
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary. This is a list of "key": "value" pairs. "key": The name of the flag. These flags are passed at instance setup time, so include both server options and system variables for Postgres. Flags are specified with underscores, not hyphens. "value": The value of the flag. Booleans are set to **on** for true and **off** for false. This field must be omitted if the flag doesn't take a value.
func (o InstanceOutput) DatabaseFlags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.DatabaseFlags }).(pulumi.StringMapOutput)
}

// Delete time stamp
func (o InstanceOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// User-settable and human-readable display name for the Instance.
func (o InstanceOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// For Resource freshness validation (https://google.aip.dev/154)
func (o InstanceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.
func (o InstanceOutput) GceZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.GceZone }).(pulumi.StringOutput)
}

// Required. ID of the requesting object.
func (o InstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The type of the instance. Specified at creation time.
func (o InstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The IP address for the Instance. This is the connection endpoint for an end-user application.
func (o InstanceOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// Labels as key value pairs
func (o InstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Configurations for the machines that host the underlying database engine.
func (o InstanceOutput) MachineConfig() MachineConfigResponseOutput {
	return o.ApplyT(func(v *Instance) MachineConfigResponseOutput { return v.MachineConfig }).(MachineConfigResponseOutput)
}

// The name of the instance resource with the format: * projects/{project}/locations/{region}/clusters/{cluster_id}/instances/{instance_id} where the cluster and instance ID segments should satisfy the regex expression `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`, e.g. 1-63 characters of lowercase letters, numbers, and dashes, starting with a letter, and ending with a letter or number. For more details see https://google.aip.dev/122. The prefix of the instance resource name is the name of the parent resource: * projects/{project}/locations/{region}/clusters/{cluster_id}
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of available read-only VMs in this instance, including the standby for a PRIMARY instance.
func (o InstanceOutput) Nodes() NodeResponseArrayOutput {
	return o.ApplyT(func(v *Instance) NodeResponseArrayOutput { return v.Nodes }).(NodeResponseArrayOutput)
}

func (o InstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Configuration for query insights.
func (o InstanceOutput) QueryInsightsConfig() QueryInsightsInstanceConfigResponseOutput {
	return o.ApplyT(func(v *Instance) QueryInsightsInstanceConfigResponseOutput { return v.QueryInsightsConfig }).(QueryInsightsInstanceConfigResponseOutput)
}

// Read pool instance configuration. This is required if the value of instanceType is READ_POOL.
func (o InstanceOutput) ReadPoolConfig() ReadPoolConfigResponseOutput {
	return o.ApplyT(func(v *Instance) ReadPoolConfigResponseOutput { return v.ReadPoolConfig }).(ReadPoolConfigResponseOutput)
}

// Reconciling (https://google.aip.dev/128#reconciliation). Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.
func (o InstanceOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o InstanceOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The current serving state of the instance.
func (o InstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
func (o InstanceOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Update policy that will be applied during instance update. This field is not persisted when you update the instance. To use a non-default update policy, you must specify explicitly specify the value in each update request.
func (o InstanceOutput) UpdatePolicy() UpdatePolicyResponseOutput {
	return o.ApplyT(func(v *Instance) UpdatePolicyResponseOutput { return v.UpdatePolicy }).(UpdatePolicyResponseOutput)
}

// Update time stamp
func (o InstanceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// This is set for the read-write VM of the PRIMARY instance only.
func (o InstanceOutput) WritableNode() NodeResponseOutput {
	return o.ApplyT(func(v *Instance) NodeResponseOutput { return v.WritableNode }).(NodeResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterOutputType(InstanceOutput{})
}