// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create an OS Config patch deployment.
 */
export class PatchDeployment extends pulumi.CustomResource {
    /**
     * Get an existing PatchDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PatchDeployment {
        return new PatchDeployment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:osconfig/v1:PatchDeployment';

    /**
     * Returns true if the given object is an instance of PatchDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PatchDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PatchDeployment.__pulumiType;
    }

    /**
     * Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Duration of the patch. After the duration ends, the patch times out.
     */
    public readonly duration!: pulumi.Output<string>;
    /**
     * Required. VM instances to patch.
     */
    public readonly instanceFilter!: pulumi.Output<outputs.osconfig.v1.PatchInstanceFilterResponse>;
    /**
     * The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
     */
    public /*out*/ readonly lastExecuteTime!: pulumi.Output<string>;
    /**
     * Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Required. Schedule a one-time execution.
     */
    public readonly oneTimeSchedule!: pulumi.Output<outputs.osconfig.v1.OneTimeScheduleResponse>;
    /**
     * Optional. Patch configuration that is applied.
     */
    public readonly patchConfig!: pulumi.Output<outputs.osconfig.v1.PatchConfigResponse>;
    /**
     * Required. Schedule recurring executions.
     */
    public readonly recurringSchedule!: pulumi.Output<outputs.osconfig.v1.RecurringScheduleResponse>;
    /**
     * Optional. Rollout strategy of the patch job.
     */
    public readonly rollout!: pulumi.Output<outputs.osconfig.v1.PatchRolloutResponse>;
    /**
     * Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a PatchDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PatchDeploymentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.patchDeploymentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'patchDeploymentId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["duration"] = args ? args.duration : undefined;
            inputs["instanceFilter"] = args ? args.instanceFilter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["oneTimeSchedule"] = args ? args.oneTimeSchedule : undefined;
            inputs["patchConfig"] = args ? args.patchConfig : undefined;
            inputs["patchDeploymentId"] = args ? args.patchDeploymentId : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["recurringSchedule"] = args ? args.recurringSchedule : undefined;
            inputs["rollout"] = args ? args.rollout : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["lastExecuteTime"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["duration"] = undefined /*out*/;
            inputs["instanceFilter"] = undefined /*out*/;
            inputs["lastExecuteTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["oneTimeSchedule"] = undefined /*out*/;
            inputs["patchConfig"] = undefined /*out*/;
            inputs["recurringSchedule"] = undefined /*out*/;
            inputs["rollout"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PatchDeployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PatchDeployment resource.
 */
export interface PatchDeploymentArgs {
    /**
     * Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Optional. Duration of the patch. After the duration ends, the patch times out.
     */
    readonly duration?: pulumi.Input<string>;
    /**
     * Required. VM instances to patch.
     */
    readonly instanceFilter?: pulumi.Input<inputs.osconfig.v1.PatchInstanceFilterArgs>;
    /**
     * Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Required. Schedule a one-time execution.
     */
    readonly oneTimeSchedule?: pulumi.Input<inputs.osconfig.v1.OneTimeScheduleArgs>;
    /**
     * Optional. Patch configuration that is applied.
     */
    readonly patchConfig?: pulumi.Input<inputs.osconfig.v1.PatchConfigArgs>;
    readonly patchDeploymentId: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Required. Schedule recurring executions.
     */
    readonly recurringSchedule?: pulumi.Input<inputs.osconfig.v1.RecurringScheduleArgs>;
    /**
     * Optional. Rollout strategy of the patch job.
     */
    readonly rollout?: pulumi.Input<inputs.osconfig.v1.PatchRolloutArgs>;
}
