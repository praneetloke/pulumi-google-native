// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Redis instance based on the specified tier and memory size. By default, the instance is accessible from the project's [default network](https://cloud.google.com/vpc/docs/vpc). The creation is executed asynchronously and callers may check the returned operation to track its progress. Once the operation is completed the Redis instance will be fully functional. Completed longrunning.Operation will contain the new instance object in the response field. The returned operation is automatically deleted after a few hours, so there is no need to call DeleteOperation.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:redis/v1:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
     */
    public readonly alternativeLocationId!: pulumi.Output<string>;
    /**
     * Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
     */
    public readonly authEnabled!: pulumi.Output<boolean>;
    /**
     * Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
     */
    public readonly authorizedNetwork!: pulumi.Output<string>;
    /**
     * Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
     */
    public readonly connectMode!: pulumi.Output<string>;
    /**
     * The time the instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the location_id provided by the user at creation time. For Standard Tier instances, this can be either location_id or alternative_location_id and can change after a failover event.
     */
    public /*out*/ readonly currentLocationId!: pulumi.Output<string>;
    /**
     * An arbitrary and optional user-provided name for the instance.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Resource labels to represent user provided metadata
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
     */
    public readonly maintenancePolicy!: pulumi.Output<outputs.redis.v1.MaintenancePolicyResponse>;
    /**
     * Date and time of upcoming maintenance events which have been scheduled.
     */
    public /*out*/ readonly maintenanceSchedule!: pulumi.Output<outputs.redis.v1.MaintenanceScheduleResponse>;
    /**
     * Redis memory size in GiB.
     */
    public readonly memorySizeGb!: pulumi.Output<number>;
    /**
     * Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
     */
    public /*out*/ readonly persistenceIamIdentity!: pulumi.Output<string>;
    /**
     * The port number of the exposed Redis endpoint.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
     */
    public readonly redisConfigs!: pulumi.Output<{[key: string]: string}>;
    /**
     * Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
     */
    public readonly redisVersion!: pulumi.Output<string>;
    /**
     * Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
     */
    public readonly reservedIpRange!: pulumi.Output<string>;
    /**
     * List of server CA certificates for the instance.
     */
    public /*out*/ readonly serverCaCerts!: pulumi.Output<outputs.redis.v1.TlsCertificateResponse[]>;
    /**
     * The current state of this instance.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current status of this instance, if available.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * The service tier of the instance.
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
     */
    public readonly transitEncryptionMode!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.memorySizeGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memorySizeGb'");
            }
            if ((!args || args.tier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tier'");
            }
            inputs["alternativeLocationId"] = args ? args.alternativeLocationId : undefined;
            inputs["authEnabled"] = args ? args.authEnabled : undefined;
            inputs["authorizedNetwork"] = args ? args.authorizedNetwork : undefined;
            inputs["connectMode"] = args ? args.connectMode : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            inputs["memorySizeGb"] = args ? args.memorySizeGb : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["redisConfigs"] = args ? args.redisConfigs : undefined;
            inputs["redisVersion"] = args ? args.redisVersion : undefined;
            inputs["reservedIpRange"] = args ? args.reservedIpRange : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["transitEncryptionMode"] = args ? args.transitEncryptionMode : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["currentLocationId"] = undefined /*out*/;
            inputs["host"] = undefined /*out*/;
            inputs["maintenanceSchedule"] = undefined /*out*/;
            inputs["persistenceIamIdentity"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["serverCaCerts"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
        } else {
            inputs["alternativeLocationId"] = undefined /*out*/;
            inputs["authEnabled"] = undefined /*out*/;
            inputs["authorizedNetwork"] = undefined /*out*/;
            inputs["connectMode"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["currentLocationId"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["host"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["maintenancePolicy"] = undefined /*out*/;
            inputs["maintenanceSchedule"] = undefined /*out*/;
            inputs["memorySizeGb"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["persistenceIamIdentity"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["redisConfigs"] = undefined /*out*/;
            inputs["redisVersion"] = undefined /*out*/;
            inputs["reservedIpRange"] = undefined /*out*/;
            inputs["serverCaCerts"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
            inputs["tier"] = undefined /*out*/;
            inputs["transitEncryptionMode"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Optional. Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in location_id.
     */
    alternativeLocationId?: pulumi.Input<string>;
    /**
     * Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
     */
    authEnabled?: pulumi.Input<boolean>;
    /**
     * Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
     */
    authorizedNetwork?: pulumi.Input<string>;
    /**
     * Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
     */
    connectMode?: pulumi.Input<enums.redis.v1.InstanceConnectMode>;
    /**
     * An arbitrary and optional user-provided name for the instance.
     */
    displayName?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    /**
     * Resource labels to represent user provided metadata
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If alternative_location_id is also provided, it must be different from location_id.
     */
    location?: pulumi.Input<string>;
    /**
     * Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
     */
    maintenancePolicy?: pulumi.Input<inputs.redis.v1.MaintenancePolicyArgs>;
    /**
     * Redis memory size in GiB.
     */
    memorySizeGb: pulumi.Input<number>;
    /**
     * Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
     */
    redisConfigs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
     */
    redisVersion?: pulumi.Input<string>;
    /**
     * Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29.
     */
    reservedIpRange?: pulumi.Input<string>;
    /**
     * The service tier of the instance.
     */
    tier: pulumi.Input<enums.redis.v1.InstanceTier>;
    /**
     * Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
     */
    transitEncryptionMode?: pulumi.Input<enums.redis.v1.InstanceTransitEncryptionMode>;
}
