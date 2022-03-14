// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new game server cluster in a given project and location.
 */
export class GameServerCluster extends pulumi.CustomResource {
    /**
     * Get an existing GameServerCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GameServerCluster {
        return new GameServerCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gameservices/v1beta:GameServerCluster';

    /**
     * Returns true if the given object is an instance of GameServerCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GameServerCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GameServerCluster.__pulumiType;
    }

    /**
     * The state of the Kubernetes cluster in preview. This will be available if view is set to FULL in the relevant list/get/preview request.
     */
    public /*out*/ readonly clusterState!: pulumi.Output<outputs.gameservices.v1beta.KubernetesClusterStateResponse>;
    /**
     * The game server cluster connection information. This information is used to manage game server clusters.
     */
    public readonly connectionInfo!: pulumi.Output<outputs.gameservices.v1beta.GameServerClusterConnectionInfoResponse>;
    /**
     * The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Human readable description of the cluster.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The labels associated with this game server cluster. Each label is a key-value pair.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name of the game server cluster, in the following form: `projects/{project}/locations/{locationId}/realms/{realmId}/gameServerClusters/{gameServerClusterId}`. For example, `projects/my-project/locations/global/realms/zanzibar/gameServerClusters/my-gke-cluster`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The last-modified time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a GameServerCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GameServerClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.gameServerClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gameServerClusterId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["connectionInfo"] = args ? args.connectionInfo : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["gameServerClusterId"] = args ? args.gameServerClusterId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["clusterState"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["clusterState"] = undefined /*out*/;
            resourceInputs["connectionInfo"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GameServerCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GameServerCluster resource.
 */
export interface GameServerClusterArgs {
    /**
     * The game server cluster connection information. This information is used to manage game server clusters.
     */
    connectionInfo?: pulumi.Input<inputs.gameservices.v1beta.GameServerClusterConnectionInfoArgs>;
    /**
     * Human readable description of the cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
     */
    etag?: pulumi.Input<string>;
    /**
     * Required. The ID of the game server cluster resource to create.
     */
    gameServerClusterId: pulumi.Input<string>;
    /**
     * The labels associated with this game server cluster. Each label is a key-value pair.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the game server cluster, in the following form: `projects/{project}/locations/{locationId}/realms/{realmId}/gameServerClusters/{gameServerClusterId}`. For example, `projects/my-project/locations/global/realms/zanzibar/gameServerClusters/my-gke-cluster`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    realmId: pulumi.Input<string>;
}
