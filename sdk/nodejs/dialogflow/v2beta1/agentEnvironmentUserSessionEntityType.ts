// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a session entity type. If the specified session entity type already exists, overrides the session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.
 */
export class AgentEnvironmentUserSessionEntityType extends pulumi.CustomResource {
    /**
     * Get an existing AgentEnvironmentUserSessionEntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentEnvironmentUserSessionEntityType {
        return new AgentEnvironmentUserSessionEntityType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v2beta1:AgentEnvironmentUserSessionEntityType';

    /**
     * Returns true if the given object is an instance of AgentEnvironmentUserSessionEntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentEnvironmentUserSessionEntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentEnvironmentUserSessionEntityType.__pulumiType;
    }

    /**
     * Required. The collection of entities associated with this session entity type.
     */
    public readonly entities!: pulumi.Output<outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1EntityTypeEntityResponse[]>;
    /**
     * Required. Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    public readonly entityOverrideMode!: pulumi.Output<string>;
    /**
     * Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a AgentEnvironmentUserSessionEntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentEnvironmentUserSessionEntityTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.entityTypeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityTypeId'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.sessionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionId'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            inputs["entities"] = args ? args.entities : undefined;
            inputs["entityOverrideMode"] = args ? args.entityOverrideMode : undefined;
            inputs["entityTypeId"] = args ? args.entityTypeId : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sessionId"] = args ? args.sessionId : undefined;
            inputs["userId"] = args ? args.userId : undefined;
        } else {
            inputs["entities"] = undefined /*out*/;
            inputs["entityOverrideMode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AgentEnvironmentUserSessionEntityType.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentEnvironmentUserSessionEntityType resource.
 */
export interface AgentEnvironmentUserSessionEntityTypeArgs {
    /**
     * Required. The collection of entities associated with this session entity type.
     */
    readonly entities?: pulumi.Input<pulumi.Input<inputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1EntityTypeEntityArgs>[]>;
    /**
     * Required. Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    readonly entityOverrideMode?: pulumi.Input<string>;
    readonly entityTypeId: pulumi.Input<string>;
    readonly environmentId: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    /**
     * Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    readonly sessionId: pulumi.Input<string>;
    readonly userId: pulumi.Input<string>;
}
