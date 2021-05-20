// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new game server deployment in a given project and location.
 */
export class GameServerDeployment extends pulumi.CustomResource {
    /**
     * Get an existing GameServerDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GameServerDeployment {
        return new GameServerDeployment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gameservices/v1:GameServerDeployment';

    /**
     * Returns true if the given object is an instance of GameServerDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GameServerDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GameServerDeployment.__pulumiType;
    }

    /**
     * The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Human readable description of the game server delpoyment.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * ETag of the resource.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The labels associated with this game server deployment. Each label is a key-value pair.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name of the game server deployment, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-deployment`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The last-modified time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a GameServerDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GameServerDeploymentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deploymentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentId'");
            }
            if ((!args || args.gameServerDeploymentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gameServerDeploymentId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["deploymentId"] = args ? args.deploymentId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["gameServerDeploymentId"] = args ? args.gameServerDeploymentId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GameServerDeployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GameServerDeployment resource.
 */
export interface GameServerDeploymentArgs {
    readonly deploymentId: pulumi.Input<string>;
    /**
     * Human readable description of the game server delpoyment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * ETag of the resource.
     */
    readonly etag?: pulumi.Input<string>;
    readonly gameServerDeploymentId: pulumi.Input<string>;
    /**
     * The labels associated with this game server deployment. Each label is a key-value pair.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly location: pulumi.Input<string>;
    /**
     * The resource name of the game server deployment, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-deployment`.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
}
