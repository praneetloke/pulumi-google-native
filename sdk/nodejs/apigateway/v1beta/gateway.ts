// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new Gateway in a given project and location.
 */
export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigateway/v1beta:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * Required. Resource name of the API Config for this Gateway. Format: projects/{project}/locations/global/apis/{api}/configs/{apiConfig}
     */
    public readonly apiConfig!: pulumi.Output<string>;
    /**
     * Created time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The default API Gateway host name of the form `{gateway_id}-{hash}.{region_code}.gateway.dev`.
     */
    public /*out*/ readonly defaultHostname!: pulumi.Output<string>;
    /**
     * Optional. Display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource name of the Gateway. Format: projects/{project}/locations/{location}/gateways/{gateway}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current state of the Gateway.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Updated time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["apiConfig"] = args ? args.apiConfig : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["gatewayId"] = args ? args.gatewayId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["defaultHostname"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["apiConfig"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["defaultHostname"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Gateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * Required. Resource name of the API Config for this Gateway. Format: projects/{project}/locations/global/apis/{api}/configs/{apiConfig}
     */
    readonly apiConfig?: pulumi.Input<string>;
    /**
     * Optional. Display name.
     */
    readonly displayName?: pulumi.Input<string>;
    readonly gatewayId: pulumi.Input<string>;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly location: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
}
