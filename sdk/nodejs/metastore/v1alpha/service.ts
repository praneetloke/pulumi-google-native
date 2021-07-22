// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a metastore service in a project and location.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:metastore/v1alpha:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
     */
    public /*out*/ readonly artifactGcsUri!: pulumi.Output<string>;
    /**
     * The time when the metastore service was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The URI of the endpoint used to access the metastore service.
     */
    public /*out*/ readonly endpointUri!: pulumi.Output<string>;
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     */
    public readonly hiveMetastoreConfig!: pulumi.Output<outputs.metastore.v1alpha.HiveMetastoreConfigResponse>;
    /**
     * User-defined labels for the metastore service.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time.
     */
    public readonly maintenanceWindow!: pulumi.Output<outputs.metastore.v1alpha.MaintenanceWindowResponse>;
    /**
     * The setting that defines how metastore metadata should be integrated with external services and systems.
     */
    public readonly metadataIntegration!: pulumi.Output<outputs.metastore.v1alpha.MetadataIntegrationResponse>;
    /**
     * The metadata management activities of the metastore service.
     */
    public /*out*/ readonly metadataManagementActivity!: pulumi.Output<outputs.metastore.v1alpha.MetadataManagementActivityResponse>;
    /**
     * Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Immutable. The release channel of the service. If unspecified, defaults to STABLE.
     */
    public readonly releaseChannel!: pulumi.Output<string>;
    /**
     * The current state of the metastore service.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current state of the metastore service, if available.
     */
    public /*out*/ readonly stateMessage!: pulumi.Output<string>;
    /**
     * The tier of the service.
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * The globally unique resource identifier of the metastore service.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the metastore service was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            inputs["hiveMetastoreConfig"] = args ? args.hiveMetastoreConfig : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            inputs["metadataIntegration"] = args ? args.metadataIntegration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["releaseChannel"] = args ? args.releaseChannel : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["serviceId"] = args ? args.serviceId : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["artifactGcsUri"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["endpointUri"] = undefined /*out*/;
            inputs["metadataManagementActivity"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["stateMessage"] = undefined /*out*/;
            inputs["uid"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["artifactGcsUri"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["endpointUri"] = undefined /*out*/;
            inputs["hiveMetastoreConfig"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["maintenanceWindow"] = undefined /*out*/;
            inputs["metadataIntegration"] = undefined /*out*/;
            inputs["metadataManagementActivity"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["network"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["releaseChannel"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["stateMessage"] = undefined /*out*/;
            inputs["tier"] = undefined /*out*/;
            inputs["uid"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     */
    hiveMetastoreConfig?: pulumi.Input<inputs.metastore.v1alpha.HiveMetastoreConfigArgs>;
    /**
     * User-defined labels for the metastore service.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time.
     */
    maintenanceWindow?: pulumi.Input<inputs.metastore.v1alpha.MaintenanceWindowArgs>;
    /**
     * The setting that defines how metastore metadata should be integrated with external services and systems.
     */
    metadataIntegration?: pulumi.Input<inputs.metastore.v1alpha.MetadataIntegrationArgs>;
    /**
     * Immutable. The relative resource name of the metastore service, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}.
     */
    name?: pulumi.Input<string>;
    /**
     * Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
     */
    network?: pulumi.Input<string>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    port?: pulumi.Input<number>;
    project?: pulumi.Input<string>;
    /**
     * Immutable. The release channel of the service. If unspecified, defaults to STABLE.
     */
    releaseChannel?: pulumi.Input<enums.metastore.v1alpha.ServiceReleaseChannel>;
    requestId?: pulumi.Input<string>;
    serviceId: pulumi.Input<string>;
    /**
     * The tier of the service.
     */
    tier?: pulumi.Input<enums.metastore.v1alpha.ServiceTier>;
}
