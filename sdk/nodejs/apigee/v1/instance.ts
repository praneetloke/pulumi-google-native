// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an Apigee runtime instance. The instance is accessible from the authorized network configured on the organization. **Note:** Not supported for Apigee hybrid.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Time the instance was created in milliseconds since epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Optional. Description of the instance.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
     */
    public readonly diskEncryptionKeyName!: pulumi.Output<string>;
    /**
     * Optional. Display name for the instance.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Time the instance was last modified in milliseconds since epoch.
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * Compute Engine location where the instance resides.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
     */
    public readonly peeringCidrRange!: pulumi.Output<string>;
    /**
     * Port number of the exposed Apigee endpoint.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["diskEncryptionKeyName"] = args ? args.diskEncryptionKeyName : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["peeringCidrRange"] = args ? args.peeringCidrRange : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["host"] = undefined /*out*/;
            inputs["lastModifiedAt"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        } else {
            inputs["createdAt"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["diskEncryptionKeyName"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["host"] = undefined /*out*/;
            inputs["lastModifiedAt"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["peeringCidrRange"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Optional. Description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
     */
    diskEncryptionKeyName?: pulumi.Input<string>;
    /**
     * Optional. Display name for the instance.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Compute Engine location where the instance resides.
     */
    location?: pulumi.Input<string>;
    /**
     * Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
     */
    peeringCidrRange?: pulumi.Input<enums.apigee.v1.InstancePeeringCidrRange>;
}
