// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new private connection in a given project and location.
 */
export class PrivateConnection extends pulumi.CustomResource {
    /**
     * Get an existing PrivateConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateConnection {
        return new PrivateConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datamigration/v1:PrivateConnection';

    /**
     * Returns true if the given object is an instance of PrivateConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateConnection.__pulumiType;
    }

    /**
     * The create time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The private connection display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The error details in case of state FAILED.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.datamigration.v1.StatusResponse>;
    /**
     * The resource labels for private connections to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Required. The private connection identifier.
     */
    public readonly privateConnectionId!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. A unique ID used to identify the request. If the server receives two requests with the same ID, then the second request is ignored. It is recommended to always set this value to a UUID. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * Optional. If set to true, will skip validations.
     */
    public readonly skipValidation!: pulumi.Output<boolean | undefined>;
    /**
     * The state of the private connection.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The last update time of the resource.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * VPC peering configuration.
     */
    public readonly vpcPeeringConfig!: pulumi.Output<outputs.datamigration.v1.VpcPeeringConfigResponse>;

    /**
     * Create a PrivateConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.privateConnectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateConnectionId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateConnectionId"] = args ? args.privateConnectionId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["skipValidation"] = args ? args.skipValidation : undefined;
            resourceInputs["vpcPeeringConfig"] = args ? args.vpcPeeringConfig : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["privateConnectionId"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["skipValidation"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vpcPeeringConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "privateConnectionId", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PrivateConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateConnection resource.
 */
export interface PrivateConnectionArgs {
    /**
     * The private connection display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The resource labels for private connections to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Required. The private connection identifier.
     */
    privateConnectionId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. A unique ID used to identify the request. If the server receives two requests with the same ID, then the second request is ignored. It is recommended to always set this value to a UUID. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    requestId?: pulumi.Input<string>;
    /**
     * Optional. If set to true, will skip validations.
     */
    skipValidation?: pulumi.Input<boolean>;
    /**
     * VPC peering configuration.
     */
    vpcPeeringConfig?: pulumi.Input<inputs.datamigration.v1.VpcPeeringConfigArgs>;
}
