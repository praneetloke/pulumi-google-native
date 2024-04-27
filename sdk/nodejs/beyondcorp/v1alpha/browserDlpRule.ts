// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new BrowserDlpRule in a given organization and PartnerTenant.
 * Auto-naming is currently not supported for this resource.
 */
export class BrowserDlpRule extends pulumi.CustomResource {
    /**
     * Get an existing BrowserDlpRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BrowserDlpRule {
        return new BrowserDlpRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:beyondcorp/v1alpha:BrowserDlpRule';

    /**
     * Returns true if the given object is an instance of BrowserDlpRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BrowserDlpRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BrowserDlpRule.__pulumiType;
    }

    /**
     * The group to which this Rule should be applied to.
     */
    public readonly group!: pulumi.Output<outputs.beyondcorp.v1alpha.GoogleCloudBeyondcorpPartnerservicesV1alphaGroupResponse>;
    /**
     * Unique resource name. The name is ignored when creating BrowserDlpRule.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly organizationId!: pulumi.Output<string>;
    public readonly partnerTenantId!: pulumi.Output<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The policy settings to apply.
     */
    public readonly ruleSetting!: pulumi.Output<outputs.beyondcorp.v1alpha.GoogleCloudBeyondcorpPartnerservicesV1alphaRuleSettingResponse>;

    /**
     * Create a BrowserDlpRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BrowserDlpRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.partnerTenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partnerTenantId'");
            }
            if ((!args || args.ruleSetting === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleSetting'");
            }
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["partnerTenantId"] = args ? args.partnerTenantId : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["ruleSetting"] = args ? args.ruleSetting : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["group"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["partnerTenantId"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["ruleSetting"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["organizationId", "partnerTenantId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(BrowserDlpRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BrowserDlpRule resource.
 */
export interface BrowserDlpRuleArgs {
    /**
     * The group to which this Rule should be applied to.
     */
    group: pulumi.Input<inputs.beyondcorp.v1alpha.GoogleCloudBeyondcorpPartnerservicesV1alphaGroupArgs>;
    organizationId: pulumi.Input<string>;
    partnerTenantId: pulumi.Input<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * The policy settings to apply.
     */
    ruleSetting: pulumi.Input<inputs.beyondcorp.v1alpha.GoogleCloudBeyondcorpPartnerservicesV1alphaRuleSettingArgs>;
}