// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an TransitionRouteGroup in the specified flow.
 */
export class AgentFlowTransitionRouteGroup extends pulumi.CustomResource {
    /**
     * Get an existing AgentFlowTransitionRouteGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentFlowTransitionRouteGroup {
        return new AgentFlowTransitionRouteGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3:AgentFlowTransitionRouteGroup';

    /**
     * Returns true if the given object is an instance of AgentFlowTransitionRouteGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentFlowTransitionRouteGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentFlowTransitionRouteGroup.__pulumiType;
    }

    /**
     * Required. The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Transition routes associated with the TransitionRouteGroup.
     */
    public readonly transitionRoutes!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3TransitionRouteResponse[]>;

    /**
     * Create a AgentFlowTransitionRouteGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentFlowTransitionRouteGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.flowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flowId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.transitionRouteGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitionRouteGroupId'");
            }
            inputs["agentId"] = args ? args.agentId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["flowId"] = args ? args.flowId : undefined;
            inputs["languageCode"] = args ? args.languageCode : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["transitionRouteGroupId"] = args ? args.transitionRouteGroupId : undefined;
            inputs["transitionRoutes"] = args ? args.transitionRoutes : undefined;
        } else {
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["transitionRoutes"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AgentFlowTransitionRouteGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentFlowTransitionRouteGroup resource.
 */
export interface AgentFlowTransitionRouteGroupArgs {
    readonly agentId: pulumi.Input<string>;
    /**
     * Required. The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
     */
    readonly displayName?: pulumi.Input<string>;
    readonly flowId: pulumi.Input<string>;
    readonly languageCode?: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    /**
     * The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    readonly transitionRouteGroupId: pulumi.Input<string>;
    /**
     * Transition routes associated with the TransitionRouteGroup.
     */
    readonly transitionRoutes?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3TransitionRouteArgs>[]>;
}
