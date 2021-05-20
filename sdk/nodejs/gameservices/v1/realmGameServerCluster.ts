// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new game server cluster in a given project and location.
 */
export class RealmGameServerCluster extends pulumi.CustomResource {
    /**
     * Get an existing RealmGameServerCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RealmGameServerCluster {
        return new RealmGameServerCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gameservices/v1:RealmGameServerCluster';

    /**
     * Returns true if the given object is an instance of RealmGameServerCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RealmGameServerCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RealmGameServerCluster.__pulumiType;
    }

    /**
     * The game server cluster connection information. This information is used to manage game server clusters.
     */
    public readonly connectionInfo!: pulumi.Output<outputs.gameservices.v1.GameServerClusterConnectionInfoResponse>;
    /**
     * The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Human readable description of the cluster.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * ETag of the resource.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The labels associated with this game server cluster. Each label is a key-value pair.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The last-modified time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a RealmGameServerCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealmGameServerClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.gameServerClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gameServerClusterId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["connectionInfo"] = args ? args.connectionInfo : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["gameServerClusterId"] = args ? args.gameServerClusterId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["connectionInfo"] = undefined /*out*/;
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
        super(RealmGameServerCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RealmGameServerCluster resource.
 */
export interface RealmGameServerClusterArgs {
    /**
     * The game server cluster connection information. This information is used to manage game server clusters.
     */
    readonly connectionInfo?: pulumi.Input<inputs.gameservices.v1.GameServerClusterConnectionInfoArgs>;
    /**
     * Human readable description of the cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * ETag of the resource.
     */
    readonly etag?: pulumi.Input<string>;
    readonly gameServerClusterId: pulumi.Input<string>;
    /**
     * The labels associated with this game server cluster. Each label is a key-value pair.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly location: pulumi.Input<string>;
    /**
     * Required. The resource name of the game server cluster, in the following form: `projects/{project}/locations/{location}/realms/{realm}/gameServerClusters/{cluster}`. For example, `projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster`.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
}
