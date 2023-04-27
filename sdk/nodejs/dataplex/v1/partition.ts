// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Create a metadata partition.
 * Auto-naming is currently not supported for this resource.
 */
export class Partition extends pulumi.CustomResource {
    /**
     * Get an existing Partition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Partition {
        return new Partition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataplex/v1:Partition';

    /**
     * Returns true if the given object is an instance of Partition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Partition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Partition.__pulumiType;
    }

    public readonly entityId!: pulumi.Output<string>;
    /**
     * Optional. The etag for this partition.
     */
    public readonly etag!: pulumi.Output<string>;
    public readonly lakeId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Partition values used in the HTTP URL must be double encoded. For example, url_encode(url_encode(value)) can be used to encode "US:CA/CA#Sunnyvale so that the request URL ends with "/partitions/US%253ACA/CA%2523Sunnyvale". The name field in the response retains the encoded format.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Immutable. The set of values representing the partition, which correspond to the partition schema defined in the parent entity.
     */
    public readonly values!: pulumi.Output<string[]>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Partition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PartitionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.entityId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityId'");
            }
            if ((!args || args.lakeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lakeId'");
            }
            if ((!args || args.values === undefined) && !opts.urn) {
                throw new Error("Missing required property 'values'");
            }
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["lakeId"] = args ? args.lakeId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["values"] = args ? args.values : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["entityId"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["lakeId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["values"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["entityId", "lakeId", "location", "project", "zone"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Partition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Partition resource.
 */
export interface PartitionArgs {
    entityId: pulumi.Input<string>;
    /**
     * Optional. The etag for this partition.
     */
    etag?: pulumi.Input<string>;
    lakeId: pulumi.Input<string>;
    /**
     * Immutable. The location of the entity data within the partition, for example, gs://bucket/path/to/entity/key1=value1/key2=value2. Or projects//datasets//tables/
     */
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Immutable. The set of values representing the partition, which correspond to the partition schema defined in the parent entity.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
    zone?: pulumi.Input<string>;
}
