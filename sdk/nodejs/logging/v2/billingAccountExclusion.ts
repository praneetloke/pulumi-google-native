// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new exclusion in a specified parent resource. Only log entries belonging to that resource can be excluded. You can have up to 10 exclusions in a resource.
 */
export class BillingAccountExclusion extends pulumi.CustomResource {
    /**
     * Get an existing BillingAccountExclusion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BillingAccountExclusion {
        return new BillingAccountExclusion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:logging/v2:BillingAccountExclusion';

    /**
     * Returns true if the given object is an instance of BillingAccountExclusion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BillingAccountExclusion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BillingAccountExclusion.__pulumiType;
    }

    /**
     * The creation timestamp of the exclusion.This field may not be present for older exclusions.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. A description of this exclusion.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * Required. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries. For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:"resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)"
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * Required. A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The last update timestamp of the exclusion.This field may not be present for older exclusions.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a BillingAccountExclusion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BillingAccountExclusionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.billingAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'billingAccountId'");
            }
            if ((!args || args.exclusionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exclusionId'");
            }
            inputs["billingAccountId"] = args ? args.billingAccountId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["exclusionId"] = args ? args.exclusionId : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["disabled"] = undefined /*out*/;
            inputs["filter"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BillingAccountExclusion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BillingAccountExclusion resource.
 */
export interface BillingAccountExclusionArgs {
    readonly billingAccountId: pulumi.Input<string>;
    /**
     * Optional. A description of this exclusion.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
     */
    readonly disabled?: pulumi.Input<boolean>;
    readonly exclusionId: pulumi.Input<string>;
    /**
     * Required. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries. For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:"resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)"
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * Required. A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
     */
    readonly name?: pulumi.Input<string>;
}
