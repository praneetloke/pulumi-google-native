// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new DatacenterConnector in a given Source.
 * Auto-naming is currently not supported for this resource.
 */
export class DatacenterConnector extends pulumi.CustomResource {
    /**
     * Get an existing DatacenterConnector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatacenterConnector {
        return new DatacenterConnector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:vmmigration/v1:DatacenterConnector';

    /**
     * Returns true if the given object is an instance of DatacenterConnector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatacenterConnector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatacenterConnector.__pulumiType;
    }

    /**
     * Appliance OVA version. This is the OVA which is manually installed by the user and contains the infrastructure for the automatically updatable components on the appliance.
     */
    public /*out*/ readonly applianceInfrastructureVersion!: pulumi.Output<string>;
    /**
     * Appliance last installed update bundle version. This is the version of the automatically updatable components on the appliance.
     */
    public /*out*/ readonly applianceSoftwareVersion!: pulumi.Output<string>;
    /**
     * The available versions for updating this appliance.
     */
    public /*out*/ readonly availableVersions!: pulumi.Output<outputs.vmmigration.v1.AvailableUpdatesResponse>;
    /**
     * The communication channel between the datacenter connector and GCP.
     */
    public /*out*/ readonly bucket!: pulumi.Output<string>;
    /**
     * The time the connector was created (as an API call, not when it was actually installed).
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Provides details on the state of the Datacenter Connector in case of an error.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.vmmigration.v1.StatusResponse>;
    /**
     * The connector's name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
     */
    public readonly registrationId!: pulumi.Output<string>;
    /**
     * The service account to use in the connector when communicating with the cloud.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * State of the DatacenterConnector, as determined by the health checks.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The time the state was last set.
     */
    public /*out*/ readonly stateTime!: pulumi.Output<string>;
    /**
     * The last time the connector was updated with an API call.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The status of the current / last upgradeAppliance operation.
     */
    public /*out*/ readonly upgradeStatus!: pulumi.Output<outputs.vmmigration.v1.UpgradeStatusResponse>;
    /**
     * The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a DatacenterConnector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatacenterConnectorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datacenterConnectorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datacenterConnectorId'");
            }
            if ((!args || args.sourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceId'");
            }
            resourceInputs["datacenterConnectorId"] = args ? args.datacenterConnectorId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["registrationId"] = args ? args.registrationId : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["sourceId"] = args ? args.sourceId : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["applianceInfrastructureVersion"] = undefined /*out*/;
            resourceInputs["applianceSoftwareVersion"] = undefined /*out*/;
            resourceInputs["availableVersions"] = undefined /*out*/;
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["upgradeStatus"] = undefined /*out*/;
        } else {
            resourceInputs["applianceInfrastructureVersion"] = undefined /*out*/;
            resourceInputs["applianceSoftwareVersion"] = undefined /*out*/;
            resourceInputs["availableVersions"] = undefined /*out*/;
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["registrationId"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["upgradeStatus"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatacenterConnector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatacenterConnector resource.
 */
export interface DatacenterConnectorArgs {
    /**
     * Required. The datacenterConnector identifier.
     */
    datacenterConnectorId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Immutable. A unique key for this connector. This key is internal to the OVA connector and is supplied with its creation during the registration process and can not be modified.
     */
    registrationId?: pulumi.Input<string>;
    /**
     * A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * The service account to use in the connector when communicating with the cloud.
     */
    serviceAccount?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
    /**
     * The version running in the DatacenterConnector. This is supplied by the OVA connector during the registration process and can not be modified.
     */
    version?: pulumi.Input<string>;
}
