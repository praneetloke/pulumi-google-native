// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a lake resource.
 * Auto-naming is currently not supported for this resource.
 */
export class Lake extends pulumi.CustomResource {
    /**
     * Get an existing Lake resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Lake {
        return new Lake(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataplex/v1:Lake';

    /**
     * Returns true if the given object is an instance of Lake.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lake {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lake.__pulumiType;
    }

    /**
     * Aggregated status of the underlying assets of the lake.
     */
    public /*out*/ readonly assetStatus!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1AssetStatusResponse>;
    /**
     * The time when the lake was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Description of the lake.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. User friendly display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. User-defined labels for the lake.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Required. Lake identifier. This ID will be used to generate names such as database and dataset names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique within the customer project / location.
     */
    public readonly lakeId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Optional. Settings to manage lake and Dataproc Metastore service instance association.
     */
    public readonly metastore!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1LakeMetastoreResponse>;
    /**
     * Metastore status of the lake.
     */
    public /*out*/ readonly metastoreStatus!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1LakeMetastoreStatusResponse>;
    /**
     * The relative resource name of the lake, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
     */
    public /*out*/ readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Current state of the lake.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the lake was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Lake resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LakeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.lakeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lakeId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["lakeId"] = args ? args.lakeId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["metastore"] = args ? args.metastore : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["assetStatus"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["metastoreStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["assetStatus"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["lakeId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["metastore"] = undefined /*out*/;
            resourceInputs["metastoreStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["lakeId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Lake.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Lake resource.
 */
export interface LakeArgs {
    /**
     * Optional. Description of the lake.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. User-defined labels for the lake.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Required. Lake identifier. This ID will be used to generate names such as database and dataset names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique within the customer project / location.
     */
    lakeId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Optional. Settings to manage lake and Dataproc Metastore service instance association.
     */
    metastore?: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1LakeMetastoreArgs>;
    project?: pulumi.Input<string>;
}
