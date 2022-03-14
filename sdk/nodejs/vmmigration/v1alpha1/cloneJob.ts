// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Initiates a Clone of a specific migrating VM.
 * Auto-naming is currently not supported for this resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class CloneJob extends pulumi.CustomResource {
    /**
     * Get an existing CloneJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CloneJob {
        return new CloneJob(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:vmmigration/v1alpha1:CloneJob';

    /**
     * Returns true if the given object is an instance of CloneJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloneJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloneJob.__pulumiType;
    }

    /**
     * Details of the target VM in Compute Engine.
     */
    public /*out*/ readonly computeEngineTargetDetails!: pulumi.Output<outputs.vmmigration.v1alpha1.ComputeEngineTargetDetailsResponse>;
    /**
     * The time the clone job was created (as an API call, not when it was actually created in the target).
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Provides details for the errors that led to the Clone Job's state.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.vmmigration.v1alpha1.StatusResponse>;
    /**
     * The name of the clone.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * State of the clone job.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The time the state was last updated.
     */
    public /*out*/ readonly stateTime!: pulumi.Output<string>;

    /**
     * Create a CloneJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloneJobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloneJobId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloneJobId'");
            }
            if ((!args || args.migratingVmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'migratingVmId'");
            }
            if ((!args || args.sourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceId'");
            }
            resourceInputs["cloneJobId"] = args ? args.cloneJobId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["migratingVmId"] = args ? args.migratingVmId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["sourceId"] = args ? args.sourceId : undefined;
            resourceInputs["computeEngineTargetDetails"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateTime"] = undefined /*out*/;
        } else {
            resourceInputs["computeEngineTargetDetails"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloneJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CloneJob resource.
 */
export interface CloneJobArgs {
    /**
     * Required. The clone job identifier.
     */
    cloneJobId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    migratingVmId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
