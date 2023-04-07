// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Asynchronously creates linked dataset in BigQuery which makes it possible to use BugQuery to read the logs stored in the bucket. A bucket may currently only contain one link.
 */
export class Link extends pulumi.CustomResource {
    /**
     * Get an existing Link resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Link {
        return new Link(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:logging/v2:Link';

    /**
     * Returns true if the given object is an instance of Link.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Link {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Link.__pulumiType;
    }

    /**
     * The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
     */
    public readonly bigqueryDataset!: pulumi.Output<outputs.logging.v2.BigQueryDatasetResponse>;
    public readonly bucketId!: pulumi.Output<string>;
    /**
     * The creation timestamp of the link.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Describes this link.The maximum length of the description is 8000 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The resource lifecycle state.
     */
    public /*out*/ readonly lifecycleState!: pulumi.Output<string>;
    /**
     * Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
     */
    public readonly linkId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Link resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucketId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketId'");
            }
            if ((!args || args.linkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'linkId'");
            }
            resourceInputs["bigqueryDataset"] = args ? args.bigqueryDataset : undefined;
            resourceInputs["bucketId"] = args ? args.bucketId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["linkId"] = args ? args.linkId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["lifecycleState"] = undefined /*out*/;
        } else {
            resourceInputs["bigqueryDataset"] = undefined /*out*/;
            resourceInputs["bucketId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["lifecycleState"] = undefined /*out*/;
            resourceInputs["linkId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["bucketId", "linkId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Link.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Link resource.
 */
export interface LinkArgs {
    /**
     * The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
     */
    bigqueryDataset?: pulumi.Input<inputs.logging.v2.BigQueryDatasetArgs>;
    bucketId: pulumi.Input<string>;
    /**
     * Describes this link.The maximum length of the description is 8000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
     */
    linkId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
