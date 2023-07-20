// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a specific Redis instance.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:redis/v1:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupInstanceResult struct {
	// Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
	AlternativeLocationId string `pulumi:"alternativeLocationId"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	AuthEnabled bool `pulumi:"authEnabled"`
	// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	AuthorizedNetwork string `pulumi:"authorizedNetwork"`
	// Optional. The available maintenance versions that an instance could update to.
	AvailableMaintenanceVersions []string `pulumi:"availableMaintenanceVersions"`
	// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	ConnectMode string `pulumi:"connectMode"`
	// The time the instance was created.
	CreateTime string `pulumi:"createTime"`
	// The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.
	CurrentLocationId string `pulumi:"currentLocationId"`
	// Optional. The KMS key reference that the customer provides when trying to create the instance.
	CustomerManagedKey string `pulumi:"customerManagedKey"`
	// An arbitrary and optional user-provided name for the instance.
	DisplayName string `pulumi:"displayName"`
	// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Host string `pulumi:"host"`
	// Resource labels to represent user provided metadata
	Labels map[string]string `pulumi:"labels"`
	// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.
	Location string `pulumi:"location"`
	// Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
	MaintenancePolicy MaintenancePolicyResponse `pulumi:"maintenancePolicy"`
	// Date and time of upcoming maintenance events which have been scheduled.
	MaintenanceSchedule MaintenanceScheduleResponse `pulumi:"maintenanceSchedule"`
	// Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
	MaintenanceVersion string `pulumi:"maintenanceVersion"`
	// Redis memory size in GiB.
	MemorySizeGb int `pulumi:"memorySizeGb"`
	// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Name string `pulumi:"name"`
	// Info per node.
	Nodes []NodeInfoResponse `pulumi:"nodes"`
	// Optional. Persistence configuration parameters
	PersistenceConfig PersistenceConfigResponse `pulumi:"persistenceConfig"`
	// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
	PersistenceIamIdentity string `pulumi:"persistenceIamIdentity"`
	// The port number of the exposed Redis endpoint.
	Port int `pulumi:"port"`
	// Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.
	ReadEndpoint string `pulumi:"readEndpoint"`
	// The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.
	ReadEndpointPort int `pulumi:"readEndpointPort"`
	// Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
	ReadReplicasMode string `pulumi:"readReplicasMode"`
	// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	RedisConfigs map[string]string `pulumi:"redisConfigs"`
	// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
	RedisVersion string `pulumi:"redisVersion"`
	// Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
	ReplicaCount int `pulumi:"replicaCount"`
	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
	ReservedIpRange string `pulumi:"reservedIpRange"`
	// Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
	SecondaryIpRange string `pulumi:"secondaryIpRange"`
	// List of server CA certificates for the instance.
	ServerCaCerts []TlsCertificateResponse `pulumi:"serverCaCerts"`
	// The current state of this instance.
	State string `pulumi:"state"`
	// Additional information about the current status of this instance, if available.
	StatusMessage string `pulumi:"statusMessage"`
	// Optional. reasons that causes instance in "SUSPENDED" state.
	SuspensionReasons []string `pulumi:"suspensionReasons"`
	// The service tier of the instance.
	Tier string `pulumi:"tier"`
	// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	TransitEncryptionMode string `pulumi:"transitEncryptionMode"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
func (o LookupInstanceResultOutput) AlternativeLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.AlternativeLocationId }).(pulumi.StringOutput)
}

// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
func (o LookupInstanceResultOutput) AuthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.AuthEnabled }).(pulumi.BoolOutput)
}

// Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
func (o LookupInstanceResultOutput) AuthorizedNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.AuthorizedNetwork }).(pulumi.StringOutput)
}

// Optional. The available maintenance versions that an instance could update to.
func (o LookupInstanceResultOutput) AvailableMaintenanceVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.AvailableMaintenanceVersions }).(pulumi.StringArrayOutput)
}

// Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
func (o LookupInstanceResultOutput) ConnectMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ConnectMode }).(pulumi.StringOutput)
}

// The time the instance was created.
func (o LookupInstanceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.
func (o LookupInstanceResultOutput) CurrentLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CurrentLocationId }).(pulumi.StringOutput)
}

// Optional. The KMS key reference that the customer provides when trying to create the instance.
func (o LookupInstanceResultOutput) CustomerManagedKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CustomerManagedKey }).(pulumi.StringOutput)
}

// An arbitrary and optional user-provided name for the instance.
func (o LookupInstanceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
func (o LookupInstanceResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Host }).(pulumi.StringOutput)
}

// Resource labels to represent user provided metadata
func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.
func (o LookupInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
func (o LookupInstanceResultOutput) MaintenancePolicy() MaintenancePolicyResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) MaintenancePolicyResponse { return v.MaintenancePolicy }).(MaintenancePolicyResponseOutput)
}

// Date and time of upcoming maintenance events which have been scheduled.
func (o LookupInstanceResultOutput) MaintenanceSchedule() MaintenanceScheduleResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) MaintenanceScheduleResponse { return v.MaintenanceSchedule }).(MaintenanceScheduleResponseOutput)
}

// Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
func (o LookupInstanceResultOutput) MaintenanceVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.MaintenanceVersion }).(pulumi.StringOutput)
}

// Redis memory size in GiB.
func (o LookupInstanceResultOutput) MemorySizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.MemorySizeGb }).(pulumi.IntOutput)
}

// Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Info per node.
func (o LookupInstanceResultOutput) Nodes() NodeInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []NodeInfoResponse { return v.Nodes }).(NodeInfoResponseArrayOutput)
}

// Optional. Persistence configuration parameters
func (o LookupInstanceResultOutput) PersistenceConfig() PersistenceConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) PersistenceConfigResponse { return v.PersistenceConfig }).(PersistenceConfigResponseOutput)
}

// Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
func (o LookupInstanceResultOutput) PersistenceIamIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.PersistenceIamIdentity }).(pulumi.StringOutput)
}

// The port number of the exposed Redis endpoint.
func (o LookupInstanceResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.Port }).(pulumi.IntOutput)
}

// Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.
func (o LookupInstanceResultOutput) ReadEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ReadEndpoint }).(pulumi.StringOutput)
}

// The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.
func (o LookupInstanceResultOutput) ReadEndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.ReadEndpointPort }).(pulumi.IntOutput)
}

// Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
func (o LookupInstanceResultOutput) ReadReplicasMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ReadReplicasMode }).(pulumi.StringOutput)
}

// Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
func (o LookupInstanceResultOutput) RedisConfigs() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.RedisConfigs }).(pulumi.StringMapOutput)
}

// Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
func (o LookupInstanceResultOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.RedisVersion }).(pulumi.StringOutput)
}

// Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
func (o LookupInstanceResultOutput) ReplicaCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.ReplicaCount }).(pulumi.IntOutput)
}

// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
func (o LookupInstanceResultOutput) ReservedIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ReservedIpRange }).(pulumi.StringOutput)
}

// Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
func (o LookupInstanceResultOutput) SecondaryIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.SecondaryIpRange }).(pulumi.StringOutput)
}

// List of server CA certificates for the instance.
func (o LookupInstanceResultOutput) ServerCaCerts() TlsCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []TlsCertificateResponse { return v.ServerCaCerts }).(TlsCertificateResponseArrayOutput)
}

// The current state of this instance.
func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current status of this instance, if available.
func (o LookupInstanceResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

// Optional. reasons that causes instance in "SUSPENDED" state.
func (o LookupInstanceResultOutput) SuspensionReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.SuspensionReasons }).(pulumi.StringArrayOutput)
}

// The service tier of the instance.
func (o LookupInstanceResultOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Tier }).(pulumi.StringOutput)
}

// Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
func (o LookupInstanceResultOutput) TransitEncryptionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.TransitEncryptionMode }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
