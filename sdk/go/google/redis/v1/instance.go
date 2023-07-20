// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Redis instance based on the specified tier and memory size. By default, the instance is accessible from the project's [default network](https://cloud.google.com/vpc/docs/vpc). The creation is executed asynchronously and callers may check the returned operation to track its progress. Once the operation is completed the Redis instance will be fully functional. Completed longrunning.Operation will contain the new instance object in the response field. The returned operation is automatically deleted after a few hours, so there is no need to call DeleteOperation.
type Instance struct {
	pulumi.CustomResourceState

	// Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
	AlternativeLocationId pulumi.StringOutput `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled pulumi.BoolOutput `pulumi:"authEnabled"`
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork pulumi.StringOutput `pulumi:"authorizedNetwork"`
	// Optional. The available maintenance versions that an instance could update to.
	AvailableMaintenanceVersions pulumi.StringArrayOutput `pulumi:"availableMaintenanceVersions"`
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode pulumi.StringOutput `pulumi:"connectMode"`
	// The time the instance was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.
	CurrentLocationId pulumi.StringOutput `pulumi:"currentLocationId"`
	// Optional. The KMS key reference that the customer provides when trying to create the instance.
	CustomerManagedKey pulumi.StringOutput `pulumi:"customerManagedKey"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host pulumi.StringOutput `pulumi:"host"`
	// Required. The logical name of the Redis instance in the customer project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-40 characters. * Must end with a number or a letter. * Must be unique within the customer project / location
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Resource labels to represent user provided metadata
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
	MaintenancePolicy MaintenancePolicyResponseOutput `pulumi:"maintenancePolicy"`
	// Date and time of upcoming maintenance events which have been scheduled.
	MaintenanceSchedule MaintenanceScheduleResponseOutput `pulumi:"maintenanceSchedule"`
	// Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
	MaintenanceVersion pulumi.StringOutput `pulumi:"maintenanceVersion"`
	// Redis memory size in GiB.
	MemorySizeGb pulumi.IntOutput `pulumi:"memorySizeGb"`
	// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name pulumi.StringOutput `pulumi:"name"`
	// Info per node.
	Nodes NodeInfoResponseArrayOutput `pulumi:"nodes"`
	// Optional. Persistence configuration parameters
	PersistenceConfig PersistenceConfigResponseOutput `pulumi:"persistenceConfig"`
	// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
	PersistenceIamIdentity pulumi.StringOutput `pulumi:"persistenceIamIdentity"`
	// The port number of the exposed Redis endpoint.
	Port    pulumi.IntOutput    `pulumi:"port"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.
	ReadEndpoint pulumi.StringOutput `pulumi:"readEndpoint"`
	// The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.
	ReadEndpointPort pulumi.IntOutput `pulumi:"readEndpointPort"`
	// Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
	ReadReplicasMode pulumi.StringOutput `pulumi:"readReplicasMode"`
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs pulumi.StringMapOutput `pulumi:"redisConfigs"`
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
	RedisVersion pulumi.StringOutput `pulumi:"redisVersion"`
	// Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
	ReplicaCount pulumi.IntOutput `pulumi:"replicaCount"`
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
	ReservedIpRange pulumi.StringOutput `pulumi:"reservedIpRange"`
	// Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
	SecondaryIpRange pulumi.StringOutput `pulumi:"secondaryIpRange"`
	// List of server CA certificates for the instance.
	ServerCaCerts TlsCertificateResponseArrayOutput `pulumi:"serverCaCerts"`
	// The current state of this instance.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current status of this instance, if available.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Optional. reasons that causes instance in "SUSPENDED" state.
	SuspensionReasons pulumi.StringArrayOutput `pulumi:"suspensionReasons"`
	// The service tier of the instance.
	Tier pulumi.StringOutput `pulumi:"tier"`
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode pulumi.StringOutput `pulumi:"transitEncryptionMode"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.MemorySizeGb == nil {
		return nil, errors.New("invalid value for required argument 'MemorySizeGb'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"instanceId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("google-native:redis/v1:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:redis/v1:Instance", name, id, state, &resource, opts...)
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
	// Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
	AlternativeLocationId *string `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled *bool `pulumi:"authEnabled"`
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Optional. The available maintenance versions that an instance could update to.
	AvailableMaintenanceVersions []string `pulumi:"availableMaintenanceVersions"`
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode *InstanceConnectMode `pulumi:"connectMode"`
	// Optional. The KMS key reference that the customer provides when trying to create the instance.
	CustomerManagedKey *string `pulumi:"customerManagedKey"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName *string `pulumi:"displayName"`
	// Required. The logical name of the Redis instance in the customer project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-40 characters. * Must end with a number or a letter. * Must be unique within the customer project / location
	InstanceId string `pulumi:"instanceId"`
	// Resource labels to represent user provided metadata
	Labels map[string]string `pulumi:"labels"`
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.
	Location *string `pulumi:"location"`
	// Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
	MaintenancePolicy *MaintenancePolicy `pulumi:"maintenancePolicy"`
	// Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
	MaintenanceVersion *string `pulumi:"maintenanceVersion"`
	// Redis memory size in GiB.
	MemorySizeGb int `pulumi:"memorySizeGb"`
	// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name *string `pulumi:"name"`
	// Optional. Persistence configuration parameters
	PersistenceConfig *PersistenceConfig `pulumi:"persistenceConfig"`
	Project           *string            `pulumi:"project"`
	// Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
	ReadReplicasMode *InstanceReadReplicasMode `pulumi:"readReplicasMode"`
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs map[string]string `pulumi:"redisConfigs"`
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
	RedisVersion *string `pulumi:"redisVersion"`
	// Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
	ReplicaCount *int `pulumi:"replicaCount"`
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
	ReservedIpRange *string `pulumi:"reservedIpRange"`
	// Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
	SecondaryIpRange *string `pulumi:"secondaryIpRange"`
	// Optional. reasons that causes instance in "SUSPENDED" state.
	SuspensionReasons []InstanceSuspensionReasonsItem `pulumi:"suspensionReasons"`
	// The service tier of the instance.
	Tier InstanceTier `pulumi:"tier"`
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode *InstanceTransitEncryptionMode `pulumi:"transitEncryptionMode"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
	AlternativeLocationId pulumi.StringPtrInput
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled pulumi.BoolPtrInput
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork pulumi.StringPtrInput
	// Optional. The available maintenance versions that an instance could update to.
	AvailableMaintenanceVersions pulumi.StringArrayInput
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode InstanceConnectModePtrInput
	// Optional. The KMS key reference that the customer provides when trying to create the instance.
	CustomerManagedKey pulumi.StringPtrInput
	// An arbitrary and optional user-provided name for the instance.
	DisplayName pulumi.StringPtrInput
	// Required. The logical name of the Redis instance in the customer project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-40 characters. * Must end with a number or a letter. * Must be unique within the customer project / location
	InstanceId pulumi.StringInput
	// Resource labels to represent user provided metadata
	Labels pulumi.StringMapInput
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.
	Location pulumi.StringPtrInput
	// Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
	MaintenancePolicy MaintenancePolicyPtrInput
	// Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
	MaintenanceVersion pulumi.StringPtrInput
	// Redis memory size in GiB.
	MemorySizeGb pulumi.IntInput
	// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name pulumi.StringPtrInput
	// Optional. Persistence configuration parameters
	PersistenceConfig PersistenceConfigPtrInput
	Project           pulumi.StringPtrInput
	// Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
	ReadReplicasMode InstanceReadReplicasModePtrInput
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs pulumi.StringMapInput
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
	RedisVersion pulumi.StringPtrInput
	// Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
	ReplicaCount pulumi.IntPtrInput
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
	ReservedIpRange pulumi.StringPtrInput
	// Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
	SecondaryIpRange pulumi.StringPtrInput
	// Optional. reasons that causes instance in "SUSPENDED" state.
	SuspensionReasons InstanceSuspensionReasonsItemArrayInput
	// The service tier of the instance.
	Tier InstanceTierInput
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode InstanceTransitEncryptionModePtrInput
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

// Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
func (o InstanceOutput) AlternativeLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AlternativeLocationId }).(pulumi.StringOutput)
}

// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
func (o InstanceOutput) AuthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.AuthEnabled }).(pulumi.BoolOutput)
}

// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
func (o InstanceOutput) AuthorizedNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AuthorizedNetwork }).(pulumi.StringOutput)
}

// Optional. The available maintenance versions that an instance could update to.
func (o InstanceOutput) AvailableMaintenanceVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.AvailableMaintenanceVersions }).(pulumi.StringArrayOutput)
}

// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
func (o InstanceOutput) ConnectMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ConnectMode }).(pulumi.StringOutput)
}

// The time the instance was created.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.
func (o InstanceOutput) CurrentLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CurrentLocationId }).(pulumi.StringOutput)
}

// Optional. The KMS key reference that the customer provides when trying to create the instance.
func (o InstanceOutput) CustomerManagedKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CustomerManagedKey }).(pulumi.StringOutput)
}

// An arbitrary and optional user-provided name for the instance.
func (o InstanceOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
func (o InstanceOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Required. The logical name of the Redis instance in the customer project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-40 characters. * Must end with a number or a letter. * Must be unique within the customer project / location
func (o InstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Resource labels to represent user provided metadata
func (o InstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
func (o InstanceOutput) MaintenancePolicy() MaintenancePolicyResponseOutput {
	return o.ApplyT(func(v *Instance) MaintenancePolicyResponseOutput { return v.MaintenancePolicy }).(MaintenancePolicyResponseOutput)
}

// Date and time of upcoming maintenance events which have been scheduled.
func (o InstanceOutput) MaintenanceSchedule() MaintenanceScheduleResponseOutput {
	return o.ApplyT(func(v *Instance) MaintenanceScheduleResponseOutput { return v.MaintenanceSchedule }).(MaintenanceScheduleResponseOutput)
}

// Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
func (o InstanceOutput) MaintenanceVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.MaintenanceVersion }).(pulumi.StringOutput)
}

// Redis memory size in GiB.
func (o InstanceOutput) MemorySizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.MemorySizeGb }).(pulumi.IntOutput)
}

// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Info per node.
func (o InstanceOutput) Nodes() NodeInfoResponseArrayOutput {
	return o.ApplyT(func(v *Instance) NodeInfoResponseArrayOutput { return v.Nodes }).(NodeInfoResponseArrayOutput)
}

// Optional. Persistence configuration parameters
func (o InstanceOutput) PersistenceConfig() PersistenceConfigResponseOutput {
	return o.ApplyT(func(v *Instance) PersistenceConfigResponseOutput { return v.PersistenceConfig }).(PersistenceConfigResponseOutput)
}

// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
func (o InstanceOutput) PersistenceIamIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PersistenceIamIdentity }).(pulumi.StringOutput)
}

// The port number of the exposed Redis endpoint.
func (o InstanceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o InstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.
func (o InstanceOutput) ReadEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ReadEndpoint }).(pulumi.StringOutput)
}

// The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.
func (o InstanceOutput) ReadEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.ReadEndpointPort }).(pulumi.IntOutput)
}

// Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
func (o InstanceOutput) ReadReplicasMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ReadReplicasMode }).(pulumi.StringOutput)
}

// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
func (o InstanceOutput) RedisConfigs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.RedisConfigs }).(pulumi.StringMapOutput)
}

// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
func (o InstanceOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RedisVersion }).(pulumi.StringOutput)
}

// Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
func (o InstanceOutput) ReplicaCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.ReplicaCount }).(pulumi.IntOutput)
}

// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
func (o InstanceOutput) ReservedIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ReservedIpRange }).(pulumi.StringOutput)
}

// Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
func (o InstanceOutput) SecondaryIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SecondaryIpRange }).(pulumi.StringOutput)
}

// List of server CA certificates for the instance.
func (o InstanceOutput) ServerCaCerts() TlsCertificateResponseArrayOutput {
	return o.ApplyT(func(v *Instance) TlsCertificateResponseArrayOutput { return v.ServerCaCerts }).(TlsCertificateResponseArrayOutput)
}

// The current state of this instance.
func (o InstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current status of this instance, if available.
func (o InstanceOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

// Optional. reasons that causes instance in "SUSPENDED" state.
func (o InstanceOutput) SuspensionReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SuspensionReasons }).(pulumi.StringArrayOutput)
}

// The service tier of the instance.
func (o InstanceOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
func (o InstanceOutput) TransitEncryptionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TransitEncryptionMode }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterOutputType(InstanceOutput{})
}
