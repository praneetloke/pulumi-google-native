// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the details of a specific Redis instance.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:redis/v1:getInstance", {
        "instanceId": args.instanceId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetInstanceArgs {
    instanceId: string;
    location: string;
    project?: string;
}

export interface GetInstanceResult {
    /**
     * Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
     */
    readonly alternativeLocationId: string;
    /**
     * Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
     */
    readonly authEnabled: boolean;
    /**
     * Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
     */
    readonly authorizedNetwork: string;
    /**
     * Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
     */
    readonly connectMode: string;
    /**
     * The time the instance was created.
     */
    readonly createTime: string;
    /**
     * The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.
     */
    readonly currentLocationId: string;
    /**
     * Optional. The KMS key reference that the customer provides when trying to create the instance.
     */
    readonly customerManagedKey: string;
    /**
     * An arbitrary and optional user-provided name for the instance.
     */
    readonly displayName: string;
    /**
     * Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
     */
    readonly host: string;
    /**
     * Resource labels to represent user provided metadata
     */
    readonly labels: {[key: string]: string};
    /**
     * Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.
     */
    readonly location: string;
    /**
     * Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
     */
    readonly maintenancePolicy: outputs.redis.v1.MaintenancePolicyResponse;
    /**
     * Date and time of upcoming maintenance events which have been scheduled.
     */
    readonly maintenanceSchedule: outputs.redis.v1.MaintenanceScheduleResponse;
    /**
     * Redis memory size in GiB.
     */
    readonly memorySizeGb: number;
    /**
     * Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
     */
    readonly name: string;
    /**
     * Info per node.
     */
    readonly nodes: outputs.redis.v1.NodeInfoResponse[];
    /**
     * Optional. Persistence configuration parameters
     */
    readonly persistenceConfig: outputs.redis.v1.PersistenceConfigResponse;
    /**
     * Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
     */
    readonly persistenceIamIdentity: string;
    /**
     * The port number of the exposed Redis endpoint.
     */
    readonly port: number;
    /**
     * Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.
     */
    readonly readEndpoint: string;
    /**
     * The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.
     */
    readonly readEndpointPort: number;
    /**
     * Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
     */
    readonly readReplicasMode: string;
    /**
     * Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
     */
    readonly redisConfigs: {[key: string]: string};
    /**
     * Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
     */
    readonly redisVersion: string;
    /**
     * Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
     */
    readonly replicaCount: number;
    /**
     * Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
     */
    readonly reservedIpRange: string;
    /**
     * Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
     */
    readonly secondaryIpRange: string;
    /**
     * List of server CA certificates for the instance.
     */
    readonly serverCaCerts: outputs.redis.v1.TlsCertificateResponse[];
    /**
     * The current state of this instance.
     */
    readonly state: string;
    /**
     * Additional information about the current status of this instance, if available.
     */
    readonly statusMessage: string;
    /**
     * Optional. reasons that causes instance in "SUSPENDED" state.
     */
    readonly suspensionReasons: string[];
    /**
     * The service tier of the instance.
     */
    readonly tier: string;
    /**
     * Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
     */
    readonly transitEncryptionMode: string;
}

export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

export interface GetInstanceOutputArgs {
    instanceId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
