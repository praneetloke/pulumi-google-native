// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a CustomerRepricingConfig. Call this method to set modifications for a specific customer's bill. You can only create configs if the RepricingConfig.effective_invoice_month is a future month. If needed, you can create a config for the current month, with some restrictions. When creating a config for a future month, make sure there are no existing configs for that RepricingConfig.effective_invoice_month. The following restrictions are for creating configs in the current month. * This functionality is reserved for recovering from an erroneous config, and should not be used for regular business cases. * The new config will not modify exports used with other configs. Changes to the config may be immediate, but may take up to 24 hours. * There is a limit of ten configs for any RepricingConfig.EntitlementGranularity.entitlement or RepricingConfig.effective_invoice_month. * The contained CustomerRepricingConfig.repricing_config vaule must be different from the value used in the current config for a RepricingConfig.EntitlementGranularity.entitlement. Possible Error Codes: * PERMISSION_DENIED: If the account making the request and the account being queried are different. * INVALID_ARGUMENT: Missing or invalid required parameters in the request. Also displays if the updated config is for the current month or past months. * NOT_FOUND: The CustomerRepricingConfig specified does not exist or is not associated with the given account. * INTERNAL: Any non-user error related to technical issues in the backend. In this case, contact Cloud Channel support. Return Value: If successful, the updated CustomerRepricingConfig resource, otherwise returns an error.
 * Auto-naming is currently not supported for this resource.
 */
export class CustomerRepricingConfig extends pulumi.CustomResource {
    /**
     * Get an existing CustomerRepricingConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomerRepricingConfig {
        return new CustomerRepricingConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudchannel/v1:CustomerRepricingConfig';

    /**
     * Returns true if the given object is an instance of CustomerRepricingConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomerRepricingConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomerRepricingConfig.__pulumiType;
    }

    public readonly accountId!: pulumi.Output<string>;
    public readonly customerId!: pulumi.Output<string>;
    /**
     * Resource name of the CustomerRepricingConfig. Format: accounts/{account_id}/customers/{customer_id}/customerRepricingConfigs/{id}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The configuration for bill modifications made by a reseller before sending it to customers.
     */
    public readonly repricingConfig!: pulumi.Output<outputs.cloudchannel.v1.GoogleCloudChannelV1RepricingConfigResponse>;
    /**
     * Timestamp of an update to the repricing rule. If `update_time` is after RepricingConfig.effective_invoice_month then it indicates this was set mid-month.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a CustomerRepricingConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomerRepricingConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            if ((!args || args.customerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerId'");
            }
            if ((!args || args.repricingConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repricingConfig'");
            }
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["customerId"] = args ? args.customerId : undefined;
            resourceInputs["repricingConfig"] = args ? args.repricingConfig : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["customerId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["repricingConfig"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["accountId", "customerId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(CustomerRepricingConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomerRepricingConfig resource.
 */
export interface CustomerRepricingConfigArgs {
    accountId: pulumi.Input<string>;
    customerId: pulumi.Input<string>;
    /**
     * The configuration for bill modifications made by a reseller before sending it to customers.
     */
    repricingConfig: pulumi.Input<inputs.cloudchannel.v1.GoogleCloudChannelV1RepricingConfigArgs>;
}