// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an Experiment in the specified Environment.
 */
export class AgentEnvironmentExperiment extends pulumi.CustomResource {
    /**
     * Get an existing AgentEnvironmentExperiment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentEnvironmentExperiment {
        return new AgentEnvironmentExperiment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3:AgentEnvironmentExperiment';

    /**
     * Returns true if the given object is an instance of AgentEnvironmentExperiment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentEnvironmentExperiment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentEnvironmentExperiment.__pulumiType;
    }

    /**
     * Creation time of this experiment.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * The definition of the experiment.
     */
    public readonly definition!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentDefinitionResponse>;
    /**
     * The human-readable description of the experiment.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * End time of this experiment.
     */
    public readonly endTime!: pulumi.Output<string>;
    /**
     * Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
     */
    public readonly experimentLength!: pulumi.Output<string>;
    /**
     * Last update time of this experiment.
     */
    public readonly lastUpdateTime!: pulumi.Output<string>;
    /**
     * The name of the experiment. Format: projects//locations//agents//environments//experiments/..
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Inference result of the experiment.
     */
    public readonly result!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentResultResponse>;
    /**
     * Start time of this experiment.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The history of updates to the experiment variants.
     */
    public readonly variantsHistory!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3VariantsHistoryResponse[]>;

    /**
     * Create a AgentEnvironmentExperiment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentEnvironmentExperimentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.experimentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'experimentId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["agentId"] = args ? args.agentId : undefined;
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["definition"] = args ? args.definition : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["experimentId"] = args ? args.experimentId : undefined;
            inputs["experimentLength"] = args ? args.experimentLength : undefined;
            inputs["lastUpdateTime"] = args ? args.lastUpdateTime : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["result"] = args ? args.result : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["variantsHistory"] = args ? args.variantsHistory : undefined;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["definition"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["experimentLength"] = undefined /*out*/;
            inputs["lastUpdateTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["result"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["variantsHistory"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AgentEnvironmentExperiment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentEnvironmentExperiment resource.
 */
export interface AgentEnvironmentExperimentArgs {
    readonly agentId: pulumi.Input<string>;
    /**
     * Creation time of this experiment.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * The definition of the experiment.
     */
    readonly definition?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentDefinitionArgs>;
    /**
     * The human-readable description of the experiment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * End time of this experiment.
     */
    readonly endTime?: pulumi.Input<string>;
    readonly environmentId: pulumi.Input<string>;
    readonly experimentId: pulumi.Input<string>;
    /**
     * Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
     */
    readonly experimentLength?: pulumi.Input<string>;
    /**
     * Last update time of this experiment.
     */
    readonly lastUpdateTime?: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    /**
     * The name of the experiment. Format: projects//locations//agents//environments//experiments/..
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Inference result of the experiment.
     */
    readonly result?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentResultArgs>;
    /**
     * Start time of this experiment.
     */
    readonly startTime?: pulumi.Input<string>;
    /**
     * The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The history of updates to the experiment variants.
     */
    readonly variantsHistory?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3VariantsHistoryArgs>[]>;
}
